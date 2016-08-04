package quickfix

import (
	"crypto/tls"
	"fmt"
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
	quitChan        chan bool
}

//Start Initiator.
func (i *Initiator) Start() error {

	i.quitChan = make(chan bool)

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
		address := fmt.Sprintf("%v:%v", socketConnectHost, socketConnectPort)
		go handleInitiatorConnection(address, i.globalLog, sessionID, i.quitChan, time.Duration(reconnectInterval)*time.Second, tlsConfig)
	}

	return nil
}

//Stop Initiator.
func (i *Initiator) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()
	close(i.quitChan)
}

//NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, storeFactory MessageStoreFactory, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	i := new(Initiator)
	i.app = app
	i.storeFactory = storeFactory
	i.settings = appSettings
	i.sessionSettings = appSettings.SessionSettings()
	i.logFactory = logFactory

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

		err = createSession(sessionID, storeFactory, s, logFactory, app)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}
