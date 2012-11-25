package session

import("fmt")

// SessionID is a unique identifer of a Session
type ID struct {
  BeginString, TargetCompID, SenderCompID, DefaultApplVerID string
}

func (s *ID) String() string{
  return fmt.Sprintf("%s:%s->%s", s.BeginString, s.SenderCompID, s.TargetCompID)
}

