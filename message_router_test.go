package quickfix

import (
	"bytes"
	"fmt"
	"testing"

	"github.com/stretchr/testify/suite"
)

type MessageRouterTestSuite struct {
	suite.Suite
	*MessageRouter
	msg             *Message
	sessionID       SessionID
	returnReject    MessageRejectError
	routedBy        string
	routedSessionID SessionID
	routedMessage   *Message
}

func TestMessageRouterTestSuite(t *testing.T) {
	suite.Run(t, new(MessageRouterTestSuite))
}

func (suite *MessageRouterTestSuite) givenTheRoute(beginString, msgType string) {
	suite.AddRoute(
		beginString,
		msgType,
		func(msg *Message, sessionID SessionID) MessageRejectError {
			suite.routedBy = fmt.Sprintf("%v:%v", beginString, msgType)
			suite.routedSessionID = sessionID
			suite.routedMessage = msg

			return suite.returnReject
		},
	)
}

func (suite *MessageRouterTestSuite) givenTheMessage(msgBytes []byte) {
	err := ParseMessage(suite.msg, bytes.NewBuffer(msgBytes))
	suite.Nil(err)

	var beginString FIXString
	suite.Require().Nil(suite.msg.Header.GetField(tagBeginString, &beginString))
	var senderCompID FIXString
	suite.Require().Nil(suite.msg.Header.GetField(tagSenderCompID, &senderCompID))
	var targetCompID FIXString
	suite.Require().Nil(suite.msg.Header.GetField(tagTargetCompID, &targetCompID))
	suite.sessionID = SessionID{BeginString: string(beginString), SenderCompID: string(targetCompID), TargetCompID: string(senderCompID)}
}

func (suite *MessageRouterTestSuite) givenTargetDefaultApplVerIDForSession(defaultApplVerID string, sessionID SessionID) {
	s := &session{
		sessionID:              sessionID,
		targetDefaultApplVerID: defaultApplVerID,
	}
	suite.Nil(registerSession(s))
}

func (suite *MessageRouterTestSuite) givenAFIX42NewOrderSingle() {
	suite.givenTheMessage([]byte("8=FIX.4.29=8735=D49=TW34=356=ISLD52=20160421-14:43:5040=160=20160421-14:43:5054=121=311=id10=235"))
}

func (suite *MessageRouterTestSuite) givenAFIXTLogonMessage() {
	suite.givenTheMessage([]byte("8=FIXT.1.19=6335=A34=149=TW52=20160420-21:21:4956=ISLD98=0108=21137=810=105"))
}

func (suite *MessageRouterTestSuite) anticipateReject(rej MessageRejectError) {
	suite.returnReject = rej
}

func (suite *MessageRouterTestSuite) verifyMessageNotRouted() {
	suite.Equal("", suite.routedBy, "Message should not be routed")
}

func (suite *MessageRouterTestSuite) verifyMessageRoutedBy(beginString, msgType string) {
	suite.NotEqual("", suite.routedBy, "Message expected to be routed")

	suite.Equal(fmt.Sprintf("%v:%v", beginString, msgType), suite.routedBy)
	suite.Equal(suite.sessionID, suite.routedSessionID)
	suite.Equal(suite.msg.String(), suite.routedMessage.String())
}

func (suite *MessageRouterTestSuite) resetRouter() {
	suite.MessageRouter = NewMessageRouter()
	suite.routedBy = ""
	suite.routedSessionID = SessionID{}
	suite.routedMessage = &Message{}
	suite.returnReject = nil
}

func (suite *MessageRouterTestSuite) SetupTest() {
	suite.resetRouter()
	sessionsLock.Lock()
	defer sessionsLock.Unlock()

	sessions = make(map[SessionID]*session)
	suite.msg = NewMessage()
}

func (suite *MessageRouterTestSuite) TestNoRoute() {
	suite.givenTheMessage([]byte("8=FIX.4.39=8735=D49=TW34=356=ISLD52=20160421-14:43:5040=160=20160421-14:43:5054=121=311=id10=235"))

	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageNotRouted()
	suite.Equal(NewBusinessMessageRejectError("Unsupported Message Type", 3, nil), rej)
}

