package quotecancel

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fixt11"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// QuoteCancel is the fix50sp2 QuoteCancel type, MsgType = Z.
type QuoteCancel struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

// FromMessage creates a QuoteCancel from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) QuoteCancel {
	return QuoteCancel{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m QuoteCancel) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a QuoteCancel initialized with the required fields for QuoteCancel.
func New(quotecanceltype field.QuoteCancelTypeField) (m QuoteCancel) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("Z"))
	m.Set(quotecanceltype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg QuoteCancel, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "Z", r
}

// SetAccount sets Account, Tag 1.
func (m QuoteCancel) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetQuoteID sets QuoteID, Tag 117.
func (m QuoteCancel) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

// SetQuoteReqID sets QuoteReqID, Tag 131.
func (m QuoteCancel) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

// SetNoQuoteEntries sets NoQuoteEntries, Tag 295.
func (m QuoteCancel) SetNoQuoteEntries(f NoQuoteEntriesRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteCancelType sets QuoteCancelType, Tag 298.
func (m QuoteCancel) SetQuoteCancelType(v enum.QuoteCancelType) {
	m.Set(field.NewQuoteCancelType(v))
}

// SetQuoteResponseLevel sets QuoteResponseLevel, Tag 301.
func (m QuoteCancel) SetQuoteResponseLevel(v enum.QuoteResponseLevel) {
	m.Set(field.NewQuoteResponseLevel(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m QuoteCancel) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m QuoteCancel) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetQuoteType sets QuoteType, Tag 537.
func (m QuoteCancel) SetQuoteType(v enum.QuoteType) {
	m.Set(field.NewQuoteType(v))
}

// SetAccountType sets AccountType, Tag 581.
func (m QuoteCancel) SetAccountType(v enum.AccountType) {
	m.Set(field.NewAccountType(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m QuoteCancel) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// SetAcctIDSource sets AcctIDSource, Tag 660.
func (m QuoteCancel) SetAcctIDSource(v enum.AcctIDSource) {
	m.Set(field.NewAcctIDSource(v))
}

// SetQuoteMsgID sets QuoteMsgID, Tag 1166.
func (m QuoteCancel) SetQuoteMsgID(v string) {
	m.Set(field.NewQuoteMsgID(v))
}

// SetNoTargetPartyIDs sets NoTargetPartyIDs, Tag 1461.
func (m QuoteCancel) SetNoTargetPartyIDs(f NoTargetPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetAccount gets Account, Tag 1.
func (m QuoteCancel) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteID gets QuoteID, Tag 117.
func (m QuoteCancel) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteReqID gets QuoteReqID, Tag 131.
func (m QuoteCancel) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoQuoteEntries gets NoQuoteEntries, Tag 295.
func (m QuoteCancel) GetNoQuoteEntries() (NoQuoteEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteCancelType gets QuoteCancelType, Tag 298.
func (m QuoteCancel) GetQuoteCancelType() (v enum.QuoteCancelType, err quickfix.MessageRejectError) {
	var f field.QuoteCancelTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteResponseLevel gets QuoteResponseLevel, Tag 301.
func (m QuoteCancel) GetQuoteResponseLevel() (v enum.QuoteResponseLevel, err quickfix.MessageRejectError) {
	var f field.QuoteResponseLevelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m QuoteCancel) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m QuoteCancel) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetQuoteType gets QuoteType, Tag 537.
func (m QuoteCancel) GetQuoteType() (v enum.QuoteType, err quickfix.MessageRejectError) {
	var f field.QuoteTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAccountType gets AccountType, Tag 581.
func (m QuoteCancel) GetAccountType() (v enum.AccountType, err quickfix.MessageRejectError) {
	var f field.AccountTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m QuoteCancel) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAcctIDSource gets AcctIDSource, Tag 660.
func (m QuoteCancel) GetAcctIDSource() (v enum.AcctIDSource, err quickfix.MessageRejectError) {
	var f field.AcctIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetQuoteMsgID gets QuoteMsgID, Tag 1166.
func (m QuoteCancel) GetQuoteMsgID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteMsgIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoTargetPartyIDs gets NoTargetPartyIDs, Tag 1461.
func (m QuoteCancel) GetNoTargetPartyIDs() (NoTargetPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTargetPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasAccount returns true if Account is present, Tag 1.
func (m QuoteCancel) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasQuoteID returns true if QuoteID is present, Tag 117.
func (m QuoteCancel) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

// HasQuoteReqID returns true if QuoteReqID is present, Tag 131.
func (m QuoteCancel) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

// HasNoQuoteEntries returns true if NoQuoteEntries is present, Tag 295.
func (m QuoteCancel) HasNoQuoteEntries() bool {
	return m.Has(tag.NoQuoteEntries)
}

// HasQuoteCancelType returns true if QuoteCancelType is present, Tag 298.
func (m QuoteCancel) HasQuoteCancelType() bool {
	return m.Has(tag.QuoteCancelType)
}

// HasQuoteResponseLevel returns true if QuoteResponseLevel is present, Tag 301.
func (m QuoteCancel) HasQuoteResponseLevel() bool {
	return m.Has(tag.QuoteResponseLevel)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m QuoteCancel) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m QuoteCancel) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasQuoteType returns true if QuoteType is present, Tag 537.
func (m QuoteCancel) HasQuoteType() bool {
	return m.Has(tag.QuoteType)
}

// HasAccountType returns true if AccountType is present, Tag 581.
func (m QuoteCancel) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m QuoteCancel) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// HasAcctIDSource returns true if AcctIDSource is present, Tag 660.
func (m QuoteCancel) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

// HasQuoteMsgID returns true if QuoteMsgID is present, Tag 1166.
func (m QuoteCancel) HasQuoteMsgID() bool {
	return m.Has(tag.QuoteMsgID)
}

// HasNoTargetPartyIDs returns true if NoTargetPartyIDs is present, Tag 1461.
func (m QuoteCancel) HasNoTargetPartyIDs() bool {
	return m.Has(tag.NoTargetPartyIDs)
}

// NoQuoteEntries is a repeating group element, Tag 295.
type NoQuoteEntries struct {
	*quickfix.Group
}

// SetSymbol sets Symbol, Tag 55.
func (m NoQuoteEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m NoQuoteEntries) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoQuoteEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m NoQuoteEntries) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m NoQuoteEntries) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m NoQuoteEntries) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m NoQuoteEntries) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m NoQuoteEntries) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetSecuritySubType sets SecuritySubType, Tag 762.
func (m NoQuoteEntries) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m NoQuoteEntries) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m NoQuoteEntries) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m NoQuoteEntries) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m NoQuoteEntries) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m NoQuoteEntries) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m NoQuoteEntries) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m NoQuoteEntries) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m NoQuoteEntries) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m NoQuoteEntries) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m NoQuoteEntries) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m NoQuoteEntries) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m NoQuoteEntries) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m NoQuoteEntries) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m NoQuoteEntries) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m NoQuoteEntries) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetStrikeCurrency sets StrikeCurrency, Tag 947.
func (m NoQuoteEntries) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m NoQuoteEntries) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m NoQuoteEntries) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m NoQuoteEntries) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m NoQuoteEntries) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m NoQuoteEntries) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m NoQuoteEntries) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m NoQuoteEntries) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m NoQuoteEntries) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m NoQuoteEntries) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m NoQuoteEntries) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetPool sets Pool, Tag 691.
func (m NoQuoteEntries) SetPool(v string) {
	m.Set(field.NewPool(v))
}

// SetContractSettlMonth sets ContractSettlMonth, Tag 667.
func (m NoQuoteEntries) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

// SetCPProgram sets CPProgram, Tag 875.
func (m NoQuoteEntries) SetCPProgram(v enum.CPProgram) {
	m.Set(field.NewCPProgram(v))
}

// SetCPRegType sets CPRegType, Tag 876.
func (m NoQuoteEntries) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

// SetNoEvents sets NoEvents, Tag 864.
func (m NoQuoteEntries) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetDatedDate sets DatedDate, Tag 873.
func (m NoQuoteEntries) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

// SetInterestAccrualDate sets InterestAccrualDate, Tag 874.
func (m NoQuoteEntries) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

// SetSecurityStatus sets SecurityStatus, Tag 965.
func (m NoQuoteEntries) SetSecurityStatus(v enum.SecurityStatus) {
	m.Set(field.NewSecurityStatus(v))
}

// SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966.
func (m NoQuoteEntries) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

// SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049.
func (m NoQuoteEntries) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

// SetStrikeMultiplier sets StrikeMultiplier, Tag 967.
func (m NoQuoteEntries) SetStrikeMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeMultiplier(value, scale))
}

// SetStrikeValue sets StrikeValue, Tag 968.
func (m NoQuoteEntries) SetStrikeValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikeValue(value, scale))
}

// SetMinPriceIncrement sets MinPriceIncrement, Tag 969.
func (m NoQuoteEntries) SetMinPriceIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrement(value, scale))
}

// SetPositionLimit sets PositionLimit, Tag 970.
func (m NoQuoteEntries) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

// SetNTPositionLimit sets NTPositionLimit, Tag 971.
func (m NoQuoteEntries) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

// SetNoInstrumentParties sets NoInstrumentParties, Tag 1018.
func (m NoQuoteEntries) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnitOfMeasure sets UnitOfMeasure, Tag 996.
func (m NoQuoteEntries) SetUnitOfMeasure(v enum.UnitOfMeasure) {
	m.Set(field.NewUnitOfMeasure(v))
}

// SetTimeUnit sets TimeUnit, Tag 997.
func (m NoQuoteEntries) SetTimeUnit(v enum.TimeUnit) {
	m.Set(field.NewTimeUnit(v))
}

// SetMaturityTime sets MaturityTime, Tag 1079.
func (m NoQuoteEntries) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

// SetSecurityGroup sets SecurityGroup, Tag 1151.
func (m NoQuoteEntries) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

// SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146.
func (m NoQuoteEntries) SetMinPriceIncrementAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinPriceIncrementAmount(value, scale))
}

// SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147.
func (m NoQuoteEntries) SetUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnitOfMeasureQty(value, scale))
}

// SetSecurityXMLLen sets SecurityXMLLen, Tag 1184.
func (m NoQuoteEntries) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

// SetSecurityXML sets SecurityXML, Tag 1185.
func (m NoQuoteEntries) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

// SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186.
func (m NoQuoteEntries) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

// SetProductComplex sets ProductComplex, Tag 1227.
func (m NoQuoteEntries) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

// SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191.
func (m NoQuoteEntries) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

// SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192.
func (m NoQuoteEntries) SetPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewPriceUnitOfMeasureQty(value, scale))
}

// SetSettlMethod sets SettlMethod, Tag 1193.
func (m NoQuoteEntries) SetSettlMethod(v enum.SettlMethod) {
	m.Set(field.NewSettlMethod(v))
}

// SetExerciseStyle sets ExerciseStyle, Tag 1194.
func (m NoQuoteEntries) SetExerciseStyle(v enum.ExerciseStyle) {
	m.Set(field.NewExerciseStyle(v))
}

// SetOptPayoutAmount sets OptPayoutAmount, Tag 1195.
func (m NoQuoteEntries) SetOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewOptPayoutAmount(value, scale))
}

// SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196.
func (m NoQuoteEntries) SetPriceQuoteMethod(v enum.PriceQuoteMethod) {
	m.Set(field.NewPriceQuoteMethod(v))
}

// SetListMethod sets ListMethod, Tag 1198.
func (m NoQuoteEntries) SetListMethod(v enum.ListMethod) {
	m.Set(field.NewListMethod(v))
}

// SetCapPrice sets CapPrice, Tag 1199.
func (m NoQuoteEntries) SetCapPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewCapPrice(value, scale))
}

// SetFloorPrice sets FloorPrice, Tag 1200.
func (m NoQuoteEntries) SetFloorPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewFloorPrice(value, scale))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m NoQuoteEntries) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetFlexibleIndicator sets FlexibleIndicator, Tag 1244.
func (m NoQuoteEntries) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

// SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242.
func (m NoQuoteEntries) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

// SetValuationMethod sets ValuationMethod, Tag 1197.
func (m NoQuoteEntries) SetValuationMethod(v enum.ValuationMethod) {
	m.Set(field.NewValuationMethod(v))
}

// SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435.
func (m NoQuoteEntries) SetContractMultiplierUnit(v enum.ContractMultiplierUnit) {
	m.Set(field.NewContractMultiplierUnit(v))
}

// SetFlowScheduleType sets FlowScheduleType, Tag 1439.
func (m NoQuoteEntries) SetFlowScheduleType(v enum.FlowScheduleType) {
	m.Set(field.NewFlowScheduleType(v))
}

// SetRestructuringType sets RestructuringType, Tag 1449.
func (m NoQuoteEntries) SetRestructuringType(v enum.RestructuringType) {
	m.Set(field.NewRestructuringType(v))
}

// SetSeniority sets Seniority, Tag 1450.
func (m NoQuoteEntries) SetSeniority(v enum.Seniority) {
	m.Set(field.NewSeniority(v))
}

// SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451.
func (m NoQuoteEntries) SetNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewNotionalPercentageOutstanding(value, scale))
}

// SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m NoQuoteEntries) SetOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(value, scale))
}

// SetAttachmentPoint sets AttachmentPoint, Tag 1457.
func (m NoQuoteEntries) SetAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewAttachmentPoint(value, scale))
}

// SetDetachmentPoint sets DetachmentPoint, Tag 1458.
func (m NoQuoteEntries) SetDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewDetachmentPoint(value, scale))
}

// SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478.
func (m NoQuoteEntries) SetStrikePriceDeterminationMethod(v enum.StrikePriceDeterminationMethod) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

// SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479.
func (m NoQuoteEntries) SetStrikePriceBoundaryMethod(v enum.StrikePriceBoundaryMethod) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

// SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480.
func (m NoQuoteEntries) SetStrikePriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePriceBoundaryPrecision(value, scale))
}

// SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m NoQuoteEntries) SetUnderlyingPriceDeterminationMethod(v enum.UnderlyingPriceDeterminationMethod) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

// SetOptPayoutType sets OptPayoutType, Tag 1482.
func (m NoQuoteEntries) SetOptPayoutType(v enum.OptPayoutType) {
	m.Set(field.NewOptPayoutType(v))
}

// SetNoComplexEvents sets NoComplexEvents, Tag 1483.
func (m NoQuoteEntries) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

// SetAgreementDesc sets AgreementDesc, Tag 913.
func (m NoQuoteEntries) SetAgreementDesc(v string) {
	m.Set(field.NewAgreementDesc(v))
}

// SetAgreementID sets AgreementID, Tag 914.
func (m NoQuoteEntries) SetAgreementID(v string) {
	m.Set(field.NewAgreementID(v))
}

// SetAgreementDate sets AgreementDate, Tag 915.
func (m NoQuoteEntries) SetAgreementDate(v string) {
	m.Set(field.NewAgreementDate(v))
}

// SetAgreementCurrency sets AgreementCurrency, Tag 918.
func (m NoQuoteEntries) SetAgreementCurrency(v string) {
	m.Set(field.NewAgreementCurrency(v))
}

