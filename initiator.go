package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/errors"
	"net"
)

//Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        *Settings
	sessionSettings map[SessionID]*SessionSettings
	logFactory      LogFactory
	globalLog       Log
}

//Start Initiator.
func (i *Initiator) Start() error {

	for sessionID, s := range i.sessionSettings {
		socketConnectHost, err := s.Setting(config.SocketConnectHost)
		if err != nil {
			return fmt.Errorf("error on SocketConnectHost: %v", err)
		}

		socketConnectPort, err := s.IntSetting(config.SocketConnectPort)
		if err != nil {
			return fmt.Errorf("error on SocketConnectPort: %v", err)
		}

		conn, err := net.Dial("tcp", fmt.Sprintf("%v:%v", socketConnectHost, socketConnectPort))
		if err != nil {
			return err
		}

		go handleInitiatorConnection(conn, i.globalLog, sessionID)
	}

	return nil
}

//Stop Initiator.
func (i *Initiator) Stop() {

}

//NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	i := new(Initiator)
	i.app = app
	i.settings = appSettings
	i.sessionSettings = make(map[SessionID]*SessionSettings)
	i.logFactory = logFactory
	i.globalLog = logFactory.Create()

	for _, s := range appSettings.SessionSettings() {

		//fail fast
		if ok := s.HasSetting(config.SocketConnectHost); !ok {
			return nil, errors.RequiredConfigurationMissing(config.SocketConnectHost)
		}

		if ok := s.HasSetting(config.SocketConnectPort); !ok {
			return nil, errors.RequiredConfigurationMissing(config.SocketConnectPort)
		}

		sessionID, err := createSession(s, logFactory, app)
		if err != nil {
			return nil, err
		}

		i.sessionSettings[sessionID] = s
	}

	return i, nil
}
