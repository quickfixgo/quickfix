package quickfix

import (
	"bufio"
	"crypto/tls"
	"fmt"
	"net"
	"strconv"
	"sync"
	"time"

	"github.com/quickfixgo/quickfix/config"
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
}

//Start Initiator.
func (i *Initiator) Start() error {
	i.stopChan = make(chan interface{})

	for sessionID, s := range i.sessionSettings {
		socketConnectHost, err := s.Setting(config.SocketConnectHost)
		if err != nil {
			return fmt.Errorf("error on SocketConnectHost: %v", err)
		}

		socketConnectPort, err := s.IntSetting(config.SocketConnectPort)
		if err != nil {
			return fmt.Errorf("error on SocketConnectPort: %v", err)
		}

		reconnectInterval := 30 // Default configuration (in seconds)
		if s.HasSetting(config.ReconnectInterval) {
			if reconnectInterval, err = s.IntSetting(config.ReconnectInterval); err != nil {
				return fmt.Errorf("error on ReconnectInterval: %v", err)
			} else if reconnectInterval <= 0 {
				return fmt.Errorf("ReconnectInterval must be greater than zero")
			}
		}

		var tlsConfig *tls.Config
		if tlsConfig, err = loadTLSConfig(i.settings); err != nil {
			return err
		}
		address := net.JoinHostPort(socketConnectHost, strconv.Itoa(socketConnectPort))

		i.wg.Add(1)
		go func(sessID SessionID) {
			i.handleConnection(address, i.sessions[sessID], time.Duration(reconnectInterval)*time.Second, tlsConfig)
			i.wg.Done()
		}(sessionID)
	}

	return nil
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
	}

	var err error
	i.globalLog, err = logFactory.Create()
	if err != nil {
		return i, err
	}

	for sessionID, s := range i.sessionSettings {

		//fail fast
		if ok := s.HasSetting(config.SocketConnectHost); !ok {
			return nil, requiredConfigurationMissing(config.SocketConnectHost)
		}

		if ok := s.HasSetting(config.SocketConnectPort); !ok {
			return nil, requiredConfigurationMissing(config.SocketConnectPort)
		}

		if ok := s.HasSetting(config.HeartBtInt); !ok {
			return nil, requiredConfigurationMissing(config.HeartBtInt)
		}

		session, err := createSession(sessionID, storeFactory, s, logFactory, app)
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

func (i *Initiator) handleConnection(address string, session *session, reconnectInterval time.Duration, tlsConfig *tls.Config) {
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

	for {
		if !i.waitForInSessionTime(session) {
			return
		}

		var disconnected chan interface{}
		var msgIn chan fixIn
		var msgOut chan []byte

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
		if err := session.initiate(msgIn, msgOut); err != nil {
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
		session.log.OnEventf("%v Reconnecting in %v", session.sessionID, reconnectInterval)
		if !i.waitForReconnectInterval(reconnectInterval) {
			return
		}
	}
}
