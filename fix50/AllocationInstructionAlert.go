package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationInstructionAlert msg type = BM.
type AllocationInstructionAlert struct {
	message.Message
}

//AllocationInstructionAlertBuilder builds AllocationInstructionAlert messages.
type AllocationInstructionAlertBuilder struct {
	message.MessageBuilder
}

//CreateAllocationInstructionAlertBuilder returns an initialized AllocationInstructionAlertBuilder with specified required fields.
func CreateAllocationInstructionAlertBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	side field.Side,
	quantity field.Quantity,
	tradedate field.TradeDate) AllocationInstructionAlertBuilder {
	var builder AllocationInstructionAlertBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(alloctype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocTransType() (*field.AllocTransType, errors.MessageRejectError) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocTransType(f *field.AllocTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocType() (*field.AllocType, errors.MessageRejectError) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocType(f *field.AllocType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryAllocID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecondaryAllocID() (*field.SecondaryAllocID, errors.MessageRejectError) {
	f := new(field.SecondaryAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryAllocID reads a SecondaryAllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecondaryAllocID(f *field.SecondaryAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RefAllocID() (*field.RefAllocID, errors.MessageRejectError) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRefAllocID(f *field.RefAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocCancReplaceReason is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocCancReplaceReason() (*field.AllocCancReplaceReason, errors.MessageRejectError) {
	f := new(field.AllocCancReplaceReason)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocCancReplaceReason reads a AllocCancReplaceReason from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocCancReplaceReason(f *field.AllocCancReplaceReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocIntermedReqType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocIntermedReqType() (*field.AllocIntermedReqType, errors.MessageRejectError) {
	f := new(field.AllocIntermedReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocIntermedReqType reads a AllocIntermedReqType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocIntermedReqType(f *field.AllocIntermedReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocLinkID() (*field.AllocLinkID, errors.MessageRejectError) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocLinkID(f *field.AllocLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocLinkType() (*field.AllocLinkType, errors.MessageRejectError) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocLinkType(f *field.AllocLinkType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BookingRefID() (*field.BookingRefID, errors.MessageRejectError) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBookingRefID(f *field.BookingRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocNoOrdersType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AllocNoOrdersType() (*field.AllocNoOrdersType, errors.MessageRejectError) {
	f := new(field.AllocNoOrdersType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocNoOrdersType reads a AllocNoOrdersType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAllocNoOrdersType(f *field.AllocNoOrdersType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PreviouslyReported() (*field.PreviouslyReported, errors.MessageRejectError) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPreviouslyReported(f *field.PreviouslyReported) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ReversalIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ReversalIndicator() (*field.ReversalIndicator, errors.MessageRejectError) {
	f := new(field.ReversalIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetReversalIndicator reads a ReversalIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetReversalIndicator(f *field.ReversalIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MatchType() (*field.MatchType, errors.MessageRejectError) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMatchType(f *field.MatchType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecuritySubType() (*field.SecuritySubType, errors.MessageRejectError) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecuritySubType(f *field.SecuritySubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeCurrency() (*field.StrikeCurrency, errors.MessageRejectError) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeCurrency(f *field.StrikeCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Pool() (*field.Pool, errors.MessageRejectError) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPool(f *field.Pool) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ContractSettlMonth() (*field.ContractSettlMonth, errors.MessageRejectError) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetContractSettlMonth(f *field.ContractSettlMonth) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CPProgram() (*field.CPProgram, errors.MessageRejectError) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCPProgram(f *field.CPProgram) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CPRegType() (*field.CPRegType, errors.MessageRejectError) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCPRegType(f *field.CPRegType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoEvents() (*field.NoEvents, errors.MessageRejectError) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoEvents(f *field.NoEvents) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DatedDate() (*field.DatedDate, errors.MessageRejectError) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDatedDate(f *field.DatedDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InterestAccrualDate() (*field.InterestAccrualDate, errors.MessageRejectError) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInterestAccrualDate(f *field.InterestAccrualDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SecurityStatus() (*field.SecurityStatus, errors.MessageRejectError) {
	f := new(field.SecurityStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSecurityStatus(f *field.SecurityStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettleOnOpenFlag() (*field.SettleOnOpenFlag, errors.MessageRejectError) {
	f := new(field.SettleOnOpenFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettleOnOpenFlag(f *field.SettleOnOpenFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethod, errors.MessageRejectError) {
	f := new(field.InstrmtAssignmentMethod)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethod) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeMultiplier() (*field.StrikeMultiplier, errors.MessageRejectError) {
	f := new(field.StrikeMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeMultiplier(f *field.StrikeMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StrikeValue() (*field.StrikeValue, errors.MessageRejectError) {
	f := new(field.StrikeValue)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStrikeValue(f *field.StrikeValue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MinPriceIncrement() (*field.MinPriceIncrement, errors.MessageRejectError) {
	f := new(field.MinPriceIncrement)
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMinPriceIncrement(f *field.MinPriceIncrement) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PositionLimit() (*field.PositionLimit, errors.MessageRejectError) {
	f := new(field.PositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPositionLimit(f *field.PositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NTPositionLimit() (*field.NTPositionLimit, errors.MessageRejectError) {
	f := new(field.NTPositionLimit)
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNTPositionLimit(f *field.NTPositionLimit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoInstrumentParties() (*field.NoInstrumentParties, errors.MessageRejectError) {
	f := new(field.NoInstrumentParties)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoInstrumentParties(f *field.NoInstrumentParties) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) UnitOfMeasure() (*field.UnitOfMeasure, errors.MessageRejectError) {
	f := new(field.UnitOfMeasure)
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetUnitOfMeasure(f *field.UnitOfMeasure) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TimeUnit() (*field.TimeUnit, errors.MessageRejectError) {
	f := new(field.TimeUnit)
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTimeUnit(f *field.TimeUnit) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MaturityTime() (*field.MaturityTime, errors.MessageRejectError) {
	f := new(field.MaturityTime)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMaturityTime(f *field.MaturityTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryForm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DeliveryForm() (*field.DeliveryForm, errors.MessageRejectError) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryForm reads a DeliveryForm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDeliveryForm(f *field.DeliveryForm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PctAtRisk is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PctAtRisk() (*field.PctAtRisk, errors.MessageRejectError) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//GetPctAtRisk reads a PctAtRisk from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPctAtRisk(f *field.PctAtRisk) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrAttrib is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoInstrAttrib() (*field.NoInstrAttrib, errors.MessageRejectError) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrAttrib reads a NoInstrAttrib from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoInstrAttrib(f *field.NoInstrAttrib) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementDesc() (*field.AgreementDesc, errors.MessageRejectError) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementDesc(f *field.AgreementDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementID() (*field.AgreementID, errors.MessageRejectError) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementID(f *field.AgreementID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementDate() (*field.AgreementDate, errors.MessageRejectError) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementDate(f *field.AgreementDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AgreementCurrency() (*field.AgreementCurrency, errors.MessageRejectError) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAgreementCurrency(f *field.AgreementCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TerminationType() (*field.TerminationType, errors.MessageRejectError) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTerminationType(f *field.TerminationType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StartDate() (*field.StartDate, errors.MessageRejectError) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStartDate(f *field.StartDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndDate() (*field.EndDate, errors.MessageRejectError) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndDate(f *field.EndDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) DeliveryType() (*field.DeliveryType, errors.MessageRejectError) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetDeliveryType(f *field.DeliveryType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MarginRatio() (*field.MarginRatio, errors.MessageRejectError) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMarginRatio(f *field.MarginRatio) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) QtyType() (*field.QtyType, errors.MessageRejectError) {
	f := new(field.QtyType)
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetQtyType(f *field.QtyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgParPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgParPx() (*field.AvgParPx, errors.MessageRejectError) {
	f := new(field.AvgParPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgParPx reads a AvgParPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgParPx(f *field.AvgParPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Spread() (*field.Spread, errors.MessageRejectError) {
	f := new(field.Spread)
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSpread(f *field.Spread) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrency, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurveName() (*field.BenchmarkCurveName, errors.MessageRejectError) {
	f := new(field.BenchmarkCurveName)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurveName(f *field.BenchmarkCurveName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkCurvePoint() (*field.BenchmarkCurvePoint, errors.MessageRejectError) {
	f := new(field.BenchmarkCurvePoint)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePoint) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkPrice() (*field.BenchmarkPrice, errors.MessageRejectError) {
	f := new(field.BenchmarkPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkPrice(f *field.BenchmarkPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkPriceType() (*field.BenchmarkPriceType, errors.MessageRejectError) {
	f := new(field.BenchmarkPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkPriceType(f *field.BenchmarkPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkSecurityID() (*field.BenchmarkSecurityID, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkSecurityID(f *field.BenchmarkSecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSource, errors.MessageRejectError) {
	f := new(field.BenchmarkSecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxPrecision is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPxPrecision() (*field.AvgPxPrecision, errors.MessageRejectError) {
	f := new(field.AvgPxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxPrecision reads a AvgPxPrecision from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPxPrecision(f *field.AvgPxPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettlType() (*field.SettlType, errors.MessageRejectError) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettlType(f *field.SettlType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) SettlDate() (*field.SettlDate, errors.MessageRejectError) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetSettlDate(f *field.SettlDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) BookingType() (*field.BookingType, errors.MessageRejectError) {
	f := new(field.BookingType)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingType reads a BookingType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetBookingType(f *field.BookingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) GrossTradeAmt() (*field.GrossTradeAmt, errors.MessageRejectError) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetGrossTradeAmt(f *field.GrossTradeAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Concession() (*field.Concession, errors.MessageRejectError) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetConcession(f *field.Concession) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotalTakedown() (*field.TotalTakedown, errors.MessageRejectError) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotalTakedown(f *field.TotalTakedown) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AutoAcceptIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AutoAcceptIndicator() (*field.AutoAcceptIndicator, errors.MessageRejectError) {
	f := new(field.AutoAcceptIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAutoAcceptIndicator reads a AutoAcceptIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAutoAcceptIndicator(f *field.AutoAcceptIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NumDaysInterest() (*field.NumDaysInterest, errors.MessageRejectError) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNumDaysInterest(f *field.NumDaysInterest) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AccruedInterestRate() (*field.AccruedInterestRate, errors.MessageRejectError) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAccruedInterestRate(f *field.AccruedInterestRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AccruedInterestAmt() (*field.AccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.AccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAccruedInterestAmt(f *field.AccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAtMaturity is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) InterestAtMaturity() (*field.InterestAtMaturity, errors.MessageRejectError) {
	f := new(field.InterestAtMaturity)
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAtMaturity reads a InterestAtMaturity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetInterestAtMaturity(f *field.InterestAtMaturity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndAccruedInterestAmt() (*field.EndAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.EndAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) StartCash() (*field.StartCash, errors.MessageRejectError) {
	f := new(field.StartCash)
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetStartCash(f *field.StartCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) EndCash() (*field.EndCash, errors.MessageRejectError) {
	f := new(field.EndCash)
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetEndCash(f *field.EndCash) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LegalConfirm() (*field.LegalConfirm, errors.MessageRejectError) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLegalConfirm(f *field.LegalConfirm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoStipulations() (*field.NoStipulations, errors.MessageRejectError) {
	f := new(field.NoStipulations)
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoStipulations(f *field.NoStipulations) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldType() (*field.YieldType, errors.MessageRejectError) {
	f := new(field.YieldType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldType reads a YieldType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldType(f *field.YieldType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Yield is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) Yield() (*field.Yield, errors.MessageRejectError) {
	f := new(field.Yield)
	err := m.Body.Get(f)
	return f, err
}

//GetYield reads a Yield from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYield(f *field.Yield) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldCalcDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldCalcDate() (*field.YieldCalcDate, errors.MessageRejectError) {
	f := new(field.YieldCalcDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldCalcDate reads a YieldCalcDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldCalcDate(f *field.YieldCalcDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionDate() (*field.YieldRedemptionDate, errors.MessageRejectError) {
	f := new(field.YieldRedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionDate reads a YieldRedemptionDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionDate(f *field.YieldRedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPrice is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionPrice() (*field.YieldRedemptionPrice, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPrice reads a YieldRedemptionPrice from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionPrice(f *field.YieldRedemptionPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//YieldRedemptionPriceType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) YieldRedemptionPriceType() (*field.YieldRedemptionPriceType, errors.MessageRejectError) {
	f := new(field.YieldRedemptionPriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetYieldRedemptionPriceType reads a YieldRedemptionPriceType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetYieldRedemptionPriceType(f *field.YieldRedemptionPriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoPosAmt() (*field.NoPosAmt, errors.MessageRejectError) {
	f := new(field.NoPosAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoPosAmt(f *field.NoPosAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoAllocs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TotNoAllocs() (*field.TotNoAllocs, errors.MessageRejectError) {
	f := new(field.TotNoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoAllocs reads a TotNoAllocs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTotNoAllocs(f *field.TotNoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) AvgPxIndicator() (*field.AvgPxIndicator, errors.MessageRejectError) {
	f := new(field.AvgPxIndicator)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetAvgPxIndicator(f *field.AvgPxIndicator) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TrdType() (*field.TrdType, errors.MessageRejectError) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTrdType(f *field.TrdType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TrdSubType() (*field.TrdSubType, errors.MessageRejectError) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTrdSubType(f *field.TrdSubType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CustOrderCapacity is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) CustOrderCapacity() (*field.CustOrderCapacity, errors.MessageRejectError) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}

//GetCustOrderCapacity reads a CustOrderCapacity from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetCustOrderCapacity(f *field.CustOrderCapacity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeInputSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) TradeInputSource() (*field.TradeInputSource, errors.MessageRejectError) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeInputSource reads a TradeInputSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetTradeInputSource(f *field.TradeInputSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MultiLegReportingType() (*field.MultiLegReportingType, errors.MessageRejectError) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMultiLegReportingType(f *field.MultiLegReportingType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) MessageEventSource() (*field.MessageEventSource, errors.MessageRejectError) {
	f := new(field.MessageEventSource)
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetMessageEventSource(f *field.MessageEventSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for AllocationInstructionAlert.
func (m AllocationInstructionAlert) RndPx() (*field.RndPx, errors.MessageRejectError) {
	f := new(field.RndPx)
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from AllocationInstructionAlert.
func (m AllocationInstructionAlert) GetRndPx(f *field.RndPx) errors.MessageRejectError {
	return m.Body.Get(f)
}
