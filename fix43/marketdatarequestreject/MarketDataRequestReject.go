package marketdatarequestreject

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataRequestReject is the fix43 MarketDataRequestReject type, MsgType = Y
type MarketDataRequestReject struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a MarketDataRequestReject from a quickfix.Message instance
func FromMessage(m quickfix.Message) MarketDataRequestReject {
	return MarketDataRequestReject{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataRequestReject) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a MarketDataRequestReject initialized with the required fields for MarketDataRequestReject
func New(mdreqid field.MDReqIDField) (m MarketDataRequestReject) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("Y"))
	m.Set(mdreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataRequestReject, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "Y", r
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
func (m MarketDataRequestReject) SetMDReqRejReason(v string) {
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
func (m MarketDataRequestReject) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDReqID gets MDReqID, Tag 262
func (m MarketDataRequestReject) GetMDReqID() (f field.MDReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDReqRejReason gets MDReqRejReason, Tag 281
func (m MarketDataRequestReject) GetMDReqRejReason() (f field.MDReqRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m MarketDataRequestReject) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m MarketDataRequestReject) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
