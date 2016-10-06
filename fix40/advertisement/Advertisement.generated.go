package advertisement

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Advertisement is the fix40 Advertisement type, MsgType = 7
type Advertisement struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Advertisement from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Advertisement {
	return Advertisement{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Advertisement) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Advertisement initialized with the required fields for Advertisement
func New(advid field.AdvIdField, advtranstype field.AdvTransTypeField, symbol field.SymbolField, advside field.AdvSideField, shares field.SharesField) (m Advertisement) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("7"))
	m.Set(advid)
	m.Set(advtranstype)
	m.Set(symbol)
	m.Set(advside)
	m.Set(shares)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Advertisement, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "7", r
}

//SetAdvId sets AdvId, Tag 2
func (m Advertisement) SetAdvId(v string) {
	m.Set(field.NewAdvId(v))
}

//SetAdvRefID sets AdvRefID, Tag 3
func (m Advertisement) SetAdvRefID(v string) {
	m.Set(field.NewAdvRefID(v))
}

//SetAdvSide sets AdvSide, Tag 4
func (m Advertisement) SetAdvSide(v enum.AdvSide) {
	m.Set(field.NewAdvSide(v))
}

//SetAdvTransType sets AdvTransType, Tag 5
func (m Advertisement) SetAdvTransType(v enum.AdvTransType) {
	m.Set(field.NewAdvTransType(v))
}

//SetCurrency sets Currency, Tag 15
func (m Advertisement) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m Advertisement) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetPrice sets Price, Tag 44
func (m Advertisement) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Advertisement) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetShares sets Shares, Tag 53
func (m Advertisement) SetShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewShares(value, scale))
}

//SetSymbol sets Symbol, Tag 55
func (m Advertisement) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m Advertisement) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Advertisement) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Advertisement) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m Advertisement) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Advertisement) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//GetAdvId gets AdvId, Tag 2
func (m Advertisement) GetAdvId() (v string, err quickfix.MessageRejectError) {
	var f field.AdvIdField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAdvRefID gets AdvRefID, Tag 3
func (m Advertisement) GetAdvRefID() (v string, err quickfix.MessageRejectError) {
	var f field.AdvRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAdvSide gets AdvSide, Tag 4
func (m Advertisement) GetAdvSide() (v enum.AdvSide, err quickfix.MessageRejectError) {
	var f field.AdvSideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAdvTransType gets AdvTransType, Tag 5
func (m Advertisement) GetAdvTransType() (v enum.AdvTransType, err quickfix.MessageRejectError) {
	var f field.AdvTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m Advertisement) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m Advertisement) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPrice gets Price, Tag 44
func (m Advertisement) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Advertisement) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetShares gets Shares, Tag 53
func (m Advertisement) GetShares() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Advertisement) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m Advertisement) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Advertisement) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Advertisement) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m Advertisement) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Advertisement) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasAdvId returns true if AdvId is present, Tag 2
func (m Advertisement) HasAdvId() bool {
	return m.Has(tag.AdvId)
}

//HasAdvRefID returns true if AdvRefID is present, Tag 3
func (m Advertisement) HasAdvRefID() bool {
	return m.Has(tag.AdvRefID)
}

//HasAdvSide returns true if AdvSide is present, Tag 4
func (m Advertisement) HasAdvSide() bool {
	return m.Has(tag.AdvSide)
}

//HasAdvTransType returns true if AdvTransType is present, Tag 5
func (m Advertisement) HasAdvTransType() bool {
	return m.Has(tag.AdvTransType)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m Advertisement) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m Advertisement) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasPrice returns true if Price is present, Tag 44
func (m Advertisement) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Advertisement) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasShares returns true if Shares is present, Tag 53
func (m Advertisement) HasShares() bool {
	return m.Has(tag.Shares)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Advertisement) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m Advertisement) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Advertisement) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Advertisement) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Advertisement) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Advertisement) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}
