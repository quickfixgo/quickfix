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
