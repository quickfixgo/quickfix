package tradecapturereport

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//TradeCaptureReport is the fix43 TradeCaptureReport type, MsgType = AE
type TradeCaptureReport struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a TradeCaptureReport from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TradeCaptureReport {
	return TradeCaptureReport{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradeCaptureReport) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a TradeCaptureReport initialized with the required fields for TradeCaptureReport
func New(tradereportid field.TradeReportIDField, exectype field.ExecTypeField, previouslyreported field.PreviouslyReportedField, lastqty field.LastQtyField, lastpx field.LastPxField, tradedate field.TradeDateField, transacttime field.TransactTimeField) (m TradeCaptureReport) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AE"))
	m.Set(tradereportid)
	m.Set(exectype)
	m.Set(previouslyreported)
	m.Set(lastqty)
	m.Set(lastpx)
	m.Set(tradedate)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradeCaptureReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "AE", r
}

//SetExecID sets ExecID, Tag 17
func (m TradeCaptureReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradeCaptureReport) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m TradeCaptureReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m TradeCaptureReport) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastQty sets LastQty, Tag 32
func (m TradeCaptureReport) SetLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastQty(value, scale))
}

//SetOrderQty sets OrderQty, Tag 38
func (m TradeCaptureReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m TradeCaptureReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m TradeCaptureReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m TradeCaptureReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m TradeCaptureReport) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m TradeCaptureReport) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradeCaptureReport) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m TradeCaptureReport) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetIssuer sets Issuer, Tag 106
func (m TradeCaptureReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m TradeCaptureReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetExecType sets ExecType, Tag 150
func (m TradeCaptureReport) SetExecType(v enum.ExecType) {
	m.Set(field.NewExecType(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m TradeCaptureReport) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m TradeCaptureReport) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m TradeCaptureReport) SetLastSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastSpotRate(value, scale))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m TradeCaptureReport) SetLastForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastForwardPoints(value, scale))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m TradeCaptureReport) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m TradeCaptureReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m TradeCaptureReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m TradeCaptureReport) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m TradeCaptureReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m TradeCaptureReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m TradeCaptureReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m TradeCaptureReport) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m TradeCaptureReport) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradeCaptureReport) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m TradeCaptureReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m TradeCaptureReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m TradeCaptureReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m TradeCaptureReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetExecRestatementReason sets ExecRestatementReason, Tag 378
