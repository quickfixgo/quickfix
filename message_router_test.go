package quickfix_test

import (
	"fmt"
	"testing"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/tag"
	"github.com/stretchr/testify/suite"
)

type MessageRouterTestSuite struct {
	suite.Suite
	*quickfix.MessageRouter
	msg             quickfix.Message
	sessionID       quickfix.SessionID
	returnReject    quickfix.MessageRejectError
	routedBy        string
	routedSessionID quickfix.SessionID
	routedMessage   quickfix.Message
}

func TestMessageRouterTestSuite(t *testing.T) {
	suite.Run(t, new(MessageRouterTestSuite))
}

func (suite *MessageRouterTestSuite) givenTheRoute(beginString, msgType string) {
	suite.AddRoute(
		beginString,
		msgType,
		func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
			suite.routedBy = fmt.Sprintf("%v:%v", beginString, msgType)
			suite.routedSessionID = sessionID
			suite.routedMessage = msg

			return suite.returnReject
		},
	)
}

func (suite *MessageRouterTestSuite) givenTheMessage(msgBytes []byte) {
	msg, err := quickfix.ParseMessage(msgBytes)
	suite.Nil(err)
	suite.NotNil(msg)

	suite.msg = msg

	var beginString quickfix.FIXString
	msg.Header.GetField(tag.BeginString, &beginString)
	var senderCompID quickfix.FIXString
	msg.Header.GetField(tag.SenderCompID, &senderCompID)
	var targetCompID quickfix.FIXString
	msg.Header.GetField(tag.TargetCompID, &targetCompID)
	suite.sessionID = quickfix.SessionID{BeginString: string(beginString), SenderCompID: string(targetCompID), TargetCompID: string(senderCompID)}
}

func (suite *MessageRouterTestSuite) givenAFIX42NewOrderSingle() {
	suite.givenTheMessage([]byte("8=FIX.4.29=8735=D49=TW34=356=ISLD52=20160421-14:43:5040=160=20160421-14:43:5054=121=311=id10=235"))
}

func (suite *MessageRouterTestSuite) givenAFIXTLogonMessage() {
	suite.givenTheMessage([]byte("8=FIXT.1.19=6335=A34=149=TW52=20160420-21:21:4956=ISLD98=0108=21137=810=105"))
}

func (suite *MessageRouterTestSuite) anticipateReject(rej quickfix.MessageRejectError) {
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
	suite.MessageRouter = quickfix.NewMessageRouter()
	suite.routedBy = ""
	suite.routedSessionID = quickfix.SessionID{}
	suite.routedMessage = quickfix.Message{}
	suite.returnReject = nil
}

func (suite *MessageRouterTestSuite) SetupTest() {
	suite.resetRouter()
}

func (suite *MessageRouterTestSuite) TestNoRoute() {
	suite.givenTheMessage([]byte("8=FIX.4.39=8735=D49=TW34=356=ISLD52=20160421-14:43:5040=160=20160421-14:43:5054=121=311=id10=235"))

	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageNotRouted()
	suite.Equal(quickfix.NewBusinessMessageRejectError("Unsupported Message Type", 3, nil), rej)
}

func (suite *MessageRouterTestSuite) TestSimpleRoute() {
	suite.givenTheRoute(enum.BeginStringFIX42, "D")
	suite.givenTheRoute(enum.BeginStringFIXT11, "A")

	suite.givenAFIX42NewOrderSingle()
	rej := suite.Route(suite.msg, suite.sessionID)

	suite.verifyMessageRoutedBy(enum.BeginStringFIX42, "D")
	suite.Nil(rej)
}

func (suite *MessageRouterTestSuite) TestSimpleRouteWithReject() {
	suite.givenTheRoute(enum.BeginStringFIX42, "D")
	suite.givenTheRoute(enum.BeginStringFIXT11, "A")
	suite.anticipateReject(quickfix.NewMessageRejectError("some error", 5, nil))

	suite.givenAFIX42NewOrderSingle()
	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(enum.BeginStringFIX42, "D")
	suite.Equal(suite.returnReject, rej)
}

func (suite *MessageRouterTestSuite) TestRouteFIXTAdminMessage() {
	suite.givenTheRoute(enum.BeginStringFIX42, "D")
	suite.givenTheRoute(enum.BeginStringFIXT11, "A")
	suite.givenAFIXTLogonMessage()

	rej := suite.Route(suite.msg, suite.sessionID)
	suite.verifyMessageRoutedBy(enum.BeginStringFIXT11, "A")
	suite.Nil(rej)
}
