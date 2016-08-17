package liststrikeprice

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//ListStrikePrice is the fix42 ListStrikePrice type, MsgType = m
type ListStrikePrice struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListStrikePrice from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListStrikePrice {
	return ListStrikePrice{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListStrikePrice) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListStrikePrice initialized with the required fields for ListStrikePrice
func New(listid field.ListIDField, totnostrikes field.TotNoStrikesField) (m ListStrikePrice) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("m"))
	m.Set(listid)
	m.Set(totnostrikes)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListStrikePrice, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "m", r
}

//SetListID sets ListID, Tag 66
func (m ListStrikePrice) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetTotNoStrikes sets TotNoStrikes, Tag 422
func (m ListStrikePrice) SetTotNoStrikes(v int) {
	m.Set(field.NewTotNoStrikes(v))
}

//SetNoStrikes sets NoStrikes, Tag 428
func (m ListStrikePrice) SetNoStrikes(f NoStrikesRepeatingGroup) {
	m.SetGroup(f)
}

//GetListID gets ListID, Tag 66
func (m ListStrikePrice) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotNoStrikes gets TotNoStrikes, Tag 422
func (m ListStrikePrice) GetTotNoStrikes() (f field.TotNoStrikesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoStrikes gets NoStrikes, Tag 428
func (m ListStrikePrice) GetNoStrikes() (NoStrikesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoStrikesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasListID returns true if ListID is present, Tag 66
func (m ListStrikePrice) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasTotNoStrikes returns true if TotNoStrikes is present, Tag 422
func (m ListStrikePrice) HasTotNoStrikes() bool {
	return m.Has(tag.TotNoStrikes)
}

//HasNoStrikes returns true if NoStrikes is present, Tag 428
func (m ListStrikePrice) HasNoStrikes() bool {
	return m.Has(tag.NoStrikes)
}

//NoStrikes is a repeating group element, Tag 428
type NoStrikes struct {
	quickfix.Group
}

//SetSymbol sets Symbol, Tag 55
func (m NoStrikes) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoStrikes) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoStrikes) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NoStrikes) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoStrikes) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoStrikes) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NoStrikes) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoStrikes) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoStrikes) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoStrikes) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoStrikes) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoStrikes) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoStrikes) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoStrikes) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoStrikes) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoStrikes) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoStrikes) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoStrikes) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoStrikes) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetPrevClosePx sets PrevClosePx, Tag 140
func (m NoStrikes) SetPrevClosePx(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrevClosePx(value, scale))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoStrikes) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSide sets Side, Tag 54
func (m NoStrikes) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetPrice sets Price, Tag 44
func (m NoStrikes) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m NoStrikes) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetText sets Text, Tag 58
func (m NoStrikes) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoStrikes) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoStrikes) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetSymbol gets Symbol, Tag 55
func (m NoStrikes) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoStrikes) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoStrikes) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NoStrikes) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoStrikes) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoStrikes) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NoStrikes) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoStrikes) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoStrikes) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoStrikes) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoStrikes) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoStrikes) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoStrikes) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoStrikes) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoStrikes) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoStrikes) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoStrikes) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoStrikes) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoStrikes) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrevClosePx gets PrevClosePx, Tag 140
func (m NoStrikes) GetPrevClosePx() (f field.PrevClosePxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoStrikes) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NoStrikes) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m NoStrikes) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoStrikes) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoStrikes) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoStrikes) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoStrikes) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoStrikes) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoStrikes) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoStrikes) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NoStrikes) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoStrikes) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoStrikes) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NoStrikes) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoStrikes) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoStrikes) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoStrikes) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoStrikes) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoStrikes) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoStrikes) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoStrikes) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoStrikes) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoStrikes) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoStrikes) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoStrikes) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoStrikes) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasPrevClosePx returns true if PrevClosePx is present, Tag 140
func (m NoStrikes) HasPrevClosePx() bool {
	return m.Has(tag.PrevClosePx)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoStrikes) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSide returns true if Side is present, Tag 54
func (m NoStrikes) HasSide() bool {
	return m.Has(tag.Side)
}

//HasPrice returns true if Price is present, Tag 44
func (m NoStrikes) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoStrikes) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasText returns true if Text is present, Tag 58
func (m NoStrikes) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoStrikes) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoStrikes) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//NoStrikesRepeatingGroup is a repeating group, Tag 428
type NoStrikesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoStrikesRepeatingGroup returns an initialized, NoStrikesRepeatingGroup
func NewNoStrikesRepeatingGroup() NoStrikesRepeatingGroup {
	return NoStrikesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoStrikes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.PrevClosePx), quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
}

//Add create and append a new NoStrikes to this group
func (m NoStrikesRepeatingGroup) Add() NoStrikes {
	g := m.RepeatingGroup.Add()
	return NoStrikes{g}
}

//Get returns the ith NoStrikes in the NoStrikesRepeatinGroup
func (m NoStrikesRepeatingGroup) Get(i int) NoStrikes {
	return NoStrikes{m.RepeatingGroup.Get(i)}
}
