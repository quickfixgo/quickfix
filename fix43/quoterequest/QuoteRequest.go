package quoterequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//QuoteRequest is the fix43 QuoteRequest type, MsgType = R
type QuoteRequest struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a QuoteRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) QuoteRequest {
	return QuoteRequest{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m QuoteRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a QuoteRequest initialized with the required fields for QuoteRequest
func New(quotereqid field.QuoteReqIDField) (m QuoteRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("R"))
	m.Set(quotereqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg QuoteRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "R", r
}

//SetText sets Text, Tag 58
func (m QuoteRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m QuoteRequest) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m QuoteRequest) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m QuoteRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m QuoteRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetRFQReqID sets RFQReqID, Tag 644
func (m QuoteRequest) SetRFQReqID(v string) {
	m.Set(field.NewRFQReqID(v))
}

//GetText gets Text, Tag 58
func (m QuoteRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m QuoteRequest) GetQuoteReqID() (f field.QuoteReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m QuoteRequest) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m QuoteRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m QuoteRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRFQReqID gets RFQReqID, Tag 644
func (m QuoteRequest) GetRFQReqID() (f field.RFQReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m QuoteRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m QuoteRequest) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m QuoteRequest) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m QuoteRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m QuoteRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasRFQReqID returns true if RFQReqID is present, Tag 644
func (m QuoteRequest) HasRFQReqID() bool {
	return m.Has(tag.RFQReqID)
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

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NoRelatedSym) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NoRelatedSym) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NoRelatedSym) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoRelatedSym) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoRelatedSym) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NoRelatedSym) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NoRelatedSym) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NoRelatedSym) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NoRelatedSym) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NoRelatedSym) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NoRelatedSym) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m NoRelatedSym) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NoRelatedSym) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NoRelatedSym) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NoRelatedSym) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NoRelatedSym) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NoRelatedSym) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NoRelatedSym) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoRelatedSym) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoRelatedSym) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoRelatedSym) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
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

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NoRelatedSym) SetPrevClosePx(v float64) {
	m.Set(field.NewPrevClosePx(v))
}

//SetQuoteRequestType sets QuoteRequestType, Tag 303
func (m NoRelatedSym) SetQuoteRequestType(v int) {
	m.Set(field.NewQuoteRequestType(v))
}

//SetQuoteType sets QuoteType, Tag 537
func (m NoRelatedSym) SetQuoteType(v int) {
	m.Set(field.NewQuoteType(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoRelatedSym) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoRelatedSym) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m NoRelatedSym) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetNoStipulations sets NoStipulations, Tag 232
func (m NoRelatedSym) SetNoStipulations(f NoStipulationsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSide sets Side, Tag 54
func (m NoRelatedSym) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetQuantityType sets QuantityType, Tag 465
func (m NoRelatedSym) SetQuantityType(v int) {
	m.Set(field.NewQuantityType(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m NoRelatedSym) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m NoRelatedSym) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m NoRelatedSym) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NoRelatedSym) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetOrdType sets OrdType, Tag 40
func (m NoRelatedSym) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m NoRelatedSym) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m NoRelatedSym) SetOrderQty2(v float64) {
	m.Set(field.NewOrderQty2(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NoRelatedSym) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m NoRelatedSym) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoRelatedSym) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSpread sets Spread, Tag 218
func (m NoRelatedSym) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m NoRelatedSym) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m NoRelatedSym) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m NoRelatedSym) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetPriceType sets PriceType, Tag 423
func (m NoRelatedSym) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetPrice sets Price, Tag 44
func (m NoRelatedSym) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetPrice2 sets Price2, Tag 640
func (m NoRelatedSym) SetPrice2(v float64) {
	m.Set(field.NewPrice2(v))
}

//SetYieldType sets YieldType, Tag 235
func (m NoRelatedSym) SetYieldType(v string) {
	m.Set(field.NewYieldType(v))
}

//SetYield sets Yield, Tag 236
func (m NoRelatedSym) SetYield(v float64) {
	m.Set(field.NewYield(v))
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

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NoRelatedSym) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NoRelatedSym) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NoRelatedSym) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoRelatedSym) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
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

//GetMaturityDate gets MaturityDate, Tag 541
func (m NoRelatedSym) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NoRelatedSym) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NoRelatedSym) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NoRelatedSym) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NoRelatedSym) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NoRelatedSym) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m NoRelatedSym) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NoRelatedSym) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NoRelatedSym) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NoRelatedSym) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NoRelatedSym) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NoRelatedSym) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NoRelatedSym) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
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

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NoRelatedSym) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteRequestType gets QuoteRequestType, Tag 303
func (m NoRelatedSym) GetQuoteRequestType() (f field.QuoteRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteType gets QuoteType, Tag 537
func (m NoRelatedSym) GetQuoteType() (f field.QuoteTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoRelatedSym) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoRelatedSym) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m NoRelatedSym) GetTradeOriginationDate() (f field.TradeOriginationDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStipulations gets NoStipulations, Tag 232
func (m NoRelatedSym) GetNoStipulations() (NoStipulationsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStipulationsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSide gets Side, Tag 54
func (m NoRelatedSym) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuantityType gets QuantityType, Tag 465
func (m NoRelatedSym) GetQuantityType() (f field.QuantityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m NoRelatedSym) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m NoRelatedSym) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m NoRelatedSym) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NoRelatedSym) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NoRelatedSym) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m NoRelatedSym) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m NoRelatedSym) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoRelatedSym) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m NoRelatedSym) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoRelatedSym) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSpread gets Spread, Tag 218
func (m NoRelatedSym) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m NoRelatedSym) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m NoRelatedSym) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m NoRelatedSym) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m NoRelatedSym) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NoRelatedSym) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice2 gets Price2, Tag 640
func (m NoRelatedSym) GetPrice2() (f field.Price2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYieldType gets YieldType, Tag 235
func (m NoRelatedSym) GetYieldType() (f field.YieldTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetYield gets Yield, Tag 236
func (m NoRelatedSym) GetYield() (f field.YieldField, err quickfix.MessageRejectError) {
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

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NoRelatedSym) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NoRelatedSym) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoRelatedSym) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoRelatedSym) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NoRelatedSym) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NoRelatedSym) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NoRelatedSym) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NoRelatedSym) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NoRelatedSym) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NoRelatedSym) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NoRelatedSym) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NoRelatedSym) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NoRelatedSym) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NoRelatedSym) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NoRelatedSym) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NoRelatedSym) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NoRelatedSym) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
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

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NoRelatedSym) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasQuoteRequestType returns true if QuoteRequestType is present, Tag 303
func (m NoRelatedSym) HasQuoteRequestType() bool {
	return m.Has(tag.QuoteRequestType)
}

