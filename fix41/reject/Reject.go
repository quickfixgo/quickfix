package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//Reject is the fix41 Reject type, MsgType = 3
type Reject struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
}

//FromMessage creates a Reject from a quickfix.Message instance
func FromMessage(m quickfix.Message) Reject {
	return Reject{
		Header:  fix41.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix41.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m Reject) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a Reject initialized with the required fields for Reject
func New(refseqnum field.RefSeqNumField) (m Reject) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("3"))
	m.Set(refseqnum)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Reject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "3", r
}

//SetRefSeqNum sets RefSeqNum, Tag 45
func (m Reject) SetRefSeqNum(v int) {
	m.Set(field.NewRefSeqNum(v))
}

//SetText sets Text, Tag 58
func (m Reject) SetText(v string) {
	m.Set(field.NewText(v))
}

//GetRefSeqNum gets RefSeqNum, Tag 45
func (m Reject) GetRefSeqNum() (f field.RefSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m Reject) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRefSeqNum returns true if RefSeqNum is present, Tag 45
func (m Reject) HasRefSeqNum() bool {
	return m.Has(tag.RefSeqNum)
}

//HasText returns true if Text is present, Tag 58
func (m Reject) HasText() bool {
	return m.Has(tag.Text)
}
