package tradecapturereport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//TradeCaptureReport is the fix43 TradeCaptureReport type, MsgType = AE
type TradeCaptureReport struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradeCaptureReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradeCaptureReport {
	return TradeCaptureReport{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradeCaptureReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradeCaptureReport initialized with the required fields for TradeCaptureReport
func New(tradereportid field.TradeReportIDField, exectype field.ExecTypeField, previouslyreported field.PreviouslyReportedField, lastqty field.LastQtyField, lastpx field.LastPxField, tradedate field.TradeDateField, transacttime field.TransactTimeField) (m TradeCaptureReport) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AE"))
	m.Header.Set(field.NewBeginString("FIX.4.3"))
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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "AE", r
}

//SetExecID sets ExecID, Tag 17
func (m TradeCaptureReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m TradeCaptureReport) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m TradeCaptureReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m TradeCaptureReport) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastQty sets LastQty, Tag 32
func (m TradeCaptureReport) SetLastQty(v float64) {
	m.Set(field.NewLastQty(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m TradeCaptureReport) SetOrderQty(v float64) {
	m.Set(field.NewOrderQty(v))
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
func (m TradeCaptureReport) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m TradeCaptureReport) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m TradeCaptureReport) SetSymbolSfx(v string) {
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
func (m TradeCaptureReport) SetExecType(v string) {
	m.Set(field.NewExecType(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m TradeCaptureReport) SetCashOrderQty(v float64) {
	m.Set(field.NewCashOrderQty(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m TradeCaptureReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m TradeCaptureReport) SetLastSpotRate(v float64) {
	m.Set(field.NewLastSpotRate(v))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m TradeCaptureReport) SetLastForwardPoints(v float64) {
	m.Set(field.NewLastForwardPoints(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m TradeCaptureReport) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
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
func (m TradeCaptureReport) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
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
func (m TradeCaptureReport) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m TradeCaptureReport) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m TradeCaptureReport) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
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
func (m TradeCaptureReport) SetExecRestatementReason(v int) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m TradeCaptureReport) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m TradeCaptureReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m TradeCaptureReport) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m TradeCaptureReport) SetRoundingModulus(v float64) {
	m.Set(field.NewRoundingModulus(v))
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
func (m TradeCaptureReport) SetTradeReportTransType(v int) {
	m.Set(field.NewTradeReportTransType(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m TradeCaptureReport) SetOrderPercent(v float64) {
	m.Set(field.NewOrderPercent(v))
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
func (m TradeCaptureReport) SetInstrRegistry(v string) {
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
func (m TradeCaptureReport) SetMatchStatus(v string) {
	m.Set(field.NewMatchStatus(v))
}

//SetMatchType sets MatchType, Tag 574
func (m TradeCaptureReport) SetMatchType(v string) {
	m.Set(field.NewMatchType(v))
}

//GetExecID gets ExecID, Tag 17
func (m TradeCaptureReport) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m TradeCaptureReport) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m TradeCaptureReport) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m TradeCaptureReport) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastQty gets LastQty, Tag 32
func (m TradeCaptureReport) GetLastQty() (f field.LastQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m TradeCaptureReport) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m TradeCaptureReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m TradeCaptureReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m TradeCaptureReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m TradeCaptureReport) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m TradeCaptureReport) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m TradeCaptureReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m TradeCaptureReport) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m TradeCaptureReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m TradeCaptureReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecType gets ExecType, Tag 150
func (m TradeCaptureReport) GetExecType() (f field.ExecTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m TradeCaptureReport) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m TradeCaptureReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m TradeCaptureReport) GetLastSpotRate() (f field.LastSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m TradeCaptureReport) GetLastForwardPoints() (f field.LastForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m TradeCaptureReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m TradeCaptureReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m TradeCaptureReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m TradeCaptureReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m TradeCaptureReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m TradeCaptureReport) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m TradeCaptureReport) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m TradeCaptureReport) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m TradeCaptureReport) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m TradeCaptureReport) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m TradeCaptureReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m TradeCaptureReport) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m TradeCaptureReport) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m TradeCaptureReport) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m TradeCaptureReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m TradeCaptureReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m TradeCaptureReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m TradeCaptureReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m TradeCaptureReport) GetExecRestatementReason() (f field.ExecRestatementReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m TradeCaptureReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m TradeCaptureReport) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m TradeCaptureReport) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m TradeCaptureReport) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m TradeCaptureReport) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m TradeCaptureReport) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m TradeCaptureReport) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m TradeCaptureReport) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportTransType gets TradeReportTransType, Tag 487
func (m TradeCaptureReport) GetTradeReportTransType() (f field.TradeReportTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m TradeCaptureReport) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m TradeCaptureReport) GetSecondaryExecID() (f field.SecondaryExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m TradeCaptureReport) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m TradeCaptureReport) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSides gets NoSides, Tag 552
func (m TradeCaptureReport) GetNoSides() (NoSidesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSidesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradeRequestID gets TradeRequestID, Tag 568
func (m TradeCaptureReport) GetTradeRequestID() (f field.TradeRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPreviouslyReported gets PreviouslyReported, Tag 570
func (m TradeCaptureReport) GetPreviouslyReported() (f field.PreviouslyReportedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportID gets TradeReportID, Tag 571
func (m TradeCaptureReport) GetTradeReportID() (f field.TradeReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeReportRefID gets TradeReportRefID, Tag 572
func (m TradeCaptureReport) GetTradeReportRefID() (f field.TradeReportRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchStatus gets MatchStatus, Tag 573
func (m TradeCaptureReport) GetMatchStatus() (f field.MatchStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMatchType gets MatchType, Tag 574
func (m TradeCaptureReport) GetMatchType() (f field.MatchTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
func (m NoSides) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoSides) SetProcessCode(v string) {
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
func (m NoSides) SetClearingFeeIndicator(v string) {
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
func (m NoSides) SetOrderCapacity(v string) {
	m.Set(field.NewOrderCapacity(v))
}

//SetOrderRestrictions sets OrderRestrictions, Tag 529
func (m NoSides) SetOrderRestrictions(v string) {
	m.Set(field.NewOrderRestrictions(v))
}

//SetCustOrderCapacity sets CustOrderCapacity, Tag 582
func (m NoSides) SetCustOrderCapacity(v int) {
	m.Set(field.NewCustOrderCapacity(v))
}

//SetTransBkdTime sets TransBkdTime, Tag 483
func (m NoSides) SetTransBkdTime(v time.Time) {
	m.Set(field.NewTransBkdTime(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoSides) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoSides) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetCommission sets Commission, Tag 12
func (m NoSides) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m NoSides) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoSides) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoSides) SetFundRenewWaiv(v string) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m NoSides) SetGrossTradeAmt(v float64) {
	m.Set(field.NewGrossTradeAmt(v))
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
func (m NoSides) SetAccruedInterestRate(v float64) {
	m.Set(field.NewAccruedInterestRate(v))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoSides) SetAccruedInterestAmt(v float64) {
	m.Set(field.NewAccruedInterestAmt(v))
}

//SetConcession sets Concession, Tag 238
func (m NoSides) SetConcession(v float64) {
	m.Set(field.NewConcession(v))
}

//SetTotalTakedown sets TotalTakedown, Tag 237
func (m NoSides) SetTotalTakedown(v float64) {
	m.Set(field.NewTotalTakedown(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m NoSides) SetNetMoney(v float64) {
	m.Set(field.NewNetMoney(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoSides) SetSettlCurrAmt(v float64) {
	m.Set(field.NewSettlCurrAmt(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoSides) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoSides) SetSettlCurrFxRate(v float64) {
	m.Set(field.NewSettlCurrFxRate(v))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m NoSides) SetSettlCurrFxRateCalc(v string) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m NoSides) SetPositionEffect(v string) {
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
func (m NoSides) SetMultiLegReportingType(v string) {
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
func (m NoSides) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoSides) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoSides) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoSides) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSides) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccount gets Account, Tag 1
func (m NoSides) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m NoSides) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoSides) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOddLot gets OddLot, Tag 575
func (m NoSides) GetOddLot() (f field.OddLotField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoClearingInstructions gets NoClearingInstructions, Tag 576
func (m NoSides) GetNoClearingInstructions() (NoClearingInstructionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoClearingInstructionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetClearingFeeIndicator gets ClearingFeeIndicator, Tag 635
func (m NoSides) GetClearingFeeIndicator() (f field.ClearingFeeIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputSource gets TradeInputSource, Tag 578
func (m NoSides) GetTradeInputSource() (f field.TradeInputSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeInputDevice gets TradeInputDevice, Tag 579
func (m NoSides) GetTradeInputDevice() (f field.TradeInputDeviceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoSides) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m NoSides) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m NoSides) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderCapacity gets OrderCapacity, Tag 528
func (m NoSides) GetOrderCapacity() (f field.OrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderRestrictions gets OrderRestrictions, Tag 529
func (m NoSides) GetOrderRestrictions() (f field.OrderRestrictionsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCustOrderCapacity gets CustOrderCapacity, Tag 582
func (m NoSides) GetCustOrderCapacity() (f field.CustOrderCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransBkdTime gets TransBkdTime, Tag 483
func (m NoSides) GetTransBkdTime() (f field.TransBkdTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoSides) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoSides) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NoSides) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NoSides) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoSides) GetCommCurrency() (f field.CommCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoSides) GetFundRenewWaiv() (f field.FundRenewWaivField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m NoSides) GetGrossTradeAmt() (f field.GrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m NoSides) GetNumDaysInterest() (f field.NumDaysInterestField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExDate gets ExDate, Tag 230
func (m NoSides) GetExDate() (f field.ExDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m NoSides) GetAccruedInterestRate() (f field.AccruedInterestRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoSides) GetAccruedInterestAmt() (f field.AccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetConcession gets Concession, Tag 238
func (m NoSides) GetConcession() (f field.ConcessionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalTakedown gets TotalTakedown, Tag 237
func (m NoSides) GetTotalTakedown() (f field.TotalTakedownField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m NoSides) GetNetMoney() (f field.NetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoSides) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoSides) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoSides) GetSettlCurrFxRate() (f field.SettlCurrFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m NoSides) GetSettlCurrFxRateCalc() (f field.SettlCurrFxRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m NoSides) GetPositionEffect() (f field.PositionEffectField, err quickfix.MessageRejectError) {
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

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m NoSides) GetMultiLegReportingType() (f field.MultiLegReportingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//NoClearingInstructions is a repeating group element, Tag 576
type NoClearingInstructions struct {
	quickfix.Group
}

//SetClearingInstruction sets ClearingInstruction, Tag 577
func (m NoClearingInstructions) SetClearingInstruction(v int) {
	m.Set(field.NewClearingInstruction(v))
}

//GetClearingInstruction gets ClearingInstruction, Tag 577
func (m NoClearingInstructions) GetClearingInstruction() (f field.ClearingInstructionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
}

//SetContAmtType sets ContAmtType, Tag 519
func (m NoContAmts) SetContAmtType(v int) {
	m.Set(field.NewContAmtType(v))
}

//SetContAmtValue sets ContAmtValue, Tag 520
func (m NoContAmts) SetContAmtValue(v float64) {
	m.Set(field.NewContAmtValue(v))
}

//SetContAmtCurr sets ContAmtCurr, Tag 521
func (m NoContAmts) SetContAmtCurr(v string) {
	m.Set(field.NewContAmtCurr(v))
}

//GetContAmtType gets ContAmtType, Tag 519
func (m NoContAmts) GetContAmtType() (f field.ContAmtTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContAmtValue gets ContAmtValue, Tag 520
func (m NoContAmts) GetContAmtValue() (f field.ContAmtValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContAmtCurr gets ContAmtCurr, Tag 521
func (m NoContAmts) GetContAmtCurr() (f field.ContAmtCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(v float64) {
	m.Set(field.NewMiscFeeAmt(v))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v string) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (f field.MiscFeeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (f field.MiscFeeCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (f field.MiscFeeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