// SetTerminationType sets TerminationType, Tag 788.
func (m NoQuoteEntries) SetTerminationType(v enum.TerminationType) {
	m.Set(field.NewTerminationType(v))
}

// SetStartDate sets StartDate, Tag 916.
func (m NoQuoteEntries) SetStartDate(v string) {
	m.Set(field.NewStartDate(v))
}

// SetEndDate sets EndDate, Tag 917.
func (m NoQuoteEntries) SetEndDate(v string) {
	m.Set(field.NewEndDate(v))
}

// SetDeliveryType sets DeliveryType, Tag 919.
func (m NoQuoteEntries) SetDeliveryType(v enum.DeliveryType) {
	m.Set(field.NewDeliveryType(v))
}

// SetMarginRatio sets MarginRatio, Tag 898.
func (m NoQuoteEntries) SetMarginRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewMarginRatio(value, scale))
}

// SetNoUnderlyings sets NoUnderlyings, Tag 711.
func (m NoQuoteEntries) SetNoUnderlyings(f NoUnderlyingsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoLegs sets NoLegs, Tag 555.
func (m NoQuoteEntries) SetNoLegs(f NoLegsRepeatingGroup) {
	m.SetGroup(f)
}

// GetSymbol gets Symbol, Tag 55.
func (m NoQuoteEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m NoQuoteEntries) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoQuoteEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m NoQuoteEntries) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m NoQuoteEntries) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m NoQuoteEntries) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m NoQuoteEntries) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m NoQuoteEntries) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecuritySubType gets SecuritySubType, Tag 762.
func (m NoQuoteEntries) GetSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.SecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m NoQuoteEntries) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m NoQuoteEntries) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m NoQuoteEntries) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m NoQuoteEntries) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m NoQuoteEntries) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m NoQuoteEntries) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m NoQuoteEntries) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m NoQuoteEntries) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m NoQuoteEntries) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m NoQuoteEntries) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m NoQuoteEntries) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m NoQuoteEntries) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m NoQuoteEntries) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m NoQuoteEntries) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m NoQuoteEntries) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeCurrency gets StrikeCurrency, Tag 947.
func (m NoQuoteEntries) GetStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.StrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m NoQuoteEntries) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m NoQuoteEntries) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m NoQuoteEntries) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m NoQuoteEntries) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m NoQuoteEntries) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m NoQuoteEntries) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m NoQuoteEntries) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m NoQuoteEntries) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m NoQuoteEntries) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m NoQuoteEntries) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPool gets Pool, Tag 691.
func (m NoQuoteEntries) GetPool() (v string, err quickfix.MessageRejectError) {
	var f field.PoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractSettlMonth gets ContractSettlMonth, Tag 667.
func (m NoQuoteEntries) GetContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.ContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPProgram gets CPProgram, Tag 875.
func (m NoQuoteEntries) GetCPProgram() (v enum.CPProgram, err quickfix.MessageRejectError) {
	var f field.CPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCPRegType gets CPRegType, Tag 876.
func (m NoQuoteEntries) GetCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.CPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoEvents gets NoEvents, Tag 864.
func (m NoQuoteEntries) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetDatedDate gets DatedDate, Tag 873.
func (m NoQuoteEntries) GetDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.DatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInterestAccrualDate gets InterestAccrualDate, Tag 874.
func (m NoQuoteEntries) GetInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.InterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityStatus gets SecurityStatus, Tag 965.
func (m NoQuoteEntries) GetSecurityStatus() (v enum.SecurityStatus, err quickfix.MessageRejectError) {
	var f field.SecurityStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966.
func (m NoQuoteEntries) GetSettleOnOpenFlag() (v string, err quickfix.MessageRejectError) {
	var f field.SettleOnOpenFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049.
func (m NoQuoteEntries) GetInstrmtAssignmentMethod() (v string, err quickfix.MessageRejectError) {
	var f field.InstrmtAssignmentMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeMultiplier gets StrikeMultiplier, Tag 967.
func (m NoQuoteEntries) GetStrikeMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikeValue gets StrikeValue, Tag 968.
func (m NoQuoteEntries) GetStrikeValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikeValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrement gets MinPriceIncrement, Tag 969.
func (m NoQuoteEntries) GetMinPriceIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPositionLimit gets PositionLimit, Tag 970.
func (m NoQuoteEntries) GetPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.PositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNTPositionLimit gets NTPositionLimit, Tag 971.
func (m NoQuoteEntries) GetNTPositionLimit() (v int, err quickfix.MessageRejectError) {
	var f field.NTPositionLimitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentParties gets NoInstrumentParties, Tag 1018.
func (m NoQuoteEntries) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnitOfMeasure gets UnitOfMeasure, Tag 996.
func (m NoQuoteEntries) GetUnitOfMeasure() (v enum.UnitOfMeasure, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTimeUnit gets TimeUnit, Tag 997.
func (m NoQuoteEntries) GetTimeUnit() (v enum.TimeUnit, err quickfix.MessageRejectError) {
	var f field.TimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityTime gets MaturityTime, Tag 1079.
func (m NoQuoteEntries) GetMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityGroup gets SecurityGroup, Tag 1151.
func (m NoQuoteEntries) GetSecurityGroup() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityGroupField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146.
func (m NoQuoteEntries) GetMinPriceIncrementAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinPriceIncrementAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147.
func (m NoQuoteEntries) GetUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLLen gets SecurityXMLLen, Tag 1184.
func (m NoQuoteEntries) GetSecurityXMLLen() (v int, err quickfix.MessageRejectError) {
	var f field.SecurityXMLLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXML gets SecurityXML, Tag 1185.
func (m NoQuoteEntries) GetSecurityXML() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186.
func (m NoQuoteEntries) GetSecurityXMLSchema() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityXMLSchemaField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetProductComplex gets ProductComplex, Tag 1227.
func (m NoQuoteEntries) GetProductComplex() (v string, err quickfix.MessageRejectError) {
	var f field.ProductComplexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191.
func (m NoQuoteEntries) GetPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192.
func (m NoQuoteEntries) GetPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSettlMethod gets SettlMethod, Tag 1193.
func (m NoQuoteEntries) GetSettlMethod() (v enum.SettlMethod, err quickfix.MessageRejectError) {
	var f field.SettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetExerciseStyle gets ExerciseStyle, Tag 1194.
func (m NoQuoteEntries) GetExerciseStyle() (v enum.ExerciseStyle, err quickfix.MessageRejectError) {
	var f field.ExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutAmount gets OptPayoutAmount, Tag 1195.
func (m NoQuoteEntries) GetOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196.
func (m NoQuoteEntries) GetPriceQuoteMethod() (v enum.PriceQuoteMethod, err quickfix.MessageRejectError) {
	var f field.PriceQuoteMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetListMethod gets ListMethod, Tag 1198.
func (m NoQuoteEntries) GetListMethod() (v enum.ListMethod, err quickfix.MessageRejectError) {
	var f field.ListMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCapPrice gets CapPrice, Tag 1199.
func (m NoQuoteEntries) GetCapPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CapPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFloorPrice gets FloorPrice, Tag 1200.
func (m NoQuoteEntries) GetFloorPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FloorPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m NoQuoteEntries) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexibleIndicator gets FlexibleIndicator, Tag 1244.
func (m NoQuoteEntries) GetFlexibleIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexibleIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242.
func (m NoQuoteEntries) GetFlexProductEligibilityIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.FlexProductEligibilityIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetValuationMethod gets ValuationMethod, Tag 1197.
func (m NoQuoteEntries) GetValuationMethod() (v enum.ValuationMethod, err quickfix.MessageRejectError) {
	var f field.ValuationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435.
func (m NoQuoteEntries) GetContractMultiplierUnit() (v enum.ContractMultiplierUnit, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFlowScheduleType gets FlowScheduleType, Tag 1439.
func (m NoQuoteEntries) GetFlowScheduleType() (v enum.FlowScheduleType, err quickfix.MessageRejectError) {
	var f field.FlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRestructuringType gets RestructuringType, Tag 1449.
func (m NoQuoteEntries) GetRestructuringType() (v enum.RestructuringType, err quickfix.MessageRejectError) {
	var f field.RestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSeniority gets Seniority, Tag 1450.
func (m NoQuoteEntries) GetSeniority() (v enum.Seniority, err quickfix.MessageRejectError) {
	var f field.SeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451.
func (m NoQuoteEntries) GetNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452.
func (m NoQuoteEntries) GetOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAttachmentPoint gets AttachmentPoint, Tag 1457.
func (m NoQuoteEntries) GetAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDetachmentPoint gets DetachmentPoint, Tag 1458.
func (m NoQuoteEntries) GetDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478.
func (m NoQuoteEntries) GetStrikePriceDeterminationMethod() (v enum.StrikePriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479.
func (m NoQuoteEntries) GetStrikePriceBoundaryMethod() (v enum.StrikePriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480.
func (m NoQuoteEntries) GetStrikePriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481.
func (m NoQuoteEntries) GetUnderlyingPriceDeterminationMethod() (v enum.UnderlyingPriceDeterminationMethod, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceDeterminationMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptPayoutType gets OptPayoutType, Tag 1482.
func (m NoQuoteEntries) GetOptPayoutType() (v enum.OptPayoutType, err quickfix.MessageRejectError) {
	var f field.OptPayoutTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEvents gets NoComplexEvents, Tag 1483.
func (m NoQuoteEntries) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetAgreementDesc gets AgreementDesc, Tag 913.
func (m NoQuoteEntries) GetAgreementDesc() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementID gets AgreementID, Tag 914.
func (m NoQuoteEntries) GetAgreementID() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementDate gets AgreementDate, Tag 915.
func (m NoQuoteEntries) GetAgreementDate() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetAgreementCurrency gets AgreementCurrency, Tag 918.
func (m NoQuoteEntries) GetAgreementCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.AgreementCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTerminationType gets TerminationType, Tag 788.
func (m NoQuoteEntries) GetTerminationType() (v enum.TerminationType, err quickfix.MessageRejectError) {
	var f field.TerminationTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStartDate gets StartDate, Tag 916.
func (m NoQuoteEntries) GetStartDate() (v string, err quickfix.MessageRejectError) {
	var f field.StartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEndDate gets EndDate, Tag 917.
func (m NoQuoteEntries) GetEndDate() (v string, err quickfix.MessageRejectError) {
	var f field.EndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetDeliveryType gets DeliveryType, Tag 919.
func (m NoQuoteEntries) GetDeliveryType() (v enum.DeliveryType, err quickfix.MessageRejectError) {
	var f field.DeliveryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMarginRatio gets MarginRatio, Tag 898.
func (m NoQuoteEntries) GetMarginRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MarginRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyings gets NoUnderlyings, Tag 711.
func (m NoQuoteEntries) GetNoUnderlyings() (NoUnderlyingsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoLegs gets NoLegs, Tag 555.
func (m NoQuoteEntries) GetNoLegs() (NoLegsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m NoQuoteEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m NoQuoteEntries) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoQuoteEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m NoQuoteEntries) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m NoQuoteEntries) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m NoQuoteEntries) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m NoQuoteEntries) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m NoQuoteEntries) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasSecuritySubType returns true if SecuritySubType is present, Tag 762.
func (m NoQuoteEntries) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m NoQuoteEntries) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m NoQuoteEntries) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m NoQuoteEntries) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m NoQuoteEntries) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m NoQuoteEntries) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m NoQuoteEntries) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m NoQuoteEntries) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m NoQuoteEntries) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m NoQuoteEntries) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m NoQuoteEntries) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m NoQuoteEntries) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m NoQuoteEntries) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m NoQuoteEntries) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m NoQuoteEntries) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m NoQuoteEntries) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947.
func (m NoQuoteEntries) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m NoQuoteEntries) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m NoQuoteEntries) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m NoQuoteEntries) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m NoQuoteEntries) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m NoQuoteEntries) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m NoQuoteEntries) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m NoQuoteEntries) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m NoQuoteEntries) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m NoQuoteEntries) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m NoQuoteEntries) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasPool returns true if Pool is present, Tag 691.
func (m NoQuoteEntries) HasPool() bool {
	return m.Has(tag.Pool)
}

// HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667.
func (m NoQuoteEntries) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

// HasCPProgram returns true if CPProgram is present, Tag 875.
func (m NoQuoteEntries) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

// HasCPRegType returns true if CPRegType is present, Tag 876.
func (m NoQuoteEntries) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

// HasNoEvents returns true if NoEvents is present, Tag 864.
func (m NoQuoteEntries) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

// HasDatedDate returns true if DatedDate is present, Tag 873.
func (m NoQuoteEntries) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

// HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874.
func (m NoQuoteEntries) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

// HasSecurityStatus returns true if SecurityStatus is present, Tag 965.
func (m NoQuoteEntries) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

// HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966.
func (m NoQuoteEntries) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

// HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049.
func (m NoQuoteEntries) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

// HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967.
func (m NoQuoteEntries) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

// HasStrikeValue returns true if StrikeValue is present, Tag 968.
func (m NoQuoteEntries) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

// HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969.
func (m NoQuoteEntries) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

// HasPositionLimit returns true if PositionLimit is present, Tag 970.
func (m NoQuoteEntries) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

// HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971.
func (m NoQuoteEntries) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

// HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018.
func (m NoQuoteEntries) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

// HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996.
func (m NoQuoteEntries) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

// HasTimeUnit returns true if TimeUnit is present, Tag 997.
func (m NoQuoteEntries) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

// HasMaturityTime returns true if MaturityTime is present, Tag 1079.
func (m NoQuoteEntries) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

// HasSecurityGroup returns true if SecurityGroup is present, Tag 1151.
func (m NoQuoteEntries) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

// HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146.
func (m NoQuoteEntries) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

// HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147.
func (m NoQuoteEntries) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

// HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184.
func (m NoQuoteEntries) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

// HasSecurityXML returns true if SecurityXML is present, Tag 1185.
func (m NoQuoteEntries) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

// HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186.
func (m NoQuoteEntries) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

// HasProductComplex returns true if ProductComplex is present, Tag 1227.
func (m NoQuoteEntries) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

// HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191.
func (m NoQuoteEntries) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

// HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192.
func (m NoQuoteEntries) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

// HasSettlMethod returns true if SettlMethod is present, Tag 1193.
func (m NoQuoteEntries) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

// HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194.
func (m NoQuoteEntries) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

// HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195.
func (m NoQuoteEntries) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

// HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196.
func (m NoQuoteEntries) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

// HasListMethod returns true if ListMethod is present, Tag 1198.
func (m NoQuoteEntries) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

// HasCapPrice returns true if CapPrice is present, Tag 1199.
func (m NoQuoteEntries) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

// HasFloorPrice returns true if FloorPrice is present, Tag 1200.
func (m NoQuoteEntries) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m NoQuoteEntries) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244.
func (m NoQuoteEntries) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

// HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242.
func (m NoQuoteEntries) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

// HasValuationMethod returns true if ValuationMethod is present, Tag 1197.
func (m NoQuoteEntries) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

// HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435.
func (m NoQuoteEntries) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

// HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439.
func (m NoQuoteEntries) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

// HasRestructuringType returns true if RestructuringType is present, Tag 1449.
func (m NoQuoteEntries) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

// HasSeniority returns true if Seniority is present, Tag 1450.
func (m NoQuoteEntries) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

// HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451.
func (m NoQuoteEntries) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

// HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452.
func (m NoQuoteEntries) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

// HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457.
func (m NoQuoteEntries) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

// HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458.
func (m NoQuoteEntries) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

// HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478.
func (m NoQuoteEntries) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

// HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479.
func (m NoQuoteEntries) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

// HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480.
func (m NoQuoteEntries) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

// HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481.
func (m NoQuoteEntries) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

// HasOptPayoutType returns true if OptPayoutType is present, Tag 1482.
func (m NoQuoteEntries) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

// HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483.
func (m NoQuoteEntries) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
}

// HasAgreementDesc returns true if AgreementDesc is present, Tag 913.
func (m NoQuoteEntries) HasAgreementDesc() bool {
	return m.Has(tag.AgreementDesc)
}

// HasAgreementID returns true if AgreementID is present, Tag 914.
func (m NoQuoteEntries) HasAgreementID() bool {
	return m.Has(tag.AgreementID)
}

// HasAgreementDate returns true if AgreementDate is present, Tag 915.
func (m NoQuoteEntries) HasAgreementDate() bool {
	return m.Has(tag.AgreementDate)
}

// HasAgreementCurrency returns true if AgreementCurrency is present, Tag 918.
func (m NoQuoteEntries) HasAgreementCurrency() bool {
	return m.Has(tag.AgreementCurrency)
}

// HasTerminationType returns true if TerminationType is present, Tag 788.
func (m NoQuoteEntries) HasTerminationType() bool {
	return m.Has(tag.TerminationType)
}

// HasStartDate returns true if StartDate is present, Tag 916.
func (m NoQuoteEntries) HasStartDate() bool {
	return m.Has(tag.StartDate)
}

// HasEndDate returns true if EndDate is present, Tag 917.
func (m NoQuoteEntries) HasEndDate() bool {
	return m.Has(tag.EndDate)
}

// HasDeliveryType returns true if DeliveryType is present, Tag 919.
func (m NoQuoteEntries) HasDeliveryType() bool {
	return m.Has(tag.DeliveryType)
}

// HasMarginRatio returns true if MarginRatio is present, Tag 898.
func (m NoQuoteEntries) HasMarginRatio() bool {
	return m.Has(tag.MarginRatio)
}

// HasNoUnderlyings returns true if NoUnderlyings is present, Tag 711.
func (m NoQuoteEntries) HasNoUnderlyings() bool {
	return m.Has(tag.NoUnderlyings)
}

// HasNoLegs returns true if NoLegs is present, Tag 555.
func (m NoQuoteEntries) HasNoLegs() bool {
	return m.Has(tag.NoLegs)
}

// NoSecurityAltID is a repeating group element, Tag 454.
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455.
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455.
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455.
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456.
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454.
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup.
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

// Add create and append a new NoSecurityAltID to this group.
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup.
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoEvents is a repeating group element, Tag 864.
type NoEvents struct {
	*quickfix.Group
}

// SetEventType sets EventType, Tag 865.
func (m NoEvents) SetEventType(v enum.EventType) {
	m.Set(field.NewEventType(v))
}

// SetEventDate sets EventDate, Tag 866.
func (m NoEvents) SetEventDate(v string) {
	m.Set(field.NewEventDate(v))
}

// SetEventPx sets EventPx, Tag 867.
func (m NoEvents) SetEventPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewEventPx(value, scale))
}

// SetEventText sets EventText, Tag 868.
func (m NoEvents) SetEventText(v string) {
	m.Set(field.NewEventText(v))
}

// SetEventTime sets EventTime, Tag 1145.
func (m NoEvents) SetEventTime(v time.Time) {
	m.Set(field.NewEventTime(v))
}

// GetEventType gets EventType, Tag 865.
func (m NoEvents) GetEventType() (v enum.EventType, err quickfix.MessageRejectError) {
	var f field.EventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventDate gets EventDate, Tag 866.
func (m NoEvents) GetEventDate() (v string, err quickfix.MessageRejectError) {
	var f field.EventDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventPx gets EventPx, Tag 867.
func (m NoEvents) GetEventPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EventPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventText gets EventText, Tag 868.
func (m NoEvents) GetEventText() (v string, err quickfix.MessageRejectError) {
	var f field.EventTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEventTime gets EventTime, Tag 1145.
func (m NoEvents) GetEventTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.EventTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasEventType returns true if EventType is present, Tag 865.
func (m NoEvents) HasEventType() bool {
	return m.Has(tag.EventType)
}

// HasEventDate returns true if EventDate is present, Tag 866.
func (m NoEvents) HasEventDate() bool {
	return m.Has(tag.EventDate)
}

// HasEventPx returns true if EventPx is present, Tag 867.
func (m NoEvents) HasEventPx() bool {
	return m.Has(tag.EventPx)
}

// HasEventText returns true if EventText is present, Tag 868.
func (m NoEvents) HasEventText() bool {
	return m.Has(tag.EventText)
}

// HasEventTime returns true if EventTime is present, Tag 1145.
func (m NoEvents) HasEventTime() bool {
	return m.Has(tag.EventTime)
}

// NoEventsRepeatingGroup is a repeating group, Tag 864.
type NoEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoEventsRepeatingGroup returns an initialized, NoEventsRepeatingGroup.
func NewNoEventsRepeatingGroup() NoEventsRepeatingGroup {
	return NoEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.EventType), quickfix.GroupElement(tag.EventDate), quickfix.GroupElement(tag.EventPx), quickfix.GroupElement(tag.EventText), quickfix.GroupElement(tag.EventTime)})}
}

// Add create and append a new NoEvents to this group.
func (m NoEventsRepeatingGroup) Add() NoEvents {
	g := m.RepeatingGroup.Add()
	return NoEvents{g}
}

// Get returns the ith NoEvents in the NoEventsRepeatinGroup.
func (m NoEventsRepeatingGroup) Get(i int) NoEvents {
	return NoEvents{m.RepeatingGroup.Get(i)}
}

// NoInstrumentParties is a repeating group element, Tag 1018.
type NoInstrumentParties struct {
	*quickfix.Group
}

// SetInstrumentPartyID sets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) SetInstrumentPartyID(v string) {
	m.Set(field.NewInstrumentPartyID(v))
}

// SetInstrumentPartyIDSource sets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) SetInstrumentPartyIDSource(v string) {
	m.Set(field.NewInstrumentPartyIDSource(v))
}

// SetInstrumentPartyRole sets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) SetInstrumentPartyRole(v int) {
	m.Set(field.NewInstrumentPartyRole(v))
}

// SetNoInstrumentPartySubIDs sets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) SetNoInstrumentPartySubIDs(f NoInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetInstrumentPartyID gets InstrumentPartyID, Tag 1019.
func (m NoInstrumentParties) GetInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyIDSource gets InstrumentPartyIDSource, Tag 1050.
func (m NoInstrumentParties) GetInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartyRole gets InstrumentPartyRole, Tag 1051.
func (m NoInstrumentParties) GetInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoInstrumentPartySubIDs gets NoInstrumentPartySubIDs, Tag 1052.
func (m NoInstrumentParties) GetNoInstrumentPartySubIDs() (NoInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasInstrumentPartyID returns true if InstrumentPartyID is present, Tag 1019.
func (m NoInstrumentParties) HasInstrumentPartyID() bool {
	return m.Has(tag.InstrumentPartyID)
}

// HasInstrumentPartyIDSource returns true if InstrumentPartyIDSource is present, Tag 1050.
func (m NoInstrumentParties) HasInstrumentPartyIDSource() bool {
	return m.Has(tag.InstrumentPartyIDSource)
}

// HasInstrumentPartyRole returns true if InstrumentPartyRole is present, Tag 1051.
func (m NoInstrumentParties) HasInstrumentPartyRole() bool {
	return m.Has(tag.InstrumentPartyRole)
}

// HasNoInstrumentPartySubIDs returns true if NoInstrumentPartySubIDs is present, Tag 1052.
func (m NoInstrumentParties) HasNoInstrumentPartySubIDs() bool {
	return m.Has(tag.NoInstrumentPartySubIDs)
}

// NoInstrumentPartySubIDs is a repeating group element, Tag 1052.
type NoInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetInstrumentPartySubID sets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubID(v string) {
	m.Set(field.NewInstrumentPartySubID(v))
}

// SetInstrumentPartySubIDType sets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) SetInstrumentPartySubIDType(v int) {
	m.Set(field.NewInstrumentPartySubIDType(v))
}

// GetInstrumentPartySubID gets InstrumentPartySubID, Tag 1053.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrumentPartySubIDType gets InstrumentPartySubIDType, Tag 1054.
func (m NoInstrumentPartySubIDs) GetInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.InstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasInstrumentPartySubID returns true if InstrumentPartySubID is present, Tag 1053.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubID() bool {
	return m.Has(tag.InstrumentPartySubID)
}

// HasInstrumentPartySubIDType returns true if InstrumentPartySubIDType is present, Tag 1054.
func (m NoInstrumentPartySubIDs) HasInstrumentPartySubIDType() bool {
	return m.Has(tag.InstrumentPartySubIDType)
}

// NoInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1052.
type NoInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartySubIDsRepeatingGroup returns an initialized, NoInstrumentPartySubIDsRepeatingGroup.
func NewNoInstrumentPartySubIDsRepeatingGroup() NoInstrumentPartySubIDsRepeatingGroup {
	return NoInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartySubID), quickfix.GroupElement(tag.InstrumentPartySubIDType)})}
}

// Add create and append a new NoInstrumentPartySubIDs to this group.
func (m NoInstrumentPartySubIDsRepeatingGroup) Add() NoInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoInstrumentPartySubIDs{g}
}

// Get returns the ith NoInstrumentPartySubIDs in the NoInstrumentPartySubIDsRepeatinGroup.
func (m NoInstrumentPartySubIDsRepeatingGroup) Get(i int) NoInstrumentPartySubIDs {
	return NoInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoInstrumentPartiesRepeatingGroup is a repeating group, Tag 1018.
type NoInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoInstrumentPartiesRepeatingGroup returns an initialized, NoInstrumentPartiesRepeatingGroup.
func NewNoInstrumentPartiesRepeatingGroup() NoInstrumentPartiesRepeatingGroup {
	return NoInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.InstrumentPartyID), quickfix.GroupElement(tag.InstrumentPartyIDSource), quickfix.GroupElement(tag.InstrumentPartyRole), NewNoInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoInstrumentParties to this group.
func (m NoInstrumentPartiesRepeatingGroup) Add() NoInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoInstrumentParties{g}
}

// Get returns the ith NoInstrumentParties in the NoInstrumentPartiesRepeatinGroup.
func (m NoInstrumentPartiesRepeatingGroup) Get(i int) NoInstrumentParties {
	return NoInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoComplexEvents is a repeating group element, Tag 1483.
type NoComplexEvents struct {
	*quickfix.Group
}

// SetComplexEventType sets ComplexEventType, Tag 1484.
func (m NoComplexEvents) SetComplexEventType(v enum.ComplexEventType) {
	m.Set(field.NewComplexEventType(v))
}

// SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) SetComplexOptPayoutAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexOptPayoutAmount(value, scale))
}

// SetComplexEventPrice sets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) SetComplexEventPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPrice(value, scale))
}

// SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v enum.ComplexEventPriceBoundaryMethod) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

// SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(value decimal.Decimal, scale int32) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(value, scale))
}

// SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) SetComplexEventPriceTimeType(v enum.ComplexEventPriceTimeType) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

// SetComplexEventCondition sets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) SetComplexEventCondition(v enum.ComplexEventCondition) {
	m.Set(field.NewComplexEventCondition(v))
}

// SetNoComplexEventDates sets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventType gets ComplexEventType, Tag 1484.
func (m NoComplexEvents) GetComplexEventType() (v enum.ComplexEventType, err quickfix.MessageRejectError) {
	var f field.ComplexEventTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485.
func (m NoComplexEvents) GetComplexOptPayoutAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexOptPayoutAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPrice gets ComplexEventPrice, Tag 1486.
func (m NoComplexEvents) GetComplexEventPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487.
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (v enum.ComplexEventPriceBoundaryMethod, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488.
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceBoundaryPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489.
func (m NoComplexEvents) GetComplexEventPriceTimeType() (v enum.ComplexEventPriceTimeType, err quickfix.MessageRejectError) {
	var f field.ComplexEventPriceTimeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventCondition gets ComplexEventCondition, Tag 1490.
func (m NoComplexEvents) GetComplexEventCondition() (v enum.ComplexEventCondition, err quickfix.MessageRejectError) {
	var f field.ComplexEventConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventDates gets NoComplexEventDates, Tag 1491.
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventType returns true if ComplexEventType is present, Tag 1484.
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

// HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485.
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

// HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486.
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

// HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487.
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

// HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488.
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

// HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489.
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

// HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490.
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

// HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491.
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

// NoComplexEventDates is a repeating group element, Tag 1491.
type NoComplexEventDates struct {
	*quickfix.Group
}

// SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

// SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

// SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

// GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492.
func (m NoComplexEventDates) GetComplexEventStartDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493.
func (m NoComplexEventDates) GetComplexEventEndDate() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494.
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492.
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

// HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493.
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

// HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494.
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

// NoComplexEventTimes is a repeating group element, Tag 1494.
type NoComplexEventTimes struct {
	*quickfix.Group
}

// SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

// SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

// GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495.
func (m NoComplexEventTimes) GetComplexEventStartTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496.
func (m NoComplexEventTimes) GetComplexEventEndTime() (v string, err quickfix.MessageRejectError) {
	var f field.ComplexEventEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495.
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

// HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496.
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

// NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494.
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup.
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

// Add create and append a new NoComplexEventTimes to this group.
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

// Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup.
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

// NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491.
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup.
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
}

// Add create and append a new NoComplexEventDates to this group.
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

// Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup.
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

// NoComplexEventsRepeatingGroup is a repeating group, Tag 1483.
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup.
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
}

// Add create and append a new NoComplexEvents to this group.
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

// Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup.
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}

// NoUnderlyings is a repeating group element, Tag 711.
type NoUnderlyings struct {
	*quickfix.Group
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSecuritySubType sets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) SetUnderlyingSecuritySubType(v string) {
	m.Set(field.NewUnderlyingSecuritySubType(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingStrikeCurrency sets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) SetUnderlyingStrikeCurrency(v string) {
	m.Set(field.NewUnderlyingStrikeCurrency(v))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCPProgram sets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) SetUnderlyingCPProgram(v string) {
	m.Set(field.NewUnderlyingCPProgram(v))
}

// SetUnderlyingCPRegType sets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) SetUnderlyingCPRegType(v string) {
	m.Set(field.NewUnderlyingCPRegType(v))
}

// SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

// SetUnderlyingQty sets UnderlyingQty, Tag 879.
func (m NoUnderlyings) SetUnderlyingQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingQty(value, scale))
}

// SetUnderlyingPx sets UnderlyingPx, Tag 810.
func (m NoUnderlyings) SetUnderlyingPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPx(value, scale))
}

