package errors

import (
	"fmt"
)

//IncorrectBeginString is a message reject specific to incorrect begin strings.
type IncorrectBeginString struct{ messageRejectError }

func (e IncorrectBeginString) Error() string { return "Incorrect BeginString" }

//TargetTooHigh is a MessageReject where the sequence number is larger than expected.
type TargetTooHigh struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e TargetTooHigh) Error() string {
	return fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}

//TargetTooLow is a MessageReject where the sequence number is less than expected.
type TargetTooLow struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e TargetTooLow) Error() string {
	return fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}
