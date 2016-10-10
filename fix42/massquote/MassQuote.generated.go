package massquote

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//MassQuote is the fix42 MassQuote type, MsgType = i
type MassQuote struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a MassQuote from a quickfix.Message instance
func FromMessage(m *quickfix.Message) MassQuote {
	return MassQuote{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MassQuote) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MassQuote initialized with the required fields for MassQuote
func New(quoteid field.QuoteIDField) (m MassQuote) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("i"))
	m.Set(quoteid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MassQuote, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "i", r
}

//SetQuoteID sets QuoteID, Tag 117
func (m MassQuote) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m MassQuote) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetDefBidSize sets DefBidSize, Tag 293
func (m MassQuote) SetDefBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewDefBidSize(value, scale))
}

//SetDefOfferSize sets DefOfferSize, Tag 294
func (m MassQuote) SetDefOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewDefOfferSize(value, scale))
}

//SetNoQuoteSets sets NoQuoteSets, Tag 296
func (m MassQuote) SetNoQuoteSets(f NoQuoteSetsRepeatingGroup) {
	m.SetGroup(f)
}

//SetQuoteResponseLevel sets QuoteResponseLevel, Tag 301
func (m MassQuote) SetQuoteResponseLevel(v enum.QuoteResponseLevel) {
	m.Set(field.NewQuoteResponseLevel(v))
}

//GetQuoteID gets QuoteID, Tag 117
func (m MassQuote) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m MassQuote) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDefBidSize gets DefBidSize, Tag 293
func (m MassQuote) GetDefBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DefBidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDefOfferSize gets DefOfferSize, Tag 294
func (m MassQuote) GetDefOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.DefOfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoQuoteSets gets NoQuoteSets, Tag 296
func (m MassQuote) GetNoQuoteSets() (NoQuoteSetsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteSetsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetQuoteResponseLevel gets QuoteResponseLevel, Tag 301
func (m MassQuote) GetQuoteResponseLevel() (v enum.QuoteResponseLevel, err quickfix.MessageRejectError) {
	var f field.QuoteResponseLevelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m MassQuote) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m MassQuote) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasDefBidSize returns true if DefBidSize is present, Tag 293
func (m MassQuote) HasDefBidSize() bool {
	return m.Has(tag.DefBidSize)
}

//HasDefOfferSize returns true if DefOfferSize is present, Tag 294
func (m MassQuote) HasDefOfferSize() bool {
	return m.Has(tag.DefOfferSize)
}

//HasNoQuoteSets returns true if NoQuoteSets is present, Tag 296
func (m MassQuote) HasNoQuoteSets() bool {
	return m.Has(tag.NoQuoteSets)
}

//HasQuoteResponseLevel returns true if QuoteResponseLevel is present, Tag 301
func (m MassQuote) HasQuoteResponseLevel() bool {
	return m.Has(tag.QuoteResponseLevel)
}

//NoQuoteSets is a repeating group element, Tag 296
type NoQuoteSets struct {
	*quickfix.Group
}

//SetQuoteSetID sets QuoteSetID, Tag 302
func (m NoQuoteSets) SetQuoteSetID(v string) {
	m.Set(field.NewQuoteSetID(v))
}

//SetUnderlyingSymbol sets UnderlyingSymbol, Tag 311
func (m NoQuoteSets) SetUnderlyingSymbol(v string) {
	m.Set(field.NewUnderlyingSymbol(v))
}

//SetUnderlyingSymbolSfx sets UnderlyingSymbolSfx, Tag 312
func (m NoQuoteSets) SetUnderlyingSymbolSfx(v string) {
	m.Set(field.NewUnderlyingSymbolSfx(v))
}

//SetUnderlyingSecurityID sets UnderlyingSecurityID, Tag 309
func (m NoQuoteSets) SetUnderlyingSecurityID(v string) {
	m.Set(field.NewUnderlyingSecurityID(v))
}

