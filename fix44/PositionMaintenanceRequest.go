package fix44

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//PositionMaintenanceRequest msg type = AL.
type PositionMaintenanceRequest struct {
	message.Message
}

//PositionMaintenanceRequestBuilder builds PositionMaintenanceRequest messages.
type PositionMaintenanceRequestBuilder struct {
	message.MessageBuilder
}

//CreatePositionMaintenanceRequestBuilder returns an initialized PositionMaintenanceRequestBuilder with specified required fields.
func CreatePositionMaintenanceRequestBuilder(
	posreqid field.PosReqID,
	postranstype field.PosTransType,
	posmaintaction field.PosMaintAction,
	clearingbusinessdate field.ClearingBusinessDate,
	account field.Account,
	accounttype field.AccountType,
	transacttime field.TransactTime) PositionMaintenanceRequestBuilder {
	var builder PositionMaintenanceRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posreqid)
	builder.Body.Set(postranstype)
	builder.Body.Set(posmaintaction)
	builder.Body.Set(clearingbusinessdate)
	builder.Body.Set(account)
	builder.Body.Set(accounttype)
	builder.Body.Set(transacttime)
	return builder
}

//PosReqID is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosReqID() (field.PosReqID, error) {
	var f field.PosReqID
	err := m.Body.Get(&f)
	return f, err
}

//PosTransType is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosTransType() (field.PosTransType, error) {
	var f field.PosTransType
	err := m.Body.Get(&f)
	return f, err
}

//PosMaintAction is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosMaintAction() (field.PosMaintAction, error) {
	var f field.PosMaintAction
	err := m.Body.Get(&f)
	return f, err
}

//OrigPosReqRefID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) OrigPosReqRefID() (field.OrigPosReqRefID, error) {
	var f field.OrigPosReqRefID
	err := m.Body.Get(&f)
	return f, err
}

//PosMaintRptRefID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PosMaintRptRefID() (field.PosMaintRptRefID, error) {
	var f field.PosMaintRptRefID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessSubID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SettlSessSubID() (field.SettlSessSubID, error) {
	var f field.SettlSessSubID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//Account is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Account() (field.Account, error) {
	var f field.Account
	err := m.Body.Get(&f)
	return f, err
}

//AcctIDSource is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AcctIDSource() (field.AcctIDSource, error) {
	var f field.AcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//AccountType is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AccountType() (field.AccountType, error) {
	var f field.AccountType
	err := m.Body.Get(&f)
	return f, err
}

//Symbol is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Symbol() (field.Symbol, error) {
	var f field.Symbol
	err := m.Body.Get(&f)
	return f, err
}

//SymbolSfx is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SymbolSfx() (field.SymbolSfx, error) {
	var f field.SymbolSfx
	err := m.Body.Get(&f)
	return f, err
}

//SecurityID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityID() (field.SecurityID, error) {
	var f field.SecurityID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityIDSource is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityIDSource() (field.SecurityIDSource, error) {
	var f field.SecurityIDSource
	err := m.Body.Get(&f)
	return f, err
}

//NoSecurityAltID is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoSecurityAltID() (field.NoSecurityAltID, error) {
	var f field.NoSecurityAltID
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//SecuritySubType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecuritySubType() (field.SecuritySubType, error) {
	var f field.SecuritySubType
	err := m.Body.Get(&f)
	return f, err
}

//MaturityMonthYear is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MaturityMonthYear() (field.MaturityMonthYear, error) {
	var f field.MaturityMonthYear
	err := m.Body.Get(&f)
	return f, err
}

//MaturityDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) MaturityDate() (field.MaturityDate, error) {
	var f field.MaturityDate
	err := m.Body.Get(&f)
	return f, err
}

//CouponPaymentDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CouponPaymentDate() (field.CouponPaymentDate, error) {
	var f field.CouponPaymentDate
	err := m.Body.Get(&f)
	return f, err
}

//IssueDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) IssueDate() (field.IssueDate, error) {
	var f field.IssueDate
	err := m.Body.Get(&f)
	return f, err
}

//RepoCollateralSecurityType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepoCollateralSecurityType() (field.RepoCollateralSecurityType, error) {
	var f field.RepoCollateralSecurityType
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseTerm is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepurchaseTerm() (field.RepurchaseTerm, error) {
	var f field.RepurchaseTerm
	err := m.Body.Get(&f)
	return f, err
}

//RepurchaseRate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RepurchaseRate() (field.RepurchaseRate, error) {
	var f field.RepurchaseRate
	err := m.Body.Get(&f)
	return f, err
}

//Factor is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Factor() (field.Factor, error) {
	var f field.Factor
	err := m.Body.Get(&f)
	return f, err
}

//CreditRating is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CreditRating() (field.CreditRating, error) {
	var f field.CreditRating
	err := m.Body.Get(&f)
	return f, err
}

