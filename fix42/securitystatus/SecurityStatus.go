package securitystatus

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityStatus is the fix42 SecurityStatus type, MsgType = f
type SecurityStatus struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
}

//FromMessage creates a SecurityStatus from a quickfix.Message instance
func FromMessage(m quickfix.Message) SecurityStatus {
	return SecurityStatus{
		Header:  fix42.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix42.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityStatus) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a SecurityStatus initialized with the required fields for SecurityStatus
func New(symbol field.SymbolField) (m SecurityStatus) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("f"))
	m.Set(symbol)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "f", r
}

//SetCurrency sets Currency, Tag 15
func (m SecurityStatus) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m SecurityStatus) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetLastPx sets LastPx, Tag 31
func (m SecurityStatus) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
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
func (m SecurityStatus) SetSymbolSfx(v string) {
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
func (m SecurityStatus) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m SecurityStatus) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m SecurityStatus) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m SecurityStatus) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
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
func (m SecurityStatus) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m SecurityStatus) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetFinancialStatus sets FinancialStatus, Tag 291
func (m SecurityStatus) SetFinancialStatus(v string) {
	m.Set(field.NewFinancialStatus(v))
}

//SetCorporateAction sets CorporateAction, Tag 292
func (m SecurityStatus) SetCorporateAction(v string) {
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
func (m SecurityStatus) SetSecurityTradingStatus(v int) {
	m.Set(field.NewSecurityTradingStatus(v))
}

//SetHaltReasonChar sets HaltReasonChar, Tag 327
func (m SecurityStatus) SetHaltReasonChar(v string) {
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
func (m SecurityStatus) SetBuyVolume(v float64) {
	m.Set(field.NewBuyVolume(v))
}

//SetSellVolume sets SellVolume, Tag 331
func (m SecurityStatus) SetSellVolume(v float64) {
	m.Set(field.NewSellVolume(v))
}

//SetHighPx sets HighPx, Tag 332
func (m SecurityStatus) SetHighPx(v float64) {
	m.Set(field.NewHighPx(v))
}

//SetLowPx sets LowPx, Tag 333
func (m SecurityStatus) SetLowPx(v float64) {
	m.Set(field.NewLowPx(v))
}

//SetAdjustment sets Adjustment, Tag 334
func (m SecurityStatus) SetAdjustment(v int) {
	m.Set(field.NewAdjustment(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityStatus) SetTradingSessionID(v string) {
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
func (m SecurityStatus) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m SecurityStatus) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m SecurityStatus) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m SecurityStatus) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m SecurityStatus) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m SecurityStatus) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m SecurityStatus) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m SecurityStatus) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m SecurityStatus) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m SecurityStatus) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m SecurityStatus) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m SecurityStatus) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m SecurityStatus) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m SecurityStatus) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m SecurityStatus) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m SecurityStatus) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m SecurityStatus) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m SecurityStatus) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFinancialStatus gets FinancialStatus, Tag 291
func (m SecurityStatus) GetFinancialStatus() (f field.FinancialStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCorporateAction gets CorporateAction, Tag 292
func (m SecurityStatus) GetCorporateAction() (f field.CorporateActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityStatusReqID gets SecurityStatusReqID, Tag 324
func (m SecurityStatus) GetSecurityStatusReqID() (f field.SecurityStatusReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m SecurityStatus) GetUnsolicitedIndicator() (f field.UnsolicitedIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityTradingStatus gets SecurityTradingStatus, Tag 326
func (m SecurityStatus) GetSecurityTradingStatus() (f field.SecurityTradingStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHaltReasonChar gets HaltReasonChar, Tag 327
func (m SecurityStatus) GetHaltReasonChar() (f field.HaltReasonCharField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInViewOfCommon gets InViewOfCommon, Tag 328
func (m SecurityStatus) GetInViewOfCommon() (f field.InViewOfCommonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDueToRelated gets DueToRelated, Tag 329
func (m SecurityStatus) GetDueToRelated() (f field.DueToRelatedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBuyVolume gets BuyVolume, Tag 330
func (m SecurityStatus) GetBuyVolume() (f field.BuyVolumeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSellVolume gets SellVolume, Tag 331
func (m SecurityStatus) GetSellVolume() (f field.SellVolumeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHighPx gets HighPx, Tag 332
func (m SecurityStatus) GetHighPx() (f field.HighPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLowPx gets LowPx, Tag 333
func (m SecurityStatus) GetLowPx() (f field.LowPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAdjustment gets Adjustment, Tag 334
func (m SecurityStatus) GetAdjustment() (f field.AdjustmentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityStatus) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m SecurityStatus) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m SecurityStatus) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m SecurityStatus) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m SecurityStatus) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
