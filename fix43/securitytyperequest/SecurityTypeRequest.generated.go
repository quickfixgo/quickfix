package securitytyperequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityTypeRequest is the fix43 SecurityTypeRequest type, MsgType = v
type SecurityTypeRequest struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a SecurityTypeRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) SecurityTypeRequest {
	return SecurityTypeRequest{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityTypeRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a SecurityTypeRequest initialized with the required fields for SecurityTypeRequest
func New(securityreqid field.SecurityReqIDField) (m SecurityTypeRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("v"))
	m.Set(securityreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityTypeRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "v", r
}

//SetText sets Text, Tag 58
func (m SecurityTypeRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m SecurityTypeRequest) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityTypeRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SecurityTypeRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SecurityTypeRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m SecurityTypeRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetText gets Text, Tag 58
func (m SecurityTypeRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m SecurityTypeRequest) GetSecurityReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityTypeRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SecurityTypeRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SecurityTypeRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m SecurityTypeRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m SecurityTypeRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m SecurityTypeRequest) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SecurityTypeRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SecurityTypeRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SecurityTypeRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m SecurityTypeRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
