package securitydefinition

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//SecurityDefinition is the fix42 SecurityDefinition type, MsgType = d
type SecurityDefinition struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a SecurityDefinition from a quickfix.Message instance
func FromMessage(m quickfix.Message) SecurityDefinition {
	return SecurityDefinition{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m SecurityDefinition) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a SecurityDefinition initialized with the required fields for SecurityDefinition
func New(securityreqid field.SecurityReqIDField, securityresponseid field.SecurityResponseIDField, totalnumsecurities field.TotalNumSecuritiesField) (m SecurityDefinition) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("d"))
	m.Set(securityreqid)
	m.Set(securityresponseid)
	m.Set(totalnumsecurities)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg SecurityDefinition, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "d", r
}

//SetCurrency sets Currency, Tag 15
func (m SecurityDefinition) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m SecurityDefinition) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m SecurityDefinition) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m SecurityDefinition) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m SecurityDefinition) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m SecurityDefinition) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m SecurityDefinition) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m SecurityDefinition) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m SecurityDefinition) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetSecurityType sets SecurityType, Tag 167
func (m SecurityDefinition) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m SecurityDefinition) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m SecurityDefinition) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m SecurityDefinition) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m SecurityDefinition) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m SecurityDefinition) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m SecurityDefinition) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m SecurityDefinition) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m SecurityDefinition) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetSecurityReqID sets SecurityReqID, Tag 320
func (m SecurityDefinition) SetSecurityReqID(v string) {
	m.Set(field.NewSecurityReqID(v))
}

//SetSecurityResponseID sets SecurityResponseID, Tag 322
func (m SecurityDefinition) SetSecurityResponseID(v string) {
	m.Set(field.NewSecurityResponseID(v))
}

//SetSecurityResponseType sets SecurityResponseType, Tag 323
func (m SecurityDefinition) SetSecurityResponseType(v int) {
	m.Set(field.NewSecurityResponseType(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m SecurityDefinition) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m SecurityDefinition) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m SecurityDefinition) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m SecurityDefinition) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m SecurityDefinition) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m SecurityDefinition) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m SecurityDefinition) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetTotalNumSecurities sets TotalNumSecurities, Tag 393
func (m SecurityDefinition) SetTotalNumSecurities(v int) {
	m.Set(field.NewTotalNumSecurities(v))
}

//GetCurrency gets Currency, Tag 15
func (m SecurityDefinition) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m SecurityDefinition) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m SecurityDefinition) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m SecurityDefinition) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m SecurityDefinition) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m SecurityDefinition) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m SecurityDefinition) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m SecurityDefinition) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m SecurityDefinition) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSecurityType gets SecurityType, Tag 167
func (m SecurityDefinition) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m SecurityDefinition) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m SecurityDefinition) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m SecurityDefinition) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m SecurityDefinition) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m SecurityDefinition) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m SecurityDefinition) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m SecurityDefinition) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m SecurityDefinition) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityReqID gets SecurityReqID, Tag 320
func (m SecurityDefinition) GetSecurityReqID() (f field.SecurityReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityResponseID gets SecurityResponseID, Tag 322
func (m SecurityDefinition) GetSecurityResponseID() (f field.SecurityResponseIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityResponseType gets SecurityResponseType, Tag 323
func (m SecurityDefinition) GetSecurityResponseType() (f field.SecurityResponseTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m SecurityDefinition) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m SecurityDefinition) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m SecurityDefinition) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m SecurityDefinition) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m SecurityDefinition) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m SecurityDefinition) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m SecurityDefinition) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalNumSecurities gets TotalNumSecurities, Tag 393
func (m SecurityDefinition) GetTotalNumSecurities() (f field.TotalNumSecuritiesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m SecurityDefinition) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m SecurityDefinition) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m SecurityDefinition) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m SecurityDefinition) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m SecurityDefinition) HasText() bool {
	return m.Has(tag.Text)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m SecurityDefinition) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m SecurityDefinition) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m SecurityDefinition) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m SecurityDefinition) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m SecurityDefinition) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m SecurityDefinition) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m SecurityDefinition) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m SecurityDefinition) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m SecurityDefinition) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m SecurityDefinition) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m SecurityDefinition) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m SecurityDefinition) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m SecurityDefinition) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasSecurityReqID returns true if SecurityReqID is present, Tag 320
func (m SecurityDefinition) HasSecurityReqID() bool {
	return m.Has(tag.SecurityReqID)
}

//HasSecurityResponseID returns true if SecurityResponseID is present, Tag 322
func (m SecurityDefinition) HasSecurityResponseID() bool {
	return m.Has(tag.SecurityResponseID)
}

//HasSecurityResponseType returns true if SecurityResponseType is present, Tag 323
func (m SecurityDefinition) HasSecurityResponseType() bool {
	return m.Has(tag.SecurityResponseType)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m SecurityDefinition) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m SecurityDefinition) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m SecurityDefinition) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m SecurityDefinition) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m SecurityDefinition) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m SecurityDefinition) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m SecurityDefinition) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasTotalNumSecurities returns true if TotalNumSecurities is present, Tag 393
