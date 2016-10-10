package allocation

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//Allocation is the fix43 Allocation type, MsgType = J
type Allocation struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Allocation from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Allocation {
	return Allocation{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Allocation) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Allocation initialized with the required fields for Allocation
func New(allocid field.AllocIDField, alloctranstype field.AllocTransTypeField, alloctype field.AllocTypeField, side field.SideField, quantity field.QuantityField, avgpx field.AvgPxField, tradedate field.TradeDateField) (m Allocation) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("J"))
	m.Set(allocid)
	m.Set(alloctranstype)
	m.Set(alloctype)
	m.Set(side)
	m.Set(quantity)
	m.Set(avgpx)
	m.Set(tradedate)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Allocation, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "J", r
}

//SetAvgPx sets AvgPx, Tag 6
func (m Allocation) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m Allocation) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m Allocation) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m Allocation) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Allocation) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetQuantity sets Quantity, Tag 53
func (m Allocation) SetQuantity(value decimal.Decimal, scale int32) {
	m.Set(field.NewQuantity(value, scale))
}

//SetSide sets Side, Tag 54
func (m Allocation) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Allocation) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m Allocation) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Allocation) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m Allocation) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Allocation) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Allocation) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetAllocID sets AllocID, Tag 70
func (m Allocation) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetAllocTransType sets AllocTransType, Tag 71
func (m Allocation) SetAllocTransType(v enum.AllocTransType) {
	m.Set(field.NewAllocTransType(v))
}

//SetRefAllocID sets RefAllocID, Tag 72
func (m Allocation) SetRefAllocID(v string) {
	m.Set(field.NewRefAllocID(v))
}

//SetNoOrders sets NoOrders, Tag 73
func (m Allocation) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

//SetAvgPrxPrecision sets AvgPrxPrecision, Tag 74
func (m Allocation) SetAvgPrxPrecision(v int) {
	m.Set(field.NewAvgPrxPrecision(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m Allocation) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetPositionEffect sets PositionEffect, Tag 77
func (m Allocation) SetPositionEffect(v enum.PositionEffect) {
	m.Set(field.NewPositionEffect(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m Allocation) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetIssuer sets Issuer, Tag 106
func (m Allocation) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Allocation) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m Allocation) SetNetMoney(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetMoney(value, scale))
}

//SetNoExecs sets NoExecs, Tag 124
func (m Allocation) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNumDaysInterest sets NumDaysInterest, Tag 157
func (m Allocation) SetNumDaysInterest(v int) {
	m.Set(field.NewNumDaysInterest(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m Allocation) SetAccruedInterestRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestRate(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m Allocation) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetAllocLinkID sets AllocLinkID, Tag 196
func (m Allocation) SetAllocLinkID(v string) {
	m.Set(field.NewAllocLinkID(v))
}

//SetAllocLinkType sets AllocLinkType, Tag 197
func (m Allocation) SetAllocLinkType(v enum.AllocLinkType) {
	m.Set(field.NewAllocLinkType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m Allocation) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m Allocation) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m Allocation) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m Allocation) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m Allocation) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m Allocation) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m Allocation) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m Allocation) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m Allocation) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m Allocation) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetTradeOriginationDate sets TradeOriginationDate, Tag 229
func (m Allocation) SetTradeOriginationDate(v string) {
	m.Set(field.NewTradeOriginationDate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m Allocation) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetTotalTakedown sets TotalTakedown, Tag 237
func (m Allocation) SetTotalTakedown(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalTakedown(value, scale))
}

//SetConcession sets Concession, Tag 238
func (m Allocation) SetConcession(value decimal.Decimal, scale int32) {
	m.Set(field.NewConcession(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m Allocation) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m Allocation) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m Allocation) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m Allocation) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m Allocation) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m Allocation) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m Allocation) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m Allocation) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m Allocation) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m Allocation) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m Allocation) SetGrossTradeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewGrossTradeAmt(value, scale))
}

