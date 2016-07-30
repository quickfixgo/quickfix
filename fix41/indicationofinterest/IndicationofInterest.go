package indicationofinterest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//IndicationofInterest is the fix41 IndicationofInterest type, MsgType = 6
type IndicationofInterest struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a IndicationofInterest from a quickfix.Message instance
func FromMessage(m quickfix.Message) IndicationofInterest {
	return IndicationofInterest{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m IndicationofInterest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a IndicationofInterest initialized with the required fields for IndicationofInterest
func New(ioiid field.IOIidField, ioitranstype field.IOITransTypeField, symbol field.SymbolField, side field.SideField, ioishares field.IOISharesField) (m IndicationofInterest) {
	m.Header = fix41.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("6"))
	m.Header.Set(field.NewBeginString("FIX.4.1"))
	m.Set(ioiid)
	m.Set(ioitranstype)
	m.Set(symbol)
	m.Set(side)
	m.Set(ioishares)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg IndicationofInterest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "6", r
}

//SetCurrency sets Currency, Tag 15
func (m IndicationofInterest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m IndicationofInterest) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetIOIid sets IOIid, Tag 23
func (m IndicationofInterest) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

//SetIOIOthSvc sets IOIOthSvc, Tag 24
func (m IndicationofInterest) SetIOIOthSvc(v string) {
	m.Set(field.NewIOIOthSvc(v))
}

//SetIOIQltyInd sets IOIQltyInd, Tag 25
func (m IndicationofInterest) SetIOIQltyInd(v string) {
	m.Set(field.NewIOIQltyInd(v))
}

//SetIOIRefID sets IOIRefID, Tag 26
func (m IndicationofInterest) SetIOIRefID(v string) {
	m.Set(field.NewIOIRefID(v))
}

//SetIOIShares sets IOIShares, Tag 27
func (m IndicationofInterest) SetIOIShares(v string) {
	m.Set(field.NewIOIShares(v))
}

//SetIOITransType sets IOITransType, Tag 28
func (m IndicationofInterest) SetIOITransType(v string) {
	m.Set(field.NewIOITransType(v))
}

//SetPrice sets Price, Tag 44
func (m IndicationofInterest) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m IndicationofInterest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m IndicationofInterest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m IndicationofInterest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m IndicationofInterest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m IndicationofInterest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m IndicationofInterest) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m IndicationofInterest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m IndicationofInterest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m IndicationofInterest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetIOINaturalFlag sets IOINaturalFlag, Tag 130
func (m IndicationofInterest) SetIOINaturalFlag(v bool) {
	m.Set(field.NewIOINaturalFlag(v))
}

