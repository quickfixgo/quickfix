package quickfix

import (
	"fmt"
)

//IncorrectBeginString is a message reject specific to incorrect begin strings.
type incorrectBeginString struct{ messageRejectError }

func (e incorrectBeginString) Error() string { return "Incorrect BeginString" }

//targetTooHigh is a MessageReject where the sequence number is larger than expected.
type targetTooHigh struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e targetTooHigh) Error() string {
	return fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}

//targetTooLow is a MessageReject where the sequence number is less than expected.
type targetTooLow struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e targetTooLow) Error() string {
	return fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}
