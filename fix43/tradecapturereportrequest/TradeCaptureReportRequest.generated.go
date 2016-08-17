package tradecapturereportrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//TradeCaptureReportRequest is the fix43 TradeCaptureReportRequest type, MsgType = AD
type TradeCaptureReportRequest struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradeCaptureReportRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradeCaptureReportRequest {
	return TradeCaptureReportRequest{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradeCaptureReportRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradeCaptureReportRequest initialized with the required fields for TradeCaptureReportRequest
func New(traderequestid field.TradeRequestIDField, traderequesttype field.TradeRequestTypeField) (m TradeCaptureReportRequest) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AD"))
	m.Set(traderequestid)
	m.Set(traderequesttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradeCaptureReportRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "AD", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m TradeCaptureReportRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetExecID sets ExecID, Tag 17
func (m TradeCaptureReportRequest) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradeCaptureReportRequest) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m TradeCaptureReportRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m TradeCaptureReportRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m TradeCaptureReportRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m TradeCaptureReportRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m TradeCaptureReportRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradeCaptureReportRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m TradeCaptureReportRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m TradeCaptureReportRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m TradeCaptureReportRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradeCaptureReportRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m TradeCaptureReportRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m TradeCaptureReportRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradeCaptureReportRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m TradeCaptureReportRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m TradeCaptureReportRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m TradeCaptureReportRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m TradeCaptureReportRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m TradeCaptureReportRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m TradeCaptureReportRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradeCaptureReportRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReportRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m TradeCaptureReportRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m TradeCaptureReportRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetSubscriptionRequestType sets SubscriptionRequestType, Tag 263
func (m TradeCaptureReportRequest) SetSubscriptionRequestType(v string) {
	m.Set(field.NewSubscriptionRequestType(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m TradeCaptureReportRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m TradeCaptureReportRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReportRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReportRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m TradeCaptureReportRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m TradeCaptureReportRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m TradeCaptureReportRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradeCaptureReportRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m TradeCaptureReportRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m TradeCaptureReportRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m TradeCaptureReportRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReportRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m TradeCaptureReportRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m TradeCaptureReportRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m TradeCaptureReportRequest) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetTradeRequestID sets TradeRequestID, Tag 568
func (m TradeCaptureReportRequest) SetTradeRequestID(v string) {
	m.Set(field.NewTradeRequestID(v))
}

//SetTradeRequestType sets TradeRequestType, Tag 569
func (m TradeCaptureReportRequest) SetTradeRequestType(v int) {
	m.Set(field.NewTradeRequestType(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m TradeCaptureReportRequest) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetTradeInputSource sets TradeInputSource, Tag 578
func (m TradeCaptureReportRequest) SetTradeInputSource(v string) {
	m.Set(field.NewTradeInputSource(v))
}

//SetTradeInputDevice sets TradeInputDevice, Tag 579
func (m TradeCaptureReportRequest) SetTradeInputDevice(v string) {
	m.Set(field.NewTradeInputDevice(v))
}

//SetNoDates sets NoDates, Tag 580
func (m TradeCaptureReportRequest) SetNoDates(f NoDatesRepeatingGroup) {
	m.SetGroup(f)
}

//GetClOrdID gets ClOrdID, Tag 11
func (m TradeCaptureReportRequest) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m TradeCaptureReportRequest) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradeCaptureReportRequest) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m TradeCaptureReportRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m TradeCaptureReportRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m TradeCaptureReportRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m TradeCaptureReportRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m TradeCaptureReportRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradeCaptureReportRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m TradeCaptureReportRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradeCaptureReportRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m TradeCaptureReportRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradeCaptureReportRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m TradeCaptureReportRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m TradeCaptureReportRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradeCaptureReportRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m TradeCaptureReportRequest) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradeCaptureReportRequest) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m TradeCaptureReportRequest) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradeCaptureReportRequest) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradeCaptureReportRequest) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m TradeCaptureReportRequest) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradeCaptureReportRequest) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReportRequest) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradeCaptureReportRequest) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m TradeCaptureReportRequest) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSubscriptionRequestType gets SubscriptionRequestType, Tag 263
func (m TradeCaptureReportRequest) GetSubscriptionRequestType() (f field.SubscriptionRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradeCaptureReportRequest) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradeCaptureReportRequest) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReportRequest) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReportRequest) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m TradeCaptureReportRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m TradeCaptureReportRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m TradeCaptureReportRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradeCaptureReportRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m TradeCaptureReportRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m TradeCaptureReportRequest) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradeCaptureReportRequest) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReportRequest) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradeCaptureReportRequest) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m TradeCaptureReportRequest) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradeCaptureReportRequest) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeRequestID gets TradeRequestID, Tag 568
func (m TradeCaptureReportRequest) GetTradeRequestID() (f field.TradeRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeRequestType gets TradeRequestType, Tag 569
func (m TradeCaptureReportRequest) GetTradeRequestType() (f field.TradeRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m TradeCaptureReportRequest) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputSource gets TradeInputSource, Tag 578
func (m TradeCaptureReportRequest) GetTradeInputSource() (f field.TradeInputSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputDevice gets TradeInputDevice, Tag 579
func (m TradeCaptureReportRequest) GetTradeInputDevice() (f field.TradeInputDeviceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoDates gets NoDates, Tag 580
func (m TradeCaptureReportRequest) GetNoDates() (NoDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m TradeCaptureReportRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m TradeCaptureReportRequest) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m TradeCaptureReportRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m TradeCaptureReportRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m TradeCaptureReportRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m TradeCaptureReportRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m TradeCaptureReportRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m TradeCaptureReportRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m TradeCaptureReportRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m TradeCaptureReportRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m TradeCaptureReportRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m TradeCaptureReportRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m TradeCaptureReportRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m TradeCaptureReportRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m TradeCaptureReportRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradeCaptureReportRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m TradeCaptureReportRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m TradeCaptureReportRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m TradeCaptureReportRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m TradeCaptureReportRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m TradeCaptureReportRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m TradeCaptureReportRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m TradeCaptureReportRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m TradeCaptureReportRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m TradeCaptureReportRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m TradeCaptureReportRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasSubscriptionRequestType returns true if SubscriptionRequestType is present, Tag 263
func (m TradeCaptureReportRequest) HasSubscriptionRequestType() bool {
	return m.Has(tag.SubscriptionRequestType)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m TradeCaptureReportRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m TradeCaptureReportRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m TradeCaptureReportRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m TradeCaptureReportRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m TradeCaptureReportRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m TradeCaptureReportRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m TradeCaptureReportRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m TradeCaptureReportRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m TradeCaptureReportRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m TradeCaptureReportRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m TradeCaptureReportRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m TradeCaptureReportRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m TradeCaptureReportRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m TradeCaptureReportRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m TradeCaptureReportRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasTradeRequestID returns true if TradeRequestID is present, Tag 568
func (m TradeCaptureReportRequest) HasTradeRequestID() bool {
	return m.Has(tag.TradeRequestID)
}

//HasTradeRequestType returns true if TradeRequestType is present, Tag 569
func (m TradeCaptureReportRequest) HasTradeRequestType() bool {
	return m.Has(tag.TradeRequestType)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m TradeCaptureReportRequest) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasTradeInputSource returns true if TradeInputSource is present, Tag 578
func (m TradeCaptureReportRequest) HasTradeInputSource() bool {
	return m.Has(tag.TradeInputSource)
}

//HasTradeInputDevice returns true if TradeInputDevice is present, Tag 579
func (m TradeCaptureReportRequest) HasTradeInputDevice() bool {
	return m.Has(tag.TradeInputDevice)
}

//HasNoDates returns true if NoDates is present, Tag 580
func (m TradeCaptureReportRequest) HasNoDates() bool {
	return m.Has(tag.NoDates)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoDates is a repeating group element, Tag 580
type NoDates struct {
	quickfix.Group
}

//SetTradeDate sets TradeDate, Tag 75
func (m NoDates) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m NoDates) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//GetTradeDate gets TradeDate, Tag 75
func (m NoDates) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m NoDates) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m NoDates) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m NoDates) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//NoDatesRepeatingGroup is a repeating group, Tag 580
type NoDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoDatesRepeatingGroup returns an initialized, NoDatesRepeatingGroup
func NewNoDatesRepeatingGroup() NoDatesRepeatingGroup {
	return NoDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradeDate), quickfix.GroupElement(tag.TransactTime)})}
}

//Add create and append a new NoDates to this group
func (m NoDatesRepeatingGroup) Add() NoDates {
	g := m.RepeatingGroup.Add()
	return NoDates{g}
}

//Get returns the ith NoDates in the NoDatesRepeatinGroup
func (m NoDatesRepeatingGroup) Get(i int) NoDates {
	return NoDates{m.RepeatingGroup.Get(i)}
}
