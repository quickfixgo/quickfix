package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReportRequest msg type = AD.
type TradeCaptureReportRequest struct {
	message.Message
}

//TradeCaptureReportRequestBuilder builds TradeCaptureReportRequest messages.
type TradeCaptureReportRequestBuilder struct {
	message.MessageBuilder
}

//NewTradeCaptureReportRequestBuilder returns an initialized TradeCaptureReportRequestBuilder with specified required fields.
func NewTradeCaptureReportRequestBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType) *TradeCaptureReportRequestBuilder {
	builder := new(TradeCaptureReportRequestBuilder)
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestType is a required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeRequestType() (*field.TradeRequestType, error) {
	f := new(field.TradeRequestType)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecondaryTradeReportID() (*field.SecondaryTradeReportID, error) {
	f := new(field.SecondaryTradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//OrderID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) OrderID() (*field.OrderID, error) {
	f := new(field.OrderID)
	err := m.Body.Get(f)
	return f, err
}

//ClOrdID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TrdType() (*field.TrdType, error) {
	f := new(field.TrdType)
	err := m.Body.Get(f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TrdSubType() (*field.TrdSubType, error) {
	f := new(field.TrdSubType)
	err := m.Body.Get(f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TransferReason() (*field.TransferReason, error) {
	f := new(field.TransferReason)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecondaryTrdType() (*field.SecondaryTrdType, error) {
	f := new(field.SecondaryTrdType)
	err := m.Body.Get(f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeLinkID() (*field.TradeLinkID, error) {
	f := new(field.TradeLinkID)
	err := m.Body.Get(f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TrdMatchID() (*field.TrdMatchID, error) {
	f := new(field.TrdMatchID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecuritySubType() (*field.SecuritySubType, error) {
	f := new(field.SecuritySubType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) StrikeCurrency() (*field.StrikeCurrency, error) {
	f := new(field.StrikeCurrency)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Pool() (*field.Pool, error) {
	f := new(field.Pool)
	err := m.Body.Get(f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ContractSettlMonth() (*field.ContractSettlMonth, error) {
	f := new(field.ContractSettlMonth)
	err := m.Body.Get(f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CPProgram() (*field.CPProgram, error) {
	f := new(field.CPProgram)
	err := m.Body.Get(f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) CPRegType() (*field.CPRegType, error) {
	f := new(field.CPRegType)
	err := m.Body.Get(f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoEvents() (*field.NoEvents, error) {
	f := new(field.NoEvents)
	err := m.Body.Get(f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) DatedDate() (*field.DatedDate, error) {
	f := new(field.DatedDate)
	err := m.Body.Get(f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) InterestAccrualDate() (*field.InterestAccrualDate, error) {
	f := new(field.InterestAccrualDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryForm is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) DeliveryForm() (*field.DeliveryForm, error) {
	f := new(field.DeliveryForm)
	err := m.Body.Get(f)
	return f, err
}

//PctAtRisk is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) PctAtRisk() (*field.PctAtRisk, error) {
	f := new(field.PctAtRisk)
	err := m.Body.Get(f)
	return f, err
}

//NoInstrAttrib is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoInstrAttrib() (*field.NoInstrAttrib, error) {
	f := new(field.NoInstrAttrib)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) AgreementDesc() (*field.AgreementDesc, error) {
	f := new(field.AgreementDesc)
	err := m.Body.Get(f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) AgreementID() (*field.AgreementID, error) {
	f := new(field.AgreementID)
	err := m.Body.Get(f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) AgreementDate() (*field.AgreementDate, error) {
	f := new(field.AgreementDate)
	err := m.Body.Get(f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) AgreementCurrency() (*field.AgreementCurrency, error) {
	f := new(field.AgreementCurrency)
	err := m.Body.Get(f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TerminationType() (*field.TerminationType, error) {
	f := new(field.TerminationType)
	err := m.Body.Get(f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) StartDate() (*field.StartDate, error) {
	f := new(field.StartDate)
	err := m.Body.Get(f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EndDate() (*field.EndDate, error) {
	f := new(field.EndDate)
	err := m.Body.Get(f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) DeliveryType() (*field.DeliveryType, error) {
	f := new(field.DeliveryType)
	err := m.Body.Get(f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) MarginRatio() (*field.MarginRatio, error) {
	f := new(field.MarginRatio)
	err := m.Body.Get(f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//NoDates is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) NoDates() (*field.NoDates, error) {
	f := new(field.NoDates)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//TimeBracket is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TimeBracket() (*field.TimeBracket, error) {
	f := new(field.TimeBracket)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) MultiLegReportingType() (*field.MultiLegReportingType, error) {
	f := new(field.MultiLegReportingType)
	err := m.Body.Get(f)
	return f, err
}

//TradeInputSource is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeInputSource() (*field.TradeInputSource, error) {
	f := new(field.TradeInputSource)
	err := m.Body.Get(f)
	return f, err
}

//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) TradeInputDevice() (*field.TradeInputDevice, error) {
	f := new(field.TradeInputDevice)
	err := m.Body.Get(f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ResponseTransportType() (*field.ResponseTransportType, error) {
	f := new(field.ResponseTransportType)
	err := m.Body.Get(f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) ResponseDestination() (*field.ResponseDestination, error) {
	f := new(field.ResponseDestination)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportRequest.
func (m *TradeCaptureReportRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