// SetUnderlyingDirtyPrice sets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) SetUnderlyingDirtyPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDirtyPrice(value, scale))
}

// SetUnderlyingEndPrice sets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) SetUnderlyingEndPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndPrice(value, scale))
}

// SetUnderlyingStartValue sets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) SetUnderlyingStartValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStartValue(value, scale))
}

// SetUnderlyingCurrentValue sets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) SetUnderlyingCurrentValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCurrentValue(value, scale))
}

// SetUnderlyingEndValue sets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) SetUnderlyingEndValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingEndValue(value, scale))
}

// SetNoUnderlyingStips sets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) SetNoUnderlyingStips(f NoUnderlyingStipsRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingAllocationPercent sets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) SetUnderlyingAllocationPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAllocationPercent(value, scale))
}

// SetUnderlyingSettlementType sets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) SetUnderlyingSettlementType(v enum.UnderlyingSettlementType) {
	m.Set(field.NewUnderlyingSettlementType(v))
}

// SetUnderlyingCashAmount sets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) SetUnderlyingCashAmount(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCashAmount(value, scale))
}

// SetUnderlyingCashType sets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) SetUnderlyingCashType(v enum.UnderlyingCashType) {
	m.Set(field.NewUnderlyingCashType(v))
}

// SetUnderlyingUnitOfMeasure sets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingUnitOfMeasure(v))
}

// SetUnderlyingTimeUnit sets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) SetUnderlyingTimeUnit(v string) {
	m.Set(field.NewUnderlyingTimeUnit(v))
}

// SetUnderlyingCapValue sets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) SetUnderlyingCapValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCapValue(value, scale))
}

// SetNoUndlyInstrumentParties sets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) SetNoUndlyInstrumentParties(f NoUndlyInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

// SetUnderlyingSettlMethod sets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) SetUnderlyingSettlMethod(v string) {
	m.Set(field.NewUnderlyingSettlMethod(v))
}

// SetUnderlyingAdjustedQuantity sets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) SetUnderlyingAdjustedQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAdjustedQuantity(value, scale))
}

// SetUnderlyingFXRate sets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) SetUnderlyingFXRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFXRate(value, scale))
}

// SetUnderlyingFXRateCalc sets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) SetUnderlyingFXRateCalc(v enum.UnderlyingFXRateCalc) {
	m.Set(field.NewUnderlyingFXRateCalc(v))
}

// SetUnderlyingMaturityTime sets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) SetUnderlyingMaturityTime(v string) {
	m.Set(field.NewUnderlyingMaturityTime(v))
}

// SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

// SetUnderlyingExerciseStyle sets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) SetUnderlyingExerciseStyle(v int) {
	m.Set(field.NewUnderlyingExerciseStyle(v))
}

// SetUnderlyingUnitOfMeasureQty sets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) SetUnderlyingUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingUnitOfMeasureQty(value, scale))
}

// SetUnderlyingPriceUnitOfMeasure sets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasure(v string) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasure(v))
}

// SetUnderlyingPriceUnitOfMeasureQty sets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) SetUnderlyingPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingPriceUnitOfMeasureQty(value, scale))
}

// SetUnderlyingContractMultiplierUnit sets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) SetUnderlyingContractMultiplierUnit(v int) {
	m.Set(field.NewUnderlyingContractMultiplierUnit(v))
}

// SetUnderlyingFlowScheduleType sets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) SetUnderlyingFlowScheduleType(v int) {
	m.Set(field.NewUnderlyingFlowScheduleType(v))
}

// SetUnderlyingRestructuringType sets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) SetUnderlyingRestructuringType(v string) {
	m.Set(field.NewUnderlyingRestructuringType(v))
}

// SetUnderlyingSeniority sets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) SetUnderlyingSeniority(v string) {
	m.Set(field.NewUnderlyingSeniority(v))
}

// SetUnderlyingNotionalPercentageOutstanding sets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) SetUnderlyingNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingOriginalNotionalPercentageOutstanding sets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) SetUnderlyingOriginalNotionalPercentageOutstanding(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingOriginalNotionalPercentageOutstanding(value, scale))
}

// SetUnderlyingAttachmentPoint sets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) SetUnderlyingAttachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingAttachmentPoint(value, scale))
}

// SetUnderlyingDetachmentPoint sets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) SetUnderlyingDetachmentPoint(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingDetachmentPoint(value, scale))
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311.
func (m NoUnderlyings) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312.
func (m NoUnderlyings) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309.
func (m NoUnderlyings) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305.
func (m NoUnderlyings) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457.
func (m NoUnderlyings) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462.
func (m NoUnderlyings) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463.
func (m NoUnderlyings) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310.
func (m NoUnderlyings) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecuritySubType gets UnderlyingSecuritySubType, Tag 763.
func (m NoUnderlyings) GetUnderlyingSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313.
func (m NoUnderlyings) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542.
func (m NoUnderlyings) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241.
func (m NoUnderlyings) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242.
func (m NoUnderlyings) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m NoUnderlyings) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244.
func (m NoUnderlyings) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245.
func (m NoUnderlyings) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246.
func (m NoUnderlyings) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256.
func (m NoUnderlyings) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595.
func (m NoUnderlyings) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592.
func (m NoUnderlyings) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m NoUnderlyings) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594.
func (m NoUnderlyings) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247.
func (m NoUnderlyings) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316.
func (m NoUnderlyings) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikeCurrency gets UnderlyingStrikeCurrency, Tag 941.
func (m NoUnderlyings) GetUnderlyingStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317.
func (m NoUnderlyings) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436.
func (m NoUnderlyings) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435.
func (m NoUnderlyings) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308.
func (m NoUnderlyings) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306.
func (m NoUnderlyings) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362.
func (m NoUnderlyings) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363.
func (m NoUnderlyings) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307.
func (m NoUnderlyings) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365.
func (m NoUnderlyings) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPProgram gets UnderlyingCPProgram, Tag 877.
func (m NoUnderlyings) GetUnderlyingCPProgram() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPProgramField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCPRegType gets UnderlyingCPRegType, Tag 878.
func (m NoUnderlyings) GetUnderlyingCPRegType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCPRegTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318.
func (m NoUnderlyings) GetUnderlyingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingQty gets UnderlyingQty, Tag 879.
func (m NoUnderlyings) GetUnderlyingQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPx gets UnderlyingPx, Tag 810.
func (m NoUnderlyings) GetUnderlyingPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDirtyPrice gets UnderlyingDirtyPrice, Tag 882.
func (m NoUnderlyings) GetUnderlyingDirtyPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDirtyPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndPrice gets UnderlyingEndPrice, Tag 883.
func (m NoUnderlyings) GetUnderlyingEndPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStartValue gets UnderlyingStartValue, Tag 884.
func (m NoUnderlyings) GetUnderlyingStartValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStartValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCurrentValue gets UnderlyingCurrentValue, Tag 885.
func (m NoUnderlyings) GetUnderlyingCurrentValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCurrentValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingEndValue gets UnderlyingEndValue, Tag 886.
func (m NoUnderlyings) GetUnderlyingEndValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingEndValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUnderlyingStips gets NoUnderlyingStips, Tag 887.
func (m NoUnderlyings) GetNoUnderlyingStips() (NoUnderlyingStipsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingStipsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingAllocationPercent gets UnderlyingAllocationPercent, Tag 972.
func (m NoUnderlyings) GetUnderlyingAllocationPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAllocationPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSettlementType gets UnderlyingSettlementType, Tag 975.
func (m NoUnderlyings) GetUnderlyingSettlementType() (v enum.UnderlyingSettlementType, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlementTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashAmount gets UnderlyingCashAmount, Tag 973.
func (m NoUnderlyings) GetUnderlyingCashAmount() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashAmountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCashType gets UnderlyingCashType, Tag 974.
func (m NoUnderlyings) GetUnderlyingCashType() (v enum.UnderlyingCashType, err quickfix.MessageRejectError) {
	var f field.UnderlyingCashTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasure gets UnderlyingUnitOfMeasure, Tag 998.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingTimeUnit gets UnderlyingTimeUnit, Tag 1000.
func (m NoUnderlyings) GetUnderlyingTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCapValue gets UnderlyingCapValue, Tag 1038.
func (m NoUnderlyings) GetUnderlyingCapValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCapValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentParties gets NoUndlyInstrumentParties, Tag 1058.
func (m NoUnderlyings) GetNoUndlyInstrumentParties() (NoUndlyInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetUnderlyingSettlMethod gets UnderlyingSettlMethod, Tag 1039.
func (m NoUnderlyings) GetUnderlyingSettlMethod() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSettlMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAdjustedQuantity gets UnderlyingAdjustedQuantity, Tag 1044.
func (m NoUnderlyings) GetUnderlyingAdjustedQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAdjustedQuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRate gets UnderlyingFXRate, Tag 1045.
func (m NoUnderlyings) GetUnderlyingFXRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFXRateCalc gets UnderlyingFXRateCalc, Tag 1046.
func (m NoUnderlyings) GetUnderlyingFXRateCalc() (v enum.UnderlyingFXRateCalc, err quickfix.MessageRejectError) {
	var f field.UnderlyingFXRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityTime gets UnderlyingMaturityTime, Tag 1213.
func (m NoUnderlyings) GetUnderlyingMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315.
func (m NoUnderlyings) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingExerciseStyle gets UnderlyingExerciseStyle, Tag 1419.
func (m NoUnderlyings) GetUnderlyingExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingUnitOfMeasureQty gets UnderlyingUnitOfMeasureQty, Tag 1423.
func (m NoUnderlyings) GetUnderlyingUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasure gets UnderlyingPriceUnitOfMeasure, Tag 1424.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingPriceUnitOfMeasureQty gets UnderlyingPriceUnitOfMeasureQty, Tag 1425.
func (m NoUnderlyings) GetUnderlyingPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplierUnit gets UnderlyingContractMultiplierUnit, Tag 1437.
func (m NoUnderlyings) GetUnderlyingContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFlowScheduleType gets UnderlyingFlowScheduleType, Tag 1441.
func (m NoUnderlyings) GetUnderlyingFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRestructuringType gets UnderlyingRestructuringType, Tag 1453.
func (m NoUnderlyings) GetUnderlyingRestructuringType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRestructuringTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSeniority gets UnderlyingSeniority, Tag 1454.
func (m NoUnderlyings) GetUnderlyingSeniority() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSeniorityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingNotionalPercentageOutstanding gets UnderlyingNotionalPercentageOutstanding, Tag 1455.
func (m NoUnderlyings) GetUnderlyingNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOriginalNotionalPercentageOutstanding gets UnderlyingOriginalNotionalPercentageOutstanding, Tag 1456.
func (m NoUnderlyings) GetUnderlyingOriginalNotionalPercentageOutstanding() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingOriginalNotionalPercentageOutstandingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingAttachmentPoint gets UnderlyingAttachmentPoint, Tag 1459.
func (m NoUnderlyings) GetUnderlyingAttachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingAttachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingDetachmentPoint gets UnderlyingDetachmentPoint, Tag 1460.
func (m NoUnderlyings) GetUnderlyingDetachmentPoint() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingDetachmentPointField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311.
func (m NoUnderlyings) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312.
func (m NoUnderlyings) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309.
func (m NoUnderlyings) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305.
func (m NoUnderlyings) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457.
func (m NoUnderlyings) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462.
func (m NoUnderlyings) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463.
func (m NoUnderlyings) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310.
func (m NoUnderlyings) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSecuritySubType returns true if UnderlyingSecuritySubType is present, Tag 763.
func (m NoUnderlyings) HasUnderlyingSecuritySubType() bool {
	return m.Has(tag.UnderlyingSecuritySubType)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313.
func (m NoUnderlyings) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542.
func (m NoUnderlyings) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241.
func (m NoUnderlyings) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242.
func (m NoUnderlyings) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243.
func (m NoUnderlyings) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244.
func (m NoUnderlyings) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245.
func (m NoUnderlyings) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246.
func (m NoUnderlyings) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256.
func (m NoUnderlyings) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595.
func (m NoUnderlyings) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592.
func (m NoUnderlyings) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593.
func (m NoUnderlyings) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594.
func (m NoUnderlyings) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247.
func (m NoUnderlyings) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316.
func (m NoUnderlyings) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingStrikeCurrency returns true if UnderlyingStrikeCurrency is present, Tag 941.
func (m NoUnderlyings) HasUnderlyingStrikeCurrency() bool {
	return m.Has(tag.UnderlyingStrikeCurrency)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317.
func (m NoUnderlyings) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436.
func (m NoUnderlyings) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435.
func (m NoUnderlyings) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308.
func (m NoUnderlyings) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306.
func (m NoUnderlyings) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362.
func (m NoUnderlyings) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363.
func (m NoUnderlyings) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307.
func (m NoUnderlyings) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365.
func (m NoUnderlyings) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCPProgram returns true if UnderlyingCPProgram is present, Tag 877.
func (m NoUnderlyings) HasUnderlyingCPProgram() bool {
	return m.Has(tag.UnderlyingCPProgram)
}

// HasUnderlyingCPRegType returns true if UnderlyingCPRegType is present, Tag 878.
func (m NoUnderlyings) HasUnderlyingCPRegType() bool {
	return m.Has(tag.UnderlyingCPRegType)
}

// HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318.
func (m NoUnderlyings) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

// HasUnderlyingQty returns true if UnderlyingQty is present, Tag 879.
func (m NoUnderlyings) HasUnderlyingQty() bool {
	return m.Has(tag.UnderlyingQty)
}

// HasUnderlyingPx returns true if UnderlyingPx is present, Tag 810.
func (m NoUnderlyings) HasUnderlyingPx() bool {
	return m.Has(tag.UnderlyingPx)
}

// HasUnderlyingDirtyPrice returns true if UnderlyingDirtyPrice is present, Tag 882.
func (m NoUnderlyings) HasUnderlyingDirtyPrice() bool {
	return m.Has(tag.UnderlyingDirtyPrice)
}

// HasUnderlyingEndPrice returns true if UnderlyingEndPrice is present, Tag 883.
func (m NoUnderlyings) HasUnderlyingEndPrice() bool {
	return m.Has(tag.UnderlyingEndPrice)
}

// HasUnderlyingStartValue returns true if UnderlyingStartValue is present, Tag 884.
func (m NoUnderlyings) HasUnderlyingStartValue() bool {
	return m.Has(tag.UnderlyingStartValue)
}

// HasUnderlyingCurrentValue returns true if UnderlyingCurrentValue is present, Tag 885.
func (m NoUnderlyings) HasUnderlyingCurrentValue() bool {
	return m.Has(tag.UnderlyingCurrentValue)
}

// HasUnderlyingEndValue returns true if UnderlyingEndValue is present, Tag 886.
func (m NoUnderlyings) HasUnderlyingEndValue() bool {
	return m.Has(tag.UnderlyingEndValue)
}

// HasNoUnderlyingStips returns true if NoUnderlyingStips is present, Tag 887.
func (m NoUnderlyings) HasNoUnderlyingStips() bool {
	return m.Has(tag.NoUnderlyingStips)
}

// HasUnderlyingAllocationPercent returns true if UnderlyingAllocationPercent is present, Tag 972.
func (m NoUnderlyings) HasUnderlyingAllocationPercent() bool {
	return m.Has(tag.UnderlyingAllocationPercent)
}

// HasUnderlyingSettlementType returns true if UnderlyingSettlementType is present, Tag 975.
func (m NoUnderlyings) HasUnderlyingSettlementType() bool {
	return m.Has(tag.UnderlyingSettlementType)
}

// HasUnderlyingCashAmount returns true if UnderlyingCashAmount is present, Tag 973.
func (m NoUnderlyings) HasUnderlyingCashAmount() bool {
	return m.Has(tag.UnderlyingCashAmount)
}

// HasUnderlyingCashType returns true if UnderlyingCashType is present, Tag 974.
func (m NoUnderlyings) HasUnderlyingCashType() bool {
	return m.Has(tag.UnderlyingCashType)
}

// HasUnderlyingUnitOfMeasure returns true if UnderlyingUnitOfMeasure is present, Tag 998.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingUnitOfMeasure)
}

// HasUnderlyingTimeUnit returns true if UnderlyingTimeUnit is present, Tag 1000.
func (m NoUnderlyings) HasUnderlyingTimeUnit() bool {
	return m.Has(tag.UnderlyingTimeUnit)
}

// HasUnderlyingCapValue returns true if UnderlyingCapValue is present, Tag 1038.
func (m NoUnderlyings) HasUnderlyingCapValue() bool {
	return m.Has(tag.UnderlyingCapValue)
}

// HasNoUndlyInstrumentParties returns true if NoUndlyInstrumentParties is present, Tag 1058.
func (m NoUnderlyings) HasNoUndlyInstrumentParties() bool {
	return m.Has(tag.NoUndlyInstrumentParties)
}

// HasUnderlyingSettlMethod returns true if UnderlyingSettlMethod is present, Tag 1039.
func (m NoUnderlyings) HasUnderlyingSettlMethod() bool {
	return m.Has(tag.UnderlyingSettlMethod)
}

// HasUnderlyingAdjustedQuantity returns true if UnderlyingAdjustedQuantity is present, Tag 1044.
func (m NoUnderlyings) HasUnderlyingAdjustedQuantity() bool {
	return m.Has(tag.UnderlyingAdjustedQuantity)
}

// HasUnderlyingFXRate returns true if UnderlyingFXRate is present, Tag 1045.
func (m NoUnderlyings) HasUnderlyingFXRate() bool {
	return m.Has(tag.UnderlyingFXRate)
}

// HasUnderlyingFXRateCalc returns true if UnderlyingFXRateCalc is present, Tag 1046.
func (m NoUnderlyings) HasUnderlyingFXRateCalc() bool {
	return m.Has(tag.UnderlyingFXRateCalc)
}

// HasUnderlyingMaturityTime returns true if UnderlyingMaturityTime is present, Tag 1213.
func (m NoUnderlyings) HasUnderlyingMaturityTime() bool {
	return m.Has(tag.UnderlyingMaturityTime)
}

// HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315.
func (m NoUnderlyings) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

// HasUnderlyingExerciseStyle returns true if UnderlyingExerciseStyle is present, Tag 1419.
func (m NoUnderlyings) HasUnderlyingExerciseStyle() bool {
	return m.Has(tag.UnderlyingExerciseStyle)
}

// HasUnderlyingUnitOfMeasureQty returns true if UnderlyingUnitOfMeasureQty is present, Tag 1423.
func (m NoUnderlyings) HasUnderlyingUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingUnitOfMeasureQty)
}

// HasUnderlyingPriceUnitOfMeasure returns true if UnderlyingPriceUnitOfMeasure is present, Tag 1424.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasure() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasure)
}

// HasUnderlyingPriceUnitOfMeasureQty returns true if UnderlyingPriceUnitOfMeasureQty is present, Tag 1425.
func (m NoUnderlyings) HasUnderlyingPriceUnitOfMeasureQty() bool {
	return m.Has(tag.UnderlyingPriceUnitOfMeasureQty)
}

// HasUnderlyingContractMultiplierUnit returns true if UnderlyingContractMultiplierUnit is present, Tag 1437.
func (m NoUnderlyings) HasUnderlyingContractMultiplierUnit() bool {
	return m.Has(tag.UnderlyingContractMultiplierUnit)
}

// HasUnderlyingFlowScheduleType returns true if UnderlyingFlowScheduleType is present, Tag 1441.
func (m NoUnderlyings) HasUnderlyingFlowScheduleType() bool {
	return m.Has(tag.UnderlyingFlowScheduleType)
}

// HasUnderlyingRestructuringType returns true if UnderlyingRestructuringType is present, Tag 1453.
func (m NoUnderlyings) HasUnderlyingRestructuringType() bool {
	return m.Has(tag.UnderlyingRestructuringType)
}

// HasUnderlyingSeniority returns true if UnderlyingSeniority is present, Tag 1454.
func (m NoUnderlyings) HasUnderlyingSeniority() bool {
	return m.Has(tag.UnderlyingSeniority)
}

// HasUnderlyingNotionalPercentageOutstanding returns true if UnderlyingNotionalPercentageOutstanding is present, Tag 1455.
func (m NoUnderlyings) HasUnderlyingNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingNotionalPercentageOutstanding)
}

// HasUnderlyingOriginalNotionalPercentageOutstanding returns true if UnderlyingOriginalNotionalPercentageOutstanding is present, Tag 1456.
func (m NoUnderlyings) HasUnderlyingOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.UnderlyingOriginalNotionalPercentageOutstanding)
}