//SetUnderlyingIDSource sets UnderlyingIDSource, Tag 305
func (m NoQuoteSets) SetUnderlyingIDSource(v string) {
	m.Set(field.NewUnderlyingIDSource(v))
}

//SetUnderlyingSecurityType sets UnderlyingSecurityType, Tag 310
func (m NoQuoteSets) SetUnderlyingSecurityType(v string) {
	m.Set(field.NewUnderlyingSecurityType(v))
}

//SetUnderlyingMaturityMonthYear sets UnderlyingMaturityMonthYear, Tag 313
func (m NoQuoteSets) SetUnderlyingMaturityMonthYear(v string) {
	m.Set(field.NewUnderlyingMaturityMonthYear(v))
}

//SetUnderlyingMaturityDay sets UnderlyingMaturityDay, Tag 314
func (m NoQuoteSets) SetUnderlyingMaturityDay(v int) {
	m.Set(field.NewUnderlyingMaturityDay(v))
}

//SetUnderlyingPutOrCall sets UnderlyingPutOrCall, Tag 315
func (m NoQuoteSets) SetUnderlyingPutOrCall(v int) {
	m.Set(field.NewUnderlyingPutOrCall(v))
}

//SetUnderlyingStrikePrice sets UnderlyingStrikePrice, Tag 316
func (m NoQuoteSets) SetUnderlyingStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingStrikePrice(value, scale))
}

//SetUnderlyingOptAttribute sets UnderlyingOptAttribute, Tag 317
func (m NoQuoteSets) SetUnderlyingOptAttribute(v string) {
	m.Set(field.NewUnderlyingOptAttribute(v))
}

//SetUnderlyingContractMultiplier sets UnderlyingContractMultiplier, Tag 436
func (m NoQuoteSets) SetUnderlyingContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingContractMultiplier(value, scale))
}

//SetUnderlyingCouponRate sets UnderlyingCouponRate, Tag 435
func (m NoQuoteSets) SetUnderlyingCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewUnderlyingCouponRate(value, scale))
}

//SetUnderlyingSecurityExchange sets UnderlyingSecurityExchange, Tag 308
func (m NoQuoteSets) SetUnderlyingSecurityExchange(v string) {
	m.Set(field.NewUnderlyingSecurityExchange(v))
}

//SetUnderlyingIssuer sets UnderlyingIssuer, Tag 306
func (m NoQuoteSets) SetUnderlyingIssuer(v string) {
	m.Set(field.NewUnderlyingIssuer(v))
}

//SetEncodedUnderlyingIssuerLen sets EncodedUnderlyingIssuerLen, Tag 362
func (m NoQuoteSets) SetEncodedUnderlyingIssuerLen(v int) {
	m.Set(field.NewEncodedUnderlyingIssuerLen(v))
}

//SetEncodedUnderlyingIssuer sets EncodedUnderlyingIssuer, Tag 363
func (m NoQuoteSets) SetEncodedUnderlyingIssuer(v string) {
	m.Set(field.NewEncodedUnderlyingIssuer(v))
}

//SetUnderlyingSecurityDesc sets UnderlyingSecurityDesc, Tag 307
func (m NoQuoteSets) SetUnderlyingSecurityDesc(v string) {
	m.Set(field.NewUnderlyingSecurityDesc(v))
}

//SetEncodedUnderlyingSecurityDescLen sets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoQuoteSets) SetEncodedUnderlyingSecurityDescLen(v int) {
	m.Set(field.NewEncodedUnderlyingSecurityDescLen(v))
}

//SetEncodedUnderlyingSecurityDesc sets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoQuoteSets) SetEncodedUnderlyingSecurityDesc(v string) {
	m.Set(field.NewEncodedUnderlyingSecurityDesc(v))
}