func (m TradeCaptureReport) SetExecRestatementReason(v enum.ExecRestatementReason) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m TradeCaptureReport) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m TradeCaptureReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m TradeCaptureReport) SetRoundingDirection(v enum.RoundingDirection) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m TradeCaptureReport) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m TradeCaptureReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m TradeCaptureReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetTradeReportTransType sets TradeReportTransType, Tag 487
func (m TradeCaptureReport) SetTradeReportTransType(v enum.TradeReportTransType) {
	m.Set(field.NewTradeReportTransType(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m TradeCaptureReport) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetSecondaryExecID sets SecondaryExecID, Tag 527
func (m TradeCaptureReport) SetSecondaryExecID(v string) {
	m.Set(field.NewSecondaryExecID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m TradeCaptureReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m TradeCaptureReport) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoSides sets NoSides, Tag 552
func (m TradeCaptureReport) SetNoSides(f NoSidesRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradeRequestID sets TradeRequestID, Tag 568
func (m TradeCaptureReport) SetTradeRequestID(v string) {
	m.Set(field.NewTradeRequestID(v))
}

//SetPreviouslyReported sets PreviouslyReported, Tag 570
func (m TradeCaptureReport) SetPreviouslyReported(v bool) {
	m.Set(field.NewPreviouslyReported(v))
}

//SetTradeReportID sets TradeReportID, Tag 571
func (m TradeCaptureReport) SetTradeReportID(v string) {
	m.Set(field.NewTradeReportID(v))
}

//SetTradeReportRefID sets TradeReportRefID, Tag 572
func (m TradeCaptureReport) SetTradeReportRefID(v string) {
	m.Set(field.NewTradeReportRefID(v))
}

//SetMatchStatus sets MatchStatus, Tag 573
func (m TradeCaptureReport) SetMatchStatus(v enum.MatchStatus) {
	m.Set(field.NewMatchStatus(v))
}

//SetMatchType sets MatchType, Tag 574
func (m TradeCaptureReport) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

//GetExecID gets ExecID, Tag 17
func (m TradeCaptureReport) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradeCaptureReport) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m TradeCaptureReport) GetLastMkt() (v string, err quickfix.MessageRejectError) {
	var f field.LastMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m TradeCaptureReport) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastQty gets LastQty, Tag 32
func (m TradeCaptureReport) GetLastQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m TradeCaptureReport) GetOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m TradeCaptureReport) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m TradeCaptureReport) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m TradeCaptureReport) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m TradeCaptureReport) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m TradeCaptureReport) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradeCaptureReport) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m TradeCaptureReport) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m TradeCaptureReport) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradeCaptureReport) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecType gets ExecType, Tag 150
func (m TradeCaptureReport) GetExecType() (v enum.ExecType, err quickfix.MessageRejectError) {
	var f field.ExecTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m TradeCaptureReport) GetCashOrderQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CashOrderQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m TradeCaptureReport) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m TradeCaptureReport) GetLastSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m TradeCaptureReport) GetLastForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m TradeCaptureReport) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m TradeCaptureReport) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradeCaptureReport) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m TradeCaptureReport) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradeCaptureReport) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m TradeCaptureReport) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradeCaptureReport) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradeCaptureReport) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m TradeCaptureReport) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradeCaptureReport) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReport) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradeCaptureReport) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m TradeCaptureReport) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradeCaptureReport) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradeCaptureReport) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReport) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReport) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m TradeCaptureReport) GetExecRestatementReason() (v enum.ExecRestatementReason, err quickfix.MessageRejectError) {
	var f field.ExecRestatementReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m TradeCaptureReport) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m TradeCaptureReport) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m TradeCaptureReport) GetRoundingDirection() (v enum.RoundingDirection, err quickfix.MessageRejectError) {
	var f field.RoundingDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m TradeCaptureReport) GetRoundingModulus() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundingModulusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradeCaptureReport) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReport) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradeCaptureReport) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeReportTransType gets TradeReportTransType, Tag 487
func (m TradeCaptureReport) GetTradeReportTransType() (v enum.TradeReportTransType, err quickfix.MessageRejectError) {
	var f field.TradeReportTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m TradeCaptureReport) GetOrderPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m TradeCaptureReport) GetSecondaryExecID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m TradeCaptureReport) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradeCaptureReport) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSides gets NoSides, Tag 552
func (m TradeCaptureReport) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeRequestID gets TradeRequestID, Tag 568
func (m TradeCaptureReport) GetTradeRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPreviouslyReported gets PreviouslyReported, Tag 570
func (m TradeCaptureReport) GetPreviouslyReported() (v bool, err quickfix.MessageRejectError) {
	var f field.PreviouslyReportedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeReportID gets TradeReportID, Tag 571
func (m TradeCaptureReport) GetTradeReportID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeReportRefID gets TradeReportRefID, Tag 572
func (m TradeCaptureReport) GetTradeReportRefID() (v string, err quickfix.MessageRejectError) {
	var f field.TradeReportRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m TradeCaptureReport) GetMatchStatus() (v enum.MatchStatus, err quickfix.MessageRejectError) {
	var f field.MatchStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchType gets MatchType, Tag 574
func (m TradeCaptureReport) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasExecID returns true if ExecID is present, Tag 17
func (m TradeCaptureReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m TradeCaptureReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m TradeCaptureReport) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m TradeCaptureReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastQty returns true if LastQty is present, Tag 32
func (m TradeCaptureReport) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m TradeCaptureReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m TradeCaptureReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m TradeCaptureReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m TradeCaptureReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m TradeCaptureReport) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m TradeCaptureReport) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m TradeCaptureReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m TradeCaptureReport) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m TradeCaptureReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m TradeCaptureReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasExecType returns true if ExecType is present, Tag 150
func (m TradeCaptureReport) HasExecType() bool {
	return m.Has(tag.ExecType)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m TradeCaptureReport) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m TradeCaptureReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasLastSpotRate returns true if LastSpotRate is present, Tag 194