// HasUnderlyingAttachmentPoint returns true if UnderlyingAttachmentPoint is present, Tag 1459.
func (m NoUnderlyings) HasUnderlyingAttachmentPoint() bool {
	return m.Has(tag.UnderlyingAttachmentPoint)
}

// HasUnderlyingDetachmentPoint returns true if UnderlyingDetachmentPoint is present, Tag 1460.
func (m NoUnderlyings) HasUnderlyingDetachmentPoint() bool {
	return m.Has(tag.UnderlyingDetachmentPoint)
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457.
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457.
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup.
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingStips is a repeating group element, Tag 887.
type NoUnderlyingStips struct {
	*quickfix.Group
}

// SetUnderlyingStipType sets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) SetUnderlyingStipType(v string) {
	m.Set(field.NewUnderlyingStipType(v))
}

// SetUnderlyingStipValue sets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) SetUnderlyingStipValue(v string) {
	m.Set(field.NewUnderlyingStipValue(v))
}

// GetUnderlyingStipType gets UnderlyingStipType, Tag 888.
func (m NoUnderlyingStips) GetUnderlyingStipType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStipValue gets UnderlyingStipValue, Tag 889.
func (m NoUnderlyingStips) GetUnderlyingStipValue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStipValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingStipType returns true if UnderlyingStipType is present, Tag 888.
func (m NoUnderlyingStips) HasUnderlyingStipType() bool {
	return m.Has(tag.UnderlyingStipType)
}

// HasUnderlyingStipValue returns true if UnderlyingStipValue is present, Tag 889.
func (m NoUnderlyingStips) HasUnderlyingStipValue() bool {
	return m.Has(tag.UnderlyingStipValue)
}

// NoUnderlyingStipsRepeatingGroup is a repeating group, Tag 887.
type NoUnderlyingStipsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingStipsRepeatingGroup returns an initialized, NoUnderlyingStipsRepeatingGroup.
func NewNoUnderlyingStipsRepeatingGroup() NoUnderlyingStipsRepeatingGroup {
	return NoUnderlyingStipsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingStips,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingStipType), quickfix.GroupElement(tag.UnderlyingStipValue)})}
}

// Add create and append a new NoUnderlyingStips to this group.
func (m NoUnderlyingStipsRepeatingGroup) Add() NoUnderlyingStips {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingStips{g}
}

// Get returns the ith NoUnderlyingStips in the NoUnderlyingStipsRepeatinGroup.
func (m NoUnderlyingStipsRepeatingGroup) Get(i int) NoUnderlyingStips {
	return NoUnderlyingStips{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentParties is a repeating group element, Tag 1058.
type NoUndlyInstrumentParties struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartyID sets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyID(v))
}

// SetUnderlyingInstrumentPartyIDSource sets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyIDSource(v string) {
	m.Set(field.NewUnderlyingInstrumentPartyIDSource(v))
}

// SetUnderlyingInstrumentPartyRole sets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) SetUnderlyingInstrumentPartyRole(v int) {
	m.Set(field.NewUnderlyingInstrumentPartyRole(v))
}

// SetNoUndlyInstrumentPartySubIDs sets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) SetNoUndlyInstrumentPartySubIDs(f NoUndlyInstrumentPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetUnderlyingInstrumentPartyID gets UnderlyingInstrumentPartyID, Tag 1059.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyIDSource gets UnderlyingInstrumentPartyIDSource, Tag 1060.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartyRole gets UnderlyingInstrumentPartyRole, Tag 1061.
func (m NoUndlyInstrumentParties) GetUnderlyingInstrumentPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoUndlyInstrumentPartySubIDs gets NoUndlyInstrumentPartySubIDs, Tag 1062.
func (m NoUndlyInstrumentParties) GetNoUndlyInstrumentPartySubIDs() (NoUndlyInstrumentPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUndlyInstrumentPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasUnderlyingInstrumentPartyID returns true if UnderlyingInstrumentPartyID is present, Tag 1059.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyID() bool {
	return m.Has(tag.UnderlyingInstrumentPartyID)
}

// HasUnderlyingInstrumentPartyIDSource returns true if UnderlyingInstrumentPartyIDSource is present, Tag 1060.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyIDSource() bool {
	return m.Has(tag.UnderlyingInstrumentPartyIDSource)
}

// HasUnderlyingInstrumentPartyRole returns true if UnderlyingInstrumentPartyRole is present, Tag 1061.
func (m NoUndlyInstrumentParties) HasUnderlyingInstrumentPartyRole() bool {
	return m.Has(tag.UnderlyingInstrumentPartyRole)
}

// HasNoUndlyInstrumentPartySubIDs returns true if NoUndlyInstrumentPartySubIDs is present, Tag 1062.
func (m NoUndlyInstrumentParties) HasNoUndlyInstrumentPartySubIDs() bool {
	return m.Has(tag.NoUndlyInstrumentPartySubIDs)
}

// NoUndlyInstrumentPartySubIDs is a repeating group element, Tag 1062.
type NoUndlyInstrumentPartySubIDs struct {
	*quickfix.Group
}

// SetUnderlyingInstrumentPartySubID sets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubID(v string) {
	m.Set(field.NewUnderlyingInstrumentPartySubID(v))
}

// SetUnderlyingInstrumentPartySubIDType sets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) SetUnderlyingInstrumentPartySubIDType(v int) {
	m.Set(field.NewUnderlyingInstrumentPartySubIDType(v))
}

// GetUnderlyingInstrumentPartySubID gets UnderlyingInstrumentPartySubID, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrumentPartySubIDType gets UnderlyingInstrumentPartySubIDType, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) GetUnderlyingInstrumentPartySubIDType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrumentPartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingInstrumentPartySubID returns true if UnderlyingInstrumentPartySubID is present, Tag 1063.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubID() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubID)
}

// HasUnderlyingInstrumentPartySubIDType returns true if UnderlyingInstrumentPartySubIDType is present, Tag 1064.
func (m NoUndlyInstrumentPartySubIDs) HasUnderlyingInstrumentPartySubIDType() bool {
	return m.Has(tag.UnderlyingInstrumentPartySubIDType)
}

// NoUndlyInstrumentPartySubIDsRepeatingGroup is a repeating group, Tag 1062.
type NoUndlyInstrumentPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartySubIDsRepeatingGroup returns an initialized, NoUndlyInstrumentPartySubIDsRepeatingGroup.
func NewNoUndlyInstrumentPartySubIDsRepeatingGroup() NoUndlyInstrumentPartySubIDsRepeatingGroup {
	return NoUndlyInstrumentPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartySubID), quickfix.GroupElement(tag.UnderlyingInstrumentPartySubIDType)})}
}

// Add create and append a new NoUndlyInstrumentPartySubIDs to this group.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Add() NoUndlyInstrumentPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentPartySubIDs{g}
}

// Get returns the ith NoUndlyInstrumentPartySubIDs in the NoUndlyInstrumentPartySubIDsRepeatinGroup.
func (m NoUndlyInstrumentPartySubIDsRepeatingGroup) Get(i int) NoUndlyInstrumentPartySubIDs {
	return NoUndlyInstrumentPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoUndlyInstrumentPartiesRepeatingGroup is a repeating group, Tag 1058.
type NoUndlyInstrumentPartiesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUndlyInstrumentPartiesRepeatingGroup returns an initialized, NoUndlyInstrumentPartiesRepeatingGroup.
func NewNoUndlyInstrumentPartiesRepeatingGroup() NoUndlyInstrumentPartiesRepeatingGroup {
	return NoUndlyInstrumentPartiesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUndlyInstrumentParties,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingInstrumentPartyID), quickfix.GroupElement(tag.UnderlyingInstrumentPartyIDSource), quickfix.GroupElement(tag.UnderlyingInstrumentPartyRole), NewNoUndlyInstrumentPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoUndlyInstrumentParties to this group.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Add() NoUndlyInstrumentParties {
	g := m.RepeatingGroup.Add()
	return NoUndlyInstrumentParties{g}
}

// Get returns the ith NoUndlyInstrumentParties in the NoUndlyInstrumentPartiesRepeatinGroup.
func (m NoUndlyInstrumentPartiesRepeatingGroup) Get(i int) NoUndlyInstrumentParties {
	return NoUndlyInstrumentParties{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingsRepeatingGroup is a repeating group, Tag 711.
type NoUnderlyingsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingsRepeatingGroup returns an initialized, NoUnderlyingsRepeatingGroup.
func NewNoUnderlyingsRepeatingGroup() NoUnderlyingsRepeatingGroup {
	return NoUnderlyingsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyings,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingSecurityIDSource), NewNoUnderlyingSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingProduct), quickfix.GroupElement(tag.UnderlyingCFICode), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingSecuritySubType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDate), quickfix.GroupElement(tag.UnderlyingCouponPaymentDate), quickfix.GroupElement(tag.UnderlyingIssueDate), quickfix.GroupElement(tag.UnderlyingRepoCollateralSecurityType), quickfix.GroupElement(tag.UnderlyingRepurchaseTerm), quickfix.GroupElement(tag.UnderlyingRepurchaseRate), quickfix.GroupElement(tag.UnderlyingFactor), quickfix.GroupElement(tag.UnderlyingCreditRating), quickfix.GroupElement(tag.UnderlyingInstrRegistry), quickfix.GroupElement(tag.UnderlyingCountryOfIssue), quickfix.GroupElement(tag.UnderlyingStateOrProvinceOfIssue), quickfix.GroupElement(tag.UnderlyingLocaleOfIssue), quickfix.GroupElement(tag.UnderlyingRedemptionDate), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingStrikeCurrency), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.UnderlyingCPProgram), quickfix.GroupElement(tag.UnderlyingCPRegType), quickfix.GroupElement(tag.UnderlyingCurrency), quickfix.GroupElement(tag.UnderlyingQty), quickfix.GroupElement(tag.UnderlyingPx), quickfix.GroupElement(tag.UnderlyingDirtyPrice), quickfix.GroupElement(tag.UnderlyingEndPrice), quickfix.GroupElement(tag.UnderlyingStartValue), quickfix.GroupElement(tag.UnderlyingCurrentValue), quickfix.GroupElement(tag.UnderlyingEndValue), NewNoUnderlyingStipsRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingAllocationPercent), quickfix.GroupElement(tag.UnderlyingSettlementType), quickfix.GroupElement(tag.UnderlyingCashAmount), quickfix.GroupElement(tag.UnderlyingCashType), quickfix.GroupElement(tag.UnderlyingUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingTimeUnit), quickfix.GroupElement(tag.UnderlyingCapValue), NewNoUndlyInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnderlyingSettlMethod), quickfix.GroupElement(tag.UnderlyingAdjustedQuantity), quickfix.GroupElement(tag.UnderlyingFXRate), quickfix.GroupElement(tag.UnderlyingFXRateCalc), quickfix.GroupElement(tag.UnderlyingMaturityTime), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingExerciseStyle), quickfix.GroupElement(tag.UnderlyingUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasure), quickfix.GroupElement(tag.UnderlyingPriceUnitOfMeasureQty), quickfix.GroupElement(tag.UnderlyingContractMultiplierUnit), quickfix.GroupElement(tag.UnderlyingFlowScheduleType), quickfix.GroupElement(tag.UnderlyingRestructuringType), quickfix.GroupElement(tag.UnderlyingSeniority), quickfix.GroupElement(tag.UnderlyingNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingOriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.UnderlyingAttachmentPoint), quickfix.GroupElement(tag.UnderlyingDetachmentPoint)})}
}

