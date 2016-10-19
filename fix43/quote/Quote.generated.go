package quote

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//Quote is the fix43 Quote type, MsgType = S
type Quote struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Quote from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Quote {
	return Quote{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Quote) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Quote initialized with the required fields for Quote
func New(quoteid field.QuoteIDField) (m Quote) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("S"))
	m.Set(quoteid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Quote, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "S", r
}

//SetAccount sets Account, Tag 1
func (m Quote) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetCommission sets Commission, Tag 12
func (m Quote) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m Quote) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m Quote) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m Quote) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrdType sets OrdType, Tag 40
func (m Quote) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Quote) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Quote) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m Quote) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Quote) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m Quote) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m Quote) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Quote) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Quote) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m Quote) SetExDestination(v enum.ExDestination) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m Quote) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Quote) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m Quote) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m Quote) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetBidPx sets BidPx, Tag 132
func (m Quote) SetBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidPx(value, scale))
}

//SetOfferPx sets OfferPx, Tag 133
func (m Quote) SetOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferPx(value, scale))
}

//SetBidSize sets BidSize, Tag 134
func (m Quote) SetBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSize(value, scale))
}

//SetOfferSize sets OfferSize, Tag 135
func (m Quote) SetOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSize(value, scale))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m Quote) SetSettlCurrFxRateCalc(v enum.SettlCurrFxRateCalc) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m Quote) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetBidSpotRate sets BidSpotRate, Tag 188
func (m Quote) SetBidSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSpotRate(value, scale))
}

//SetBidForwardPoints sets BidForwardPoints, Tag 189
func (m Quote) SetBidForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidForwardPoints(value, scale))
}

//SetOfferSpotRate sets OfferSpotRate, Tag 190
func (m Quote) SetOfferSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSpotRate(value, scale))
}

//SetOfferForwardPoints sets OfferForwardPoints, Tag 191
func (m Quote) SetOfferForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferForwardPoints(value, scale))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m Quote) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m Quote) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m Quote) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m Quote) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m Quote) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m Quote) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m Quote) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m Quote) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m Quote) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m Quote) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m Quote) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m Quote) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m Quote) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m Quote) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m Quote) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m Quote) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetQuoteResponseLevel sets QuoteResponseLevel, Tag 301
func (m Quote) SetQuoteResponseLevel(v enum.QuoteResponseLevel) {
	m.Set(field.NewQuoteResponseLevel(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m Quote) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m Quote) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m Quote) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m Quote) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m Quote) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m Quote) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m Quote) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m Quote) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m Quote) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m Quote) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m Quote) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m Quote) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m Quote) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m Quote) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetQuoteType sets QuoteType, Tag 537
func (m Quote) SetQuoteType(v enum.QuoteType) {
	m.Set(field.NewQuoteType(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m Quote) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m Quote) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetAccountType sets AccountType, Tag 581
func (m Quote) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m Quote) SetCustOrderCapacity(v enum.CustOrderCapacity) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m Quote) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetMidPx sets MidPx, Tag 631
func (m Quote) SetMidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMidPx(value, scale))
}

//SetBidYield sets BidYield, Tag 632
func (m Quote) SetBidYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidYield(value, scale))
}

//SetMidYield sets MidYield, Tag 633
func (m Quote) SetMidYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewMidYield(value, scale))
}

//SetOfferYield sets OfferYield, Tag 634
func (m Quote) SetOfferYield(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferYield(value, scale))
}

//SetBidForwardPoints2 sets BidForwardPoints2, Tag 642
func (m Quote) SetBidForwardPoints2(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidForwardPoints2(value, scale))
}

//SetOfferForwardPoints2 sets OfferForwardPoints2, Tag 643
func (m Quote) SetOfferForwardPoints2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferForwardPoints2(value, scale))
}

//SetMktBidPx sets MktBidPx, Tag 645
func (m Quote) SetMktBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMktBidPx(value, scale))
}

//SetMktOfferPx sets MktOfferPx, Tag 646
func (m Quote) SetMktOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMktOfferPx(value, scale))
}

//SetMinBidSize sets MinBidSize, Tag 647
func (m Quote) SetMinBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinBidSize(value, scale))
}

//SetMinOfferSize sets MinOfferSize, Tag 648
func (m Quote) SetMinOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinOfferSize(value, scale))
}

