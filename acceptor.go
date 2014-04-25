package quickfix

import (
	"errors"
	"github.com/quickfixgo/quickfix/log"
	"github.com/quickfixgo/quickfix/settings"
	"net"
	"strconv"
)

//Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app        Application
	settings   settings.ApplicationSettings
	logFactory log.LogFactory
	globalLog  log.Log
}

//Start accepting connections.
func (a *Acceptor) Start() (e error) {
	port, hasPort := a.settings.GetGlobalSettings().GetInt(settings.SocketAcceptPort)
	if !hasPort {
		return errors.New("config error: must provide SocketAcceptPort")
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
func NewAcceptor(app Application, settings settings.ApplicationSettings, logFactory log.LogFactory) (*Acceptor, error) {
	a := new(Acceptor)
	a.app = app
	a.settings = settings
	a.logFactory = logFactory
	a.globalLog = logFactory.Create()

	for _, s := range settings.GetSessionSettings() {
		if _, err := createSession(s, logFactory, app); err != nil {
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
