package securitystatus

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityStatus is the fix42 SecurityStatus type, MsgType = f
type SecurityStatus struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a SecurityStatus from a quickfix.Message instance
func FromMessage(m *quickfix.Message) SecurityStatus {
	return SecurityStatus{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityStatus) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a SecurityStatus initialized with the required fields for SecurityStatus
func New(symbol field.SymbolField) (m SecurityStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("f"))
	m.Set(symbol)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "f", r
}

//SetCurrency sets Currency, Tag 15
func (m SecurityStatus) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m SecurityStatus) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetLastPx sets LastPx, Tag 31
func (m SecurityStatus) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m SecurityStatus) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m SecurityStatus) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m SecurityStatus) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m SecurityStatus) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m SecurityStatus) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m SecurityStatus) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m SecurityStatus) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m SecurityStatus) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m SecurityStatus) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m SecurityStatus) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m SecurityStatus) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m SecurityStatus) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m SecurityStatus) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m SecurityStatus) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m SecurityStatus) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetFinancialStatus sets FinancialStatus, Tag 291
func (m SecurityStatus) SetFinancialStatus(v enum.FinancialStatus) {
	m.Set(field.NewFinancialStatus(v))
}

//SetCorporateAction sets CorporateAction, Tag 292
func (m SecurityStatus) SetCorporateAction(v enum.CorporateAction) {
	m.Set(field.NewCorporateAction(v))
}

//SetSecurityStatusReqID sets SecurityStatusReqID, Tag 324
func (m SecurityStatus) SetSecurityStatusReqID(v string) {
	m.Set(field.NewSecurityStatusReqID(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m SecurityStatus) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetSecurityTradingStatus sets SecurityTradingStatus, Tag 326
func (m SecurityStatus) SetSecurityTradingStatus(v enum.SecurityTradingStatus) {
	m.Set(field.NewSecurityTradingStatus(v))
}

//SetHaltReasonChar sets HaltReasonChar, Tag 327
func (m SecurityStatus) SetHaltReasonChar(v enum.HaltReasonChar) {
	m.Set(field.NewHaltReasonChar(v))
}

//SetInViewOfCommon sets InViewOfCommon, Tag 328
func (m SecurityStatus) SetInViewOfCommon(v bool) {
	m.Set(field.NewInViewOfCommon(v))
}

//SetDueToRelated sets DueToRelated, Tag 329
func (m SecurityStatus) SetDueToRelated(v bool) {
	m.Set(field.NewDueToRelated(v))
}

//SetBuyVolume sets BuyVolume, Tag 330
func (m SecurityStatus) SetBuyVolume(value decimal.Decimal, scale int32) {
	m.Set(field.NewBuyVolume(value, scale))
}

//SetSellVolume sets SellVolume, Tag 331
func (m SecurityStatus) SetSellVolume(value decimal.Decimal, scale int32) {
	m.Set(field.NewSellVolume(value, scale))
}

//SetHighPx sets HighPx, Tag 332
func (m SecurityStatus) SetHighPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewHighPx(value, scale))
}

//SetLowPx sets LowPx, Tag 333
func (m SecurityStatus) SetLowPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLowPx(value, scale))
}

//SetAdjustment sets Adjustment, Tag 334
func (m SecurityStatus) SetAdjustment(v enum.Adjustment) {
	m.Set(field.NewAdjustment(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityStatus) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m SecurityStatus) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m SecurityStatus) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m SecurityStatus) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m SecurityStatus) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//GetCurrency gets Currency, Tag 15
