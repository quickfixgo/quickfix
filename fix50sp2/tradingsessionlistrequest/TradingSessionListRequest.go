package tradingsessionlistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionListRequest is the fix50sp2 TradingSessionListRequest type, MsgType = BI
type TradingSessionListRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a TradingSessionListRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradingSessionListRequest {
	return TradingSessionListRequest{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionListRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a TradingSessionListRequest initialized with the required fields for TradingSessionListRequest
func New(tradsesreqid field.TradSesReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField) (m TradingSessionListRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BI"))
	m.Set(tradsesreqid)
	m.Set(subscriptionrequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionListRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BI", r
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradingSessionListRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradingSessionListRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionListRequest) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m TradingSessionListRequest) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionListRequest) SetTradSesMethod(v int) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionListRequest) SetTradSesMode(v int) {
	m.Set(field.NewTradSesMode(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m TradingSessionListRequest) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m TradingSessionListRequest) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetMarketID sets MarketID, Tag 1301
func (m TradingSessionListRequest) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradingSessionListRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradingSessionListRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionListRequest) GetTradSesReqID() (f field.TradSesReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionListRequest) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionListRequest) GetTradSesMethod() (f field.TradSesMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionListRequest) GetTradSesMode() (f field.TradSesModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m TradingSessionListRequest) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m TradingSessionListRequest) GetMarketSegmentID() (f field.MarketSegmentIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m TradingSessionListRequest) GetMarketID() (f field.MarketIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradingSessionListRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m TradingSessionListRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionListRequest) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m TradingSessionListRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradSesMethod returns true if TradSesMethod is present, Tag 338
func (m TradingSessionListRequest) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

//HasTradSesMode returns true if TradSesMode is present, Tag 339
func (m TradingSessionListRequest) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m TradingSessionListRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m TradingSessionListRequest) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasMarketID returns true if MarketID is present, Tag 1301
func (m TradingSessionListRequest) HasMarketID() bool {
	return m.Has(tag.MarketID)
}
