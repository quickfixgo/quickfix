package settlementobligationreport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//SettlementObligationReport is the fix50sp1 SettlementObligationReport type, MsgType = BQ
type SettlementObligationReport struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a SettlementObligationReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) SettlementObligationReport {
	return SettlementObligationReport{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SettlementObligationReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a SettlementObligationReport initialized with the required fields for SettlementObligationReport
func New(settlobligmsgid field.SettlObligMsgIDField, settlobligmode field.SettlObligModeField) (m SettlementObligationReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BQ"))
	m.Set(settlobligmsgid)
	m.Set(settlobligmode)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SettlementObligationReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "BQ", r
}

//SetText sets Text, Tag 58
func (m SettlementObligationReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m SettlementObligationReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SettlementObligationReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SettlementObligationReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetClearingBusinessDate sets ClearingBusinessDate, Tag 715
func (m SettlementObligationReport) SetClearingBusinessDate(v string) {
	m.Set(field.NewClearingBusinessDate(v))
}

//SetSettlementCycleNo sets SettlementCycleNo, Tag 1153
func (m SettlementObligationReport) SetSettlementCycleNo(v int) {
	m.Set(field.NewSettlementCycleNo(v))
}

//SetSettlObligMode sets SettlObligMode, Tag 1159
func (m SettlementObligationReport) SetSettlObligMode(v int) {
	m.Set(field.NewSettlObligMode(v))
}

//SetSettlObligMsgID sets SettlObligMsgID, Tag 1160
func (m SettlementObligationReport) SetSettlObligMsgID(v string) {
	m.Set(field.NewSettlObligMsgID(v))
}

//SetNoSettlOblig sets NoSettlOblig, Tag 1165
func (m SettlementObligationReport) SetNoSettlOblig(f NoSettlObligRepeatingGroup) {
	m.SetGroup(f)
}

//SetApplID sets ApplID, Tag 1180
func (m SettlementObligationReport) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m SettlementObligationReport) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m SettlementObligationReport) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m SettlementObligationReport) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//GetText gets Text, Tag 58
func (m SettlementObligationReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SettlementObligationReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SettlementObligationReport) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SettlementObligationReport) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingBusinessDate gets ClearingBusinessDate, Tag 715
func (m SettlementObligationReport) GetClearingBusinessDate() (f field.ClearingBusinessDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlementCycleNo gets SettlementCycleNo, Tag 1153
func (m SettlementObligationReport) GetSettlementCycleNo() (f field.SettlementCycleNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlObligMode gets SettlObligMode, Tag 1159
func (m SettlementObligationReport) GetSettlObligMode() (f field.SettlObligModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlObligMsgID gets SettlObligMsgID, Tag 1160
func (m SettlementObligationReport) GetSettlObligMsgID() (f field.SettlObligMsgIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlOblig gets NoSettlOblig, Tag 1165
func (m SettlementObligationReport) GetNoSettlOblig() (NoSettlObligRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlObligRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetApplID gets ApplID, Tag 1180
func (m SettlementObligationReport) GetApplID() (f field.ApplIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m SettlementObligationReport) GetApplSeqNum() (f field.ApplSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m SettlementObligationReport) GetApplLastSeqNum() (f field.ApplLastSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m SettlementObligationReport) GetApplResendFlag() (f field.ApplResendFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m SettlementObligationReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m SettlementObligationReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SettlementObligationReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SettlementObligationReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasClearingBusinessDate returns true if ClearingBusinessDate is present, Tag 715
func (m SettlementObligationReport) HasClearingBusinessDate() bool {
	return m.Has(tag.ClearingBusinessDate)
}

//HasSettlementCycleNo returns true if SettlementCycleNo is present, Tag 1153
func (m SettlementObligationReport) HasSettlementCycleNo() bool {
	return m.Has(tag.SettlementCycleNo)
}

//HasSettlObligMode returns true if SettlObligMode is present, Tag 1159
func (m SettlementObligationReport) HasSettlObligMode() bool {
	return m.Has(tag.SettlObligMode)
}

//HasSettlObligMsgID returns true if SettlObligMsgID is present, Tag 1160
func (m SettlementObligationReport) HasSettlObligMsgID() bool {
	return m.Has(tag.SettlObligMsgID)
}

//HasNoSettlOblig returns true if NoSettlOblig is present, Tag 1165
func (m SettlementObligationReport) HasNoSettlOblig() bool {
	return m.Has(tag.NoSettlOblig)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m SettlementObligationReport) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m SettlementObligationReport) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m SettlementObligationReport) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m SettlementObligationReport) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//NoSettlOblig is a repeating group element, Tag 1165
type NoSettlOblig struct {
	quickfix.Group
}

//SetNetGrossInd sets NetGrossInd, Tag 430
func (m NoSettlOblig) SetNetGrossInd(v int) {
	m.Set(field.NewNetGrossInd(v))
}

//SetSettlObligID sets SettlObligID, Tag 1161
func (m NoSettlOblig) SetSettlObligID(v string) {
	m.Set(field.NewSettlObligID(v))
}

//SetSettlObligTransType sets SettlObligTransType, Tag 1162
func (m NoSettlOblig) SetSettlObligTransType(v string) {
	m.Set(field.NewSettlObligTransType(v))
}

//SetSettlObligRefID sets SettlObligRefID, Tag 1163
func (m NoSettlOblig) SetSettlObligRefID(v string) {
	m.Set(field.NewSettlObligRefID(v))
}

//SetCcyAmt sets CcyAmt, Tag 1157
func (m NoSettlOblig) SetCcyAmt(v float64) {
	m.Set(field.NewCcyAmt(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoSettlOblig) SetSettlCurrAmt(v float64) {
	m.Set(field.NewSettlCurrAmt(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoSettlOblig) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoSettlOblig) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoSettlOblig) SetSettlCurrFxRate(v float64) {
	m.Set(field.NewSettlCurrFxRate(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m NoSettlOblig) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NoSettlOblig) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoSettlOblig) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoSettlOblig) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NoSettlOblig) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NoSettlOblig) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NoSettlOblig) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoSettlOblig) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoSettlOblig) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m NoSettlOblig) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoSettlOblig) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NoSettlOblig) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NoSettlOblig) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NoSettlOblig) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NoSettlOblig) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NoSettlOblig) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NoSettlOblig) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m NoSettlOblig) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NoSettlOblig) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NoSettlOblig) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NoSettlOblig) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NoSettlOblig) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NoSettlOblig) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NoSettlOblig) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoSettlOblig) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m NoSettlOblig) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoSettlOblig) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoSettlOblig) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoSettlOblig) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoSettlOblig) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoSettlOblig) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoSettlOblig) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoSettlOblig) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoSettlOblig) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoSettlOblig) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoSettlOblig) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetPool sets Pool, Tag 691
