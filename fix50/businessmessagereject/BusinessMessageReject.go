package businessmessagereject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//BusinessMessageReject is the fix50 BusinessMessageReject type, MsgType = j
type BusinessMessageReject struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a BusinessMessageReject from a quickfix.Message instance
func FromMessage(m quickfix.Message) BusinessMessageReject {
	return BusinessMessageReject{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m BusinessMessageReject) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a BusinessMessageReject initialized with the required fields for BusinessMessageReject
func New(refmsgtype field.RefMsgTypeField, businessrejectreason field.BusinessRejectReasonField) (m BusinessMessageReject) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("j"))
	m.Header.Set(field.NewBeginString("7"))
	m.Set(refmsgtype)
	m.Set(businessrejectreason)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg BusinessMessageReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "j", r
}

//SetRefSeqNum sets RefSeqNum, Tag 45
func (m BusinessMessageReject) SetRefSeqNum(v int) {
	m.Set(field.NewRefSeqNum(v))
}

//SetText sets Text, Tag 58
func (m BusinessMessageReject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m BusinessMessageReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m BusinessMessageReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetRefMsgType sets RefMsgType, Tag 372
func (m BusinessMessageReject) SetRefMsgType(v string) {
	m.Set(field.NewRefMsgType(v))
}

//SetBusinessRejectRefID sets BusinessRejectRefID, Tag 379
func (m BusinessMessageReject) SetBusinessRejectRefID(v string) {
	m.Set(field.NewBusinessRejectRefID(v))
}

//SetBusinessRejectReason sets BusinessRejectReason, Tag 380
func (m BusinessMessageReject) SetBusinessRejectReason(v int) {
	m.Set(field.NewBusinessRejectReason(v))
}

//GetRefSeqNum gets RefSeqNum, Tag 45
func (m BusinessMessageReject) GetRefSeqNum() (f field.RefSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m BusinessMessageReject) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m BusinessMessageReject) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m BusinessMessageReject) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefMsgType gets RefMsgType, Tag 372
func (m BusinessMessageReject) GetRefMsgType() (f field.RefMsgTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBusinessRejectRefID gets BusinessRejectRefID, Tag 379
func (m BusinessMessageReject) GetBusinessRejectRefID() (f field.BusinessRejectRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBusinessRejectReason gets BusinessRejectReason, Tag 380
func (m BusinessMessageReject) GetBusinessRejectReason() (f field.BusinessRejectReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRefSeqNum returns true if RefSeqNum is present, Tag 45
func (m BusinessMessageReject) HasRefSeqNum() bool {
	return m.Has(tag.RefSeqNum)
}

//HasText returns true if Text is present, Tag 58
func (m BusinessMessageReject) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m BusinessMessageReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m BusinessMessageReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasRefMsgType returns true if RefMsgType is present, Tag 372
func (m BusinessMessageReject) HasRefMsgType() bool {
	return m.Has(tag.RefMsgType)
}

//HasBusinessRejectRefID returns true if BusinessRejectRefID is present, Tag 379
func (m BusinessMessageReject) HasBusinessRejectRefID() bool {
	return m.Has(tag.BusinessRejectRefID)
}

//HasBusinessRejectReason returns true if BusinessRejectReason is present, Tag 380
func (m BusinessMessageReject) HasBusinessRejectReason() bool {
	return m.Has(tag.BusinessRejectReason)
}
