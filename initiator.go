package quickfix

type Initiator interface {
	//Start Initiator.
	Start() (bool, error)

	//Stop Initiator.
	Stop()
}