func (m NoSettlOblig) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m NoSettlOblig) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m NoSettlOblig) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m NoSettlOblig) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m NoSettlOblig) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m NoSettlOblig) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m NoSettlOblig) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m NoSettlOblig) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m NoSettlOblig) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m NoSettlOblig) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m NoSettlOblig) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m NoSettlOblig) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m NoSettlOblig) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m NoSettlOblig) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m NoSettlOblig) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m NoSettlOblig) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m NoSettlOblig) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m NoSettlOblig) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m NoSettlOblig) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m NoSettlOblig) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m NoSettlOblig) SetMinPriceIncrementAmount(v float64) {
	m.Set(field.NewMinPriceIncrementAmount(v))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m NoSettlOblig) SetUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnitOfMeasureQty(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m NoSettlOblig) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m NoSettlOblig) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m NoSettlOblig) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m NoSettlOblig) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m NoSettlOblig) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m NoSettlOblig) SetPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewPriceUnitOfMeasureQty(v))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m NoSettlOblig) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m NoSettlOblig) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayAmount sets OptPayAmount, Tag 1195
func (m NoSettlOblig) SetOptPayAmount(v float64) {
	m.Set(field.NewOptPayAmount(v))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m NoSettlOblig) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m NoSettlOblig) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m NoSettlOblig) SetCapPrice(v float64) {
	m.Set(field.NewCapPrice(v))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m NoSettlOblig) SetFloorPrice(v float64) {
	m.Set(field.NewFloorPrice(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoSettlOblig) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m NoSettlOblig) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m NoSettlOblig) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetFuturesValuationMethod sets FuturesValuationMethod, Tag 1197
