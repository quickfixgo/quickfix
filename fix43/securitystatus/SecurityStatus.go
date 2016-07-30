package securitystatus

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityStatus is the fix43 SecurityStatus type, MsgType = f
type SecurityStatus struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a SecurityStatus from a quickfix.Message instance
func FromMessage(m quickfix.Message) SecurityStatus {
	return SecurityStatus{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityStatus) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a SecurityStatus initialized with the required fields for SecurityStatus
func New() (m SecurityStatus) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("f"))
	m.Header.Set(field.NewBeginString("FIX.4.3"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "f", r
}

//SetCurrency sets Currency, Tag 15
func (m SecurityStatus) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m SecurityStatus) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
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

//SetText sets Text, Tag 58
func (m SecurityStatus) SetText(v string) {
	m.Set(field.NewText(v))
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

//SetStrikePrice sets StrikePrice, Tag 202
func (m SecurityStatus) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
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

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m SecurityStatus) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m SecurityStatus) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m SecurityStatus) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m SecurityStatus) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m SecurityStatus) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m SecurityStatus) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m SecurityStatus) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m SecurityStatus) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m SecurityStatus) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
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

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SecurityStatus) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SecurityStatus) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m SecurityStatus) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m SecurityStatus) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m SecurityStatus) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m SecurityStatus) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m SecurityStatus) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m SecurityStatus) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m SecurityStatus) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m SecurityStatus) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m SecurityStatus) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetCurrency gets Currency, Tag 15
func (m SecurityStatus) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m SecurityStatus) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
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

//GetText gets Text, Tag 58
func (m SecurityStatus) GetText() (f field.TextField, err quickfix.MessageRejectError) {
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

//GetStrikePrice gets StrikePrice, Tag 202
func (m SecurityStatus) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
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

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m SecurityStatus) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m SecurityStatus) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m SecurityStatus) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m SecurityStatus) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m SecurityStatus) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m SecurityStatus) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m SecurityStatus) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m SecurityStatus) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m SecurityStatus) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
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

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SecurityStatus) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SecurityStatus) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m SecurityStatus) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m SecurityStatus) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m SecurityStatus) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m SecurityStatus) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m SecurityStatus) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m SecurityStatus) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m SecurityStatus) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m SecurityStatus) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m SecurityStatus) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m SecurityStatus) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m SecurityStatus) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
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

//HasText returns true if Text is present, Tag 58
func (m SecurityStatus) HasText() bool {
	return m.Has(tag.Text)
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

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m SecurityStatus) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
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

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m SecurityStatus) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m SecurityStatus) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m SecurityStatus) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m SecurityStatus) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m SecurityStatus) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m SecurityStatus) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m SecurityStatus) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m SecurityStatus) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m SecurityStatus) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
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

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SecurityStatus) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SecurityStatus) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m SecurityStatus) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m SecurityStatus) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m SecurityStatus) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m SecurityStatus) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m SecurityStatus) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m SecurityStatus) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m SecurityStatus) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m SecurityStatus) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m SecurityStatus) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
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
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
