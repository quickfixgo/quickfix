package tradingsessionlistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fixt11"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// TradingSessionListRequest is the fix50 TradingSessionListRequest type, MsgType = BI.
type TradingSessionListRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a TradingSessionListRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) TradingSessionListRequest {
	return TradingSessionListRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m TradingSessionListRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a TradingSessionListRequest initialized with the required fields for TradingSessionListRequest.
func New(tradsesreqid field.TradSesReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m TradingSessionListRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BI"))
	m.Set(tradsesreqid)
	m.Set(subscriptionrequesttype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg TradingSessionListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "BI", r
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m TradingSessionListRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263.
func (m TradingSessionListRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

// SetTradSesReqID sets TradSesReqID, Tag 335.
func (m TradingSessionListRequest) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m TradingSessionListRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetTradSesMethod sets TradSesMethod, Tag 338.
func (m TradingSessionListRequest) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

// SetTradSesMode sets TradSesMode, Tag 339.
func (m TradingSessionListRequest) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m TradingSessionListRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m TradingSessionListRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263.
func (m TradingSessionListRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesReqID gets TradSesReqID, Tag 335.
func (m TradingSessionListRequest) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m TradingSessionListRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMethod gets TradSesMethod, Tag 338.
func (m TradingSessionListRequest) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradSesMode gets TradSesMode, Tag 339.
func (m TradingSessionListRequest) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m TradingSessionListRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m TradingSessionListRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263.
func (m TradingSessionListRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

// HasTradSesReqID returns true if TradSesReqID is present, Tag 335.
func (m TradingSessionListRequest) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m TradingSessionListRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasTradSesMethod returns true if TradSesMethod is present, Tag 338.
func (m TradingSessionListRequest) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

// HasTradSesMode returns true if TradSesMode is present, Tag 339.
func (m TradingSessionListRequest) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m TradingSessionListRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