func (m NoSettlOblig) SetFuturesValuationMethod(v string) {
	m.Set(field.NewFuturesValuationMethod(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoSettlOblig) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m NoSettlOblig) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NoSettlOblig) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetLastUpdateTime sets LastUpdateTime, Tag 779
func (m NoSettlOblig) SetLastUpdateTime(v time.Time) {
	m.Set(field.NewLastUpdateTime(v))
}

//SetNoSettlDetails sets NoSettlDetails, Tag 1158
func (m NoSettlOblig) SetNoSettlDetails(f NoSettlDetailsRepeatingGroup) {
	m.SetGroup(f)
}

//GetNetGrossInd gets NetGrossInd, Tag 430
func (m NoSettlOblig) GetNetGrossInd() (f field.NetGrossIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlObligID gets SettlObligID, Tag 1161
func (m NoSettlOblig) GetSettlObligID() (f field.SettlObligIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlObligTransType gets SettlObligTransType, Tag 1162
func (m NoSettlOblig) GetSettlObligTransType() (f field.SettlObligTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlObligRefID gets SettlObligRefID, Tag 1163
func (m NoSettlOblig) GetSettlObligRefID() (f field.SettlObligRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCcyAmt gets CcyAmt, Tag 1157
func (m NoSettlOblig) GetCcyAmt() (f field.CcyAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoSettlOblig) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoSettlOblig) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoSettlOblig) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoSettlOblig) GetSettlCurrFxRate() (f field.SettlCurrFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m NoSettlOblig) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NoSettlOblig) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoSettlOblig) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoSettlOblig) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NoSettlOblig) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NoSettlOblig) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NoSettlOblig) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoSettlOblig) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoSettlOblig) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m NoSettlOblig) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoSettlOblig) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m NoSettlOblig) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NoSettlOblig) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NoSettlOblig) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NoSettlOblig) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NoSettlOblig) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NoSettlOblig) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m NoSettlOblig) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NoSettlOblig) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NoSettlOblig) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NoSettlOblig) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NoSettlOblig) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NoSettlOblig) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NoSettlOblig) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoSettlOblig) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m NoSettlOblig) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoSettlOblig) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoSettlOblig) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoSettlOblig) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoSettlOblig) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoSettlOblig) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoSettlOblig) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoSettlOblig) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoSettlOblig) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoSettlOblig) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoSettlOblig) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPool gets Pool, Tag 691
func (m NoSettlOblig) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m NoSettlOblig) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m NoSettlOblig) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m NoSettlOblig) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m NoSettlOblig) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m NoSettlOblig) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m NoSettlOblig) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m NoSettlOblig) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m NoSettlOblig) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m NoSettlOblig) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m NoSettlOblig) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m NoSettlOblig) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m NoSettlOblig) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m NoSettlOblig) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m NoSettlOblig) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m NoSettlOblig) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m NoSettlOblig) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m NoSettlOblig) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m NoSettlOblig) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m NoSettlOblig) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m NoSettlOblig) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m NoSettlOblig) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m NoSettlOblig) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m NoSettlOblig) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m NoSettlOblig) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m NoSettlOblig) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m NoSettlOblig) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m NoSettlOblig) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m NoSettlOblig) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m NoSettlOblig) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayAmount gets OptPayAmount, Tag 1195
func (m NoSettlOblig) GetOptPayAmount() (f field.OptPayAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m NoSettlOblig) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m NoSettlOblig) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m NoSettlOblig) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m NoSettlOblig) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoSettlOblig) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m NoSettlOblig) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m NoSettlOblig) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFuturesValuationMethod gets FuturesValuationMethod, Tag 1197
func (m NoSettlOblig) GetFuturesValuationMethod() (f field.FuturesValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoSettlOblig) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m NoSettlOblig) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoSettlOblig) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastUpdateTime gets LastUpdateTime, Tag 779
func (m NoSettlOblig) GetLastUpdateTime() (f field.LastUpdateTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlDetails gets NoSettlDetails, Tag 1158
func (m NoSettlOblig) GetNoSettlDetails() (NoSettlDetailsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlDetailsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNetGrossInd returns true if NetGrossInd is present, Tag 430
func (m NoSettlOblig) HasNetGrossInd() bool {
	return m.Has(tag.NetGrossInd)
}

//HasSettlObligID returns true if SettlObligID is present, Tag 1161
func (m NoSettlOblig) HasSettlObligID() bool {
	return m.Has(tag.SettlObligID)
}

//HasSettlObligTransType returns true if SettlObligTransType is present, Tag 1162
func (m NoSettlOblig) HasSettlObligTransType() bool {
	return m.Has(tag.SettlObligTransType)
}

//HasSettlObligRefID returns true if SettlObligRefID is present, Tag 1163
func (m NoSettlOblig) HasSettlObligRefID() bool {
	return m.Has(tag.SettlObligRefID)
}

//HasCcyAmt returns true if CcyAmt is present, Tag 1157
func (m NoSettlOblig) HasCcyAmt() bool {
	return m.Has(tag.CcyAmt)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m NoSettlOblig) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoSettlOblig) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoSettlOblig) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m NoSettlOblig) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m NoSettlOblig) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoSettlOblig) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoSettlOblig) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoSettlOblig) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NoSettlOblig) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NoSettlOblig) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoSettlOblig) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoSettlOblig) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoSettlOblig) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m NoSettlOblig) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoSettlOblig) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NoSettlOblig) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NoSettlOblig) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NoSettlOblig) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NoSettlOblig) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NoSettlOblig) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NoSettlOblig) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NoSettlOblig) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NoSettlOblig) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NoSettlOblig) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NoSettlOblig) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NoSettlOblig) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NoSettlOblig) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NoSettlOblig) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoSettlOblig) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m NoSettlOblig) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoSettlOblig) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoSettlOblig) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoSettlOblig) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoSettlOblig) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoSettlOblig) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoSettlOblig) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoSettlOblig) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoSettlOblig) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoSettlOblig) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoSettlOblig) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasPool returns true if Pool is present, Tag 691
