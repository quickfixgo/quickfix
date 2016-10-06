package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//Reject is the fix42 Reject type, MsgType = 3
type Reject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Reject from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Reject {
	return Reject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Reject) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Reject initialized with the required fields for Reject
func New(refseqnum field.RefSeqNumField) (m Reject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("3"))
	m.Set(refseqnum)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Reject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m Reject) SetSessionRejectReason(v enum.SessionRejectReason) {
	m.Set(field.NewSessionRejectReason(v))
}

//GetRefSeqNum gets RefSeqNum, Tag 45
func (m Reject) GetRefSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.RefSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m Reject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m Reject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m Reject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefTagID gets RefTagID, Tag 371
func (m Reject) GetRefTagID() (v int, err quickfix.MessageRejectError) {
	var f field.RefTagIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefMsgType gets RefMsgType, Tag 372
func (m Reject) GetRefMsgType() (v string, err quickfix.MessageRejectError) {
	var f field.RefMsgTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSessionRejectReason gets SessionRejectReason, Tag 373
func (m Reject) GetSessionRejectReason() (v enum.SessionRejectReason, err quickfix.MessageRejectError) {
	var f field.SessionRejectReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
