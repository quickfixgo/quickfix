package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/log"
	"github.com/quickfixgo/quickfix/settings"
	"net"
)

//Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        settings.ApplicationSettings
	sessionSettings map[SessionID]settings.Dictionary
	logFactory      log.LogFactory
	globalLog       log.Log
}

//Start Initiator.
func (i *Initiator) Start() error {

	for sessionID, s := range i.sessionSettings {
		socketConnectHost, _ := s.GetString(settings.SocketConnectHost)
		socketConnectPort, _ := s.GetInt(settings.SocketConnectPort)

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
func NewInitiator(app Application, appSettings settings.ApplicationSettings, logFactory log.LogFactory) (*Initiator, error) {
	i := new(Initiator)
	i.app = app
	i.settings = appSettings
	i.sessionSettings = make(map[SessionID]settings.Dictionary)
	i.logFactory = logFactory
	i.globalLog = logFactory.Create()

	for _, s := range appSettings.GetSessionSettings() {

		//fail fast
		if _, ok := s.GetString(settings.SocketConnectHost); !ok {
			return nil, settings.RequiredConfigurationMissing(settings.SocketConnectHost)
		}

		if _, ok := s.GetInt(settings.SocketConnectPort); !ok {
			return nil, settings.RequiredConfigurationMissing(settings.SocketConnectPort)
		}

		sessionID, err := createSession(s, logFactory, app)
		if err != nil {
			return nil, err
		}

		i.sessionSettings[sessionID] = s
	}

	return i, nil
}
