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
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/config"

	proxyproto "github.com/pires/go-proxyproto"
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

func (c *mockConn) SetDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetReadDeadline(t time.Time) error {
	return nil
}

func (c *mockConn) SetWriteDeadline(t time.Time) error {
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

	sessionId1 SessionID
	sessionId2 SessionID
	sessionId3 SessionID

	testDynamicSessionID SessionID
	logonSessionID       SessionID
	seqNum               int

	dynamicSessionProvider AcceptorSessionProvider
}

func (suite *AcceptorTemplateTestSuite) BeforeTest(suiteName, testName string) {
	settings := NewSettings()
	settings.globalSettings.Set(config.SocketAcceptPort, "5001")
	sessionId1 := SessionID{BeginString: BeginStringFIX42, SenderCompID: "sender1", TargetCompID: "target1"}
	sessionSettings1 := NewSessionSettings()
	sessionSettings1.Set(config.BeginString, sessionId1.BeginString)
	sessionSettings1.Set(config.SenderCompID, sessionId1.SenderCompID)
	sessionSettings1.Set(config.TargetCompID, sessionId1.TargetCompID)
	suite.sessionId1 = sessionId1
	settings.AddSession(sessionSettings1)

	sessionId2 := SessionID{BeginString: BeginStringFIX43, SenderCompID: "sender2", TargetCompID: "target2"}
	sessionSettings2 := NewSessionSettings()
	sessionSettings2.Set(config.BeginString, sessionId2.BeginString)
	sessionSettings2.Set(config.SenderCompID, sessionId2.SenderCompID)
	sessionSettings2.Set(config.TargetCompID, sessionId2.TargetCompID)
	suite.sessionId2 = sessionId2
	settings.AddSession(sessionSettings2)

	// acceptor template
	sessionId3 := SessionID{BeginString: BeginStringFIX43, SenderCompID: "*", SenderSubID: "*", SenderLocationID: "*",
		TargetCompID: "target3", TargetSubID: "*", TargetLocationID: "*"}
	sessionSettings3 := NewSessionSettings()
	sessionSettings3.Set(config.BeginString, sessionId3.BeginString)
	sessionSettings3.Set(config.SenderCompID, sessionId3.SenderCompID)
	sessionSettings3.Set(config.SenderSubID, sessionId3.SenderSubID)
	sessionSettings3.Set(config.SenderLocationID, sessionId3.SenderLocationID)
	sessionSettings3.Set(config.TargetCompID, sessionId3.TargetCompID)
	sessionSettings3.Set(config.TargetSubID, sessionId3.TargetSubID)
	sessionSettings3.Set(config.TargetLocationID, sessionId3.TargetLocationID)
	sessionSettings3.Set(config.ResetOnLogout, "Y")
	sessionSettings3.Set(config.AcceptorTemplate, "Y")
	suite.sessionId3 = sessionId3
	settings.AddSession(sessionSettings3)

	app := &noopApp{}
	a, err := NewAcceptor(app, memoryStoreFactory{}, settings, NewScreenLogFactory())
	if err != nil {
		suite.Fail("Failed to create acceptor: %v", err)
	}
	suite.acceptor = a

	templateMappings := make([]*TemplateMapping, 0)
	templateMappings = append(templateMappings, &TemplateMapping{
		Pattern:    suite.sessionId3,
		TemplateID: suite.sessionId3,
	})
	suite.dynamicSessionProvider = NewDynamicAcceptorSessionProvider(suite.acceptor.settings, suite.acceptor.storeFactory, suite.acceptor.logFactory, suite.acceptor.app, templateMappings)
	suite.acceptor.SetSessionProvider(suite.dynamicSessionProvider)

	suite.testDynamicSessionID = SessionID{BeginString: BeginStringFIX43, SenderCompID: "target3", TargetCompID: "dynamicSender"}
	suite.logonSessionID = SessionID{BeginString: BeginStringFIX43, SenderCompID: "dynamicSender", TargetCompID: "target3"}
	if err := suite.acceptor.Start(); err != nil {
		suite.FailNow("acceptor start failed: %v", err)
	}

	suite.verifySessionCount(2)
	suite.seqNum = 1
}

func (suite *AcceptorTemplateTestSuite) logonAndDisconnectAfterCheck(sessionID SessionID,
	checkFuncAfterLogon func(),
	mustHaveResponse bool) {
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
		reponseMsg := NewMessage()
		err := ParseMessage(reponseMsg, bytes.NewBuffer(b))
		suite.Require().NoError(err, "parse responding message failed")
		msgType, err := reponseMsg.Header.GetString(tagMsgType)
		suite.Require().NoError(err, "unexpected mssage")
		suite.Require().Equalf("A", msgType, "expected logon message in reponse %s", reponseMsg.String())
		respondedLogonMessageReceived = true
		if checkFuncAfterLogon != nil {
			checkFuncAfterLogon()
		}
		close(mockConn1.closeChan)
	}
	suite.acceptor.handleConnection(mockConn1)
	if mustHaveResponse {
		suite.Require().Equal(true, respondedLogonMessageReceived, "expected responding logon message")
	}
}

