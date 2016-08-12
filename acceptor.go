package quickfix

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"sync"

	"github.com/quickfixgo/quickfix/config"
)

//Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app              Application
	settings         *Settings
	logFactory       LogFactory
	storeFactory     MessageStoreFactory
	globalLog        Log
	sessions         map[SessionID]*session
	sessionGroup     sync.WaitGroup
	listener         net.Listener
	listenerShutdown sync.WaitGroup
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

	address := ":" + strconv.Itoa(port)
	if tlsConfig != nil {
		if a.listener, err = tls.Listen("tcp", address, tlsConfig); err != nil {
			return err
		}
	} else {
		if a.listener, err = net.Listen("tcp", address); err != nil {
			return err
		}
	}

	for sessionID := range a.sessions {
		session := a.sessions[sessionID]
		a.sessionGroup.Add(1)
		go func() {
			session.run()
			a.sessionGroup.Done()
		}()
	}

	a.listenerShutdown.Add(1)
	go a.listenForConnections()
	return nil
}

//Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()

	a.listener.Close()
	a.listenerShutdown.Wait()
	for _, session := range a.sessions {
		session.stop()
	}
	a.sessionGroup.Wait()
}

//NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (*Acceptor, error) {
	a := &Acceptor{
		app:          app,
		storeFactory: storeFactory,
		settings:     settings,
		logFactory:   logFactory,
		sessions:     make(map[SessionID]*session),
	}

	var err error
	if a.globalLog, err = logFactory.Create(); err != nil {
		return a, err
	}

	for sessionID, sessionSettings := range settings.SessionSettings() {
		sessID := SessionID{
			BeginString:  sessionID.BeginString,
			TargetCompID: sessionID.TargetCompID,
			SenderCompID: sessionID.SenderCompID,
		}

		if _, dup := a.sessions[sessID]; dup {
			return a, errDuplicateSessionID
		}

		if a.sessions[sessID], err = createSession(sessionID, storeFactory, sessionSettings, logFactory, app); err != nil {
			return a, err
		}
	}

	return a, nil
}

func (a *Acceptor) listenForConnections() {
	defer a.listenerShutdown.Done()

	for {
		netConn, err := a.listener.Accept()
		if err != nil {
			return
		}

		go func() {
			a.handleConnection(netConn)
		}()
	}
}

func (a *Acceptor) invalidMessage(msg []byte, err error) {
	a.globalLog.OnEventf("Invalid Message: %s, %v", msg, err.Error())
}

func (a *Acceptor) handleConnection(netConn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			a.globalLog.OnEventf("Connection Terminated: %v", err)
		}

		if err := netConn.Close(); err != nil {
			a.globalLog.OnEvent(err.Error())
		}
	}()

	reader := bufio.NewReader(netConn)
	parser := newParser(reader)

	msgBytes, err := parser.ReadMessage()
	if err != nil {
		a.globalLog.OnEvent(err.Error())
		return
	}

	msg, err := ParseMessage(msgBytes)
	if err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var beginString FIXString
	if err := msg.Header.GetField(tagBeginString, &beginString); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var senderCompID FIXString
	if err := msg.Header.GetField(tagSenderCompID, &senderCompID); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var targetCompID FIXString
	if err := msg.Header.GetField(tagTargetCompID, &targetCompID); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	sessID := SessionID{BeginString: string(beginString), SenderCompID: string(targetCompID), TargetCompID: string(senderCompID)}
	session, ok := a.sessions[sessID]
	if !ok {
		a.globalLog.OnEventf("Session %v not found for incoming message: %s", sessID, msgBytes)
		return
	}

	msgIn := make(chan fixIn)
	msgOut := make(chan []byte)

	if err := session.accept(msgIn, msgOut); err != nil {
		a.globalLog.OnEventf("Unable to accept %v", err.Error())
		return
	}

	go func() {
		msgIn <- fixIn{msgBytes, parser.lastRead}
		readLoop(parser, msgIn)
	}()

	writeLoop(netConn, msgOut, a.globalLog)
}
