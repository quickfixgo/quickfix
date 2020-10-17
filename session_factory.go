package quickfix

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/quickfixgo/quickfix/internal"
)

var dayLookup = map[string]time.Weekday{
	"Sunday":    time.Sunday,
	"Monday":    time.Monday,
	"Tuesday":   time.Tuesday,
	"Wednesday": time.Wednesday,
	"Thursday":  time.Thursday,
	"Friday":    time.Friday,
	"Saturday":  time.Saturday,

	"Sun": time.Sunday,
	"Mon": time.Monday,
	"Tue": time.Tuesday,
	"Wed": time.Wednesday,
	"Thu": time.Thursday,
	"Fri": time.Friday,
	"Sat": time.Saturday,
}

var applVerIDLookup = map[string]string{
	BeginStringFIX40: "2",
	BeginStringFIX41: "3",
	BeginStringFIX42: "4",
	BeginStringFIX43: "5",
	BeginStringFIX44: "6",
	"FIX.5.0":        "7",
	"FIX.5.0SP1":     "8",
	"FIX.5.0SP2":     "9",
}

type sessionFactory struct {
	//True if building sessions that initiate logon
	BuildInitiators bool
}

//Creates Session, associates with internal session registry
func (f sessionFactory) createSession(
	sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings,
	logFactory LogFactory, application Application,
) (session *session, err error) {

	if session, err = f.newSession(sessionID, storeFactory, settings, logFactory, application); err != nil {
		return
	}

	if err = registerSession(session); err != nil {
		return
	}
	application.OnCreate(session.sessionID)
	session.log.OnEvent("Created session")

	return
}

func (f sessionFactory) newSession(
	sessionID SessionID, storeFactory MessageStoreFactory, settings *SessionSettings, logFactory LogFactory,
	application Application) (s *session, err error) {
	s = &session{sessionID: sessionID}

	var validatorSettings = defaultValidatorSettings
	if settings.HasSetting(config.ValidateFieldsOutOfOrder) {
		if validatorSettings.CheckFieldsOutOfOrder, err = settings.BoolSetting(config.ValidateFieldsOutOfOrder); err != nil {
			return
		}
	}

	if settings.HasSetting(config.RejectInvalidMessage) {
		if validatorSettings.RejectInvalidMessage, err = settings.BoolSetting(config.RejectInvalidMessage); err != nil {
			return
		}
	}

	if sessionID.IsFIXT() {
		if s.DefaultApplVerID, err = settings.Setting(config.DefaultApplVerID); err != nil {
			return
		}

		if applVerID, ok := applVerIDLookup[s.DefaultApplVerID]; ok {
			s.DefaultApplVerID = applVerID
		}

		//If the transport or app data dictionary setting is set, the other also needs to be set.
		if settings.HasSetting(config.TransportDataDictionary) || settings.HasSetting(config.AppDataDictionary) {
			var transportDataDictionaryPath, appDataDictionaryPath string
			if transportDataDictionaryPath, err = settings.Setting(config.TransportDataDictionary); err != nil {
				return
			}

			if appDataDictionaryPath, err = settings.Setting(config.AppDataDictionary); err != nil {
				return
			}

			if s.transportDataDictionary, err = datadictionary.Parse(transportDataDictionaryPath); err != nil {
				return
			}

			if s.appDataDictionary, err = datadictionary.Parse(appDataDictionaryPath); err != nil {
				return
			}

			s.validator = &fixtValidator{s.transportDataDictionary, s.appDataDictionary, validatorSettings}
		}
	} else if settings.HasSetting(config.DataDictionary) {
		var dataDictionaryPath string
		if dataDictionaryPath, err = settings.Setting(config.DataDictionary); err != nil {
			return
		}

		if s.appDataDictionary, err = datadictionary.Parse(dataDictionaryPath); err != nil {
			return
		}

		s.validator = &fixValidator{s.appDataDictionary, validatorSettings}
	}

	if settings.HasSetting(config.ResetOnLogon) {
		if s.ResetOnLogon, err = settings.BoolSetting(config.ResetOnLogon); err != nil {
			return
		}
	}

	if settings.HasSetting(config.RefreshOnLogon) {
		if s.RefreshOnLogon, err = settings.BoolSetting(config.RefreshOnLogon); err != nil {
			return
		}
	}

	if settings.HasSetting(config.ResetOnLogout) {
		if s.ResetOnLogout, err = settings.BoolSetting(config.ResetOnLogout); err != nil {
			return
		}
	}

	if settings.HasSetting(config.ResetOnDisconnect) {
		if s.ResetOnDisconnect, err = settings.BoolSetting(config.ResetOnDisconnect); err != nil {
			return
		}
	}

	if settings.HasSetting(config.EnableLastMsgSeqNumProcessed) {
		if s.EnableLastMsgSeqNumProcessed, err = settings.BoolSetting(config.EnableLastMsgSeqNumProcessed); err != nil {
			return
		}
	}

	if settings.HasSetting(config.CheckLatency) {
		var doCheckLatency bool
		if doCheckLatency, err = settings.BoolSetting(config.CheckLatency); err != nil {
			return
		}

		s.SkipCheckLatency = !doCheckLatency
	}

	if !settings.HasSetting(config.MaxLatency) {
		s.MaxLatency = 120 * time.Second
	} else {
		var maxLatency int
		if maxLatency, err = settings.IntSetting(config.MaxLatency); err != nil {
			return
		}

		if maxLatency <= 0 {
			err = errors.New("MaxLatency must be a positive integer")
			return
		}

		s.MaxLatency = time.Duration(maxLatency) * time.Second
	}

	if settings.HasSetting(config.ResendRequestChunkSize) {
		if s.ResendRequestChunkSize, err = settings.IntSetting(config.ResendRequestChunkSize); err != nil {
			return
		}
	}

	if settings.HasSetting(config.StartTime) || settings.HasSetting(config.EndTime) {
		var startTimeStr, endTimeStr string
		if startTimeStr, err = settings.Setting(config.StartTime); err != nil {
			return
		}

		if endTimeStr, err = settings.Setting(config.EndTime); err != nil {
			return
		}

		var start, end internal.TimeOfDay
		if start, err = internal.ParseTimeOfDay(startTimeStr); err != nil {
			return
		}

		if end, err = internal.ParseTimeOfDay(endTimeStr); err != nil {
			return
		}

		loc := time.UTC
		if settings.HasSetting(config.TimeZone) {
			var locStr string
			if locStr, err = settings.Setting(config.TimeZone); err != nil {
				return
			}

			loc, err = time.LoadLocation(locStr)
			if err != nil {
				return
			}
		}

		if !settings.HasSetting(config.StartDay) && !settings.HasSetting(config.EndDay) {
			s.SessionTime = internal.NewTimeRangeInLocation(start, end, loc)
		} else {
			var startDayStr, endDayStr string
			if startDayStr, err = settings.Setting(config.StartDay); err != nil {
				return
			}

			if endDayStr, err = settings.Setting(config.EndDay); err != nil {
				return
			}

			parseDay := func(setting, dayStr string) (day time.Weekday, err error) {
				day, ok := dayLookup[dayStr]
				if !ok {
					return day, IncorrectFormatForSetting{Setting: setting, Value: dayStr}
				}
				return
			}

			var startDay, endDay time.Weekday
			if startDay, err = parseDay(config.StartDay, startDayStr); err != nil {
				return
			}

			if endDay, err = parseDay(config.EndDay, endDayStr); err != nil {
				return
			}

			s.SessionTime = internal.NewWeekRangeInLocation(start, end, startDay, endDay, loc)
		}
	}

	if settings.HasSetting(config.TimeStampPrecision) {
		var precisionStr string
		if precisionStr, err = settings.Setting(config.TimeStampPrecision); err != nil {
			return
		}

		switch precisionStr {
		case "SECONDS":
			s.timestampPrecision = Seconds
		case "MILLIS":
			s.timestampPrecision = Millis
		case "MICROS":
			s.timestampPrecision = Micros
		case "NANOS":
			s.timestampPrecision = Nanos

		default:
			err = IncorrectFormatForSetting{Setting: config.TimeStampPrecision, Value: precisionStr}
			return
		}
	}

	if settings.HasSetting(config.PersistMessages) {
		var persistMessages bool
		if persistMessages, err = settings.BoolSetting(config.PersistMessages); err != nil {
			return
		}

		s.DisableMessagePersist = !persistMessages
	}

	if f.BuildInitiators {
		if err = f.buildInitiatorSettings(s, settings); err != nil {
			return
		}
	}

	if s.log, err = logFactory.CreateSessionLog(s.sessionID); err != nil {
		return
	}

	if s.store, err = storeFactory.Create(s.sessionID); err != nil {
		return
	}

	s.sessionEvent = make(chan internal.Event)
	s.messageEvent = make(chan bool, 1)
	s.admin = make(chan interface{})
	s.application = application
	return
}

