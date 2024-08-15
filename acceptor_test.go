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
	"bytes"
	"io"
	"net"
	"strings"
	"testing"
	"time"

	proxyproto "github.com/pires/go-proxyproto"
	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/suite"
)

func TestAcceptor_Start(t *testing.T) {
	sessionSettings := NewSessionSettings()
	sessionSettings.Set(config.BeginString, BeginStringFIX42)
	sessionSettings.Set(config.SenderCompID, "sender")
	sessionSettings.Set(config.TargetCompID, "target")

	settingsWithTCPProxy := NewSettings()
	settingsWithTCPProxy.GlobalSettings().Set("UseTCPProxy", "Y")

	settingsWithNoTCPProxy := NewSettings()
	settingsWithNoTCPProxy.GlobalSettings().Set("UseTCPProxy", "N")

	genericSettings := NewSettings()

	const (
		GenericListener = iota
		ProxyListener
	)

	acceptorStartTests := []struct {
		name         string
		settings     *Settings
		listenerType int
	}{
		{"with TCP proxy set", settingsWithTCPProxy, ProxyListener},
		{"with no TCP proxy set", settingsWithNoTCPProxy, GenericListener},
		{"no TCP proxy configuration set", genericSettings, GenericListener},
	}

	for _, tt := range acceptorStartTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.settings.GlobalSettings().Set("SocketAcceptPort", "5001")
			if _, err := tt.settings.AddSession(sessionSettings); err != nil {
				assert.Nil(t, err)
			}

			acceptor := &Acceptor{settings: tt.settings}
			if err := acceptor.Start(); err != nil {
				assert.NotNil(t, err)
			}
			assert.Len(t, acceptor.listeners, 1)

			for _, listener := range acceptor.listeners {
				if tt.listenerType == ProxyListener {
					_, ok := listener.(*proxyproto.Listener)
					assert.True(t, ok)
				}

				if tt.listenerType == GenericListener {
					_, ok := listener.(*net.TCPListener)
					assert.True(t, ok)
				}
			}

			acceptor.Stop()
		})
	}
}

var _ net.Conn = &mockConn{}

type mockConn struct {
	closeChan  chan struct{}
	localAddr  net.Addr
	remoteAddr net.Addr

	onWriteback     func([]byte)
	inboundMessages []*Message
}

func (c *mockConn) Read(b []byte) (n int, err error) {
	if len(c.inboundMessages) > 0 {
		messageBytes := c.inboundMessages[0].build()
		copy(b, messageBytes)
		c.inboundMessages = c.inboundMessages[1:]
		return len(messageBytes), err
	}
	<-c.closeChan
	return 0, io.EOF
}

func (c *mockConn) LocalAddr() net.Addr {
	return c.localAddr
}

func (c *mockConn) RemoteAddr() net.Addr {
	return c.remoteAddr
}

func (c *mockConn) SetDeadline(_ time.Time) error {
	return nil
}

func (c *mockConn) SetReadDeadline(_ time.Time) error {
	return nil
}

func (c *mockConn) SetWriteDeadline(_ time.Time) error {
	return nil
}

func (c *mockConn) Write(b []byte) (n int, err error) {
	if c.onWriteback != nil {
		c.onWriteback(b)
	}
	return len(b), nil
}

func (c *mockConn) Close() error {
	return nil
}

func mockLogonMessage(sessionID SessionID, msgSeqNum int) *Message {
	msg := NewMessage()
	msg.Header.SetField(tagMsgType, FIXString("A"))
	msg.Header.SetInt(tagMsgSeqNum, msgSeqNum)
	msg.Header.SetString(tagBeginString, sessionID.BeginString)
	msg.Header.SetString(tagSenderCompID, sessionID.SenderCompID)
	msg.Header.SetString(tagSenderSubID, sessionID.SenderSubID)
	msg.Header.SetString(tagSenderLocationID, sessionID.SenderLocationID)
	msg.Header.SetString(tagTargetCompID, sessionID.TargetCompID)
	msg.Header.SetString(tagTargetSubID, sessionID.TargetSubID)
	msg.Header.SetString(tagTargetLocationID, sessionID.TargetLocationID)
	msg.Header.SetField(tagSendingTime, FIXUTCTimestamp{Time: time.Now()})
	msg.Body.SetInt(tagHeartBtInt, 30)
	return msg
}

