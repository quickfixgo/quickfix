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

func (s *LatentStateTestSuite) TestIsLoggedOn() {
	s.False(s.session.IsLoggedOn())
}

func (s *LatentStateTestSuite) TestIsConnected() {
	s.False(s.session.IsConnected())
}