//SetQuoteSetValidUntilTime sets QuoteSetValidUntilTime, Tag 367
func (m NoQuoteSets) SetQuoteSetValidUntilTime(v time.Time) {
	m.Set(field.NewQuoteSetValidUntilTime(v))
}

//SetTotQuoteEntries sets TotQuoteEntries, Tag 304
func (m NoQuoteSets) SetTotQuoteEntries(v int) {
	m.Set(field.NewTotQuoteEntries(v))
}

//SetNoQuoteEntries sets NoQuoteEntries, Tag 295
func (m NoQuoteSets) SetNoQuoteEntries(f NoQuoteEntriesRepeatingGroup) {
	m.SetGroup(f)
}

//GetQuoteSetID gets QuoteSetID, Tag 302
func (m NoQuoteSets) GetQuoteSetID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteSetIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbol gets UnderlyingSymbol, Tag 311
func (m NoQuoteSets) GetUnderlyingSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSymbolSfx gets UnderlyingSymbolSfx, Tag 312
func (m NoQuoteSets) GetUnderlyingSymbolSfx() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityID gets UnderlyingSecurityID, Tag 309
func (m NoQuoteSets) GetUnderlyingSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIDSource gets UnderlyingIDSource, Tag 305
func (m NoQuoteSets) GetUnderlyingIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityType gets UnderlyingSecurityType, Tag 310
func (m NoQuoteSets) GetUnderlyingSecurityType() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityMonthYear gets UnderlyingMaturityMonthYear, Tag 313
func (m NoQuoteSets) GetUnderlyingMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingMaturityDay gets UnderlyingMaturityDay, Tag 314
func (m NoQuoteSets) GetUnderlyingMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingMaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingPutOrCall gets UnderlyingPutOrCall, Tag 315
func (m NoQuoteSets) GetUnderlyingPutOrCall() (v int, err quickfix.MessageRejectError) {
	var f field.UnderlyingPutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingStrikePrice gets UnderlyingStrikePrice, Tag 316
func (m NoQuoteSets) GetUnderlyingStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingStrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingOptAttribute gets UnderlyingOptAttribute, Tag 317
func (m NoQuoteSets) GetUnderlyingOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingOptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingContractMultiplier gets UnderlyingContractMultiplier, Tag 436
func (m NoQuoteSets) GetUnderlyingContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingCouponRate gets UnderlyingCouponRate, Tag 435
func (m NoQuoteSets) GetUnderlyingCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.UnderlyingCouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityExchange gets UnderlyingSecurityExchange, Tag 308
func (m NoQuoteSets) GetUnderlyingSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingIssuer gets UnderlyingIssuer, Tag 306
func (m NoQuoteSets) GetUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuerLen gets EncodedUnderlyingIssuerLen, Tag 362
func (m NoQuoteSets) GetEncodedUnderlyingIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingIssuer gets EncodedUnderlyingIssuer, Tag 363
func (m NoQuoteSets) GetEncodedUnderlyingIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnderlyingSecurityDesc gets UnderlyingSecurityDesc, Tag 307
func (m NoQuoteSets) GetUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.UnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDescLen gets EncodedUnderlyingSecurityDescLen, Tag 364
func (m NoQuoteSets) GetEncodedUnderlyingSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedUnderlyingSecurityDesc gets EncodedUnderlyingSecurityDesc, Tag 365
func (m NoQuoteSets) GetEncodedUnderlyingSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedUnderlyingSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteSetValidUntilTime gets QuoteSetValidUntilTime, Tag 367
func (m NoQuoteSets) GetQuoteSetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.QuoteSetValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotQuoteEntries gets TotQuoteEntries, Tag 304
func (m NoQuoteSets) GetTotQuoteEntries() (v int, err quickfix.MessageRejectError) {
	var f field.TotQuoteEntriesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoQuoteEntries gets NoQuoteEntries, Tag 295
func (m NoQuoteSets) GetNoQuoteEntries() (NoQuoteEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoQuoteEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasQuoteSetID returns true if QuoteSetID is present, Tag 302
func (m NoQuoteSets) HasQuoteSetID() bool {
	return m.Has(tag.QuoteSetID)
}

//HasUnderlyingSymbol returns true if UnderlyingSymbol is present, Tag 311
func (m NoQuoteSets) HasUnderlyingSymbol() bool {
	return m.Has(tag.UnderlyingSymbol)
}

//HasUnderlyingSymbolSfx returns true if UnderlyingSymbolSfx is present, Tag 312
func (m NoQuoteSets) HasUnderlyingSymbolSfx() bool {
	return m.Has(tag.UnderlyingSymbolSfx)
}

//HasUnderlyingSecurityID returns true if UnderlyingSecurityID is present, Tag 309
func (m NoQuoteSets) HasUnderlyingSecurityID() bool {
	return m.Has(tag.UnderlyingSecurityID)
}

//HasUnderlyingIDSource returns true if UnderlyingIDSource is present, Tag 305
func (m NoQuoteSets) HasUnderlyingIDSource() bool {
	return m.Has(tag.UnderlyingIDSource)
}

//HasUnderlyingSecurityType returns true if UnderlyingSecurityType is present, Tag 310
func (m NoQuoteSets) HasUnderlyingSecurityType() bool {
	return m.Has(tag.UnderlyingSecurityType)
}

//HasUnderlyingMaturityMonthYear returns true if UnderlyingMaturityMonthYear is present, Tag 313
func (m NoQuoteSets) HasUnderlyingMaturityMonthYear() bool {
	return m.Has(tag.UnderlyingMaturityMonthYear)
}

//HasUnderlyingMaturityDay returns true if UnderlyingMaturityDay is present, Tag 314
func (m NoQuoteSets) HasUnderlyingMaturityDay() bool {
	return m.Has(tag.UnderlyingMaturityDay)
}

//HasUnderlyingPutOrCall returns true if UnderlyingPutOrCall is present, Tag 315
func (m NoQuoteSets) HasUnderlyingPutOrCall() bool {
	return m.Has(tag.UnderlyingPutOrCall)
}

//HasUnderlyingStrikePrice returns true if UnderlyingStrikePrice is present, Tag 316
func (m NoQuoteSets) HasUnderlyingStrikePrice() bool {
	return m.Has(tag.UnderlyingStrikePrice)
}

//HasUnderlyingOptAttribute returns true if UnderlyingOptAttribute is present, Tag 317
func (m NoQuoteSets) HasUnderlyingOptAttribute() bool {
	return m.Has(tag.UnderlyingOptAttribute)
}

//HasUnderlyingContractMultiplier returns true if UnderlyingContractMultiplier is present, Tag 436
func (m NoQuoteSets) HasUnderlyingContractMultiplier() bool {
	return m.Has(tag.UnderlyingContractMultiplier)
}

//HasUnderlyingCouponRate returns true if UnderlyingCouponRate is present, Tag 435
func (m NoQuoteSets) HasUnderlyingCouponRate() bool {
	return m.Has(tag.UnderlyingCouponRate)
}

//HasUnderlyingSecurityExchange returns true if UnderlyingSecurityExchange is present, Tag 308
func (m NoQuoteSets) HasUnderlyingSecurityExchange() bool {
	return m.Has(tag.UnderlyingSecurityExchange)
}

//HasUnderlyingIssuer returns true if UnderlyingIssuer is present, Tag 306
func (m NoQuoteSets) HasUnderlyingIssuer() bool {
	return m.Has(tag.UnderlyingIssuer)
}

//HasEncodedUnderlyingIssuerLen returns true if EncodedUnderlyingIssuerLen is present, Tag 362
func (m NoQuoteSets) HasEncodedUnderlyingIssuerLen() bool {
	return m.Has(tag.EncodedUnderlyingIssuerLen)
}

//HasEncodedUnderlyingIssuer returns true if EncodedUnderlyingIssuer is present, Tag 363
func (m NoQuoteSets) HasEncodedUnderlyingIssuer() bool {
	return m.Has(tag.EncodedUnderlyingIssuer)
}

//HasUnderlyingSecurityDesc returns true if UnderlyingSecurityDesc is present, Tag 307
func (m NoQuoteSets) HasUnderlyingSecurityDesc() bool {
	return m.Has(tag.UnderlyingSecurityDesc)
}

//HasEncodedUnderlyingSecurityDescLen returns true if EncodedUnderlyingSecurityDescLen is present, Tag 364
func (m NoQuoteSets) HasEncodedUnderlyingSecurityDescLen() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDescLen)
}

//HasEncodedUnderlyingSecurityDesc returns true if EncodedUnderlyingSecurityDesc is present, Tag 365
func (m NoQuoteSets) HasEncodedUnderlyingSecurityDesc() bool {
	return m.Has(tag.EncodedUnderlyingSecurityDesc)
}

//HasQuoteSetValidUntilTime returns true if QuoteSetValidUntilTime is present, Tag 367
func (m NoQuoteSets) HasQuoteSetValidUntilTime() bool {
	return m.Has(tag.QuoteSetValidUntilTime)
}

//HasTotQuoteEntries returns true if TotQuoteEntries is present, Tag 304
func (m NoQuoteSets) HasTotQuoteEntries() bool {
	return m.Has(tag.TotQuoteEntries)
}

//HasNoQuoteEntries returns true if NoQuoteEntries is present, Tag 295
func (m NoQuoteSets) HasNoQuoteEntries() bool {
	return m.Has(tag.NoQuoteEntries)
}

//NoQuoteEntries is a repeating group element, Tag 295
type NoQuoteEntries struct {
	*quickfix.Group
}

//SetQuoteEntryID sets QuoteEntryID, Tag 299
func (m NoQuoteEntries) SetQuoteEntryID(v string) {
	m.Set(field.NewQuoteEntryID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NoQuoteEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoQuoteEntries) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoQuoteEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NoQuoteEntries) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoQuoteEntries) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoQuoteEntries) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NoQuoteEntries) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoQuoteEntries) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoQuoteEntries) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoQuoteEntries) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoQuoteEntries) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoQuoteEntries) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoQuoteEntries) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoQuoteEntries) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoQuoteEntries) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoQuoteEntries) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoQuoteEntries) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoQuoteEntries) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoQuoteEntries) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetBidPx sets BidPx, Tag 132