type AcceptorTemplateTestSuite struct {
	suite.Suite
	acceptor *Acceptor

	sessionID1 SessionID
	sessionID2 SessionID
	sessionID3 SessionID

	cliSessionID   SessionID
	logonSessionID SessionID
	seqNum         int
}

func (suite *AcceptorTemplateTestSuite) BeforeTest(_, _ string) {
	cfg := `
[default]
ConnectionType=acceptor
SocketAcceptPort=5001
TimeZone=America/New_York
StartTime=00:00:01
EndTime=23:59:59
HeartBtInt=30

[session]
BeginString=FIX.4.2
SenderCompID=sender1
TargetCompID=target1

[session]
BeginString=FIX.4.3
SenderCompID=sender2
TargetCompID=target2
ResetOnLogon=Y

[session]
AcceptorTemplate=Y
BeginString=FIX.4.3
SenderCompID=*
SenderSubID=*
SenderLocationID=*
TargetCompID=target3
TargetSubID=*
TargetLocationID=*
ResetOnLogout=Y
`
	stringReader := strings.NewReader(cfg)
	settings, err := ParseSettings(stringReader)
	if err != nil {
		suite.FailNow("parse setting failed", err)
	}
	suite.sessionID1 = SessionID{BeginString: BeginStringFIX42, SenderCompID: "sender1", TargetCompID: "target1"}
	suite.sessionID2 = SessionID{BeginString: BeginStringFIX43, SenderCompID: "sender2", TargetCompID: "target2"}
	suite.sessionID3 = SessionID{BeginString: BeginStringFIX43, SenderCompID: "*", SenderSubID: "*", SenderLocationID: "*",
		TargetCompID: "target3", TargetSubID: "*", TargetLocationID: "*"}

	app := &noopApp{}
	a, err := NewAcceptor(app, memoryStoreFactory{}, settings, NewNullLogFactory())
	if err != nil {
		suite.Fail("Failed to create acceptor", err)
	}
	suite.acceptor = a

	suite.cliSessionID = SessionID{BeginString: BeginStringFIX43, SenderCompID: "target3", TargetCompID: "dynamicSender"}
	suite.logonSessionID = SessionID{BeginString: BeginStringFIX43, SenderCompID: "dynamicSender", TargetCompID: "target3"}
	suite.seqNum = 1
}

func (suite *AcceptorTemplateTestSuite) TearDownTest() {
	suite.acceptor.Stop()
	suite.acceptor = nil
	suite.seqNum = 1
}

func (suite *AcceptorTemplateTestSuite) logonAndDisconnectAfterCheck(sessionID SessionID,
	checkFuncAfterLogon func(),
	wantLogonSuccess bool) {
	inboundMessages := []*Message{mockLogonMessage(sessionID, suite.seqNum)}
	suite.seqNum++
	var respondedLogonMessageReceived bool
	mockConn1 := &mockConn{
		closeChan:       make(chan struct{}),
		inboundMessages: inboundMessages,
		localAddr:       &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5001},
		remoteAddr:      &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5002},
	}
	mockConn1.onWriteback = func(b []byte) {
		responseMsg := NewMessage()
		err := ParseMessage(responseMsg, bytes.NewBuffer(b))
		suite.Require().NoError(err, "parse responding message failed")
		msgType, err := responseMsg.Header.GetString(tagMsgType)
		suite.Require().NoError(err, "unexpected mssage")
		if wantLogonSuccess && msgType != "A" {
			return
		}
		respondedLogonMessageReceived = true
		if checkFuncAfterLogon != nil {
			checkFuncAfterLogon()
		}
		close(mockConn1.closeChan)
	}
	suite.acceptor.handleConnection(mockConn1)
	if wantLogonSuccess {
		suite.Require().Equal(true, respondedLogonMessageReceived, "expected responding logon message")
	}
}

func (suite *AcceptorTemplateTestSuite) verifySessionCount(expectedSessionCount int) {
	suite.Require().Equalf(expectedSessionCount, len(suite.acceptor.sessions), "expected %v sessions but found %v", expectedSessionCount, len(suite.acceptor.sessions))
	suite.Require().Equalf(expectedSessionCount, len(sessions), "expected %v sessions but found %v in registry", expectedSessionCount, sessions)
}

