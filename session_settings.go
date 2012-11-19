package quickfixgo

import(
    "fmt"
    )

type SessionSettings interface {
  GetSessionSettings(sessionID SessionID) (dict Dictionary, ok bool)
  SetSessionSettings(sessionID SessionID, dict Dictionary ) error

  GetGlobalSettings() Dictionary
}

type sessionSettings struct {
  globalSettings Dictionary
  sessionSettings map[SessionID] Dictionary
}

func (s * sessionSettings) GetGlobalSettings() Dictionary {
  return s.globalSettings
}

func (s * sessionSettings) GetSessionSettings(sessionID SessionID) (dict Dictionary, ok bool) {
  return 
}

func (s * sessionSettings) SetSessionSettings(sessionID SessionID, settings Dictionary) error {
  if _, sessionExists:=s.sessionSettings[sessionID]; sessionExists {
    return fmt.Errorf("Duplicate Session %v ", sessionID)
  }

  settings.SetString(BeginString, sessionID.BeginString)
  settings.SetString(SenderCompID, sessionID.SenderCompID)
  settings.SetString(TargetCompID, sessionID.TargetCompID)

  s.sessionSettings[sessionID]=settings

  return nil
}


func NewSessionSettings() SessionSettings {
  s:=sessionSettings{globalSettings: NewDictionary(), 
      sessionSettings: make(map[SessionID] Dictionary)}

  return &s
}