// Add create and append a new NoUnderlyings to this group.
func (m NoUnderlyingsRepeatingGroup) Add() NoUnderlyings {
	g := m.RepeatingGroup.Add()
	return NoUnderlyings{g}
}

// Get returns the ith NoUnderlyings in the NoUnderlyingsRepeatinGroup.
func (m NoUnderlyingsRepeatingGroup) Get(i int) NoUnderlyings {
	return NoUnderlyings{m.RepeatingGroup.Get(i)}
}

// NoLegs is a repeating group element, Tag 555.
type NoLegs struct {
	*quickfix.Group
}

// SetLegSymbol sets LegSymbol, Tag 600.
func (m NoLegs) SetLegSymbol(v string) {
	m.Set(field.NewLegSymbol(v))
}

// SetLegSymbolSfx sets LegSymbolSfx, Tag 601.
func (m NoLegs) SetLegSymbolSfx(v string) {
	m.Set(field.NewLegSymbolSfx(v))
}

// SetLegSecurityID sets LegSecurityID, Tag 602.
func (m NoLegs) SetLegSecurityID(v string) {
	m.Set(field.NewLegSecurityID(v))
}

// SetLegSecurityIDSource sets LegSecurityIDSource, Tag 603.
func (m NoLegs) SetLegSecurityIDSource(v string) {
	m.Set(field.NewLegSecurityIDSource(v))
}

// SetNoLegSecurityAltID sets NoLegSecurityAltID, Tag 604.
func (m NoLegs) SetNoLegSecurityAltID(f NoLegSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetLegProduct sets LegProduct, Tag 607.
func (m NoLegs) SetLegProduct(v int) {
	m.Set(field.NewLegProduct(v))
}

// SetLegCFICode sets LegCFICode, Tag 608.
func (m NoLegs) SetLegCFICode(v string) {
	m.Set(field.NewLegCFICode(v))
}

// SetLegSecurityType sets LegSecurityType, Tag 609.
func (m NoLegs) SetLegSecurityType(v string) {
	m.Set(field.NewLegSecurityType(v))
}

// SetLegSecuritySubType sets LegSecuritySubType, Tag 764.
func (m NoLegs) SetLegSecuritySubType(v string) {
	m.Set(field.NewLegSecuritySubType(v))
}

// SetLegMaturityMonthYear sets LegMaturityMonthYear, Tag 610.
func (m NoLegs) SetLegMaturityMonthYear(v string) {
	m.Set(field.NewLegMaturityMonthYear(v))
}

// SetLegMaturityDate sets LegMaturityDate, Tag 611.
func (m NoLegs) SetLegMaturityDate(v string) {
	m.Set(field.NewLegMaturityDate(v))
}

// SetLegCouponPaymentDate sets LegCouponPaymentDate, Tag 248.
func (m NoLegs) SetLegCouponPaymentDate(v string) {
	m.Set(field.NewLegCouponPaymentDate(v))
}

// SetLegIssueDate sets LegIssueDate, Tag 249.
func (m NoLegs) SetLegIssueDate(v string) {
	m.Set(field.NewLegIssueDate(v))
}

// SetLegRepoCollateralSecurityType sets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) SetLegRepoCollateralSecurityType(v int) {
	m.Set(field.NewLegRepoCollateralSecurityType(v))
}

// SetLegRepurchaseTerm sets LegRepurchaseTerm, Tag 251.
func (m NoLegs) SetLegRepurchaseTerm(v int) {
	m.Set(field.NewLegRepurchaseTerm(v))
}

// SetLegRepurchaseRate sets LegRepurchaseRate, Tag 252.
func (m NoLegs) SetLegRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRepurchaseRate(value, scale))
}

// SetLegFactor sets LegFactor, Tag 253.
func (m NoLegs) SetLegFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegFactor(value, scale))
}

// SetLegCreditRating sets LegCreditRating, Tag 257.
func (m NoLegs) SetLegCreditRating(v string) {
	m.Set(field.NewLegCreditRating(v))
}

// SetLegInstrRegistry sets LegInstrRegistry, Tag 599.
func (m NoLegs) SetLegInstrRegistry(v string) {
	m.Set(field.NewLegInstrRegistry(v))
}

// SetLegCountryOfIssue sets LegCountryOfIssue, Tag 596.
func (m NoLegs) SetLegCountryOfIssue(v string) {
	m.Set(field.NewLegCountryOfIssue(v))
}

// SetLegStateOrProvinceOfIssue sets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) SetLegStateOrProvinceOfIssue(v string) {
	m.Set(field.NewLegStateOrProvinceOfIssue(v))
}

// SetLegLocaleOfIssue sets LegLocaleOfIssue, Tag 598.
func (m NoLegs) SetLegLocaleOfIssue(v string) {
	m.Set(field.NewLegLocaleOfIssue(v))
}

// SetLegRedemptionDate sets LegRedemptionDate, Tag 254.
func (m NoLegs) SetLegRedemptionDate(v string) {
	m.Set(field.NewLegRedemptionDate(v))
}

// SetLegStrikePrice sets LegStrikePrice, Tag 612.
func (m NoLegs) SetLegStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegStrikePrice(value, scale))
}

// SetLegStrikeCurrency sets LegStrikeCurrency, Tag 942.
func (m NoLegs) SetLegStrikeCurrency(v string) {
	m.Set(field.NewLegStrikeCurrency(v))
}

// SetLegOptAttribute sets LegOptAttribute, Tag 613.
func (m NoLegs) SetLegOptAttribute(v string) {
	m.Set(field.NewLegOptAttribute(v))
}

// SetLegContractMultiplier sets LegContractMultiplier, Tag 614.
func (m NoLegs) SetLegContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegContractMultiplier(value, scale))
}

// SetLegCouponRate sets LegCouponRate, Tag 615.
func (m NoLegs) SetLegCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegCouponRate(value, scale))
}

// SetLegSecurityExchange sets LegSecurityExchange, Tag 616.
func (m NoLegs) SetLegSecurityExchange(v string) {
	m.Set(field.NewLegSecurityExchange(v))
}

// SetLegIssuer sets LegIssuer, Tag 617.
func (m NoLegs) SetLegIssuer(v string) {
	m.Set(field.NewLegIssuer(v))
}

// SetEncodedLegIssuerLen sets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) SetEncodedLegIssuerLen(v int) {
	m.Set(field.NewEncodedLegIssuerLen(v))
}

// SetEncodedLegIssuer sets EncodedLegIssuer, Tag 619.
func (m NoLegs) SetEncodedLegIssuer(v string) {
	m.Set(field.NewEncodedLegIssuer(v))
}

// SetLegSecurityDesc sets LegSecurityDesc, Tag 620.
func (m NoLegs) SetLegSecurityDesc(v string) {
	m.Set(field.NewLegSecurityDesc(v))
}

// SetEncodedLegSecurityDescLen sets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) SetEncodedLegSecurityDescLen(v int) {
	m.Set(field.NewEncodedLegSecurityDescLen(v))
}

// SetEncodedLegSecurityDesc sets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) SetEncodedLegSecurityDesc(v string) {
	m.Set(field.NewEncodedLegSecurityDesc(v))
}

// SetLegRatioQty sets LegRatioQty, Tag 623.
func (m NoLegs) SetLegRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegRatioQty(value, scale))
}

// SetLegSide sets LegSide, Tag 624.
func (m NoLegs) SetLegSide(v string) {
	m.Set(field.NewLegSide(v))
}

// SetLegCurrency sets LegCurrency, Tag 556.
func (m NoLegs) SetLegCurrency(v string) {
	m.Set(field.NewLegCurrency(v))
}

// SetLegPool sets LegPool, Tag 740.
func (m NoLegs) SetLegPool(v string) {
	m.Set(field.NewLegPool(v))
}

// SetLegDatedDate sets LegDatedDate, Tag 739.
func (m NoLegs) SetLegDatedDate(v string) {
	m.Set(field.NewLegDatedDate(v))
}

// SetLegContractSettlMonth sets LegContractSettlMonth, Tag 955.
func (m NoLegs) SetLegContractSettlMonth(v string) {
	m.Set(field.NewLegContractSettlMonth(v))
}

// SetLegInterestAccrualDate sets LegInterestAccrualDate, Tag 956.
func (m NoLegs) SetLegInterestAccrualDate(v string) {
	m.Set(field.NewLegInterestAccrualDate(v))
}

// SetLegUnitOfMeasure sets LegUnitOfMeasure, Tag 999.
func (m NoLegs) SetLegUnitOfMeasure(v string) {
	m.Set(field.NewLegUnitOfMeasure(v))
}

// SetLegTimeUnit sets LegTimeUnit, Tag 1001.
func (m NoLegs) SetLegTimeUnit(v string) {
	m.Set(field.NewLegTimeUnit(v))
}

// SetLegOptionRatio sets LegOptionRatio, Tag 1017.
func (m NoLegs) SetLegOptionRatio(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegOptionRatio(value, scale))
}

// SetLegPrice sets LegPrice, Tag 566.
func (m NoLegs) SetLegPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPrice(value, scale))
}

// SetLegMaturityTime sets LegMaturityTime, Tag 1212.
func (m NoLegs) SetLegMaturityTime(v string) {
	m.Set(field.NewLegMaturityTime(v))
}

// SetLegPutOrCall sets LegPutOrCall, Tag 1358.
func (m NoLegs) SetLegPutOrCall(v int) {
	m.Set(field.NewLegPutOrCall(v))
}

// SetLegExerciseStyle sets LegExerciseStyle, Tag 1420.
func (m NoLegs) SetLegExerciseStyle(v int) {
	m.Set(field.NewLegExerciseStyle(v))
}

// SetLegUnitOfMeasureQty sets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) SetLegUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegUnitOfMeasureQty(value, scale))
}

// SetLegPriceUnitOfMeasure sets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) SetLegPriceUnitOfMeasure(v string) {
	m.Set(field.NewLegPriceUnitOfMeasure(v))
}

// SetLegPriceUnitOfMeasureQty sets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) SetLegPriceUnitOfMeasureQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLegPriceUnitOfMeasureQty(value, scale))
}

// SetLegContractMultiplierUnit sets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) SetLegContractMultiplierUnit(v int) {
	m.Set(field.NewLegContractMultiplierUnit(v))
}

// SetLegFlowScheduleType sets LegFlowScheduleType, Tag 1440.
func (m NoLegs) SetLegFlowScheduleType(v int) {
	m.Set(field.NewLegFlowScheduleType(v))
}

