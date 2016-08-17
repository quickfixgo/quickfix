package crossordercancelrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//CrossOrderCancelRequest is the fix43 CrossOrderCancelRequest type, MsgType = u
type CrossOrderCancelRequest struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a CrossOrderCancelRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) CrossOrderCancelRequest {
	return CrossOrderCancelRequest{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m CrossOrderCancelRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a CrossOrderCancelRequest initialized with the required fields for CrossOrderCancelRequest
func New(crossid field.CrossIDField, origcrossid field.OrigCrossIDField, crosstype field.CrossTypeField, crossprioritization field.CrossPrioritizationField, transacttime field.TransactTimeField) (m CrossOrderCancelRequest) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("u"))
	m.Set(crossid)
	m.Set(origcrossid)
	m.Set(crosstype)
	m.Set(crossprioritization)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg CrossOrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "u", r
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m CrossOrderCancelRequest) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m CrossOrderCancelRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m CrossOrderCancelRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m CrossOrderCancelRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m CrossOrderCancelRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m CrossOrderCancelRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m CrossOrderCancelRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m CrossOrderCancelRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m CrossOrderCancelRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m CrossOrderCancelRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m CrossOrderCancelRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m CrossOrderCancelRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m CrossOrderCancelRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m CrossOrderCancelRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m CrossOrderCancelRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m CrossOrderCancelRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m CrossOrderCancelRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m CrossOrderCancelRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m CrossOrderCancelRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m CrossOrderCancelRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m CrossOrderCancelRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m CrossOrderCancelRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m CrossOrderCancelRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m CrossOrderCancelRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m CrossOrderCancelRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m CrossOrderCancelRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m CrossOrderCancelRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m CrossOrderCancelRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m CrossOrderCancelRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m CrossOrderCancelRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m CrossOrderCancelRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m CrossOrderCancelRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m CrossOrderCancelRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m CrossOrderCancelRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m CrossOrderCancelRequest) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCrossID sets CrossID, Tag 548
func (m CrossOrderCancelRequest) SetCrossID(v string) {
	m.Set(field.NewCrossID(v))
}

//SetCrossType sets CrossType, Tag 549
func (m CrossOrderCancelRequest) SetCrossType(v int) {
	m.Set(field.NewCrossType(v))
}

//SetCrossPrioritization sets CrossPrioritization, Tag 550
func (m CrossOrderCancelRequest) SetCrossPrioritization(v int) {
	m.Set(field.NewCrossPrioritization(v))
}

//SetOrigCrossID sets OrigCrossID, Tag 551
func (m CrossOrderCancelRequest) SetOrigCrossID(v string) {
	m.Set(field.NewOrigCrossID(v))
}