//SetPriceType sets PriceType, Tag 423
func (m Allocation) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m Allocation) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m Allocation) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m Allocation) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m Allocation) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetBookingRefID sets BookingRefID, Tag 466
func (m Allocation) SetBookingRefID(v string) {
	m.Set(field.NewBookingRefID(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m Allocation) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m Allocation) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m Allocation) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetTotalAccruedInterestAmt sets TotalAccruedInterestAmt, Tag 540
func (m Allocation) SetTotalAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalAccruedInterestAmt(value, scale))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m Allocation) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m Allocation) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m Allocation) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetAllocType sets AllocType, Tag 626
func (m Allocation) SetAllocType(v enum.AllocType) {
	m.Set(field.NewAllocType(v))
}

//SetLegalConfirm sets LegalConfirm, Tag 650
func (m Allocation) SetLegalConfirm(v bool) {
	m.Set(field.NewLegalConfirm(v))
}

//GetAvgPx gets AvgPx, Tag 6
func (m Allocation) GetAvgPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m Allocation) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m Allocation) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m Allocation) GetLastMkt() (v string, err quickfix.MessageRejectError) {
	var f field.LastMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Allocation) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuantity gets Quantity, Tag 53
func (m Allocation) GetQuantity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.QuantityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m Allocation) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Allocation) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m Allocation) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Allocation) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m Allocation) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Allocation) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Allocation) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m Allocation) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocTransType gets AllocTransType, Tag 71
func (m Allocation) GetAllocTransType() (v enum.AllocTransType, err quickfix.MessageRejectError) {
	var f field.AllocTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefAllocID gets RefAllocID, Tag 72
func (m Allocation) GetRefAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.RefAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m Allocation) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAvgPrxPrecision gets AvgPrxPrecision, Tag 74
func (m Allocation) GetAvgPrxPrecision() (v int, err quickfix.MessageRejectError) {
	var f field.AvgPrxPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m Allocation) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPositionEffect gets PositionEffect, Tag 77
func (m Allocation) GetPositionEffect() (v enum.PositionEffect, err quickfix.MessageRejectError) {
	var f field.PositionEffectField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m Allocation) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetIssuer gets Issuer, Tag 106
func (m Allocation) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Allocation) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m Allocation) GetNetMoney() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NetMoneyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoExecs gets NoExecs, Tag 124
func (m Allocation) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m Allocation) GetNumDaysInterest() (v int, err quickfix.MessageRejectError) {
	var f field.NumDaysInterestField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m Allocation) GetAccruedInterestRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AccruedInterestRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m Allocation) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocLinkID gets AllocLinkID, Tag 196
func (m Allocation) GetAllocLinkID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocLinkIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocLinkType gets AllocLinkType, Tag 197
func (m Allocation) GetAllocLinkType() (v enum.AllocLinkType, err quickfix.MessageRejectError) {
	var f field.AllocLinkTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m Allocation) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m Allocation) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m Allocation) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m Allocation) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m Allocation) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m Allocation) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m Allocation) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m Allocation) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m Allocation) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m Allocation) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeOriginationDate gets TradeOriginationDate, Tag 229
func (m Allocation) GetTradeOriginationDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeOriginationDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m Allocation) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalTakedown gets TotalTakedown, Tag 237
func (m Allocation) GetTotalTakedown() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalTakedownField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetConcession gets Concession, Tag 238
func (m Allocation) GetConcession() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ConcessionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m Allocation) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m Allocation) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m Allocation) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m Allocation) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m Allocation) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m Allocation) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m Allocation) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m Allocation) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m Allocation) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m Allocation) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m Allocation) GetGrossTradeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.GrossTradeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m Allocation) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m Allocation) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m Allocation) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m Allocation) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m Allocation) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBookingRefID gets BookingRefID, Tag 466
func (m Allocation) GetBookingRefID() (v string, err quickfix.MessageRejectError) {
	var f field.BookingRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m Allocation) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m Allocation) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m Allocation) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalAccruedInterestAmt gets TotalAccruedInterestAmt, Tag 540
func (m Allocation) GetTotalAccruedInterestAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalAccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m Allocation) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m Allocation) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m Allocation) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocType gets AllocType, Tag 626
func (m Allocation) GetAllocType() (v enum.AllocType, err quickfix.MessageRejectError) {
	var f field.AllocTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLegalConfirm gets LegalConfirm, Tag 650
func (m Allocation) GetLegalConfirm() (v bool, err quickfix.MessageRejectError) {
	var f field.LegalConfirmField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m Allocation) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m Allocation) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m Allocation) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m Allocation) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Allocation) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasQuantity returns true if Quantity is present, Tag 53
func (m Allocation) HasQuantity() bool {
	return m.Has(tag.Quantity)
}

//HasSide returns true if Side is present, Tag 54
func (m Allocation) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Allocation) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m Allocation) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Allocation) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m Allocation) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m Allocation) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Allocation) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m Allocation) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasAllocTransType returns true if AllocTransType is present, Tag 71
func (m Allocation) HasAllocTransType() bool {
	return m.Has(tag.AllocTransType)
}

//HasRefAllocID returns true if RefAllocID is present, Tag 72
func (m Allocation) HasRefAllocID() bool {
	return m.Has(tag.RefAllocID)
}

//HasNoOrders returns true if NoOrders is present, Tag 73
func (m Allocation) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

//HasAvgPrxPrecision returns true if AvgPrxPrecision is present, Tag 74
func (m Allocation) HasAvgPrxPrecision() bool {
	return m.Has(tag.AvgPrxPrecision)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m Allocation) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasPositionEffect returns true if PositionEffect is present, Tag 77
func (m Allocation) HasPositionEffect() bool {
	return m.Has(tag.PositionEffect)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m Allocation) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Allocation) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Allocation) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m Allocation) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasNoExecs returns true if NoExecs is present, Tag 124
func (m Allocation) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

//HasNumDaysInterest returns true if NumDaysInterest is present, Tag 157
func (m Allocation) HasNumDaysInterest() bool {
	return m.Has(tag.NumDaysInterest)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m Allocation) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m Allocation) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasAllocLinkID returns true if AllocLinkID is present, Tag 196
func (m Allocation) HasAllocLinkID() bool {
	return m.Has(tag.AllocLinkID)
}

//HasAllocLinkType returns true if AllocLinkType is present, Tag 197
func (m Allocation) HasAllocLinkType() bool {
	return m.Has(tag.AllocLinkType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m Allocation) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m Allocation) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m Allocation) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m Allocation) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m Allocation) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m Allocation) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m Allocation) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m Allocation) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m Allocation) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m Allocation) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasTradeOriginationDate returns true if TradeOriginationDate is present, Tag 229
func (m Allocation) HasTradeOriginationDate() bool {
	return m.Has(tag.TradeOriginationDate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m Allocation) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasTotalTakedown returns true if TotalTakedown is present, Tag 237
func (m Allocation) HasTotalTakedown() bool {
	return m.Has(tag.TotalTakedown)
}

//HasConcession returns true if Concession is present, Tag 238
func (m Allocation) HasConcession() bool {
	return m.Has(tag.Concession)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m Allocation) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m Allocation) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m Allocation) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m Allocation) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m Allocation) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m Allocation) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m Allocation) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m Allocation) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m Allocation) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m Allocation) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m Allocation) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m Allocation) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m Allocation) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m Allocation) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m Allocation) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m Allocation) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasBookingRefID returns true if BookingRefID is present, Tag 466
func (m Allocation) HasBookingRefID() bool {
	return m.Has(tag.BookingRefID)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m Allocation) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m Allocation) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m Allocation) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasTotalAccruedInterestAmt returns true if TotalAccruedInterestAmt is present, Tag 540
func (m Allocation) HasTotalAccruedInterestAmt() bool {
	return m.Has(tag.TotalAccruedInterestAmt)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m Allocation) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m Allocation) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m Allocation) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasAllocType returns true if AllocType is present, Tag 626
func (m Allocation) HasAllocType() bool {
	return m.Has(tag.AllocType)
}

//HasLegalConfirm returns true if LegalConfirm is present, Tag 650
func (m Allocation) HasLegalConfirm() bool {
	return m.Has(tag.LegalConfirm)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	*quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoOrders) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m NoOrders) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m NoOrders) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetListID sets ListID, Tag 66
