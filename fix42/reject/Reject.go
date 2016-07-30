package reject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//Reject is the fix42 Reject type, MsgType = 3
type Reject struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Reject from a quickfix.Message instance
func FromMessage(m quickfix.Message) Reject {
	return Reject{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Reject) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Reject initialized with the required fields for Reject
func New(refseqnum field.RefSeqNumField) (m Reject) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("3"))
	m.Header.Set(field.NewBeginString("FIX.4.2"))
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
	return "FIX.4.2", "3", r
}

//SetRefSeqNum sets RefSeqNum, Tag 45
func (m Reject) SetRefSeqNum(v int) {
	m.Set(field.NewRefSeqNum(v))
}

//SetText sets Text, Tag 58
func (m Reject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m Reject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m Reject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetRefTagID sets RefTagID, Tag 371
func (m Reject) SetRefTagID(v int) {
	m.Set(field.NewRefTagID(v))
}

//SetRefMsgType sets RefMsgType, Tag 372
func (m Reject) SetRefMsgType(v string) {
	m.Set(field.NewRefMsgType(v))
}

//SetSessionRejectReason sets SessionRejectReason, Tag 373
func (m Reject) SetSessionRejectReason(v int) {
	m.Set(field.NewSessionRejectReason(v))
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

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m Reject) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m Reject) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefTagID gets RefTagID, Tag 371
func (m Reject) GetRefTagID() (f field.RefTagIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefMsgType gets RefMsgType, Tag 372
func (m Reject) GetRefMsgType() (f field.RefMsgTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSessionRejectReason gets SessionRejectReason, Tag 373
func (m Reject) GetSessionRejectReason() (f field.SessionRejectReasonField, err quickfix.MessageRejectError) {
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

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m Reject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m Reject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasRefTagID returns true if RefTagID is present, Tag 371
func (m Reject) HasRefTagID() bool {
	return m.Has(tag.RefTagID)
}

//HasRefMsgType returns true if RefMsgType is present, Tag 372
func (m Reject) HasRefMsgType() bool {
	return m.Has(tag.RefMsgType)
}

//HasSessionRejectReason returns true if SessionRejectReason is present, Tag 373
func (m Reject) HasSessionRejectReason() bool {
	return m.Has(tag.SessionRejectReason)
}
