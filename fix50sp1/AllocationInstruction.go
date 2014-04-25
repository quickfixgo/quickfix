package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationInstruction msg type = J.
type AllocationInstruction struct {
	message.Message
}

//AllocationInstructionBuilder builds AllocationInstruction messages.
type AllocationInstructionBuilder struct {
	message.MessageBuilder
}

//CreateAllocationInstructionBuilder returns an initialized AllocationInstructionBuilder with specified required fields.
func CreateAllocationInstructionBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	side field.Side,
	quantity field.Quantity,
	tradedate field.TradeDate) AllocationInstructionBuilder {
	var builder AllocationInstructionBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("J"))
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(alloctype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationInstruction.
func (m AllocationInstruction) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocTransType() (*field.AllocTransType, errors.MessageRejectError) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationInstruction.
func (m AllocationInstruction) GetAllocTransType(f *field.AllocTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a required field for AllocationInstruction.
func (m AllocationInstruction) AllocType() (*field.AllocType, errors.MessageRejectError) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from AllocationInstruction.
func (m AllocationInstruction) GetAllocType(f *field.AllocType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecondaryAllocID() (*field.SecondaryAllocID, errors.MessageRejectError) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationInstruction.
func (m AllocationInstruction) GetSecondaryAllocID(f *field.SecondaryAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RefAllocID() (*field.RefAllocID, errors.MessageRejectError) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from AllocationInstruction.
func (m AllocationInstruction) GetRefAllocID(f *field.RefAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocCancReplaceReason is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocCancReplaceReason() (*field.AllocCancReplaceReason, errors.MessageRejectError) {
	f := new(field.AllocCancReplaceReason)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocCancReplaceReason reads a AllocCancReplaceReason from AllocationInstruction.
func (m AllocationInstruction) GetAllocCancReplaceReason(f *field.AllocCancReplaceReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocIntermedReqType() (*field.AllocIntermedReqType, errors.MessageRejectError) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationInstruction.
func (m AllocationInstruction) GetAllocIntermedReqType(f *field.AllocIntermedReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocLinkID() (*field.AllocLinkID, errors.MessageRejectError) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from AllocationInstruction.
func (m AllocationInstruction) GetAllocLinkID(f *field.AllocLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocLinkType() (*field.AllocLinkType, errors.MessageRejectError) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from AllocationInstruction.
func (m AllocationInstruction) GetAllocLinkType(f *field.AllocLinkType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BookingRefID() (*field.BookingRefID, errors.MessageRejectError) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from AllocationInstruction.
func (m AllocationInstruction) GetBookingRefID(f *field.BookingRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocNoOrdersType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AllocNoOrdersType() (*field.AllocNoOrdersType, errors.MessageRejectError) {
	f := new(field.AllocNoOrdersType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocNoOrdersType reads a AllocNoOrdersType from AllocationInstruction.
func (m AllocationInstruction) GetAllocNoOrdersType(f *field.AllocNoOrdersType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from AllocationInstruction.
func (m AllocationInstruction) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from AllocationInstruction.
func (m AllocationInstruction) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PreviouslyReported() (*field.PreviouslyReported, errors.MessageRejectError) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from AllocationInstruction.
func (m AllocationInstruction) GetPreviouslyReported(f *field.PreviouslyReported) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReversalIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ReversalIndicator() (*field.ReversalIndicator, errors.MessageRejectError) {
	f := new(field.ReversalIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetReversalIndicator reads a ReversalIndicator from AllocationInstruction.
func (m AllocationInstruction) GetReversalIndicator(f *field.ReversalIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MatchType() (*field.MatchType, errors.MessageRejectError) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from AllocationInstruction.
func (m AllocationInstruction) GetMatchType(f *field.MatchType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for AllocationInstruction.
func (m AllocationInstruction) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from AllocationInstruction.
func (m AllocationInstruction) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AllocationInstruction.
func (m AllocationInstruction) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AllocationInstruction.
func (m AllocationInstruction) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AllocationInstruction.
func (m AllocationInstruction) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AllocationInstruction.
func (m AllocationInstruction) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AllocationInstruction.
func (m AllocationInstruction) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationInstruction.
func (m AllocationInstruction) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AllocationInstruction.
func (m AllocationInstruction) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationInstruction.
func (m AllocationInstruction) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AllocationInstruction.
func (m AllocationInstruction) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AllocationInstruction.
func (m AllocationInstruction) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AllocationInstruction.
func (m AllocationInstruction) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AllocationInstruction.
func (m AllocationInstruction) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AllocationInstruction.
func (m AllocationInstruction) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AllocationInstruction.
func (m AllocationInstruction) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AllocationInstruction.
func (m AllocationInstruction) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AllocationInstruction.
func (m AllocationInstruction) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AllocationInstruction.
func (m AllocationInstruction) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AllocationInstruction.
func (m AllocationInstruction) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AllocationInstruction.
func (m AllocationInstruction) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AllocationInstruction.
func (m AllocationInstruction) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AllocationInstruction.
func (m AllocationInstruction) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AllocationInstruction.
func (m AllocationInstruction) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AllocationInstruction.
func (m AllocationInstruction) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AllocationInstruction.
func (m AllocationInstruction) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AllocationInstruction.
func (m AllocationInstruction) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AllocationInstruction.
func (m AllocationInstruction) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AllocationInstruction.
func (m AllocationInstruction) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AllocationInstruction.
func (m AllocationInstruction) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AllocationInstruction.
func (m AllocationInstruction) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AllocationInstruction.
func (m AllocationInstruction) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AllocationInstruction.
func (m AllocationInstruction) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AllocationInstruction.
func (m AllocationInstruction) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AllocationInstruction.
func (m AllocationInstruction) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AllocationInstruction.
func (m AllocationInstruction) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AllocationInstruction.
func (m AllocationInstruction) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AllocationInstruction.
func (m AllocationInstruction) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AllocationInstruction.
func (m AllocationInstruction) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AllocationInstruction.
func (m AllocationInstruction) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AllocationInstruction.
func (m AllocationInstruction) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AllocationInstruction.
func (m AllocationInstruction) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AllocationInstruction.
func (m AllocationInstruction) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AllocationInstruction.
func (m AllocationInstruction) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AllocationInstruction.
func (m AllocationInstruction) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from AllocationInstruction.
func (m AllocationInstruction) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from AllocationInstruction.
func (m AllocationInstruction) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from AllocationInstruction.
func (m AllocationInstruction) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from AllocationInstruction.
func (m AllocationInstruction) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from AllocationInstruction.
func (m AllocationInstruction) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from AllocationInstruction.
func (m AllocationInstruction) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from AllocationInstruction.
func (m AllocationInstruction) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from AllocationInstruction.
func (m AllocationInstruction) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from AllocationInstruction.
func (m AllocationInstruction) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for AllocationInstruction.
func (m AllocationInstruction) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from AllocationInstruction.
func (m AllocationInstruction) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from AllocationInstruction.
func (m AllocationInstruction) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from AllocationInstruction.
func (m AllocationInstruction) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityGroup() (*field.SecurityGroup, errors.MessageRejectError) {
	f := new(field.SecurityGroup)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from AllocationInstruction.
func (m AllocationInstruction) GetSecurityGroup(f *field.SecurityGroup) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MinPriceIncrementAmount() (*field.MinPriceIncrementAmount, errors.MessageRejectError) {
	f := new(field.MinPriceIncrementAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from AllocationInstruction.
func (m AllocationInstruction) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for AllocationInstruction.
func (m AllocationInstruction) UnitOfMeasureQty() (*field.UnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.UnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from AllocationInstruction.
func (m AllocationInstruction) GetUnitOfMeasureQty(f *field.UnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityXMLLen() (*field.SecurityXMLLen, errors.MessageRejectError) {
	f := new(field.SecurityXMLLen)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from AllocationInstruction.
func (m AllocationInstruction) GetSecurityXMLLen(f *field.SecurityXMLLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityXML() (*field.SecurityXML, errors.MessageRejectError) {
	f := new(field.SecurityXML)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from AllocationInstruction.
func (m AllocationInstruction) GetSecurityXML(f *field.SecurityXML) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SecurityXMLSchema() (*field.SecurityXMLSchema, errors.MessageRejectError) {
	f := new(field.SecurityXMLSchema)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from AllocationInstruction.
func (m AllocationInstruction) GetSecurityXMLSchema(f *field.SecurityXMLSchema) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ProductComplex() (*field.ProductComplex, errors.MessageRejectError) {
	f := new(field.ProductComplex)
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from AllocationInstruction.
func (m AllocationInstruction) GetProductComplex(f *field.ProductComplex) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PriceUnitOfMeasure() (*field.PriceUnitOfMeasure, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from AllocationInstruction.
func (m AllocationInstruction) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQty, errors.MessageRejectError) {
	f := new(field.PriceUnitOfMeasureQty)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from AllocationInstruction.
func (m AllocationInstruction) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQty) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettlMethod() (*field.SettlMethod, errors.MessageRejectError) {
	f := new(field.SettlMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from AllocationInstruction.
func (m AllocationInstruction) GetSettlMethod(f *field.SettlMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ExerciseStyle() (*field.ExerciseStyle, errors.MessageRejectError) {
	f := new(field.ExerciseStyle)
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from AllocationInstruction.
func (m AllocationInstruction) GetExerciseStyle(f *field.ExerciseStyle) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for AllocationInstruction.
func (m AllocationInstruction) OptPayAmount() (*field.OptPayAmount, errors.MessageRejectError) {
	f := new(field.OptPayAmount)
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from AllocationInstruction.
func (m AllocationInstruction) GetOptPayAmount(f *field.OptPayAmount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PriceQuoteMethod() (*field.PriceQuoteMethod, errors.MessageRejectError) {
	f := new(field.PriceQuoteMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from AllocationInstruction.
func (m AllocationInstruction) GetPriceQuoteMethod(f *field.PriceQuoteMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ListMethod() (*field.ListMethod, errors.MessageRejectError) {
	f := new(field.ListMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from AllocationInstruction.
func (m AllocationInstruction) GetListMethod(f *field.ListMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CapPrice() (*field.CapPrice, errors.MessageRejectError) {
	f := new(field.CapPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from AllocationInstruction.
func (m AllocationInstruction) GetCapPrice(f *field.CapPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) FloorPrice() (*field.FloorPrice, errors.MessageRejectError) {
	f := new(field.FloorPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from AllocationInstruction.
func (m AllocationInstruction) GetFloorPrice(f *field.FloorPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PutOrCall() (*field.PutOrCall, errors.MessageRejectError) {
	f := new(field.PutOrCall)
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from AllocationInstruction.
func (m AllocationInstruction) GetPutOrCall(f *field.PutOrCall) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) FlexibleIndicator() (*field.FlexibleIndicator, errors.MessageRejectError) {
	f := new(field.FlexibleIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from AllocationInstruction.
func (m AllocationInstruction) GetFlexibleIndicator(f *field.FlexibleIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicator, errors.MessageRejectError) {
	f := new(field.FlexProductEligibilityIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from AllocationInstruction.
func (m AllocationInstruction) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for AllocationInstruction.
func (m AllocationInstruction) FuturesValuationMethod() (*field.FuturesValuationMethod, errors.MessageRejectError) {
	f := new(field.FuturesValuationMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from AllocationInstruction.
func (m AllocationInstruction) GetFuturesValuationMethod(f *field.FuturesValuationMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DeliveryForm() (*field.DeliveryForm, errors.MessageRejectError) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from AllocationInstruction.
func (m AllocationInstruction) GetDeliveryForm(f *field.DeliveryForm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PctAtRisk() (*field.PctAtRisk, errors.MessageRejectError) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from AllocationInstruction.
func (m AllocationInstruction) GetPctAtRisk(f *field.PctAtRisk) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoInstrAttrib() (*field.NoInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from AllocationInstruction.
func (m AllocationInstruction) GetNoInstrAttrib(f *field.NoInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from AllocationInstruction.
func (m AllocationInstruction) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from AllocationInstruction.
func (m AllocationInstruction) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from AllocationInstruction.
func (m AllocationInstruction) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from AllocationInstruction.
func (m AllocationInstruction) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from AllocationInstruction.
func (m AllocationInstruction) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from AllocationInstruction.
func (m AllocationInstruction) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from AllocationInstruction.
func (m AllocationInstruction) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from AllocationInstruction.
func (m AllocationInstruction) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from AllocationInstruction.
func (m AllocationInstruction) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AllocationInstruction.
func (m AllocationInstruction) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AllocationInstruction.
func (m AllocationInstruction) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for AllocationInstruction.
func (m AllocationInstruction) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationInstruction.
func (m AllocationInstruction) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from AllocationInstruction.
func (m AllocationInstruction) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from AllocationInstruction.
func (m AllocationInstruction) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from AllocationInstruction.
func (m AllocationInstruction) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from AllocationInstruction.
func (m AllocationInstruction) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from AllocationInstruction.
func (m AllocationInstruction) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from AllocationInstruction.
func (m AllocationInstruction) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from AllocationInstruction.
func (m AllocationInstruction) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgParPx() (*field.AvgParPx, errors.MessageRejectError) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from AllocationInstruction.
func (m AllocationInstruction) GetAvgParPx(f *field.AvgParPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from AllocationInstruction.
func (m AllocationInstruction) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from AllocationInstruction.
func (m AllocationInstruction) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AllocationInstruction.
func (m AllocationInstruction) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgPxPrecision() (*field.AvgPxPrecision, errors.MessageRejectError) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from AllocationInstruction.
func (m AllocationInstruction) GetAvgPxPrecision(f *field.AvgPxPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationInstruction.
func (m AllocationInstruction) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationInstruction.
func (m AllocationInstruction) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationInstruction.
func (m AllocationInstruction) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationInstruction.
func (m AllocationInstruction) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from AllocationInstruction.
func (m AllocationInstruction) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from AllocationInstruction.
func (m AllocationInstruction) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from AllocationInstruction.
func (m AllocationInstruction) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) GrossTradeAmt() (*field.GrossTradeAmt, errors.MessageRejectError) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from AllocationInstruction.
func (m AllocationInstruction) GetGrossTradeAmt(f *field.GrossTradeAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Concession() (*field.Concession, errors.MessageRejectError) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from AllocationInstruction.
func (m AllocationInstruction) GetConcession(f *field.Concession) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotalTakedown() (*field.TotalTakedown, errors.MessageRejectError) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from AllocationInstruction.
func (m AllocationInstruction) GetTotalTakedown(f *field.TotalTakedown) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from AllocationInstruction.
func (m AllocationInstruction) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for AllocationInstruction.
func (m AllocationInstruction) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from AllocationInstruction.
func (m AllocationInstruction) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AutoAcceptIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AutoAcceptIndicator() (*field.AutoAcceptIndicator, errors.MessageRejectError) {
	f := new(field.AutoAcceptIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAutoAcceptIndicator reads a AutoAcceptIndicator from AllocationInstruction.
func (m AllocationInstruction) GetAutoAcceptIndicator(f *field.AutoAcceptIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationInstruction.
func (m AllocationInstruction) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationInstruction.
func (m AllocationInstruction) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationInstruction.
func (m AllocationInstruction) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NumDaysInterest() (*field.NumDaysInterest, errors.MessageRejectError) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from AllocationInstruction.
func (m AllocationInstruction) GetNumDaysInterest(f *field.NumDaysInterest) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AccruedInterestRate() (*field.AccruedInterestRate, errors.MessageRejectError) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from AllocationInstruction.
func (m AllocationInstruction) GetAccruedInterestRate(f *field.AccruedInterestRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from AllocationInstruction.
func (m AllocationInstruction) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from AllocationInstruction.
func (m AllocationInstruction) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for AllocationInstruction.
func (m AllocationInstruction) InterestAtMaturity() (*field.InterestAtMaturity, errors.MessageRejectError) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from AllocationInstruction.
func (m AllocationInstruction) GetInterestAtMaturity(f *field.InterestAtMaturity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from AllocationInstruction.
func (m AllocationInstruction) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for AllocationInstruction.
func (m AllocationInstruction) StartCash() (*field.StartCash, errors.MessageRejectError) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from AllocationInstruction.
func (m AllocationInstruction) GetStartCash(f *field.StartCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for AllocationInstruction.
func (m AllocationInstruction) EndCash() (*field.EndCash, errors.MessageRejectError) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from AllocationInstruction.
func (m AllocationInstruction) GetEndCash(f *field.EndCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LegalConfirm() (*field.LegalConfirm, errors.MessageRejectError) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationInstruction.
func (m AllocationInstruction) GetLegalConfirm(f *field.LegalConfirm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from AllocationInstruction.
func (m AllocationInstruction) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from AllocationInstruction.
func (m AllocationInstruction) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for AllocationInstruction.
func (m AllocationInstruction) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from AllocationInstruction.
func (m AllocationInstruction) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from AllocationInstruction.
func (m AllocationInstruction) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from AllocationInstruction.
func (m AllocationInstruction) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from AllocationInstruction.
func (m AllocationInstruction) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from AllocationInstruction.
func (m AllocationInstruction) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoAllocs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TotNoAllocs() (*field.TotNoAllocs, errors.MessageRejectError) {
	f := new(field.TotNoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoAllocs reads a TotNoAllocs from AllocationInstruction.
func (m AllocationInstruction) GetTotNoAllocs(f *field.TotNoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for AllocationInstruction.
func (m AllocationInstruction) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from AllocationInstruction.
func (m AllocationInstruction) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationInstruction.
func (m AllocationInstruction) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AllocationInstruction.
func (m AllocationInstruction) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AllocationInstruction.
func (m AllocationInstruction) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationInstruction.
func (m AllocationInstruction) AvgPxIndicator() (*field.AvgPxIndicator, errors.MessageRejectError) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationInstruction.
func (m AllocationInstruction) GetAvgPxIndicator(f *field.AvgPxIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationInstruction.
func (m AllocationInstruction) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationInstruction.
func (m AllocationInstruction) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TrdType() (*field.TrdType, errors.MessageRejectError) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from AllocationInstruction.
func (m AllocationInstruction) GetTrdType(f *field.TrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TrdSubType() (*field.TrdSubType, errors.MessageRejectError) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from AllocationInstruction.
func (m AllocationInstruction) GetTrdSubType(f *field.TrdSubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for AllocationInstruction.
func (m AllocationInstruction) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from AllocationInstruction.
func (m AllocationInstruction) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeInputSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) TradeInputSource() (*field.TradeInputSource, errors.MessageRejectError) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeInputSource reads a TradeInputSource from AllocationInstruction.
func (m AllocationInstruction) GetTradeInputSource(f *field.TradeInputSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MultiLegReportingType() (*field.MultiLegReportingType, errors.MessageRejectError) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from AllocationInstruction.
func (m AllocationInstruction) GetMultiLegReportingType(f *field.MultiLegReportingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for AllocationInstruction.
func (m AllocationInstruction) MessageEventSource() (*field.MessageEventSource, errors.MessageRejectError) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from AllocationInstruction.
func (m AllocationInstruction) GetMessageEventSource(f *field.MessageEventSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for AllocationInstruction.
func (m AllocationInstruction) RndPx() (*field.RndPx, errors.MessageRejectError) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from AllocationInstruction.
func (m AllocationInstruction) GetRndPx(f *field.RndPx) errors.MessageRejectError {
	return m.Body.Get(f)
}