func (suite *MessageRouterTestSuite) TestNoRouteWhitelistedMessageTypes() {
	var tests = []string{"0", "A", "1", "2", "3", "4", "5", "j"}

	for _, test := range tests {
		suite.SetupTest()

		msg := fmt.Sprintf("8=FIX.4.39=8735=%v49=TW34=356=ISLD52=20160421-14:43:5040=160=20160421-14:43:5054=121=311=id10=235", test)
		suite.givenTheMessage([]byte(msg))

		rej := suite.Route(suite.msg, suite.sessionID)
		suite.verifyMessageNotRouted()
		suite.Nil(rej, "Message type '%v' should not be rejected by the MessageRouter", test)
	}
}

func (suite *MessageRouterTestSuite) TestSimpleRoute() {
	suite.givenTheRoute(string(BeginStringFIX42), "D")
	suite.givenTheRoute(string(BeginStringFIXT11), "A")

	suite.givenAFIX42NewOrderSingle()
	rej := suite.Route(suite.msg, suite.sessionID)

	suite.verifyMessageRoutedBy(string(BeginStringFIX42), "D")
	suite.Nil(rej)
}

func (suite *MessageRouterTestSuite) TestSimpleRouteWithReject() {
	suite.givenTheRoute(string(BeginStringFIX42), "D")
	suite.givenTheRoute(string(BeginStringFIXT11), "A")
	suite.anticipateReject(NewMessageRejectError("some error", 5, nil))

	suite.givenAFIX42NewOrderSingle()
	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(string(BeginStringFIX42), "D")
	suite.Equal(suite.returnReject, rej)
}

func (suite *MessageRouterTestSuite) TestRouteFIXTAdminMessage() {
	suite.givenTheRoute(string(BeginStringFIX42), "D")
	suite.givenTheRoute(string(BeginStringFIXT11), "A")
	suite.givenAFIXTLogonMessage()

	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(string(BeginStringFIXT11), "A")
	suite.Nil(rej)
}

func (suite *MessageRouterTestSuite) TestRouteFIXT50AppWithApplVerID() {
	suite.givenTheRoute(BeginStringFIX42, "D")
	suite.givenTheRoute(ApplVerIDFIX50, "D")
	suite.givenTheRoute(ApplVerIDFIX50SP1, "D")

	suite.givenTheMessage([]byte("8=FIXT.1.19=8935=D49=TW34=356=ISLD52=20160424-16:48:261128=740=160=20160424-16:48:2611=id21=310=120"))
	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(ApplVerIDFIX50, "D")
	suite.Nil(rej)
}

func (suite *MessageRouterTestSuite) TestRouteFIXTAppWithApplVerID() {
	suite.givenTheRoute(BeginStringFIX42, "D")
	suite.givenTheRoute(ApplVerIDFIX50, "D")
	suite.givenTheRoute(ApplVerIDFIX50SP1, "D")

	suite.givenTheMessage([]byte("8=FIXT.1.19=8935=D49=TW34=356=ISLD52=20160424-16:48:261128=840=160=20160424-16:48:2611=id21=310=120"))
	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(ApplVerIDFIX50SP1, "D")
	suite.Nil(rej)
}

func (suite *MessageRouterTestSuite) TestRouteFIXTAppWithDefaultApplVerID() {
	suite.givenTheRoute(BeginStringFIX42, "D")
	suite.givenTheRoute(ApplVerIDFIX50, "D")
	suite.givenTheRoute(ApplVerIDFIX50SP1, "D")
	suite.givenTargetDefaultApplVerIDForSession(
		"8",
		SessionID{BeginString: string(BeginStringFIXT11), SenderCompID: "ISLD", TargetCompID: "TW"},
	)

	suite.givenTheMessage([]byte("8=FIXT.1.19=8235=D49=TW34=356=ISLD52=20160424-16:48:2640=160=20160424-16:48:2611=id21=310=120"))
	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(ApplVerIDFIX50SP1, "D")
	suite.Nil(rej)
}