// GetLegSymbol gets LegSymbol, Tag 600.
func (m NoLegs) GetLegSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSymbolSfx gets LegSymbolSfx, Tag 601.
func (m NoLegs) GetLegSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.LegSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityID gets LegSecurityID, Tag 602.
func (m NoLegs) GetLegSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityIDSource gets LegSecurityIDSource, Tag 603.
func (m NoLegs) GetLegSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoLegSecurityAltID gets NoLegSecurityAltID, Tag 604.
func (m NoLegs) GetNoLegSecurityAltID() (NoLegSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLegSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetLegProduct gets LegProduct, Tag 607.
func (m NoLegs) GetLegProduct() (v int, err quickfix.MessageRejectError) {
	var f field.LegProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCFICode gets LegCFICode, Tag 608.
func (m NoLegs) GetLegCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.LegCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityType gets LegSecurityType, Tag 609.
func (m NoLegs) GetLegSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecuritySubType gets LegSecuritySubType, Tag 764.
func (m NoLegs) GetLegSecuritySubType() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecuritySubTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityMonthYear gets LegMaturityMonthYear, Tag 610.
func (m NoLegs) GetLegMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityDate gets LegMaturityDate, Tag 611.
func (m NoLegs) GetLegMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponPaymentDate gets LegCouponPaymentDate, Tag 248.
func (m NoLegs) GetLegCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssueDate gets LegIssueDate, Tag 249.
func (m NoLegs) GetLegIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepoCollateralSecurityType gets LegRepoCollateralSecurityType, Tag 250.
func (m NoLegs) GetLegRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseTerm gets LegRepurchaseTerm, Tag 251.
func (m NoLegs) GetLegRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRepurchaseRate gets LegRepurchaseRate, Tag 252.
func (m NoLegs) GetLegRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFactor gets LegFactor, Tag 253.
func (m NoLegs) GetLegFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCreditRating gets LegCreditRating, Tag 257.
func (m NoLegs) GetLegCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.LegCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInstrRegistry gets LegInstrRegistry, Tag 599.
func (m NoLegs) GetLegInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.LegInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCountryOfIssue gets LegCountryOfIssue, Tag 596.
func (m NoLegs) GetLegCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStateOrProvinceOfIssue gets LegStateOrProvinceOfIssue, Tag 597.
func (m NoLegs) GetLegStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegLocaleOfIssue gets LegLocaleOfIssue, Tag 598.
func (m NoLegs) GetLegLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LegLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRedemptionDate gets LegRedemptionDate, Tag 254.
func (m NoLegs) GetLegRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikePrice gets LegStrikePrice, Tag 612.
func (m NoLegs) GetLegStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegStrikeCurrency gets LegStrikeCurrency, Tag 942.
func (m NoLegs) GetLegStrikeCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegStrikeCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptAttribute gets LegOptAttribute, Tag 613.
func (m NoLegs) GetLegOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.LegOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplier gets LegContractMultiplier, Tag 614.
func (m NoLegs) GetLegContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCouponRate gets LegCouponRate, Tag 615.
func (m NoLegs) GetLegCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityExchange gets LegSecurityExchange, Tag 616.
func (m NoLegs) GetLegSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegIssuer gets LegIssuer, Tag 617.
func (m NoLegs) GetLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.LegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuerLen gets EncodedLegIssuerLen, Tag 618.
func (m NoLegs) GetEncodedLegIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegIssuer gets EncodedLegIssuer, Tag 619.
func (m NoLegs) GetEncodedLegIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityDesc gets LegSecurityDesc, Tag 620.
func (m NoLegs) GetLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDescLen gets EncodedLegSecurityDescLen, Tag 621.
func (m NoLegs) GetEncodedLegSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedLegSecurityDesc gets EncodedLegSecurityDesc, Tag 622.
func (m NoLegs) GetEncodedLegSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedLegSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegRatioQty gets LegRatioQty, Tag 623.
func (m NoLegs) GetLegRatioQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegRatioQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSide gets LegSide, Tag 624.
func (m NoLegs) GetLegSide() (v string, err quickfix.MessageRejectError) {
	var f field.LegSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegCurrency gets LegCurrency, Tag 556.
func (m NoLegs) GetLegCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.LegCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPool gets LegPool, Tag 740.
func (m NoLegs) GetLegPool() (v string, err quickfix.MessageRejectError) {
	var f field.LegPoolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegDatedDate gets LegDatedDate, Tag 739.
func (m NoLegs) GetLegDatedDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegDatedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractSettlMonth gets LegContractSettlMonth, Tag 955.
func (m NoLegs) GetLegContractSettlMonth() (v string, err quickfix.MessageRejectError) {
	var f field.LegContractSettlMonthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegInterestAccrualDate gets LegInterestAccrualDate, Tag 956.
func (m NoLegs) GetLegInterestAccrualDate() (v string, err quickfix.MessageRejectError) {
	var f field.LegInterestAccrualDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasure gets LegUnitOfMeasure, Tag 999.
func (m NoLegs) GetLegUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegTimeUnit gets LegTimeUnit, Tag 1001.
func (m NoLegs) GetLegTimeUnit() (v string, err quickfix.MessageRejectError) {
	var f field.LegTimeUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegOptionRatio gets LegOptionRatio, Tag 1017.
func (m NoLegs) GetLegOptionRatio() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegOptionRatioField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPrice gets LegPrice, Tag 566.
func (m NoLegs) GetLegPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegMaturityTime gets LegMaturityTime, Tag 1212.
func (m NoLegs) GetLegMaturityTime() (v string, err quickfix.MessageRejectError) {
	var f field.LegMaturityTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPutOrCall gets LegPutOrCall, Tag 1358.
func (m NoLegs) GetLegPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.LegPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegExerciseStyle gets LegExerciseStyle, Tag 1420.
func (m NoLegs) GetLegExerciseStyle() (v int, err quickfix.MessageRejectError) {
	var f field.LegExerciseStyleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegUnitOfMeasureQty gets LegUnitOfMeasureQty, Tag 1224.
func (m NoLegs) GetLegUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasure gets LegPriceUnitOfMeasure, Tag 1421.
func (m NoLegs) GetLegPriceUnitOfMeasure() (v string, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegPriceUnitOfMeasureQty gets LegPriceUnitOfMeasureQty, Tag 1422.
func (m NoLegs) GetLegPriceUnitOfMeasureQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LegPriceUnitOfMeasureQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegContractMultiplierUnit gets LegContractMultiplierUnit, Tag 1436.
func (m NoLegs) GetLegContractMultiplierUnit() (v int, err quickfix.MessageRejectError) {
	var f field.LegContractMultiplierUnitField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegFlowScheduleType gets LegFlowScheduleType, Tag 1440.
func (m NoLegs) GetLegFlowScheduleType() (v int, err quickfix.MessageRejectError) {
	var f field.LegFlowScheduleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSymbol returns true if LegSymbol is present, Tag 600.
func (m NoLegs) HasLegSymbol() bool {
	return m.Has(tag.LegSymbol)
}

// HasLegSymbolSfx returns true if LegSymbolSfx is present, Tag 601.
func (m NoLegs) HasLegSymbolSfx() bool {
	return m.Has(tag.LegSymbolSfx)
}

// HasLegSecurityID returns true if LegSecurityID is present, Tag 602.
func (m NoLegs) HasLegSecurityID() bool {
	return m.Has(tag.LegSecurityID)
}

// HasLegSecurityIDSource returns true if LegSecurityIDSource is present, Tag 603.
func (m NoLegs) HasLegSecurityIDSource() bool {
	return m.Has(tag.LegSecurityIDSource)
}

// HasNoLegSecurityAltID returns true if NoLegSecurityAltID is present, Tag 604.
func (m NoLegs) HasNoLegSecurityAltID() bool {
	return m.Has(tag.NoLegSecurityAltID)
}

// HasLegProduct returns true if LegProduct is present, Tag 607.
func (m NoLegs) HasLegProduct() bool {
	return m.Has(tag.LegProduct)
}

// HasLegCFICode returns true if LegCFICode is present, Tag 608.
func (m NoLegs) HasLegCFICode() bool {
	return m.Has(tag.LegCFICode)
}

// HasLegSecurityType returns true if LegSecurityType is present, Tag 609.
func (m NoLegs) HasLegSecurityType() bool {
	return m.Has(tag.LegSecurityType)
}

// HasLegSecuritySubType returns true if LegSecuritySubType is present, Tag 764.
func (m NoLegs) HasLegSecuritySubType() bool {
	return m.Has(tag.LegSecuritySubType)
}

// HasLegMaturityMonthYear returns true if LegMaturityMonthYear is present, Tag 610.
func (m NoLegs) HasLegMaturityMonthYear() bool {
	return m.Has(tag.LegMaturityMonthYear)
}

// HasLegMaturityDate returns true if LegMaturityDate is present, Tag 611.
func (m NoLegs) HasLegMaturityDate() bool {
	return m.Has(tag.LegMaturityDate)
}

// HasLegCouponPaymentDate returns true if LegCouponPaymentDate is present, Tag 248.
func (m NoLegs) HasLegCouponPaymentDate() bool {
	return m.Has(tag.LegCouponPaymentDate)
}

// HasLegIssueDate returns true if LegIssueDate is present, Tag 249.
func (m NoLegs) HasLegIssueDate() bool {
	return m.Has(tag.LegIssueDate)
}

// HasLegRepoCollateralSecurityType returns true if LegRepoCollateralSecurityType is present, Tag 250.
func (m NoLegs) HasLegRepoCollateralSecurityType() bool {
	return m.Has(tag.LegRepoCollateralSecurityType)
}

// HasLegRepurchaseTerm returns true if LegRepurchaseTerm is present, Tag 251.
func (m NoLegs) HasLegRepurchaseTerm() bool {
	return m.Has(tag.LegRepurchaseTerm)
}

// HasLegRepurchaseRate returns true if LegRepurchaseRate is present, Tag 252.
func (m NoLegs) HasLegRepurchaseRate() bool {
	return m.Has(tag.LegRepurchaseRate)
}

// HasLegFactor returns true if LegFactor is present, Tag 253.
func (m NoLegs) HasLegFactor() bool {
	return m.Has(tag.LegFactor)
}

// HasLegCreditRating returns true if LegCreditRating is present, Tag 257.
func (m NoLegs) HasLegCreditRating() bool {
	return m.Has(tag.LegCreditRating)
}

// HasLegInstrRegistry returns true if LegInstrRegistry is present, Tag 599.
func (m NoLegs) HasLegInstrRegistry() bool {
	return m.Has(tag.LegInstrRegistry)
}

// HasLegCountryOfIssue returns true if LegCountryOfIssue is present, Tag 596.
func (m NoLegs) HasLegCountryOfIssue() bool {
	return m.Has(tag.LegCountryOfIssue)
}

// HasLegStateOrProvinceOfIssue returns true if LegStateOrProvinceOfIssue is present, Tag 597.
func (m NoLegs) HasLegStateOrProvinceOfIssue() bool {
	return m.Has(tag.LegStateOrProvinceOfIssue)
}

// HasLegLocaleOfIssue returns true if LegLocaleOfIssue is present, Tag 598.
func (m NoLegs) HasLegLocaleOfIssue() bool {
	return m.Has(tag.LegLocaleOfIssue)
}

// HasLegRedemptionDate returns true if LegRedemptionDate is present, Tag 254.
func (m NoLegs) HasLegRedemptionDate() bool {
	return m.Has(tag.LegRedemptionDate)
}

// HasLegStrikePrice returns true if LegStrikePrice is present, Tag 612.
func (m NoLegs) HasLegStrikePrice() bool {
	return m.Has(tag.LegStrikePrice)
}

// HasLegStrikeCurrency returns true if LegStrikeCurrency is present, Tag 942.
func (m NoLegs) HasLegStrikeCurrency() bool {
	return m.Has(tag.LegStrikeCurrency)
}

// HasLegOptAttribute returns true if LegOptAttribute is present, Tag 613.
func (m NoLegs) HasLegOptAttribute() bool {
	return m.Has(tag.LegOptAttribute)
}

// HasLegContractMultiplier returns true if LegContractMultiplier is present, Tag 614.
func (m NoLegs) HasLegContractMultiplier() bool {
	return m.Has(tag.LegContractMultiplier)
}

// HasLegCouponRate returns true if LegCouponRate is present, Tag 615.
func (m NoLegs) HasLegCouponRate() bool {
	return m.Has(tag.LegCouponRate)
}

// HasLegSecurityExchange returns true if LegSecurityExchange is present, Tag 616.
func (m NoLegs) HasLegSecurityExchange() bool {
	return m.Has(tag.LegSecurityExchange)
}

// HasLegIssuer returns true if LegIssuer is present, Tag 617.
func (m NoLegs) HasLegIssuer() bool {
	return m.Has(tag.LegIssuer)
}

// HasEncodedLegIssuerLen returns true if EncodedLegIssuerLen is present, Tag 618.
func (m NoLegs) HasEncodedLegIssuerLen() bool {
	return m.Has(tag.EncodedLegIssuerLen)
}

// HasEncodedLegIssuer returns true if EncodedLegIssuer is present, Tag 619.
func (m NoLegs) HasEncodedLegIssuer() bool {
	return m.Has(tag.EncodedLegIssuer)
}

// HasLegSecurityDesc returns true if LegSecurityDesc is present, Tag 620.
func (m NoLegs) HasLegSecurityDesc() bool {
	return m.Has(tag.LegSecurityDesc)
}

// HasEncodedLegSecurityDescLen returns true if EncodedLegSecurityDescLen is present, Tag 621.
func (m NoLegs) HasEncodedLegSecurityDescLen() bool {
	return m.Has(tag.EncodedLegSecurityDescLen)
}

// HasEncodedLegSecurityDesc returns true if EncodedLegSecurityDesc is present, Tag 622.
func (m NoLegs) HasEncodedLegSecurityDesc() bool {
	return m.Has(tag.EncodedLegSecurityDesc)
}

// HasLegRatioQty returns true if LegRatioQty is present, Tag 623.
func (m NoLegs) HasLegRatioQty() bool {
	return m.Has(tag.LegRatioQty)
}

// HasLegSide returns true if LegSide is present, Tag 624.
func (m NoLegs) HasLegSide() bool {
	return m.Has(tag.LegSide)
}

// HasLegCurrency returns true if LegCurrency is present, Tag 556.
func (m NoLegs) HasLegCurrency() bool {
	return m.Has(tag.LegCurrency)
}

// HasLegPool returns true if LegPool is present, Tag 740.
func (m NoLegs) HasLegPool() bool {
	return m.Has(tag.LegPool)
}

// HasLegDatedDate returns true if LegDatedDate is present, Tag 739.
func (m NoLegs) HasLegDatedDate() bool {
	return m.Has(tag.LegDatedDate)
}

// HasLegContractSettlMonth returns true if LegContractSettlMonth is present, Tag 955.
func (m NoLegs) HasLegContractSettlMonth() bool {
	return m.Has(tag.LegContractSettlMonth)
}

// HasLegInterestAccrualDate returns true if LegInterestAccrualDate is present, Tag 956.
func (m NoLegs) HasLegInterestAccrualDate() bool {
	return m.Has(tag.LegInterestAccrualDate)
}

// HasLegUnitOfMeasure returns true if LegUnitOfMeasure is present, Tag 999.
func (m NoLegs) HasLegUnitOfMeasure() bool {
	return m.Has(tag.LegUnitOfMeasure)
}

// HasLegTimeUnit returns true if LegTimeUnit is present, Tag 1001.
func (m NoLegs) HasLegTimeUnit() bool {
	return m.Has(tag.LegTimeUnit)
}

// HasLegOptionRatio returns true if LegOptionRatio is present, Tag 1017.
func (m NoLegs) HasLegOptionRatio() bool {
	return m.Has(tag.LegOptionRatio)
}

// HasLegPrice returns true if LegPrice is present, Tag 566.
func (m NoLegs) HasLegPrice() bool {
	return m.Has(tag.LegPrice)
}

// HasLegMaturityTime returns true if LegMaturityTime is present, Tag 1212.
func (m NoLegs) HasLegMaturityTime() bool {
	return m.Has(tag.LegMaturityTime)
}

// HasLegPutOrCall returns true if LegPutOrCall is present, Tag 1358.
func (m NoLegs) HasLegPutOrCall() bool {
	return m.Has(tag.LegPutOrCall)
}

// HasLegExerciseStyle returns true if LegExerciseStyle is present, Tag 1420.
func (m NoLegs) HasLegExerciseStyle() bool {
	return m.Has(tag.LegExerciseStyle)
}

// HasLegUnitOfMeasureQty returns true if LegUnitOfMeasureQty is present, Tag 1224.
func (m NoLegs) HasLegUnitOfMeasureQty() bool {
	return m.Has(tag.LegUnitOfMeasureQty)
}

// HasLegPriceUnitOfMeasure returns true if LegPriceUnitOfMeasure is present, Tag 1421.
func (m NoLegs) HasLegPriceUnitOfMeasure() bool {
	return m.Has(tag.LegPriceUnitOfMeasure)
}

// HasLegPriceUnitOfMeasureQty returns true if LegPriceUnitOfMeasureQty is present, Tag 1422.
func (m NoLegs) HasLegPriceUnitOfMeasureQty() bool {
	return m.Has(tag.LegPriceUnitOfMeasureQty)
}

// HasLegContractMultiplierUnit returns true if LegContractMultiplierUnit is present, Tag 1436.
func (m NoLegs) HasLegContractMultiplierUnit() bool {
	return m.Has(tag.LegContractMultiplierUnit)
}

// HasLegFlowScheduleType returns true if LegFlowScheduleType is present, Tag 1440.
func (m NoLegs) HasLegFlowScheduleType() bool {
	return m.Has(tag.LegFlowScheduleType)
}

// NoLegSecurityAltID is a repeating group element, Tag 604.
type NoLegSecurityAltID struct {
	*quickfix.Group
}

// SetLegSecurityAltID sets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) SetLegSecurityAltID(v string) {
	m.Set(field.NewLegSecurityAltID(v))
}

// SetLegSecurityAltIDSource sets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) SetLegSecurityAltIDSource(v string) {
	m.Set(field.NewLegSecurityAltIDSource(v))
}

// GetLegSecurityAltID gets LegSecurityAltID, Tag 605.
func (m NoLegSecurityAltID) GetLegSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLegSecurityAltIDSource gets LegSecurityAltIDSource, Tag 606.
func (m NoLegSecurityAltID) GetLegSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.LegSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLegSecurityAltID returns true if LegSecurityAltID is present, Tag 605.
func (m NoLegSecurityAltID) HasLegSecurityAltID() bool {
	return m.Has(tag.LegSecurityAltID)
}

// HasLegSecurityAltIDSource returns true if LegSecurityAltIDSource is present, Tag 606.
func (m NoLegSecurityAltID) HasLegSecurityAltIDSource() bool {
	return m.Has(tag.LegSecurityAltIDSource)
}

// NoLegSecurityAltIDRepeatingGroup is a repeating group, Tag 604.
type NoLegSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegSecurityAltIDRepeatingGroup returns an initialized, NoLegSecurityAltIDRepeatingGroup.
func NewNoLegSecurityAltIDRepeatingGroup() NoLegSecurityAltIDRepeatingGroup {
	return NoLegSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSecurityAltID), quickfix.GroupElement(tag.LegSecurityAltIDSource)})}
}

// Add create and append a new NoLegSecurityAltID to this group.
func (m NoLegSecurityAltIDRepeatingGroup) Add() NoLegSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoLegSecurityAltID{g}
}

// Get returns the ith NoLegSecurityAltID in the NoLegSecurityAltIDRepeatinGroup.
func (m NoLegSecurityAltIDRepeatingGroup) Get(i int) NoLegSecurityAltID {
	return NoLegSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoLegsRepeatingGroup is a repeating group, Tag 555.
type NoLegsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoLegsRepeatingGroup returns an initialized, NoLegsRepeatingGroup.
func NewNoLegsRepeatingGroup() NoLegsRepeatingGroup {
	return NoLegsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLegs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LegSymbol), quickfix.GroupElement(tag.LegSymbolSfx), quickfix.GroupElement(tag.LegSecurityID), quickfix.GroupElement(tag.LegSecurityIDSource), NewNoLegSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.LegProduct), quickfix.GroupElement(tag.LegCFICode), quickfix.GroupElement(tag.LegSecurityType), quickfix.GroupElement(tag.LegSecuritySubType), quickfix.GroupElement(tag.LegMaturityMonthYear), quickfix.GroupElement(tag.LegMaturityDate), quickfix.GroupElement(tag.LegCouponPaymentDate), quickfix.GroupElement(tag.LegIssueDate), quickfix.GroupElement(tag.LegRepoCollateralSecurityType), quickfix.GroupElement(tag.LegRepurchaseTerm), quickfix.GroupElement(tag.LegRepurchaseRate), quickfix.GroupElement(tag.LegFactor), quickfix.GroupElement(tag.LegCreditRating), quickfix.GroupElement(tag.LegInstrRegistry), quickfix.GroupElement(tag.LegCountryOfIssue), quickfix.GroupElement(tag.LegStateOrProvinceOfIssue), quickfix.GroupElement(tag.LegLocaleOfIssue), quickfix.GroupElement(tag.LegRedemptionDate), quickfix.GroupElement(tag.LegStrikePrice), quickfix.GroupElement(tag.LegStrikeCurrency), quickfix.GroupElement(tag.LegOptAttribute), quickfix.GroupElement(tag.LegContractMultiplier), quickfix.GroupElement(tag.LegCouponRate), quickfix.GroupElement(tag.LegSecurityExchange), quickfix.GroupElement(tag.LegIssuer), quickfix.GroupElement(tag.EncodedLegIssuerLen), quickfix.GroupElement(tag.EncodedLegIssuer), quickfix.GroupElement(tag.LegSecurityDesc), quickfix.GroupElement(tag.EncodedLegSecurityDescLen), quickfix.GroupElement(tag.EncodedLegSecurityDesc), quickfix.GroupElement(tag.LegRatioQty), quickfix.GroupElement(tag.LegSide), quickfix.GroupElement(tag.LegCurrency), quickfix.GroupElement(tag.LegPool), quickfix.GroupElement(tag.LegDatedDate), quickfix.GroupElement(tag.LegContractSettlMonth), quickfix.GroupElement(tag.LegInterestAccrualDate), quickfix.GroupElement(tag.LegUnitOfMeasure), quickfix.GroupElement(tag.LegTimeUnit), quickfix.GroupElement(tag.LegOptionRatio), quickfix.GroupElement(tag.LegPrice), quickfix.GroupElement(tag.LegMaturityTime), quickfix.GroupElement(tag.LegPutOrCall), quickfix.GroupElement(tag.LegExerciseStyle), quickfix.GroupElement(tag.LegUnitOfMeasureQty), quickfix.GroupElement(tag.LegPriceUnitOfMeasure), quickfix.GroupElement(tag.LegPriceUnitOfMeasureQty), quickfix.GroupElement(tag.LegContractMultiplierUnit), quickfix.GroupElement(tag.LegFlowScheduleType)})}
}

// Add create and append a new NoLegs to this group.
func (m NoLegsRepeatingGroup) Add() NoLegs {
	g := m.RepeatingGroup.Add()
	return NoLegs{g}
}

// Get returns the ith NoLegs in the NoLegsRepeatinGroup.
func (m NoLegsRepeatingGroup) Get(i int) NoLegs {
	return NoLegs{m.RepeatingGroup.Get(i)}
}

// NoQuoteEntriesRepeatingGroup is a repeating group, Tag 295.
type NoQuoteEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoQuoteEntriesRepeatingGroup returns an initialized, NoQuoteEntriesRepeatingGroup.
func NewNoQuoteEntriesRepeatingGroup() NoQuoteEntriesRepeatingGroup {
	return NoQuoteEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoQuoteEntries,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), NewNoEventsRepeatingGroup(), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.SecurityStatus), quickfix.GroupElement(tag.SettleOnOpenFlag), quickfix.GroupElement(tag.InstrmtAssignmentMethod), quickfix.GroupElement(tag.StrikeMultiplier), quickfix.GroupElement(tag.StrikeValue), quickfix.GroupElement(tag.MinPriceIncrement), quickfix.GroupElement(tag.PositionLimit), quickfix.GroupElement(tag.NTPositionLimit), NewNoInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnitOfMeasure), quickfix.GroupElement(tag.TimeUnit), quickfix.GroupElement(tag.MaturityTime), quickfix.GroupElement(tag.SecurityGroup), quickfix.GroupElement(tag.MinPriceIncrementAmount), quickfix.GroupElement(tag.UnitOfMeasureQty), quickfix.GroupElement(tag.SecurityXMLLen), quickfix.GroupElement(tag.SecurityXML), quickfix.GroupElement(tag.SecurityXMLSchema), quickfix.GroupElement(tag.ProductComplex), quickfix.GroupElement(tag.PriceUnitOfMeasure), quickfix.GroupElement(tag.PriceUnitOfMeasureQty), quickfix.GroupElement(tag.SettlMethod), quickfix.GroupElement(tag.ExerciseStyle), quickfix.GroupElement(tag.OptPayoutAmount), quickfix.GroupElement(tag.PriceQuoteMethod), quickfix.GroupElement(tag.ListMethod), quickfix.GroupElement(tag.CapPrice), quickfix.GroupElement(tag.FloorPrice), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.FlexibleIndicator), quickfix.GroupElement(tag.FlexProductEligibilityIndicator), quickfix.GroupElement(tag.ValuationMethod), quickfix.GroupElement(tag.ContractMultiplierUnit), quickfix.GroupElement(tag.FlowScheduleType), quickfix.GroupElement(tag.RestructuringType), quickfix.GroupElement(tag.Seniority), quickfix.GroupElement(tag.NotionalPercentageOutstanding), quickfix.GroupElement(tag.OriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.AttachmentPoint), quickfix.GroupElement(tag.DetachmentPoint), quickfix.GroupElement(tag.StrikePriceDeterminationMethod), quickfix.GroupElement(tag.StrikePriceBoundaryMethod), quickfix.GroupElement(tag.StrikePriceBoundaryPrecision), quickfix.GroupElement(tag.UnderlyingPriceDeterminationMethod), quickfix.GroupElement(tag.OptPayoutType), NewNoComplexEventsRepeatingGroup(), quickfix.GroupElement(tag.AgreementDesc), quickfix.GroupElement(tag.AgreementID), quickfix.GroupElement(tag.AgreementDate), quickfix.GroupElement(tag.AgreementCurrency), quickfix.GroupElement(tag.TerminationType), quickfix.GroupElement(tag.StartDate), quickfix.GroupElement(tag.EndDate), quickfix.GroupElement(tag.DeliveryType), quickfix.GroupElement(tag.MarginRatio), NewNoUnderlyingsRepeatingGroup(), NewNoLegsRepeatingGroup()})}
}