func (m NoQuoteEntries) SetBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidPx(value, scale))
}

//SetOfferPx sets OfferPx, Tag 133
func (m NoQuoteEntries) SetOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferPx(value, scale))
}

//SetBidSize sets BidSize, Tag 134
func (m NoQuoteEntries) SetBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSize(value, scale))
}

//SetOfferSize sets OfferSize, Tag 135
func (m NoQuoteEntries) SetOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSize(value, scale))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m NoQuoteEntries) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetBidSpotRate sets BidSpotRate, Tag 188
func (m NoQuoteEntries) SetBidSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSpotRate(value, scale))
}

//SetOfferSpotRate sets OfferSpotRate, Tag 190
func (m NoQuoteEntries) SetOfferSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSpotRate(value, scale))
}

//SetBidForwardPoints sets BidForwardPoints, Tag 189
func (m NoQuoteEntries) SetBidForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidForwardPoints(value, scale))
}

//SetOfferForwardPoints sets OfferForwardPoints, Tag 191
func (m NoQuoteEntries) SetOfferForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferForwardPoints(value, scale))
}

//SetTransactTime sets TransactTime, Tag 60
func (m NoQuoteEntries) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoQuoteEntries) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NoQuoteEntries) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetOrdType sets OrdType, Tag 40
func (m NoQuoteEntries) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m NoQuoteEntries) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m NoQuoteEntries) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m NoQuoteEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//GetQuoteEntryID gets QuoteEntryID, Tag 299
func (m NoQuoteEntries) GetQuoteEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NoQuoteEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoQuoteEntries) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoQuoteEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NoQuoteEntries) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoQuoteEntries) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoQuoteEntries) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NoQuoteEntries) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoQuoteEntries) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoQuoteEntries) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoQuoteEntries) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoQuoteEntries) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoQuoteEntries) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoQuoteEntries) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoQuoteEntries) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoQuoteEntries) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoQuoteEntries) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoQuoteEntries) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoQuoteEntries) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoQuoteEntries) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidPx gets BidPx, Tag 132
func (m NoQuoteEntries) GetBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferPx gets OfferPx, Tag 133
func (m NoQuoteEntries) GetOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSize gets BidSize, Tag 134
func (m NoQuoteEntries) GetBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSize gets OfferSize, Tag 135
func (m NoQuoteEntries) GetOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m NoQuoteEntries) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSpotRate gets BidSpotRate, Tag 188
func (m NoQuoteEntries) GetBidSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSpotRate gets OfferSpotRate, Tag 190
func (m NoQuoteEntries) GetOfferSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidForwardPoints gets BidForwardPoints, Tag 189
func (m NoQuoteEntries) GetBidForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferForwardPoints gets OfferForwardPoints, Tag 191
func (m NoQuoteEntries) GetOfferForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m NoQuoteEntries) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoQuoteEntries) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NoQuoteEntries) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m NoQuoteEntries) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m NoQuoteEntries) GetFutSettDate2() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m NoQuoteEntries) GetOrderQty2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoQuoteEntries) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasQuoteEntryID returns true if QuoteEntryID is present, Tag 299
func (m NoQuoteEntries) HasQuoteEntryID() bool {
	return m.Has(tag.QuoteEntryID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoQuoteEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoQuoteEntries) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoQuoteEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NoQuoteEntries) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoQuoteEntries) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoQuoteEntries) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NoQuoteEntries) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoQuoteEntries) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoQuoteEntries) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoQuoteEntries) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoQuoteEntries) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoQuoteEntries) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoQuoteEntries) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoQuoteEntries) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoQuoteEntries) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoQuoteEntries) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoQuoteEntries) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoQuoteEntries) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoQuoteEntries) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasBidPx returns true if BidPx is present, Tag 132
