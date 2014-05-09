package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//MarketDataSnapshotFullRefresh msg type = W.
type MarketDataSnapshotFullRefresh struct {
	message.Message
}

//MarketDataSnapshotFullRefreshBuilder builds MarketDataSnapshotFullRefresh messages.
type MarketDataSnapshotFullRefreshBuilder struct {
	message.MessageBuilder
}

//CreateMarketDataSnapshotFullRefreshBuilder returns an initialized MarketDataSnapshotFullRefreshBuilder with specified required fields.
func CreateMarketDataSnapshotFullRefreshBuilder(
	symbol *field.SymbolField,
	nomdentries *field.NoMDEntriesField) MarketDataSnapshotFullRefreshBuilder {
	var builder MarketDataSnapshotFullRefreshBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("W"))
	builder.Body.Set(symbol)
	builder.Body.Set(nomdentries)
	return builder
}

//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MDReqID() (*field.MDReqIDField, errors.MessageRejectError) {
	f := &field.MDReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMDReqID reads a MDReqID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMDReqID(f *field.MDReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Symbol() (*field.SymbolField, errors.MessageRejectError) {
	f := &field.SymbolField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbol(f *field.SymbolField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SymbolSfx() (*field.SymbolSfxField, errors.MessageRejectError) {
	f := &field.SymbolSfxField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSymbolSfx(f *field.SymbolSfxField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityID() (*field.SecurityIDField, errors.MessageRejectError) {
	f := &field.SecurityIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityID(f *field.SecurityIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IDSource is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) IDSource() (*field.IDSourceField, errors.MessageRejectError) {
	f := &field.IDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIDSource reads a IDSource from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIDSource(f *field.IDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityMonthYear() (*field.MaturityMonthYearField, errors.MessageRejectError) {
	f := &field.MaturityMonthYearField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityMonthYear(f *field.MaturityMonthYearField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDay is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) MaturityDay() (*field.MaturityDayField, errors.MessageRejectError) {
	f := &field.MaturityDayField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDay reads a MaturityDay from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetMaturityDay(f *field.MaturityDayField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PutOrCall is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) PutOrCall() (*field.PutOrCallField, errors.MessageRejectError) {
	f := &field.PutOrCallField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPutOrCall reads a PutOrCall from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetPutOrCall(f *field.PutOrCallField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) StrikePrice() (*field.StrikePriceField, errors.MessageRejectError) {
	f := &field.StrikePriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetStrikePrice(f *field.StrikePriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) OptAttribute() (*field.OptAttributeField, errors.MessageRejectError) {
	f := &field.OptAttributeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetOptAttribute(f *field.OptAttributeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) ContractMultiplier() (*field.ContractMultiplierField, errors.MessageRejectError) {
	f := &field.ContractMultiplierField{}
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetContractMultiplier(f *field.ContractMultiplierField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CouponRate() (*field.CouponRateField, errors.MessageRejectError) {
	f := &field.CouponRateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCouponRate(f *field.CouponRateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityExchange() (*field.SecurityExchangeField, errors.MessageRejectError) {
	f := &field.SecurityExchangeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityExchange(f *field.SecurityExchangeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) Issuer() (*field.IssuerField, errors.MessageRejectError) {
	f := &field.IssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetIssuer(f *field.IssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuerLen() (*field.EncodedIssuerLenField, errors.MessageRejectError) {
	f := &field.EncodedIssuerLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuerLen(f *field.EncodedIssuerLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedIssuer() (*field.EncodedIssuerField, errors.MessageRejectError) {
	f := &field.EncodedIssuerField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedIssuer(f *field.EncodedIssuerField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) SecurityDesc() (*field.SecurityDescField, errors.MessageRejectError) {
	f := &field.SecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetSecurityDesc(f *field.SecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDescLen() (*field.EncodedSecurityDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) EncodedSecurityDesc() (*field.EncodedSecurityDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetEncodedSecurityDesc(f *field.EncodedSecurityDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) FinancialStatus() (*field.FinancialStatusField, errors.MessageRejectError) {
	f := &field.FinancialStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetFinancialStatus reads a FinancialStatus from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetFinancialStatus(f *field.FinancialStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) CorporateAction() (*field.CorporateActionField, errors.MessageRejectError) {
	f := &field.CorporateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetCorporateAction(f *field.CorporateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalVolumeTraded is a non-required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) TotalVolumeTraded() (*field.TotalVolumeTradedField, errors.MessageRejectError) {
	f := &field.TotalVolumeTradedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalVolumeTraded reads a TotalVolumeTraded from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetTotalVolumeTraded(f *field.TotalVolumeTradedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) NoMDEntries() (*field.NoMDEntriesField, errors.MessageRejectError) {
	f := &field.NoMDEntriesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoMDEntries reads a NoMDEntries from MarketDataSnapshotFullRefresh.
func (m MarketDataSnapshotFullRefresh) GetNoMDEntries(f *field.NoMDEntriesField) errors.MessageRejectError {
	return m.Body.Get(f)
}