//SetNoSides sets NoSides, Tag 552
func (m CrossOrderCancelRequest) SetNoSides(f NoSidesRepeatingGroup) {
	m.SetGroup(f)
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m CrossOrderCancelRequest) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m CrossOrderCancelRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m CrossOrderCancelRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m CrossOrderCancelRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m CrossOrderCancelRequest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m CrossOrderCancelRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m CrossOrderCancelRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m CrossOrderCancelRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m CrossOrderCancelRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m CrossOrderCancelRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m CrossOrderCancelRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m CrossOrderCancelRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m CrossOrderCancelRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m CrossOrderCancelRequest) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m CrossOrderCancelRequest) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m CrossOrderCancelRequest) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m CrossOrderCancelRequest) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m CrossOrderCancelRequest) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m CrossOrderCancelRequest) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m CrossOrderCancelRequest) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m CrossOrderCancelRequest) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m CrossOrderCancelRequest) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m CrossOrderCancelRequest) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m CrossOrderCancelRequest) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m CrossOrderCancelRequest) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m CrossOrderCancelRequest) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m CrossOrderCancelRequest) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m CrossOrderCancelRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m CrossOrderCancelRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m CrossOrderCancelRequest) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m CrossOrderCancelRequest) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m CrossOrderCancelRequest) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m CrossOrderCancelRequest) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m CrossOrderCancelRequest) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m CrossOrderCancelRequest) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCrossID gets CrossID, Tag 548
func (m CrossOrderCancelRequest) GetCrossID() (f field.CrossIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCrossType gets CrossType, Tag 549
func (m CrossOrderCancelRequest) GetCrossType() (f field.CrossTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCrossPrioritization gets CrossPrioritization, Tag 550
func (m CrossOrderCancelRequest) GetCrossPrioritization() (f field.CrossPrioritizationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigCrossID gets OrigCrossID, Tag 551
func (m CrossOrderCancelRequest) GetOrigCrossID() (f field.OrigCrossIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSides gets NoSides, Tag 552
func (m CrossOrderCancelRequest) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m CrossOrderCancelRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m CrossOrderCancelRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m CrossOrderCancelRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m CrossOrderCancelRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m CrossOrderCancelRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m CrossOrderCancelRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m CrossOrderCancelRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m CrossOrderCancelRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m CrossOrderCancelRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m CrossOrderCancelRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m CrossOrderCancelRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m CrossOrderCancelRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m CrossOrderCancelRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m CrossOrderCancelRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m CrossOrderCancelRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m CrossOrderCancelRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m CrossOrderCancelRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m CrossOrderCancelRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m CrossOrderCancelRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m CrossOrderCancelRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m CrossOrderCancelRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m CrossOrderCancelRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m CrossOrderCancelRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m CrossOrderCancelRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m CrossOrderCancelRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m CrossOrderCancelRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m CrossOrderCancelRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m CrossOrderCancelRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m CrossOrderCancelRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m CrossOrderCancelRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m CrossOrderCancelRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m CrossOrderCancelRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m CrossOrderCancelRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m CrossOrderCancelRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m CrossOrderCancelRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCrossID returns true if CrossID is present, Tag 548
func (m CrossOrderCancelRequest) HasCrossID() bool {
	return m.Has(tag.CrossID)
}

//HasCrossType returns true if CrossType is present, Tag 549
func (m CrossOrderCancelRequest) HasCrossType() bool {
	return m.Has(tag.CrossType)
}

//HasCrossPrioritization returns true if CrossPrioritization is present, Tag 550
func (m CrossOrderCancelRequest) HasCrossPrioritization() bool {
	return m.Has(tag.CrossPrioritization)
}

//HasOrigCrossID returns true if OrigCrossID is present, Tag 551
func (m CrossOrderCancelRequest) HasOrigCrossID() bool {
	return m.Has(tag.OrigCrossID)
}

//HasNoSides returns true if NoSides is present, Tag 552
func (m CrossOrderCancelRequest) HasNoSides() bool {
	return m.Has(tag.NoSides)
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

//NoSides is a repeating group element, Tag 552
type NoSides struct {
	quickfix.Group
}

//SetSide sets Side, Tag 54
func (m NoSides) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m NoSides) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoSides) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoSides) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m NoSides) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetOrigOrdModTime sets OrigOrdModTime, Tag 586
func (m NoSides) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSides) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m NoSides) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NoSides) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NoSides) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m NoSides) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m NoSides) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m NoSides) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m NoSides) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetText sets Text, Tag 58
func (m NoSides) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoSides) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoSides) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetSide gets Side, Tag 54
func (m NoSides) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m NoSides) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoSides) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoSides) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m NoSides) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigOrdModTime gets OrigOrdModTime, Tag 586
func (m NoSides) GetOrigOrdModTime() (f field.OrigOrdModTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSides) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m NoSides) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NoSides) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NoSides) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m NoSides) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m NoSides) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m NoSides) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NoSides) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoSides) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoSides) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoSides) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSide returns true if Side is present, Tag 54
func (m NoSides) HasSide() bool {
	return m.Has(tag.Side)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m NoSides) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoSides) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoSides) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m NoSides) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586
func (m NoSides) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSides) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m NoSides) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NoSides) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NoSides) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m NoSides) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m NoSides) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m NoSides) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m NoSides) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasText returns true if Text is present, Tag 58
func (m NoSides) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoSides) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoSides) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
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

//NoSidesRepeatingGroup is a repeating group, Tag 552
type NoSidesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSidesRepeatingGroup returns an initialized, NoSidesRepeatingGroup
func NewNoSidesRepeatingGroup() NoSidesRepeatingGroup {
	return NoSidesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSides,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.OrigClOrdID), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ClOrdLinkID), quickfix.GroupElement(tag.OrigOrdModTime), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.TradeOriginationDate), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.CashOrderQty), quickfix.GroupElement(tag.OrderPercent), quickfix.GroupElement(tag.RoundingDirection), quickfix.GroupElement(tag.RoundingModulus), quickfix.GroupElement(tag.ComplianceID), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
}

//Add create and append a new NoSides to this group
func (m NoSidesRepeatingGroup) Add() NoSides {
	g := m.RepeatingGroup.Add()
	return NoSides{g}
}

//Get returns the ith NoSides in the NoSidesRepeatinGroup
func (m NoSidesRepeatingGroup) Get(i int) NoSides {
	return NoSides{m.RepeatingGroup.Get(i)}
}
