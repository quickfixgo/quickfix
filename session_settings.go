package quickfixgo

type SessionSettings interface {
  GetSessionSettings(sessionID SessionID) Dictionary
  GetGlobalSettings() Dictionary
}

type sessionSettings struct {
  globalSettings Dictionary
}

func (s * sessionSettings) GetGlobalSettings() Dictionary {
  return s.globalSettings
}

func (s * sessionSettings) GetSessionSettings(sessionID SessionID) Dictionary {
  return nil
}

func NewSessionSettings() SessionSettings {
  s:=sessionSettings{}
  s.globalSettings=newDictionary()

  return &s
}