//HasQuoteType returns true if QuoteType is present, Tag 537
func (m NoRelatedSym) HasQuoteType() bool {
	return m.Has(tag.QuoteType)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoRelatedSym) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoRelatedSym) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m NoRelatedSym) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasNoStipulations returns true if NoStipulations is present, Tag 232
func (m NoRelatedSym) HasNoStipulations() bool {
	return m.Has(tag.NoStipulations)
}

//HasSide returns true if Side is present, Tag 54
func (m NoRelatedSym) HasSide() bool {
	return m.Has(tag.Side)
}

//HasQuantityType returns true if QuantityType is present, Tag 465
func (m NoRelatedSym) HasQuantityType() bool {
	return m.Has(tag.QuantityType)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m NoRelatedSym) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m NoRelatedSym) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m NoRelatedSym) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NoRelatedSym) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoRelatedSym) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m NoRelatedSym) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m NoRelatedSym) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NoRelatedSym) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m NoRelatedSym) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoRelatedSym) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSpread returns true if Spread is present, Tag 218
func (m NoRelatedSym) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m NoRelatedSym) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m NoRelatedSym) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m NoRelatedSym) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m NoRelatedSym) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasPrice returns true if Price is present, Tag 44
func (m NoRelatedSym) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasPrice2 returns true if Price2 is present, Tag 640
func (m NoRelatedSym) HasPrice2() bool {
	return m.Has(tag.Price2)
}

//HasYieldType returns true if YieldType is present, Tag 235
func (m NoRelatedSym) HasYieldType() bool {
	return m.Has(tag.YieldType)
}

//HasYield returns true if Yield is present, Tag 236
func (m NoRelatedSym) HasYield() bool {
	return m.Has(tag.Yield)
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

//NoStipulations is a repeating group element, Tag 232
type NoStipulations struct {
	quickfix.Group
}

//SetStipulationType sets StipulationType, Tag 233
func (m NoStipulations) SetStipulationType(v string) {
	m.Set(field.NewStipulationType(v))
}

//SetStipulationValue sets StipulationValue, Tag 234
func (m NoStipulations) SetStipulationValue(v string) {
	m.Set(field.NewStipulationValue(v))
}

//GetStipulationType gets StipulationType, Tag 233
func (m NoStipulations) GetStipulationType() (f field.StipulationTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStipulationValue gets StipulationValue, Tag 234
func (m NoStipulations) GetStipulationValue() (f field.StipulationValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasStipulationType returns true if StipulationType is present, Tag 233
func (m NoStipulations) HasStipulationType() bool {
	return m.Has(tag.StipulationType)
}

//HasStipulationValue returns true if StipulationValue is present, Tag 234
func (m NoStipulations) HasStipulationValue() bool {
	return m.Has(tag.StipulationValue)
}

//NoStipulationsRepeatingGroup is a repeating group, Tag 232
type NoStipulationsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStipulationsRepeatingGroup returns an initialized, NoStipulationsRepeatingGroup
func NewNoStipulationsRepeatingGroup() NoStipulationsRepeatingGroup {
	return NoStipulationsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStipulations,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StipulationType), quickfix.GroupElement(tag.StipulationValue)})}
}

//Add create and append a new NoStipulations to this group
func (m NoStipulationsRepeatingGroup) Add() NoStipulations {
	g := m.RepeatingGroup.Add()
	return NoStipulations{g}
}

//Get returns the ith NoStipulations in the NoStipulationsRepeatinGroup
func (m NoStipulationsRepeatingGroup) Get(i int) NoStipulations {
	return NoStipulations{m.RepeatingGroup.Get(i)}
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), quickfix.GroupElement(tag.NoSecurityAltID), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.PrevClosePx), quickfix.GroupElement(tag.QuoteRequestType), quickfix.GroupElement(tag.QuoteType), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.TradeOriginationDate), quickfix.GroupElement(tag.NoStipulations), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.QuantityType), quickfix.GroupElement(tag.OrderQty), quickfix.GroupElement(tag.CashOrderQty), quickfix.GroupElement(tag.SettlmntTyp), quickfix.GroupElement(tag.FutSettDate), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.FutSettDate2), quickfix.GroupElement(tag.OrderQty2), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.TransactTime), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.Spread), quickfix.GroupElement(tag.BenchmarkCurveCurrency), quickfix.GroupElement(tag.BenchmarkCurveName), quickfix.GroupElement(tag.BenchmarkCurvePoint), quickfix.GroupElement(tag.PriceType), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.Price2), quickfix.GroupElement(tag.YieldType), quickfix.GroupElement(tag.Yield)})}
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
