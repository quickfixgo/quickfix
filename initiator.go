package quickfixgo

type Initiator interface {
  //Start Initiator.
  Start() (bool, error)

  //Stop Initiator.
  Stop()
}
