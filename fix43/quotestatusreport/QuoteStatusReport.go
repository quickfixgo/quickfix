package quotestatusreport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//QuoteStatusReport is the fix43 QuoteStatusReport type, MsgType = AI
type QuoteStatusReport struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
}

//FromMessage creates a QuoteStatusReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) QuoteStatusReport {
	return QuoteStatusReport{
		Header:  fix43.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix43.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m QuoteStatusReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a QuoteStatusReport initialized with the required fields for QuoteStatusReport
func New(quoteid field.QuoteIDField) (m QuoteStatusReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AI"))
	m.Set(quoteid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg QuoteStatusReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "AI", r
}

//SetAccount sets Account, Tag 1
func (m QuoteStatusReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetCommission sets Commission, Tag 12
func (m QuoteStatusReport) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m QuoteStatusReport) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCurrency sets Currency, Tag 15
func (m QuoteStatusReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m QuoteStatusReport) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrdType sets OrdType, Tag 40
func (m QuoteStatusReport) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m QuoteStatusReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m QuoteStatusReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m QuoteStatusReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m QuoteStatusReport) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m QuoteStatusReport) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m QuoteStatusReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetExDestination sets ExDestination, Tag 100
func (m QuoteStatusReport) SetExDestination(v string) {
	m.Set(field.NewExDestination(v))
}

//SetIssuer sets Issuer, Tag 106
func (m QuoteStatusReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m QuoteStatusReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m QuoteStatusReport) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m QuoteStatusReport) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetBidPx sets BidPx, Tag 132
func (m QuoteStatusReport) SetBidPx(v float64) {
	m.Set(field.NewBidPx(v))
}

//SetOfferPx sets OfferPx, Tag 133
func (m QuoteStatusReport) SetOfferPx(v float64) {
	m.Set(field.NewOfferPx(v))
}

//SetBidSize sets BidSize, Tag 134
func (m QuoteStatusReport) SetBidSize(v float64) {
	m.Set(field.NewBidSize(v))
}

//SetOfferSize sets OfferSize, Tag 135
func (m QuoteStatusReport) SetOfferSize(v float64) {
	m.Set(field.NewOfferSize(v))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m QuoteStatusReport) SetSettlCurrFxRateCalc(v string) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m QuoteStatusReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetBidSpotRate sets BidSpotRate, Tag 188
func (m QuoteStatusReport) SetBidSpotRate(v float64) {
	m.Set(field.NewBidSpotRate(v))
}

//SetBidForwardPoints sets BidForwardPoints, Tag 189
func (m QuoteStatusReport) SetBidForwardPoints(v float64) {
	m.Set(field.NewBidForwardPoints(v))
}

//SetOfferSpotRate sets OfferSpotRate, Tag 190
func (m QuoteStatusReport) SetOfferSpotRate(v float64) {
	m.Set(field.NewOfferSpotRate(v))
}

//SetOfferForwardPoints sets OfferForwardPoints, Tag 191
func (m QuoteStatusReport) SetOfferForwardPoints(v float64) {
	m.Set(field.NewOfferForwardPoints(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m QuoteStatusReport) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m QuoteStatusReport) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m QuoteStatusReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m QuoteStatusReport) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m QuoteStatusReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m QuoteStatusReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m QuoteStatusReport) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m QuoteStatusReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m QuoteStatusReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m QuoteStatusReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m QuoteStatusReport) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m QuoteStatusReport) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m QuoteStatusReport) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m QuoteStatusReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m QuoteStatusReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m QuoteStatusReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetQuoteStatus sets QuoteStatus, Tag 297
