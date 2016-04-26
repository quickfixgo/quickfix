package quickfix

import (
	"fmt"
	"github.com/quickfixgo/quickfix/enum"
)

// SessionID is a unique identifer of a Session
type SessionID struct {
	BeginString, TargetCompID, TargetSubID, TargetLocationID, SenderCompID, SenderSubID, SenderLocationID, Qualifier string
}

//IsFIXT returns true if the SessionID has a FIXT BeginString
func (s SessionID) IsFIXT() bool {
	return s.BeginString == enum.BeginStringFIXT11
}

func (s SessionID) String() string {
	if len(s.Qualifier) > 0 {
		return fmt.Sprintf("%s:%s->%s:%s", s.BeginString, s.SenderCompID, s.TargetCompID, s.Qualifier)
	}

	return fmt.Sprintf("%s:%s->%s", s.BeginString, s.SenderCompID, s.TargetCompID)
}
