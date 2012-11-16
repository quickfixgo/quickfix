package quickfixgo

type SessionSettings interface {
  GetSessionSettings(sessionID SessionID) Dictionary
  GetGlobalSettings() Dictionary
}