//SetURLLink sets URLLink, Tag 149
func (m IndicationofInterest) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m IndicationofInterest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetNoIOIQualifiers sets NoIOIQualifiers, Tag 199
func (m IndicationofInterest) SetNoIOIQualifiers(f NoIOIQualifiersRepeatingGroup) {
	m.SetGroup(f)
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m IndicationofInterest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m IndicationofInterest) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m IndicationofInterest) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m IndicationofInterest) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m IndicationofInterest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m IndicationofInterest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//GetCurrency gets Currency, Tag 15
func (m IndicationofInterest) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m IndicationofInterest) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIid gets IOIid, Tag 23
func (m IndicationofInterest) GetIOIid() (f field.IOIidField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIOthSvc gets IOIOthSvc, Tag 24
func (m IndicationofInterest) GetIOIOthSvc() (f field.IOIOthSvcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIQltyInd gets IOIQltyInd, Tag 25
func (m IndicationofInterest) GetIOIQltyInd() (f field.IOIQltyIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIRefID gets IOIRefID, Tag 26
func (m IndicationofInterest) GetIOIRefID() (f field.IOIRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIShares gets IOIShares, Tag 27
func (m IndicationofInterest) GetIOIShares() (f field.IOISharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOITransType gets IOITransType, Tag 28
func (m IndicationofInterest) GetIOITransType() (f field.IOITransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m IndicationofInterest) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m IndicationofInterest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m IndicationofInterest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m IndicationofInterest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m IndicationofInterest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m IndicationofInterest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m IndicationofInterest) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m IndicationofInterest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m IndicationofInterest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m IndicationofInterest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOINaturalFlag gets IOINaturalFlag, Tag 130
func (m IndicationofInterest) GetIOINaturalFlag() (f field.IOINaturalFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetURLLink gets URLLink, Tag 149
func (m IndicationofInterest) GetURLLink() (f field.URLLinkField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m IndicationofInterest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoIOIQualifiers gets NoIOIQualifiers, Tag 199
func (m IndicationofInterest) GetNoIOIQualifiers() (NoIOIQualifiersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoIOIQualifiersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m IndicationofInterest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m IndicationofInterest) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m IndicationofInterest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m IndicationofInterest) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m IndicationofInterest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m IndicationofInterest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m IndicationofInterest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m IndicationofInterest) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasIOIid returns true if IOIid is present, Tag 23
func (m IndicationofInterest) HasIOIid() bool {
	return m.Has(tag.IOIid)
}

//HasIOIOthSvc returns true if IOIOthSvc is present, Tag 24
func (m IndicationofInterest) HasIOIOthSvc() bool {
	return m.Has(tag.IOIOthSvc)
}

//HasIOIQltyInd returns true if IOIQltyInd is present, Tag 25
func (m IndicationofInterest) HasIOIQltyInd() bool {
	return m.Has(tag.IOIQltyInd)
}

//HasIOIRefID returns true if IOIRefID is present, Tag 26
func (m IndicationofInterest) HasIOIRefID() bool {
	return m.Has(tag.IOIRefID)
}

//HasIOIShares returns true if IOIShares is present, Tag 27
func (m IndicationofInterest) HasIOIShares() bool {
	return m.Has(tag.IOIShares)
}

//HasIOITransType returns true if IOITransType is present, Tag 28
func (m IndicationofInterest) HasIOITransType() bool {
	return m.Has(tag.IOITransType)
}

//HasPrice returns true if Price is present, Tag 44
func (m IndicationofInterest) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m IndicationofInterest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m IndicationofInterest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m IndicationofInterest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m IndicationofInterest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m IndicationofInterest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m IndicationofInterest) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m IndicationofInterest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m IndicationofInterest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m IndicationofInterest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasIOINaturalFlag returns true if IOINaturalFlag is present, Tag 130
func (m IndicationofInterest) HasIOINaturalFlag() bool {
	return m.Has(tag.IOINaturalFlag)
}

//HasURLLink returns true if URLLink is present, Tag 149
func (m IndicationofInterest) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m IndicationofInterest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasNoIOIQualifiers returns true if NoIOIQualifiers is present, Tag 199
func (m IndicationofInterest) HasNoIOIQualifiers() bool {
	return m.Has(tag.NoIOIQualifiers)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m IndicationofInterest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m IndicationofInterest) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m IndicationofInterest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m IndicationofInterest) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m IndicationofInterest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m IndicationofInterest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//NoIOIQualifiers is a repeating group element, Tag 199
type NoIOIQualifiers struct {
	quickfix.Group
}

//SetIOIQualifier sets IOIQualifier, Tag 104
func (m NoIOIQualifiers) SetIOIQualifier(v string) {
	m.Set(field.NewIOIQualifier(v))
}

//GetIOIQualifier gets IOIQualifier, Tag 104
func (m NoIOIQualifiers) GetIOIQualifier() (f field.IOIQualifierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIOIQualifier returns true if IOIQualifier is present, Tag 104
func (m NoIOIQualifiers) HasIOIQualifier() bool {
	return m.Has(tag.IOIQualifier)
}

//NoIOIQualifiersRepeatingGroup is a repeating group, Tag 199
type NoIOIQualifiersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoIOIQualifiersRepeatingGroup returns an initialized, NoIOIQualifiersRepeatingGroup
func NewNoIOIQualifiersRepeatingGroup() NoIOIQualifiersRepeatingGroup {
	return NoIOIQualifiersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoIOIQualifiers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.IOIQualifier)})}
}

//Add create and append a new NoIOIQualifiers to this group
func (m NoIOIQualifiersRepeatingGroup) Add() NoIOIQualifiers {
	g := m.RepeatingGroup.Add()
	return NoIOIQualifiers{g}
}

//Get returns the ith NoIOIQualifiers in the NoIOIQualifiersRepeatinGroup
func (m NoIOIQualifiersRepeatingGroup) Get(i int) NoIOIQualifiers {
	return NoIOIQualifiers{m.RepeatingGroup.Get(i)}
}