func (m NoOrders) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoOrders) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoOrders) GetSecondaryOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryOrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m NoOrders) GetSecondaryClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m NoOrders) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoOrders) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m NoOrders) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m NoOrders) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasListID returns true if ListID is present, Tag 66
func (m NoOrders) HasListID() bool {
	return m.Has(tag.ListID)
}

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.SecondaryClOrdID), quickfix.GroupElement(tag.ListID)})}
}

//Add create and append a new NoOrders to this group
func (m NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

//Get returns the ith NoOrders in the NoOrdersRepeatinGroup
func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	*quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocPrice sets AllocPrice, Tag 366
func (m NoAllocs) SetAllocPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocPrice(value, scale))
}

//SetAllocQty sets AllocQty, Tag 80
func (m NoAllocs) SetAllocQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocQty(value, scale))
}

//SetIndividualAllocID sets IndividualAllocID, Tag 467
func (m NoAllocs) SetIndividualAllocID(v string) {
	m.Set(field.NewIndividualAllocID(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoAllocs) SetProcessCode(v enum.ProcessCode) {
	m.Set(field.NewProcessCode(v))
}

//SetNoNestedPartyIDs sets NoNestedPartyIDs, Tag 539
func (m NoAllocs) SetNoNestedPartyIDs(f NoNestedPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNotifyBrokerOfCredit sets NotifyBrokerOfCredit, Tag 208
func (m NoAllocs) SetNotifyBrokerOfCredit(v bool) {
	m.Set(field.NewNotifyBrokerOfCredit(v))
}

//SetAllocHandlInst sets AllocHandlInst, Tag 209
func (m NoAllocs) SetAllocHandlInst(v enum.AllocHandlInst) {
	m.Set(field.NewAllocHandlInst(v))
}

//SetAllocText sets AllocText, Tag 161
func (m NoAllocs) SetAllocText(v string) {
	m.Set(field.NewAllocText(v))
}

//SetEncodedAllocTextLen sets EncodedAllocTextLen, Tag 360
func (m NoAllocs) SetEncodedAllocTextLen(v int) {
	m.Set(field.NewEncodedAllocTextLen(v))
}

//SetEncodedAllocText sets EncodedAllocText, Tag 361
func (m NoAllocs) SetEncodedAllocText(v string) {
	m.Set(field.NewEncodedAllocText(v))
}

//SetCommission sets Commission, Tag 12
func (m NoAllocs) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NoAllocs) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetCommCurrency sets CommCurrency, Tag 479
func (m NoAllocs) SetCommCurrency(v string) {
	m.Set(field.NewCommCurrency(v))
}

//SetFundRenewWaiv sets FundRenewWaiv, Tag 497
func (m NoAllocs) SetFundRenewWaiv(v enum.FundRenewWaiv) {
	m.Set(field.NewFundRenewWaiv(v))
}

//SetAllocAvgPx sets AllocAvgPx, Tag 153
func (m NoAllocs) SetAllocAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocAvgPx(value, scale))
}