func (m NoSettlOblig) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m NoSettlOblig) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m NoSettlOblig) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m NoSettlOblig) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m NoSettlOblig) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m NoSettlOblig) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m NoSettlOblig) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m NoSettlOblig) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m NoSettlOblig) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m NoSettlOblig) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m NoSettlOblig) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m NoSettlOblig) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m NoSettlOblig) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m NoSettlOblig) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m NoSettlOblig) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m NoSettlOblig) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m NoSettlOblig) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m NoSettlOblig) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m NoSettlOblig) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m NoSettlOblig) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m NoSettlOblig) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m NoSettlOblig) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m NoSettlOblig) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m NoSettlOblig) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m NoSettlOblig) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m NoSettlOblig) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m NoSettlOblig) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m NoSettlOblig) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m NoSettlOblig) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m NoSettlOblig) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayAmount returns true if OptPayAmount is present, Tag 1195
func (m NoSettlOblig) HasOptPayAmount() bool {
	return m.Has(tag.OptPayAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m NoSettlOblig) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m NoSettlOblig) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m NoSettlOblig) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m NoSettlOblig) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoSettlOblig) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m NoSettlOblig) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m NoSettlOblig) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasFuturesValuationMethod returns true if FuturesValuationMethod is present, Tag 1197
func (m NoSettlOblig) HasFuturesValuationMethod() bool {
	return m.Has(tag.FuturesValuationMethod)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoSettlOblig) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m NoSettlOblig) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NoSettlOblig) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasLastUpdateTime returns true if LastUpdateTime is present, Tag 779
