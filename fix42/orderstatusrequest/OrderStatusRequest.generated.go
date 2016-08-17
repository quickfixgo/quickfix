package orderstatusrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderStatusRequest is the fix42 OrderStatusRequest type, MsgType = H
type OrderStatusRequest struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderStatusRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderStatusRequest {
	return OrderStatusRequest{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderStatusRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a OrderStatusRequest initialized with the required fields for OrderStatusRequest
func New(clordid field.ClOrdIDField, symbol field.SymbolField, side field.SideField) (m OrderStatusRequest) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("H"))
	m.Set(clordid)
	m.Set(symbol)
	m.Set(side)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "H", r
}

//SetAccount sets Account, Tag 1
func (m OrderStatusRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderStatusRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m OrderStatusRequest) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderStatusRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m OrderStatusRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m OrderStatusRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m OrderStatusRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderStatusRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m OrderStatusRequest) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetIssuer sets Issuer, Tag 106
func (m OrderStatusRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m OrderStatusRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m OrderStatusRequest) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m OrderStatusRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m OrderStatusRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m OrderStatusRequest) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m OrderStatusRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m OrderStatusRequest) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m OrderStatusRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m OrderStatusRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m OrderStatusRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m OrderStatusRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m OrderStatusRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m OrderStatusRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m OrderStatusRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m OrderStatusRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//GetAccount gets Account, Tag 1
func (m OrderStatusRequest) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderStatusRequest) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m OrderStatusRequest) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderStatusRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m OrderStatusRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m OrderStatusRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m OrderStatusRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderStatusRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m OrderStatusRequest) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m OrderStatusRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderStatusRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m OrderStatusRequest) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m OrderStatusRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderStatusRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m OrderStatusRequest) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m OrderStatusRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m OrderStatusRequest) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m OrderStatusRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderStatusRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m OrderStatusRequest) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m OrderStatusRequest) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m OrderStatusRequest) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m OrderStatusRequest) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m OrderStatusRequest) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m OrderStatusRequest) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m OrderStatusRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderStatusRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m OrderStatusRequest) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderStatusRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m OrderStatusRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m OrderStatusRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m OrderStatusRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m OrderStatusRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m OrderStatusRequest) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m OrderStatusRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m OrderStatusRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m OrderStatusRequest) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m OrderStatusRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m OrderStatusRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m OrderStatusRequest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m OrderStatusRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m OrderStatusRequest) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m OrderStatusRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m OrderStatusRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m OrderStatusRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m OrderStatusRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m OrderStatusRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m OrderStatusRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m OrderStatusRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m OrderStatusRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}
