package indicationofinterest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//IndicationofInterest is the fix40 IndicationofInterest type, MsgType = 6
type IndicationofInterest struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a IndicationofInterest from a quickfix.Message instance
func FromMessage(m quickfix.Message) IndicationofInterest {
	return IndicationofInterest{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
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
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("6"))
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
	return "FIX.4.0", "6", r
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

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m IndicationofInterest) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m IndicationofInterest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIOIQualifier sets IOIQualifier, Tag 104
func (m IndicationofInterest) SetIOIQualifier(v string) {
	m.Set(field.NewIOIQualifier(v))
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

//GetIOIQualifier gets IOIQualifier, Tag 104
func (m IndicationofInterest) GetIOIQualifier() (f field.IOIQualifierField, err quickfix.MessageRejectError) {
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

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m IndicationofInterest) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m IndicationofInterest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIOIQualifier returns true if IOIQualifier is present, Tag 104
func (m IndicationofInterest) HasIOIQualifier() bool {
	return m.Has(tag.IOIQualifier)
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