//SetAllocNetMoney sets AllocNetMoney, Tag 154
func (m NoAllocs) SetAllocNetMoney(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocNetMoney(value, scale))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoAllocs) SetSettlCurrAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrAmt(value, scale))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoAllocs) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoAllocs) SetSettlCurrFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrFxRate(value, scale))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m NoAllocs) SetSettlCurrFxRateCalc(v enum.SettlCurrFxRateCalc) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoAllocs) SetAccruedInterestAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewAccruedInterestAmt(value, scale))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m NoAllocs) SetSettlInstMode(v enum.SettlInstMode) {
	m.Set(field.NewSettlInstMode(v))
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m NoAllocs) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocPrice gets AllocPrice, Tag 366
func (m NoAllocs) GetAllocPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocQty gets AllocQty, Tag 80
func (m NoAllocs) GetAllocQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIndividualAllocID gets IndividualAllocID, Tag 467
func (m NoAllocs) GetIndividualAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.IndividualAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoAllocs) GetProcessCode() (v enum.ProcessCode, err quickfix.MessageRejectError) {
	var f field.ProcessCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoNestedPartyIDs gets NoNestedPartyIDs, Tag 539
func (m NoAllocs) GetNoNestedPartyIDs() (NoNestedPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoNestedPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNotifyBrokerOfCredit gets NotifyBrokerOfCredit, Tag 208
func (m NoAllocs) GetNotifyBrokerOfCredit() (v bool, err quickfix.MessageRejectError) {
	var f field.NotifyBrokerOfCreditField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocHandlInst gets AllocHandlInst, Tag 209
func (m NoAllocs) GetAllocHandlInst() (v enum.AllocHandlInst, err quickfix.MessageRejectError) {
	var f field.AllocHandlInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocText gets AllocText, Tag 161
func (m NoAllocs) GetAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.AllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedAllocTextLen gets EncodedAllocTextLen, Tag 360
func (m NoAllocs) GetEncodedAllocTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedAllocText gets EncodedAllocText, Tag 361
func (m NoAllocs) GetEncodedAllocText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedAllocTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m NoAllocs) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m NoAllocs) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommCurrency gets CommCurrency, Tag 479
func (m NoAllocs) GetCommCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CommCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFundRenewWaiv gets FundRenewWaiv, Tag 497
func (m NoAllocs) GetFundRenewWaiv() (v enum.FundRenewWaiv, err quickfix.MessageRejectError) {
	var f field.FundRenewWaivField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocAvgPx gets AllocAvgPx, Tag 153
func (m NoAllocs) GetAllocAvgPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocAvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocNetMoney gets AllocNetMoney, Tag 154
func (m NoAllocs) GetAllocNetMoney() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocNetMoneyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoAllocs) GetSettlCurrAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoAllocs) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoAllocs) GetSettlCurrFxRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m NoAllocs) GetSettlCurrFxRateCalc() (v enum.SettlCurrFxRateCalc, err quickfix.MessageRejectError) {
	var f field.SettlCurrFxRateCalcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoAllocs) GetAccruedInterestAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AccruedInterestAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m NoAllocs) GetSettlInstMode() (v enum.SettlInstMode, err quickfix.MessageRejectError) {
	var f field.SettlInstModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m NoAllocs) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasAllocPrice returns true if AllocPrice is present, Tag 366
func (m NoAllocs) HasAllocPrice() bool {
	return m.Has(tag.AllocPrice)
}

//HasAllocQty returns true if AllocQty is present, Tag 80
func (m NoAllocs) HasAllocQty() bool {
	return m.Has(tag.AllocQty)
}

//HasIndividualAllocID returns true if IndividualAllocID is present, Tag 467
func (m NoAllocs) HasIndividualAllocID() bool {
	return m.Has(tag.IndividualAllocID)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NoAllocs) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasNoNestedPartyIDs returns true if NoNestedPartyIDs is present, Tag 539
func (m NoAllocs) HasNoNestedPartyIDs() bool {
	return m.Has(tag.NoNestedPartyIDs)
}

//HasNotifyBrokerOfCredit returns true if NotifyBrokerOfCredit is present, Tag 208
func (m NoAllocs) HasNotifyBrokerOfCredit() bool {
	return m.Has(tag.NotifyBrokerOfCredit)
}

//HasAllocHandlInst returns true if AllocHandlInst is present, Tag 209
func (m NoAllocs) HasAllocHandlInst() bool {
	return m.Has(tag.AllocHandlInst)
}

//HasAllocText returns true if AllocText is present, Tag 161
func (m NoAllocs) HasAllocText() bool {
	return m.Has(tag.AllocText)
}

//HasEncodedAllocTextLen returns true if EncodedAllocTextLen is present, Tag 360
func (m NoAllocs) HasEncodedAllocTextLen() bool {
	return m.Has(tag.EncodedAllocTextLen)
}

//HasEncodedAllocText returns true if EncodedAllocText is present, Tag 361
func (m NoAllocs) HasEncodedAllocText() bool {
	return m.Has(tag.EncodedAllocText)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoAllocs) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoAllocs) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCommCurrency returns true if CommCurrency is present, Tag 479
