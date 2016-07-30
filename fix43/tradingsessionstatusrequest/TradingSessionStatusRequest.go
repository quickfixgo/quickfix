package tradingsessionstatusrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionStatusRequest is the fix43 TradingSessionStatusRequest type, MsgType = g
type TradingSessionStatusRequest struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradingSessionStatusRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradingSessionStatusRequest {
	return TradingSessionStatusRequest{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionStatusRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradingSessionStatusRequest initialized with the required fields for TradingSessionStatusRequest
func New(tradsesreqid field.TradSesReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m TradingSessionStatusRequest) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("g"))
	m.Header.Set(field.NewBeginString("FIX.4.3"))
	m.Set(tradsesreqid)
	m.Set(subscriptionrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "g", r
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradingSessionStatusRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionStatusRequest) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m TradingSessionStatusRequest) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionStatusRequest) SetTradSesMethod(v int) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionStatusRequest) SetTradSesMode(v int) {
	m.Set(field.NewTradSesMode(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m TradingSessionStatusRequest) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradingSessionStatusRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionStatusRequest) GetTradSesReqID() (f field.TradSesReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionStatusRequest) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionStatusRequest) GetTradSesMethod() (f field.TradSesMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionStatusRequest) GetTradSesMode() (f field.TradSesModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m TradingSessionStatusRequest) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m TradingSessionStatusRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
