package quickfix

import (
	"testing"

	"github.com/stretchr/testify/suite"
)

type LatentStateTestSuite struct {
	SessionSuite
}

func TestLatentStateTestSuite(t *testing.T) {
	suite.Run(t, new(LatentStateTestSuite))
}

func (s *LatentStateTestSuite) SetupTest() {
	s.Init()
	s.session.State = latentState{}
}

func (s *LatentStateTestSuite) TestPreliminary() {
	s.False(s.session.IsLoggedOn())
	s.False(s.session.IsConnected())
	s.True(s.session.IsSessionTime())
}

func (s *LatentStateTestSuite) TestDisconnected() {
	s.session.Disconnected(s.session)
	s.State(latentState{})
}

func (s *LatentStateTestSuite) TestStop() {
	notify := make(chan interface{})
	s.session.Stop(s.session, notify)
	s.State(latentState{})
	_, ok := <-notify
	s.False(ok)
}
