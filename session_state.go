package quickfix

import "fmt"

//sessionState is the current state of the session state machine. The session state determines how the session responds to
//incoming messages, timeouts, and requests to send application messages.
type sessionState interface {
	//FixMsgIn is called by the session on incoming messages from the counter party.  The return type is the next session state
	//following message processing
	FixMsgIn(*session, Message) (nextState sessionState)

	//Timeout is called by the session on a timeout event.
	Timeout(*session, event) (nextState sessionState)

	//IsLoggedOn returns true if state is logged on an in session, false otherwise
	IsLoggedOn() bool

	//debugging convenience
	fmt.Stringer
}