func (m NoSettlOblig) HasLastUpdateTime() bool {
	return m.Has(tag.LastUpdateTime)
}

//HasNoSettlDetails returns true if NoSettlDetails is present, Tag 1158
func (m NoSettlOblig) HasNoSettlDetails() bool {
	return m.Has(tag.NoSettlDetails)
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), quickfix.GroupElement(tag.NoInstrumentPartySubIDs)})}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.NoPartySubIDs)})}
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

//NoSettlDetails is a repeating group element, Tag 1158
type NoSettlDetails struct {
	quickfix.Group
}

//SetSettlObligSource sets SettlObligSource, Tag 1164
func (m NoSettlDetails) SetSettlObligSource(v string) {
	m.Set(field.NewSettlObligSource(v))
}

//SetNoSettlPartyIDs sets NoSettlPartyIDs, Tag 781
func (m NoSettlDetails) SetNoSettlPartyIDs(f NoSettlPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlObligSource gets SettlObligSource, Tag 1164
func (m NoSettlDetails) GetSettlObligSource() (f field.SettlObligSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartyIDs gets NoSettlPartyIDs, Tag 781
func (m NoSettlDetails) GetNoSettlPartyIDs() (NoSettlPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlObligSource returns true if SettlObligSource is present, Tag 1164
func (m NoSettlDetails) HasSettlObligSource() bool {
	return m.Has(tag.SettlObligSource)
}

//HasNoSettlPartyIDs returns true if NoSettlPartyIDs is present, Tag 781
func (m NoSettlDetails) HasNoSettlPartyIDs() bool {
	return m.Has(tag.NoSettlPartyIDs)
}

//NoSettlPartyIDs is a repeating group element, Tag 781
type NoSettlPartyIDs struct {
	quickfix.Group
}

//SetSettlPartyID sets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) SetSettlPartyID(v string) {
	m.Set(field.NewSettlPartyID(v))
}

//SetSettlPartyIDSource sets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) SetSettlPartyIDSource(v string) {
	m.Set(field.NewSettlPartyIDSource(v))
}

//SetSettlPartyRole sets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) SetSettlPartyRole(v int) {
	m.Set(field.NewSettlPartyRole(v))
}

