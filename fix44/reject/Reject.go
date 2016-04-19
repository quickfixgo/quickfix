//Package reject msg type = 3.
package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//Message is a Reject FIX Message
type Message struct {
	FIXMsgType string `fix:"3"`
	fix44.Header
	//RefSeqNum is a required field for Reject.
	RefSeqNum int `fix:"45"`
	//RefTagID is a non-required field for Reject.
	RefTagID *int `fix:"371"`
	//RefMsgType is a non-required field for Reject.
	RefMsgType *string `fix:"372"`
	//SessionRejectReason is a non-required field for Reject.
	SessionRejectReason *int `fix:"373"`
	//Text is a non-required field for Reject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Reject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Reject.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized Reject instance
func New(refseqnum int) *Message {
	var m Message
	m.SetRefSeqNum(refseqnum)
	return &m
}

func (m *Message) SetRefSeqNum(v int)           { m.RefSeqNum = v }
func (m *Message) SetRefTagID(v int)            { m.RefTagID = &v }
func (m *Message) SetRefMsgType(v string)       { m.RefMsgType = &v }
func (m *Message) SetSessionRejectReason(v int) { m.SessionRejectReason = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "3", r
}