func (f sessionFactory) buildInitiatorSettings(session *session, settings *SessionSettings) error {
	session.InitiateLogon = true

	heartBtInt, err := settings.IntSetting(config.HeartBtInt)
	if err != nil {
		return err
	}

	if heartBtInt <= 0 {
		return errors.New("Heartbeat must be greater than zero")
	}
	session.HeartBtInt = time.Duration(heartBtInt) * time.Second

	session.ReconnectInterval = 30 * time.Second
	if settings.HasSetting(config.ReconnectInterval) {

		interval, err := settings.IntSetting(config.ReconnectInterval)
		if err != nil {
			return err
		}

		if interval <= 0 {
			return errors.New("ReconnectInterval must be greater than zero")
		}

		session.ReconnectInterval = time.Duration(interval) * time.Second
	}

	session.LogoutTimeout = 2 * time.Second
	if settings.HasSetting(config.LogoutTimeout) {

		timeout, err := settings.IntSetting(config.LogoutTimeout)
		if err != nil {
			return err
		}

		if timeout <= 0 {
			return errors.New("LogoutTimeout must be greater than zero")
		}

		session.LogoutTimeout = time.Duration(timeout) * time.Second
	}

	session.LogonTimeout = 10 * time.Second
	if settings.HasSetting(config.LogonTimeout) {

		timeout, err := settings.IntSetting(config.LogonTimeout)
		if err != nil {
			return err
		}

		if timeout <= 0 {
			return errors.New("LogonTimeout must be greater than zero")
		}

		session.LogonTimeout = time.Duration(timeout) * time.Second
	}

	return f.configureSocketConnectAddress(session, settings)
}

func (f sessionFactory) configureSocketConnectAddress(session *session, settings *SessionSettings) (err error) {
	session.SocketConnectAddress = []string{}

	var socketConnectHost, socketConnectPort string
	for i := 0; ; {

		hostConfig := config.SocketConnectHost
		portConfig := config.SocketConnectPort

		if i > 0 {
			hostConfig = hostConfig + strconv.Itoa(i)
			portConfig = portConfig + strconv.Itoa(i)

			if !(settings.HasSetting(hostConfig) || settings.HasSetting(portConfig)) {
				return
			}
		}

		if socketConnectHost, err = settings.Setting(hostConfig); err != nil {
			return
		}

		if socketConnectPort, err = settings.Setting(portConfig); err != nil {
			return
		}

		session.SocketConnectAddress = append(session.SocketConnectAddress, net.JoinHostPort(socketConnectHost, socketConnectPort))
		i++
	}
}
