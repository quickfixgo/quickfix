package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Allocation msg type = J.
type Allocation struct {
	message.Message
}

//AllocationBuilder builds Allocation messages.
type AllocationBuilder struct {
	message.MessageBuilder
}

//CreateAllocationBuilder returns an initialized AllocationBuilder with specified required fields.
func CreateAllocationBuilder(
	allocid field.AllocID,
	alloctranstype field.AllocTransType,
	alloctype field.AllocType,
	side field.Side,
	quantity field.Quantity,
	avgpx field.AvgPx,
	tradedate field.TradeDate) AllocationBuilder {
	var builder AllocationBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("J"))
	builder.Body.Set(allocid)
	builder.Body.Set(alloctranstype)
	builder.Body.Set(alloctype)
	builder.Body.Set(side)
	builder.Body.Set(quantity)
	builder.Body.Set(avgpx)
	builder.Body.Set(tradedate)
	return builder
}

//AllocID is a required field for Allocation.
func (m Allocation) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from Allocation.
func (m Allocation) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocTransType is a required field for Allocation.
func (m Allocation) AllocTransType() (*field.AllocTransType, errors.MessageRejectError) {
	f := new(field.AllocTransType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocTransType reads a AllocTransType from Allocation.
func (m Allocation) GetAllocTransType(f *field.AllocTransType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocType is a required field for Allocation.
func (m Allocation) AllocType() (*field.AllocType, errors.MessageRejectError) {
	f := new(field.AllocType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocType reads a AllocType from Allocation.
func (m Allocation) GetAllocType(f *field.AllocType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RefAllocID is a non-required field for Allocation.
func (m Allocation) RefAllocID() (*field.RefAllocID, errors.MessageRejectError) {
	f := new(field.RefAllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetRefAllocID reads a RefAllocID from Allocation.
func (m Allocation) GetRefAllocID(f *field.RefAllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkID is a non-required field for Allocation.
func (m Allocation) AllocLinkID() (*field.AllocLinkID, errors.MessageRejectError) {
	f := new(field.AllocLinkID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkID reads a AllocLinkID from Allocation.
func (m Allocation) GetAllocLinkID(f *field.AllocLinkID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocLinkType is a non-required field for Allocation.
func (m Allocation) AllocLinkType() (*field.AllocLinkType, errors.MessageRejectError) {
	f := new(field.AllocLinkType)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocLinkType reads a AllocLinkType from Allocation.
func (m Allocation) GetAllocLinkType(f *field.AllocLinkType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//BookingRefID is a non-required field for Allocation.
func (m Allocation) BookingRefID() (*field.BookingRefID, errors.MessageRejectError) {
	f := new(field.BookingRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetBookingRefID reads a BookingRefID from Allocation.
func (m Allocation) GetBookingRefID(f *field.BookingRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a non-required field for Allocation.
func (m Allocation) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from Allocation.
func (m Allocation) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoExecs is a non-required field for Allocation.
func (m Allocation) NoExecs() (*field.NoExecs, errors.MessageRejectError) {
	f := new(field.NoExecs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoExecs reads a NoExecs from Allocation.
func (m Allocation) GetNoExecs(f *field.NoExecs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a required field for Allocation.
func (m Allocation) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from Allocation.
func (m Allocation) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Symbol is a non-required field for Allocation.
func (m Allocation) Symbol() (*field.Symbol, errors.MessageRejectError) {
	f := new(field.Symbol)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbol reads a Symbol from Allocation.
func (m Allocation) GetSymbol(f *field.Symbol) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SymbolSfx is a non-required field for Allocation.
func (m Allocation) SymbolSfx() (*field.SymbolSfx, errors.MessageRejectError) {
	f := new(field.SymbolSfx)
	err := m.Body.Get(f)
	return f, err
}

//GetSymbolSfx reads a SymbolSfx from Allocation.
func (m Allocation) GetSymbolSfx(f *field.SymbolSfx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityID is a non-required field for Allocation.
func (m Allocation) SecurityID() (*field.SecurityID, errors.MessageRejectError) {
	f := new(field.SecurityID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityID reads a SecurityID from Allocation.
func (m Allocation) GetSecurityID(f *field.SecurityID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityIDSource is a non-required field for Allocation.
func (m Allocation) SecurityIDSource() (*field.SecurityIDSource, errors.MessageRejectError) {
	f := new(field.SecurityIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityIDSource reads a SecurityIDSource from Allocation.
func (m Allocation) GetSecurityIDSource(f *field.SecurityIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSecurityAltID is a non-required field for Allocation.
func (m Allocation) NoSecurityAltID() (*field.NoSecurityAltID, errors.MessageRejectError) {
	f := new(field.NoSecurityAltID)
	err := m.Body.Get(f)
	return f, err
}

//GetNoSecurityAltID reads a NoSecurityAltID from Allocation.
func (m Allocation) GetNoSecurityAltID(f *field.NoSecurityAltID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for Allocation.
func (m Allocation) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from Allocation.
func (m Allocation) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for Allocation.
func (m Allocation) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from Allocation.
func (m Allocation) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for Allocation.
func (m Allocation) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from Allocation.
func (m Allocation) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityMonthYear is a non-required field for Allocation.
func (m Allocation) MaturityMonthYear() (*field.MaturityMonthYear, errors.MessageRejectError) {
	f := new(field.MaturityMonthYear)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityMonthYear reads a MaturityMonthYear from Allocation.
func (m Allocation) GetMaturityMonthYear(f *field.MaturityMonthYear) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MaturityDate is a non-required field for Allocation.
func (m Allocation) MaturityDate() (*field.MaturityDate, errors.MessageRejectError) {
	f := new(field.MaturityDate)
	err := m.Body.Get(f)
	return f, err
}

//GetMaturityDate reads a MaturityDate from Allocation.
func (m Allocation) GetMaturityDate(f *field.MaturityDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponPaymentDate is a non-required field for Allocation.
func (m Allocation) CouponPaymentDate() (*field.CouponPaymentDate, errors.MessageRejectError) {
	f := new(field.CouponPaymentDate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponPaymentDate reads a CouponPaymentDate from Allocation.
func (m Allocation) GetCouponPaymentDate(f *field.CouponPaymentDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//IssueDate is a non-required field for Allocation.
func (m Allocation) IssueDate() (*field.IssueDate, errors.MessageRejectError) {
	f := new(field.IssueDate)
	err := m.Body.Get(f)
	return f, err
}

//GetIssueDate reads a IssueDate from Allocation.
func (m Allocation) GetIssueDate(f *field.IssueDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepoCollateralSecurityType is a non-required field for Allocation.
func (m Allocation) RepoCollateralSecurityType() (*field.RepoCollateralSecurityType, errors.MessageRejectError) {
	f := new(field.RepoCollateralSecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetRepoCollateralSecurityType reads a RepoCollateralSecurityType from Allocation.
func (m Allocation) GetRepoCollateralSecurityType(f *field.RepoCollateralSecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseTerm is a non-required field for Allocation.
func (m Allocation) RepurchaseTerm() (*field.RepurchaseTerm, errors.MessageRejectError) {
	f := new(field.RepurchaseTerm)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseTerm reads a RepurchaseTerm from Allocation.
func (m Allocation) GetRepurchaseTerm(f *field.RepurchaseTerm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RepurchaseRate is a non-required field for Allocation.
func (m Allocation) RepurchaseRate() (*field.RepurchaseRate, errors.MessageRejectError) {
	f := new(field.RepurchaseRate)
	err := m.Body.Get(f)
	return f, err
}

//GetRepurchaseRate reads a RepurchaseRate from Allocation.
func (m Allocation) GetRepurchaseRate(f *field.RepurchaseRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Factor is a non-required field for Allocation.
func (m Allocation) Factor() (*field.Factor, errors.MessageRejectError) {
	f := new(field.Factor)
	err := m.Body.Get(f)
	return f, err
}

//GetFactor reads a Factor from Allocation.
func (m Allocation) GetFactor(f *field.Factor) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CreditRating is a non-required field for Allocation.
func (m Allocation) CreditRating() (*field.CreditRating, errors.MessageRejectError) {
	f := new(field.CreditRating)
	err := m.Body.Get(f)
	return f, err
}

//GetCreditRating reads a CreditRating from Allocation.
func (m Allocation) GetCreditRating(f *field.CreditRating) errors.MessageRejectError {
	return m.Body.Get(f)
}

//InstrRegistry is a non-required field for Allocation.
func (m Allocation) InstrRegistry() (*field.InstrRegistry, errors.MessageRejectError) {
	f := new(field.InstrRegistry)
	err := m.Body.Get(f)
	return f, err
}

//GetInstrRegistry reads a InstrRegistry from Allocation.
func (m Allocation) GetInstrRegistry(f *field.InstrRegistry) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CountryOfIssue is a non-required field for Allocation.
func (m Allocation) CountryOfIssue() (*field.CountryOfIssue, errors.MessageRejectError) {
	f := new(field.CountryOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetCountryOfIssue reads a CountryOfIssue from Allocation.
func (m Allocation) GetCountryOfIssue(f *field.CountryOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StateOrProvinceOfIssue is a non-required field for Allocation.
func (m Allocation) StateOrProvinceOfIssue() (*field.StateOrProvinceOfIssue, errors.MessageRejectError) {
	f := new(field.StateOrProvinceOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetStateOrProvinceOfIssue reads a StateOrProvinceOfIssue from Allocation.
func (m Allocation) GetStateOrProvinceOfIssue(f *field.StateOrProvinceOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LocaleOfIssue is a non-required field for Allocation.
func (m Allocation) LocaleOfIssue() (*field.LocaleOfIssue, errors.MessageRejectError) {
	f := new(field.LocaleOfIssue)
	err := m.Body.Get(f)
	return f, err
}

//GetLocaleOfIssue reads a LocaleOfIssue from Allocation.
func (m Allocation) GetLocaleOfIssue(f *field.LocaleOfIssue) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RedemptionDate is a non-required field for Allocation.
func (m Allocation) RedemptionDate() (*field.RedemptionDate, errors.MessageRejectError) {
	f := new(field.RedemptionDate)
	err := m.Body.Get(f)
	return f, err
}

//GetRedemptionDate reads a RedemptionDate from Allocation.
func (m Allocation) GetRedemptionDate(f *field.RedemptionDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StrikePrice is a non-required field for Allocation.
func (m Allocation) StrikePrice() (*field.StrikePrice, errors.MessageRejectError) {
	f := new(field.StrikePrice)
	err := m.Body.Get(f)
	return f, err
}

//GetStrikePrice reads a StrikePrice from Allocation.
func (m Allocation) GetStrikePrice(f *field.StrikePrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//OptAttribute is a non-required field for Allocation.
func (m Allocation) OptAttribute() (*field.OptAttribute, errors.MessageRejectError) {
	f := new(field.OptAttribute)
	err := m.Body.Get(f)
	return f, err
}

//GetOptAttribute reads a OptAttribute from Allocation.
func (m Allocation) GetOptAttribute(f *field.OptAttribute) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContractMultiplier is a non-required field for Allocation.
func (m Allocation) ContractMultiplier() (*field.ContractMultiplier, errors.MessageRejectError) {
	f := new(field.ContractMultiplier)
	err := m.Body.Get(f)
	return f, err
}

//GetContractMultiplier reads a ContractMultiplier from Allocation.
func (m Allocation) GetContractMultiplier(f *field.ContractMultiplier) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CouponRate is a non-required field for Allocation.
func (m Allocation) CouponRate() (*field.CouponRate, errors.MessageRejectError) {
	f := new(field.CouponRate)
	err := m.Body.Get(f)
	return f, err
}

//GetCouponRate reads a CouponRate from Allocation.
func (m Allocation) GetCouponRate(f *field.CouponRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityExchange is a non-required field for Allocation.
func (m Allocation) SecurityExchange() (*field.SecurityExchange, errors.MessageRejectError) {
	f := new(field.SecurityExchange)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityExchange reads a SecurityExchange from Allocation.
func (m Allocation) GetSecurityExchange(f *field.SecurityExchange) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Issuer is a non-required field for Allocation.
func (m Allocation) Issuer() (*field.Issuer, errors.MessageRejectError) {
	f := new(field.Issuer)
	err := m.Body.Get(f)
	return f, err
}

//GetIssuer reads a Issuer from Allocation.
func (m Allocation) GetIssuer(f *field.Issuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuerLen is a non-required field for Allocation.
func (m Allocation) EncodedIssuerLen() (*field.EncodedIssuerLen, errors.MessageRejectError) {
	f := new(field.EncodedIssuerLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuerLen reads a EncodedIssuerLen from Allocation.
func (m Allocation) GetEncodedIssuerLen(f *field.EncodedIssuerLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedIssuer is a non-required field for Allocation.
func (m Allocation) EncodedIssuer() (*field.EncodedIssuer, errors.MessageRejectError) {
	f := new(field.EncodedIssuer)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedIssuer reads a EncodedIssuer from Allocation.
func (m Allocation) GetEncodedIssuer(f *field.EncodedIssuer) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityDesc is a non-required field for Allocation.
func (m Allocation) SecurityDesc() (*field.SecurityDesc, errors.MessageRejectError) {
	f := new(field.SecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityDesc reads a SecurityDesc from Allocation.
func (m Allocation) GetSecurityDesc(f *field.SecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDescLen is a non-required field for Allocation.
func (m Allocation) EncodedSecurityDescLen() (*field.EncodedSecurityDescLen, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDescLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDescLen reads a EncodedSecurityDescLen from Allocation.
func (m Allocation) GetEncodedSecurityDescLen(f *field.EncodedSecurityDescLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityDesc is a non-required field for Allocation.
func (m Allocation) EncodedSecurityDesc() (*field.EncodedSecurityDesc, errors.MessageRejectError) {
	f := new(field.EncodedSecurityDesc)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityDesc reads a EncodedSecurityDesc from Allocation.
func (m Allocation) GetEncodedSecurityDesc(f *field.EncodedSecurityDesc) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Quantity is a required field for Allocation.
func (m Allocation) Quantity() (*field.Quantity, errors.MessageRejectError) {
	f := new(field.Quantity)
	err := m.Body.Get(f)
	return f, err
}

//GetQuantity reads a Quantity from Allocation.
func (m Allocation) GetQuantity(f *field.Quantity) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastMkt is a non-required field for Allocation.
func (m Allocation) LastMkt() (*field.LastMkt, errors.MessageRejectError) {
	f := new(field.LastMkt)
	err := m.Body.Get(f)
	return f, err
}

//GetLastMkt reads a LastMkt from Allocation.
func (m Allocation) GetLastMkt(f *field.LastMkt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeOriginationDate is a non-required field for Allocation.
func (m Allocation) TradeOriginationDate() (*field.TradeOriginationDate, errors.MessageRejectError) {
	f := new(field.TradeOriginationDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeOriginationDate reads a TradeOriginationDate from Allocation.
func (m Allocation) GetTradeOriginationDate(f *field.TradeOriginationDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a non-required field for Allocation.
func (m Allocation) TradingSessionID() (*field.TradingSessionID, errors.MessageRejectError) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from Allocation.
func (m Allocation) GetTradingSessionID(f *field.TradingSessionID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for Allocation.
func (m Allocation) TradingSessionSubID() (*field.TradingSessionSubID, errors.MessageRejectError) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from Allocation.
func (m Allocation) GetTradingSessionSubID(f *field.TradingSessionSubID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriceType is a non-required field for Allocation.
func (m Allocation) PriceType() (*field.PriceType, errors.MessageRejectError) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}

//GetPriceType reads a PriceType from Allocation.
func (m Allocation) GetPriceType(f *field.PriceType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPx is a required field for Allocation.
func (m Allocation) AvgPx() (*field.AvgPx, errors.MessageRejectError) {
	f := new(field.AvgPx)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPx reads a AvgPx from Allocation.
func (m Allocation) GetAvgPx(f *field.AvgPx) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Currency is a non-required field for Allocation.
func (m Allocation) Currency() (*field.Currency, errors.MessageRejectError) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}

//GetCurrency reads a Currency from Allocation.
func (m Allocation) GetCurrency(f *field.Currency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AvgPrxPrecision is a non-required field for Allocation.
func (m Allocation) AvgPrxPrecision() (*field.AvgPrxPrecision, errors.MessageRejectError) {
	f := new(field.AvgPrxPrecision)
	err := m.Body.Get(f)
	return f, err
}

//GetAvgPrxPrecision reads a AvgPrxPrecision from Allocation.
func (m Allocation) GetAvgPrxPrecision(f *field.AvgPrxPrecision) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for Allocation.
func (m Allocation) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from Allocation.
func (m Allocation) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for Allocation.
func (m Allocation) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from Allocation.
func (m Allocation) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for Allocation.
func (m Allocation) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from Allocation.
func (m Allocation) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlmntTyp is a non-required field for Allocation.
func (m Allocation) SettlmntTyp() (*field.SettlmntTyp, errors.MessageRejectError) {
	f := new(field.SettlmntTyp)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlmntTyp reads a SettlmntTyp from Allocation.
func (m Allocation) GetSettlmntTyp(f *field.SettlmntTyp) errors.MessageRejectError {
	return m.Body.Get(f)
}

//FutSettDate is a non-required field for Allocation.
func (m Allocation) FutSettDate() (*field.FutSettDate, errors.MessageRejectError) {
	f := new(field.FutSettDate)
	err := m.Body.Get(f)
	return f, err
}

//GetFutSettDate reads a FutSettDate from Allocation.
func (m Allocation) GetFutSettDate(f *field.FutSettDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//GrossTradeAmt is a non-required field for Allocation.
func (m Allocation) GrossTradeAmt() (*field.GrossTradeAmt, errors.MessageRejectError) {
	f := new(field.GrossTradeAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetGrossTradeAmt reads a GrossTradeAmt from Allocation.
func (m Allocation) GetGrossTradeAmt(f *field.GrossTradeAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Concession is a non-required field for Allocation.
func (m Allocation) Concession() (*field.Concession, errors.MessageRejectError) {
	f := new(field.Concession)
	err := m.Body.Get(f)
	return f, err
}

//GetConcession reads a Concession from Allocation.
func (m Allocation) GetConcession(f *field.Concession) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalTakedown is a non-required field for Allocation.
func (m Allocation) TotalTakedown() (*field.TotalTakedown, errors.MessageRejectError) {
	f := new(field.TotalTakedown)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalTakedown reads a TotalTakedown from Allocation.
func (m Allocation) GetTotalTakedown(f *field.TotalTakedown) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NetMoney is a non-required field for Allocation.
func (m Allocation) NetMoney() (*field.NetMoney, errors.MessageRejectError) {
	f := new(field.NetMoney)
	err := m.Body.Get(f)
	return f, err
}

//GetNetMoney reads a NetMoney from Allocation.
func (m Allocation) GetNetMoney(f *field.NetMoney) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PositionEffect is a non-required field for Allocation.
func (m Allocation) PositionEffect() (*field.PositionEffect, errors.MessageRejectError) {
	f := new(field.PositionEffect)
	err := m.Body.Get(f)
	return f, err
}

//GetPositionEffect reads a PositionEffect from Allocation.
func (m Allocation) GetPositionEffect(f *field.PositionEffect) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for Allocation.
func (m Allocation) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from Allocation.
func (m Allocation) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for Allocation.
func (m Allocation) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from Allocation.
func (m Allocation) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for Allocation.
func (m Allocation) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from Allocation.
func (m Allocation) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NumDaysInterest is a non-required field for Allocation.
func (m Allocation) NumDaysInterest() (*field.NumDaysInterest, errors.MessageRejectError) {
	f := new(field.NumDaysInterest)
	err := m.Body.Get(f)
	return f, err
}

//GetNumDaysInterest reads a NumDaysInterest from Allocation.
func (m Allocation) GetNumDaysInterest(f *field.NumDaysInterest) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AccruedInterestRate is a non-required field for Allocation.
func (m Allocation) AccruedInterestRate() (*field.AccruedInterestRate, errors.MessageRejectError) {
	f := new(field.AccruedInterestRate)
	err := m.Body.Get(f)
	return f, err
}

//GetAccruedInterestRate reads a AccruedInterestRate from Allocation.
func (m Allocation) GetAccruedInterestRate(f *field.AccruedInterestRate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalAccruedInterestAmt is a non-required field for Allocation.
func (m Allocation) TotalAccruedInterestAmt() (*field.TotalAccruedInterestAmt, errors.MessageRejectError) {
	f := new(field.TotalAccruedInterestAmt)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalAccruedInterestAmt reads a TotalAccruedInterestAmt from Allocation.
func (m Allocation) GetTotalAccruedInterestAmt(f *field.TotalAccruedInterestAmt) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LegalConfirm is a non-required field for Allocation.
func (m Allocation) LegalConfirm() (*field.LegalConfirm, errors.MessageRejectError) {
	f := new(field.LegalConfirm)
	err := m.Body.Get(f)
	return f, err
}

//GetLegalConfirm reads a LegalConfirm from Allocation.
func (m Allocation) GetLegalConfirm(f *field.LegalConfirm) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAllocs is a non-required field for Allocation.
func (m Allocation) NoAllocs() (*field.NoAllocs, errors.MessageRejectError) {
	f := new(field.NoAllocs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAllocs reads a NoAllocs from Allocation.
func (m Allocation) GetNoAllocs(f *field.NoAllocs) errors.MessageRejectError {
	return m.Body.Get(f)
}
