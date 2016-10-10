package indicationofinterest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//IndicationofInterest is the fix40 IndicationofInterest type, MsgType = 6
type IndicationofInterest struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a IndicationofInterest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) IndicationofInterest {
	return IndicationofInterest{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m IndicationofInterest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a IndicationofInterest initialized with the required fields for IndicationofInterest
func New(ioiid field.IOIidField, ioitranstype field.IOITransTypeField, symbol field.SymbolField, side field.SideField, ioishares field.IOISharesField) (m IndicationofInterest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

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
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "6", r
}

//SetCurrency sets Currency, Tag 15
func (m IndicationofInterest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m IndicationofInterest) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetIOIid sets IOIid, Tag 23
func (m IndicationofInterest) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

//SetIOIOthSvc sets IOIOthSvc, Tag 24
func (m IndicationofInterest) SetIOIOthSvc(v enum.IOIOthSvc) {
	m.Set(field.NewIOIOthSvc(v))
}

//SetIOIQltyInd sets IOIQltyInd, Tag 25
func (m IndicationofInterest) SetIOIQltyInd(v enum.IOIQltyInd) {
	m.Set(field.NewIOIQltyInd(v))
}

//SetIOIRefID sets IOIRefID, Tag 26
func (m IndicationofInterest) SetIOIRefID(v string) {
	m.Set(field.NewIOIRefID(v))
}

//SetIOIShares sets IOIShares, Tag 27
func (m IndicationofInterest) SetIOIShares(v enum.IOIShares) {
	m.Set(field.NewIOIShares(v))
}

//SetIOITransType sets IOITransType, Tag 28
func (m IndicationofInterest) SetIOITransType(v enum.IOITransType) {
	m.Set(field.NewIOITransType(v))
}

//SetPrice sets Price, Tag 44
func (m IndicationofInterest) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m IndicationofInterest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m IndicationofInterest) SetSide(v enum.Side) {
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
func (m IndicationofInterest) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIOIQualifier sets IOIQualifier, Tag 104
func (m IndicationofInterest) SetIOIQualifier(v enum.IOIQualifier) {
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
func (m IndicationofInterest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m IndicationofInterest) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIid gets IOIid, Tag 23
func (m IndicationofInterest) GetIOIid() (v string, err quickfix.MessageRejectError) {
	var f field.IOIidField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIOthSvc gets IOIOthSvc, Tag 24
func (m IndicationofInterest) GetIOIOthSvc() (v enum.IOIOthSvc, err quickfix.MessageRejectError) {
	var f field.IOIOthSvcField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIQltyInd gets IOIQltyInd, Tag 25
func (m IndicationofInterest) GetIOIQltyInd() (v enum.IOIQltyInd, err quickfix.MessageRejectError) {
	var f field.IOIQltyIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIRefID gets IOIRefID, Tag 26
func (m IndicationofInterest) GetIOIRefID() (v string, err quickfix.MessageRejectError) {
	var f field.IOIRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIShares gets IOIShares, Tag 27
func (m IndicationofInterest) GetIOIShares() (v enum.IOIShares, err quickfix.MessageRejectError) {
	var f field.IOISharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOITransType gets IOITransType, Tag 28
func (m IndicationofInterest) GetIOITransType() (v enum.IOITransType, err quickfix.MessageRejectError) {
	var f field.IOITransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrice gets Price, Tag 44
func (m IndicationofInterest) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m IndicationofInterest) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m IndicationofInterest) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m IndicationofInterest) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m IndicationofInterest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m IndicationofInterest) GetValidUntilTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ValidUntilTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m IndicationofInterest) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOIQualifier gets IOIQualifier, Tag 104
func (m IndicationofInterest) GetIOIQualifier() (v enum.IOIQualifier, err quickfix.MessageRejectError) {
	var f field.IOIQualifierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m IndicationofInterest) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m IndicationofInterest) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIOINaturalFlag gets IOINaturalFlag, Tag 130
func (m IndicationofInterest) GetIOINaturalFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.IOINaturalFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