func (m TradeCaptureReport) HasLastSpotRate() bool {
	return m.Has(tag.LastSpotRate)
}

//HasLastForwardPoints returns true if LastForwardPoints is present, Tag 195
func (m TradeCaptureReport) HasLastForwardPoints() bool {
	return m.Has(tag.LastForwardPoints)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m TradeCaptureReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m TradeCaptureReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m TradeCaptureReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m TradeCaptureReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m TradeCaptureReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m TradeCaptureReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m TradeCaptureReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m TradeCaptureReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m TradeCaptureReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m TradeCaptureReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m TradeCaptureReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m TradeCaptureReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m TradeCaptureReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m TradeCaptureReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m TradeCaptureReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m TradeCaptureReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m TradeCaptureReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m TradeCaptureReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasExecRestatementReason returns true if ExecRestatementReason is present, Tag 378
func (m TradeCaptureReport) HasExecRestatementReason() bool {
	return m.Has(tag.ExecRestatementReason)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m TradeCaptureReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m TradeCaptureReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m TradeCaptureReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m TradeCaptureReport) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m TradeCaptureReport) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m TradeCaptureReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m TradeCaptureReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m TradeCaptureReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasTradeReportTransType returns true if TradeReportTransType is present, Tag 487
func (m TradeCaptureReport) HasTradeReportTransType() bool {
	return m.Has(tag.TradeReportTransType)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m TradeCaptureReport) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryExecID returns true if SecondaryExecID is present, Tag 527
func (m TradeCaptureReport) HasSecondaryExecID() bool {
	return m.Has(tag.SecondaryExecID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m TradeCaptureReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m TradeCaptureReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoSides returns true if NoSides is present, Tag 552
func (m TradeCaptureReport) HasNoSides() bool {
	return m.Has(tag.NoSides)
}

//HasTradeRequestID returns true if TradeRequestID is present, Tag 568
func (m TradeCaptureReport) HasTradeRequestID() bool {
	return m.Has(tag.TradeRequestID)
}

//HasPreviouslyReported returns true if PreviouslyReported is present, Tag 570
func (m TradeCaptureReport) HasPreviouslyReported() bool {
	return m.Has(tag.PreviouslyReported)
}

//HasTradeReportID returns true if TradeReportID is present, Tag 571
func (m TradeCaptureReport) HasTradeReportID() bool {
	return m.Has(tag.TradeReportID)
}

//HasTradeReportRefID returns true if TradeReportRefID is present, Tag 572
func (m TradeCaptureReport) HasTradeReportRefID() bool {
	return m.Has(tag.TradeReportRefID)
}

//HasMatchStatus returns true if MatchStatus is present, Tag 573
func (m TradeCaptureReport) HasMatchStatus() bool {
	return m.Has(tag.MatchStatus)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m TradeCaptureReport) HasMatchType() bool {
	return m.Has(tag.MatchType)
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

//NoSides is a repeating group element, Tag 552
type NoSides struct {
	*quickfix.Group
}

//SetSide sets Side, Tag 54
func (m NoSides) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoSides) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m NoSides) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoSides) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSides) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAccount sets Account, Tag 1
func (m NoSides) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAccountType sets AccountType, Tag 581
func (m NoSides) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoSides) SetProcessCode(v enum.ProcessCode) {
	m.Set(field.NewProcessCode(v))
}

//SetOddLot sets OddLot, Tag 575
func (m NoSides) SetOddLot(v bool) {
	m.Set(field.NewOddLot(v))
}