func (suite *AcceptorTemplateTestSuite) verifySessionCount(expectedSessionCount int) {
	suite.Require().Equalf(expectedSessionCount, len(suite.acceptor.sessions), "expected %v sessions but found %v", expectedSessionCount, len(suite.acceptor.sessions))
	suite.Require().Equalf(expectedSessionCount, len(sessions), "expected %v sessions but found %v in registry", expectedSessionCount, len(suite.acceptor.sessions))
}

func (suite *AcceptorTemplateTestSuite) TestCreateDynamicSessionBySessionProvider() {
	logonSessionID := suite.logonSessionID
	suite.logonAndDisconnectAfterCheck(suite.testDynamicSessionID, func() {
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
	suite.acceptor.Stop()
}

func (suite *AcceptorTemplateTestSuite) TestSessionCreatedBySessionProviderShouldBeKept() {
	logonSessionID := suite.logonSessionID
	suite.logonAndDisconnectAfterCheck(suite.testDynamicSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	err := SendToTarget(createFIX43NewOrderSingle(), logonSessionID)
	suite.NoError(err, "expected message can still be sent after session disconnected")
	suite.acceptor.Stop()
}

func (suite *AcceptorTemplateTestSuite) TestNoNewSessionCreatedWhenSameSessionIDLogons() {
	suite.logonAndDisconnectAfterCheck(suite.testDynamicSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	suite.logonAndDisconnectAfterCheck(suite.testDynamicSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	suite.logonAndDisconnectAfterCheck(suite.testDynamicSessionID, func() {
		suite.verifySessionCount(3)
	}, true)
	suite.acceptor.Stop()
}

func (suite *AcceptorTemplateTestSuite) TestSessionNotFoundBySessionProvider() {
	sessionID := SessionID{BeginString: BeginStringFIX43, SenderCompID: "unknownSender", TargetCompID: "unknownTarget"}
	suite.logonAndDisconnectAfterCheck(sessionID, func() {}, false)
	suite.verifySessionCount(2)
	suite.acceptor.Stop()
}

func TestAcceptorTemplateTestSuite(t *testing.T) {
	suite.Run(t, new(AcceptorTemplateTestSuite))
}

type DynamicSessionTestSuite struct {
	suite.Suite
}

func (suite *DynamicSessionTestSuite) TestDynamicSession() {
	settings := NewSettings()
	settings.globalSettings.Set(config.SocketAcceptPort, "5001")
	settings.globalSettings.Set(config.DynamicSessions, "Y")
	sessionId1 := SessionID{BeginString: BeginStringFIX42, SenderCompID: "sender1", TargetCompID: "target1"}
	sessionSettings1 := NewSessionSettings()
	sessionSettings1.Set(config.BeginString, sessionId1.BeginString)
	sessionSettings1.Set(config.SenderCompID, sessionId1.SenderCompID)
	sessionSettings1.Set(config.TargetCompID, sessionId1.TargetCompID)
	settings.AddSession(sessionSettings1)

	a, err := NewAcceptor(&noopApp{}, memoryStoreFactory{}, settings, NewNullLogFactory())
	suite.Require().NoError(err, "create acceptor with DynamicSession=Y failed")

	if err := a.Start(); err != nil {
		suite.FailNow("acceptor start failed: %v", err)
	}

	inboundSessionID := SessionID{BeginString: BeginStringFIX43, SenderCompID: "X", TargetCompID: "Y"}
	inboundMessages := []*Message{mockLogonMessage(inboundSessionID, 1)}
	reversedInboundSessionID := SessionID{BeginString: BeginStringFIX43, SenderCompID: "Y", TargetCompID: "X"}

	mockConn1 := &mockConn{
		closeChan:       make(chan struct{}),
		inboundMessages: inboundMessages,
		localAddr:       &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5001},
		remoteAddr:      &net.TCPAddr{IP: net.ParseIP("127.0.0.1"), Port: 5002},
	}

	var respondedLogonMessageReceived bool
	mockConn1.onWriteback = func(_ []byte) {
		respondedLogonMessageReceived = true
		// close conn
		close(mockConn1.closeChan)
	}

	a.handleConnection(mockConn1)
	suite.Require().Equal(true, respondedLogonMessageReceived, "expected responding logon message")
	err = SendToTarget(createFIX43NewOrderSingle(), reversedInboundSessionID)
	suite.Error(err, "session created by DynamicSession is unregistered after session connected")
	a.Stop()
}

func TestDynamicSessionTestSuite(t *testing.T) {
	suite.Run(t, new(DynamicSessionTestSuite))
}
