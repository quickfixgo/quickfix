//Package resendrequest msg type = 2.
package resendrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
)

//Message is a ResendRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"2"`
	fix43.Header
	//BeginSeqNo is a required field for ResendRequest.
	BeginSeqNo int `fix:"7"`
	//EndSeqNo is a required field for ResendRequest.
	EndSeqNo int `fix:"16"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized ResendRequest instance
func New(beginseqno int, endseqno int) *Message {
	var m Message
	m.SetBeginSeqNo(beginseqno)
	m.SetEndSeqNo(endseqno)
	return &m
}

func (m *Message) SetBeginSeqNo(v int) { m.BeginSeqNo = v }
func (m *Message) SetEndSeqNo(v int)   { m.EndSeqNo = v }

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
	return enum.BeginStringFIX43, "2", r
}
