package advertisement

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Advertisement is the fix40 Advertisement type, MsgType = 7
type Advertisement struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Advertisement from a quickfix.Message instance
func FromMessage(m quickfix.Message) Advertisement {
	return Advertisement{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Advertisement) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Advertisement initialized with the required fields for Advertisement
func New(advid field.AdvIdField, advtranstype field.AdvTransTypeField, symbol field.SymbolField, advside field.AdvSideField, shares field.SharesField) (m Advertisement) {
	m.Header = fix40.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("7"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m Advertisement) SetAdvSide(v string) {
	m.Set(field.NewAdvSide(v))
}

//SetAdvTransType sets AdvTransType, Tag 5
func (m Advertisement) SetAdvTransType(v string) {
	m.Set(field.NewAdvTransType(v))
}

//SetCurrency sets Currency, Tag 15
func (m Advertisement) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m Advertisement) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetPrice sets Price, Tag 44
func (m Advertisement) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Advertisement) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetShares sets Shares, Tag 53
func (m Advertisement) SetShares(v float64) {
	m.Set(field.NewShares(v))
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
func (m Advertisement) SetSymbolSfx(v string) {
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
func (m Advertisement) GetAdvId() (f field.AdvIdField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAdvRefID gets AdvRefID, Tag 3
func (m Advertisement) GetAdvRefID() (f field.AdvRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAdvSide gets AdvSide, Tag 4
func (m Advertisement) GetAdvSide() (f field.AdvSideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAdvTransType gets AdvTransType, Tag 5
func (m Advertisement) GetAdvTransType() (f field.AdvTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m Advertisement) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m Advertisement) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m Advertisement) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Advertisement) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShares gets Shares, Tag 53
func (m Advertisement) GetShares() (f field.SharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Advertisement) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m Advertisement) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Advertisement) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Advertisement) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m Advertisement) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Advertisement) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
