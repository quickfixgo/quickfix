package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//CollateralAssignment msg type = AY.
type CollateralAssignment struct {
	message.Message
}

//CollateralAssignmentBuilder builds CollateralAssignment messages.
type CollateralAssignmentBuilder struct {
	message.MessageBuilder
}

//CreateCollateralAssignmentBuilder returns an initialized CollateralAssignmentBuilder with specified required fields.
func CreateCollateralAssignmentBuilder(
	collasgnid *field.CollAsgnIDField,
	collasgnreason *field.CollAsgnReasonField,
	collasgntranstype *field.CollAsgnTransTypeField,
	transacttime *field.TransactTimeField) CollateralAssignmentBuilder {
	var builder CollateralAssignmentBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("AY"))
	builder.Body().Set(collasgnid)
	builder.Body().Set(collasgnreason)
	builder.Body().Set(collasgntranstype)
	builder.Body().Set(transacttime)
	return builder
}

//CollAsgnID is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnID() (*field.CollAsgnIDField, errors.MessageRejectError) {
	f := &field.CollAsgnIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnID reads a CollAsgnID from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnID(f *field.CollAsgnIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollReqID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollReqID() (*field.CollReqIDField, errors.MessageRejectError) {
	f := &field.CollReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollReqID reads a CollReqID from CollateralAssignment.
func (m CollateralAssignment) GetCollReqID(f *field.CollReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnReason is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnReason() (*field.CollAsgnReasonField, errors.MessageRejectError) {
	f := &field.CollAsgnReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnReason reads a CollAsgnReason from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnReason(f *field.CollAsgnReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnTransType is a required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnTransType() (*field.CollAsgnTransTypeField, errors.MessageRejectError) {
	f := &field.CollAsgnTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnTransType reads a CollAsgnTransType from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnTransType(f *field.CollAsgnTransTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollAsgnRefID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CollAsgnRefID() (*field.CollAsgnRefIDField, errors.MessageRejectError) {
	f := &field.CollAsgnRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollAsgnRefID reads a CollAsgnRefID from CollateralAssignment.
func (m CollateralAssignment) GetCollAsgnRefID(f *field.CollAsgnRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for CollateralAssignment.
func (m CollateralAssignment) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from CollateralAssignment.
func (m CollateralAssignment) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from CollateralAssignment.
func (m CollateralAssignment) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from CollateralAssignment.
func (m CollateralAssignment) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from CollateralAssignment.
func (m CollateralAssignment) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from CollateralAssignment.
func (m CollateralAssignment) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from CollateralAssignment.
func (m CollateralAssignment) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CollateralAssignment.
func (m CollateralAssignment) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from CollateralAssignment.
func (m CollateralAssignment) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from CollateralAssignment.
func (m CollateralAssignment) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoExecs() (*field.NoExecsField, errors.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from CollateralAssignment.
func (m CollateralAssignment) GetNoExecs(f *field.NoExecsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrades is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrades() (*field.NoTradesField, errors.MessageRejectError) {
	f := &field.NoTradesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrades reads a NoTrades from CollateralAssignment.
func (m CollateralAssignment) GetNoTrades(f *field.NoTradesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CollateralAssignment.
func (m CollateralAssignment) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CollateralAssignment.
func (m CollateralAssignment) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CollateralAssignment.
func (m CollateralAssignment) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CollateralAssignment.
func (m CollateralAssignment) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CollateralAssignment.
func (m CollateralAssignment) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CollateralAssignment.
func (m CollateralAssignment) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CollateralAssignment.
func (m CollateralAssignment) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CollateralAssignment.
func (m CollateralAssignment) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CollateralAssignment.
func (m CollateralAssignment) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CollateralAssignment.
func (m CollateralAssignment) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CollateralAssignment.
func (m CollateralAssignment) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CollateralAssignment.
func (m CollateralAssignment) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CollateralAssignment.
func (m CollateralAssignment) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CollateralAssignment.
func (m CollateralAssignment) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CollateralAssignment.
func (m CollateralAssignment) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CollateralAssignment.
func (m CollateralAssignment) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CollateralAssignment.
func (m CollateralAssignment) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CollateralAssignment.
func (m CollateralAssignment) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CollateralAssignment.
func (m CollateralAssignment) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CollateralAssignment.
func (m CollateralAssignment) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CollateralAssignment.
func (m CollateralAssignment) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CollateralAssignment.
func (m CollateralAssignment) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CollateralAssignment.
func (m CollateralAssignment) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CollateralAssignment.
func (m CollateralAssignment) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CollateralAssignment.
func (m CollateralAssignment) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CollateralAssignment.
func (m CollateralAssignment) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CollateralAssignment.
func (m CollateralAssignment) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CollateralAssignment.
func (m CollateralAssignment) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CollateralAssignment.
func (m CollateralAssignment) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CollateralAssignment.
func (m CollateralAssignment) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CollateralAssignment.
func (m CollateralAssignment) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CollateralAssignment.
func (m CollateralAssignment) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CollateralAssignment.
func (m CollateralAssignment) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CollateralAssignment.
func (m CollateralAssignment) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CollateralAssignment.
func (m CollateralAssignment) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CollateralAssignment.
func (m CollateralAssignment) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CollateralAssignment.
func (m CollateralAssignment) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CollateralAssignment.
func (m CollateralAssignment) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityStatus() (*field.SecurityStatusField, errors.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from CollateralAssignment.
func (m CollateralAssignment) GetSecurityStatus(f *field.SecurityStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, errors.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from CollateralAssignment.
func (m CollateralAssignment) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, errors.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from CollateralAssignment.
func (m CollateralAssignment) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeMultiplier() (*field.StrikeMultiplierField, errors.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from CollateralAssignment.
func (m CollateralAssignment) GetStrikeMultiplier(f *field.StrikeMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StrikeValue() (*field.StrikeValueField, errors.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from CollateralAssignment.
func (m CollateralAssignment) GetStrikeValue(f *field.StrikeValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MinPriceIncrement() (*field.MinPriceIncrementField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from CollateralAssignment.
func (m CollateralAssignment) GetMinPriceIncrement(f *field.MinPriceIncrementField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PositionLimit() (*field.PositionLimitField, errors.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from CollateralAssignment.
func (m CollateralAssignment) GetPositionLimit(f *field.PositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NTPositionLimit() (*field.NTPositionLimitField, errors.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from CollateralAssignment.
func (m CollateralAssignment) GetNTPositionLimit(f *field.NTPositionLimitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoInstrumentParties() (*field.NoInstrumentPartiesField, errors.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from CollateralAssignment.
func (m CollateralAssignment) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for CollateralAssignment.
func (m CollateralAssignment) UnitOfMeasure() (*field.UnitOfMeasureField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from CollateralAssignment.
func (m CollateralAssignment) GetUnitOfMeasure(f *field.UnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TimeUnit() (*field.TimeUnitField, errors.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from CollateralAssignment.
func (m CollateralAssignment) GetTimeUnit(f *field.TimeUnitField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MaturityTime() (*field.MaturityTimeField, errors.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from CollateralAssignment.
func (m CollateralAssignment) GetMaturityTime(f *field.MaturityTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityGroup() (*field.SecurityGroupField, errors.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from CollateralAssignment.
func (m CollateralAssignment) GetSecurityGroup(f *field.SecurityGroupField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, errors.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from CollateralAssignment.
func (m CollateralAssignment) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for CollateralAssignment.
func (m CollateralAssignment) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from CollateralAssignment.
func (m CollateralAssignment) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityXMLLen() (*field.SecurityXMLLenField, errors.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from CollateralAssignment.
func (m CollateralAssignment) GetSecurityXMLLen(f *field.SecurityXMLLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityXML() (*field.SecurityXMLField, errors.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from CollateralAssignment.
func (m CollateralAssignment) GetSecurityXML(f *field.SecurityXMLField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SecurityXMLSchema() (*field.SecurityXMLSchemaField, errors.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from CollateralAssignment.
func (m CollateralAssignment) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ProductComplex() (*field.ProductComplexField, errors.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from CollateralAssignment.
func (m CollateralAssignment) GetProductComplex(f *field.ProductComplexField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from CollateralAssignment.
func (m CollateralAssignment) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, errors.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from CollateralAssignment.
func (m CollateralAssignment) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlMethod() (*field.SettlMethodField, errors.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from CollateralAssignment.
func (m CollateralAssignment) GetSettlMethod(f *field.SettlMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ExerciseStyle() (*field.ExerciseStyleField, errors.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from CollateralAssignment.
func (m CollateralAssignment) GetExerciseStyle(f *field.ExerciseStyleField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayAmount is a non-required field for CollateralAssignment.
func (m CollateralAssignment) OptPayAmount() (*field.OptPayAmountField, errors.MessageRejectError) {
	f := &field.OptPayAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayAmount reads a OptPayAmount from CollateralAssignment.
func (m CollateralAssignment) GetOptPayAmount(f *field.OptPayAmountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceQuoteMethod() (*field.PriceQuoteMethodField, errors.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from CollateralAssignment.
func (m CollateralAssignment) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ListMethod() (*field.ListMethodField, errors.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from CollateralAssignment.
func (m CollateralAssignment) GetListMethod(f *field.ListMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CapPrice() (*field.CapPriceField, errors.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from CollateralAssignment.
func (m CollateralAssignment) GetCapPrice(f *field.CapPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) FloorPrice() (*field.FloorPriceField, errors.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from CollateralAssignment.
func (m CollateralAssignment) GetFloorPrice(f *field.FloorPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from CollateralAssignment.
func (m CollateralAssignment) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for CollateralAssignment.
func (m CollateralAssignment) FlexibleIndicator() (*field.FlexibleIndicatorField, errors.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from CollateralAssignment.
func (m CollateralAssignment) GetFlexibleIndicator(f *field.FlexibleIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for CollateralAssignment.
func (m CollateralAssignment) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, errors.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from CollateralAssignment.
func (m CollateralAssignment) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FuturesValuationMethod is a non-required field for CollateralAssignment.
func (m CollateralAssignment) FuturesValuationMethod() (*field.FuturesValuationMethodField, errors.MessageRejectError) {
	f := &field.FuturesValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFuturesValuationMethod reads a FuturesValuationMethod from CollateralAssignment.
func (m CollateralAssignment) GetFuturesValuationMethod(f *field.FuturesValuationMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from CollateralAssignment.
func (m CollateralAssignment) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from CollateralAssignment.
func (m CollateralAssignment) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from CollateralAssignment.
func (m CollateralAssignment) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from CollateralAssignment.
func (m CollateralAssignment) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from CollateralAssignment.
func (m CollateralAssignment) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from CollateralAssignment.
func (m CollateralAssignment) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from CollateralAssignment.
func (m CollateralAssignment) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from CollateralAssignment.
func (m CollateralAssignment) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from CollateralAssignment.
func (m CollateralAssignment) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CollateralAssignment.
func (m CollateralAssignment) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Quantity() (*field.QuantityField, errors.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from CollateralAssignment.
func (m CollateralAssignment) GetQuantity(f *field.QuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from CollateralAssignment.
func (m CollateralAssignment) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CollateralAssignment.
func (m CollateralAssignment) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CollateralAssignment.
func (m CollateralAssignment) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CollateralAssignment.
func (m CollateralAssignment) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginExcess is a non-required field for CollateralAssignment.
func (m CollateralAssignment) MarginExcess() (*field.MarginExcessField, errors.MessageRejectError) {
	f := &field.MarginExcessField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginExcess reads a MarginExcess from CollateralAssignment.
func (m CollateralAssignment) GetMarginExcess(f *field.MarginExcessField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNetValue is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TotalNetValue() (*field.TotalNetValueField, errors.MessageRejectError) {
	f := &field.TotalNetValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNetValue reads a TotalNetValue from CollateralAssignment.
func (m CollateralAssignment) GetTotalNetValue(f *field.TotalNetValueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CashOutstanding is a non-required field for CollateralAssignment.
func (m CollateralAssignment) CashOutstanding() (*field.CashOutstandingField, errors.MessageRejectError) {
	f := &field.CashOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCashOutstanding reads a CashOutstanding from CollateralAssignment.
func (m CollateralAssignment) GetCashOutstanding(f *field.CashOutstandingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, errors.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from CollateralAssignment.
func (m CollateralAssignment) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from CollateralAssignment.
func (m CollateralAssignment) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMiscFees is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoMiscFees() (*field.NoMiscFeesField, errors.MessageRejectError) {
	f := &field.NoMiscFeesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMiscFees reads a NoMiscFees from CollateralAssignment.
func (m CollateralAssignment) GetNoMiscFees(f *field.NoMiscFeesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Price is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Price() (*field.PriceField, errors.MessageRejectError) {
	f := &field.PriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPrice reads a Price from CollateralAssignment.
func (m CollateralAssignment) GetPrice(f *field.PriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) PriceType() (*field.PriceTypeField, errors.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from CollateralAssignment.
func (m CollateralAssignment) GetPriceType(f *field.PriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) AccruedInterestAmt() (*field.AccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.AccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestAmt reads a AccruedInterestAmt from CollateralAssignment.
func (m CollateralAssignment) GetAccruedInterestAmt(f *field.AccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndAccruedInterestAmt is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndAccruedInterestAmt() (*field.EndAccruedInterestAmtField, errors.MessageRejectError) {
	f := &field.EndAccruedInterestAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndAccruedInterestAmt reads a EndAccruedInterestAmt from CollateralAssignment.
func (m CollateralAssignment) GetEndAccruedInterestAmt(f *field.EndAccruedInterestAmtField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StartCash() (*field.StartCashField, errors.MessageRejectError) {
	f := &field.StartCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartCash reads a StartCash from CollateralAssignment.
func (m CollateralAssignment) GetStartCash(f *field.StartCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndCash is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EndCash() (*field.EndCashField, errors.MessageRejectError) {
	f := &field.EndCashField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndCash reads a EndCash from CollateralAssignment.
func (m CollateralAssignment) GetEndCash(f *field.EndCashField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Spread is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Spread() (*field.SpreadField, errors.MessageRejectError) {
	f := &field.SpreadField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSpread reads a Spread from CollateralAssignment.
func (m CollateralAssignment) GetSpread(f *field.SpreadField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveCurrency is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveCurrency() (*field.BenchmarkCurveCurrencyField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveCurrency reads a BenchmarkCurveCurrency from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurveCurrency(f *field.BenchmarkCurveCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurveName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurveName() (*field.BenchmarkCurveNameField, errors.MessageRejectError) {
	f := &field.BenchmarkCurveNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurveName reads a BenchmarkCurveName from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurveName(f *field.BenchmarkCurveNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkCurvePoint is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkCurvePoint() (*field.BenchmarkCurvePointField, errors.MessageRejectError) {
	f := &field.BenchmarkCurvePointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkCurvePoint reads a BenchmarkCurvePoint from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkCurvePoint(f *field.BenchmarkCurvePointField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPrice is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPrice() (*field.BenchmarkPriceField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPrice reads a BenchmarkPrice from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkPrice(f *field.BenchmarkPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkPriceType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkPriceType() (*field.BenchmarkPriceTypeField, errors.MessageRejectError) {
	f := &field.BenchmarkPriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkPriceType reads a BenchmarkPriceType from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkPriceType(f *field.BenchmarkPriceTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityID() (*field.BenchmarkSecurityIDField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityID reads a BenchmarkSecurityID from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkSecurityID(f *field.BenchmarkSecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BenchmarkSecurityIDSource is a non-required field for CollateralAssignment.
func (m CollateralAssignment) BenchmarkSecurityIDSource() (*field.BenchmarkSecurityIDSourceField, errors.MessageRejectError) {
	f := &field.BenchmarkSecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetBenchmarkSecurityIDSource reads a BenchmarkSecurityIDSource from CollateralAssignment.
func (m CollateralAssignment) GetBenchmarkSecurityIDSource(f *field.BenchmarkSecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoStipulations is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoStipulations() (*field.NoStipulationsField, errors.MessageRejectError) {
	f := &field.NoStipulationsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoStipulations reads a NoStipulations from CollateralAssignment.
func (m CollateralAssignment) GetNoStipulations(f *field.NoStipulationsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDeliveryType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlDeliveryType() (*field.SettlDeliveryTypeField, errors.MessageRejectError) {
	f := &field.SettlDeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDeliveryType reads a SettlDeliveryType from CollateralAssignment.
func (m CollateralAssignment) GetSettlDeliveryType(f *field.SettlDeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbType() (*field.StandInstDbTypeField, errors.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbType(f *field.StandInstDbTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbName() (*field.StandInstDbNameField, errors.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbName(f *field.StandInstDbNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) StandInstDbID() (*field.StandInstDbIDField, errors.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from CollateralAssignment.
func (m CollateralAssignment) GetStandInstDbID(f *field.StandInstDbIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoDlvyInst is a non-required field for CollateralAssignment.
func (m CollateralAssignment) NoDlvyInst() (*field.NoDlvyInstField, errors.MessageRejectError) {
	f := &field.NoDlvyInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoDlvyInst reads a NoDlvyInst from CollateralAssignment.
func (m CollateralAssignment) GetNoDlvyInst(f *field.NoDlvyInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from CollateralAssignment.
func (m CollateralAssignment) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from CollateralAssignment.
func (m CollateralAssignment) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from CollateralAssignment.
func (m CollateralAssignment) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for CollateralAssignment.
func (m CollateralAssignment) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from CollateralAssignment.
func (m CollateralAssignment) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for CollateralAssignment.
func (m CollateralAssignment) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from CollateralAssignment.
func (m CollateralAssignment) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for CollateralAssignment.
func (m CollateralAssignment) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from CollateralAssignment.
func (m CollateralAssignment) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from CollateralAssignment.
func (m CollateralAssignment) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for CollateralAssignment.
func (m CollateralAssignment) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from CollateralAssignment.
func (m CollateralAssignment) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
