package positionmaintenancereport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//PositionMaintenanceReport is the fix50sp1 PositionMaintenanceReport type, MsgType = AM
type PositionMaintenanceReport struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a PositionMaintenanceReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) PositionMaintenanceReport {
	return PositionMaintenanceReport{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m PositionMaintenanceReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a PositionMaintenanceReport initialized with the required fields for PositionMaintenanceReport
func New(posmaintrptid field.PosMaintRptIDField, postranstype field.PosTransTypeField, posmaintaction field.PosMaintActionField, posmaintstatus field.PosMaintStatusField, clearingbusinessdate field.ClearingBusinessDateField) (m PositionMaintenanceReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("AM"))
	m.Header.Set(field.NewBeginString("8"))
	m.Set(posmaintrptid)
	m.Set(postranstype)
	m.Set(posmaintaction)
	m.Set(posmaintstatus)
	m.Set(clearingbusinessdate)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg PositionMaintenanceReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "AM", r
}

//SetAccount sets Account, Tag 1
func (m PositionMaintenanceReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetCurrency sets Currency, Tag 15
func (m PositionMaintenanceReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m PositionMaintenanceReport) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m PositionMaintenanceReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m PositionMaintenanceReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m PositionMaintenanceReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m PositionMaintenanceReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m PositionMaintenanceReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m PositionMaintenanceReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m PositionMaintenanceReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m PositionMaintenanceReport) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m PositionMaintenanceReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m PositionMaintenanceReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m PositionMaintenanceReport) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m PositionMaintenanceReport) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m PositionMaintenanceReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m PositionMaintenanceReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m PositionMaintenanceReport) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m PositionMaintenanceReport) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m PositionMaintenanceReport) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m PositionMaintenanceReport) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m PositionMaintenanceReport) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m PositionMaintenanceReport) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m PositionMaintenanceReport) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m PositionMaintenanceReport) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m PositionMaintenanceReport) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m PositionMaintenanceReport) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m PositionMaintenanceReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m PositionMaintenanceReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m PositionMaintenanceReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m PositionMaintenanceReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m PositionMaintenanceReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m PositionMaintenanceReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m PositionMaintenanceReport) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m PositionMaintenanceReport) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m PositionMaintenanceReport) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m PositionMaintenanceReport) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m PositionMaintenanceReport) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m PositionMaintenanceReport) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m PositionMaintenanceReport) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m PositionMaintenanceReport) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m PositionMaintenanceReport) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m PositionMaintenanceReport) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetNoLegs sets NoLegs, Tag 555
func (m PositionMaintenanceReport) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

//SetAccountType sets AccountType, Tag 581
func (m PositionMaintenanceReport) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m PositionMaintenanceReport) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m PositionMaintenanceReport) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetPool sets Pool, Tag 691
func (m PositionMaintenanceReport) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetNoPositions sets NoPositions, Tag 702
func (m PositionMaintenanceReport) SetNoPositions(f NoPositionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetPosTransType sets PosTransType, Tag 709
func (m PositionMaintenanceReport) SetPosTransType(v int) {
	m.Set(field.NewPosTransType(v))
}

//SetPosReqID sets PosReqID, Tag 710
func (m PositionMaintenanceReport) SetPosReqID(v string) {
	m.Set(field.NewPosReqID(v))
}

//SetNoUnderlyings sets NoUnderlyings, Tag 711
func (m PositionMaintenanceReport) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

//SetPosMaintAction sets PosMaintAction, Tag 712
func (m PositionMaintenanceReport) SetPosMaintAction(v int) {
	m.Set(field.NewPosMaintAction(v))
}

//SetOrigPosReqRefID sets OrigPosReqRefID, Tag 713
func (m PositionMaintenanceReport) SetOrigPosReqRefID(v string) {
	m.Set(field.NewOrigPosReqRefID(v))
}

//SetPosMaintRptRefID sets PosMaintRptRefID, Tag 714
func (m PositionMaintenanceReport) SetPosMaintRptRefID(v string) {
	m.Set(field.NewPosMaintRptRefID(v))
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m PositionMaintenanceReport) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetSettlSessID sets SettlSessID, Tag 716
func (m PositionMaintenanceReport) SetSettlSessID(v string) {
	m.Set(field.NewSettlSessID(v))
}

//SetSettlSessSubID sets SettlSessSubID, Tag 717
func (m PositionMaintenanceReport) SetSettlSessSubID(v string) {
	m.Set(field.NewSettlSessSubID(v))
}

//SetAdjustmentType sets AdjustmentType, Tag 718
func (m PositionMaintenanceReport) SetAdjustmentType(v int) {
	m.Set(field.NewAdjustmentType(v))
}

//SetContraryInstructionIndicator sets ContraryInstructionIndicator, Tag 719
func (m PositionMaintenanceReport) SetContraryInstructionIndicator(v bool) {
	m.Set(field.NewContraryInstructionIndicator(v))
}

//SetPriorSpreadIndicator sets PriorSpreadIndicator, Tag 720
func (m PositionMaintenanceReport) SetPriorSpreadIndicator(v bool) {
	m.Set(field.NewPriorSpreadIndicator(v))
}

//SetPosMaintRptID sets PosMaintRptID, Tag 721
func (m PositionMaintenanceReport) SetPosMaintRptID(v string) {
	m.Set(field.NewPosMaintRptID(v))
}

//SetPosMaintStatus sets PosMaintStatus, Tag 722
func (m PositionMaintenanceReport) SetPosMaintStatus(v int) {
	m.Set(field.NewPosMaintStatus(v))
}

//SetPosMaintResult sets PosMaintResult, Tag 723
func (m PositionMaintenanceReport) SetPosMaintResult(v int) {
	m.Set(field.NewPosMaintResult(v))
}

//SetNoPosAmt sets NoPosAmt, Tag 753
func (m PositionMaintenanceReport) SetNoPosAmt(f NoPosAmtRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m PositionMaintenanceReport) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetThresholdAmount sets ThresholdAmount, Tag 834
func (m PositionMaintenanceReport) SetThresholdAmount(v float64) {
	m.Set(field.NewThresholdAmount(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m PositionMaintenanceReport) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m PositionMaintenanceReport) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m PositionMaintenanceReport) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m PositionMaintenanceReport) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m PositionMaintenanceReport) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m PositionMaintenanceReport) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m PositionMaintenanceReport) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m PositionMaintenanceReport) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m PositionMaintenanceReport) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m PositionMaintenanceReport) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m PositionMaintenanceReport) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m PositionMaintenanceReport) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m PositionMaintenanceReport) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m PositionMaintenanceReport) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m PositionMaintenanceReport) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m PositionMaintenanceReport) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m PositionMaintenanceReport) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m PositionMaintenanceReport) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m PositionMaintenanceReport) SetMinPriceIncrementAmount(v float64) {
	m.Set(field.NewMinPriceIncrementAmount(v))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m PositionMaintenanceReport) SetUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnitOfMeasureQty(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m PositionMaintenanceReport) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m PositionMaintenanceReport) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m PositionMaintenanceReport) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m PositionMaintenanceReport) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m PositionMaintenanceReport) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m PositionMaintenanceReport) SetPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewPriceUnitOfMeasureQty(v))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m PositionMaintenanceReport) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m PositionMaintenanceReport) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayAmount sets OptPayAmount, Tag 1195