//SetNoClearingInstructions sets NoClearingInstructions, Tag 576
func (m NoSides) SetNoClearingInstructions(f NoClearingInstructionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetClearingFeeIndicator sets ClearingFeeIndicator, Tag 635
func (m NoSides) SetClearingFeeIndicator(v enum.ClearingFeeIndicator) {
	m.Set(field.NewClearingFeeIndicator(v))
}

//SetTradeInputSource sets TradeInputSource, Tag 578
func (m NoSides) SetTradeInputSource(v string) {
	m.Set(field.NewTradeInputSource(v))
}

//SetTradeInputDevice sets TradeInputDevice, Tag 579
func (m NoSides) SetTradeInputDevice(v string) {
	m.Set(field.NewTradeInputDevice(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoSides) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m NoSides) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m NoSides) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetOrderCapacity sets OrderCapacity, Tag 528
func (m NoSides) SetOrderCapacity(v enum.OrderCapacity) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NoSides) SetOrderRestrictions(v enum.OrderRestrictions) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NoSides) SetCustOrderCapacity(v enum.CustOrderCapacity) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetTransBkdTime sets TransBkdTime, Tag 483
func (m NoSides) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoSides) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoSides) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetCommission sets Commission, Tag 12
func (m NoSides) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NoSides) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoSides) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoSides) SetFundRenewWaiv(v enum.FundRenewWaiv) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m NoSides) SetGrossTradeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewGrossTradeAmt(value, scale))
}

//SetNumDaysInterest sets NumDaysInterest, Tag 157
func (m NoSides) SetNumDaysInterest(v int) {
	m.Set(field.NewNumDaysInterest(v))
}

//SetExDate sets ExDate, Tag 230
func (m NoSides) SetExDate(v string) {
	m.Set(field.NewExDate(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m NoSides) SetAccruedInterestRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestRate(value, scale))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoSides) SetAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestAmt(value, scale))
}

//SetConcession sets Concession, Tag 238
func (m NoSides) SetConcession(value decimal.Decimal, scale int32) {
	m.Set(field.NewConcession(value, scale))
}

//SetTotalTakedown sets TotalTakedown, Tag 237
func (m NoSides) SetTotalTakedown(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalTakedown(value, scale))
}

