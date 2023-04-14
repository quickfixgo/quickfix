package ordermassstatusrequest

import (
	"github.com/shopspring/decimal"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fix43"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// OrderMassStatusRequest is the fix43 OrderMassStatusRequest type, MsgType = AF.
type OrderMassStatusRequest struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

// FromMessage creates a OrderMassStatusRequest from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) OrderMassStatusRequest {
	return OrderMassStatusRequest{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m OrderMassStatusRequest) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a OrderMassStatusRequest initialized with the required fields for OrderMassStatusRequest.
func New(massstatusreqid field.MassStatusReqIDField, massstatusreqtype field.MassStatusReqTypeField) (m OrderMassStatusRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("AF"))
	m.Set(massstatusreqid)
	m.Set(massstatusreqtype)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg OrderMassStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "AF", r
}

// SetAccount sets Account, Tag 1.
func (m OrderMassStatusRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

// SetSecurityIDSource sets SecurityIDSource, Tag 22.
func (m OrderMassStatusRequest) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m OrderMassStatusRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetSide sets Side, Tag 54.
func (m OrderMassStatusRequest) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

// SetSymbol sets Symbol, Tag 55.
func (m OrderMassStatusRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m OrderMassStatusRequest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m OrderMassStatusRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m OrderMassStatusRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m OrderMassStatusRequest) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m OrderMassStatusRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m OrderMassStatusRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m OrderMassStatusRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m OrderMassStatusRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetCouponRate sets CouponRate, Tag 223.
func (m OrderMassStatusRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

// SetCouponPaymentDate sets CouponPaymentDate, Tag 224.
func (m OrderMassStatusRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

// SetIssueDate sets IssueDate, Tag 225.
func (m OrderMassStatusRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

// SetRepurchaseTerm sets RepurchaseTerm, Tag 226.
func (m OrderMassStatusRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

// SetRepurchaseRate sets RepurchaseRate, Tag 227.
func (m OrderMassStatusRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

// SetFactor sets Factor, Tag 228.
func (m OrderMassStatusRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

// SetContractMultiplier sets ContractMultiplier, Tag 231.
func (m OrderMassStatusRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

// SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239.
func (m OrderMassStatusRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

// SetRedemptionDate sets RedemptionDate, Tag 240.
func (m OrderMassStatusRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

// SetUnderlyingCouponPaymentDate sets UnderlyingCouponPaymentDate, Tag 241.
func (m OrderMassStatusRequest) SetUnderlyingCouponPaymentDate(v string) {
	m.Set(field.NewUnderlyingCouponPaymentDate(v))
}

// SetUnderlyingIssueDate sets UnderlyingIssueDate, Tag 242.
func (m OrderMassStatusRequest) SetUnderlyingIssueDate(v string) {
	m.Set(field.NewUnderlyingIssueDate(v))
}

// SetUnderlyingRepoCollateralSecurityType sets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m OrderMassStatusRequest) SetUnderlyingRepoCollateralSecurityType(v int) {
	m.Set(field.NewUnderlyingRepoCollateralSecurityType(v))
}

// SetUnderlyingRepurchaseTerm sets UnderlyingRepurchaseTerm, Tag 244.
func (m OrderMassStatusRequest) SetUnderlyingRepurchaseTerm(v int) {
	m.Set(field.NewUnderlyingRepurchaseTerm(v))
}

// SetUnderlyingRepurchaseRate sets UnderlyingRepurchaseRate, Tag 245.
func (m OrderMassStatusRequest) SetUnderlyingRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingRepurchaseRate(value, scale))
}

// SetUnderlyingFactor sets UnderlyingFactor, Tag 246.
func (m OrderMassStatusRequest) SetUnderlyingFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingFactor(value, scale))
}

// SetUnderlyingRedemptionDate sets UnderlyingRedemptionDate, Tag 247.
func (m OrderMassStatusRequest) SetUnderlyingRedemptionDate(v string) {
	m.Set(field.NewUnderlyingRedemptionDate(v))
}

// SetCreditRating sets CreditRating, Tag 255.
func (m OrderMassStatusRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

// SetUnderlyingCreditRating sets UnderlyingCreditRating, Tag 256.
func (m OrderMassStatusRequest) SetUnderlyingCreditRating(v string) {
	m.Set(field.NewUnderlyingCreditRating(v))
}

// SetUnderlyingSecurityIDSource sets UnderlyingSecurityIDSource, Tag 305.
func (m OrderMassStatusRequest) SetUnderlyingSecurityIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityIDSource(v))
}

// SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306.
func (m OrderMassStatusRequest) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

// SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307.
func (m OrderMassStatusRequest) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

// SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308.
func (m OrderMassStatusRequest) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

// SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309.
func (m OrderMassStatusRequest) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

// SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310.
func (m OrderMassStatusRequest) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

// SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311.
func (m OrderMassStatusRequest) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

// SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312.
func (m OrderMassStatusRequest) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

// SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313.
func (m OrderMassStatusRequest) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

// SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316.
func (m OrderMassStatusRequest) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

// SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317.
func (m OrderMassStatusRequest) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

// SetTradingSessionID sets TradingSessionID, Tag 336.
func (m OrderMassStatusRequest) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

// SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348.
func (m OrderMassStatusRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

// SetEncodedIssuer sets EncodedIssuer, Tag 349.
func (m OrderMassStatusRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

// SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350.
func (m OrderMassStatusRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

// SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351.
func (m OrderMassStatusRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

// SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362.
func (m OrderMassStatusRequest) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

// SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363.
func (m OrderMassStatusRequest) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

// SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m OrderMassStatusRequest) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

// SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365.
func (m OrderMassStatusRequest) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

// SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435.
func (m OrderMassStatusRequest) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

// SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436.
func (m OrderMassStatusRequest) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

// SetNoPartyIDs sets NoPartyIDs, Tag 453.
func (m OrderMassStatusRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoSecurityAltID sets NoSecurityAltID, Tag 454.
func (m OrderMassStatusRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetNoUnderlyingSecurityAltID sets NoUnderlyingSecurityAltID, Tag 457.
func (m OrderMassStatusRequest) SetNoUnderlyingSecurityAltID(f NoUnderlyingSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

// SetProduct sets Product, Tag 460.
func (m OrderMassStatusRequest) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

// SetCFICode sets CFICode, Tag 461.
func (m OrderMassStatusRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

// SetUnderlyingProduct sets UnderlyingProduct, Tag 462.
func (m OrderMassStatusRequest) SetUnderlyingProduct(v int) {
	m.Set(field.NewUnderlyingProduct(v))
}

// SetUnderlyingCFICode sets UnderlyingCFICode, Tag 463.
func (m OrderMassStatusRequest) SetUnderlyingCFICode(v string) {
	m.Set(field.NewUnderlyingCFICode(v))
}

// SetCountryOfIssue sets CountryOfIssue, Tag 470.
func (m OrderMassStatusRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

// SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471.
func (m OrderMassStatusRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

// SetLocaleOfIssue sets LocaleOfIssue, Tag 472.
func (m OrderMassStatusRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

// SetMaturityDate sets MaturityDate, Tag 541.
func (m OrderMassStatusRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

// SetUnderlyingMaturityDate sets UnderlyingMaturityDate, Tag 542.
func (m OrderMassStatusRequest) SetUnderlyingMaturityDate(v string) {
	m.Set(field.NewUnderlyingMaturityDate(v))
}

// SetInstrRegistry sets InstrRegistry, Tag 543.
func (m OrderMassStatusRequest) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

// SetMassStatusReqID sets MassStatusReqID, Tag 584.
func (m OrderMassStatusRequest) SetMassStatusReqID(v string) {
	m.Set(field.NewMassStatusReqID(v))
}

// SetMassStatusReqType sets MassStatusReqType, Tag 585.
func (m OrderMassStatusRequest) SetMassStatusReqType(v enum.MassStatusReqType) {
	m.Set(field.NewMassStatusReqType(v))
}

// SetUnderlyingCountryOfIssue sets UnderlyingCountryOfIssue, Tag 592.
func (m OrderMassStatusRequest) SetUnderlyingCountryOfIssue(v string) {
	m.Set(field.NewUnderlyingCountryOfIssue(v))
}

// SetUnderlyingStateOrProvinceOfIssue sets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m OrderMassStatusRequest) SetUnderlyingStateOrProvinceOfIssue(v string) {
	m.Set(field.NewUnderlyingStateOrProvinceOfIssue(v))
}

// SetUnderlyingLocaleOfIssue sets UnderlyingLocaleOfIssue, Tag 594.
func (m OrderMassStatusRequest) SetUnderlyingLocaleOfIssue(v string) {
	m.Set(field.NewUnderlyingLocaleOfIssue(v))
}

// SetUnderlyingInstrRegistry sets UnderlyingInstrRegistry, Tag 595.
func (m OrderMassStatusRequest) SetUnderlyingInstrRegistry(v string) {
	m.Set(field.NewUnderlyingInstrRegistry(v))
}

// SetTradingSessionSubID sets TradingSessionSubID, Tag 625.
func (m OrderMassStatusRequest) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

// GetAccount gets Account, Tag 1.
func (m OrderMassStatusRequest) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityIDSource gets SecurityIDSource, Tag 22.
func (m OrderMassStatusRequest) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m OrderMassStatusRequest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSide gets Side, Tag 54.
func (m OrderMassStatusRequest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbol gets Symbol, Tag 55.
func (m OrderMassStatusRequest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m OrderMassStatusRequest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m OrderMassStatusRequest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m OrderMassStatusRequest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m OrderMassStatusRequest) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m OrderMassStatusRequest) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m OrderMassStatusRequest) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m OrderMassStatusRequest) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m OrderMassStatusRequest) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponRate gets CouponRate, Tag 223.
func (m OrderMassStatusRequest) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCouponPaymentDate gets CouponPaymentDate, Tag 224.
func (m OrderMassStatusRequest) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssueDate gets IssueDate, Tag 225.
func (m OrderMassStatusRequest) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseTerm gets RepurchaseTerm, Tag 226.
func (m OrderMassStatusRequest) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepurchaseRate gets RepurchaseRate, Tag 227.
func (m OrderMassStatusRequest) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetFactor gets Factor, Tag 228.
func (m OrderMassStatusRequest) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetContractMultiplier gets ContractMultiplier, Tag 231.
func (m OrderMassStatusRequest) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239.
func (m OrderMassStatusRequest) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRedemptionDate gets RedemptionDate, Tag 240.
func (m OrderMassStatusRequest) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponPaymentDate gets UnderlyingCouponPaymentDate, Tag 241.
func (m OrderMassStatusRequest) GetUnderlyingCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssueDate gets UnderlyingIssueDate, Tag 242.
func (m OrderMassStatusRequest) GetUnderlyingIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepoCollateralSecurityType gets UnderlyingRepoCollateralSecurityType, Tag 243.
func (m OrderMassStatusRequest) GetUnderlyingRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseTerm gets UnderlyingRepurchaseTerm, Tag 244.
func (m OrderMassStatusRequest) GetUnderlyingRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRepurchaseRate gets UnderlyingRepurchaseRate, Tag 245.
func (m OrderMassStatusRequest) GetUnderlyingRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingRepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingFactor gets UnderlyingFactor, Tag 246.
func (m OrderMassStatusRequest) GetUnderlyingFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingFactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingRedemptionDate gets UnderlyingRedemptionDate, Tag 247.
func (m OrderMassStatusRequest) GetUnderlyingRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingRedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCreditRating gets CreditRating, Tag 255.
func (m OrderMassStatusRequest) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCreditRating gets UnderlyingCreditRating, Tag 256.
func (m OrderMassStatusRequest) GetUnderlyingCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityIDSource gets UnderlyingSecurityIDSource, Tag 305.
func (m OrderMassStatusRequest) GetUnderlyingSecurityIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306.
func (m OrderMassStatusRequest) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307.
func (m OrderMassStatusRequest) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308.
func (m OrderMassStatusRequest) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309.
func (m OrderMassStatusRequest) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310.
func (m OrderMassStatusRequest) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311.
func (m OrderMassStatusRequest) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312.
func (m OrderMassStatusRequest) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313.
func (m OrderMassStatusRequest) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316.
func (m OrderMassStatusRequest) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317.
func (m OrderMassStatusRequest) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionID gets TradingSessionID, Tag 336.
func (m OrderMassStatusRequest) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348.
func (m OrderMassStatusRequest) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedIssuer gets EncodedIssuer, Tag 349.
func (m OrderMassStatusRequest) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350.
func (m OrderMassStatusRequest) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351.
func (m OrderMassStatusRequest) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362.
func (m OrderMassStatusRequest) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363.
func (m OrderMassStatusRequest) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364.
func (m OrderMassStatusRequest) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365.
func (m OrderMassStatusRequest) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435.
func (m OrderMassStatusRequest) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436.
func (m OrderMassStatusRequest) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoPartyIDs gets NoPartyIDs, Tag 453.
func (m OrderMassStatusRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoSecurityAltID gets NoSecurityAltID, Tag 454.
func (m OrderMassStatusRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetNoUnderlyingSecurityAltID gets NoUnderlyingSecurityAltID, Tag 457.
func (m OrderMassStatusRequest) GetNoUnderlyingSecurityAltID() (NoUnderlyingSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoUnderlyingSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetProduct gets Product, Tag 460.
func (m OrderMassStatusRequest) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCFICode gets CFICode, Tag 461.
func (m OrderMassStatusRequest) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingProduct gets UnderlyingProduct, Tag 462.
func (m OrderMassStatusRequest) GetUnderlyingProduct() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCFICode gets UnderlyingCFICode, Tag 463.
func (m OrderMassStatusRequest) GetUnderlyingCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetCountryOfIssue gets CountryOfIssue, Tag 470.
func (m OrderMassStatusRequest) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471.
func (m OrderMassStatusRequest) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetLocaleOfIssue gets LocaleOfIssue, Tag 472.
func (m OrderMassStatusRequest) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDate gets MaturityDate, Tag 541.
func (m OrderMassStatusRequest) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingMaturityDate gets UnderlyingMaturityDate, Tag 542.
func (m OrderMassStatusRequest) GetUnderlyingMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetInstrRegistry gets InstrRegistry, Tag 543.
func (m OrderMassStatusRequest) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMassStatusReqID gets MassStatusReqID, Tag 584.
func (m OrderMassStatusRequest) GetMassStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MassStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMassStatusReqType gets MassStatusReqType, Tag 585.
func (m OrderMassStatusRequest) GetMassStatusReqType() (v enum.MassStatusReqType, err quickfix.MessageRejectError) {
	var f field.MassStatusReqTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingCountryOfIssue gets UnderlyingCountryOfIssue, Tag 592.
func (m OrderMassStatusRequest) GetUnderlyingCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingCountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingStateOrProvinceOfIssue gets UnderlyingStateOrProvinceOfIssue, Tag 593.
func (m OrderMassStatusRequest) GetUnderlyingStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingStateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingLocaleOfIssue gets UnderlyingLocaleOfIssue, Tag 594.
func (m OrderMassStatusRequest) GetUnderlyingLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingLocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingInstrRegistry gets UnderlyingInstrRegistry, Tag 595.
func (m OrderMassStatusRequest) GetUnderlyingInstrRegistry() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingInstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetTradingSessionSubID gets TradingSessionSubID, Tag 625.
func (m OrderMassStatusRequest) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasAccount returns true if Account is present, Tag 1.
func (m OrderMassStatusRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

// HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22.
func (m OrderMassStatusRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m OrderMassStatusRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasSide returns true if Side is present, Tag 54.
func (m OrderMassStatusRequest) HasSide() bool {
	return m.Has(tag.Side)
}

// HasSymbol returns true if Symbol is present, Tag 55.
func (m OrderMassStatusRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m OrderMassStatusRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m OrderMassStatusRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m OrderMassStatusRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m OrderMassStatusRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m OrderMassStatusRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m OrderMassStatusRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m OrderMassStatusRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m OrderMassStatusRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasCouponRate returns true if CouponRate is present, Tag 223.
func (m OrderMassStatusRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

// HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224.
func (m OrderMassStatusRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

// HasIssueDate returns true if IssueDate is present, Tag 225.
func (m OrderMassStatusRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

// HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226.
func (m OrderMassStatusRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

// HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227.
func (m OrderMassStatusRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

// HasFactor returns true if Factor is present, Tag 228.
func (m OrderMassStatusRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

// HasContractMultiplier returns true if ContractMultiplier is present, Tag 231.
func (m OrderMassStatusRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

// HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239.
func (m OrderMassStatusRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

// HasRedemptionDate returns true if RedemptionDate is present, Tag 240.
func (m OrderMassStatusRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

// HasUnderlyingCouponPaymentDate returns true if UnderlyingCouponPaymentDate is present, Tag 241.
func (m OrderMassStatusRequest) HasUnderlyingCouponPaymentDate() bool {
	return m.Has(tag.UnderlyingCouponPaymentDate)
}

// HasUnderlyingIssueDate returns true if UnderlyingIssueDate is present, Tag 242.
func (m OrderMassStatusRequest) HasUnderlyingIssueDate() bool {
	return m.Has(tag.UnderlyingIssueDate)
}

// HasUnderlyingRepoCollateralSecurityType returns true if UnderlyingRepoCollateralSecurityType is present, Tag 243.
func (m OrderMassStatusRequest) HasUnderlyingRepoCollateralSecurityType() bool {
	return m.Has(tag.UnderlyingRepoCollateralSecurityType)
}

// HasUnderlyingRepurchaseTerm returns true if UnderlyingRepurchaseTerm is present, Tag 244.
func (m OrderMassStatusRequest) HasUnderlyingRepurchaseTerm() bool {
	return m.Has(tag.UnderlyingRepurchaseTerm)
}

// HasUnderlyingRepurchaseRate returns true if UnderlyingRepurchaseRate is present, Tag 245.
func (m OrderMassStatusRequest) HasUnderlyingRepurchaseRate() bool {
	return m.Has(tag.UnderlyingRepurchaseRate)
}

// HasUnderlyingFactor returns true if UnderlyingFactor is present, Tag 246.
func (m OrderMassStatusRequest) HasUnderlyingFactor() bool {
	return m.Has(tag.UnderlyingFactor)
}

// HasUnderlyingRedemptionDate returns true if UnderlyingRedemptionDate is present, Tag 247.
func (m OrderMassStatusRequest) HasUnderlyingRedemptionDate() bool {
	return m.Has(tag.UnderlyingRedemptionDate)
}

// HasCreditRating returns true if CreditRating is present, Tag 255.
func (m OrderMassStatusRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

// HasUnderlyingCreditRating returns true if UnderlyingCreditRating is present, Tag 256.
func (m OrderMassStatusRequest) HasUnderlyingCreditRating() bool {
	return m.Has(tag.UnderlyingCreditRating)
}

// HasUnderlyingSecurityIDSource returns true if UnderlyingSecurityIDSource is present, Tag 305.
func (m OrderMassStatusRequest) HasUnderlyingSecurityIDSource() bool {
	return m.Has(tag.UnderlyingSecurityIDSource)
}

// HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306.
func (m OrderMassStatusRequest) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

// HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307.
func (m OrderMassStatusRequest) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

// HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308.
func (m OrderMassStatusRequest) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

// HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309.
func (m OrderMassStatusRequest) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

// HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310.
func (m OrderMassStatusRequest) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

// HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311.
func (m OrderMassStatusRequest) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

// HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312.
func (m OrderMassStatusRequest) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

// HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313.
func (m OrderMassStatusRequest) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

// HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316.
func (m OrderMassStatusRequest) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

// HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317.
func (m OrderMassStatusRequest) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

// HasTradingSessionID returns true if TradingSessionID is present, Tag 336.
func (m OrderMassStatusRequest) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

// HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348.
func (m OrderMassStatusRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

// HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349.
func (m OrderMassStatusRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

// HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350.
func (m OrderMassStatusRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

// HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351.
func (m OrderMassStatusRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

// HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362.
func (m OrderMassStatusRequest) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

// HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363.
func (m OrderMassStatusRequest) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

// HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364.
func (m OrderMassStatusRequest) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

// HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365.
func (m OrderMassStatusRequest) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

// HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435.
func (m OrderMassStatusRequest) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

// HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436.
func (m OrderMassStatusRequest) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

// HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453.
func (m OrderMassStatusRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

// HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454.
func (m OrderMassStatusRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

// HasNoUnderlyingSecurityAltID returns true if NoUnderlyingSecurityAltID is present, Tag 457.
func (m OrderMassStatusRequest) HasNoUnderlyingSecurityAltID() bool {
	return m.Has(tag.NoUnderlyingSecurityAltID)
}

// HasProduct returns true if Product is present, Tag 460.
func (m OrderMassStatusRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

// HasCFICode returns true if CFICode is present, Tag 461.
func (m OrderMassStatusRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

// HasUnderlyingProduct returns true if UnderlyingProduct is present, Tag 462.
func (m OrderMassStatusRequest) HasUnderlyingProduct() bool {
	return m.Has(tag.UnderlyingProduct)
}

// HasUnderlyingCFICode returns true if UnderlyingCFICode is present, Tag 463.
func (m OrderMassStatusRequest) HasUnderlyingCFICode() bool {
	return m.Has(tag.UnderlyingCFICode)
}

// HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470.
func (m OrderMassStatusRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

// HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471.
func (m OrderMassStatusRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

// HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472.
func (m OrderMassStatusRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

// HasMaturityDate returns true if MaturityDate is present, Tag 541.
func (m OrderMassStatusRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

// HasUnderlyingMaturityDate returns true if UnderlyingMaturityDate is present, Tag 542.
func (m OrderMassStatusRequest) HasUnderlyingMaturityDate() bool {
	return m.Has(tag.UnderlyingMaturityDate)
}

// HasInstrRegistry returns true if InstrRegistry is present, Tag 543.
func (m OrderMassStatusRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

// HasMassStatusReqID returns true if MassStatusReqID is present, Tag 584.
func (m OrderMassStatusRequest) HasMassStatusReqID() bool {
	return m.Has(tag.MassStatusReqID)
}

// HasMassStatusReqType returns true if MassStatusReqType is present, Tag 585.
func (m OrderMassStatusRequest) HasMassStatusReqType() bool {
	return m.Has(tag.MassStatusReqType)
}

// HasUnderlyingCountryOfIssue returns true if UnderlyingCountryOfIssue is present, Tag 592.
func (m OrderMassStatusRequest) HasUnderlyingCountryOfIssue() bool {
	return m.Has(tag.UnderlyingCountryOfIssue)
}

// HasUnderlyingStateOrProvinceOfIssue returns true if UnderlyingStateOrProvinceOfIssue is present, Tag 593.
func (m OrderMassStatusRequest) HasUnderlyingStateOrProvinceOfIssue() bool {
	return m.Has(tag.UnderlyingStateOrProvinceOfIssue)
}

// HasUnderlyingLocaleOfIssue returns true if UnderlyingLocaleOfIssue is present, Tag 594.
func (m OrderMassStatusRequest) HasUnderlyingLocaleOfIssue() bool {
	return m.Has(tag.UnderlyingLocaleOfIssue)
}

// HasUnderlyingInstrRegistry returns true if UnderlyingInstrRegistry is present, Tag 595.
func (m OrderMassStatusRequest) HasUnderlyingInstrRegistry() bool {
	return m.Has(tag.UnderlyingInstrRegistry)
}

// HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625.
func (m OrderMassStatusRequest) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

// NoPartyIDs is a repeating group element, Tag 453.
type NoPartyIDs struct {
	*quickfix.Group
}

// SetPartyID sets PartyID, Tag 448.
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

// SetPartyIDSource sets PartyIDSource, Tag 447.
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

// SetPartyRole sets PartyRole, Tag 452.
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

// SetPartySubID sets PartySubID, Tag 523.
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

// GetPartyID gets PartyID, Tag 448.
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyIDSource gets PartyIDSource, Tag 447.
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartyRole gets PartyRole, Tag 452.
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPartySubID gets PartySubID, Tag 523.
func (m NoPartyIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasPartyID returns true if PartyID is present, Tag 448.
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

// HasPartyIDSource returns true if PartyIDSource is present, Tag 447.
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

// HasPartyRole returns true if PartyRole is present, Tag 452.
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

// HasPartySubID returns true if PartySubID is present, Tag 523.
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

// NoPartyIDsRepeatingGroup is a repeating group, Tag 453.
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup.
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

// Add create and append a new NoPartyIDs to this group.
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

// Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup.
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

// NoSecurityAltID is a repeating group element, Tag 454.
type NoSecurityAltID struct {
	*quickfix.Group
}

// SetSecurityAltID sets SecurityAltID, Tag 455.
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

// SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

// GetSecurityAltID gets SecurityAltID, Tag 455.
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456.
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasSecurityAltID returns true if SecurityAltID is present, Tag 455.
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

// HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456.
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

// NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454.
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup.
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

// Add create and append a new NoSecurityAltID to this group.
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

// Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup.
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

// NoUnderlyingSecurityAltID is a repeating group element, Tag 457.
type NoUnderlyingSecurityAltID struct {
	*quickfix.Group
}

// SetUnderlyingSecurityAltID sets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltID(v string) {
	m.Set(field.NewUnderlyingSecurityAltID(v))
}

// SetUnderlyingSecurityAltIDSource sets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) SetUnderlyingSecurityAltIDSource(v string) {
	m.Set(field.NewUnderlyingSecurityAltIDSource(v))
}

// GetUnderlyingSecurityAltID gets UnderlyingSecurityAltID, Tag 458.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUnderlyingSecurityAltIDSource gets UnderlyingSecurityAltIDSource, Tag 459.
func (m NoUnderlyingSecurityAltID) GetUnderlyingSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUnderlyingSecurityAltID returns true if UnderlyingSecurityAltID is present, Tag 458.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltID() bool {
	return m.Has(tag.UnderlyingSecurityAltID)
}

// HasUnderlyingSecurityAltIDSource returns true if UnderlyingSecurityAltIDSource is present, Tag 459.
func (m NoUnderlyingSecurityAltID) HasUnderlyingSecurityAltIDSource() bool {
	return m.Has(tag.UnderlyingSecurityAltIDSource)
}

// NoUnderlyingSecurityAltIDRepeatingGroup is a repeating group, Tag 457.
type NoUnderlyingSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoUnderlyingSecurityAltIDRepeatingGroup returns an initialized, NoUnderlyingSecurityAltIDRepeatingGroup.
func NewNoUnderlyingSecurityAltIDRepeatingGroup() NoUnderlyingSecurityAltIDRepeatingGroup {
	return NoUnderlyingSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoUnderlyingSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSecurityAltID), quickfix.GroupElement(tag.UnderlyingSecurityAltIDSource)})}
}

// Add create and append a new NoUnderlyingSecurityAltID to this group.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Add() NoUnderlyingSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoUnderlyingSecurityAltID{g}
}

// Get returns the ith NoUnderlyingSecurityAltID in the NoUnderlyingSecurityAltIDRepeatinGroup.
func (m NoUnderlyingSecurityAltIDRepeatingGroup) Get(i int) NoUnderlyingSecurityAltID {
	return NoUnderlyingSecurityAltID{m.RepeatingGroup.Get(i)}
}