// Add create and append a new NoQuoteEntries to this group.
func (m NoQuoteEntriesRepeatingGroup) Add() NoQuoteEntries {
	g := m.RepeatingGroup.Add()
	return NoQuoteEntries{g}
}

// Get returns the ith NoQuoteEntries in the NoQuoteEntriesRepeatinGroup.
func (m NoQuoteEntriesRepeatingGroup) Get(i int) NoQuoteEntries {
	return NoQuoteEntries{m.RepeatingGroup.Get(i)}
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetNoPartySubIDs sets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) SetNoPartySubIDs(f NoPartySubIDsRepeatingGroup) {
	m.SetGroup(f)
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartySubIDs gets NoPartySubIDs, Tag 802.
func (m NoPartyIDs) GetNoPartySubIDs() (NoPartySubIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartySubIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasNoPartySubIDs returns true if NoPartySubIDs is present, Tag 802.
func (m NoPartyIDs) HasNoPartySubIDs() bool {
	return m.Has(tag.NoPartySubIDs)
}

// NoPartySubIDs is a repeating group element, Tag 802.
type NoPartySubIDs struct {
	*quickfix.Group
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartySubIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// SetPartySubIDType sets PartySubIDType, Tag 803.
func (m NoPartySubIDs) SetPartySubIDType(v enum.PartySubIDType) {
	m.Set(field.NewPartySubIDType(v))
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartySubIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubIDType gets PartySubIDType, Tag 803.
func (m NoPartySubIDs) GetPartySubIDType() (v enum.PartySubIDType, err quickfix.MessageRejectError) {
	var f field.PartySubIDTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartySubIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// HasPartySubIDType returns true if PartySubIDType is present, Tag 803.
func (m NoPartySubIDs) HasPartySubIDType() bool {
	return m.Has(tag.PartySubIDType)
}

// NoPartySubIDsRepeatingGroup is a repeating group, Tag 802.
type NoPartySubIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartySubIDsRepeatingGroup returns an initialized, NoPartySubIDsRepeatingGroup.
func NewNoPartySubIDsRepeatingGroup() NoPartySubIDsRepeatingGroup {
	return NoPartySubIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartySubIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartySubID), quickfix.GroupElement(tag.PartySubIDType)})}
}

// Add create and append a new NoPartySubIDs to this group.
func (m NoPartySubIDsRepeatingGroup) Add() NoPartySubIDs {
	g := m.RepeatingGroup.Add()
	return NoPartySubIDs{g}
}

// Get returns the ith NoPartySubIDs in the NoPartySubIDsRepeatinGroup.
func (m NoPartySubIDsRepeatingGroup) Get(i int) NoPartySubIDs {
	return NoPartySubIDs{m.RepeatingGroup.Get(i)}
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), NewNoPartySubIDsRepeatingGroup()})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoTargetPartyIDs is a repeating group element, Tag 1461.
type NoTargetPartyIDs struct {
	*quickfix.Group
}

// SetTargetPartyID sets TargetPartyID, Tag 1462.
func (m NoTargetPartyIDs) SetTargetPartyID(v string) {
	m.Set(field.NewTargetPartyID(v))
}

// SetTargetPartyIDSource sets TargetPartyIDSource, Tag 1463.
func (m NoTargetPartyIDs) SetTargetPartyIDSource(v string) {
	m.Set(field.NewTargetPartyIDSource(v))
}

// SetTargetPartyRole sets TargetPartyRole, Tag 1464.
func (m NoTargetPartyIDs) SetTargetPartyRole(v int) {
	m.Set(field.NewTargetPartyRole(v))
}

// GetTargetPartyID gets TargetPartyID, Tag 1462.
func (m NoTargetPartyIDs) GetTargetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.TargetPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartyIDSource gets TargetPartyIDSource, Tag 1463.
func (m NoTargetPartyIDs) GetTargetPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.TargetPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTargetPartyRole gets TargetPartyRole, Tag 1464.
func (m NoTargetPartyIDs) GetTargetPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.TargetPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasTargetPartyID returns true if TargetPartyID is present, Tag 1462.
func (m NoTargetPartyIDs) HasTargetPartyID() bool {
	return m.Has(tag.TargetPartyID)
}

// HasTargetPartyIDSource returns true if TargetPartyIDSource is present, Tag 1463.
func (m NoTargetPartyIDs) HasTargetPartyIDSource() bool {
	return m.Has(tag.TargetPartyIDSource)
}

// HasTargetPartyRole returns true if TargetPartyRole is present, Tag 1464.
func (m NoTargetPartyIDs) HasTargetPartyRole() bool {
	return m.Has(tag.TargetPartyRole)
}

// NoTargetPartyIDsRepeatingGroup is a repeating group, Tag 1461.
type NoTargetPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoTargetPartyIDsRepeatingGroup returns an initialized, NoTargetPartyIDsRepeatingGroup.
func NewNoTargetPartyIDsRepeatingGroup() NoTargetPartyIDsRepeatingGroup {
	return NoTargetPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTargetPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TargetPartyID), quickfix.GroupElement(tag.TargetPartyIDSource), quickfix.GroupElement(tag.TargetPartyRole)})}
}

// Add create and append a new NoTargetPartyIDs to this group.
func (m NoTargetPartyIDsRepeatingGroup) Add() NoTargetPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoTargetPartyIDs{g}
}

// Get returns the ith NoTargetPartyIDs in the NoTargetPartyIDsRepeatinGroup.
func (m NoTargetPartyIDsRepeatingGroup) Get(i int) NoTargetPartyIDs {
	return NoTargetPartyIDs{m.RepeatingGroup.Get(i)}
}