//SetNetMoney sets NetMoney, Tag 118
func (m NoSides) SetNetMoney(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetMoney(value, scale))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoSides) SetSettlCurrAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrAmt(value, scale))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoSides) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoSides) SetSettlCurrFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrFxRate(value, scale))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m NoSides) SetSettlCurrFxRateCalc(v enum.SettlCurrFxRateCalc) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NoSides) SetPositionEffect(v enum.PositionEffect) {
	m.Set(field.NewPositionEffect(v))
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

//SetMultiLegReportingType sets MultiLegReportingType, Tag 442
func (m NoSides) SetMultiLegReportingType(v enum.MultiLegReportingType) {
	m.Set(field.NewMultiLegReportingType(v))
}

//SetNoContAmts sets NoContAmts, Tag 518
func (m NoSides) SetNoContAmts(f NoContAmtsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m NoSides) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//GetSide gets Side, Tag 54
func (m NoSides) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoSides) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoSides) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoSides) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSides) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccount gets Account, Tag 1
func (m NoSides) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NoSides) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoSides) GetProcessCode() (v enum.ProcessCode, err quickfix.MessageRejectError) {
	var f field.ProcessCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOddLot gets OddLot, Tag 575
func (m NoSides) GetOddLot() (v bool, err quickfix.MessageRejectError) {
	var f field.OddLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoClearingInstructions gets NoClearingInstructions, Tag 576
func (m NoSides) GetNoClearingInstructions() (NoClearingInstructionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoClearingInstructionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NoSides) GetClearingFeeIndicator() (v enum.ClearingFeeIndicator, err quickfix.MessageRejectError) {
	var f field.ClearingFeeIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeInputSource gets TradeInputSource, Tag 578
func (m NoSides) GetTradeInputSource() (v string, err quickfix.MessageRejectError) {
	var f field.TradeInputSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeInputDevice gets TradeInputDevice, Tag 579
func (m NoSides) GetTradeInputDevice() (v string, err quickfix.MessageRejectError) {
	var f field.TradeInputDeviceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoSides) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NoSides) GetComplianceID() (v string, err quickfix.MessageRejectError) {
	var f field.ComplianceIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NoSides) GetSolicitedFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.SolicitedFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NoSides) GetOrderCapacity() (v enum.OrderCapacity, err quickfix.MessageRejectError) {
	var f field.OrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NoSides) GetOrderRestrictions() (v enum.OrderRestrictions, err quickfix.MessageRejectError) {
	var f field.OrderRestrictionsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NoSides) GetCustOrderCapacity() (v enum.CustOrderCapacity, err quickfix.MessageRejectError) {
	var f field.CustOrderCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransBkdTime gets TransBkdTime, Tag 483
func (m NoSides) GetTransBkdTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransBkdTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoSides) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoSides) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m NoSides) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m NoSides) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoSides) GetCommCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CommCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoSides) GetFundRenewWaiv() (v enum.FundRenewWaiv, err quickfix.MessageRejectError) {
	var f field.FundRenewWaivField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m NoSides) GetGrossTradeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.GrossTradeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m NoSides) GetNumDaysInterest() (v int, err quickfix.MessageRejectError) {
	var f field.NumDaysInterestField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExDate gets ExDate, Tag 230
func (m NoSides) GetExDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m NoSides) GetAccruedInterestRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AccruedInterestRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoSides) GetAccruedInterestAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetConcession gets Concession, Tag 238
func (m NoSides) GetConcession() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ConcessionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalTakedown gets TotalTakedown, Tag 237
func (m NoSides) GetTotalTakedown() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalTakedownField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m NoSides) GetNetMoney() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NetMoneyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoSides) GetSettlCurrAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoSides) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoSides) GetSettlCurrFxRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m NoSides) GetSettlCurrFxRateCalc() (v enum.SettlCurrFxRateCalc, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NoSides) GetPositionEffect() (v enum.PositionEffect, err quickfix.MessageRejectError) {
	var f field.PositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m NoSides) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoSides) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoSides) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m NoSides) GetMultiLegReportingType() (v enum.MultiLegReportingType, err quickfix.MessageRejectError) {
	var f field.MultiLegReportingTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoContAmts gets NoContAmts, Tag 518
func (m NoSides) GetNoContAmts() (NoContAmtsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContAmtsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m NoSides) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSide returns true if Side is present, Tag 54
func (m NoSides) HasSide() bool {
	return m.Has(tag.Side)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoSides) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m NoSides) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoSides) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSides) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasAccount returns true if Account is present, Tag 1
func (m NoSides) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m NoSides) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NoSides) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasOddLot returns true if OddLot is present, Tag 575
func (m NoSides) HasOddLot() bool {
	return m.Has(tag.OddLot)
}

//HasNoClearingInstructions returns true if NoClearingInstructions is present, Tag 576
func (m NoSides) HasNoClearingInstructions() bool {
	return m.Has(tag.NoClearingInstructions)
}

//HasClearingFeeIndicator returns true if ClearingFeeIndicator is present, Tag 635
func (m NoSides) HasClearingFeeIndicator() bool {
	return m.Has(tag.ClearingFeeIndicator)
}

//HasTradeInputSource returns true if TradeInputSource is present, Tag 578
func (m NoSides) HasTradeInputSource() bool {
	return m.Has(tag.TradeInputSource)
}

//HasTradeInputDevice returns true if TradeInputDevice is present, Tag 579
func (m NoSides) HasTradeInputDevice() bool {
	return m.Has(tag.TradeInputDevice)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoSides) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m NoSides) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m NoSides) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasOrderCapacity returns true if OrderCapacity is present, Tag 528
