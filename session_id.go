package quickfix

import (
	"fmt"
)

// SessionID is a unique identifer of a Session
type SessionID struct {
	BeginString, TargetCompID, SenderCompID, DefaultApplVerID string
}

func (s SessionID) String() string {
	return fmt.Sprintf("%s:%s->%s", s.BeginString, s.SenderCompID, s.TargetCompID)
}