//SetNoSettlPartySubIDs sets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) SetNoSettlPartySubIDs(f NoSettlPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetSettlPartyID gets SettlPartyID, Tag 782
func (m NoSettlPartyIDs) GetSettlPartyID() (f field.SettlPartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyIDSource gets SettlPartyIDSource, Tag 783
func (m NoSettlPartyIDs) GetSettlPartyIDSource() (f field.SettlPartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartyRole gets SettlPartyRole, Tag 784
func (m NoSettlPartyIDs) GetSettlPartyRole() (f field.SettlPartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSettlPartySubIDs gets NoSettlPartySubIDs, Tag 801
func (m NoSettlPartyIDs) GetNoSettlPartySubIDs() (NoSettlPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSettlPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasSettlPartyID returns true if SettlPartyID is present, Tag 782
func (m NoSettlPartyIDs) HasSettlPartyID() bool {
	return m.Has(tag.SettlPartyID)
}

//HasSettlPartyIDSource returns true if SettlPartyIDSource is present, Tag 783
func (m NoSettlPartyIDs) HasSettlPartyIDSource() bool {
	return m.Has(tag.SettlPartyIDSource)
}

//HasSettlPartyRole returns true if SettlPartyRole is present, Tag 784
func (m NoSettlPartyIDs) HasSettlPartyRole() bool {
	return m.Has(tag.SettlPartyRole)
}

//HasNoSettlPartySubIDs returns true if NoSettlPartySubIDs is present, Tag 801
func (m NoSettlPartyIDs) HasNoSettlPartySubIDs() bool {
	return m.Has(tag.NoSettlPartySubIDs)
}

//NoSettlPartySubIDs is a repeating group element, Tag 801
type NoSettlPartySubIDs struct {
	quickfix.Group
}

//SetSettlPartySubID sets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) SetSettlPartySubID(v string) {
	m.Set(field.NewSettlPartySubID(v))
}

//SetSettlPartySubIDType sets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) SetSettlPartySubIDType(v int) {
	m.Set(field.NewSettlPartySubIDType(v))
}

//GetSettlPartySubID gets SettlPartySubID, Tag 785
func (m NoSettlPartySubIDs) GetSettlPartySubID() (f field.SettlPartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlPartySubIDType gets SettlPartySubIDType, Tag 786
func (m NoSettlPartySubIDs) GetSettlPartySubIDType() (f field.SettlPartySubIDTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSettlPartySubID returns true if SettlPartySubID is present, Tag 785
func (m NoSettlPartySubIDs) HasSettlPartySubID() bool {
	return m.Has(tag.SettlPartySubID)
}

//HasSettlPartySubIDType returns true if SettlPartySubIDType is present, Tag 786
func (m NoSettlPartySubIDs) HasSettlPartySubIDType() bool {
	return m.Has(tag.SettlPartySubIDType)
}

//NoSettlPartySubIDsRepeatingGroup is a repeating group, Tag 801
type NoSettlPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartySubIDsRepeatingGroup returns an initialized, NoSettlPartySubIDsRepeatingGroup
func NewNoSettlPartySubIDsRepeatingGroup() NoSettlPartySubIDsRepeatingGroup {
	return NoSettlPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartySubID), quickfix.GroupElement(tag.SettlPartySubIDType)})}
}

//Add create and append a new NoSettlPartySubIDs to this group
func (m NoSettlPartySubIDsRepeatingGroup) Add() NoSettlPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartySubIDs{g}
}

//Get returns the ith NoSettlPartySubIDs in the NoSettlPartySubIDsRepeatinGroup
func (m NoSettlPartySubIDsRepeatingGroup) Get(i int) NoSettlPartySubIDs {
	return NoSettlPartySubIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlPartyIDsRepeatingGroup is a repeating group, Tag 781
type NoSettlPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlPartyIDsRepeatingGroup returns an initialized, NoSettlPartyIDsRepeatingGroup
func NewNoSettlPartyIDsRepeatingGroup() NoSettlPartyIDsRepeatingGroup {
	return NoSettlPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlPartyID), quickfix.GroupElement(tag.SettlPartyIDSource), quickfix.GroupElement(tag.SettlPartyRole), quickfix.GroupElement(tag.NoSettlPartySubIDs)})}
}

//Add create and append a new NoSettlPartyIDs to this group
func (m NoSettlPartyIDsRepeatingGroup) Add() NoSettlPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoSettlPartyIDs{g}
}

//Get returns the ith NoSettlPartyIDs in the NoSettlPartyIDsRepeatinGroup
func (m NoSettlPartyIDsRepeatingGroup) Get(i int) NoSettlPartyIDs {
	return NoSettlPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSettlDetailsRepeatingGroup is a repeating group, Tag 1158
type NoSettlDetailsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlDetailsRepeatingGroup returns an initialized, NoSettlDetailsRepeatingGroup
func NewNoSettlDetailsRepeatingGroup() NoSettlDetailsRepeatingGroup {
	return NoSettlDetailsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlDetails,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SettlObligSource), quickfix.GroupElement(tag.NoSettlPartyIDs)})}
}