func (suite *AcceptorTemplateTestSuite) TestCreateDynamicSessionBySessionProvider() {
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	logonSessionID := suite.logonSessionID
	suite.logonAndDisconnectAfterCheck(suite.cliSessionID, func() {
		suite.verifySessionCount(3)

		createdSession, ok := suite.acceptor.sessions[logonSessionID]
		suite.Require().Equal(true, ok, "expected dynamic session to be created")
		suite.Require().Equal(logonSessionID, createdSession.sessionID, "expected session ID to match inbound session ID")
		suite.Require().Equal(createdSession.ResetOnLogout, true, "expected ResetOnLogout=Y for createdSession")

		remoteAddr, ok := suite.acceptor.RemoteAddr(logonSessionID)
		if !ok {
			suite.Fail("Failed to get remote address for dynamic session")
		}
		suite.Require().Equal("127.0.0.1:5002", remoteAddr.String(), "expect remoteAddr for dynamic session to be 127.0.0.1:5002 but got %v", remoteAddr.String())
	}, true)
}

func (suite *AcceptorTemplateTestSuite) TestSessionCreatedBySessionProviderShouldBeKept() {
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	logonSessionID := suite.logonSessionID
	suite.logonAndDisconnectAfterCheck(suite.cliSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	err := SendToTarget(createFIX43NewOrderSingle(), logonSessionID)
	suite.NoError(err, "expected message can still be sent after session disconnected")
}

func (suite *AcceptorTemplateTestSuite) TestNoNewSessionCreatedWhenSameSessionIDLogons() {
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	suite.logonAndDisconnectAfterCheck(suite.cliSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	suite.logonAndDisconnectAfterCheck(suite.cliSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	suite.logonAndDisconnectAfterCheck(suite.cliSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
}

func (suite *AcceptorTemplateTestSuite) TestSessionNotFoundBySessionProvider() {
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	sessionID := SessionID{BeginString: BeginStringFIX43, SenderCompID: "unknownSender", TargetCompID: "unknownTarget"}
	suite.logonAndDisconnectAfterCheck(sessionID, func() {}, false)
	suite.verifySessionCount(2)
}

type mockCustomTemplateIDProvider struct {
	staticTemplateID SessionID
}

// mockCustomTemplateIDProvider always returns the same templateID
func (p *mockCustomTemplateIDProvider) GetTemplateID(inboundSessionID SessionID) *SessionID {
	return &p.staticTemplateID
}

func (suite *AcceptorTemplateTestSuite) TestCustomTemplateIDProvider_NoSessionCreated() {
	// this templateIDProvider selects session FIX.4.3:sender2->sender2 as the template
	templateIDProvider := &mockCustomTemplateIDProvider{staticTemplateID: SessionID{
		BeginString: BeginStringFIX43, SenderCompID: "sender2", TargetCompID: "target2",
	}}
	suite.acceptor.SetTemplateIDProvider(templateIDProvider)
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	// no session created
	logon1 := SessionID{BeginString: BeginStringFIX42, SenderCompID: "target1", TargetCompID: "sender1"}
	suite.logonAndDisconnectAfterCheck(logon1, func() {
		suite.verifySessionCount(2)
	}, true)
}

func (suite *AcceptorTemplateTestSuite) TestCustomTemplateIDProvider_SessionCreated() {
	// this templateIDProvider selects session FIX.4.3:sender2->sender2 as the template
	templateIDProvider := &mockCustomTemplateIDProvider{staticTemplateID: SessionID{
		BeginString: BeginStringFIX43, SenderCompID: "sender2", TargetCompID: "target2",
	}}
	suite.acceptor.SetTemplateIDProvider(templateIDProvider)
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed", err)
	}
	suite.verifySessionCount(2)

	// session created
	logonSessionID2 := SessionID{BeginString: BeginStringFIX42, SenderCompID: "any", TargetCompID: "any"}
	suite.logonAndDisconnectAfterCheck(logonSessionID2, func() {
		suite.verifySessionCount(3)
	}, true)
	// logon again
	suite.logonAndDisconnectAfterCheck(logonSessionID2, func() {
		suite.verifySessionCount(3)
	}, true)

	session2 := suite.acceptor.sessions[logonSessionID2]
	suite.Require().Equal(true, session2.ResetOnLogon, "expected session2 ResetOnLogon=Y")
}

func TestAcceptorTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(AcceptorTemplateTestSuite))
}
