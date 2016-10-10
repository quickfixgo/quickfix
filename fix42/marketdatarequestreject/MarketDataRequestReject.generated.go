package marketdatarequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataRequestReject is the fix42 MarketDataRequestReject type, MsgType = Y
type MarketDataRequestReject struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a MarketDataRequestReject from a quickfix.Message instance
func FromMessage(m *quickfix.Message) MarketDataRequestReject {
	return MarketDataRequestReject{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataRequestReject) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MarketDataRequestReject initialized with the required fields for MarketDataRequestReject
func New(mdreqid field.MDReqIDField) (m MarketDataRequestReject) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Y"))
	m.Set(mdreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "Y", r
}

//SetText sets Text, Tag 58
func (m MarketDataRequestReject) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetMDReqID sets MDReqID, Tag 262
func (m MarketDataRequestReject) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

//SetMDReqRejReason sets MDReqRejReason, Tag 281
func (m MarketDataRequestReject) SetMDReqRejReason(v enum.MDReqRejReason) {
	m.Set(field.NewMDReqRejReason(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m MarketDataRequestReject) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m MarketDataRequestReject) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetText gets Text, Tag 58
func (m MarketDataRequestReject) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDReqID gets MDReqID, Tag 262
func (m MarketDataRequestReject) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDReqRejReason gets MDReqRejReason, Tag 281
func (m MarketDataRequestReject) GetMDReqRejReason() (v enum.MDReqRejReason, err quickfix.MessageRejectError) {
	var f field.MDReqRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m MarketDataRequestReject) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m MarketDataRequestReject) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m MarketDataRequestReject) HasText() bool {
	return m.Has(tag.Text)
}

//HasMDReqID returns true if MDReqID is present, Tag 262
func (m MarketDataRequestReject) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

//HasMDReqRejReason returns true if MDReqRejReason is present, Tag 281
func (m MarketDataRequestReject) HasMDReqRejReason() bool {
	return m.Has(tag.MDReqRejReason)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m MarketDataRequestReject) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m MarketDataRequestReject) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}
