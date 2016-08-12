package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type NotSessionTimeTestSuite struct {
	SessionSuiteRig
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
	s.session.Stop(s.session)
	s.Stopped()
}