//SetSettlCurrBidFxRate sets SettlCurrBidFxRate, Tag 656
func (m Quote) SetSettlCurrBidFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrBidFxRate(value, scale))
}

//SetSettlCurrOfferFxRate sets SettlCurrOfferFxRate, Tag 657
func (m Quote) SetSettlCurrOfferFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrOfferFxRate(value, scale))
}

//GetAccount gets Account, Tag 1
func (m Quote) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m Quote) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m Quote) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m Quote) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m Quote) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m Quote) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Quote) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Quote) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m Quote) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Quote) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m Quote) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m Quote) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Quote) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Quote) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m Quote) GetExDestination() (v enum.ExDestination, err quickfix.MessageRejectError) {
	var f field.ExDestinationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m Quote) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Quote) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m Quote) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m Quote) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidPx gets BidPx, Tag 132
func (m Quote) GetBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferPx gets OfferPx, Tag 133
func (m Quote) GetOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSize gets BidSize, Tag 134
func (m Quote) GetBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSize gets OfferSize, Tag 135
func (m Quote) GetOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m Quote) GetSettlCurrFxRateCalc() (v enum.SettlCurrFxRateCalc, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m Quote) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSpotRate gets BidSpotRate, Tag 188
func (m Quote) GetBidSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidForwardPoints gets BidForwardPoints, Tag 189
func (m Quote) GetBidForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSpotRate gets OfferSpotRate, Tag 190
func (m Quote) GetOfferSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferForwardPoints gets OfferForwardPoints, Tag 191
func (m Quote) GetOfferForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m Quote) GetOrderQty2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m Quote) GetFutSettDate2() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m Quote) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m Quote) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m Quote) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m Quote) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m Quote) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m Quote) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m Quote) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m Quote) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m Quote) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m Quote) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m Quote) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m Quote) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m Quote) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m Quote) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteResponseLevel gets QuoteResponseLevel, Tag 301
func (m Quote) GetQuoteResponseLevel() (v enum.QuoteResponseLevel, err quickfix.MessageRejectError) {
	var f field.QuoteResponseLevelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m Quote) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m Quote) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m Quote) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m Quote) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m Quote) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m Quote) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m Quote) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m Quote) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m Quote) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m Quote) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m Quote) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m Quote) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m Quote) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m Quote) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteType gets QuoteType, Tag 537
func (m Quote) GetQuoteType() (v enum.QuoteType, err quickfix.MessageRejectError) {
	var f field.QuoteTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m Quote) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m Quote) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccountType gets AccountType, Tag 581
func (m Quote) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m Quote) GetCustOrderCapacity() (v enum.CustOrderCapacity, err quickfix.MessageRejectError) {
	var f field.CustOrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m Quote) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMidPx gets MidPx, Tag 631
func (m Quote) GetMidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidYield gets BidYield, Tag 632
func (m Quote) GetBidYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidYieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMidYield gets MidYield, Tag 633
func (m Quote) GetMidYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MidYieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferYield gets OfferYield, Tag 634
func (m Quote) GetOfferYield() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferYieldField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidForwardPoints2 gets BidForwardPoints2, Tag 642
func (m Quote) GetBidForwardPoints2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidForwardPoints2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferForwardPoints2 gets OfferForwardPoints2, Tag 643
func (m Quote) GetOfferForwardPoints2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferForwardPoints2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMktBidPx gets MktBidPx, Tag 645
func (m Quote) GetMktBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MktBidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMktOfferPx gets MktOfferPx, Tag 646
func (m Quote) GetMktOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MktOfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinBidSize gets MinBidSize, Tag 647
func (m Quote) GetMinBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinBidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinOfferSize gets MinOfferSize, Tag 648
func (m Quote) GetMinOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinOfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrBidFxRate gets SettlCurrBidFxRate, Tag 656
func (m Quote) GetSettlCurrBidFxRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrBidFxRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrOfferFxRate gets SettlCurrOfferFxRate, Tag 657
func (m Quote) GetSettlCurrOfferFxRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrOfferFxRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m Quote) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasCommission returns true if Commission is present, Tag 12
func (m Quote) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m Quote) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m Quote) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m Quote) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m Quote) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Quote) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Quote) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m Quote) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Quote) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m Quote) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m Quote) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m Quote) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Quote) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m Quote) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Quote) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Quote) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m Quote) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m Quote) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasBidPx returns true if BidPx is present, Tag 132
func (m Quote) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

