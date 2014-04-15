package fix50

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

//CreateTradeCaptureReportRequestBuilder returns an initialized TradeCaptureReportRequestBuilder with specified required fields.
func CreateTradeCaptureReportRequestBuilder(
	traderequestid field.TradeRequestID,
	traderequesttype field.TradeRequestType) TradeCaptureReportRequestBuilder {
	var builder TradeCaptureReportRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(traderequestid)
	builder.Body.Set(traderequesttype)
	return builder
}

//TradeRequestID is a required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeRequestID() (field.TradeRequestID, error) {
	var f field.TradeRequestID
	err := m.Body.Get(&f)
	return f, err
}

//TradeRequestType is a required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeRequestType() (field.TradeRequestType, error) {
	var f field.TradeRequestType
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SubscriptionRequestType() (field.SubscriptionRequestType, error) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}

//TradeReportID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeReportID() (field.TradeReportID, error) {
	var f field.TradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeReportID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecondaryTradeReportID() (field.SecondaryTradeReportID, error) {
	var f field.SecondaryTradeReportID
	err := m.Body.Get(&f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ExecID() (field.ExecID, error) {
	var f field.ExecID
	err := m.Body.Get(&f)
	return f, err
}

//ExecType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ExecType() (field.ExecType, error) {
	var f field.ExecType
	err := m.Body.Get(&f)
	return f, err
}

//OrderID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) OrderID() (field.OrderID, error) {
	var f field.OrderID
	err := m.Body.Get(&f)
	return f, err
}

//ClOrdID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ClOrdID() (field.ClOrdID, error) {
	var f field.ClOrdID
	err := m.Body.Get(&f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MatchStatus() (field.MatchStatus, error) {
	var f field.MatchStatus
	err := m.Body.Get(&f)
	return f, err
}

//TrdType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TrdType() (field.TrdType, error) {
	var f field.TrdType
	err := m.Body.Get(&f)
	return f, err
}

//TrdSubType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TrdSubType() (field.TrdSubType, error) {
	var f field.TrdSubType
	err := m.Body.Get(&f)
	return f, err
}

//TransferReason is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TransferReason() (field.TransferReason, error) {
	var f field.TransferReason
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTrdType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecondaryTrdType() (field.SecondaryTrdType, error) {
	var f field.SecondaryTrdType
	err := m.Body.Get(&f)
	return f, err
}

//TradeLinkID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeLinkID() (field.TradeLinkID, error) {
	var f field.TradeLinkID
	err := m.Body.Get(&f)
	return f, err
}

//TrdMatchID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TrdMatchID() (field.TrdMatchID, error) {
	var f field.TrdMatchID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityStatus is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecurityStatus() (field.SecurityStatus, error) {
	var f field.SecurityStatus
	err := m.Body.Get(&f)
	return f, err
}

//SettleOnOpenFlag is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SettleOnOpenFlag() (field.SettleOnOpenFlag, error) {
	var f field.SettleOnOpenFlag
	err := m.Body.Get(&f)
	return f, err
}

//InstrmtAssignmentMethod is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) InstrmtAssignmentMethod() (field.InstrmtAssignmentMethod, error) {
	var f field.InstrmtAssignmentMethod
	err := m.Body.Get(&f)
	return f, err
}

//StrikeMultiplier is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StrikeMultiplier() (field.StrikeMultiplier, error) {
	var f field.StrikeMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//StrikeValue is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StrikeValue() (field.StrikeValue, error) {
	var f field.StrikeValue
	err := m.Body.Get(&f)
	return f, err
}

//MinPriceIncrement is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MinPriceIncrement() (field.MinPriceIncrement, error) {
	var f field.MinPriceIncrement
	err := m.Body.Get(&f)
	return f, err
}

//PositionLimit is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) PositionLimit() (field.PositionLimit, error) {
	var f field.PositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NTPositionLimit is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NTPositionLimit() (field.NTPositionLimit, error) {
	var f field.NTPositionLimit
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrumentParties is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoInstrumentParties() (field.NoInstrumentParties, error) {
	var f field.NoInstrumentParties
	err := m.Body.Get(&f)
	return f, err
}

//UnitOfMeasure is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) UnitOfMeasure() (field.UnitOfMeasure, error) {
	var f field.UnitOfMeasure
	err := m.Body.Get(&f)
	return f, err
}