func (m NoAllocs) HasCommCurrency() bool {
	return m.Has(tag.CommCurrency)
}

//HasFundRenewWaiv returns true if FundRenewWaiv is present, Tag 497
func (m NoAllocs) HasFundRenewWaiv() bool {
	return m.Has(tag.FundRenewWaiv)
}

//HasAllocAvgPx returns true if AllocAvgPx is present, Tag 153
func (m NoAllocs) HasAllocAvgPx() bool {
	return m.Has(tag.AllocAvgPx)
}

//HasAllocNetMoney returns true if AllocNetMoney is present, Tag 154
func (m NoAllocs) HasAllocNetMoney() bool {
	return m.Has(tag.AllocNetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m NoAllocs) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoAllocs) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m NoAllocs) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m NoAllocs) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m NoAllocs) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasSettlInstMode returns true if SettlInstMode is present, Tag 160
func (m NoAllocs) HasSettlInstMode() bool {
	return m.Has(tag.SettlInstMode)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m NoAllocs) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//NoNestedPartyIDs is a repeating group element, Tag 539
type NoNestedPartyIDs struct {
	*quickfix.Group
}

//SetNestedPartyID sets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) SetNestedPartyID(v string) {
	m.Set(field.NewNestedPartyID(v))
}

//SetNestedPartyIDSource sets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) SetNestedPartyIDSource(v string) {
	m.Set(field.NewNestedPartyIDSource(v))
}

//SetNestedPartyRole sets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) SetNestedPartyRole(v int) {
	m.Set(field.NewNestedPartyRole(v))
}

//SetNestedPartySubID sets NestedPartySubID, Tag 545
func (m NoNestedPartyIDs) SetNestedPartySubID(v string) {
	m.Set(field.NewNestedPartySubID(v))
}

//GetNestedPartyID gets NestedPartyID, Tag 524
func (m NoNestedPartyIDs) GetNestedPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyIDSource gets NestedPartyIDSource, Tag 525
func (m NoNestedPartyIDs) GetNestedPartyIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartyRole gets NestedPartyRole, Tag 538
func (m NoNestedPartyIDs) GetNestedPartyRole() (v int, err quickfix.MessageRejectError) {
	var f field.NestedPartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNestedPartySubID gets NestedPartySubID, Tag 545
func (m NoNestedPartyIDs) GetNestedPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.NestedPartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasNestedPartyID returns true if NestedPartyID is present, Tag 524
func (m NoNestedPartyIDs) HasNestedPartyID() bool {
	return m.Has(tag.NestedPartyID)
}

//HasNestedPartyIDSource returns true if NestedPartyIDSource is present, Tag 525
func (m NoNestedPartyIDs) HasNestedPartyIDSource() bool {
	return m.Has(tag.NestedPartyIDSource)
}

//HasNestedPartyRole returns true if NestedPartyRole is present, Tag 538
func (m NoNestedPartyIDs) HasNestedPartyRole() bool {
	return m.Has(tag.NestedPartyRole)
}

//HasNestedPartySubID returns true if NestedPartySubID is present, Tag 545
func (m NoNestedPartyIDs) HasNestedPartySubID() bool {
	return m.Has(tag.NestedPartySubID)
}

//NoNestedPartyIDsRepeatingGroup is a repeating group, Tag 539
type NoNestedPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoNestedPartyIDsRepeatingGroup returns an initialized, NoNestedPartyIDsRepeatingGroup
func NewNoNestedPartyIDsRepeatingGroup() NoNestedPartyIDsRepeatingGroup {
	return NoNestedPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoNestedPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.NestedPartyID), quickfix.GroupElement(tag.NestedPartyIDSource), quickfix.GroupElement(tag.NestedPartyRole), quickfix.GroupElement(tag.NestedPartySubID)})}
}

//Add create and append a new NoNestedPartyIDs to this group
func (m NoNestedPartyIDsRepeatingGroup) Add() NoNestedPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoNestedPartyIDs{g}
}