func (m PositionMaintenanceReport) SetOptPayAmount(v float64) {
	m.Set(field.NewOptPayAmount(v))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m PositionMaintenanceReport) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetFuturesValuationMethod sets FuturesValuationMethod, Tag 1197
func (m PositionMaintenanceReport) SetFuturesValuationMethod(v string) {
	m.Set(field.NewFuturesValuationMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m PositionMaintenanceReport) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m PositionMaintenanceReport) SetCapPrice(v float64) {
	m.Set(field.NewCapPrice(v))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m PositionMaintenanceReport) SetFloorPrice(v float64) {
	m.Set(field.NewFloorPrice(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m PositionMaintenanceReport) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m PositionMaintenanceReport) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m PositionMaintenanceReport) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//GetAccount gets Account, Tag 1
func (m PositionMaintenanceReport) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m PositionMaintenanceReport) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m PositionMaintenanceReport) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m PositionMaintenanceReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m PositionMaintenanceReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m PositionMaintenanceReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m PositionMaintenanceReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m PositionMaintenanceReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m PositionMaintenanceReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m PositionMaintenanceReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m PositionMaintenanceReport) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m PositionMaintenanceReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m PositionMaintenanceReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m PositionMaintenanceReport) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m PositionMaintenanceReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m PositionMaintenanceReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m PositionMaintenanceReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m PositionMaintenanceReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m PositionMaintenanceReport) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m PositionMaintenanceReport) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m PositionMaintenanceReport) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m PositionMaintenanceReport) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m PositionMaintenanceReport) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m PositionMaintenanceReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m PositionMaintenanceReport) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m PositionMaintenanceReport) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m PositionMaintenanceReport) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m PositionMaintenanceReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m PositionMaintenanceReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m PositionMaintenanceReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m PositionMaintenanceReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m PositionMaintenanceReport) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m PositionMaintenanceReport) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m PositionMaintenanceReport) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m PositionMaintenanceReport) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m PositionMaintenanceReport) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m PositionMaintenanceReport) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m PositionMaintenanceReport) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m PositionMaintenanceReport) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m PositionMaintenanceReport) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m PositionMaintenanceReport) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m PositionMaintenanceReport) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m PositionMaintenanceReport) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegs gets NoLegs, Tag 555
func (m PositionMaintenanceReport) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAccountType gets AccountType, Tag 581
func (m PositionMaintenanceReport) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m PositionMaintenanceReport) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m PositionMaintenanceReport) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m PositionMaintenanceReport) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPositions gets NoPositions, Tag 702
func (m PositionMaintenanceReport) GetNoPositions() (NoPositionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPositionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPosTransType gets PosTransType, Tag 709
func (m PositionMaintenanceReport) GetPosTransType() (f field.PosTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosReqID gets PosReqID, Tag 710
func (m PositionMaintenanceReport) GetPosReqID() (f field.PosReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyings gets NoUnderlyings, Tag 711
func (m PositionMaintenanceReport) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetPosMaintAction gets PosMaintAction, Tag 712
func (m PositionMaintenanceReport) GetPosMaintAction() (f field.PosMaintActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigPosReqRefID gets OrigPosReqRefID, Tag 713
func (m PositionMaintenanceReport) GetOrigPosReqRefID() (f field.OrigPosReqRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosMaintRptRefID gets PosMaintRptRefID, Tag 714
func (m PositionMaintenanceReport) GetPosMaintRptRefID() (f field.PosMaintRptRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m PositionMaintenanceReport) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessID gets SettlSessID, Tag 716
func (m PositionMaintenanceReport) GetSettlSessID() (f field.SettlSessIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlSessSubID gets SettlSessSubID, Tag 717
func (m PositionMaintenanceReport) GetSettlSessSubID() (f field.SettlSessSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAdjustmentType gets AdjustmentType, Tag 718
func (m PositionMaintenanceReport) GetAdjustmentType() (f field.AdjustmentTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContraryInstructionIndicator gets ContraryInstructionIndicator, Tag 719
func (m PositionMaintenanceReport) GetContraryInstructionIndicator() (f field.ContraryInstructionIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriorSpreadIndicator gets PriorSpreadIndicator, Tag 720
func (m PositionMaintenanceReport) GetPriorSpreadIndicator() (f field.PriorSpreadIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosMaintRptID gets PosMaintRptID, Tag 721
func (m PositionMaintenanceReport) GetPosMaintRptID() (f field.PosMaintRptIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosMaintStatus gets PosMaintStatus, Tag 722
func (m PositionMaintenanceReport) GetPosMaintStatus() (f field.PosMaintStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosMaintResult gets PosMaintResult, Tag 723
func (m PositionMaintenanceReport) GetPosMaintResult() (f field.PosMaintResultField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPosAmt gets NoPosAmt, Tag 753
func (m PositionMaintenanceReport) GetNoPosAmt() (NoPosAmtRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPosAmtRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m PositionMaintenanceReport) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetThresholdAmount gets ThresholdAmount, Tag 834
func (m PositionMaintenanceReport) GetThresholdAmount() (f field.ThresholdAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m PositionMaintenanceReport) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m PositionMaintenanceReport) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m PositionMaintenanceReport) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m PositionMaintenanceReport) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m PositionMaintenanceReport) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m PositionMaintenanceReport) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m PositionMaintenanceReport) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m PositionMaintenanceReport) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m PositionMaintenanceReport) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m PositionMaintenanceReport) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m PositionMaintenanceReport) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m PositionMaintenanceReport) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m PositionMaintenanceReport) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m PositionMaintenanceReport) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m PositionMaintenanceReport) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m PositionMaintenanceReport) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m PositionMaintenanceReport) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m PositionMaintenanceReport) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m PositionMaintenanceReport) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m PositionMaintenanceReport) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m PositionMaintenanceReport) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m PositionMaintenanceReport) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m PositionMaintenanceReport) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m PositionMaintenanceReport) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m PositionMaintenanceReport) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m PositionMaintenanceReport) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m PositionMaintenanceReport) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m PositionMaintenanceReport) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayAmount gets OptPayAmount, Tag 1195
func (m PositionMaintenanceReport) GetOptPayAmount() (f field.OptPayAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m PositionMaintenanceReport) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFuturesValuationMethod gets FuturesValuationMethod, Tag 1197
func (m PositionMaintenanceReport) GetFuturesValuationMethod() (f field.FuturesValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m PositionMaintenanceReport) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m PositionMaintenanceReport) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m PositionMaintenanceReport) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m PositionMaintenanceReport) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m PositionMaintenanceReport) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m PositionMaintenanceReport) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m PositionMaintenanceReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m PositionMaintenanceReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m PositionMaintenanceReport) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m PositionMaintenanceReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m PositionMaintenanceReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m PositionMaintenanceReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m PositionMaintenanceReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m PositionMaintenanceReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m PositionMaintenanceReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m PositionMaintenanceReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m PositionMaintenanceReport) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m PositionMaintenanceReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m PositionMaintenanceReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m PositionMaintenanceReport) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m PositionMaintenanceReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m PositionMaintenanceReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m PositionMaintenanceReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m PositionMaintenanceReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m PositionMaintenanceReport) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m PositionMaintenanceReport) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m PositionMaintenanceReport) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m PositionMaintenanceReport) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m PositionMaintenanceReport) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m PositionMaintenanceReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m PositionMaintenanceReport) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m PositionMaintenanceReport) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m PositionMaintenanceReport) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m PositionMaintenanceReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m PositionMaintenanceReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m PositionMaintenanceReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m PositionMaintenanceReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m PositionMaintenanceReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m PositionMaintenanceReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m PositionMaintenanceReport) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m PositionMaintenanceReport) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m PositionMaintenanceReport) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m PositionMaintenanceReport) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m PositionMaintenanceReport) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m PositionMaintenanceReport) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m PositionMaintenanceReport) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m PositionMaintenanceReport) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m PositionMaintenanceReport) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m PositionMaintenanceReport) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasNoLegs returns true if NoLegs is present, Tag 555
func (m PositionMaintenanceReport) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m PositionMaintenanceReport) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m PositionMaintenanceReport) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m PositionMaintenanceReport) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasPool returns true if Pool is present, Tag 691
func (m PositionMaintenanceReport) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasNoPositions returns true if NoPositions is present, Tag 702
func (m PositionMaintenanceReport) HasNoPositions() bool {
	return m.Has(tag.NoPositions)
}

