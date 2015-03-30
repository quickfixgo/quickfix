package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"net"
	"strconv"
)

//Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app                 Application
	settings            *Settings
	logFactory          LogFactory
	storeFactory        MessageStoreFactory
	globalLog           Log
	qualifiedSessionIDs map[SessionID]SessionID
	quitChans           map[int]chan bool
	doneChannel         chan int
}

//Start accepting connections.
func (a *Acceptor) Start() (e error) {
	port, err := a.settings.GlobalSettings().IntSetting(config.SocketAcceptPort)
	if err != nil {
		return fmt.Errorf("error fetching required SocketAcceptPort: %v", err)
	}

	server, err := net.Listen("tcp", ":"+strconv.Itoa(port))
	if server == nil {
		return err
	}

	connections := a.listenForConnections(server)

	go func() {
		var i = 0
		for {
			select {
			case cxn := <-connections:
				a.quitChans[i] = make(chan bool)
				go handleAcceptorConnection(cxn, a.qualifiedSessionIDs, a.globalLog, a.quitChans[i], a.doneChannel, i)
				i++
			case j := <-a.doneChannel:
				delete(a.quitChans, j)
			}
		}
	}()

	return
}

//Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()
	for _, channel := range a.quitChans {
		channel <- true
	}
}

//NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (*Acceptor, error) {
	a := new(Acceptor)
	a.app = app
	a.storeFactory = storeFactory
	a.settings = settings
	a.logFactory = logFactory
	a.qualifiedSessionIDs = make(map[SessionID]SessionID)
	a.quitChans = make(map[int]chan bool)
	a.doneChannel = make(chan int, 10)

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
			if netConn == nil {
				a.globalLog.OnEventf("Couldn't Accept: %v", err.Error())
				continue
			}

			ch <- netConn
		}
	}()

	return ch
}
