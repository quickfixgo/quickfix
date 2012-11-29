package session

//Callback interface for receiving session events
type Callback interface {
  OnCreate(sessionID ID)
  OnLogon(sessionID ID)
  OnLogout(sessionID ID)
}
