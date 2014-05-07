package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"net"
	"strconv"
)

//Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app        Application
	settings   *Settings
	logFactory LogFactory
	globalLog  Log
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
		for {
			cxn := <-connections
			go handleAcceptorConnection(cxn, a.globalLog)
		}
	}()

	return
}

//Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {}

//NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, settings *Settings, logFactory LogFactory) (*Acceptor, error) {
	a := new(Acceptor)
	a.app = app
	a.settings = settings
	a.logFactory = logFactory
	a.globalLog = logFactory.Create()

	for sessionID, sessionSettings := range settings.SessionSettings() {
		if err := createSession(sessionID, sessionSettings, logFactory, app); err != nil {
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
