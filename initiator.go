package quickfix

import (
	"bufio"
	"crypto/tls"
	"net"
	"sync"
	"time"
)

//Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        *Settings
	sessionSettings map[SessionID]*SessionSettings
	storeFactory    MessageStoreFactory
	logFactory      LogFactory
	globalLog       Log
	stopChan        chan interface{}
	wg              sync.WaitGroup
	sessions        map[SessionID]*session
	sessionFactory
}

//Start Initiator.
func (i *Initiator) Start() (err error) {
	i.stopChan = make(chan interface{})

	for sessionID, settings := range i.sessionSettings {
		//TODO: move into session factory
		var tlsConfig *tls.Config
		if tlsConfig, err = loadTLSConfig(settings); err != nil {
			return
		}

		i.wg.Add(1)
		go func(sessID SessionID) {
			i.handleConnection(i.sessions[sessID], tlsConfig)
			i.wg.Done()
		}(sessionID)
	}

	return
}

//Stop Initiator.
func (i *Initiator) Stop() {
	close(i.stopChan)
	i.wg.Wait()
}

//NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, storeFactory MessageStoreFactory, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	i := &Initiator{
		app:             app,
		storeFactory:    storeFactory,
		settings:        appSettings,
		sessionSettings: appSettings.SessionSettings(),
		logFactory:      logFactory,
		sessions:        make(map[SessionID]*session),
		sessionFactory:  sessionFactory{true},
	}

	var err error
	i.globalLog, err = logFactory.Create()
	if err != nil {
		return i, err
	}

	for sessionID, s := range i.sessionSettings {
		session, err := i.createSession(sessionID, storeFactory, s, logFactory, app)
		if err != nil {
			return nil, err
		}

		i.sessions[sessionID] = session
	}

	return i, nil
}

//waitForInSessionTime returns true if the session is in session, false if the handler should stop
func (i *Initiator) waitForInSessionTime(session *session) bool {
	inSessionTime := make(chan interface{})
	go func() {
		session.waitForInSessionTime()
		close(inSessionTime)
	}()

	select {
	case <-inSessionTime:
	case <-i.stopChan:
		return false
	}

	return true
}

//watiForReconnectInterval returns true if a reconnect should be re-attempted, false if handler should stop
func (i *Initiator) waitForReconnectInterval(reconnectInterval time.Duration) bool {
	select {
	case <-time.After(reconnectInterval):
	case <-i.stopChan:
		return false
	}

	return true
}

func (i *Initiator) handleConnection(session *session, tlsConfig *tls.Config) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		session.run()
		wg.Done()
	}()

	defer func() {
		session.stop()
		wg.Wait()
	}()

	connectionAttempt := 0

	for {
		if !i.waitForInSessionTime(session) {
			return
		}

		var disconnected chan interface{}
		var msgIn chan fixIn
		var msgOut chan []byte

		address := session.SocketConnectAddress[connectionAttempt%len(session.SocketConnectAddress)]
		session.log.OnEventf("Connecting to: %v", address)

		var netConn net.Conn
		if tlsConfig != nil {
			tlsConn, err := tls.Dial("tcp", address, tlsConfig)
			if err != nil {
				session.log.OnEventf("Failed to connect: %v", err)
				goto reconnect
			}

			err = tlsConn.Handshake()
			if err != nil {
				session.log.OnEventf("Failed handshake:%v", err)
				goto reconnect
			}
			netConn = tlsConn
		} else {
			var err error
			netConn, err = net.Dial("tcp", address)
			if err != nil {
				session.log.OnEventf("Failed to connect: %v", err)
				goto reconnect
			}
		}

		msgIn = make(chan fixIn)
		msgOut = make(chan []byte)
		if err := session.connect(msgIn, msgOut); err != nil {
			session.log.OnEventf("Failed to initiate: %v", err)
			goto reconnect
		}

		go readLoop(newParser(bufio.NewReader(netConn)), msgIn)
		disconnected = make(chan interface{})
		go func() {
			writeLoop(netConn, msgOut, session.log)
			if err := netConn.Close(); err != nil {
				session.log.OnEvent(err.Error())
			}
			close(disconnected)
		}()

		select {
		case <-disconnected:
		case <-i.stopChan:
			return
		}

	reconnect:
		connectionAttempt++
		session.log.OnEventf("Reconnecting in %v", session.ReconnectInterval)
		if !i.waitForReconnectInterval(session.ReconnectInterval) {
			return
		}
	}
}
