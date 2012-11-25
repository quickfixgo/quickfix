package session

type Callback interface {
  OnCreate(sessionID ID)
}