func (m QuoteStatusReport) SetQuoteStatus(v int) {
	m.Set(field.NewQuoteStatus(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m QuoteStatusReport) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m QuoteStatusReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m QuoteStatusReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m QuoteStatusReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m QuoteStatusReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m QuoteStatusReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m QuoteStatusReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m QuoteStatusReport) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m QuoteStatusReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m QuoteStatusReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m QuoteStatusReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m QuoteStatusReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetQuoteType sets QuoteType, Tag 537
func (m QuoteStatusReport) SetQuoteType(v int) {
	m.Set(field.NewQuoteType(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m QuoteStatusReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m QuoteStatusReport) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetAccountType sets AccountType, Tag 581
func (m QuoteStatusReport) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m QuoteStatusReport) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m QuoteStatusReport) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetMidPx sets MidPx, Tag 631
func (m QuoteStatusReport) SetMidPx(v float64) {
	m.Set(field.NewMidPx(v))
}

//SetBidYield sets BidYield, Tag 632
func (m QuoteStatusReport) SetBidYield(v float64) {
	m.Set(field.NewBidYield(v))
}

//SetMidYield sets MidYield, Tag 633
func (m QuoteStatusReport) SetMidYield(v float64) {
	m.Set(field.NewMidYield(v))
}

//SetOfferYield sets OfferYield, Tag 634
func (m QuoteStatusReport) SetOfferYield(v float64) {
	m.Set(field.NewOfferYield(v))
}

//SetBidForwardPoints2 sets BidForwardPoints2, Tag 642
func (m QuoteStatusReport) SetBidForwardPoints2(v float64) {
	m.Set(field.NewBidForwardPoints2(v))
}

//SetOfferForwardPoints2 sets OfferForwardPoints2, Tag 643
func (m QuoteStatusReport) SetOfferForwardPoints2(v float64) {
	m.Set(field.NewOfferForwardPoints2(v))
}

//SetMktBidPx sets MktBidPx, Tag 645
func (m QuoteStatusReport) SetMktBidPx(v float64) {
	m.Set(field.NewMktBidPx(v))
}

//SetMktOfferPx sets MktOfferPx, Tag 646
func (m QuoteStatusReport) SetMktOfferPx(v float64) {
	m.Set(field.NewMktOfferPx(v))
}

//SetMinBidSize sets MinBidSize, Tag 647
func (m QuoteStatusReport) SetMinBidSize(v float64) {
	m.Set(field.NewMinBidSize(v))
}

//SetMinOfferSize sets MinOfferSize, Tag 648
func (m QuoteStatusReport) SetMinOfferSize(v float64) {
	m.Set(field.NewMinOfferSize(v))
}

//SetQuoteStatusReqID sets QuoteStatusReqID, Tag 649
func (m QuoteStatusReport) SetQuoteStatusReqID(v string) {
	m.Set(field.NewQuoteStatusReqID(v))
}

//SetSettlCurrBidFxRate sets SettlCurrBidFxRate, Tag 656
func (m QuoteStatusReport) SetSettlCurrBidFxRate(v float64) {
	m.Set(field.NewSettlCurrBidFxRate(v))
}

//SetSettlCurrOfferFxRate sets SettlCurrOfferFxRate, Tag 657
func (m QuoteStatusReport) SetSettlCurrOfferFxRate(v float64) {
	m.Set(field.NewSettlCurrOfferFxRate(v))
}

//GetAccount gets Account, Tag 1
func (m QuoteStatusReport) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m QuoteStatusReport) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m QuoteStatusReport) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m QuoteStatusReport) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m QuoteStatusReport) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m QuoteStatusReport) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m QuoteStatusReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m QuoteStatusReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m QuoteStatusReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m QuoteStatusReport) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m QuoteStatusReport) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m QuoteStatusReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDestination gets ExDestination, Tag 100
func (m QuoteStatusReport) GetExDestination() (f field.ExDestinationField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m QuoteStatusReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m QuoteStatusReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m QuoteStatusReport) GetQuoteID() (f field.QuoteIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m QuoteStatusReport) GetQuoteReqID() (f field.QuoteReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidPx gets BidPx, Tag 132
func (m QuoteStatusReport) GetBidPx() (f field.BidPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferPx gets OfferPx, Tag 133
func (m QuoteStatusReport) GetOfferPx() (f field.OfferPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidSize gets BidSize, Tag 134
func (m QuoteStatusReport) GetBidSize() (f field.BidSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferSize gets OfferSize, Tag 135
func (m QuoteStatusReport) GetOfferSize() (f field.OfferSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m QuoteStatusReport) GetSettlCurrFxRateCalc() (f field.SettlCurrFxRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m QuoteStatusReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidSpotRate gets BidSpotRate, Tag 188
func (m QuoteStatusReport) GetBidSpotRate() (f field.BidSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidForwardPoints gets BidForwardPoints, Tag 189
func (m QuoteStatusReport) GetBidForwardPoints() (f field.BidForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferSpotRate gets OfferSpotRate, Tag 190
func (m QuoteStatusReport) GetOfferSpotRate() (f field.OfferSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferForwardPoints gets OfferForwardPoints, Tag 191
func (m QuoteStatusReport) GetOfferForwardPoints() (f field.OfferForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m QuoteStatusReport) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m QuoteStatusReport) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m QuoteStatusReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m QuoteStatusReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m QuoteStatusReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m QuoteStatusReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m QuoteStatusReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m QuoteStatusReport) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m QuoteStatusReport) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m QuoteStatusReport) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m QuoteStatusReport) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m QuoteStatusReport) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m QuoteStatusReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m QuoteStatusReport) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m QuoteStatusReport) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m QuoteStatusReport) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteStatus gets QuoteStatus, Tag 297
func (m QuoteStatusReport) GetQuoteStatus() (f field.QuoteStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m QuoteStatusReport) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m QuoteStatusReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m QuoteStatusReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m QuoteStatusReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m QuoteStatusReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m QuoteStatusReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m QuoteStatusReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m QuoteStatusReport) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m QuoteStatusReport) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m QuoteStatusReport) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m QuoteStatusReport) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m QuoteStatusReport) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteType gets QuoteType, Tag 537
func (m QuoteStatusReport) GetQuoteType() (f field.QuoteTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m QuoteStatusReport) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m QuoteStatusReport) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m QuoteStatusReport) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m QuoteStatusReport) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m QuoteStatusReport) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMidPx gets MidPx, Tag 631
func (m QuoteStatusReport) GetMidPx() (f field.MidPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidYield gets BidYield, Tag 632
func (m QuoteStatusReport) GetBidYield() (f field.BidYieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMidYield gets MidYield, Tag 633
func (m QuoteStatusReport) GetMidYield() (f field.MidYieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferYield gets OfferYield, Tag 634
func (m QuoteStatusReport) GetOfferYield() (f field.OfferYieldField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidForwardPoints2 gets BidForwardPoints2, Tag 642
func (m QuoteStatusReport) GetBidForwardPoints2() (f field.BidForwardPoints2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOfferForwardPoints2 gets OfferForwardPoints2, Tag 643
func (m QuoteStatusReport) GetOfferForwardPoints2() (f field.OfferForwardPoints2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMktBidPx gets MktBidPx, Tag 645
func (m QuoteStatusReport) GetMktBidPx() (f field.MktBidPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMktOfferPx gets MktOfferPx, Tag 646
func (m QuoteStatusReport) GetMktOfferPx() (f field.MktOfferPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinBidSize gets MinBidSize, Tag 647
func (m QuoteStatusReport) GetMinBidSize() (f field.MinBidSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinOfferSize gets MinOfferSize, Tag 648
func (m QuoteStatusReport) GetMinOfferSize() (f field.MinOfferSizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteStatusReqID gets QuoteStatusReqID, Tag 649
func (m QuoteStatusReport) GetQuoteStatusReqID() (f field.QuoteStatusReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrBidFxRate gets SettlCurrBidFxRate, Tag 656
func (m QuoteStatusReport) GetSettlCurrBidFxRate() (f field.SettlCurrBidFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrOfferFxRate gets SettlCurrOfferFxRate, Tag 657
func (m QuoteStatusReport) GetSettlCurrOfferFxRate() (f field.SettlCurrOfferFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m QuoteStatusReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasCommission returns true if Commission is present, Tag 12
func (m QuoteStatusReport) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m QuoteStatusReport) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m QuoteStatusReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m QuoteStatusReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m QuoteStatusReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m QuoteStatusReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m QuoteStatusReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m QuoteStatusReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m QuoteStatusReport) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m QuoteStatusReport) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m QuoteStatusReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasExDestination returns true if ExDestination is present, Tag 100
func (m QuoteStatusReport) HasExDestination() bool {
	return m.Has(tag.ExDestination)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m QuoteStatusReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m QuoteStatusReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m QuoteStatusReport) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m QuoteStatusReport) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasBidPx returns true if BidPx is present, Tag 132
func (m QuoteStatusReport) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

//HasOfferPx returns true if OfferPx is present, Tag 133
func (m QuoteStatusReport) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

//HasBidSize returns true if BidSize is present, Tag 134
func (m QuoteStatusReport) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

//HasOfferSize returns true if OfferSize is present, Tag 135
func (m QuoteStatusReport) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m QuoteStatusReport) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m QuoteStatusReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasBidSpotRate returns true if BidSpotRate is present, Tag 188
func (m QuoteStatusReport) HasBidSpotRate() bool {
	return m.Has(tag.BidSpotRate)
}

//HasBidForwardPoints returns true if BidForwardPoints is present, Tag 189
func (m QuoteStatusReport) HasBidForwardPoints() bool {
	return m.Has(tag.BidForwardPoints)
}

//HasOfferSpotRate returns true if OfferSpotRate is present, Tag 190
func (m QuoteStatusReport) HasOfferSpotRate() bool {
	return m.Has(tag.OfferSpotRate)
}

//HasOfferForwardPoints returns true if OfferForwardPoints is present, Tag 191
func (m QuoteStatusReport) HasOfferForwardPoints() bool {
	return m.Has(tag.OfferForwardPoints)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m QuoteStatusReport) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m QuoteStatusReport) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m QuoteStatusReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m QuoteStatusReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m QuoteStatusReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m QuoteStatusReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m QuoteStatusReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m QuoteStatusReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m QuoteStatusReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m QuoteStatusReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m QuoteStatusReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m QuoteStatusReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m QuoteStatusReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m QuoteStatusReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m QuoteStatusReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m QuoteStatusReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasQuoteStatus returns true if QuoteStatus is present, Tag 297
func (m QuoteStatusReport) HasQuoteStatus() bool {
	return m.Has(tag.QuoteStatus)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m QuoteStatusReport) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m QuoteStatusReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m QuoteStatusReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m QuoteStatusReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m QuoteStatusReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m QuoteStatusReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m QuoteStatusReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m QuoteStatusReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m QuoteStatusReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m QuoteStatusReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m QuoteStatusReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m QuoteStatusReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasQuoteType returns true if QuoteType is present, Tag 537
func (m QuoteStatusReport) HasQuoteType() bool {
	return m.Has(tag.QuoteType)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m QuoteStatusReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m QuoteStatusReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m QuoteStatusReport) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m QuoteStatusReport) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m QuoteStatusReport) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasMidPx returns true if MidPx is present, Tag 631
func (m QuoteStatusReport) HasMidPx() bool {
	return m.Has(tag.MidPx)
}

//HasBidYield returns true if BidYield is present, Tag 632
func (m QuoteStatusReport) HasBidYield() bool {
	return m.Has(tag.BidYield)
}

//HasMidYield returns true if MidYield is present, Tag 633
func (m QuoteStatusReport) HasMidYield() bool {
	return m.Has(tag.MidYield)
}

//HasOfferYield returns true if OfferYield is present, Tag 634
func (m QuoteStatusReport) HasOfferYield() bool {
	return m.Has(tag.OfferYield)
}

//HasBidForwardPoints2 returns true if BidForwardPoints2 is present, Tag 642
func (m QuoteStatusReport) HasBidForwardPoints2() bool {
	return m.Has(tag.BidForwardPoints2)
}

//HasOfferForwardPoints2 returns true if OfferForwardPoints2 is present, Tag 643
func (m QuoteStatusReport) HasOfferForwardPoints2() bool {
	return m.Has(tag.OfferForwardPoints2)
}

//HasMktBidPx returns true if MktBidPx is present, Tag 645
func (m QuoteStatusReport) HasMktBidPx() bool {
	return m.Has(tag.MktBidPx)
}

//HasMktOfferPx returns true if MktOfferPx is present, Tag 646
func (m QuoteStatusReport) HasMktOfferPx() bool {
	return m.Has(tag.MktOfferPx)
}

//HasMinBidSize returns true if MinBidSize is present, Tag 647
func (m QuoteStatusReport) HasMinBidSize() bool {
	return m.Has(tag.MinBidSize)
}

//HasMinOfferSize returns true if MinOfferSize is present, Tag 648
func (m QuoteStatusReport) HasMinOfferSize() bool {
	return m.Has(tag.MinOfferSize)
}

//HasQuoteStatusReqID returns true if QuoteStatusReqID is present, Tag 649
func (m QuoteStatusReport) HasQuoteStatusReqID() bool {
	return m.Has(tag.QuoteStatusReqID)
}

//HasSettlCurrBidFxRate returns true if SettlCurrBidFxRate is present, Tag 656
func (m QuoteStatusReport) HasSettlCurrBidFxRate() bool {
	return m.Has(tag.SettlCurrBidFxRate)
}

//HasSettlCurrOfferFxRate returns true if SettlCurrOfferFxRate is present, Tag 657
func (m QuoteStatusReport) HasSettlCurrOfferFxRate() bool {
	return m.Has(tag.SettlCurrOfferFxRate)
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
