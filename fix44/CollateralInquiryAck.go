package fix44

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//CollateralInquiryAck msg type = BG.
type CollateralInquiryAck struct {
	message.Message
}

//CollateralInquiryAckBuilder builds CollateralInquiryAck messages.
type CollateralInquiryAckBuilder struct {
	message.MessageBuilder
}

//CreateCollateralInquiryAckBuilder returns an initialized CollateralInquiryAckBuilder with specified required fields.
func CreateCollateralInquiryAckBuilder(
	collinquiryid *field.CollInquiryIDField,
	collinquirystatus *field.CollInquiryStatusField) CollateralInquiryAckBuilder {
	var builder CollateralInquiryAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("BG"))
	builder.Body.Set(collinquiryid)
	builder.Body.Set(collinquirystatus)
	return builder
}

//CollInquiryID is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryID() (*field.CollInquiryIDField, errors.MessageRejectError) {
	f := &field.CollInquiryIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollInquiryID reads a CollInquiryID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCollInquiryID(f *field.CollInquiryIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollInquiryStatus is a required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryStatus() (*field.CollInquiryStatusField, errors.MessageRejectError) {
	f := &field.CollInquiryStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollInquiryStatus reads a CollInquiryStatus from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCollInquiryStatus(f *field.CollInquiryStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CollInquiryResult is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CollInquiryResult() (*field.CollInquiryResultField, errors.MessageRejectError) {
	f := &field.CollInquiryResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCollInquiryResult reads a CollInquiryResult from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCollInquiryResult(f *field.CollInquiryResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoCollInquiryQualifier() (*field.NoCollInquiryQualifierField, errors.MessageRejectError) {
	f := &field.NoCollInquiryQualifierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoCollInquiryQualifier reads a NoCollInquiryQualifier from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoCollInquiryQualifier(f *field.NoCollInquiryQualifierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNumReports is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TotNumReports() (*field.TotNumReportsField, errors.MessageRejectError) {
	f := &field.TotNumReportsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNumReports reads a TotNumReports from CollateralInquiryAck.
func (m CollateralInquiryAck) GetTotNumReports(f *field.TotNumReportsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Account is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Account() (*field.AccountField, errors.MessageRejectError) {
	f := &field.AccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccount reads a Account from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAccount(f *field.AccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccountType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AccountType() (*field.AccountTypeField, errors.MessageRejectError) {
	f := &field.AccountTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAccountType reads a AccountType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAccountType(f *field.AccountTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OrderID() (*field.OrderIDField, errors.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetOrderID(f *field.OrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryOrderID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryOrderID() (*field.SecondaryOrderIDField, errors.MessageRejectError) {
	f := &field.SecondaryOrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryOrderID reads a SecondaryOrderID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecondaryOrderID(f *field.SecondaryOrderIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecondaryClOrdID() (*field.SecondaryClOrdIDField, errors.MessageRejectError) {
	f := &field.SecondaryClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryClOrdID reads a SecondaryClOrdID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecondaryClOrdID(f *field.SecondaryClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoExecs() (*field.NoExecsField, errors.MessageRejectError) {
	f := &field.NoExecsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoExecs(f *field.NoExecsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrades is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoTrades() (*field.NoTradesField, errors.MessageRejectError) {
	f := &field.NoTradesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrades reads a NoTrades from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoTrades(f *field.NoTradesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityIDSource() (*field.SecurityIDSourceField, errors.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecurityIDSource(f *field.SecurityIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoSecurityAltID() (*field.NoSecurityAltIDField, errors.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoSecurityAltID(f *field.NoSecurityAltIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from CollateralInquiryAck.
func (m CollateralInquiryAck) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecuritySubType() (*field.SecuritySubTypeField, errors.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecuritySubType(f *field.SecuritySubTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from CollateralInquiryAck.
func (m CollateralInquiryAck) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MaturityDate() (*field.MaturityDateField, errors.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetMaturityDate(f *field.MaturityDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponPaymentDate() (*field.CouponPaymentDateField, errors.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCouponPaymentDate(f *field.CouponPaymentDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) IssueDate() (*field.IssueDateField, errors.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetIssueDate(f *field.IssueDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, errors.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseTerm() (*field.RepurchaseTermField, errors.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from CollateralInquiryAck.
func (m CollateralInquiryAck) GetRepurchaseTerm(f *field.RepurchaseTermField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RepurchaseRate() (*field.RepurchaseRateField, errors.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetRepurchaseRate(f *field.RepurchaseRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Factor() (*field.FactorField, errors.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from CollateralInquiryAck.
func (m CollateralInquiryAck) GetFactor(f *field.FactorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CreditRating() (*field.CreditRatingField, errors.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCreditRating(f *field.CreditRatingField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InstrRegistry() (*field.InstrRegistryField, errors.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from CollateralInquiryAck.
func (m CollateralInquiryAck) GetInstrRegistry(f *field.InstrRegistryField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CountryOfIssue() (*field.CountryOfIssueField, errors.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCountryOfIssue(f *field.CountryOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, errors.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from CollateralInquiryAck.
func (m CollateralInquiryAck) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) LocaleOfIssue() (*field.LocaleOfIssueField, errors.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from CollateralInquiryAck.
func (m CollateralInquiryAck) GetLocaleOfIssue(f *field.LocaleOfIssueField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) RedemptionDate() (*field.RedemptionDateField, errors.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetRedemptionDate(f *field.RedemptionDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from CollateralInquiryAck.
func (m CollateralInquiryAck) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StrikeCurrency() (*field.StrikeCurrencyField, errors.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from CollateralInquiryAck.
func (m CollateralInquiryAck) GetStrikeCurrency(f *field.StrikeCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from CollateralInquiryAck.
func (m CollateralInquiryAck) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from CollateralInquiryAck.
func (m CollateralInquiryAck) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from CollateralInquiryAck.
func (m CollateralInquiryAck) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Pool() (*field.PoolField, errors.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from CollateralInquiryAck.
func (m CollateralInquiryAck) GetPool(f *field.PoolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ContractSettlMonth() (*field.ContractSettlMonthField, errors.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from CollateralInquiryAck.
func (m CollateralInquiryAck) GetContractSettlMonth(f *field.ContractSettlMonthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPProgram() (*field.CPProgramField, errors.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCPProgram(f *field.CPProgramField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) CPRegType() (*field.CPRegTypeField, errors.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCPRegType(f *field.CPRegTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoEvents() (*field.NoEventsField, errors.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoEvents(f *field.NoEventsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DatedDate() (*field.DatedDateField, errors.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetDatedDate(f *field.DatedDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) InterestAccrualDate() (*field.InterestAccrualDateField, errors.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetInterestAccrualDate(f *field.InterestAccrualDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDesc is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDesc() (*field.AgreementDescField, errors.MessageRejectError) {
	f := &field.AgreementDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDesc reads a AgreementDesc from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAgreementDesc(f *field.AgreementDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementID() (*field.AgreementIDField, errors.MessageRejectError) {
	f := &field.AgreementIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementID reads a AgreementID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAgreementID(f *field.AgreementIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementDate() (*field.AgreementDateField, errors.MessageRejectError) {
	f := &field.AgreementDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementDate reads a AgreementDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAgreementDate(f *field.AgreementDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AgreementCurrency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) AgreementCurrency() (*field.AgreementCurrencyField, errors.MessageRejectError) {
	f := &field.AgreementCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAgreementCurrency reads a AgreementCurrency from CollateralInquiryAck.
func (m CollateralInquiryAck) GetAgreementCurrency(f *field.AgreementCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TerminationType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TerminationType() (*field.TerminationTypeField, errors.MessageRejectError) {
	f := &field.TerminationTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTerminationType reads a TerminationType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetTerminationType(f *field.TerminationTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StartDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) StartDate() (*field.StartDateField, errors.MessageRejectError) {
	f := &field.StartDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStartDate reads a StartDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetStartDate(f *field.StartDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EndDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EndDate() (*field.EndDateField, errors.MessageRejectError) {
	f := &field.EndDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEndDate reads a EndDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEndDate(f *field.EndDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//DeliveryType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) DeliveryType() (*field.DeliveryTypeField, errors.MessageRejectError) {
	f := &field.DeliveryTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDeliveryType reads a DeliveryType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetDeliveryType(f *field.DeliveryTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarginRatio is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) MarginRatio() (*field.MarginRatioField, errors.MessageRejectError) {
	f := &field.MarginRatioField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarginRatio reads a MarginRatio from CollateralInquiryAck.
func (m CollateralInquiryAck) GetMarginRatio(f *field.MarginRatioField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlDate() (*field.SettlDateField, errors.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSettlDate(f *field.SettlDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Quantity() (*field.QuantityField, errors.MessageRejectError) {
	f := &field.QuantityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from CollateralInquiryAck.
func (m CollateralInquiryAck) GetQuantity(f *field.QuantityField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) QtyType() (*field.QtyTypeField, errors.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetQtyType(f *field.QtyTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Currency() (*field.CurrencyField, errors.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from CollateralInquiryAck.
func (m CollateralInquiryAck) GetCurrency(f *field.CurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from CollateralInquiryAck.
func (m CollateralInquiryAck) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) TradingSessionSubID() (*field.TradingSessionSubIDField, errors.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetTradingSessionSubID(f *field.TradingSessionSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) SettlSessSubID() (*field.SettlSessSubIDField, errors.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from CollateralInquiryAck.
func (m CollateralInquiryAck) GetSettlSessSubID(f *field.SettlSessSubIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from CollateralInquiryAck.
func (m CollateralInquiryAck) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseTransportType() (*field.ResponseTransportTypeField, errors.MessageRejectError) {
	f := &field.ResponseTransportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from CollateralInquiryAck.
func (m CollateralInquiryAck) GetResponseTransportType(f *field.ResponseTransportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) ResponseDestination() (*field.ResponseDestinationField, errors.MessageRejectError) {
	f := &field.ResponseDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from CollateralInquiryAck.
func (m CollateralInquiryAck) GetResponseDestination(f *field.ResponseDestinationField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from CollateralInquiryAck.
func (m CollateralInquiryAck) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for CollateralInquiryAck.
func (m CollateralInquiryAck) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from CollateralInquiryAck.
func (m CollateralInquiryAck) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
