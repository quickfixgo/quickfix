package quickfix

import (
	"crypto/tls"
	"fmt"
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
		address := fmt.Sprintf("%v:%v", socketConnectHost, socketConnectPort)

		i.wg.Add(1)
		go func(sessID SessionID) {
			handleInitiatorConnection(address, i.globalLog, sessID, time.Duration(reconnectInterval)*time.Second, tlsConfig, i.stopChan)
			i.wg.Done()
		}(sessionID)
	}

	return nil
}

//Stop Initiator.
func (i *Initiator) Stop() {
	close(i.stopChan)

	for sessionID := range i.sessionSettings {
		session, err := lookupSession(sessionID)
		if err != nil {
			i.globalLog.OnEventf("Error getting session: %v", err)
		} else {
			go session.disconnect()
		}
	}

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

		err = createSession(sessionID, storeFactory, s, logFactory, app)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}