func (m SecurityStatus) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m SecurityStatus) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m SecurityStatus) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m SecurityStatus) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m SecurityStatus) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SecurityStatus) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m SecurityStatus) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m SecurityStatus) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m SecurityStatus) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SecurityStatus) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m SecurityStatus) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m SecurityStatus) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m SecurityStatus) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m SecurityStatus) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m SecurityStatus) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m SecurityStatus) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m SecurityStatus) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m SecurityStatus) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFinancialStatus gets FinancialStatus, Tag 291
func (m SecurityStatus) GetFinancialStatus() (v enum.FinancialStatus, err quickfix.MessageRejectError) {
	var f field.FinancialStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCorporateAction gets CorporateAction, Tag 292
func (m SecurityStatus) GetCorporateAction() (v enum.CorporateAction, err quickfix.MessageRejectError) {
	var f field.CorporateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityStatusReqID gets SecurityStatusReqID, Tag 324
func (m SecurityStatus) GetSecurityStatusReqID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityStatusReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m SecurityStatus) GetUnsolicitedIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.UnsolicitedIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityTradingStatus gets SecurityTradingStatus, Tag 326
func (m SecurityStatus) GetSecurityTradingStatus() (v enum.SecurityTradingStatus, err quickfix.MessageRejectError) {
	var f field.SecurityTradingStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHaltReasonChar gets HaltReasonChar, Tag 327
func (m SecurityStatus) GetHaltReasonChar() (v enum.HaltReasonChar, err quickfix.MessageRejectError) {
	var f field.HaltReasonCharField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInViewOfCommon gets InViewOfCommon, Tag 328
func (m SecurityStatus) GetInViewOfCommon() (v bool, err quickfix.MessageRejectError) {
	var f field.InViewOfCommonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDueToRelated gets DueToRelated, Tag 329
func (m SecurityStatus) GetDueToRelated() (v bool, err quickfix.MessageRejectError) {
	var f field.DueToRelatedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBuyVolume gets BuyVolume, Tag 330
func (m SecurityStatus) GetBuyVolume() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BuyVolumeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSellVolume gets SellVolume, Tag 331
func (m SecurityStatus) GetSellVolume() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SellVolumeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHighPx gets HighPx, Tag 332
func (m SecurityStatus) GetHighPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.HighPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLowPx gets LowPx, Tag 333
func (m SecurityStatus) GetLowPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LowPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAdjustment gets Adjustment, Tag 334
func (m SecurityStatus) GetAdjustment() (v enum.Adjustment, err quickfix.MessageRejectError) {
	var f field.AdjustmentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityStatus) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m SecurityStatus) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m SecurityStatus) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m SecurityStatus) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m SecurityStatus) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m SecurityStatus) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m SecurityStatus) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m SecurityStatus) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m SecurityStatus) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m SecurityStatus) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m SecurityStatus) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m SecurityStatus) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m SecurityStatus) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m SecurityStatus) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m SecurityStatus) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m SecurityStatus) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m SecurityStatus) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m SecurityStatus) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m SecurityStatus) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m SecurityStatus) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m SecurityStatus) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m SecurityStatus) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m SecurityStatus) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasFinancialStatus returns true if FinancialStatus is present, Tag 291
func (m SecurityStatus) HasFinancialStatus() bool {
	return m.Has(tag.FinancialStatus)
}

//HasCorporateAction returns true if CorporateAction is present, Tag 292
func (m SecurityStatus) HasCorporateAction() bool {
	return m.Has(tag.CorporateAction)
}

//HasSecurityStatusReqID returns true if SecurityStatusReqID is present, Tag 324
func (m SecurityStatus) HasSecurityStatusReqID() bool {
	return m.Has(tag.SecurityStatusReqID)
}

//HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m SecurityStatus) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

//HasSecurityTradingStatus returns true if SecurityTradingStatus is present, Tag 326
func (m SecurityStatus) HasSecurityTradingStatus() bool {
	return m.Has(tag.SecurityTradingStatus)
}

//HasHaltReasonChar returns true if HaltReasonChar is present, Tag 327
func (m SecurityStatus) HasHaltReasonChar() bool {
	return m.Has(tag.HaltReasonChar)
}

//HasInViewOfCommon returns true if InViewOfCommon is present, Tag 328
func (m SecurityStatus) HasInViewOfCommon() bool {
	return m.Has(tag.InViewOfCommon)
}

//HasDueToRelated returns true if DueToRelated is present, Tag 329
func (m SecurityStatus) HasDueToRelated() bool {
	return m.Has(tag.DueToRelated)
}

//HasBuyVolume returns true if BuyVolume is present, Tag 330
func (m SecurityStatus) HasBuyVolume() bool {
	return m.Has(tag.BuyVolume)
}

//HasSellVolume returns true if SellVolume is present, Tag 331
func (m SecurityStatus) HasSellVolume() bool {
	return m.Has(tag.SellVolume)
}

//HasHighPx returns true if HighPx is present, Tag 332
func (m SecurityStatus) HasHighPx() bool {
	return m.Has(tag.HighPx)
}

//HasLowPx returns true if LowPx is present, Tag 333
func (m SecurityStatus) HasLowPx() bool {
	return m.Has(tag.LowPx)
}

//HasAdjustment returns true if Adjustment is present, Tag 334
func (m SecurityStatus) HasAdjustment() bool {
	return m.Has(tag.Adjustment)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SecurityStatus) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m SecurityStatus) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m SecurityStatus) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m SecurityStatus) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m SecurityStatus) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}