func (m SecurityDefinition) HasTotalNumSecurities() bool {
	return m.Has(tag.TotalNumSecurities)
}

//NoRelatedSym is a repeating group element, Tag 146
type NoRelatedSym struct {
	quickfix.Group
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoRelatedSym) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoRelatedSym) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoRelatedSym) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingIDSource sets UnderlyingIDSource, Tag 305
func (m NoRelatedSym) SetUnderlyingIDSource(v string) {
	m.Set(field.NewUnderlyingIDSource(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoRelatedSym) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoRelatedSym) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDay sets UnderlyingMaturityDay, Tag 314
func (m NoRelatedSym) SetUnderlyingMaturityDay(v int) {
	m.Set(field.NewUnderlyingMaturityDay(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoRelatedSym) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoRelatedSym) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoRelatedSym) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoRelatedSym) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoRelatedSym) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoRelatedSym) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoRelatedSym) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoRelatedSym) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoRelatedSym) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoRelatedSym) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoRelatedSym) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoRelatedSym) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetRatioQty sets RatioQty, Tag 319
func (m NoRelatedSym) SetRatioQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewRatioQty(value, scale))
}

//SetSide sets Side, Tag 54
func (m NoRelatedSym) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetUnderlyingCurrency sets UnderlyingCurrency, Tag 318
func (m NoRelatedSym) SetUnderlyingCurrency(v string) {
	m.Set(field.NewUnderlyingCurrency(v))
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoRelatedSym) GetUnderlyingSymbol() (f field.UnderlyingSymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoRelatedSym) GetUnderlyingSymbolSfx() (f field.UnderlyingSymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoRelatedSym) GetUnderlyingSecurityID() (f field.UnderlyingSecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIDSource gets UnderlyingIDSource, Tag 305
func (m NoRelatedSym) GetUnderlyingIDSource() (f field.UnderlyingIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoRelatedSym) GetUnderlyingSecurityType() (f field.UnderlyingSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoRelatedSym) GetUnderlyingMaturityMonthYear() (f field.UnderlyingMaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingMaturityDay gets UnderlyingMaturityDay, Tag 314
func (m NoRelatedSym) GetUnderlyingMaturityDay() (f field.UnderlyingMaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoRelatedSym) GetUnderlyingPutOrCall() (f field.UnderlyingPutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoRelatedSym) GetUnderlyingStrikePrice() (f field.UnderlyingStrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoRelatedSym) GetUnderlyingOptAttribute() (f field.UnderlyingOptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoRelatedSym) GetUnderlyingContractMultiplier() (f field.UnderlyingContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoRelatedSym) GetUnderlyingCouponRate() (f field.UnderlyingCouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoRelatedSym) GetUnderlyingSecurityExchange() (f field.UnderlyingSecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoRelatedSym) GetUnderlyingIssuer() (f field.UnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoRelatedSym) GetEncodedUnderlyingIssuerLen() (f field.EncodedUnderlyingIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoRelatedSym) GetEncodedUnderlyingIssuer() (f field.EncodedUnderlyingIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoRelatedSym) GetUnderlyingSecurityDesc() (f field.UnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoRelatedSym) GetEncodedUnderlyingSecurityDescLen() (f field.EncodedUnderlyingSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoRelatedSym) GetEncodedUnderlyingSecurityDesc() (f field.EncodedUnderlyingSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRatioQty gets RatioQty, Tag 319
func (m NoRelatedSym) GetRatioQty() (f field.RatioQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NoRelatedSym) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnderlyingCurrency gets UnderlyingCurrency, Tag 318
func (m NoRelatedSym) GetUnderlyingCurrency() (f field.UnderlyingCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoRelatedSym) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoRelatedSym) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoRelatedSym) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingIDSource returns true if UnderlyingIDSource is present, Tag 305
func (m NoRelatedSym) HasUnderlyingIDSource() bool {
	return m.Has(tag.UnderlyingIDSource)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoRelatedSym) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoRelatedSym) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDay returns true if UnderlyingMaturityDay is present, Tag 314
func (m NoRelatedSym) HasUnderlyingMaturityDay() bool {
	return m.Has(tag.UnderlyingMaturityDay)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoRelatedSym) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoRelatedSym) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoRelatedSym) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoRelatedSym) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoRelatedSym) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoRelatedSym) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoRelatedSym) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoRelatedSym) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoRelatedSym) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoRelatedSym) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoRelatedSym) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoRelatedSym) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasRatioQty returns true if RatioQty is present, Tag 319
func (m NoRelatedSym) HasRatioQty() bool {
	return m.Has(tag.RatioQty)
}

//HasSide returns true if Side is present, Tag 54
func (m NoRelatedSym) HasSide() bool {
	return m.Has(tag.Side)
}

//HasUnderlyingCurrency returns true if UnderlyingCurrency is present, Tag 318
func (m NoRelatedSym) HasUnderlyingCurrency() bool {
	return m.Has(tag.UnderlyingCurrency)
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingIDSource), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDay), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.RatioQty), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.UnderlyingCurrency)})}
}

//Add create and append a new NoRelatedSym to this group
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

//Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}
