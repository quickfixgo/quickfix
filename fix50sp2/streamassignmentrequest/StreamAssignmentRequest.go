package streamassignmentrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//StreamAssignmentRequest is the fix50sp2 StreamAssignmentRequest type, MsgType = CC
type StreamAssignmentRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a StreamAssignmentRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) StreamAssignmentRequest {
	return StreamAssignmentRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m StreamAssignmentRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a StreamAssignmentRequest initialized with the required fields for StreamAssignmentRequest
func New(streamasgnreqid field.StreamAsgnReqIDField, streamasgnreqtype field.StreamAsgnReqTypeField) (m StreamAssignmentRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("CC"))
	m.Set(streamasgnreqid)
	m.Set(streamasgnreqtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg StreamAssignmentRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "CC", r
}

//SetStreamAsgnReqID sets StreamAsgnReqID, Tag 1497
func (m StreamAssignmentRequest) SetStreamAsgnReqID(v string) {
	m.Set(field.NewStreamAsgnReqID(v))
}

//SetStreamAsgnReqType sets StreamAsgnReqType, Tag 1498
func (m StreamAssignmentRequest) SetStreamAsgnReqType(v int) {
	m.Set(field.NewStreamAsgnReqType(v))
}

//SetNoAsgnReqs sets NoAsgnReqs, Tag 1499
func (m StreamAssignmentRequest) SetNoAsgnReqs(f NoAsgnReqsRepeatingGroup) {
	m.SetGroup(f)
}

//GetStreamAsgnReqID gets StreamAsgnReqID, Tag 1497
func (m StreamAssignmentRequest) GetStreamAsgnReqID() (f field.StreamAsgnReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStreamAsgnReqType gets StreamAsgnReqType, Tag 1498
func (m StreamAssignmentRequest) GetStreamAsgnReqType() (f field.StreamAsgnReqTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAsgnReqs gets NoAsgnReqs, Tag 1499
func (m StreamAssignmentRequest) GetNoAsgnReqs() (NoAsgnReqsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAsgnReqsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasStreamAsgnReqID returns true if StreamAsgnReqID is present, Tag 1497
func (m StreamAssignmentRequest) HasStreamAsgnReqID() bool {
	return m.Has(tag.StreamAsgnReqID)
}

//HasStreamAsgnReqType returns true if StreamAsgnReqType is present, Tag 1498
func (m StreamAssignmentRequest) HasStreamAsgnReqType() bool {
	return m.Has(tag.StreamAsgnReqType)
}

//HasNoAsgnReqs returns true if NoAsgnReqs is present, Tag 1499
func (m StreamAssignmentRequest) HasNoAsgnReqs() bool {
	return m.Has(tag.NoAsgnReqs)
}

//NoAsgnReqs is a repeating group element, Tag 1499
type NoAsgnReqs struct {
	quickfix.Group
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m NoAsgnReqs) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m NoAsgnReqs) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m NoAsgnReqs) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m NoAsgnReqs) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m NoAsgnReqs) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m NoAsgnReqs) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
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

