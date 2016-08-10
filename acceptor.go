package quickfix

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/quickfixgo/quickfix/config"
)

//Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app                 Application
	settings            *Settings
	logFactory          LogFactory
	storeFactory        MessageStoreFactory
	globalLog           Log
	qualifiedSessionIDs map[SessionID]SessionID
	stopChan            chan interface{}
	wg                  sync.WaitGroup
}

//Start accepting connections.
func (a *Acceptor) Start() error {
	port, err := a.settings.GlobalSettings().IntSetting(config.SocketAcceptPort)
	if err != nil {
		return fmt.Errorf("error fetching required SocketAcceptPort: %v", err)
	}

	var tlsConfig *tls.Config
	if tlsConfig, err = loadTLSConfig(a.settings); err != nil {
		return err
	}

	var server net.Listener
	address := ":" + strconv.Itoa(port)
	if tlsConfig != nil {
		if server, err = tls.Listen("tcp", address, tlsConfig); err != nil {
			return err
		}
	} else {
		if server, err = net.Listen("tcp", address); err != nil {
			return err
		}
	}

	connections := a.listenForConnections(server)
	a.wg.Add(1)
	go a.run(server, connections)

	return nil
}

//Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()

	a.stopChan <- true
	a.wg.Wait()
}

//NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (*Acceptor, error) {
	a := &Acceptor{
		stopChan:            make(chan interface{}),
		app:                 app,
		storeFactory:        storeFactory,
		settings:            settings,
		logFactory:          logFactory,
		qualifiedSessionIDs: make(map[SessionID]SessionID),
	}

	var err error
	a.globalLog, err = logFactory.Create()
	if err != nil {
		return a, err
	}

	for sessionID, sessionSettings := range settings.SessionSettings() {
		//unqualified sessionIDs must be unique
		unqualifiedSessionID := SessionID{
			BeginString:  sessionID.BeginString,
			TargetCompID: sessionID.TargetCompID,
			SenderCompID: sessionID.SenderCompID}

		if _, dup := a.qualifiedSessionIDs[unqualifiedSessionID]; dup {
			return nil, fmt.Errorf("duplicate SessionID %v", unqualifiedSessionID)
		}
		a.qualifiedSessionIDs[unqualifiedSessionID] = sessionID

		if err = createSession(sessionID, storeFactory, sessionSettings, logFactory, app); err != nil {
			return nil, err
		}
	}

	return a, nil
}

func (a *Acceptor) listenForConnections(listener net.Listener) (ch chan net.Conn) {
	ch = make(chan net.Conn)

	go func() {
		for {
			netConn, err := listener.Accept()
			if err != nil {
				close(ch)
				return
			}

			ch <- netConn
		}
	}()

	return ch
}

func (a *Acceptor) run(server net.Listener, connections chan net.Conn) {
	defer func() {
		a.wg.Done()
	}()

	var connGroup sync.WaitGroup
	var pendingStop bool

	for {
		select {

		case <-a.stopChan:
			if err := server.Close(); err != nil {
				a.globalLog.OnEvent(err.Error())
			}

			for sessionID := range a.qualifiedSessionIDs {
				session, err := lookupSession(sessionID)
				if err != nil {
					a.globalLog.OnEventf("Error getting session: %v", err)
				} else {
					go session.disconnect()
				}
			}
			pendingStop = true

		case cxn, ok := <-connections:
			if !ok {
				connGroup.Wait()
				return
			}

			if pendingStop {
				if err := cxn.Close(); err != nil {
					a.globalLog.OnEvent(err.Error())
				}
				continue
			}

			connGroup.Add(1)
			go func() {
				handleAcceptorConnection(cxn, a.qualifiedSessionIDs, a.globalLog)
				connGroup.Done()
			}()
		}
	}
}