func (m NoQuoteEntries) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

//HasOfferPx returns true if OfferPx is present, Tag 133
func (m NoQuoteEntries) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

//HasBidSize returns true if BidSize is present, Tag 134
func (m NoQuoteEntries) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

//HasOfferSize returns true if OfferSize is present, Tag 135
func (m NoQuoteEntries) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m NoQuoteEntries) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasBidSpotRate returns true if BidSpotRate is present, Tag 188
func (m NoQuoteEntries) HasBidSpotRate() bool {
	return m.Has(tag.BidSpotRate)
}

//HasOfferSpotRate returns true if OfferSpotRate is present, Tag 190
func (m NoQuoteEntries) HasOfferSpotRate() bool {
	return m.Has(tag.OfferSpotRate)
}

//HasBidForwardPoints returns true if BidForwardPoints is present, Tag 189
func (m NoQuoteEntries) HasBidForwardPoints() bool {
	return m.Has(tag.BidForwardPoints)
}

//HasOfferForwardPoints returns true if OfferForwardPoints is present, Tag 191
func (m NoQuoteEntries) HasOfferForwardPoints() bool {
	return m.Has(tag.OfferForwardPoints)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m NoQuoteEntries) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoQuoteEntries) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NoQuoteEntries) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoQuoteEntries) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m NoQuoteEntries) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m NoQuoteEntries) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoQuoteEntries) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//NoQuoteEntriesRepeatingGroup is a repeating group, Tag 295
type NoQuoteEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoQuoteEntriesRepeatingGroup returns an initialized, NoQuoteEntriesRepeatingGroup
func NewNoQuoteEntriesRepeatingGroup() NoQuoteEntriesRepeatingGroup {
	return NoQuoteEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoQuoteEntries,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.QuoteEntryID), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.BidPx), quickfix.GroupElement(tag.OfferPx), quickfix.GroupElement(tag.BidSize), quickfix.GroupElement(tag.OfferSize), quickfix.GroupElement(tag.ValidUntilTime), quickfix.GroupElement(tag.BidSpotRate), quickfix.GroupElement(tag.OfferSpotRate), quickfix.GroupElement(tag.BidForwardPoints), quickfix.GroupElement(tag.OfferForwardPoints), quickfix.GroupElement(tag.TransactTime), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.FutSettDate), quickfix.GroupElement(tag.OrdType), quickfix.GroupElement(tag.FutSettDate2), quickfix.GroupElement(tag.OrderQty2), quickfix.GroupElement(tag.Currency)})}
}

//Add create and append a new NoQuoteEntries to this group
func (m NoQuoteEntriesRepeatingGroup) Add() NoQuoteEntries {
	g := m.RepeatingGroup.Add()
	return NoQuoteEntries{g}
}

//Get returns the ith NoQuoteEntries in the NoQuoteEntriesRepeatinGroup
func (m NoQuoteEntriesRepeatingGroup) Get(i int) NoQuoteEntries {
	return NoQuoteEntries{m.RepeatingGroup.Get(i)}
}

