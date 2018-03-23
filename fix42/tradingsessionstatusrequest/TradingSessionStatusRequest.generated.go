package tradingsessionstatusrequest

import (
	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/enum"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fix42"
	"github.com/alpacahq/quickfix/tag"
)

//TradingSessionStatusRequest is the fix42 TradingSessionStatusRequest type, MsgType = g
type TradingSessionStatusRequest struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a TradingSessionStatusRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TradingSessionStatusRequest {
	return TradingSessionStatusRequest{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a TradingSessionStatusRequest initialized with the required fields for TradingSessionStatusRequest
func New(tradsesreqid field.TradSesReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m TradingSessionStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("g"))
	m.Set(tradsesreqid)
	m.Set(subscriptionrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "g", r
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradingSessionStatusRequest) SetSubscriptionRequestType(v enum.SubscriptionRequestType) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionStatusRequest) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m TradingSessionStatusRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionStatusRequest) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionStatusRequest) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradingSessionStatusRequest) GetSubscriptionRequestType() (v enum.SubscriptionRequestType, err quickfix.MessageRejectError) {
	var f field.SubscriptionRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionStatusRequest) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionStatusRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionStatusRequest) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionStatusRequest) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m TradingSessionStatusRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionStatusRequest) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m TradingSessionStatusRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradSesMethod returns true if TradSesMethod is present, Tag 338
func (m TradingSessionStatusRequest) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

//HasTradSesMode returns true if TradSesMode is present, Tag 339
func (m TradingSessionStatusRequest) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}
