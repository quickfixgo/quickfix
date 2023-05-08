// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"fmt"
)

// IncorrectBeginString is a message reject specific to incorrect begin strings.
type incorrectBeginString struct{ messageRejectError }

func (e incorrectBeginString) Error() string { return "Incorrect BeginString" }

// targetTooHigh is a MessageReject where the sequence number is larger than expected.
type targetTooHigh struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e targetTooHigh) Error() string {
	return fmt.Sprintf("MsgSeqNum too high, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}

// targetTooLow is a MessageReject where the sequence number is less than expected.
type targetTooLow struct {
	messageRejectError
	ReceivedTarget int
	ExpectedTarget int
}

func (e targetTooLow) Error() string {
	return fmt.Sprintf("MsgSeqNum too low, expecting %d but received %d", e.ExpectedTarget, e.ReceivedTarget)
}