//NoQuoteSetsRepeatingGroup is a repeating group, Tag 296
type NoQuoteSetsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoQuoteSetsRepeatingGroup returns an initialized, NoQuoteSetsRepeatingGroup
func NewNoQuoteSetsRepeatingGroup() NoQuoteSetsRepeatingGroup {
	return NoQuoteSetsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoQuoteSets,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.QuoteSetID), quickfix.GroupElement(tag.UnderlyingSymbol), quickfix.GroupElement(tag.UnderlyingSymbolSfx), quickfix.GroupElement(tag.UnderlyingSecurityID), quickfix.GroupElement(tag.UnderlyingIDSource), quickfix.GroupElement(tag.UnderlyingSecurityType), quickfix.GroupElement(tag.UnderlyingMaturityMonthYear), quickfix.GroupElement(tag.UnderlyingMaturityDay), quickfix.GroupElement(tag.UnderlyingPutOrCall), quickfix.GroupElement(tag.UnderlyingStrikePrice), quickfix.GroupElement(tag.UnderlyingOptAttribute), quickfix.GroupElement(tag.UnderlyingContractMultiplier), quickfix.GroupElement(tag.UnderlyingCouponRate), quickfix.GroupElement(tag.UnderlyingSecurityExchange), quickfix.GroupElement(tag.UnderlyingIssuer), quickfix.GroupElement(tag.EncodedUnderlyingIssuerLen), quickfix.GroupElement(tag.EncodedUnderlyingIssuer), quickfix.GroupElement(tag.UnderlyingSecurityDesc), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDescLen), quickfix.GroupElement(tag.EncodedUnderlyingSecurityDesc), quickfix.GroupElement(tag.QuoteSetValidUntilTime), quickfix.GroupElement(tag.TotQuoteEntries), NewNoQuoteEntriesRepeatingGroup()})}
}

//Add create and append a new NoQuoteSets to this group
func (m NoQuoteSetsRepeatingGroup) Add() NoQuoteSets {
	g := m.RepeatingGroup.Add()
	return NoQuoteSets{g}
}

//Get returns the ith NoQuoteSets in the NoQuoteSetsRepeatinGroup
func (m NoQuoteSetsRepeatingGroup) Get(i int) NoQuoteSets {
	return NoQuoteSets{m.RepeatingGroup.Get(i)}
}
