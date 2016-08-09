package quickfix

import (
	"crypto/tls"
	"fmt"
	"net"
	"strconv"

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
	quitChan            chan bool
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

	go func() {
		for cxn := range connections {
			go handleAcceptorConnection(cxn, a.qualifiedSessionIDs, a.globalLog)
		}
	}()

	return nil
}

//Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()
	close(a.quitChan)
}

//NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (*Acceptor, error) {
	a := &Acceptor{
		quitChan:            make(chan bool),
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

		if err = createSession(sessionID, storeFactory, sessionSettings, logFactory, app, a.quitChan); err != nil {
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
			if netConn == nil {
				a.globalLog.OnEventf("Couldn't Accept: %v", err.Error())
				continue
			}

			ch <- netConn
		}
	}()

	return ch
}
