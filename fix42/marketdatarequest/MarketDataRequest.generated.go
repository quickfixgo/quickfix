package marketdatarequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataRequest is the fix42 MarketDataRequest type, MsgType = V
type MarketDataRequest struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a MarketDataRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) MarketDataRequest {
	return MarketDataRequest{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a MarketDataRequest initialized with the required fields for MarketDataRequest
func New(mdreqid field.MDReqIDField, subscriptionrequesttype field.SubscriptionRequestTypeField, marketdepth field.MarketDepthField) (m MarketDataRequest) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("V"))
	m.Set(mdreqid)
	m.Set(subscriptionrequesttype)
	m.Set(marketdepth)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "V", r
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m MarketDataRequest) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetMDReqID sets MDReqID, Tag 262
func (m MarketDataRequest) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m MarketDataRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetMarketDepth sets MarketDepth, Tag 264
func (m MarketDataRequest) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

//SetMDUpdateType sets MDUpdateType, Tag 265
func (m MarketDataRequest) SetMDUpdateType(v int) {
	m.Set(field.NewMDUpdateType(v))
}

//SetAggregatedBook sets AggregatedBook, Tag 266
func (m MarketDataRequest) SetAggregatedBook(v bool) {
	m.Set(field.NewAggregatedBook(v))
}

//SetNoMDEntryTypes sets NoMDEntryTypes, Tag 267
func (m MarketDataRequest) SetNoMDEntryTypes(f NoMDEntryTypesRepeatingGroup) {
	m.SetGroup(f)
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m MarketDataRequest) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMDReqID gets MDReqID, Tag 262
func (m MarketDataRequest) GetMDReqID() (f field.MDReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m MarketDataRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMarketDepth gets MarketDepth, Tag 264
func (m MarketDataRequest) GetMarketDepth() (f field.MarketDepthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDUpdateType gets MDUpdateType, Tag 265
func (m MarketDataRequest) GetMDUpdateType() (f field.MDUpdateTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAggregatedBook gets AggregatedBook, Tag 266
func (m MarketDataRequest) GetAggregatedBook() (f field.AggregatedBookField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoMDEntryTypes gets NoMDEntryTypes, Tag 267
func (m MarketDataRequest) GetNoMDEntryTypes() (NoMDEntryTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDEntryTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m MarketDataRequest) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasMDReqID returns true if MDReqID is present, Tag 262
func (m MarketDataRequest) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m MarketDataRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasMarketDepth returns true if MarketDepth is present, Tag 264
func (m MarketDataRequest) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

//HasMDUpdateType returns true if MDUpdateType is present, Tag 265
func (m MarketDataRequest) HasMDUpdateType() bool {
	return m.Has(tag.MDUpdateType)
}

//HasAggregatedBook returns true if AggregatedBook is present, Tag 266
func (m MarketDataRequest) HasAggregatedBook() bool {
	return m.Has(tag.AggregatedBook)
}

//HasNoMDEntryTypes returns true if NoMDEntryTypes is present, Tag 267
func (m MarketDataRequest) HasNoMDEntryTypes() bool {
	return m.Has(tag.NoMDEntryTypes)
}

//NoRelatedSym is a repeating group element, Tag 146
type NoRelatedSym struct {
	quickfix.Group
}

//SetSymbol sets Symbol, Tag 55
func (m NoRelatedSym) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoRelatedSym) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NoRelatedSym) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoRelatedSym) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NoRelatedSym) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoRelatedSym) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoRelatedSym) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoRelatedSym) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoRelatedSym) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoRelatedSym) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoRelatedSym) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//GetSymbol gets Symbol, Tag 55
func (m NoRelatedSym) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoRelatedSym) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoRelatedSym) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NoRelatedSym) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoRelatedSym) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoRelatedSym) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NoRelatedSym) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoRelatedSym) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoRelatedSym) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoRelatedSym) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoRelatedSym) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoRelatedSym) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoRelatedSym) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoRelatedSym) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoRelatedSym) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoRelatedSym) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoRelatedSym) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoRelatedSym) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoRelatedSym) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NoRelatedSym) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NoRelatedSym) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoRelatedSym) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoRelatedSym) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoRelatedSym) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoRelatedSym) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoRelatedSym) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoRelatedSym) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoRelatedSym) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoRelatedSym) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoRelatedSym) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoRelatedSym) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.TradingSessionID)})}
}

//Add create and append a new NoRelatedSym to this group
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

//Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}

//NoMDEntryTypes is a repeating group element, Tag 267
type NoMDEntryTypes struct {
	quickfix.Group
}

//SetMDEntryType sets MDEntryType, Tag 269
func (m NoMDEntryTypes) SetMDEntryType(v string) {
	m.Set(field.NewMDEntryType(v))
}

//GetMDEntryType gets MDEntryType, Tag 269
func (m NoMDEntryTypes) GetMDEntryType() (f field.MDEntryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasMDEntryType returns true if MDEntryType is present, Tag 269
func (m NoMDEntryTypes) HasMDEntryType() bool {
	return m.Has(tag.MDEntryType)
}

//NoMDEntryTypesRepeatingGroup is a repeating group, Tag 267
type NoMDEntryTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntryTypesRepeatingGroup returns an initialized, NoMDEntryTypesRepeatingGroup
func NewNoMDEntryTypesRepeatingGroup() NoMDEntryTypesRepeatingGroup {
	return NoMDEntryTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDEntryTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDEntryType)})}
}

//Add create and append a new NoMDEntryTypes to this group
func (m NoMDEntryTypesRepeatingGroup) Add() NoMDEntryTypes {
	g := m.RepeatingGroup.Add()
	return NoMDEntryTypes{g}
}

//Get returns the ith NoMDEntryTypes in the NoMDEntryTypesRepeatinGroup
func (m NoMDEntryTypesRepeatingGroup) Get(i int) NoMDEntryTypes {
	return NoMDEntryTypes{m.RepeatingGroup.Get(i)}
}