//HasPosTransType returns true if PosTransType is present, Tag 709
func (m PositionMaintenanceReport) HasPosTransType() bool {
	return m.Has(tag.PosTransType)
}

//HasPosReqID returns true if PosReqID is present, Tag 710
func (m PositionMaintenanceReport) HasPosReqID() bool {
	return m.Has(tag.PosReqID)
}

//HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711
func (m PositionMaintenanceReport) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

//HasPosMaintAction returns true if PosMaintAction is present, Tag 712
func (m PositionMaintenanceReport) HasPosMaintAction() bool {
	return m.Has(tag.PosMaintAction)
}

//HasOrigPosReqRefID returns true if OrigPosReqRefID is present, Tag 713
func (m PositionMaintenanceReport) HasOrigPosReqRefID() bool {
	return m.Has(tag.OrigPosReqRefID)
}

//HasPosMaintRptRefID returns true if PosMaintRptRefID is present, Tag 714
func (m PositionMaintenanceReport) HasPosMaintRptRefID() bool {
	return m.Has(tag.PosMaintRptRefID)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m PositionMaintenanceReport) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasSettlSessID returns true if SettlSessID is present, Tag 716
func (m PositionMaintenanceReport) HasSettlSessID() bool {
	return m.Has(tag.SettlSessID)
}

//HasSettlSessSubID returns true if SettlSessSubID is present, Tag 717
func (m PositionMaintenanceReport) HasSettlSessSubID() bool {
	return m.Has(tag.SettlSessSubID)
}

//HasAdjustmentType returns true if AdjustmentType is present, Tag 718
func (m PositionMaintenanceReport) HasAdjustmentType() bool {
	return m.Has(tag.AdjustmentType)
}

//HasContraryInstructionIndicator returns true if ContraryInstructionIndicator is present, Tag 719
func (m PositionMaintenanceReport) HasContraryInstructionIndicator() bool {
	return m.Has(tag.ContraryInstructionIndicator)
}

//HasPriorSpreadIndicator returns true if PriorSpreadIndicator is present, Tag 720
func (m PositionMaintenanceReport) HasPriorSpreadIndicator() bool {
	return m.Has(tag.PriorSpreadIndicator)
}

//HasPosMaintRptID returns true if PosMaintRptID is present, Tag 721
func (m PositionMaintenanceReport) HasPosMaintRptID() bool {
	return m.Has(tag.PosMaintRptID)
}

//HasPosMaintStatus returns true if PosMaintStatus is present, Tag 722
func (m PositionMaintenanceReport) HasPosMaintStatus() bool {
	return m.Has(tag.PosMaintStatus)
}

//HasPosMaintResult returns true if PosMaintResult is present, Tag 723
func (m PositionMaintenanceReport) HasPosMaintResult() bool {
	return m.Has(tag.PosMaintResult)
}

//HasNoPosAmt returns true if NoPosAmt is present, Tag 753
func (m PositionMaintenanceReport) HasNoPosAmt() bool {
	return m.Has(tag.NoPosAmt)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m PositionMaintenanceReport) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasThresholdAmount returns true if ThresholdAmount is present, Tag 834
func (m PositionMaintenanceReport) HasThresholdAmount() bool {
	return m.Has(tag.ThresholdAmount)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m PositionMaintenanceReport) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m PositionMaintenanceReport) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m PositionMaintenanceReport) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m PositionMaintenanceReport) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m PositionMaintenanceReport) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m PositionMaintenanceReport) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m PositionMaintenanceReport) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m PositionMaintenanceReport) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m PositionMaintenanceReport) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m PositionMaintenanceReport) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m PositionMaintenanceReport) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m PositionMaintenanceReport) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m PositionMaintenanceReport) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m PositionMaintenanceReport) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m PositionMaintenanceReport) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m PositionMaintenanceReport) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m PositionMaintenanceReport) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m PositionMaintenanceReport) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m PositionMaintenanceReport) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m PositionMaintenanceReport) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m PositionMaintenanceReport) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m PositionMaintenanceReport) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m PositionMaintenanceReport) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m PositionMaintenanceReport) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m PositionMaintenanceReport) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m PositionMaintenanceReport) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m PositionMaintenanceReport) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m PositionMaintenanceReport) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayAmount returns true if OptPayAmount is present, Tag 1195