func (m NoSides) HasOrderCapacity() bool {
	return m.Has(tag.OrderCapacity)
}

//HasOrderRestrictions returns true if OrderRestrictions is present, Tag 529
func (m NoSides) HasOrderRestrictions() bool {
	return m.Has(tag.OrderRestrictions)
}

//HasCustOrderCapacity returns true if CustOrderCapacity is present, Tag 582
func (m NoSides) HasCustOrderCapacity() bool {
	return m.Has(tag.CustOrderCapacity)
}

//HasTransBkdTime returns true if TransBkdTime is present, Tag 483
func (m NoSides) HasTransBkdTime() bool {
	return m.Has(tag.TransBkdTime)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoSides) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoSides) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoSides) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoSides) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m NoSides) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m NoSides) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m NoSides) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasNumDaysInterest returns true if NumDaysInterest is present, Tag 157
func (m NoSides) HasNumDaysInterest() bool {
	return m.Has(tag.NumDaysInterest)
}

//HasExDate returns true if ExDate is present, Tag 230
func (m NoSides) HasExDate() bool {
	return m.Has(tag.ExDate)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m NoSides) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m NoSides) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasConcession returns true if Concession is present, Tag 238
func (m NoSides) HasConcession() bool {
	return m.Has(tag.Concession)
}

//HasTotalTakedown returns true if TotalTakedown is present, Tag 237
func (m NoSides) HasTotalTakedown() bool {
	return m.Has(tag.TotalTakedown)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m NoSides) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m NoSides) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoSides) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m NoSides) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m NoSides) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m NoSides) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
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

//HasMultiLegReportingType returns true if MultiLegReportingType is present, Tag 442
func (m NoSides) HasMultiLegReportingType() bool {
	return m.Has(tag.MultiLegReportingType)
}

//HasNoContAmts returns true if NoContAmts is present, Tag 518
func (m NoSides) HasNoContAmts() bool {
	return m.Has(tag.NoContAmts)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m NoSides) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
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

//NoClearingInstructions is a repeating group element, Tag 576
type NoClearingInstructions struct {
	*quickfix.Group
}

//SetClearingInstruction sets ClearingInstruction, Tag 577
func (m NoClearingInstructions) SetClearingInstruction(v enum.ClearingInstruction) {
	m.Set(field.NewClearingInstruction(v))
}

//GetClearingInstruction gets ClearingInstruction, Tag 577
func (m NoClearingInstructions) GetClearingInstruction() (v enum.ClearingInstruction, err quickfix.MessageRejectError) {
	var f field.ClearingInstructionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasClearingInstruction returns true if ClearingInstruction is present, Tag 577
func (m NoClearingInstructions) HasClearingInstruction() bool {
	return m.Has(tag.ClearingInstruction)
}

//NoClearingInstructionsRepeatingGroup is a repeating group, Tag 576
type NoClearingInstructionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoClearingInstructionsRepeatingGroup returns an initialized, NoClearingInstructionsRepeatingGroup
func NewNoClearingInstructionsRepeatingGroup() NoClearingInstructionsRepeatingGroup {
	return NoClearingInstructionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoClearingInstructions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClearingInstruction)})}
}

//Add create and append a new NoClearingInstructions to this group
func (m NoClearingInstructionsRepeatingGroup) Add() NoClearingInstructions {
	g := m.RepeatingGroup.Add()
	return NoClearingInstructions{g}
}

//Get returns the ith NoClearingInstructions in the NoClearingInstructionsRepeatinGroup
func (m NoClearingInstructionsRepeatingGroup) Get(i int) NoClearingInstructions {
	return NoClearingInstructions{m.RepeatingGroup.Get(i)}
}

//NoContAmts is a repeating group element, Tag 518
type NoContAmts struct {
	*quickfix.Group
}

//SetContAmtType sets ContAmtType, Tag 519
func (m NoContAmts) SetContAmtType(v enum.ContAmtType) {
	m.Set(field.NewContAmtType(v))
}

