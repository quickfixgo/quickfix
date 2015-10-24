//Package tradecapturereportack msg type = AR.
package tradecapturereportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a TradeCaptureReportAck wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//TradeReportID is a non-required field for TradeCaptureReportAck.
func (m Message) TradeReportID() (*field.TradeReportIDField, quickfix.MessageRejectError) {
	f := &field.TradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportID reads a TradeReportID from TradeCaptureReportAck.
func (m Message) GetTradeReportID(f *field.TradeReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportTransType is a non-required field for TradeCaptureReportAck.
func (m Message) TradeReportTransType() (*field.TradeReportTransTypeField, quickfix.MessageRejectError) {
	f := &field.TradeReportTransTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportTransType reads a TradeReportTransType from TradeCaptureReportAck.
func (m Message) GetTradeReportTransType(f *field.TradeReportTransTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportType is a non-required field for TradeCaptureReportAck.
func (m Message) TradeReportType() (*field.TradeReportTypeField, quickfix.MessageRejectError) {
	f := &field.TradeReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportType reads a TradeReportType from TradeCaptureReportAck.
func (m Message) GetTradeReportType(f *field.TradeReportTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdType is a non-required field for TradeCaptureReportAck.
func (m Message) TrdType() (*field.TrdTypeField, quickfix.MessageRejectError) {
	f := &field.TrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdType reads a TrdType from TradeCaptureReportAck.
func (m Message) GetTrdType(f *field.TrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdSubType is a non-required field for TradeCaptureReportAck.
func (m Message) TrdSubType() (*field.TrdSubTypeField, quickfix.MessageRejectError) {
	f := &field.TrdSubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdSubType reads a TrdSubType from TradeCaptureReportAck.
func (m Message) GetTrdSubType(f *field.TrdSubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryTrdType() (*field.SecondaryTrdTypeField, quickfix.MessageRejectError) {
	f := &field.SecondaryTrdTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTrdType reads a SecondaryTrdType from TradeCaptureReportAck.
func (m Message) GetSecondaryTrdType(f *field.SecondaryTrdTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransferReason is a non-required field for TradeCaptureReportAck.
func (m Message) TransferReason() (*field.TransferReasonField, quickfix.MessageRejectError) {
	f := &field.TransferReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransferReason reads a TransferReason from TradeCaptureReportAck.
func (m Message) GetTransferReason(f *field.TransferReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecType is a non-required field for TradeCaptureReportAck.
func (m Message) ExecType() (*field.ExecTypeField, quickfix.MessageRejectError) {
	f := &field.ExecTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecType reads a ExecType from TradeCaptureReportAck.
func (m Message) GetExecType(f *field.ExecTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m Message) TradeReportRefID() (*field.TradeReportRefIDField, quickfix.MessageRejectError) {
	f := &field.TradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRefID reads a TradeReportRefID from TradeCaptureReportAck.
func (m Message) GetTradeReportRefID(f *field.TradeReportRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryTradeReportRefID() (*field.SecondaryTradeReportRefIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeReportRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportRefID reads a SecondaryTradeReportRefID from TradeCaptureReportAck.
func (m Message) GetSecondaryTradeReportRefID(f *field.SecondaryTradeReportRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdRptStatus is a non-required field for TradeCaptureReportAck.
func (m Message) TrdRptStatus() (*field.TrdRptStatusField, quickfix.MessageRejectError) {
	f := &field.TrdRptStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdRptStatus reads a TrdRptStatus from TradeCaptureReportAck.
func (m Message) GetTrdRptStatus(f *field.TrdRptStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
func (m Message) TradeReportRejectReason() (*field.TradeReportRejectReasonField, quickfix.MessageRejectError) {
	f := &field.TradeReportRejectReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeReportRejectReason reads a TradeReportRejectReason from TradeCaptureReportAck.
func (m Message) GetTradeReportRejectReason(f *field.TradeReportRejectReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryTradeReportID() (*field.SecondaryTradeReportIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeReportID reads a SecondaryTradeReportID from TradeCaptureReportAck.
func (m Message) GetSecondaryTradeReportID(f *field.SecondaryTradeReportIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
func (m Message) SubscriptionRequestType() (*field.SubscriptionRequestTypeField, quickfix.MessageRejectError) {
	f := &field.SubscriptionRequestTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from TradeCaptureReportAck.
func (m Message) GetSubscriptionRequestType(f *field.SubscriptionRequestTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLinkID is a non-required field for TradeCaptureReportAck.
func (m Message) TradeLinkID() (*field.TradeLinkIDField, quickfix.MessageRejectError) {
	f := &field.TradeLinkIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLinkID reads a TradeLinkID from TradeCaptureReportAck.
func (m Message) GetTradeLinkID(f *field.TradeLinkIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TrdMatchID is a non-required field for TradeCaptureReportAck.
func (m Message) TrdMatchID() (*field.TrdMatchIDField, quickfix.MessageRejectError) {
	f := &field.TrdMatchIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTrdMatchID reads a TrdMatchID from TradeCaptureReportAck.
func (m Message) GetTrdMatchID(f *field.TrdMatchIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecID is a non-required field for TradeCaptureReportAck.
func (m Message) ExecID() (*field.ExecIDField, quickfix.MessageRejectError) {
	f := &field.ExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecID reads a ExecID from TradeCaptureReportAck.
func (m Message) GetExecID(f *field.ExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryExecID is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryExecID() (*field.SecondaryExecIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryExecIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryExecID reads a SecondaryExecID from TradeCaptureReportAck.
func (m Message) GetSecondaryExecID(f *field.SecondaryExecIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for TradeCaptureReportAck.
func (m Message) Symbol() (*field.SymbolField, quickfix.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from TradeCaptureReportAck.
func (m Message) GetSymbol(f *field.SymbolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for TradeCaptureReportAck.
func (m Message) SymbolSfx() (*field.SymbolSfxField, quickfix.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from TradeCaptureReportAck.
func (m Message) GetSymbolSfx(f *field.SymbolSfxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityID() (*field.SecurityIDField, quickfix.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from TradeCaptureReportAck.
func (m Message) GetSecurityID(f *field.SecurityIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityIDSource() (*field.SecurityIDSourceField, quickfix.MessageRejectError) {
	f := &field.SecurityIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from TradeCaptureReportAck.
func (m Message) GetSecurityIDSource(f *field.SecurityIDSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for TradeCaptureReportAck.
func (m Message) NoSecurityAltID() (*field.NoSecurityAltIDField, quickfix.MessageRejectError) {
	f := &field.NoSecurityAltIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from TradeCaptureReportAck.
func (m Message) GetNoSecurityAltID(f *field.NoSecurityAltIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for TradeCaptureReportAck.
func (m Message) Product() (*field.ProductField, quickfix.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from TradeCaptureReportAck.
func (m Message) GetProduct(f *field.ProductField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for TradeCaptureReportAck.
func (m Message) CFICode() (*field.CFICodeField, quickfix.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from TradeCaptureReportAck.
func (m Message) GetCFICode(f *field.CFICodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityType() (*field.SecurityTypeField, quickfix.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from TradeCaptureReportAck.
func (m Message) GetSecurityType(f *field.SecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecuritySubType is a non-required field for TradeCaptureReportAck.
func (m Message) SecuritySubType() (*field.SecuritySubTypeField, quickfix.MessageRejectError) {
	f := &field.SecuritySubTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecuritySubType reads a SecuritySubType from TradeCaptureReportAck.
func (m Message) GetSecuritySubType(f *field.SecuritySubTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for TradeCaptureReportAck.
func (m Message) MaturityMonthYear() (*field.MaturityMonthYearField, quickfix.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from TradeCaptureReportAck.
func (m Message) GetMaturityMonthYear(f *field.MaturityMonthYearField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for TradeCaptureReportAck.
func (m Message) MaturityDate() (*field.MaturityDateField, quickfix.MessageRejectError) {
	f := &field.MaturityDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from TradeCaptureReportAck.
func (m Message) GetMaturityDate(f *field.MaturityDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for TradeCaptureReportAck.
func (m Message) CouponPaymentDate() (*field.CouponPaymentDateField, quickfix.MessageRejectError) {
	f := &field.CouponPaymentDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from TradeCaptureReportAck.
func (m Message) GetCouponPaymentDate(f *field.CouponPaymentDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for TradeCaptureReportAck.
func (m Message) IssueDate() (*field.IssueDateField, quickfix.MessageRejectError) {
	f := &field.IssueDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from TradeCaptureReportAck.
func (m Message) GetIssueDate(f *field.IssueDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportAck.
func (m Message) RepoCollateralSecurityType() (*field.RepoCollateralSecurityTypeField, quickfix.MessageRejectError) {
	f := &field.RepoCollateralSecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from TradeCaptureReportAck.
func (m Message) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for TradeCaptureReportAck.
func (m Message) RepurchaseTerm() (*field.RepurchaseTermField, quickfix.MessageRejectError) {
	f := &field.RepurchaseTermField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from TradeCaptureReportAck.
func (m Message) GetRepurchaseTerm(f *field.RepurchaseTermField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for TradeCaptureReportAck.
func (m Message) RepurchaseRate() (*field.RepurchaseRateField, quickfix.MessageRejectError) {
	f := &field.RepurchaseRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from TradeCaptureReportAck.
func (m Message) GetRepurchaseRate(f *field.RepurchaseRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for TradeCaptureReportAck.
func (m Message) Factor() (*field.FactorField, quickfix.MessageRejectError) {
	f := &field.FactorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from TradeCaptureReportAck.
func (m Message) GetFactor(f *field.FactorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for TradeCaptureReportAck.
func (m Message) CreditRating() (*field.CreditRatingField, quickfix.MessageRejectError) {
	f := &field.CreditRatingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from TradeCaptureReportAck.
func (m Message) GetCreditRating(f *field.CreditRatingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for TradeCaptureReportAck.
func (m Message) InstrRegistry() (*field.InstrRegistryField, quickfix.MessageRejectError) {
	f := &field.InstrRegistryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from TradeCaptureReportAck.
func (m Message) GetInstrRegistry(f *field.InstrRegistryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for TradeCaptureReportAck.
func (m Message) CountryOfIssue() (*field.CountryOfIssueField, quickfix.MessageRejectError) {
	f := &field.CountryOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from TradeCaptureReportAck.
func (m Message) GetCountryOfIssue(f *field.CountryOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportAck.
func (m Message) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssueField, quickfix.MessageRejectError) {
	f := &field.StateOrProvinceOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from TradeCaptureReportAck.
func (m Message) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for TradeCaptureReportAck.
func (m Message) LocaleOfIssue() (*field.LocaleOfIssueField, quickfix.MessageRejectError) {
	f := &field.LocaleOfIssueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from TradeCaptureReportAck.
func (m Message) GetLocaleOfIssue(f *field.LocaleOfIssueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for TradeCaptureReportAck.
func (m Message) RedemptionDate() (*field.RedemptionDateField, quickfix.MessageRejectError) {
	f := &field.RedemptionDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from TradeCaptureReportAck.
func (m Message) GetRedemptionDate(f *field.RedemptionDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for TradeCaptureReportAck.
func (m Message) StrikePrice() (*field.StrikePriceField, quickfix.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from TradeCaptureReportAck.
func (m Message) GetStrikePrice(f *field.StrikePriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeCurrency is a non-required field for TradeCaptureReportAck.
func (m Message) StrikeCurrency() (*field.StrikeCurrencyField, quickfix.MessageRejectError) {
	f := &field.StrikeCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeCurrency reads a StrikeCurrency from TradeCaptureReportAck.
func (m Message) GetStrikeCurrency(f *field.StrikeCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for TradeCaptureReportAck.
func (m Message) OptAttribute() (*field.OptAttributeField, quickfix.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from TradeCaptureReportAck.
func (m Message) GetOptAttribute(f *field.OptAttributeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for TradeCaptureReportAck.
func (m Message) ContractMultiplier() (*field.ContractMultiplierField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from TradeCaptureReportAck.
func (m Message) GetContractMultiplier(f *field.ContractMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for TradeCaptureReportAck.
func (m Message) CouponRate() (*field.CouponRateField, quickfix.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from TradeCaptureReportAck.
func (m Message) GetCouponRate(f *field.CouponRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityExchange() (*field.SecurityExchangeField, quickfix.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from TradeCaptureReportAck.
func (m Message) GetSecurityExchange(f *field.SecurityExchangeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for TradeCaptureReportAck.
func (m Message) Issuer() (*field.IssuerField, quickfix.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from TradeCaptureReportAck.
func (m Message) GetIssuer(f *field.IssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedIssuerLen() (*field.EncodedIssuerLenField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from TradeCaptureReportAck.
func (m Message) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedIssuer() (*field.EncodedIssuerField, quickfix.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from TradeCaptureReportAck.
func (m Message) GetEncodedIssuer(f *field.EncodedIssuerField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityDesc() (*field.SecurityDescField, quickfix.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from TradeCaptureReportAck.
func (m Message) GetSecurityDesc(f *field.SecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from TradeCaptureReportAck.
func (m Message) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedSecurityDesc() (*field.EncodedSecurityDescField, quickfix.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from TradeCaptureReportAck.
func (m Message) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Pool is a non-required field for TradeCaptureReportAck.
func (m Message) Pool() (*field.PoolField, quickfix.MessageRejectError) {
	f := &field.PoolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPool reads a Pool from TradeCaptureReportAck.
func (m Message) GetPool(f *field.PoolField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractSettlMonth is a non-required field for TradeCaptureReportAck.
func (m Message) ContractSettlMonth() (*field.ContractSettlMonthField, quickfix.MessageRejectError) {
	f := &field.ContractSettlMonthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractSettlMonth reads a ContractSettlMonth from TradeCaptureReportAck.
func (m Message) GetContractSettlMonth(f *field.ContractSettlMonthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPProgram is a non-required field for TradeCaptureReportAck.
func (m Message) CPProgram() (*field.CPProgramField, quickfix.MessageRejectError) {
	f := &field.CPProgramField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPProgram reads a CPProgram from TradeCaptureReportAck.
func (m Message) GetCPProgram(f *field.CPProgramField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CPRegType is a non-required field for TradeCaptureReportAck.
func (m Message) CPRegType() (*field.CPRegTypeField, quickfix.MessageRejectError) {
	f := &field.CPRegTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCPRegType reads a CPRegType from TradeCaptureReportAck.
func (m Message) GetCPRegType(f *field.CPRegTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoEvents is a non-required field for TradeCaptureReportAck.
func (m Message) NoEvents() (*field.NoEventsField, quickfix.MessageRejectError) {
	f := &field.NoEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoEvents reads a NoEvents from TradeCaptureReportAck.
func (m Message) GetNoEvents(f *field.NoEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DatedDate is a non-required field for TradeCaptureReportAck.
func (m Message) DatedDate() (*field.DatedDateField, quickfix.MessageRejectError) {
	f := &field.DatedDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDatedDate reads a DatedDate from TradeCaptureReportAck.
func (m Message) GetDatedDate(f *field.DatedDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InterestAccrualDate is a non-required field for TradeCaptureReportAck.
func (m Message) InterestAccrualDate() (*field.InterestAccrualDateField, quickfix.MessageRejectError) {
	f := &field.InterestAccrualDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInterestAccrualDate reads a InterestAccrualDate from TradeCaptureReportAck.
func (m Message) GetInterestAccrualDate(f *field.InterestAccrualDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityStatus is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityStatus() (*field.SecurityStatusField, quickfix.MessageRejectError) {
	f := &field.SecurityStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityStatus reads a SecurityStatus from TradeCaptureReportAck.
func (m Message) GetSecurityStatus(f *field.SecurityStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportAck.
func (m Message) SettleOnOpenFlag() (*field.SettleOnOpenFlagField, quickfix.MessageRejectError) {
	f := &field.SettleOnOpenFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettleOnOpenFlag reads a SettleOnOpenFlag from TradeCaptureReportAck.
func (m Message) GetSettleOnOpenFlag(f *field.SettleOnOpenFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportAck.
func (m Message) InstrmtAssignmentMethod() (*field.InstrmtAssignmentMethodField, quickfix.MessageRejectError) {
	f := &field.InstrmtAssignmentMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetInstrmtAssignmentMethod reads a InstrmtAssignmentMethod from TradeCaptureReportAck.
func (m Message) GetInstrmtAssignmentMethod(f *field.InstrmtAssignmentMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeMultiplier is a non-required field for TradeCaptureReportAck.
func (m Message) StrikeMultiplier() (*field.StrikeMultiplierField, quickfix.MessageRejectError) {
	f := &field.StrikeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeMultiplier reads a StrikeMultiplier from TradeCaptureReportAck.
func (m Message) GetStrikeMultiplier(f *field.StrikeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikeValue is a non-required field for TradeCaptureReportAck.
func (m Message) StrikeValue() (*field.StrikeValueField, quickfix.MessageRejectError) {
	f := &field.StrikeValueField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikeValue reads a StrikeValue from TradeCaptureReportAck.
func (m Message) GetStrikeValue(f *field.StrikeValueField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrement is a non-required field for TradeCaptureReportAck.
func (m Message) MinPriceIncrement() (*field.MinPriceIncrementField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrement reads a MinPriceIncrement from TradeCaptureReportAck.
func (m Message) GetMinPriceIncrement(f *field.MinPriceIncrementField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PositionLimit is a non-required field for TradeCaptureReportAck.
func (m Message) PositionLimit() (*field.PositionLimitField, quickfix.MessageRejectError) {
	f := &field.PositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPositionLimit reads a PositionLimit from TradeCaptureReportAck.
func (m Message) GetPositionLimit(f *field.PositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NTPositionLimit is a non-required field for TradeCaptureReportAck.
func (m Message) NTPositionLimit() (*field.NTPositionLimitField, quickfix.MessageRejectError) {
	f := &field.NTPositionLimitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNTPositionLimit reads a NTPositionLimit from TradeCaptureReportAck.
func (m Message) GetNTPositionLimit(f *field.NTPositionLimitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoInstrumentParties is a non-required field for TradeCaptureReportAck.
func (m Message) NoInstrumentParties() (*field.NoInstrumentPartiesField, quickfix.MessageRejectError) {
	f := &field.NoInstrumentPartiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoInstrumentParties reads a NoInstrumentParties from TradeCaptureReportAck.
func (m Message) GetNoInstrumentParties(f *field.NoInstrumentPartiesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m Message) UnitOfMeasure() (*field.UnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasure reads a UnitOfMeasure from TradeCaptureReportAck.
func (m Message) GetUnitOfMeasure(f *field.UnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TimeUnit is a non-required field for TradeCaptureReportAck.
func (m Message) TimeUnit() (*field.TimeUnitField, quickfix.MessageRejectError) {
	f := &field.TimeUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTimeUnit reads a TimeUnit from TradeCaptureReportAck.
func (m Message) GetTimeUnit(f *field.TimeUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityTime is a non-required field for TradeCaptureReportAck.
func (m Message) MaturityTime() (*field.MaturityTimeField, quickfix.MessageRejectError) {
	f := &field.MaturityTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityTime reads a MaturityTime from TradeCaptureReportAck.
func (m Message) GetMaturityTime(f *field.MaturityTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityGroup is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityGroup() (*field.SecurityGroupField, quickfix.MessageRejectError) {
	f := &field.SecurityGroupField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityGroup reads a SecurityGroup from TradeCaptureReportAck.
func (m Message) GetSecurityGroup(f *field.SecurityGroupField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MinPriceIncrementAmount is a non-required field for TradeCaptureReportAck.
func (m Message) MinPriceIncrementAmount() (*field.MinPriceIncrementAmountField, quickfix.MessageRejectError) {
	f := &field.MinPriceIncrementAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMinPriceIncrementAmount reads a MinPriceIncrementAmount from TradeCaptureReportAck.
func (m Message) GetMinPriceIncrementAmount(f *field.MinPriceIncrementAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m Message) UnitOfMeasureQty() (*field.UnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.UnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnitOfMeasureQty reads a UnitOfMeasureQty from TradeCaptureReportAck.
func (m Message) GetUnitOfMeasureQty(f *field.UnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLLen is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityXMLLen() (*field.SecurityXMLLenField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLLen reads a SecurityXMLLen from TradeCaptureReportAck.
func (m Message) GetSecurityXMLLen(f *field.SecurityXMLLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXML is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityXML() (*field.SecurityXMLField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXML reads a SecurityXML from TradeCaptureReportAck.
func (m Message) GetSecurityXML(f *field.SecurityXMLField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityXMLSchema is a non-required field for TradeCaptureReportAck.
func (m Message) SecurityXMLSchema() (*field.SecurityXMLSchemaField, quickfix.MessageRejectError) {
	f := &field.SecurityXMLSchemaField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityXMLSchema reads a SecurityXMLSchema from TradeCaptureReportAck.
func (m Message) GetSecurityXMLSchema(f *field.SecurityXMLSchemaField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ProductComplex is a non-required field for TradeCaptureReportAck.
func (m Message) ProductComplex() (*field.ProductComplexField, quickfix.MessageRejectError) {
	f := &field.ProductComplexField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProductComplex reads a ProductComplex from TradeCaptureReportAck.
func (m Message) GetProductComplex(f *field.ProductComplexField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasure is a non-required field for TradeCaptureReportAck.
func (m Message) PriceUnitOfMeasure() (*field.PriceUnitOfMeasureField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasure reads a PriceUnitOfMeasure from TradeCaptureReportAck.
func (m Message) GetPriceUnitOfMeasure(f *field.PriceUnitOfMeasureField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceUnitOfMeasureQty is a non-required field for TradeCaptureReportAck.
func (m Message) PriceUnitOfMeasureQty() (*field.PriceUnitOfMeasureQtyField, quickfix.MessageRejectError) {
	f := &field.PriceUnitOfMeasureQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceUnitOfMeasureQty reads a PriceUnitOfMeasureQty from TradeCaptureReportAck.
func (m Message) GetPriceUnitOfMeasureQty(f *field.PriceUnitOfMeasureQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlMethod is a non-required field for TradeCaptureReportAck.
func (m Message) SettlMethod() (*field.SettlMethodField, quickfix.MessageRejectError) {
	f := &field.SettlMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlMethod reads a SettlMethod from TradeCaptureReportAck.
func (m Message) GetSettlMethod(f *field.SettlMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExerciseStyle is a non-required field for TradeCaptureReportAck.
func (m Message) ExerciseStyle() (*field.ExerciseStyleField, quickfix.MessageRejectError) {
	f := &field.ExerciseStyleField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExerciseStyle reads a ExerciseStyle from TradeCaptureReportAck.
func (m Message) GetExerciseStyle(f *field.ExerciseStyleField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutAmount is a non-required field for TradeCaptureReportAck.
func (m Message) OptPayoutAmount() (*field.OptPayoutAmountField, quickfix.MessageRejectError) {
	f := &field.OptPayoutAmountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutAmount reads a OptPayoutAmount from TradeCaptureReportAck.
func (m Message) GetOptPayoutAmount(f *field.OptPayoutAmountField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceQuoteMethod is a non-required field for TradeCaptureReportAck.
func (m Message) PriceQuoteMethod() (*field.PriceQuoteMethodField, quickfix.MessageRejectError) {
	f := &field.PriceQuoteMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceQuoteMethod reads a PriceQuoteMethod from TradeCaptureReportAck.
func (m Message) GetPriceQuoteMethod(f *field.PriceQuoteMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ListMethod is a non-required field for TradeCaptureReportAck.
func (m Message) ListMethod() (*field.ListMethodField, quickfix.MessageRejectError) {
	f := &field.ListMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListMethod reads a ListMethod from TradeCaptureReportAck.
func (m Message) GetListMethod(f *field.ListMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CapPrice is a non-required field for TradeCaptureReportAck.
func (m Message) CapPrice() (*field.CapPriceField, quickfix.MessageRejectError) {
	f := &field.CapPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCapPrice reads a CapPrice from TradeCaptureReportAck.
func (m Message) GetCapPrice(f *field.CapPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FloorPrice is a non-required field for TradeCaptureReportAck.
func (m Message) FloorPrice() (*field.FloorPriceField, quickfix.MessageRejectError) {
	f := &field.FloorPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFloorPrice reads a FloorPrice from TradeCaptureReportAck.
func (m Message) GetFloorPrice(f *field.FloorPriceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for TradeCaptureReportAck.
func (m Message) PutOrCall() (*field.PutOrCallField, quickfix.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from TradeCaptureReportAck.
func (m Message) GetPutOrCall(f *field.PutOrCallField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexibleIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) FlexibleIndicator() (*field.FlexibleIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexibleIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexibleIndicator reads a FlexibleIndicator from TradeCaptureReportAck.
func (m Message) GetFlexibleIndicator(f *field.FlexibleIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlexProductEligibilityIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) FlexProductEligibilityIndicator() (*field.FlexProductEligibilityIndicatorField, quickfix.MessageRejectError) {
	f := &field.FlexProductEligibilityIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlexProductEligibilityIndicator reads a FlexProductEligibilityIndicator from TradeCaptureReportAck.
func (m Message) GetFlexProductEligibilityIndicator(f *field.FlexProductEligibilityIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ValuationMethod is a non-required field for TradeCaptureReportAck.
func (m Message) ValuationMethod() (*field.ValuationMethodField, quickfix.MessageRejectError) {
	f := &field.ValuationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetValuationMethod reads a ValuationMethod from TradeCaptureReportAck.
func (m Message) GetValuationMethod(f *field.ValuationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplierUnit is a non-required field for TradeCaptureReportAck.
func (m Message) ContractMultiplierUnit() (*field.ContractMultiplierUnitField, quickfix.MessageRejectError) {
	f := &field.ContractMultiplierUnitField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplierUnit reads a ContractMultiplierUnit from TradeCaptureReportAck.
func (m Message) GetContractMultiplierUnit(f *field.ContractMultiplierUnitField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FlowScheduleType is a non-required field for TradeCaptureReportAck.
func (m Message) FlowScheduleType() (*field.FlowScheduleTypeField, quickfix.MessageRejectError) {
	f := &field.FlowScheduleTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFlowScheduleType reads a FlowScheduleType from TradeCaptureReportAck.
func (m Message) GetFlowScheduleType(f *field.FlowScheduleTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RestructuringType is a non-required field for TradeCaptureReportAck.
func (m Message) RestructuringType() (*field.RestructuringTypeField, quickfix.MessageRejectError) {
	f := &field.RestructuringTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRestructuringType reads a RestructuringType from TradeCaptureReportAck.
func (m Message) GetRestructuringType(f *field.RestructuringTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Seniority is a non-required field for TradeCaptureReportAck.
func (m Message) Seniority() (*field.SeniorityField, quickfix.MessageRejectError) {
	f := &field.SeniorityField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSeniority reads a Seniority from TradeCaptureReportAck.
func (m Message) GetSeniority(f *field.SeniorityField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NotionalPercentageOutstanding is a non-required field for TradeCaptureReportAck.
func (m Message) NotionalPercentageOutstanding() (*field.NotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.NotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNotionalPercentageOutstanding reads a NotionalPercentageOutstanding from TradeCaptureReportAck.
func (m Message) GetNotionalPercentageOutstanding(f *field.NotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OriginalNotionalPercentageOutstanding is a non-required field for TradeCaptureReportAck.
func (m Message) OriginalNotionalPercentageOutstanding() (*field.OriginalNotionalPercentageOutstandingField, quickfix.MessageRejectError) {
	f := &field.OriginalNotionalPercentageOutstandingField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOriginalNotionalPercentageOutstanding reads a OriginalNotionalPercentageOutstanding from TradeCaptureReportAck.
func (m Message) GetOriginalNotionalPercentageOutstanding(f *field.OriginalNotionalPercentageOutstandingField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AttachmentPoint is a non-required field for TradeCaptureReportAck.
func (m Message) AttachmentPoint() (*field.AttachmentPointField, quickfix.MessageRejectError) {
	f := &field.AttachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAttachmentPoint reads a AttachmentPoint from TradeCaptureReportAck.
func (m Message) GetAttachmentPoint(f *field.AttachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//DetachmentPoint is a non-required field for TradeCaptureReportAck.
func (m Message) DetachmentPoint() (*field.DetachmentPointField, quickfix.MessageRejectError) {
	f := &field.DetachmentPointField{}
	err := m.Body.Get(f)
	return f, err
}

//GetDetachmentPoint reads a DetachmentPoint from TradeCaptureReportAck.
func (m Message) GetDetachmentPoint(f *field.DetachmentPointField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceDeterminationMethod is a non-required field for TradeCaptureReportAck.
func (m Message) StrikePriceDeterminationMethod() (*field.StrikePriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceDeterminationMethod reads a StrikePriceDeterminationMethod from TradeCaptureReportAck.
func (m Message) GetStrikePriceDeterminationMethod(f *field.StrikePriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryMethod is a non-required field for TradeCaptureReportAck.
func (m Message) StrikePriceBoundaryMethod() (*field.StrikePriceBoundaryMethodField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryMethod reads a StrikePriceBoundaryMethod from TradeCaptureReportAck.
func (m Message) GetStrikePriceBoundaryMethod(f *field.StrikePriceBoundaryMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePriceBoundaryPrecision is a non-required field for TradeCaptureReportAck.
func (m Message) StrikePriceBoundaryPrecision() (*field.StrikePriceBoundaryPrecisionField, quickfix.MessageRejectError) {
	f := &field.StrikePriceBoundaryPrecisionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePriceBoundaryPrecision reads a StrikePriceBoundaryPrecision from TradeCaptureReportAck.
func (m Message) GetStrikePriceBoundaryPrecision(f *field.StrikePriceBoundaryPrecisionField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingPriceDeterminationMethod is a non-required field for TradeCaptureReportAck.
func (m Message) UnderlyingPriceDeterminationMethod() (*field.UnderlyingPriceDeterminationMethodField, quickfix.MessageRejectError) {
	f := &field.UnderlyingPriceDeterminationMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingPriceDeterminationMethod reads a UnderlyingPriceDeterminationMethod from TradeCaptureReportAck.
func (m Message) GetUnderlyingPriceDeterminationMethod(f *field.UnderlyingPriceDeterminationMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OptPayoutType is a non-required field for TradeCaptureReportAck.
func (m Message) OptPayoutType() (*field.OptPayoutTypeField, quickfix.MessageRejectError) {
	f := &field.OptPayoutTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptPayoutType reads a OptPayoutType from TradeCaptureReportAck.
func (m Message) GetOptPayoutType(f *field.OptPayoutTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoComplexEvents is a non-required field for TradeCaptureReportAck.
func (m Message) NoComplexEvents() (*field.NoComplexEventsField, quickfix.MessageRejectError) {
	f := &field.NoComplexEventsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoComplexEvents reads a NoComplexEvents from TradeCaptureReportAck.
func (m Message) GetNoComplexEvents(f *field.NoComplexEventsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for TradeCaptureReportAck.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from TradeCaptureReportAck.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRegTimestamps is a non-required field for TradeCaptureReportAck.
func (m Message) NoTrdRegTimestamps() (*field.NoTrdRegTimestampsField, quickfix.MessageRejectError) {
	f := &field.NoTrdRegTimestampsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRegTimestamps reads a NoTrdRegTimestamps from TradeCaptureReportAck.
func (m Message) GetNoTrdRegTimestamps(f *field.NoTrdRegTimestampsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseTransportType is a non-required field for TradeCaptureReportAck.
func (m Message) ResponseTransportType() (*field.ResponseTransportTypeField, quickfix.MessageRejectError) {
	f := &field.ResponseTransportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseTransportType reads a ResponseTransportType from TradeCaptureReportAck.
func (m Message) GetResponseTransportType(f *field.ResponseTransportTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ResponseDestination is a non-required field for TradeCaptureReportAck.
func (m Message) ResponseDestination() (*field.ResponseDestinationField, quickfix.MessageRejectError) {
	f := &field.ResponseDestinationField{}
	err := m.Body.Get(f)
	return f, err
}

//GetResponseDestination reads a ResponseDestination from TradeCaptureReportAck.
func (m Message) GetResponseDestination(f *field.ResponseDestinationField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradeCaptureReportAck.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradeCaptureReportAck.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradeCaptureReportAck.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradeCaptureReportAck.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradeCaptureReportAck.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for TradeCaptureReportAck.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from TradeCaptureReportAck.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) ClearingFeeIndicator() (*field.ClearingFeeIndicatorField, quickfix.MessageRejectError) {
	f := &field.ClearingFeeIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingFeeIndicator reads a ClearingFeeIndicator from TradeCaptureReportAck.
func (m Message) GetClearingFeeIndicator(f *field.ClearingFeeIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ExecRestatementReason is a non-required field for TradeCaptureReportAck.
func (m Message) ExecRestatementReason() (*field.ExecRestatementReasonField, quickfix.MessageRejectError) {
	f := &field.ExecRestatementReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExecRestatementReason reads a ExecRestatementReason from TradeCaptureReportAck.
func (m Message) GetExecRestatementReason(f *field.ExecRestatementReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PreviouslyReported is a non-required field for TradeCaptureReportAck.
func (m Message) PreviouslyReported() (*field.PreviouslyReportedField, quickfix.MessageRejectError) {
	f := &field.PreviouslyReportedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPreviouslyReported reads a PreviouslyReported from TradeCaptureReportAck.
func (m Message) GetPreviouslyReported(f *field.PreviouslyReportedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for TradeCaptureReportAck.
func (m Message) PriceType() (*field.PriceTypeField, quickfix.MessageRejectError) {
	f := &field.PriceTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from TradeCaptureReportAck.
func (m Message) GetPriceType(f *field.PriceTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionID is a non-required field for TradeCaptureReportAck.
func (m Message) UnderlyingTradingSessionID() (*field.UnderlyingTradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingTradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionID reads a UnderlyingTradingSessionID from TradeCaptureReportAck.
func (m Message) GetUnderlyingTradingSessionID(f *field.UnderlyingTradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//QtyType is a non-required field for TradeCaptureReportAck.
func (m Message) QtyType() (*field.QtyTypeField, quickfix.MessageRejectError) {
	f := &field.QtyTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetQtyType reads a QtyType from TradeCaptureReportAck.
func (m Message) GetQtyType(f *field.QtyTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnderlyingTradingSessionSubID is a non-required field for TradeCaptureReportAck.
func (m Message) UnderlyingTradingSessionSubID() (*field.UnderlyingTradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.UnderlyingTradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnderlyingTradingSessionSubID reads a UnderlyingTradingSessionSubID from TradeCaptureReportAck.
func (m Message) GetUnderlyingTradingSessionSubID(f *field.UnderlyingTradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastQty is a non-required field for TradeCaptureReportAck.
func (m Message) LastQty() (*field.LastQtyField, quickfix.MessageRejectError) {
	f := &field.LastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastQty reads a LastQty from TradeCaptureReportAck.
func (m Message) GetLastQty(f *field.LastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastPx is a non-required field for TradeCaptureReportAck.
func (m Message) LastPx() (*field.LastPxField, quickfix.MessageRejectError) {
	f := &field.LastPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastPx reads a LastPx from TradeCaptureReportAck.
func (m Message) GetLastPx(f *field.LastPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastParPx is a non-required field for TradeCaptureReportAck.
func (m Message) LastParPx() (*field.LastParPxField, quickfix.MessageRejectError) {
	f := &field.LastParPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastParPx reads a LastParPx from TradeCaptureReportAck.
func (m Message) GetLastParPx(f *field.LastParPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastSpotRate is a non-required field for TradeCaptureReportAck.
func (m Message) LastSpotRate() (*field.LastSpotRateField, quickfix.MessageRejectError) {
	f := &field.LastSpotRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSpotRate reads a LastSpotRate from TradeCaptureReportAck.
func (m Message) GetLastSpotRate(f *field.LastSpotRateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastForwardPoints is a non-required field for TradeCaptureReportAck.
func (m Message) LastForwardPoints() (*field.LastForwardPointsField, quickfix.MessageRejectError) {
	f := &field.LastForwardPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastForwardPoints reads a LastForwardPoints from TradeCaptureReportAck.
func (m Message) GetLastForwardPoints(f *field.LastForwardPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for TradeCaptureReportAck.
func (m Message) LastMkt() (*field.LastMktField, quickfix.MessageRejectError) {
	f := &field.LastMktField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from TradeCaptureReportAck.
func (m Message) GetLastMkt(f *field.LastMktField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a non-required field for TradeCaptureReportAck.
func (m Message) TradeDate() (*field.TradeDateField, quickfix.MessageRejectError) {
	f := &field.TradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from TradeCaptureReportAck.
func (m Message) GetTradeDate(f *field.TradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportAck.
func (m Message) ClearingBusinessDate() (*field.ClearingBusinessDateField, quickfix.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from TradeCaptureReportAck.
func (m Message) GetClearingBusinessDate(f *field.ClearingBusinessDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a non-required field for TradeCaptureReportAck.
func (m Message) AvgPx() (*field.AvgPxField, quickfix.MessageRejectError) {
	f := &field.AvgPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from TradeCaptureReportAck.
func (m Message) GetAvgPx(f *field.AvgPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPxIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) AvgPxIndicator() (*field.AvgPxIndicatorField, quickfix.MessageRejectError) {
	f := &field.AvgPxIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPxIndicator reads a AvgPxIndicator from TradeCaptureReportAck.
func (m Message) GetAvgPxIndicator(f *field.AvgPxIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MultiLegReportingType is a non-required field for TradeCaptureReportAck.
func (m Message) MultiLegReportingType() (*field.MultiLegReportingTypeField, quickfix.MessageRejectError) {
	f := &field.MultiLegReportingTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMultiLegReportingType reads a MultiLegReportingType from TradeCaptureReportAck.
func (m Message) GetMultiLegReportingType(f *field.MultiLegReportingTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeLegRefID is a non-required field for TradeCaptureReportAck.
func (m Message) TradeLegRefID() (*field.TradeLegRefIDField, quickfix.MessageRejectError) {
	f := &field.TradeLegRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeLegRefID reads a TradeLegRefID from TradeCaptureReportAck.
func (m Message) GetTradeLegRefID(f *field.TradeLegRefIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlType is a non-required field for TradeCaptureReportAck.
func (m Message) SettlType() (*field.SettlTypeField, quickfix.MessageRejectError) {
	f := &field.SettlTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlType reads a SettlType from TradeCaptureReportAck.
func (m Message) GetSettlType(f *field.SettlTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchStatus is a non-required field for TradeCaptureReportAck.
func (m Message) MatchStatus() (*field.MatchStatusField, quickfix.MessageRejectError) {
	f := &field.MatchStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchStatus reads a MatchStatus from TradeCaptureReportAck.
func (m Message) GetMatchStatus(f *field.MatchStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MatchType is a non-required field for TradeCaptureReportAck.
func (m Message) MatchType() (*field.MatchTypeField, quickfix.MessageRejectError) {
	f := &field.MatchTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMatchType reads a MatchType from TradeCaptureReportAck.
func (m Message) GetMatchType(f *field.MatchTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CopyMsgIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) CopyMsgIndicator() (*field.CopyMsgIndicatorField, quickfix.MessageRejectError) {
	f := &field.CopyMsgIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCopyMsgIndicator reads a CopyMsgIndicator from TradeCaptureReportAck.
func (m Message) GetCopyMsgIndicator(f *field.CopyMsgIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//PublishTrdIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) PublishTrdIndicator() (*field.PublishTrdIndicatorField, quickfix.MessageRejectError) {
	f := &field.PublishTrdIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPublishTrdIndicator reads a PublishTrdIndicator from TradeCaptureReportAck.
func (m Message) GetPublishTrdIndicator(f *field.PublishTrdIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ShortSaleReason is a non-required field for TradeCaptureReportAck.
func (m Message) ShortSaleReason() (*field.ShortSaleReasonField, quickfix.MessageRejectError) {
	f := &field.ShortSaleReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetShortSaleReason reads a ShortSaleReason from TradeCaptureReportAck.
func (m Message) GetShortSaleReason(f *field.ShortSaleReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlDate is a non-required field for TradeCaptureReportAck.
func (m Message) SettlDate() (*field.SettlDateField, quickfix.MessageRejectError) {
	f := &field.SettlDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlDate reads a SettlDate from TradeCaptureReportAck.
func (m Message) GetSettlDate(f *field.SettlDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for TradeCaptureReportAck.
func (m Message) SettlSessID() (*field.SettlSessIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from TradeCaptureReportAck.
func (m Message) GetSettlSessID(f *field.SettlSessIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessSubID is a non-required field for TradeCaptureReportAck.
func (m Message) SettlSessSubID() (*field.SettlSessSubIDField, quickfix.MessageRejectError) {
	f := &field.SettlSessSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessSubID reads a SettlSessSubID from TradeCaptureReportAck.
func (m Message) GetSettlSessSubID(f *field.SettlSessSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoPosAmt is a non-required field for TradeCaptureReportAck.
func (m Message) NoPosAmt() (*field.NoPosAmtField, quickfix.MessageRejectError) {
	f := &field.NoPosAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPosAmt reads a NoPosAmt from TradeCaptureReportAck.
func (m Message) GetNoPosAmt(f *field.NoPosAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TierCode is a non-required field for TradeCaptureReportAck.
func (m Message) TierCode() (*field.TierCodeField, quickfix.MessageRejectError) {
	f := &field.TierCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTierCode reads a TierCode from TradeCaptureReportAck.
func (m Message) GetTierCode(f *field.TierCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageEventSource is a non-required field for TradeCaptureReportAck.
func (m Message) MessageEventSource() (*field.MessageEventSourceField, quickfix.MessageRejectError) {
	f := &field.MessageEventSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMessageEventSource reads a MessageEventSource from TradeCaptureReportAck.
func (m Message) GetMessageEventSource(f *field.MessageEventSourceField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for TradeCaptureReportAck.
func (m Message) LastUpdateTime() (*field.LastUpdateTimeField, quickfix.MessageRejectError) {
	f := &field.LastUpdateTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from TradeCaptureReportAck.
func (m Message) GetLastUpdateTime(f *field.LastUpdateTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RndPx is a non-required field for TradeCaptureReportAck.
func (m Message) RndPx() (*field.RndPxField, quickfix.MessageRejectError) {
	f := &field.RndPxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRndPx reads a RndPx from TradeCaptureReportAck.
func (m Message) GetRndPx(f *field.RndPxField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSides is a required field for TradeCaptureReportAck.
func (m Message) NoSides() (*field.NoSidesField, quickfix.MessageRejectError) {
	f := &field.NoSidesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSides reads a NoSides from TradeCaptureReportAck.
func (m Message) GetNoSides(f *field.NoSidesField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//AsOfIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) AsOfIndicator() (*field.AsOfIndicatorField, quickfix.MessageRejectError) {
	f := &field.AsOfIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAsOfIndicator reads a AsOfIndicator from TradeCaptureReportAck.
func (m Message) GetAsOfIndicator(f *field.AsOfIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeID is a non-required field for TradeCaptureReportAck.
func (m Message) TradeID() (*field.TradeIDField, quickfix.MessageRejectError) {
	f := &field.TradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeID reads a TradeID from TradeCaptureReportAck.
func (m Message) GetTradeID(f *field.TradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryTradeID() (*field.SecondaryTradeIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryTradeID reads a SecondaryTradeID from TradeCaptureReportAck.
func (m Message) GetSecondaryTradeID(f *field.SecondaryTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FirmTradeID is a non-required field for TradeCaptureReportAck.
func (m Message) FirmTradeID() (*field.FirmTradeIDField, quickfix.MessageRejectError) {
	f := &field.FirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFirmTradeID reads a FirmTradeID from TradeCaptureReportAck.
func (m Message) GetFirmTradeID(f *field.FirmTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportAck.
func (m Message) SecondaryFirmTradeID() (*field.SecondaryFirmTradeIDField, quickfix.MessageRejectError) {
	f := &field.SecondaryFirmTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecondaryFirmTradeID reads a SecondaryFirmTradeID from TradeCaptureReportAck.
func (m Message) GetSecondaryFirmTradeID(f *field.SecondaryFirmTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//CalculatedCcyLastQty is a non-required field for TradeCaptureReportAck.
func (m Message) CalculatedCcyLastQty() (*field.CalculatedCcyLastQtyField, quickfix.MessageRejectError) {
	f := &field.CalculatedCcyLastQtyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCalculatedCcyLastQty reads a CalculatedCcyLastQty from TradeCaptureReportAck.
func (m Message) GetCalculatedCcyLastQty(f *field.CalculatedCcyLastQtyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LastSwapPoints is a non-required field for TradeCaptureReportAck.
func (m Message) LastSwapPoints() (*field.LastSwapPointsField, quickfix.MessageRejectError) {
	f := &field.LastSwapPointsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastSwapPoints reads a LastSwapPoints from TradeCaptureReportAck.
func (m Message) GetLastSwapPoints(f *field.LastSwapPointsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for TradeCaptureReportAck.
func (m Message) GrossTradeAmt() (*field.GrossTradeAmtField, quickfix.MessageRejectError) {
	f := &field.GrossTradeAmtField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from TradeCaptureReportAck.
func (m Message) GetGrossTradeAmt(f *field.GrossTradeAmtField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRootPartyIDs is a non-required field for TradeCaptureReportAck.
func (m Message) NoRootPartyIDs() (*field.NoRootPartyIDsField, quickfix.MessageRejectError) {
	f := &field.NoRootPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRootPartyIDs reads a NoRootPartyIDs from TradeCaptureReportAck.
func (m Message) GetNoRootPartyIDs(f *field.NoRootPartyIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m Message) TradeHandlingInstr() (*field.TradeHandlingInstrField, quickfix.MessageRejectError) {
	f := &field.TradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradeHandlingInstr reads a TradeHandlingInstr from TradeCaptureReportAck.
func (m Message) GetTradeHandlingInstr(f *field.TradeHandlingInstrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeHandlingInstr is a non-required field for TradeCaptureReportAck.
func (m Message) OrigTradeHandlingInstr() (*field.OrigTradeHandlingInstrField, quickfix.MessageRejectError) {
	f := &field.OrigTradeHandlingInstrField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeHandlingInstr reads a OrigTradeHandlingInstr from TradeCaptureReportAck.
func (m Message) GetOrigTradeHandlingInstr(f *field.OrigTradeHandlingInstrField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeDate is a non-required field for TradeCaptureReportAck.
func (m Message) OrigTradeDate() (*field.OrigTradeDateField, quickfix.MessageRejectError) {
	f := &field.OrigTradeDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeDate reads a OrigTradeDate from TradeCaptureReportAck.
func (m Message) GetOrigTradeDate(f *field.OrigTradeDateField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTradeID is a non-required field for TradeCaptureReportAck.
func (m Message) OrigTradeID() (*field.OrigTradeIDField, quickfix.MessageRejectError) {
	f := &field.OrigTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTradeID reads a OrigTradeID from TradeCaptureReportAck.
func (m Message) GetOrigTradeID(f *field.OrigTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigSecondaryTradeID is a non-required field for TradeCaptureReportAck.
func (m Message) OrigSecondaryTradeID() (*field.OrigSecondaryTradeIDField, quickfix.MessageRejectError) {
	f := &field.OrigSecondaryTradeIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigSecondaryTradeID reads a OrigSecondaryTradeID from TradeCaptureReportAck.
func (m Message) GetOrigSecondaryTradeID(f *field.OrigSecondaryTradeIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for TradeCaptureReportAck.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from TradeCaptureReportAck.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RptSys is a non-required field for TradeCaptureReportAck.
func (m Message) RptSys() (*field.RptSysField, quickfix.MessageRejectError) {
	f := &field.RptSysField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRptSys reads a RptSys from TradeCaptureReportAck.
func (m Message) GetRptSys(f *field.RptSysField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for TradeCaptureReportAck.
func (m Message) Currency() (*field.CurrencyField, quickfix.MessageRejectError) {
	f := &field.CurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from TradeCaptureReportAck.
func (m Message) GetCurrency(f *field.CurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for TradeCaptureReportAck.
func (m Message) SettlCurrency() (*field.SettlCurrencyField, quickfix.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from TradeCaptureReportAck.
func (m Message) GetSettlCurrency(f *field.SettlCurrencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//FeeMultiplier is a non-required field for TradeCaptureReportAck.
func (m Message) FeeMultiplier() (*field.FeeMultiplierField, quickfix.MessageRejectError) {
	f := &field.FeeMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFeeMultiplier reads a FeeMultiplier from TradeCaptureReportAck.
func (m Message) GetFeeMultiplier(f *field.FeeMultiplierField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTrdRepIndicators is a non-required field for TradeCaptureReportAck.
func (m Message) NoTrdRepIndicators() (*field.NoTrdRepIndicatorsField, quickfix.MessageRejectError) {
	f := &field.NoTrdRepIndicatorsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTrdRepIndicators reads a NoTrdRepIndicators from TradeCaptureReportAck.
func (m Message) GetNoTrdRepIndicators(f *field.NoTrdRepIndicatorsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradePublishIndicator is a non-required field for TradeCaptureReportAck.
func (m Message) TradePublishIndicator() (*field.TradePublishIndicatorField, quickfix.MessageRejectError) {
	f := &field.TradePublishIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradePublishIndicator reads a TradePublishIndicator from TradeCaptureReportAck.
func (m Message) GetTradePublishIndicator(f *field.TradePublishIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//VenueType is a non-required field for TradeCaptureReportAck.
func (m Message) VenueType() (*field.VenueTypeField, quickfix.MessageRejectError) {
	f := &field.VenueTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetVenueType reads a VenueType from TradeCaptureReportAck.
func (m Message) GetVenueType(f *field.VenueTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for TradeCaptureReportAck.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from TradeCaptureReportAck.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for TradeCaptureReportAck.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from TradeCaptureReportAck.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for TradeCaptureReportAck.
func New(
	nosides *field.NoSidesField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("AR"))
	builder.Body.Set(nosides)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "AR", r
}
