package fix43

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradeCaptureReport msg type = AE.
type TradeCaptureReport struct {
	message.Message
}

//TradeCaptureReportBuilder builds TradeCaptureReport messages.
type TradeCaptureReportBuilder struct {
	message.MessageBuilder
}

//NewTradeCaptureReportBuilder returns an initialized TradeCaptureReportBuilder with specified required fields.
func NewTradeCaptureReportBuilder(
	tradereportid field.TradeReportID,
	exectype field.ExecType,
	previouslyreported field.PreviouslyReported,
	lastqty field.LastQty,
	lastpx field.LastPx,
	tradedate field.TradeDate,
	transacttime field.TransactTime,
	nosides field.NoSides) *TradeCaptureReportBuilder {
	builder := new(TradeCaptureReportBuilder)
	builder.Body.Set(tradereportid)
	builder.Body.Set(exectype)
	builder.Body.Set(previouslyreported)
	builder.Body.Set(lastqty)
	builder.Body.Set(lastpx)
	builder.Body.Set(tradedate)
	builder.Body.Set(transacttime)
	builder.Body.Set(nosides)
	return builder
}

//TradeReportID is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportID() (*field.TradeReportID, error) {
	f := new(field.TradeReportID)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportTransType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportTransType() (*field.TradeReportTransType, error) {
	f := new(field.TradeReportTransType)
	err := m.Body.Get(f)
	return f, err
}

//TradeRequestID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeRequestID() (*field.TradeRequestID, error) {
	f := new(field.TradeRequestID)
	err := m.Body.Get(f)
	return f, err
}

//ExecType is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecType() (*field.ExecType, error) {
	f := new(field.ExecType)
	err := m.Body.Get(f)
	return f, err
}

//TradeReportRefID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeReportRefID() (*field.TradeReportRefID, error) {
	f := new(field.TradeReportRefID)
	err := m.Body.Get(f)
	return f, err
}

//ExecID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecID() (*field.ExecID, error) {
	f := new(field.ExecID)
	err := m.Body.Get(f)
	return f, err
}

//SecondaryExecID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecondaryExecID() (*field.SecondaryExecID, error) {
	f := new(field.SecondaryExecID)
	err := m.Body.Get(f)
	return f, err
}

//ExecRestatementReason is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ExecRestatementReason() (*field.ExecRestatementReason, error) {
	f := new(field.ExecRestatementReason)
	err := m.Body.Get(f)
	return f, err
}

//PreviouslyReported is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) PreviouslyReported() (*field.PreviouslyReported, error) {
	f := new(field.PreviouslyReported)
	err := m.Body.Get(f)
	return f, err
}

//Symbol is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Symbol() (*field.Symbol, error) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//SymbolSfx is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SymbolSfx() (*field.SymbolSfx, error) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//SecurityID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityID() (*field.SecurityID, error) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityIDSource is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityIDSource() (*field.SecurityIDSource, error) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//NoSecurityAltID is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoSecurityAltID() (*field.NoSecurityAltID, error) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//MaturityMonthYear is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MaturityMonthYear() (*field.MaturityMonthYear, error) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//MaturityDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MaturityDate() (*field.MaturityDate, error) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//CouponPaymentDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CouponPaymentDate() (*field.CouponPaymentDate, error) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//IssueDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) IssueDate() (*field.IssueDate, error) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, error) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseTerm is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepurchaseTerm() (*field.RepurchaseTerm, error) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//RepurchaseRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RepurchaseRate() (*field.RepurchaseRate, error) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//Factor is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Factor() (*field.Factor, error) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//CreditRating is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CreditRating() (*field.CreditRating, error) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//InstrRegistry is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) InstrRegistry() (*field.InstrRegistry, error) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//CountryOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CountryOfIssue() (*field.CountryOfIssue, error) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, error) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//LocaleOfIssue is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LocaleOfIssue() (*field.LocaleOfIssue, error) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//RedemptionDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RedemptionDate() (*field.RedemptionDate, error) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//StrikePrice is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) StrikePrice() (*field.StrikePrice, error) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//OptAttribute is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OptAttribute() (*field.OptAttribute, error) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//ContractMultiplier is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) ContractMultiplier() (*field.ContractMultiplier, error) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//CouponRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CouponRate() (*field.CouponRate, error) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//SecurityExchange is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityExchange() (*field.SecurityExchange, error) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//Issuer is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) Issuer() (*field.Issuer, error) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuerLen is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedIssuerLen() (*field.EncodedIssuerLen, error) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedIssuer is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedIssuer() (*field.EncodedIssuer, error) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//SecurityDesc is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SecurityDesc() (*field.SecurityDesc, error) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, error) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) EncodedSecurityDesc() (*field.EncodedSecurityDesc, error) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//OrderQty is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrderQty() (*field.OrderQty, error) {
	f := new(field.OrderQty)
	err := m.Body.Get(f)
	return f, err
}

//CashOrderQty is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) CashOrderQty() (*field.CashOrderQty, error) {
	f := new(field.CashOrderQty)
	err := m.Body.Get(f)
	return f, err
}

//OrderPercent is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) OrderPercent() (*field.OrderPercent, error) {
	f := new(field.OrderPercent)
	err := m.Body.Get(f)
	return f, err
}

//RoundingDirection is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RoundingDirection() (*field.RoundingDirection, error) {
	f := new(field.RoundingDirection)
	err := m.Body.Get(f)
	return f, err
}

//RoundingModulus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) RoundingModulus() (*field.RoundingModulus, error) {
	f := new(field.RoundingModulus)
	err := m.Body.Get(f)
	return f, err
}

//LastQty is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastQty() (*field.LastQty, error) {
	f := new(field.LastQty)
	err := m.Body.Get(f)
	return f, err
}

//LastPx is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastPx() (*field.LastPx, error) {
	f := new(field.LastPx)
	err := m.Body.Get(f)
	return f, err
}

//LastSpotRate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastSpotRate() (*field.LastSpotRate, error) {
	f := new(field.LastSpotRate)
	err := m.Body.Get(f)
	return f, err
}

//LastForwardPoints is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastForwardPoints() (*field.LastForwardPoints, error) {
	f := new(field.LastForwardPoints)
	err := m.Body.Get(f)
	return f, err
}

//LastMkt is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) LastMkt() (*field.LastMkt, error) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//TradeDate is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) TradeDate() (*field.TradeDate, error) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//SettlmntTyp is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) SettlmntTyp() (*field.SettlmntTyp, error) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//FutSettDate is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) FutSettDate() (*field.FutSettDate, error) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//MatchStatus is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MatchStatus() (*field.MatchStatus, error) {
	f := new(field.MatchStatus)
	err := m.Body.Get(f)
	return f, err
}

//MatchType is a non-required field for TradeCaptureReport.
func (m *TradeCaptureReport) MatchType() (*field.MatchType, error) {
	f := new(field.MatchType)
	err := m.Body.Get(f)
	return f, err
}

//NoSides is a required field for TradeCaptureReport.
func (m *TradeCaptureReport) NoSides() (*field.NoSides, error) {
	f := new(field.NoSides)
	err := m.Body.Get(f)
	return f, err
}