//Get returns the ith NoNestedPartyIDs in the NoNestedPartyIDsRepeatinGroup
func (m NoNestedPartyIDsRepeatingGroup) Get(i int) NoNestedPartyIDs {
	return NoNestedPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	*quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMiscFeeAmt(value, scale))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v enum.MiscFeeType) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MiscFeeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (v string, err quickfix.MessageRejectError) {
	var f field.MiscFeeCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (v enum.MiscFeeType, err quickfix.MessageRejectError) {
	var f field.MiscFeeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocPrice), quickfix.GroupElement(tag.AllocQty), quickfix.GroupElement(tag.IndividualAllocID), quickfix.GroupElement(tag.ProcessCode), NewNoNestedPartyIDsRepeatingGroup(), quickfix.GroupElement(tag.NotifyBrokerOfCredit), quickfix.GroupElement(tag.AllocHandlInst), quickfix.GroupElement(tag.AllocText), quickfix.GroupElement(tag.EncodedAllocTextLen), quickfix.GroupElement(tag.EncodedAllocText), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.CommCurrency), quickfix.GroupElement(tag.FundRenewWaiv), quickfix.GroupElement(tag.AllocAvgPx), quickfix.GroupElement(tag.AllocNetMoney), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlCurrFxRateCalc), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.SettlInstMode), NewNoMiscFeesRepeatingGroup()})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

//NoExecs is a repeating group element, Tag 124
type NoExecs struct {
	*quickfix.Group
}

//SetLastQty sets LastQty, Tag 32
func (m NoExecs) SetLastQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastQty(value, scale))
}

//SetExecID sets ExecID, Tag 17
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetSecondaryExecID sets SecondaryExecID, Tag 527
func (m NoExecs) SetSecondaryExecID(v string) {
	m.Set(field.NewSecondaryExecID(v))
}

//SetLastPx sets LastPx, Tag 31
func (m NoExecs) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastCapacity sets LastCapacity, Tag 29
func (m NoExecs) SetLastCapacity(v enum.LastCapacity) {
	m.Set(field.NewLastCapacity(v))
}

//GetLastQty gets LastQty, Tag 32
func (m NoExecs) GetLastQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecID gets ExecID, Tag 17
func (m NoExecs) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecondaryExecID gets SecondaryExecID, Tag 527
func (m NoExecs) GetSecondaryExecID() (v string, err quickfix.MessageRejectError) {
	var f field.SecondaryExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m NoExecs) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastCapacity gets LastCapacity, Tag 29
func (m NoExecs) GetLastCapacity() (v enum.LastCapacity, err quickfix.MessageRejectError) {
	var f field.LastCapacityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLastQty returns true if LastQty is present, Tag 32
func (m NoExecs) HasLastQty() bool {
	return m.Has(tag.LastQty)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasSecondaryExecID returns true if SecondaryExecID is present, Tag 527
func (m NoExecs) HasSecondaryExecID() bool {
	return m.Has(tag.SecondaryExecID)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m NoExecs) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastCapacity returns true if LastCapacity is present, Tag 29
func (m NoExecs) HasLastCapacity() bool {
	return m.Has(tag.LastCapacity)
}

//NoExecsRepeatingGroup is a repeating group, Tag 124
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LastQty), quickfix.GroupElement(tag.ExecID), quickfix.GroupElement(tag.SecondaryExecID), quickfix.GroupElement(tag.LastPx), quickfix.GroupElement(tag.LastCapacity)})}
}

//Add create and append a new NoExecs to this group
func (m NoExecsRepeatingGroup) Add() NoExecs {
	g := m.RepeatingGroup.Add()
	return NoExecs{g}
}

//Get returns the ith NoExecs in the NoExecsRepeatinGroup
func (m NoExecsRepeatingGroup) Get(i int) NoExecs {
	return NoExecs{m.RepeatingGroup.Get(i)}
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	*quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v enum.PartyIDSource) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v enum.PartyRole) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (v string, err quickfix.MessageRejectError) {
	var f field.PartyIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (v enum.PartyIDSource, err quickfix.MessageRejectError) {
	var f field.PartyIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (v enum.PartyRole, err quickfix.MessageRejectError) {
	var f field.PartyRoleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (v string, err quickfix.MessageRejectError) {
	var f field.PartySubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	*quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}