//Add create and append a new NoSettlDetails to this group
func (m NoSettlDetailsRepeatingGroup) Add() NoSettlDetails {
	g := m.RepeatingGroup.Add()
	return NoSettlDetails{g}
}

//Get returns the ith NoSettlDetails in the NoSettlDetailsRepeatinGroup
func (m NoSettlDetailsRepeatingGroup) Get(i int) NoSettlDetails {
	return NoSettlDetails{m.RepeatingGroup.Get(i)}
}

//NoSettlObligRepeatingGroup is a repeating group, Tag 1165
type NoSettlObligRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSettlObligRepeatingGroup returns an initialized, NoSettlObligRepeatingGroup
func NewNoSettlObligRepeatingGroup() NoSettlObligRepeatingGroup {
	return NoSettlObligRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSettlOblig,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NetGrossInd), quickfix.GroupElement(tag.SettlObligID), quickfix.GroupElement(tag.SettlObligTransType), quickfix.GroupElement(tag.SettlObligRefID), quickfix.GroupElement(tag.CcyAmt), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlDate), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), quickfix.GroupElement(tag.NoSecurityAltID), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), quickfix.GroupElement(tag.NoEvents), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.SecurityStatus), quickfix.GroupElement(tag.SettleOnOpenFlag), quickfix.GroupElement(tag.InstrmtAssignmentMethod), quickfix.GroupElement(tag.StrikeMultiplier), quickfix.GroupElement(tag.StrikeValue), quickfix.GroupElement(tag.MinPriceIncrement), quickfix.GroupElement(tag.PositionLimit), quickfix.GroupElement(tag.NTPositionLimit), quickfix.GroupElement(tag.NoInstrumentParties), quickfix.GroupElement(tag.UnitOfMeasure), quickfix.GroupElement(tag.TimeUnit), quickfix.GroupElement(tag.MaturityTime), quickfix.GroupElement(tag.SecurityGroup), quickfix.GroupElement(tag.MinPriceIncrementAmount), quickfix.GroupElement(tag.UnitOfMeasureQty), quickfix.GroupElement(tag.SecurityXMLLen), quickfix.GroupElement(tag.SecurityXML), quickfix.GroupElement(tag.SecurityXMLSchema), quickfix.GroupElement(tag.ProductComplex), quickfix.GroupElement(tag.PriceUnitOfMeasure), quickfix.GroupElement(tag.PriceUnitOfMeasureQty), quickfix.GroupElement(tag.SettlMethod), quickfix.GroupElement(tag.ExerciseStyle), quickfix.GroupElement(tag.OptPayAmount), quickfix.GroupElement(tag.PriceQuoteMethod), quickfix.GroupElement(tag.ListMethod), quickfix.GroupElement(tag.CapPrice), quickfix.GroupElement(tag.FloorPrice), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.FlexibleIndicator), quickfix.GroupElement(tag.FlexProductEligibilityIndicator), quickfix.GroupElement(tag.FuturesValuationMethod), quickfix.GroupElement(tag.NoPartyIDs), quickfix.GroupElement(tag.EffectiveTime), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.LastUpdateTime), quickfix.GroupElement(tag.NoSettlDetails)})}
}

//Add create and append a new NoSettlOblig to this group
func (m NoSettlObligRepeatingGroup) Add() NoSettlOblig {
	g := m.RepeatingGroup.Add()
	return NoSettlOblig{g}
}

//Get returns the ith NoSettlOblig in the NoSettlObligRepeatinGroup
func (m NoSettlObligRepeatingGroup) Get(i int) NoSettlOblig {
	return NoSettlOblig{m.RepeatingGroup.Get(i)}
}