//HasOfferPx returns true if OfferPx is present, Tag 133
func (m Quote) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

//HasBidSize returns true if BidSize is present, Tag 134
func (m Quote) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

//HasOfferSize returns true if OfferSize is present, Tag 135
func (m Quote) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m Quote) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m Quote) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasBidSpotRate returns true if BidSpotRate is present, Tag 188
func (m Quote) HasBidSpotRate() bool {
	return m.Has(tag.BidSpotRate)
}

//HasBidForwardPoints returns true if BidForwardPoints is present, Tag 189
func (m Quote) HasBidForwardPoints() bool {
	return m.Has(tag.BidForwardPoints)
}

//HasOfferSpotRate returns true if OfferSpotRate is present, Tag 190
func (m Quote) HasOfferSpotRate() bool {
	return m.Has(tag.OfferSpotRate)
}

//HasOfferForwardPoints returns true if OfferForwardPoints is present, Tag 191
func (m Quote) HasOfferForwardPoints() bool {
	return m.Has(tag.OfferForwardPoints)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m Quote) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m Quote) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m Quote) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m Quote) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m Quote) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m Quote) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m Quote) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m Quote) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m Quote) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m Quote) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m Quote) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m Quote) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m Quote) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m Quote) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m Quote) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m Quote) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasQuoteResponseLevel returns true if QuoteResponseLevel is present, Tag 301
func (m Quote) HasQuoteResponseLevel() bool {
	return m.Has(tag.QuoteResponseLevel)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m Quote) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m Quote) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m Quote) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m Quote) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m Quote) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m Quote) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m Quote) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m Quote) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m Quote) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m Quote) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m Quote) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m Quote) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m Quote) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m Quote) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasQuoteType returns true if QuoteType is present, Tag 537
func (m Quote) HasQuoteType() bool {
	return m.Has(tag.QuoteType)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m Quote) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m Quote) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m Quote) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m Quote) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m Quote) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasMidPx returns true if MidPx is present, Tag 631
func (m Quote) HasMidPx() bool {
	return m.Has(tag.MidPx)
}

//HasBidYield returns true if BidYield is present, Tag 632
func (m Quote) HasBidYield() bool {
	return m.Has(tag.BidYield)
}

//HasMidYield returns true if MidYield is present, Tag 633
func (m Quote) HasMidYield() bool {
	return m.Has(tag.MidYield)
}

//HasOfferYield returns true if OfferYield is present, Tag 634
func (m Quote) HasOfferYield() bool {
	return m.Has(tag.OfferYield)
}

//HasBidForwardPoints2 returns true if BidForwardPoints2 is present, Tag 642
func (m Quote) HasBidForwardPoints2() bool {
	return m.Has(tag.BidForwardPoints2)
}

//HasOfferForwardPoints2 returns true if OfferForwardPoints2 is present, Tag 643
func (m Quote) HasOfferForwardPoints2() bool {
	return m.Has(tag.OfferForwardPoints2)
}

//HasMktBidPx returns true if MktBidPx is present, Tag 645
func (m Quote) HasMktBidPx() bool {
	return m.Has(tag.MktBidPx)
}

//HasMktOfferPx returns true if MktOfferPx is present, Tag 646
func (m Quote) HasMktOfferPx() bool {
	return m.Has(tag.MktOfferPx)
}

//HasMinBidSize returns true if MinBidSize is present, Tag 647
func (m Quote) HasMinBidSize() bool {
	return m.Has(tag.MinBidSize)
}

//HasMinOfferSize returns true if MinOfferSize is present, Tag 648
func (m Quote) HasMinOfferSize() bool {
	return m.Has(tag.MinOfferSize)
}

//HasSettlCurrBidFxRate returns true if SettlCurrBidFxRate is present, Tag 656
func (m Quote) HasSettlCurrBidFxRate() bool {
	return m.Has(tag.SettlCurrBidFxRate)
}

//HasSettlCurrOfferFxRate returns true if SettlCurrOfferFxRate is present, Tag 657
func (m Quote) HasSettlCurrOfferFxRate() bool {
	return m.Has(tag.SettlCurrOfferFxRate)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
	*quickfix.Group
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
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