//SetContAmtValue sets ContAmtValue, Tag 520
func (m NoContAmts) SetContAmtValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewContAmtValue(value, scale))
}

//SetContAmtCurr sets ContAmtCurr, Tag 521
func (m NoContAmts) SetContAmtCurr(v string) {
	m.Set(field.NewContAmtCurr(v))
}

//GetContAmtType gets ContAmtType, Tag 519
func (m NoContAmts) GetContAmtType() (v enum.ContAmtType, err quickfix.MessageRejectError) {
	var f field.ContAmtTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContAmtValue gets ContAmtValue, Tag 520
func (m NoContAmts) GetContAmtValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContAmtValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContAmtCurr gets ContAmtCurr, Tag 521
func (m NoContAmts) GetContAmtCurr() (v string, err quickfix.MessageRejectError) {
	var f field.ContAmtCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasContAmtType returns true if ContAmtType is present, Tag 519
func (m NoContAmts) HasContAmtType() bool {
	return m.Has(tag.ContAmtType)
}

//HasContAmtValue returns true if ContAmtValue is present, Tag 520
func (m NoContAmts) HasContAmtValue() bool {
	return m.Has(tag.ContAmtValue)
}

//HasContAmtCurr returns true if ContAmtCurr is present, Tag 521
func (m NoContAmts) HasContAmtCurr() bool {
	return m.Has(tag.ContAmtCurr)
}

//NoContAmtsRepeatingGroup is a repeating group, Tag 518
type NoContAmtsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContAmtsRepeatingGroup returns an initialized, NoContAmtsRepeatingGroup
func NewNoContAmtsRepeatingGroup() NoContAmtsRepeatingGroup {
	return NoContAmtsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContAmts,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContAmtType), quickfix.GroupElement(tag.ContAmtValue), quickfix.GroupElement(tag.ContAmtCurr)})}
}

//Add create and append a new NoContAmts to this group
func (m NoContAmtsRepeatingGroup) Add() NoContAmts {
	g := m.RepeatingGroup.Add()
	return NoContAmts{g}
}

//Get returns the ith NoContAmts in the NoContAmtsRepeatinGroup
func (m NoContAmtsRepeatingGroup) Get(i int) NoContAmts {
	return NoContAmts{m.RepeatingGroup.Get(i)}
}

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	*quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMiscFeeAmt(value, scale))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v enum.MiscFeeType) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MiscFeeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (v string, err quickfix.MessageRejectError) {
	var f field.MiscFeeCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (v enum.MiscFeeType, err quickfix.MessageRejectError) {
	var f field.MiscFeeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}

//NoSidesRepeatingGroup is a repeating group, Tag 552
type NoSidesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSidesRepeatingGroup returns an initialized, NoSidesRepeatingGroup
func NewNoSidesRepeatingGroup() NoSidesRepeatingGroup {
	return NoSidesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSides,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.ClOrdID), NewNoPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AccountType), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.OddLot), NewNoClearingInstructionsRepeatingGroup(), quickfix.GroupElement(tag.ClearingFeeIndicator), quickfix.GroupElement(tag.TradeInputSource), quickfix.GroupElement(tag.TradeInputDevice), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.ComplianceID), quickfix.GroupElement(tag.SolicitedFlag), quickfix.GroupElement(tag.OrderCapacity), quickfix.GroupElement(tag.OrderRestrictions), quickfix.GroupElement(tag.CustOrderCapacity), quickfix.GroupElement(tag.TransBkdTime), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.GrossTradeAmt), quickfix.GroupElement(tag.NumDaysInterest), quickfix.GroupElement(tag.ExDate), quickfix.GroupElement(tag.AccruedInterestRate), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.Concession), quickfix.GroupElement(tag.TotalTakedown), quickfix.GroupElement(tag.NetMoney), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlCurrFxRateCalc), quickfix.GroupElement(tag.PositionEffect), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.MultiLegReportingType), NewNoContAmtsRepeatingGroup(), NewNoMiscFeesRepeatingGroup()})}
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
