//Package testrequest msg type = 1.
package testrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//Message is a TestRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"1"`
	fix44.Header
	//TestReqID is a required field for TestRequest.
	TestReqID string `fix:"112"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized TestRequest instance
func New(testreqid string) *Message {
	var m Message
	m.SetTestReqID(testreqid)
	return &m
}

func (m *Message) SetTestReqID(v string) { m.TestReqID = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX44, "1", r
}