//TimeUnit is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TimeUnit() (field.TimeUnit, error) {
	var f field.TimeUnit
	err := m.Body.Get(&f)
	return f, err
}

//MaturityTime is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MaturityTime() (field.MaturityTime, error) {
	var f field.MaturityTime
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryForm is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) DeliveryForm() (field.DeliveryForm, error) {
	var f field.DeliveryForm
	err := m.Body.Get(&f)
	return f, err
}

//PctAtRisk is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) PctAtRisk() (field.PctAtRisk, error) {
	var f field.PctAtRisk
	err := m.Body.Get(&f)
	return f, err
}

//NoInstrAttrib is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoInstrAttrib() (field.NoInstrAttrib, error) {
	var f field.NoInstrAttrib
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDesc is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) AgreementDesc() (field.AgreementDesc, error) {
	var f field.AgreementDesc
	err := m.Body.Get(&f)
	return f, err
}

//AgreementID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) AgreementID() (field.AgreementID, error) {
	var f field.AgreementID
	err := m.Body.Get(&f)
	return f, err
}

//AgreementDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) AgreementDate() (field.AgreementDate, error) {
	var f field.AgreementDate
	err := m.Body.Get(&f)
	return f, err
}

//AgreementCurrency is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) AgreementCurrency() (field.AgreementCurrency, error) {
	var f field.AgreementCurrency
	err := m.Body.Get(&f)
	return f, err
}

//TerminationType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TerminationType() (field.TerminationType, error) {
	var f field.TerminationType
	err := m.Body.Get(&f)
	return f, err
}

//StartDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) StartDate() (field.StartDate, error) {
	var f field.StartDate
	err := m.Body.Get(&f)
	return f, err
}

//EndDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EndDate() (field.EndDate, error) {
	var f field.EndDate
	err := m.Body.Get(&f)
	return f, err
}

//DeliveryType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) DeliveryType() (field.DeliveryType, error) {
	var f field.DeliveryType
	err := m.Body.Get(&f)
	return f, err
}

//MarginRatio is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MarginRatio() (field.MarginRatio, error) {
	var f field.MarginRatio
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoDates is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) NoDates() (field.NoDates, error) {
	var f field.NoDates
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionSubID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradingSessionSubID() (field.TradingSessionSubID, error) {
	var f field.TradingSessionSubID
	err := m.Body.Get(&f)
	return f, err
}

//TimeBracket is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TimeBracket() (field.TimeBracket, error) {
	var f field.TimeBracket
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//MultiLegReportingType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MultiLegReportingType() (field.MultiLegReportingType, error) {
	var f field.MultiLegReportingType
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputSource is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeInputSource() (field.TradeInputSource, error) {
	var f field.TradeInputSource
	err := m.Body.Get(&f)
	return f, err
}

//TradeInputDevice is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeInputDevice() (field.TradeInputDevice, error) {
	var f field.TradeInputDevice
	err := m.Body.Get(&f)
	return f, err
}

//ResponseTransportType is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ResponseTransportType() (field.ResponseTransportType, error) {
	var f field.ResponseTransportType
	err := m.Body.Get(&f)
	return f, err
}

//ResponseDestination is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) ResponseDestination() (field.ResponseDestination, error) {
	var f field.ResponseDestination
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}

//MessageEventSource is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) MessageEventSource() (field.MessageEventSource, error) {
	var f field.MessageEventSource
	err := m.Body.Get(&f)
	return f, err
}

//TradeID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeID() (field.TradeID, error) {
	var f field.TradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryTradeID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecondaryTradeID() (field.SecondaryTradeID, error) {
	var f field.SecondaryTradeID
	err := m.Body.Get(&f)
	return f, err
}

//FirmTradeID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) FirmTradeID() (field.FirmTradeID, error) {
	var f field.FirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//SecondaryFirmTradeID is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) SecondaryFirmTradeID() (field.SecondaryFirmTradeID, error) {
	var f field.SecondaryFirmTradeID
	err := m.Body.Get(&f)
	return f, err
}

//TradeHandlingInstr is a non-required field for TradeCaptureReportRequest.
func (m TradeCaptureReportRequest) TradeHandlingInstr() (field.TradeHandlingInstr, error) {
	var f field.TradeHandlingInstr
	err := m.Body.Get(&f)
	return f, err
}
