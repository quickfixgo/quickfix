package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotSessionTimeTestSuite struct {
	SessionSuite
}

func TestNotSessionTime(t *testing.T) {
	suite.Run(t, new(NotSessionTimeTestSuite))
}

func (s *NotSessionTimeTestSuite) SetupTest() {
	s.Init()
	s.session.State = notSessionTime{}
}

func (s *NotSessionTimeTestSuite) TestPreliminary() {
	s.False(s.session.IsLoggedOn())
	s.False(s.session.IsConnected())
	s.False(s.session.IsSessionTime())
}

func (s *NotSessionTimeTestSuite) TestDisconnected() {
	s.session.Disconnected(s.session)
	s.State(notSessionTime{})
}

func (s *NotSessionTimeTestSuite) TestStop() {
	notify := make(chan interface{})
	s.session.Stop(s.session, notify)
	s.State(notSessionTime{})
	_, ok := <-notify
	s.False(ok)
}