//InstrRegistry is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) InstrRegistry() (field.InstrRegistry, error) {
	var f field.InstrRegistry
	err := m.Body.Get(&f)
	return f, err
}

//CountryOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CountryOfIssue() (field.CountryOfIssue, error) {
	var f field.CountryOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//StateOrProvinceOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StateOrProvinceOfIssue() (field.StateOrProvinceOfIssue, error) {
	var f field.StateOrProvinceOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//LocaleOfIssue is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) LocaleOfIssue() (field.LocaleOfIssue, error) {
	var f field.LocaleOfIssue
	err := m.Body.Get(&f)
	return f, err
}

//RedemptionDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) RedemptionDate() (field.RedemptionDate, error) {
	var f field.RedemptionDate
	err := m.Body.Get(&f)
	return f, err
}

//StrikePrice is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikePrice() (field.StrikePrice, error) {
	var f field.StrikePrice
	err := m.Body.Get(&f)
	return f, err
}

//StrikeCurrency is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) StrikeCurrency() (field.StrikeCurrency, error) {
	var f field.StrikeCurrency
	err := m.Body.Get(&f)
	return f, err
}

//OptAttribute is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) OptAttribute() (field.OptAttribute, error) {
	var f field.OptAttribute
	err := m.Body.Get(&f)
	return f, err
}

//ContractMultiplier is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContractMultiplier() (field.ContractMultiplier, error) {
	var f field.ContractMultiplier
	err := m.Body.Get(&f)
	return f, err
}

//CouponRate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CouponRate() (field.CouponRate, error) {
	var f field.CouponRate
	err := m.Body.Get(&f)
	return f, err
}

//SecurityExchange is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityExchange() (field.SecurityExchange, error) {
	var f field.SecurityExchange
	err := m.Body.Get(&f)
	return f, err
}

//Issuer is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Issuer() (field.Issuer, error) {
	var f field.Issuer
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuerLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedIssuerLen() (field.EncodedIssuerLen, error) {
	var f field.EncodedIssuerLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedIssuer is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedIssuer() (field.EncodedIssuer, error) {
	var f field.EncodedIssuer
	err := m.Body.Get(&f)
	return f, err
}

//SecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) SecurityDesc() (field.SecurityDesc, error) {
	var f field.SecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDescLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedSecurityDescLen() (field.EncodedSecurityDescLen, error) {
	var f field.EncodedSecurityDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityDesc is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedSecurityDesc() (field.EncodedSecurityDesc, error) {
	var f field.EncodedSecurityDesc
	err := m.Body.Get(&f)
	return f, err
}

//Pool is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Pool() (field.Pool, error) {
	var f field.Pool
	err := m.Body.Get(&f)
	return f, err
}

//ContractSettlMonth is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContractSettlMonth() (field.ContractSettlMonth, error) {
	var f field.ContractSettlMonth
	err := m.Body.Get(&f)
	return f, err
}

//CPProgram is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CPProgram() (field.CPProgram, error) {
	var f field.CPProgram
	err := m.Body.Get(&f)
	return f, err
}

//CPRegType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) CPRegType() (field.CPRegType, error) {
	var f field.CPRegType
	err := m.Body.Get(&f)
	return f, err
}

//NoEvents is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoEvents() (field.NoEvents, error) {
	var f field.NoEvents
	err := m.Body.Get(&f)
	return f, err
}

//DatedDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) DatedDate() (field.DatedDate, error) {
	var f field.DatedDate
	err := m.Body.Get(&f)
	return f, err
}

//InterestAccrualDate is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) InterestAccrualDate() (field.InterestAccrualDate, error) {
	var f field.InterestAccrualDate
	err := m.Body.Get(&f)
	return f, err
}

//Currency is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Currency() (field.Currency, error) {
	var f field.Currency
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoLegs() (field.NoLegs, error) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoUnderlyings() (field.NoUnderlyings, error) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoTradingSessions is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoTradingSessions() (field.NoTradingSessions, error) {
	var f field.NoTradingSessions
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPositions is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) NoPositions() (field.NoPositions, error) {
	var f field.NoPositions
	err := m.Body.Get(&f)
	return f, err
}

//AdjustmentType is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) AdjustmentType() (field.AdjustmentType, error) {
	var f field.AdjustmentType
	err := m.Body.Get(&f)
	return f, err
}

//ContraryInstructionIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ContraryInstructionIndicator() (field.ContraryInstructionIndicator, error) {
	var f field.ContraryInstructionIndicator
	err := m.Body.Get(&f)
	return f, err
}

//PriorSpreadIndicator is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) PriorSpreadIndicator() (field.PriorSpreadIndicator, error) {
	var f field.PriorSpreadIndicator
	err := m.Body.Get(&f)
	return f, err
}

//ThresholdAmount is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) ThresholdAmount() (field.ThresholdAmount, error) {
	var f field.ThresholdAmount
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for PositionMaintenanceRequest.
func (m PositionMaintenanceRequest) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
