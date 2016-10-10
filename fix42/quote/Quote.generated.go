package quote

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//Quote is the fix42 Quote type, MsgType = S
type Quote struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Quote from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Quote {
	return Quote{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Quote) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Quote initialized with the required fields for Quote
func New(quoteid field.QuoteIDField, symbol field.SymbolField) (m Quote) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("S"))
	m.Set(quoteid)
	m.Set(symbol)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Quote, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "S", r
}

//SetCurrency sets Currency, Tag 15
func (m Quote) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m Quote) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetOrdType sets OrdType, Tag 40
func (m Quote) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Quote) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Quote) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Quote) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m Quote) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Quote) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Quote) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m Quote) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Quote) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetQuoteID sets QuoteID, Tag 117
func (m Quote) SetQuoteID(v string) {
	m.Set(field.NewQuoteID(v))
}

//SetQuoteReqID sets QuoteReqID, Tag 131
func (m Quote) SetQuoteReqID(v string) {
	m.Set(field.NewQuoteReqID(v))
}

//SetBidPx sets BidPx, Tag 132
func (m Quote) SetBidPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidPx(value, scale))
}

//SetOfferPx sets OfferPx, Tag 133
func (m Quote) SetOfferPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferPx(value, scale))
}

//SetBidSize sets BidSize, Tag 134
func (m Quote) SetBidSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSize(value, scale))
}

//SetOfferSize sets OfferSize, Tag 135
func (m Quote) SetOfferSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSize(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m Quote) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetBidSpotRate sets BidSpotRate, Tag 188
func (m Quote) SetBidSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidSpotRate(value, scale))
}

//SetBidForwardPoints sets BidForwardPoints, Tag 189
func (m Quote) SetBidForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewBidForwardPoints(value, scale))
}

//SetOfferSpotRate sets OfferSpotRate, Tag 190
func (m Quote) SetOfferSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferSpotRate(value, scale))
}

//SetOfferForwardPoints sets OfferForwardPoints, Tag 191
func (m Quote) SetOfferForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewOfferForwardPoints(value, scale))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m Quote) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m Quote) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m Quote) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m Quote) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m Quote) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m Quote) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m Quote) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m Quote) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m Quote) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m Quote) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetQuoteResponseLevel sets QuoteResponseLevel, Tag 301
func (m Quote) SetQuoteResponseLevel(v enum.QuoteResponseLevel) {
	m.Set(field.NewQuoteResponseLevel(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m Quote) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m Quote) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m Quote) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m Quote) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m Quote) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//GetCurrency gets Currency, Tag 15
func (m Quote) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m Quote) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrdType gets OrdType, Tag 40
func (m Quote) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Quote) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Quote) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Quote) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m Quote) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Quote) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Quote) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m Quote) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Quote) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteID gets QuoteID, Tag 117
func (m Quote) GetQuoteID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteReqID gets QuoteReqID, Tag 131
func (m Quote) GetQuoteReqID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidPx gets BidPx, Tag 132
func (m Quote) GetBidPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferPx gets OfferPx, Tag 133
func (m Quote) GetOfferPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSize gets BidSize, Tag 134
func (m Quote) GetBidSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSize gets OfferSize, Tag 135
func (m Quote) GetOfferSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m Quote) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidSpotRate gets BidSpotRate, Tag 188
func (m Quote) GetBidSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidForwardPoints gets BidForwardPoints, Tag 189
func (m Quote) GetBidForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.BidForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferSpotRate gets OfferSpotRate, Tag 190
func (m Quote) GetOfferSpotRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferSpotRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOfferForwardPoints gets OfferForwardPoints, Tag 191
func (m Quote) GetOfferForwardPoints() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OfferForwardPointsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m Quote) GetOrderQty2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OrderQty2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m Quote) GetFutSettDate2() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDate2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m Quote) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m Quote) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m Quote) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m Quote) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m Quote) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m Quote) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m Quote) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m Quote) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteResponseLevel gets QuoteResponseLevel, Tag 301
func (m Quote) GetQuoteResponseLevel() (v enum.QuoteResponseLevel, err quickfix.MessageRejectError) {
	var f field.QuoteResponseLevelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m Quote) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m Quote) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m Quote) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m Quote) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m Quote) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m Quote) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m Quote) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m Quote) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Quote) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Quote) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Quote) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m Quote) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m Quote) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Quote) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Quote) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Quote) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasQuoteID returns true if QuoteID is present, Tag 117
func (m Quote) HasQuoteID() bool {
	return m.Has(tag.QuoteID)
}

//HasQuoteReqID returns true if QuoteReqID is present, Tag 131
func (m Quote) HasQuoteReqID() bool {
	return m.Has(tag.QuoteReqID)
}

//HasBidPx returns true if BidPx is present, Tag 132
func (m Quote) HasBidPx() bool {
	return m.Has(tag.BidPx)
}

//HasOfferPx returns true if OfferPx is present, Tag 133
func (m Quote) HasOfferPx() bool {
	return m.Has(tag.OfferPx)
}

//HasBidSize returns true if BidSize is present, Tag 134
func (m Quote) HasBidSize() bool {
	return m.Has(tag.BidSize)
}

//HasOfferSize returns true if OfferSize is present, Tag 135
func (m Quote) HasOfferSize() bool {
	return m.Has(tag.OfferSize)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m Quote) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasBidSpotRate returns true if BidSpotRate is present, Tag 188
func (m Quote) HasBidSpotRate() bool {
	return m.Has(tag.BidSpotRate)
}

//HasBidForwardPoints returns true if BidForwardPoints is present, Tag 189
func (m Quote) HasBidForwardPoints() bool {
	return m.Has(tag.BidForwardPoints)
}

//HasOfferSpotRate returns true if OfferSpotRate is present, Tag 190
func (m Quote) HasOfferSpotRate() bool {
	return m.Has(tag.OfferSpotRate)
}

//HasOfferForwardPoints returns true if OfferForwardPoints is present, Tag 191
func (m Quote) HasOfferForwardPoints() bool {
	return m.Has(tag.OfferForwardPoints)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m Quote) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m Quote) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m Quote) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m Quote) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m Quote) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m Quote) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m Quote) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m Quote) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m Quote) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m Quote) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasQuoteResponseLevel returns true if QuoteResponseLevel is present, Tag 301
func (m Quote) HasQuoteResponseLevel() bool {
	return m.Has(tag.QuoteResponseLevel)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m Quote) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m Quote) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m Quote) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m Quote) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m Quote) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}