func (m PositionMaintenanceReport) HasOptPayAmount() bool {
	return m.Has(tag.OptPayAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m PositionMaintenanceReport) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasFuturesValuationMethod returns true if FuturesValuationMethod is present, Tag 1197
func (m PositionMaintenanceReport) HasFuturesValuationMethod() bool {
	return m.Has(tag.FuturesValuationMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m PositionMaintenanceReport) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m PositionMaintenanceReport) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m PositionMaintenanceReport) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m PositionMaintenanceReport) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m PositionMaintenanceReport) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m PositionMaintenanceReport) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID)})}
}

//Add create and append a new NoTradingSessions to this group
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

//Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
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

//SetNoPartySubIDs sets NoPartySubIDs, Tag 802
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
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

//GetNoPartySubIDs gets NoPartySubIDs, Tag 802
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

//NoPartySubIDs is a repeating group element, Tag 802
type NoPartySubIDs struct {
	quickfix.Group
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//SetPartySubIDType sets PartySubIDType, Tag 803
func (m NoPartySubIDs) SetPartySubIDType(v int) {
	m.Set(field.NewPartySubIDType(v))
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartySubIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubIDType gets PartySubIDType, Tag 803
func (m NoPartySubIDs) GetPartySubIDType() (f field.PartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//HasPartySubIDType returns true if PartySubIDType is present, Tag 803
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

//NoPartySubIDsRepeatingGroup is a repeating group, Tag 802
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

//Add create and append a new NoPartySubIDs to this group
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

//Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
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

//NoLegs is a repeating group element, Tag 555
type NoLegs struct {
	quickfix.Group
}

//SetLegSymbol sets LegSymbol, Tag 600
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

//SetLegSymbolSfx sets LegSymbolSfx, Tag 601
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

//SetLegSecurityID sets LegSecurityID, Tag 602
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

//SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

//SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetLegProduct sets LegProduct, Tag 607
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

//SetLegCFICode sets LegCFICode, Tag 608
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

//SetLegSecurityType sets LegSecurityType, Tag 609
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

//SetLegSecuritySubType sets LegSecuritySubType, Tag 764
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

//SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

//SetLegMaturityDate sets LegMaturityDate, Tag 611
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

//SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

//SetLegIssueDate sets LegIssueDate, Tag 249
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

//SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

//SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

//SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252
func (m NoLegs) SetLegRepurchaseRate(v float64) {
	m.Set(field.NewLegRepurchaseRate(v))
}

//SetLegFactor sets LegFactor, Tag 253
func (m NoLegs) SetLegFactor(v float64) {
	m.Set(field.NewLegFactor(v))
}

//SetLegCreditRating sets LegCreditRating, Tag 257
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

//SetLegInstrRegistry sets LegInstrRegistry, Tag 599
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

//SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

//SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

//SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

//SetLegRedemptionDate sets LegRedemptionDate, Tag 254
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

//SetLegStrikePrice sets LegStrikePrice, Tag 612
func (m NoLegs) SetLegStrikePrice(v float64) {
	m.Set(field.NewLegStrikePrice(v))
}

//SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

//SetLegOptAttribute sets LegOptAttribute, Tag 613
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

//SetLegContractMultiplier sets LegContractMultiplier, Tag 614
func (m NoLegs) SetLegContractMultiplier(v float64) {
	m.Set(field.NewLegContractMultiplier(v))
}

//SetLegCouponRate sets LegCouponRate, Tag 615
func (m NoLegs) SetLegCouponRate(v float64) {
	m.Set(field.NewLegCouponRate(v))
}

//SetLegSecurityExchange sets LegSecurityExchange, Tag 616
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

//SetLegIssuer sets LegIssuer, Tag 617
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

//SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

//SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

//SetLegSecurityDesc sets LegSecurityDesc, Tag 620
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

//SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

//SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

//SetLegRatioQty sets LegRatioQty, Tag 623
func (m NoLegs) SetLegRatioQty(v float64) {
	m.Set(field.NewLegRatioQty(v))
}

//SetLegSide sets LegSide, Tag 624
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

//SetLegCurrency sets LegCurrency, Tag 556
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

//SetLegPool sets LegPool, Tag 740
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

//SetLegDatedDate sets LegDatedDate, Tag 739
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

//SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

//SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

//SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

//SetLegTimeUnit sets LegTimeUnit, Tag 1001
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

//SetLegOptionRatio sets LegOptionRatio, Tag 1017
func (m NoLegs) SetLegOptionRatio(v float64) {
	m.Set(field.NewLegOptionRatio(v))
}

//SetLegPrice sets LegPrice, Tag 566
func (m NoLegs) SetLegPrice(v float64) {
	m.Set(field.NewLegPrice(v))
}

//SetLegMaturityTime sets LegMaturityTime, Tag 1212
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

//SetLegPutOrCall sets LegPutOrCall, Tag 1358
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

//SetLegExerciseStyle sets LegExerciseStyle, Tag 1420
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

//SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) SetLegUnitOfMeasureQty(v float64) {
	m.Set(field.NewLegUnitOfMeasureQty(v))
}

//SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

//SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) SetLegPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(v))
}

//GetLegSymbol gets LegSymbol, Tag 600
func (m NoLegs) GetLegSymbol() (f field.LegSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSymbolSfx gets LegSymbolSfx, Tag 601
func (m NoLegs) GetLegSymbolSfx() (f field.LegSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityID gets LegSecurityID, Tag 602
func (m NoLegs) GetLegSecurityID() (f field.LegSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603
func (m NoLegs) GetLegSecurityIDSource() (f field.LegSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLegProduct gets LegProduct, Tag 607
func (m NoLegs) GetLegProduct() (f field.LegProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCFICode gets LegCFICode, Tag 608
func (m NoLegs) GetLegCFICode() (f field.LegCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityType gets LegSecurityType, Tag 609
func (m NoLegs) GetLegSecurityType() (f field.LegSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecuritySubType gets LegSecuritySubType, Tag 764
func (m NoLegs) GetLegSecuritySubType() (f field.LegSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610
func (m NoLegs) GetLegMaturityMonthYear() (f field.LegMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityDate gets LegMaturityDate, Tag 611
func (m NoLegs) GetLegMaturityDate() (f field.LegMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248
func (m NoLegs) GetLegCouponPaymentDate() (f field.LegCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssueDate gets LegIssueDate, Tag 249
func (m NoLegs) GetLegIssueDate() (f field.LegIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250
func (m NoLegs) GetLegRepoCollateralSecurityType() (f field.LegRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251
func (m NoLegs) GetLegRepurchaseTerm() (f field.LegRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252
func (m NoLegs) GetLegRepurchaseRate() (f field.LegRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegFactor gets LegFactor, Tag 253
func (m NoLegs) GetLegFactor() (f field.LegFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCreditRating gets LegCreditRating, Tag 257
func (m NoLegs) GetLegCreditRating() (f field.LegCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInstrRegistry gets LegInstrRegistry, Tag 599
func (m NoLegs) GetLegInstrRegistry() (f field.LegInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596
func (m NoLegs) GetLegCountryOfIssue() (f field.LegCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597
func (m NoLegs) GetLegStateOrProvinceOfIssue() (f field.LegStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598
func (m NoLegs) GetLegLocaleOfIssue() (f field.LegLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRedemptionDate gets LegRedemptionDate, Tag 254
func (m NoLegs) GetLegRedemptionDate() (f field.LegRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikePrice gets LegStrikePrice, Tag 612
func (m NoLegs) GetLegStrikePrice() (f field.LegStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942
func (m NoLegs) GetLegStrikeCurrency() (f field.LegStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptAttribute gets LegOptAttribute, Tag 613
func (m NoLegs) GetLegOptAttribute() (f field.LegOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractMultiplier gets LegContractMultiplier, Tag 614
func (m NoLegs) GetLegContractMultiplier() (f field.LegContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCouponRate gets LegCouponRate, Tag 615
func (m NoLegs) GetLegCouponRate() (f field.LegCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityExchange gets LegSecurityExchange, Tag 616
func (m NoLegs) GetLegSecurityExchange() (f field.LegSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegIssuer gets LegIssuer, Tag 617
func (m NoLegs) GetLegIssuer() (f field.LegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618
func (m NoLegs) GetEncodedLegIssuerLen() (f field.EncodedLegIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619
func (m NoLegs) GetEncodedLegIssuer() (f field.EncodedLegIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityDesc gets LegSecurityDesc, Tag 620
func (m NoLegs) GetLegSecurityDesc() (f field.LegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621
func (m NoLegs) GetEncodedLegSecurityDescLen() (f field.EncodedLegSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622
func (m NoLegs) GetEncodedLegSecurityDesc() (f field.EncodedLegSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegRatioQty gets LegRatioQty, Tag 623
func (m NoLegs) GetLegRatioQty() (f field.LegRatioQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSide gets LegSide, Tag 624
func (m NoLegs) GetLegSide() (f field.LegSideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegCurrency gets LegCurrency, Tag 556
func (m NoLegs) GetLegCurrency() (f field.LegCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPool gets LegPool, Tag 740
func (m NoLegs) GetLegPool() (f field.LegPoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegDatedDate gets LegDatedDate, Tag 739
func (m NoLegs) GetLegDatedDate() (f field.LegDatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955
func (m NoLegs) GetLegContractSettlMonth() (f field.LegContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956
func (m NoLegs) GetLegInterestAccrualDate() (f field.LegInterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999
func (m NoLegs) GetLegUnitOfMeasure() (f field.LegUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegTimeUnit gets LegTimeUnit, Tag 1001
func (m NoLegs) GetLegTimeUnit() (f field.LegTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegOptionRatio gets LegOptionRatio, Tag 1017
func (m NoLegs) GetLegOptionRatio() (f field.LegOptionRatioField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPrice gets LegPrice, Tag 566
func (m NoLegs) GetLegPrice() (f field.LegPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegMaturityTime gets LegMaturityTime, Tag 1212
func (m NoLegs) GetLegMaturityTime() (f field.LegMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPutOrCall gets LegPutOrCall, Tag 1358
func (m NoLegs) GetLegPutOrCall() (f field.LegPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegExerciseStyle gets LegExerciseStyle, Tag 1420
func (m NoLegs) GetLegExerciseStyle() (f field.LegExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224
func (m NoLegs) GetLegUnitOfMeasureQty() (f field.LegUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421
func (m NoLegs) GetLegPriceUnitOfMeasure() (f field.LegPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (f field.LegPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSymbol returns true if LegSymbol is present, Tag 600
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

//HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

//HasLegSecurityID returns true if LegSecurityID is present, Tag 602
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

//HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

//HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

//HasLegProduct returns true if LegProduct is present, Tag 607
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

//HasLegCFICode returns true if LegCFICode is present, Tag 608
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

//HasLegSecurityType returns true if LegSecurityType is present, Tag 609
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

//HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

//HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

//HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

//HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

//HasLegIssueDate returns true if LegIssueDate is present, Tag 249
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

//HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

//HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

//HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

//HasLegFactor returns true if LegFactor is present, Tag 253
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

//HasLegCreditRating returns true if LegCreditRating is present, Tag 257
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

//HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

//HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

//HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

//HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

//HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

//HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

//HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

//HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

//HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

//HasLegCouponRate returns true if LegCouponRate is present, Tag 615
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

//HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

//HasLegIssuer returns true if LegIssuer is present, Tag 617
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

//HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

//HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

//HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

//HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

//HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

//HasLegRatioQty returns true if LegRatioQty is present, Tag 623
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

//HasLegSide returns true if LegSide is present, Tag 624
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

//HasLegCurrency returns true if LegCurrency is present, Tag 556
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

//HasLegPool returns true if LegPool is present, Tag 740
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

//HasLegDatedDate returns true if LegDatedDate is present, Tag 739
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

//HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

//HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

//HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

//HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

//HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

//HasLegPrice returns true if LegPrice is present, Tag 566
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

//HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

//HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

//HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

//HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

//HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

//HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

//NoLegSecurityAltID is a repeating group element, Tag 604
type NoLegSecurityAltID struct {
	quickfix.Group
}

//SetLegSecurityAltID sets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

//SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

//GetLegSecurityAltID gets LegSecurityAltID, Tag 605
func (m NoLegSecurityAltID) GetLegSecurityAltID() (f field.LegSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (f field.LegSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

//HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

//NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

//Add create and append a new NoLegSecurityAltID to this group
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

//Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoLegsRepeatingGroup is a repeating group, Tag 555
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty)})}
}

//Add create and append a new NoLegs to this group
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

//Get returns the ith NoLegs in the NoLegsRepeatinGroup
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

//NoPositions is a repeating group element, Tag 702
type NoPositions struct {
	quickfix.Group
}

//SetPosType sets PosType, Tag 703
func (m NoPositions) SetPosType(v string) {
	m.Set(field.NewPosType(v))
}

//SetLongQty sets LongQty, Tag 704
func (m NoPositions) SetLongQty(v float64) {
	m.Set(field.NewLongQty(v))
}

//SetShortQty sets ShortQty, Tag 705
func (m NoPositions) SetShortQty(v float64) {
	m.Set(field.NewShortQty(v))
}

//SetPosQtyStatus sets PosQtyStatus, Tag 706
func (m NoPositions) SetPosQtyStatus(v int) {
	m.Set(field.NewPosQtyStatus(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoPositions) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetQuantityDate sets QuantityDate, Tag 976
func (m NoPositions) SetQuantityDate(v string) {
	m.Set(field.NewQuantityDate(v))
}

//GetPosType gets PosType, Tag 703
func (m NoPositions) GetPosType() (f field.PosTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLongQty gets LongQty, Tag 704
func (m NoPositions) GetLongQty() (f field.LongQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShortQty gets ShortQty, Tag 705
func (m NoPositions) GetShortQty() (f field.ShortQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosQtyStatus gets PosQtyStatus, Tag 706
func (m NoPositions) GetPosQtyStatus() (f field.PosQtyStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoPositions) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetQuantityDate gets QuantityDate, Tag 976
func (m NoPositions) GetQuantityDate() (f field.QuantityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPosType returns true if PosType is present, Tag 703
func (m NoPositions) HasPosType() bool {
	return m.Has(tag.PosType)
}

//HasLongQty returns true if LongQty is present, Tag 704
func (m NoPositions) HasLongQty() bool {
	return m.Has(tag.LongQty)
}

//HasShortQty returns true if ShortQty is present, Tag 705
func (m NoPositions) HasShortQty() bool {
	return m.Has(tag.ShortQty)
}

//HasPosQtyStatus returns true if PosQtyStatus is present, Tag 706
func (m NoPositions) HasPosQtyStatus() bool {
	return m.Has(tag.PosQtyStatus)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoPositions) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasQuantityDate returns true if QuantityDate is present, Tag 976
func (m NoPositions) HasQuantityDate() bool {
	return m.Has(tag.QuantityDate)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNoNestedPartySubIDs sets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) SetNoNestedPartySubIDs(f NoNestedPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (f field.NestedPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (f field.NestedPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (f field.NestedPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoNestedPartySubIDs gets NoNestedPartySubIDs, Tag 804
func (m NoNestedPartyIDs) GetNoNestedPartySubIDs() (NoNestedPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNoNestedPartySubIDs returns true if NoNestedPartySubIDs is present, Tag 804
func (m NoNestedPartyIDs) HasNoNestedPartySubIDs() bool {
	return m.Has(tag.NoNestedPartySubIDs)
}

//NoNestedPartySubIDs is a repeating group element, Tag 804
type NoNestedPartySubIDs struct {
	quickfix.Group
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//SetNestedPartySubIDType sets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) SetNestedPartySubIDType(v int) {
	m.Set(field.NewNestedPartySubIDType(v))
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartySubIDs) GetNestedPartySubID() (f field.NestedPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNestedPartySubIDType gets NestedPartySubIDType, Tag 805
func (m NoNestedPartySubIDs) GetNestedPartySubIDType() (f field.NestedPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartySubIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//HasNestedPartySubIDType returns true if NestedPartySubIDType is present, Tag 805
func (m NoNestedPartySubIDs) HasNestedPartySubIDType() bool {
	return m.Has(tag.NestedPartySubIDType)
}

//NoNestedPartySubIDsRepeatingGroup is a repeating group, Tag 804
type NoNestedPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartySubIDsRepeatingGroup returns an initialized, NoNestedPartySubIDsRepeatingGroup
func NewNoNestedPartySubIDsRepeatingGroup() NoNestedPartySubIDsRepeatingGroup {
	return NoNestedPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartySubID), quickfix.GroupElement(tag.NestedPartySubIDType)})}
}

//Add create and append a new NoNestedPartySubIDs to this group
func (m NoNestedPartySubIDsRepeatingGroup) Add() NoNestedPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartySubIDs{g}
}

//Get returns the ith NoNestedPartySubIDs in the NoNestedPartySubIDsRepeatinGroup
func (m NoNestedPartySubIDsRepeatingGroup) Get(i int) NoNestedPartySubIDs {
	return NoNestedPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), NewNoNestedPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoPositionsRepeatingGroup is a repeating group, Tag 702
type NoPositionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPositionsRepeatingGroup returns an initialized, NoPositionsRepeatingGroup
func NewNoPositionsRepeatingGroup() NoPositionsRepeatingGroup {
	return NoPositionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPositions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PosType), quickfix.GroupElement(tag.LongQty), quickfix.GroupElement(tag.ShortQty), quickfix.GroupElement(tag.PosQtyStatus), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.QuantityDate)})}
}

//Add create and append a new NoPositions to this group
func (m NoPositionsRepeatingGroup) Add() NoPositions {
	g := m.RepeatingGroup.Add()
	return NoPositions{g}
}

//Get returns the ith NoPositions in the NoPositionsRepeatinGroup
func (m NoPositionsRepeatingGroup) Get(i int) NoPositions {
	return NoPositions{m.RepeatingGroup.Get(i)}
}

//NoUnderlyings is a repeating group element, Tag 711
type NoUnderlyings struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

//SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingProduct sets UnderlyingProduct, Tag 462
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

//SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

//SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

//SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

//SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

//SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

//SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(v float64) {
	m.Set(field.NewUnderlyingRepurchaseRate(v))
}

//SetUnderlyingFactor sets UnderlyingFactor, Tag 246
func (m NoUnderlyings) SetUnderlyingFactor(v float64) {
	m.Set(field.NewUnderlyingFactor(v))
}

//SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

//SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

//SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

//SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

//SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

//SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) SetUnderlyingStrikePrice(v float64) {
	m.Set(field.NewUnderlyingStrikePrice(v))
}

//SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) SetUnderlyingContractMultiplier(v float64) {
	m.Set(field.NewUnderlyingContractMultiplier(v))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) SetUnderlyingCouponRate(v float64) {
	m.Set(field.NewUnderlyingCouponRate(v))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

//SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//SetUnderlyingQty sets UnderlyingQty, Tag 879
func (m NoUnderlyings) SetUnderlyingQty(v float64) {
	m.Set(field.NewUnderlyingQty(v))
}

//SetUnderlyingPx sets UnderlyingPx, Tag 810
func (m NoUnderlyings) SetUnderlyingPx(v float64) {
	m.Set(field.NewUnderlyingPx(v))
}

//SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) SetUnderlyingDirtyPrice(v float64) {
	m.Set(field.NewUnderlyingDirtyPrice(v))
}

//SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) SetUnderlyingEndPrice(v float64) {
	m.Set(field.NewUnderlyingEndPrice(v))
}

//SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) SetUnderlyingStartValue(v float64) {
	m.Set(field.NewUnderlyingStartValue(v))
}

//SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) SetUnderlyingCurrentValue(v float64) {
	m.Set(field.NewUnderlyingCurrentValue(v))
}

//SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) SetUnderlyingEndValue(v float64) {
	m.Set(field.NewUnderlyingEndValue(v))
}

//SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) SetUnderlyingAllocationPercent(v float64) {
	m.Set(field.NewUnderlyingAllocationPercent(v))
}

//SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) SetUnderlyingSettlementType(v int) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

//SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) SetUnderlyingCashAmount(v float64) {
	m.Set(field.NewUnderlyingCashAmount(v))
}

//SetUnderlyingCashType sets UnderlyingCashType, Tag 974
func (m NoUnderlyings) SetUnderlyingCashType(v string) {
	m.Set(field.NewUnderlyingCashType(v))
}

//SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

//SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

//SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) SetUnderlyingCapValue(v float64) {
	m.Set(field.NewUnderlyingCapValue(v))
}

//SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

//SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(v float64) {
	m.Set(field.NewUnderlyingAdjustedQuantity(v))
}

//SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) SetUnderlyingFXRate(v float64) {
	m.Set(field.NewUnderlyingFXRate(v))
}

//SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v string) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

//SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

//SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(v))
}

//SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

//SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(v))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoUnderlyings) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoUnderlyings) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (f field.UnderlyingSecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingProduct gets UnderlyingProduct, Tag 462
func (m NoUnderlyings) GetUnderlyingProduct() (f field.UnderlyingProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463
func (m NoUnderlyings) GetUnderlyingCFICode() (f field.UnderlyingCFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoUnderlyings) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (f field.UnderlyingSecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542
func (m NoUnderlyings) GetUnderlyingMaturityDate() (f field.UnderlyingMaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (f field.UnderlyingCouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242
func (m NoUnderlyings) GetUnderlyingIssueDate() (f field.UnderlyingIssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (f field.UnderlyingRepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (f field.UnderlyingRepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (f field.UnderlyingRepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFactor gets UnderlyingFactor, Tag 246
func (m NoUnderlyings) GetUnderlyingFactor() (f field.UnderlyingFactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256
func (m NoUnderlyings) GetUnderlyingCreditRating() (f field.UnderlyingCreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (f field.UnderlyingInstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (f field.UnderlyingCountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (f field.UnderlyingStateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (f field.UnderlyingLocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (f field.UnderlyingRedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoUnderlyings) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (f field.UnderlyingStrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoUnderlyings) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoUnderlyings) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoUnderlyings) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877
func (m NoUnderlyings) GetUnderlyingCPProgram() (f field.UnderlyingCPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878
func (m NoUnderlyings) GetUnderlyingCPRegType() (f field.UnderlyingCPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoUnderlyings) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingQty gets UnderlyingQty, Tag 879
func (m NoUnderlyings) GetUnderlyingQty() (f field.UnderlyingQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPx gets UnderlyingPx, Tag 810
func (m NoUnderlyings) GetUnderlyingPx() (f field.UnderlyingPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (f field.UnderlyingDirtyPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883
func (m NoUnderlyings) GetUnderlyingEndPrice() (f field.UnderlyingEndPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884
func (m NoUnderlyings) GetUnderlyingStartValue() (f field.UnderlyingStartValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885
func (m NoUnderlyings) GetUnderlyingCurrentValue() (f field.UnderlyingCurrentValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886
func (m NoUnderlyings) GetUnderlyingEndValue() (f field.UnderlyingEndValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (f field.UnderlyingAllocationPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975
func (m NoUnderlyings) GetUnderlyingSettlementType() (f field.UnderlyingSettlementTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973
func (m NoUnderlyings) GetUnderlyingCashAmount() (f field.UnderlyingCashAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCashType gets UnderlyingCashType, Tag 974
func (m NoUnderlyings) GetUnderlyingCashType() (f field.UnderlyingCashTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (f field.UnderlyingUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000
func (m NoUnderlyings) GetUnderlyingTimeUnit() (f field.UnderlyingTimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038
func (m NoUnderlyings) GetUnderlyingCapValue() (f field.UnderlyingCapValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039
func (m NoUnderlyings) GetUnderlyingSettlMethod() (f field.UnderlyingSettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (f field.UnderlyingAdjustedQuantityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045
func (m NoUnderlyings) GetUnderlyingFXRate() (f field.UnderlyingFXRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (f field.UnderlyingFXRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213
func (m NoUnderlyings) GetUnderlyingMaturityTime() (f field.UnderlyingMaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoUnderlyings) GetUnderlyingPutOrCall() (f field.UnderlyingPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (f field.UnderlyingExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (f field.UnderlyingUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (f field.UnderlyingPriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (f field.UnderlyingPriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

//HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

//HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

//HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

//HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

//HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

//HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

//HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

//HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

//HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

//HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

//HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

//HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

//HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

//HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

//HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

//HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

//HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

//HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

//HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

//HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

//HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

//HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

//HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

//HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

//HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

//HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

//HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

//HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

//HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

//HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

//HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

//HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

//HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

//HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

//HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

//HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

//HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

//HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

//HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

//NoUnderlyingSecurityAltID is a repeating group element, Tag 457
type NoUnderlyingSecurityAltID struct {
	quickfix.Group
}

//SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

//SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

//GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (f field.UnderlyingSecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (f field.UnderlyingSecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

//HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

//NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

//Add create and append a new NoUnderlyingSecurityAltID to this group
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

//Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingStips is a repeating group element, Tag 887
type NoUnderlyingStips struct {
	quickfix.Group
}

//SetUnderlyingStipType sets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

//SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

//GetUnderlyingStipType gets UnderlyingStipType, Tag 888
func (m NoUnderlyingStips) GetUnderlyingStipType() (f field.UnderlyingStipTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889
func (m NoUnderlyingStips) GetUnderlyingStipValue() (f field.UnderlyingStipValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

//HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

//NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

//Add create and append a new NoUnderlyingStips to this group
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

//Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentParties is a repeating group element, Tag 1058
type NoUndlyInstrumentParties struct {
	quickfix.Group
}

//SetUndlyInstrumentPartyID sets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyID(v string) {
	m.Set(field.NewUndlyInstrumentPartyID(v))
}

//SetUndlyInstrumentPartyIDSource sets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyIDSource(v string) {
	m.Set(field.NewUndlyInstrumentPartyIDSource(v))
}

//SetUndlyInstrumentPartyRole sets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) SetUndlyInstrumentPartyRole(v int) {
	m.Set(field.NewUndlyInstrumentPartyRole(v))
}

//SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetUndlyInstrumentPartyID gets UndlyInstrumentPartyID, Tag 1059
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyID() (f field.UndlyInstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyIDSource gets UndlyInstrumentPartyIDSource, Tag 1060
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyIDSource() (f field.UndlyInstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartyRole gets UndlyInstrumentPartyRole, Tag 1061
func (m NoUndlyInstrumentParties) GetUndlyInstrumentPartyRole() (f field.UndlyInstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasUndlyInstrumentPartyID returns true if UndlyInstrumentPartyID is present, Tag 1059
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyID() bool {
	return m.Has(tag.UndlyInstrumentPartyID)
}

//HasUndlyInstrumentPartyIDSource returns true if UndlyInstrumentPartyIDSource is present, Tag 1060
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyIDSource() bool {
	return m.Has(tag.UndlyInstrumentPartyIDSource)
}

//HasUndlyInstrumentPartyRole returns true if UndlyInstrumentPartyRole is present, Tag 1061
func (m NoUndlyInstrumentParties) HasUndlyInstrumentPartyRole() bool {
	return m.Has(tag.UndlyInstrumentPartyRole)
}

//HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

//NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062
type NoUndlyInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetUndlyInstrumentPartySubID sets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubID(v string) {
	m.Set(field.NewUndlyInstrumentPartySubID(v))
}

//SetUndlyInstrumentPartySubIDType sets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) SetUndlyInstrumentPartySubIDType(v int) {
	m.Set(field.NewUndlyInstrumentPartySubIDType(v))
}

//GetUndlyInstrumentPartySubID gets UndlyInstrumentPartySubID, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubID() (f field.UndlyInstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUndlyInstrumentPartySubIDType gets UndlyInstrumentPartySubIDType, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) GetUndlyInstrumentPartySubIDType() (f field.UndlyInstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUndlyInstrumentPartySubID returns true if UndlyInstrumentPartySubID is present, Tag 1063
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubID() bool {
	return m.Has(tag.UndlyInstrumentPartySubID)
}

//HasUndlyInstrumentPartySubIDType returns true if UndlyInstrumentPartySubIDType is present, Tag 1064
func (m NoUndlyInstrumentPartySubIDs) HasUndlyInstrumentPartySubIDType() bool {
	return m.Has(tag.UndlyInstrumentPartySubIDType)
}

//NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartySubID), quickfix.GroupElement(tag.UndlyInstrumentPartySubIDType)})}
}

//Add create and append a new NoUndlyInstrumentPartySubIDs to this group
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

//Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UndlyInstrumentPartyID), quickfix.GroupElement(tag.UndlyInstrumentPartyIDSource), quickfix.GroupElement(tag.UndlyInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoUndlyInstrumentParties to this group
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

//Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

//NoUnderlyingsRepeatingGroup is a repeating group, Tag 711
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty)})}
}

//Add create and append a new NoUnderlyings to this group
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

//Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

//NoPosAmt is a repeating group element, Tag 753
type NoPosAmt struct {
	quickfix.Group
}

//SetPosAmtType sets PosAmtType, Tag 707
func (m NoPosAmt) SetPosAmtType(v string) {
	m.Set(field.NewPosAmtType(v))
}

//SetPosAmt sets PosAmt, Tag 708
func (m NoPosAmt) SetPosAmt(v float64) {
	m.Set(field.NewPosAmt(v))
}

//SetPositionCurrency sets PositionCurrency, Tag 1055
func (m NoPosAmt) SetPositionCurrency(v string) {
	m.Set(field.NewPositionCurrency(v))
}

//GetPosAmtType gets PosAmtType, Tag 707
func (m NoPosAmt) GetPosAmtType() (f field.PosAmtTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPosAmt gets PosAmt, Tag 708
func (m NoPosAmt) GetPosAmt() (f field.PosAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionCurrency gets PositionCurrency, Tag 1055
func (m NoPosAmt) GetPositionCurrency() (f field.PositionCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPosAmtType returns true if PosAmtType is present, Tag 707
func (m NoPosAmt) HasPosAmtType() bool {
	return m.Has(tag.PosAmtType)
}

//HasPosAmt returns true if PosAmt is present, Tag 708
func (m NoPosAmt) HasPosAmt() bool {
	return m.Has(tag.PosAmt)
}

//HasPositionCurrency returns true if PositionCurrency is present, Tag 1055
func (m NoPosAmt) HasPositionCurrency() bool {
	return m.Has(tag.PositionCurrency)
}

//NoPosAmtRepeatingGroup is a repeating group, Tag 753
type NoPosAmtRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPosAmtRepeatingGroup returns an initialized, NoPosAmtRepeatingGroup
func NewNoPosAmtRepeatingGroup() NoPosAmtRepeatingGroup {
	return NoPosAmtRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPosAmt,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PosAmtType), quickfix.GroupElement(tag.PosAmt), quickfix.GroupElement(tag.PositionCurrency)})}
}

//Add create and append a new NoPosAmt to this group
func (m NoPosAmtRepeatingGroup) Add() NoPosAmt {
	g := m.RepeatingGroup.Add()
	return NoPosAmt{g}
}

//Get returns the ith NoPosAmt in the NoPosAmtRepeatinGroup
func (m NoPosAmtRepeatingGroup) Get(i int) NoPosAmt {
	return NoPosAmt{m.RepeatingGroup.Get(i)}
}

//NoEvents is a repeating group element, Tag 864
type NoEvents struct {
	quickfix.Group
}

//SetEventType sets EventType, Tag 865
func (m NoEvents) SetEventType(v int) {
	m.Set(field.NewEventType(v))
}

//SetEventDate sets EventDate, Tag 866
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

//SetEventPx sets EventPx, Tag 867
func (m NoEvents) SetEventPx(v float64) {
	m.Set(field.NewEventPx(v))
}

//SetEventText sets EventText, Tag 868
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

//SetEventTime sets EventTime, Tag 1145
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

//GetEventType gets EventType, Tag 865
func (m NoEvents) GetEventType() (f field.EventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventDate gets EventDate, Tag 866
func (m NoEvents) GetEventDate() (f field.EventDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventPx gets EventPx, Tag 867
func (m NoEvents) GetEventPx() (f field.EventPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventText gets EventText, Tag 868
func (m NoEvents) GetEventText() (f field.EventTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEventTime gets EventTime, Tag 1145
func (m NoEvents) GetEventTime() (f field.EventTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasEventType returns true if EventType is present, Tag 865
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

//HasEventDate returns true if EventDate is present, Tag 866
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

//HasEventPx returns true if EventPx is present, Tag 867
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

//HasEventText returns true if EventText is present, Tag 868
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

//HasEventTime returns true if EventTime is present, Tag 1145
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

//NoEventsRepeatingGroup is a repeating group, Tag 864
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

//Add create and append a new NoEvents to this group
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

//Get returns the ith NoEvents in the NoEventsRepeatinGroup
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

//NoInstrumentParties is a repeating group element, Tag 1018
type NoInstrumentParties struct {
	quickfix.Group
}

//SetInstrumentPartyID sets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

//SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

//SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

//SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetInstrumentPartyID gets InstrumentPartyID, Tag 1019
func (m NoInstrumentParties) GetInstrumentPartyID() (f field.InstrumentPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (f field.InstrumentPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051
func (m NoInstrumentParties) GetInstrumentPartyRole() (f field.InstrumentPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

//HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

//HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

//HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

//NoInstrumentPartySubIDs is a repeating group element, Tag 1052
type NoInstrumentPartySubIDs struct {
	quickfix.Group
}

//SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

//SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

//GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (f field.InstrumentPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (f field.InstrumentPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

//HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

//NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

//Add create and append a new NoInstrumentPartySubIDs to this group
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

//Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

//Add create and append a new NoInstrumentParties to this group
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

//Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}