//SetSecuritySubType sets SecuritySubType, Tag 762
func (m NoRelatedSym) SetSecuritySubType(v string) {
	m.Set(field.NewSecuritySubType(v))
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

//SetStrikeCurrency sets StrikeCurrency, Tag 947
func (m NoRelatedSym) SetStrikeCurrency(v string) {
	m.Set(field.NewStrikeCurrency(v))
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

//SetPool sets Pool, Tag 691
func (m NoRelatedSym) SetPool(v string) {
	m.Set(field.NewPool(v))
}

//SetContractSettlMonth sets ContractSettlMonth, Tag 667
func (m NoRelatedSym) SetContractSettlMonth(v string) {
	m.Set(field.NewContractSettlMonth(v))
}

//SetCPProgram sets CPProgram, Tag 875
func (m NoRelatedSym) SetCPProgram(v int) {
	m.Set(field.NewCPProgram(v))
}

//SetCPRegType sets CPRegType, Tag 876
func (m NoRelatedSym) SetCPRegType(v string) {
	m.Set(field.NewCPRegType(v))
}

//SetNoEvents sets NoEvents, Tag 864
func (m NoRelatedSym) SetNoEvents(f NoEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetDatedDate sets DatedDate, Tag 873
func (m NoRelatedSym) SetDatedDate(v string) {
	m.Set(field.NewDatedDate(v))
}

//SetInterestAccrualDate sets InterestAccrualDate, Tag 874
func (m NoRelatedSym) SetInterestAccrualDate(v string) {
	m.Set(field.NewInterestAccrualDate(v))
}

//SetSecurityStatus sets SecurityStatus, Tag 965
func (m NoRelatedSym) SetSecurityStatus(v string) {
	m.Set(field.NewSecurityStatus(v))
}

//SetSettleOnOpenFlag sets SettleOnOpenFlag, Tag 966
func (m NoRelatedSym) SetSettleOnOpenFlag(v string) {
	m.Set(field.NewSettleOnOpenFlag(v))
}

//SetInstrmtAssignmentMethod sets InstrmtAssignmentMethod, Tag 1049
func (m NoRelatedSym) SetInstrmtAssignmentMethod(v string) {
	m.Set(field.NewInstrmtAssignmentMethod(v))
}

//SetStrikeMultiplier sets StrikeMultiplier, Tag 967
func (m NoRelatedSym) SetStrikeMultiplier(v float64) {
	m.Set(field.NewStrikeMultiplier(v))
}

//SetStrikeValue sets StrikeValue, Tag 968
func (m NoRelatedSym) SetStrikeValue(v float64) {
	m.Set(field.NewStrikeValue(v))
}

//SetMinPriceIncrement sets MinPriceIncrement, Tag 969
func (m NoRelatedSym) SetMinPriceIncrement(v float64) {
	m.Set(field.NewMinPriceIncrement(v))
}

//SetPositionLimit sets PositionLimit, Tag 970
func (m NoRelatedSym) SetPositionLimit(v int) {
	m.Set(field.NewPositionLimit(v))
}

//SetNTPositionLimit sets NTPositionLimit, Tag 971
func (m NoRelatedSym) SetNTPositionLimit(v int) {
	m.Set(field.NewNTPositionLimit(v))
}

//SetNoInstrumentParties sets NoInstrumentParties, Tag 1018
func (m NoRelatedSym) SetNoInstrumentParties(f NoInstrumentPartiesRepeatingGroup) {
	m.SetGroup(f)
}

//SetUnitOfMeasure sets UnitOfMeasure, Tag 996
func (m NoRelatedSym) SetUnitOfMeasure(v string) {
	m.Set(field.NewUnitOfMeasure(v))
}

//SetTimeUnit sets TimeUnit, Tag 997
func (m NoRelatedSym) SetTimeUnit(v string) {
	m.Set(field.NewTimeUnit(v))
}

//SetMaturityTime sets MaturityTime, Tag 1079
func (m NoRelatedSym) SetMaturityTime(v string) {
	m.Set(field.NewMaturityTime(v))
}

//SetSecurityGroup sets SecurityGroup, Tag 1151
func (m NoRelatedSym) SetSecurityGroup(v string) {
	m.Set(field.NewSecurityGroup(v))
}

//SetMinPriceIncrementAmount sets MinPriceIncrementAmount, Tag 1146
func (m NoRelatedSym) SetMinPriceIncrementAmount(v float64) {
	m.Set(field.NewMinPriceIncrementAmount(v))
}

//SetUnitOfMeasureQty sets UnitOfMeasureQty, Tag 1147
func (m NoRelatedSym) SetUnitOfMeasureQty(v float64) {
	m.Set(field.NewUnitOfMeasureQty(v))
}

//SetSecurityXMLLen sets SecurityXMLLen, Tag 1184
func (m NoRelatedSym) SetSecurityXMLLen(v int) {
	m.Set(field.NewSecurityXMLLen(v))
}

//SetSecurityXML sets SecurityXML, Tag 1185
func (m NoRelatedSym) SetSecurityXML(v string) {
	m.Set(field.NewSecurityXML(v))
}

//SetSecurityXMLSchema sets SecurityXMLSchema, Tag 1186
func (m NoRelatedSym) SetSecurityXMLSchema(v string) {
	m.Set(field.NewSecurityXMLSchema(v))
}

//SetProductComplex sets ProductComplex, Tag 1227
func (m NoRelatedSym) SetProductComplex(v string) {
	m.Set(field.NewProductComplex(v))
}

//SetPriceUnitOfMeasure sets PriceUnitOfMeasure, Tag 1191
func (m NoRelatedSym) SetPriceUnitOfMeasure(v string) {
	m.Set(field.NewPriceUnitOfMeasure(v))
}

//SetPriceUnitOfMeasureQty sets PriceUnitOfMeasureQty, Tag 1192
func (m NoRelatedSym) SetPriceUnitOfMeasureQty(v float64) {
	m.Set(field.NewPriceUnitOfMeasureQty(v))
}

//SetSettlMethod sets SettlMethod, Tag 1193
func (m NoRelatedSym) SetSettlMethod(v string) {
	m.Set(field.NewSettlMethod(v))
}

//SetExerciseStyle sets ExerciseStyle, Tag 1194
func (m NoRelatedSym) SetExerciseStyle(v int) {
	m.Set(field.NewExerciseStyle(v))
}

//SetOptPayoutAmount sets OptPayoutAmount, Tag 1195
func (m NoRelatedSym) SetOptPayoutAmount(v float64) {
	m.Set(field.NewOptPayoutAmount(v))
}

//SetPriceQuoteMethod sets PriceQuoteMethod, Tag 1196
func (m NoRelatedSym) SetPriceQuoteMethod(v string) {
	m.Set(field.NewPriceQuoteMethod(v))
}

//SetListMethod sets ListMethod, Tag 1198
func (m NoRelatedSym) SetListMethod(v int) {
	m.Set(field.NewListMethod(v))
}

//SetCapPrice sets CapPrice, Tag 1199
func (m NoRelatedSym) SetCapPrice(v float64) {
	m.Set(field.NewCapPrice(v))
}

//SetFloorPrice sets FloorPrice, Tag 1200
func (m NoRelatedSym) SetFloorPrice(v float64) {
	m.Set(field.NewFloorPrice(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoRelatedSym) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetFlexibleIndicator sets FlexibleIndicator, Tag 1244
func (m NoRelatedSym) SetFlexibleIndicator(v bool) {
	m.Set(field.NewFlexibleIndicator(v))
}

//SetFlexProductEligibilityIndicator sets FlexProductEligibilityIndicator, Tag 1242
func (m NoRelatedSym) SetFlexProductEligibilityIndicator(v bool) {
	m.Set(field.NewFlexProductEligibilityIndicator(v))
}

//SetValuationMethod sets ValuationMethod, Tag 1197
func (m NoRelatedSym) SetValuationMethod(v string) {
	m.Set(field.NewValuationMethod(v))
}

//SetContractMultiplierUnit sets ContractMultiplierUnit, Tag 1435
func (m NoRelatedSym) SetContractMultiplierUnit(v int) {
	m.Set(field.NewContractMultiplierUnit(v))
}

//SetFlowScheduleType sets FlowScheduleType, Tag 1439
func (m NoRelatedSym) SetFlowScheduleType(v int) {
	m.Set(field.NewFlowScheduleType(v))
}

//SetRestructuringType sets RestructuringType, Tag 1449
func (m NoRelatedSym) SetRestructuringType(v string) {
	m.Set(field.NewRestructuringType(v))
}

//SetSeniority sets Seniority, Tag 1450
func (m NoRelatedSym) SetSeniority(v string) {
	m.Set(field.NewSeniority(v))
}

//SetNotionalPercentageOutstanding sets NotionalPercentageOutstanding, Tag 1451
func (m NoRelatedSym) SetNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewNotionalPercentageOutstanding(v))
}

//SetOriginalNotionalPercentageOutstanding sets OriginalNotionalPercentageOutstanding, Tag 1452
func (m NoRelatedSym) SetOriginalNotionalPercentageOutstanding(v float64) {
	m.Set(field.NewOriginalNotionalPercentageOutstanding(v))
}

//SetAttachmentPoint sets AttachmentPoint, Tag 1457
func (m NoRelatedSym) SetAttachmentPoint(v float64) {
	m.Set(field.NewAttachmentPoint(v))
}

//SetDetachmentPoint sets DetachmentPoint, Tag 1458
func (m NoRelatedSym) SetDetachmentPoint(v float64) {
	m.Set(field.NewDetachmentPoint(v))
}

//SetStrikePriceDeterminationMethod sets StrikePriceDeterminationMethod, Tag 1478
func (m NoRelatedSym) SetStrikePriceDeterminationMethod(v int) {
	m.Set(field.NewStrikePriceDeterminationMethod(v))
}

//SetStrikePriceBoundaryMethod sets StrikePriceBoundaryMethod, Tag 1479
func (m NoRelatedSym) SetStrikePriceBoundaryMethod(v int) {
	m.Set(field.NewStrikePriceBoundaryMethod(v))
}

//SetStrikePriceBoundaryPrecision sets StrikePriceBoundaryPrecision, Tag 1480
func (m NoRelatedSym) SetStrikePriceBoundaryPrecision(v float64) {
	m.Set(field.NewStrikePriceBoundaryPrecision(v))
}

//SetUnderlyingPriceDeterminationMethod sets UnderlyingPriceDeterminationMethod, Tag 1481
func (m NoRelatedSym) SetUnderlyingPriceDeterminationMethod(v int) {
	m.Set(field.NewUnderlyingPriceDeterminationMethod(v))
}

//SetOptPayoutType sets OptPayoutType, Tag 1482
func (m NoRelatedSym) SetOptPayoutType(v int) {
	m.Set(field.NewOptPayoutType(v))
}

//SetNoComplexEvents sets NoComplexEvents, Tag 1483
func (m NoRelatedSym) SetNoComplexEvents(f NoComplexEventsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSettlType sets SettlType, Tag 63
func (m NoRelatedSym) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetMDEntrySize sets MDEntrySize, Tag 271
func (m NoRelatedSym) SetMDEntrySize(v float64) {
	m.Set(field.NewMDEntrySize(v))
}

//SetMDStreamID sets MDStreamID, Tag 1500
func (m NoRelatedSym) SetMDStreamID(v string) {
	m.Set(field.NewMDStreamID(v))
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

//GetSecuritySubType gets SecuritySubType, Tag 762
func (m NoRelatedSym) GetSecuritySubType() (f field.SecuritySubTypeField, err quickfix.MessageRejectError) {
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

//GetStrikeCurrency gets StrikeCurrency, Tag 947
func (m NoRelatedSym) GetStrikeCurrency() (f field.StrikeCurrencyField, err quickfix.MessageRejectError) {
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

//GetPool gets Pool, Tag 691
func (m NoRelatedSym) GetPool() (f field.PoolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractSettlMonth gets ContractSettlMonth, Tag 667
func (m NoRelatedSym) GetContractSettlMonth() (f field.ContractSettlMonthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPProgram gets CPProgram, Tag 875
func (m NoRelatedSym) GetCPProgram() (f field.CPProgramField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCPRegType gets CPRegType, Tag 876
func (m NoRelatedSym) GetCPRegType() (f field.CPRegTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoEvents gets NoEvents, Tag 864
func (m NoRelatedSym) GetNoEvents() (NoEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDatedDate gets DatedDate, Tag 873
func (m NoRelatedSym) GetDatedDate() (f field.DatedDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInterestAccrualDate gets InterestAccrualDate, Tag 874
func (m NoRelatedSym) GetInterestAccrualDate() (f field.InterestAccrualDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatus gets SecurityStatus, Tag 965
func (m NoRelatedSym) GetSecurityStatus() (f field.SecurityStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettleOnOpenFlag gets SettleOnOpenFlag, Tag 966
func (m NoRelatedSym) GetSettleOnOpenFlag() (f field.SettleOnOpenFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrmtAssignmentMethod gets InstrmtAssignmentMethod, Tag 1049
func (m NoRelatedSym) GetInstrmtAssignmentMethod() (f field.InstrmtAssignmentMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeMultiplier gets StrikeMultiplier, Tag 967
func (m NoRelatedSym) GetStrikeMultiplier() (f field.StrikeMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikeValue gets StrikeValue, Tag 968
func (m NoRelatedSym) GetStrikeValue() (f field.StrikeValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrement gets MinPriceIncrement, Tag 969
func (m NoRelatedSym) GetMinPriceIncrement() (f field.MinPriceIncrementField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPositionLimit gets PositionLimit, Tag 970
func (m NoRelatedSym) GetPositionLimit() (f field.PositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNTPositionLimit gets NTPositionLimit, Tag 971
func (m NoRelatedSym) GetNTPositionLimit() (f field.NTPositionLimitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoInstrumentParties gets NoInstrumentParties, Tag 1018
func (m NoRelatedSym) GetNoInstrumentParties() (NoInstrumentPartiesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoInstrumentPartiesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetUnitOfMeasure gets UnitOfMeasure, Tag 996
func (m NoRelatedSym) GetUnitOfMeasure() (f field.UnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeUnit gets TimeUnit, Tag 997
func (m NoRelatedSym) GetTimeUnit() (f field.TimeUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityTime gets MaturityTime, Tag 1079
func (m NoRelatedSym) GetMaturityTime() (f field.MaturityTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityGroup gets SecurityGroup, Tag 1151
func (m NoRelatedSym) GetSecurityGroup() (f field.SecurityGroupField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinPriceIncrementAmount gets MinPriceIncrementAmount, Tag 1146
func (m NoRelatedSym) GetMinPriceIncrementAmount() (f field.MinPriceIncrementAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnitOfMeasureQty gets UnitOfMeasureQty, Tag 1147
func (m NoRelatedSym) GetUnitOfMeasureQty() (f field.UnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLLen gets SecurityXMLLen, Tag 1184
func (m NoRelatedSym) GetSecurityXMLLen() (f field.SecurityXMLLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXML gets SecurityXML, Tag 1185
func (m NoRelatedSym) GetSecurityXML() (f field.SecurityXMLField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityXMLSchema gets SecurityXMLSchema, Tag 1186
func (m NoRelatedSym) GetSecurityXMLSchema() (f field.SecurityXMLSchemaField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProductComplex gets ProductComplex, Tag 1227
func (m NoRelatedSym) GetProductComplex() (f field.ProductComplexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasure gets PriceUnitOfMeasure, Tag 1191
func (m NoRelatedSym) GetPriceUnitOfMeasure() (f field.PriceUnitOfMeasureField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceUnitOfMeasureQty gets PriceUnitOfMeasureQty, Tag 1192
func (m NoRelatedSym) GetPriceUnitOfMeasureQty() (f field.PriceUnitOfMeasureQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlMethod gets SettlMethod, Tag 1193
func (m NoRelatedSym) GetSettlMethod() (f field.SettlMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExerciseStyle gets ExerciseStyle, Tag 1194
func (m NoRelatedSym) GetExerciseStyle() (f field.ExerciseStyleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutAmount gets OptPayoutAmount, Tag 1195
func (m NoRelatedSym) GetOptPayoutAmount() (f field.OptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceQuoteMethod gets PriceQuoteMethod, Tag 1196
func (m NoRelatedSym) GetPriceQuoteMethod() (f field.PriceQuoteMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListMethod gets ListMethod, Tag 1198
func (m NoRelatedSym) GetListMethod() (f field.ListMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCapPrice gets CapPrice, Tag 1199
func (m NoRelatedSym) GetCapPrice() (f field.CapPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFloorPrice gets FloorPrice, Tag 1200
func (m NoRelatedSym) GetFloorPrice() (f field.FloorPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoRelatedSym) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexibleIndicator gets FlexibleIndicator, Tag 1244
func (m NoRelatedSym) GetFlexibleIndicator() (f field.FlexibleIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlexProductEligibilityIndicator gets FlexProductEligibilityIndicator, Tag 1242
func (m NoRelatedSym) GetFlexProductEligibilityIndicator() (f field.FlexProductEligibilityIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValuationMethod gets ValuationMethod, Tag 1197
func (m NoRelatedSym) GetValuationMethod() (f field.ValuationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplierUnit gets ContractMultiplierUnit, Tag 1435
func (m NoRelatedSym) GetContractMultiplierUnit() (f field.ContractMultiplierUnitField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFlowScheduleType gets FlowScheduleType, Tag 1439
func (m NoRelatedSym) GetFlowScheduleType() (f field.FlowScheduleTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRestructuringType gets RestructuringType, Tag 1449
func (m NoRelatedSym) GetRestructuringType() (f field.RestructuringTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSeniority gets Seniority, Tag 1450
func (m NoRelatedSym) GetSeniority() (f field.SeniorityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNotionalPercentageOutstanding gets NotionalPercentageOutstanding, Tag 1451
func (m NoRelatedSym) GetNotionalPercentageOutstanding() (f field.NotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOriginalNotionalPercentageOutstanding gets OriginalNotionalPercentageOutstanding, Tag 1452
func (m NoRelatedSym) GetOriginalNotionalPercentageOutstanding() (f field.OriginalNotionalPercentageOutstandingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAttachmentPoint gets AttachmentPoint, Tag 1457
func (m NoRelatedSym) GetAttachmentPoint() (f field.AttachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDetachmentPoint gets DetachmentPoint, Tag 1458
func (m NoRelatedSym) GetDetachmentPoint() (f field.DetachmentPointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceDeterminationMethod gets StrikePriceDeterminationMethod, Tag 1478
func (m NoRelatedSym) GetStrikePriceDeterminationMethod() (f field.StrikePriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryMethod gets StrikePriceBoundaryMethod, Tag 1479
func (m NoRelatedSym) GetStrikePriceBoundaryMethod() (f field.StrikePriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePriceBoundaryPrecision gets StrikePriceBoundaryPrecision, Tag 1480
func (m NoRelatedSym) GetStrikePriceBoundaryPrecision() (f field.StrikePriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPriceDeterminationMethod gets UnderlyingPriceDeterminationMethod, Tag 1481
func (m NoRelatedSym) GetUnderlyingPriceDeterminationMethod() (f field.UnderlyingPriceDeterminationMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptPayoutType gets OptPayoutType, Tag 1482
func (m NoRelatedSym) GetOptPayoutType() (f field.OptPayoutTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEvents gets NoComplexEvents, Tag 1483
func (m NoRelatedSym) GetNoComplexEvents() (NoComplexEventsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSettlType gets SettlType, Tag 63
func (m NoRelatedSym) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntrySize gets MDEntrySize, Tag 271
func (m NoRelatedSym) GetMDEntrySize() (f field.MDEntrySizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDStreamID gets MDStreamID, Tag 1500
func (m NoRelatedSym) GetMDStreamID() (f field.MDStreamIDField, err quickfix.MessageRejectError) {
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

//HasSecuritySubType returns true if SecuritySubType is present, Tag 762
func (m NoRelatedSym) HasSecuritySubType() bool {
	return m.Has(tag.SecuritySubType)
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

//HasStrikeCurrency returns true if StrikeCurrency is present, Tag 947
func (m NoRelatedSym) HasStrikeCurrency() bool {
	return m.Has(tag.StrikeCurrency)
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

//HasPool returns true if Pool is present, Tag 691
func (m NoRelatedSym) HasPool() bool {
	return m.Has(tag.Pool)
}

//HasContractSettlMonth returns true if ContractSettlMonth is present, Tag 667
func (m NoRelatedSym) HasContractSettlMonth() bool {
	return m.Has(tag.ContractSettlMonth)
}

//HasCPProgram returns true if CPProgram is present, Tag 875
func (m NoRelatedSym) HasCPProgram() bool {
	return m.Has(tag.CPProgram)
}

//HasCPRegType returns true if CPRegType is present, Tag 876
func (m NoRelatedSym) HasCPRegType() bool {
	return m.Has(tag.CPRegType)
}

//HasNoEvents returns true if NoEvents is present, Tag 864
func (m NoRelatedSym) HasNoEvents() bool {
	return m.Has(tag.NoEvents)
}

//HasDatedDate returns true if DatedDate is present, Tag 873
func (m NoRelatedSym) HasDatedDate() bool {
	return m.Has(tag.DatedDate)
}

//HasInterestAccrualDate returns true if InterestAccrualDate is present, Tag 874
func (m NoRelatedSym) HasInterestAccrualDate() bool {
	return m.Has(tag.InterestAccrualDate)
}

//HasSecurityStatus returns true if SecurityStatus is present, Tag 965
func (m NoRelatedSym) HasSecurityStatus() bool {
	return m.Has(tag.SecurityStatus)
}

//HasSettleOnOpenFlag returns true if SettleOnOpenFlag is present, Tag 966
func (m NoRelatedSym) HasSettleOnOpenFlag() bool {
	return m.Has(tag.SettleOnOpenFlag)
}

//HasInstrmtAssignmentMethod returns true if InstrmtAssignmentMethod is present, Tag 1049
func (m NoRelatedSym) HasInstrmtAssignmentMethod() bool {
	return m.Has(tag.InstrmtAssignmentMethod)
}

//HasStrikeMultiplier returns true if StrikeMultiplier is present, Tag 967
func (m NoRelatedSym) HasStrikeMultiplier() bool {
	return m.Has(tag.StrikeMultiplier)
}

//HasStrikeValue returns true if StrikeValue is present, Tag 968
func (m NoRelatedSym) HasStrikeValue() bool {
	return m.Has(tag.StrikeValue)
}

//HasMinPriceIncrement returns true if MinPriceIncrement is present, Tag 969
func (m NoRelatedSym) HasMinPriceIncrement() bool {
	return m.Has(tag.MinPriceIncrement)
}

//HasPositionLimit returns true if PositionLimit is present, Tag 970
func (m NoRelatedSym) HasPositionLimit() bool {
	return m.Has(tag.PositionLimit)
}

//HasNTPositionLimit returns true if NTPositionLimit is present, Tag 971
func (m NoRelatedSym) HasNTPositionLimit() bool {
	return m.Has(tag.NTPositionLimit)
}

//HasNoInstrumentParties returns true if NoInstrumentParties is present, Tag 1018
func (m NoRelatedSym) HasNoInstrumentParties() bool {
	return m.Has(tag.NoInstrumentParties)
}

//HasUnitOfMeasure returns true if UnitOfMeasure is present, Tag 996
func (m NoRelatedSym) HasUnitOfMeasure() bool {
	return m.Has(tag.UnitOfMeasure)
}

//HasTimeUnit returns true if TimeUnit is present, Tag 997
func (m NoRelatedSym) HasTimeUnit() bool {
	return m.Has(tag.TimeUnit)
}

//HasMaturityTime returns true if MaturityTime is present, Tag 1079
func (m NoRelatedSym) HasMaturityTime() bool {
	return m.Has(tag.MaturityTime)
}

//HasSecurityGroup returns true if SecurityGroup is present, Tag 1151
func (m NoRelatedSym) HasSecurityGroup() bool {
	return m.Has(tag.SecurityGroup)
}

//HasMinPriceIncrementAmount returns true if MinPriceIncrementAmount is present, Tag 1146
func (m NoRelatedSym) HasMinPriceIncrementAmount() bool {
	return m.Has(tag.MinPriceIncrementAmount)
}

//HasUnitOfMeasureQty returns true if UnitOfMeasureQty is present, Tag 1147
func (m NoRelatedSym) HasUnitOfMeasureQty() bool {
	return m.Has(tag.UnitOfMeasureQty)
}

//HasSecurityXMLLen returns true if SecurityXMLLen is present, Tag 1184
func (m NoRelatedSym) HasSecurityXMLLen() bool {
	return m.Has(tag.SecurityXMLLen)
}

//HasSecurityXML returns true if SecurityXML is present, Tag 1185
func (m NoRelatedSym) HasSecurityXML() bool {
	return m.Has(tag.SecurityXML)
}

//HasSecurityXMLSchema returns true if SecurityXMLSchema is present, Tag 1186
func (m NoRelatedSym) HasSecurityXMLSchema() bool {
	return m.Has(tag.SecurityXMLSchema)
}

//HasProductComplex returns true if ProductComplex is present, Tag 1227
func (m NoRelatedSym) HasProductComplex() bool {
	return m.Has(tag.ProductComplex)
}

//HasPriceUnitOfMeasure returns true if PriceUnitOfMeasure is present, Tag 1191
func (m NoRelatedSym) HasPriceUnitOfMeasure() bool {
	return m.Has(tag.PriceUnitOfMeasure)
}

//HasPriceUnitOfMeasureQty returns true if PriceUnitOfMeasureQty is present, Tag 1192
func (m NoRelatedSym) HasPriceUnitOfMeasureQty() bool {
	return m.Has(tag.PriceUnitOfMeasureQty)
}

//HasSettlMethod returns true if SettlMethod is present, Tag 1193
func (m NoRelatedSym) HasSettlMethod() bool {
	return m.Has(tag.SettlMethod)
}

//HasExerciseStyle returns true if ExerciseStyle is present, Tag 1194
func (m NoRelatedSym) HasExerciseStyle() bool {
	return m.Has(tag.ExerciseStyle)
}

//HasOptPayoutAmount returns true if OptPayoutAmount is present, Tag 1195
func (m NoRelatedSym) HasOptPayoutAmount() bool {
	return m.Has(tag.OptPayoutAmount)
}

//HasPriceQuoteMethod returns true if PriceQuoteMethod is present, Tag 1196
func (m NoRelatedSym) HasPriceQuoteMethod() bool {
	return m.Has(tag.PriceQuoteMethod)
}

//HasListMethod returns true if ListMethod is present, Tag 1198
func (m NoRelatedSym) HasListMethod() bool {
	return m.Has(tag.ListMethod)
}

//HasCapPrice returns true if CapPrice is present, Tag 1199
func (m NoRelatedSym) HasCapPrice() bool {
	return m.Has(tag.CapPrice)
}

//HasFloorPrice returns true if FloorPrice is present, Tag 1200
func (m NoRelatedSym) HasFloorPrice() bool {
	return m.Has(tag.FloorPrice)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoRelatedSym) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasFlexibleIndicator returns true if FlexibleIndicator is present, Tag 1244
func (m NoRelatedSym) HasFlexibleIndicator() bool {
	return m.Has(tag.FlexibleIndicator)
}

//HasFlexProductEligibilityIndicator returns true if FlexProductEligibilityIndicator is present, Tag 1242
func (m NoRelatedSym) HasFlexProductEligibilityIndicator() bool {
	return m.Has(tag.FlexProductEligibilityIndicator)
}

//HasValuationMethod returns true if ValuationMethod is present, Tag 1197
func (m NoRelatedSym) HasValuationMethod() bool {
	return m.Has(tag.ValuationMethod)
}

//HasContractMultiplierUnit returns true if ContractMultiplierUnit is present, Tag 1435
func (m NoRelatedSym) HasContractMultiplierUnit() bool {
	return m.Has(tag.ContractMultiplierUnit)
}

//HasFlowScheduleType returns true if FlowScheduleType is present, Tag 1439
func (m NoRelatedSym) HasFlowScheduleType() bool {
	return m.Has(tag.FlowScheduleType)
}

//HasRestructuringType returns true if RestructuringType is present, Tag 1449
func (m NoRelatedSym) HasRestructuringType() bool {
	return m.Has(tag.RestructuringType)
}

//HasSeniority returns true if Seniority is present, Tag 1450
func (m NoRelatedSym) HasSeniority() bool {
	return m.Has(tag.Seniority)
}

//HasNotionalPercentageOutstanding returns true if NotionalPercentageOutstanding is present, Tag 1451
func (m NoRelatedSym) HasNotionalPercentageOutstanding() bool {
	return m.Has(tag.NotionalPercentageOutstanding)
}

//HasOriginalNotionalPercentageOutstanding returns true if OriginalNotionalPercentageOutstanding is present, Tag 1452
func (m NoRelatedSym) HasOriginalNotionalPercentageOutstanding() bool {
	return m.Has(tag.OriginalNotionalPercentageOutstanding)
}

//HasAttachmentPoint returns true if AttachmentPoint is present, Tag 1457
func (m NoRelatedSym) HasAttachmentPoint() bool {
	return m.Has(tag.AttachmentPoint)
}

//HasDetachmentPoint returns true if DetachmentPoint is present, Tag 1458
func (m NoRelatedSym) HasDetachmentPoint() bool {
	return m.Has(tag.DetachmentPoint)
}

//HasStrikePriceDeterminationMethod returns true if StrikePriceDeterminationMethod is present, Tag 1478
func (m NoRelatedSym) HasStrikePriceDeterminationMethod() bool {
	return m.Has(tag.StrikePriceDeterminationMethod)
}

//HasStrikePriceBoundaryMethod returns true if StrikePriceBoundaryMethod is present, Tag 1479
func (m NoRelatedSym) HasStrikePriceBoundaryMethod() bool {
	return m.Has(tag.StrikePriceBoundaryMethod)
}

//HasStrikePriceBoundaryPrecision returns true if StrikePriceBoundaryPrecision is present, Tag 1480
func (m NoRelatedSym) HasStrikePriceBoundaryPrecision() bool {
	return m.Has(tag.StrikePriceBoundaryPrecision)
}

//HasUnderlyingPriceDeterminationMethod returns true if UnderlyingPriceDeterminationMethod is present, Tag 1481
func (m NoRelatedSym) HasUnderlyingPriceDeterminationMethod() bool {
	return m.Has(tag.UnderlyingPriceDeterminationMethod)
}

//HasOptPayoutType returns true if OptPayoutType is present, Tag 1482
func (m NoRelatedSym) HasOptPayoutType() bool {
	return m.Has(tag.OptPayoutType)
}

//HasNoComplexEvents returns true if NoComplexEvents is present, Tag 1483
func (m NoRelatedSym) HasNoComplexEvents() bool {
	return m.Has(tag.NoComplexEvents)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m NoRelatedSym) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasMDEntrySize returns true if MDEntrySize is present, Tag 271
func (m NoRelatedSym) HasMDEntrySize() bool {
	return m.Has(tag.MDEntrySize)
}

//HasMDStreamID returns true if MDStreamID is present, Tag 1500
func (m NoRelatedSym) HasMDStreamID() bool {
	return m.Has(tag.MDStreamID)
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

//NoComplexEvents is a repeating group element, Tag 1483
type NoComplexEvents struct {
	quickfix.Group
}

//SetComplexEventType sets ComplexEventType, Tag 1484
func (m NoComplexEvents) SetComplexEventType(v int) {
	m.Set(field.NewComplexEventType(v))
}

//SetComplexOptPayoutAmount sets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) SetComplexOptPayoutAmount(v float64) {
	m.Set(field.NewComplexOptPayoutAmount(v))
}

//SetComplexEventPrice sets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) SetComplexEventPrice(v float64) {
	m.Set(field.NewComplexEventPrice(v))
}

//SetComplexEventPriceBoundaryMethod sets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) SetComplexEventPriceBoundaryMethod(v int) {
	m.Set(field.NewComplexEventPriceBoundaryMethod(v))
}

//SetComplexEventPriceBoundaryPrecision sets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) SetComplexEventPriceBoundaryPrecision(v float64) {
	m.Set(field.NewComplexEventPriceBoundaryPrecision(v))
}

//SetComplexEventPriceTimeType sets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) SetComplexEventPriceTimeType(v int) {
	m.Set(field.NewComplexEventPriceTimeType(v))
}

//SetComplexEventCondition sets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) SetComplexEventCondition(v int) {
	m.Set(field.NewComplexEventCondition(v))
}

//SetNoComplexEventDates sets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) SetNoComplexEventDates(f NoComplexEventDatesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventType gets ComplexEventType, Tag 1484
func (m NoComplexEvents) GetComplexEventType() (f field.ComplexEventTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexOptPayoutAmount gets ComplexOptPayoutAmount, Tag 1485
func (m NoComplexEvents) GetComplexOptPayoutAmount() (f field.ComplexOptPayoutAmountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPrice gets ComplexEventPrice, Tag 1486
func (m NoComplexEvents) GetComplexEventPrice() (f field.ComplexEventPriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryMethod gets ComplexEventPriceBoundaryMethod, Tag 1487
func (m NoComplexEvents) GetComplexEventPriceBoundaryMethod() (f field.ComplexEventPriceBoundaryMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceBoundaryPrecision gets ComplexEventPriceBoundaryPrecision, Tag 1488
func (m NoComplexEvents) GetComplexEventPriceBoundaryPrecision() (f field.ComplexEventPriceBoundaryPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventPriceTimeType gets ComplexEventPriceTimeType, Tag 1489
func (m NoComplexEvents) GetComplexEventPriceTimeType() (f field.ComplexEventPriceTimeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventCondition gets ComplexEventCondition, Tag 1490
func (m NoComplexEvents) GetComplexEventCondition() (f field.ComplexEventConditionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventDates gets NoComplexEventDates, Tag 1491
func (m NoComplexEvents) GetNoComplexEventDates() (NoComplexEventDatesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventDatesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventType returns true if ComplexEventType is present, Tag 1484
func (m NoComplexEvents) HasComplexEventType() bool {
	return m.Has(tag.ComplexEventType)
}

//HasComplexOptPayoutAmount returns true if ComplexOptPayoutAmount is present, Tag 1485
func (m NoComplexEvents) HasComplexOptPayoutAmount() bool {
	return m.Has(tag.ComplexOptPayoutAmount)
}

//HasComplexEventPrice returns true if ComplexEventPrice is present, Tag 1486
func (m NoComplexEvents) HasComplexEventPrice() bool {
	return m.Has(tag.ComplexEventPrice)
}

//HasComplexEventPriceBoundaryMethod returns true if ComplexEventPriceBoundaryMethod is present, Tag 1487
func (m NoComplexEvents) HasComplexEventPriceBoundaryMethod() bool {
	return m.Has(tag.ComplexEventPriceBoundaryMethod)
}

//HasComplexEventPriceBoundaryPrecision returns true if ComplexEventPriceBoundaryPrecision is present, Tag 1488
func (m NoComplexEvents) HasComplexEventPriceBoundaryPrecision() bool {
	return m.Has(tag.ComplexEventPriceBoundaryPrecision)
}

//HasComplexEventPriceTimeType returns true if ComplexEventPriceTimeType is present, Tag 1489
func (m NoComplexEvents) HasComplexEventPriceTimeType() bool {
	return m.Has(tag.ComplexEventPriceTimeType)
}

//HasComplexEventCondition returns true if ComplexEventCondition is present, Tag 1490
func (m NoComplexEvents) HasComplexEventCondition() bool {
	return m.Has(tag.ComplexEventCondition)
}

//HasNoComplexEventDates returns true if NoComplexEventDates is present, Tag 1491
func (m NoComplexEvents) HasNoComplexEventDates() bool {
	return m.Has(tag.NoComplexEventDates)
}

//NoComplexEventDates is a repeating group element, Tag 1491
type NoComplexEventDates struct {
	quickfix.Group
}

//SetComplexEventStartDate sets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) SetComplexEventStartDate(v time.Time) {
	m.Set(field.NewComplexEventStartDate(v))
}

//SetComplexEventEndDate sets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) SetComplexEventEndDate(v time.Time) {
	m.Set(field.NewComplexEventEndDate(v))
}

//SetNoComplexEventTimes sets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) SetNoComplexEventTimes(f NoComplexEventTimesRepeatingGroup) {
	m.SetGroup(f)
}

//GetComplexEventStartDate gets ComplexEventStartDate, Tag 1492
func (m NoComplexEventDates) GetComplexEventStartDate() (f field.ComplexEventStartDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndDate gets ComplexEventEndDate, Tag 1493
func (m NoComplexEventDates) GetComplexEventEndDate() (f field.ComplexEventEndDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoComplexEventTimes gets NoComplexEventTimes, Tag 1494
func (m NoComplexEventDates) GetNoComplexEventTimes() (NoComplexEventTimesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoComplexEventTimesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasComplexEventStartDate returns true if ComplexEventStartDate is present, Tag 1492
func (m NoComplexEventDates) HasComplexEventStartDate() bool {
	return m.Has(tag.ComplexEventStartDate)
}

//HasComplexEventEndDate returns true if ComplexEventEndDate is present, Tag 1493
func (m NoComplexEventDates) HasComplexEventEndDate() bool {
	return m.Has(tag.ComplexEventEndDate)
}

//HasNoComplexEventTimes returns true if NoComplexEventTimes is present, Tag 1494
func (m NoComplexEventDates) HasNoComplexEventTimes() bool {
	return m.Has(tag.NoComplexEventTimes)
}

//NoComplexEventTimes is a repeating group element, Tag 1494
type NoComplexEventTimes struct {
	quickfix.Group
}

//SetComplexEventStartTime sets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) SetComplexEventStartTime(v string) {
	m.Set(field.NewComplexEventStartTime(v))
}

//SetComplexEventEndTime sets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) SetComplexEventEndTime(v string) {
	m.Set(field.NewComplexEventEndTime(v))
}

//GetComplexEventStartTime gets ComplexEventStartTime, Tag 1495
func (m NoComplexEventTimes) GetComplexEventStartTime() (f field.ComplexEventStartTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplexEventEndTime gets ComplexEventEndTime, Tag 1496
func (m NoComplexEventTimes) GetComplexEventEndTime() (f field.ComplexEventEndTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasComplexEventStartTime returns true if ComplexEventStartTime is present, Tag 1495
func (m NoComplexEventTimes) HasComplexEventStartTime() bool {
	return m.Has(tag.ComplexEventStartTime)
}

//HasComplexEventEndTime returns true if ComplexEventEndTime is present, Tag 1496
func (m NoComplexEventTimes) HasComplexEventEndTime() bool {
	return m.Has(tag.ComplexEventEndTime)
}

//NoComplexEventTimesRepeatingGroup is a repeating group, Tag 1494
type NoComplexEventTimesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventTimesRepeatingGroup returns an initialized, NoComplexEventTimesRepeatingGroup
func NewNoComplexEventTimesRepeatingGroup() NoComplexEventTimesRepeatingGroup {
	return NoComplexEventTimesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventTimes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartTime), quickfix.GroupElement(tag.ComplexEventEndTime)})}
}

//Add create and append a new NoComplexEventTimes to this group
func (m NoComplexEventTimesRepeatingGroup) Add() NoComplexEventTimes {
	g := m.RepeatingGroup.Add()
	return NoComplexEventTimes{g}
}

//Get returns the ith NoComplexEventTimes in the NoComplexEventTimesRepeatinGroup
func (m NoComplexEventTimesRepeatingGroup) Get(i int) NoComplexEventTimes {
	return NoComplexEventTimes{m.RepeatingGroup.Get(i)}
}

//NoComplexEventDatesRepeatingGroup is a repeating group, Tag 1491
type NoComplexEventDatesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventDatesRepeatingGroup returns an initialized, NoComplexEventDatesRepeatingGroup
func NewNoComplexEventDatesRepeatingGroup() NoComplexEventDatesRepeatingGroup {
	return NoComplexEventDatesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEventDates,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventStartDate), quickfix.GroupElement(tag.ComplexEventEndDate), NewNoComplexEventTimesRepeatingGroup()})}
}

//Add create and append a new NoComplexEventDates to this group
func (m NoComplexEventDatesRepeatingGroup) Add() NoComplexEventDates {
	g := m.RepeatingGroup.Add()
	return NoComplexEventDates{g}
}

//Get returns the ith NoComplexEventDates in the NoComplexEventDatesRepeatinGroup
func (m NoComplexEventDatesRepeatingGroup) Get(i int) NoComplexEventDates {
	return NoComplexEventDates{m.RepeatingGroup.Get(i)}
}

//NoComplexEventsRepeatingGroup is a repeating group, Tag 1483
type NoComplexEventsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoComplexEventsRepeatingGroup returns an initialized, NoComplexEventsRepeatingGroup
func NewNoComplexEventsRepeatingGroup() NoComplexEventsRepeatingGroup {
	return NoComplexEventsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoComplexEvents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ComplexEventType), quickfix.GroupElement(tag.ComplexOptPayoutAmount), quickfix.GroupElement(tag.ComplexEventPrice), quickfix.GroupElement(tag.ComplexEventPriceBoundaryMethod), quickfix.GroupElement(tag.ComplexEventPriceBoundaryPrecision), quickfix.GroupElement(tag.ComplexEventPriceTimeType), quickfix.GroupElement(tag.ComplexEventCondition), NewNoComplexEventDatesRepeatingGroup()})}
}

//Add create and append a new NoComplexEvents to this group
func (m NoComplexEventsRepeatingGroup) Add() NoComplexEvents {
	g := m.RepeatingGroup.Add()
	return NoComplexEvents{g}
}

//Get returns the ith NoComplexEvents in the NoComplexEventsRepeatinGroup
func (m NoComplexEventsRepeatingGroup) Get(i int) NoComplexEvents {
	return NoComplexEvents{m.RepeatingGroup.Get(i)}
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.SecuritySubType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.StrikeCurrency), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.Pool), quickfix.GroupElement(tag.ContractSettlMonth), quickfix.GroupElement(tag.CPProgram), quickfix.GroupElement(tag.CPRegType), NewNoEventsRepeatingGroup(), quickfix.GroupElement(tag.DatedDate), quickfix.GroupElement(tag.InterestAccrualDate), quickfix.GroupElement(tag.SecurityStatus), quickfix.GroupElement(tag.SettleOnOpenFlag), quickfix.GroupElement(tag.InstrmtAssignmentMethod), quickfix.GroupElement(tag.StrikeMultiplier), quickfix.GroupElement(tag.StrikeValue), quickfix.GroupElement(tag.MinPriceIncrement), quickfix.GroupElement(tag.PositionLimit), quickfix.GroupElement(tag.NTPositionLimit), NewNoInstrumentPartiesRepeatingGroup(), quickfix.GroupElement(tag.UnitOfMeasure), quickfix.GroupElement(tag.TimeUnit), quickfix.GroupElement(tag.MaturityTime), quickfix.GroupElement(tag.SecurityGroup), quickfix.GroupElement(tag.MinPriceIncrementAmount), quickfix.GroupElement(tag.UnitOfMeasureQty), quickfix.GroupElement(tag.SecurityXMLLen), quickfix.GroupElement(tag.SecurityXML), quickfix.GroupElement(tag.SecurityXMLSchema), quickfix.GroupElement(tag.ProductComplex), quickfix.GroupElement(tag.PriceUnitOfMeasure), quickfix.GroupElement(tag.PriceUnitOfMeasureQty), quickfix.GroupElement(tag.SettlMethod), quickfix.GroupElement(tag.ExerciseStyle), quickfix.GroupElement(tag.OptPayoutAmount), quickfix.GroupElement(tag.PriceQuoteMethod), quickfix.GroupElement(tag.ListMethod), quickfix.GroupElement(tag.CapPrice), quickfix.GroupElement(tag.FloorPrice), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.FlexibleIndicator), quickfix.GroupElement(tag.FlexProductEligibilityIndicator), quickfix.GroupElement(tag.ValuationMethod), quickfix.GroupElement(tag.ContractMultiplierUnit), quickfix.GroupElement(tag.FlowScheduleType), quickfix.GroupElement(tag.RestructuringType), quickfix.GroupElement(tag.Seniority), quickfix.GroupElement(tag.NotionalPercentageOutstanding), quickfix.GroupElement(tag.OriginalNotionalPercentageOutstanding), quickfix.GroupElement(tag.AttachmentPoint), quickfix.GroupElement(tag.DetachmentPoint), quickfix.GroupElement(tag.StrikePriceDeterminationMethod), quickfix.GroupElement(tag.StrikePriceBoundaryMethod), quickfix.GroupElement(tag.StrikePriceBoundaryPrecision), quickfix.GroupElement(tag.UnderlyingPriceDeterminationMethod), quickfix.GroupElement(tag.OptPayoutType), NewNoComplexEventsRepeatingGroup(), quickfix.GroupElement(tag.SettlType), quickfix.GroupElement(tag.MDEntrySize), quickfix.GroupElement(tag.MDStreamID)})}
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

//NoAsgnReqsRepeatingGroup is a repeating group, Tag 1499
type NoAsgnReqsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAsgnReqsRepeatingGroup returns an initialized, NoAsgnReqsRepeatingGroup
func NewNoAsgnReqsRepeatingGroup() NoAsgnReqsRepeatingGroup {
	return NoAsgnReqsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAsgnReqs,
			quickfix.GroupTemplate{NewNoPartyIDsRepeatingGroup(), NewNoRelatedSymRepeatingGroup()})}
}

//Add create and append a new NoAsgnReqs to this group
func (m NoAsgnReqsRepeatingGroup) Add() NoAsgnReqs {
	g := m.RepeatingGroup.Add()
	return NoAsgnReqs{g}
}

//Get returns the ith NoAsgnReqs in the NoAsgnReqsRepeatinGroup
func (m NoAsgnReqsRepeatingGroup) Get(i int) NoAsgnReqs {
	return NoAsgnReqs{m.RepeatingGroup.Get(i)}
}
