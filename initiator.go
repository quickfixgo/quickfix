package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/config"
	"net"
)

//Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        *Settings
	sessionSettings map[SessionID]*SessionSettings
	storeFactory    MessageStoreFactory
	logFactory      LogFactory
	globalLog       Log
	quitChans       map[SessionID]chan bool
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

		i.quitChans[sessionID] = make(chan bool)
		go handleInitiatorConnection(conn, i.globalLog, sessionID, i.quitChans[sessionID])
	}

	return nil
}

//Stop Initiator.
func (i *Initiator) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()
	for _, channel := range i.quitChans {
		channel <- true
	}
}

//NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, storeFactory MessageStoreFactory, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	i := new(Initiator)
	i.app = app
	i.storeFactory = storeFactory
	i.settings = appSettings
	i.sessionSettings = appSettings.SessionSettings()
	i.logFactory = logFactory
	i.quitChans = make(map[SessionID]chan bool)

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

		err = createSession(sessionID, storeFactory, s, logFactory, app)
		if err != nil {
			return nil, err
		}
	}

	return i, nil
}
