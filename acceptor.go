package quickfixgo

type Acceptor interface {
  //Start Acceptor.
  Start() (bool, error)

  //Stop Acceptor.
  Stop()
}
