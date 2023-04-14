package news

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/gen/enum"
	"github.com/quickfixgo/quickfix/gen/field"
	"github.com/quickfixgo/quickfix/gen/fix41"
	"github.com/quickfixgo/quickfix/gen/tag"
)

// News is the fix41 News type, MsgType = B.
type News struct {
	fix41.Header
	*quickfix.Body
	fix41.Trailer
	Message *quickfix.Message
}

// FromMessage creates a News from a quickfix.Message instance.
func FromMessage(m *quickfix.Message) News {
	return News{
		Header:  fix41.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix41.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance.
func (m News) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a News initialized with the required fields for News.
func New(headline field.HeadlineField) (m News) {
	m.Message = quickfix.NewMessage()
	m.Header = fix41.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("B"))
	m.Set(headline)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message.
type RouteOut func(msg News, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type.
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "B", r
}

// SetLinesOfText sets LinesOfText, Tag 33.
func (m News) SetLinesOfText(f LinesOfTextRepeatingGroup) {
	m.SetGroup(f)
}

// SetOrigTime sets OrigTime, Tag 42.
func (m News) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

// SetUrgency sets Urgency, Tag 61.
func (m News) SetUrgency(v enum.Urgency) {
	m.Set(field.NewUrgency(v))
}

// SetRawDataLength sets RawDataLength, Tag 95.
func (m News) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

// SetRawData sets RawData, Tag 96.
func (m News) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

// SetNoRelatedSym sets NoRelatedSym, Tag 146.
func (m News) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

// SetHeadline sets Headline, Tag 148.
func (m News) SetHeadline(v string) {
	m.Set(field.NewHeadline(v))
}

// SetURLLink sets URLLink, Tag 149.
func (m News) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

// GetLinesOfText gets LinesOfText, Tag 33.
func (m News) GetLinesOfText() (LinesOfTextRepeatingGroup, quickfix.MessageRejectError) {
	f := NewLinesOfTextRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetOrigTime gets OrigTime, Tag 42.
func (m News) GetOrigTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUrgency gets Urgency, Tag 61.
func (m News) GetUrgency() (v enum.Urgency, err quickfix.MessageRejectError) {
	var f field.UrgencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRawDataLength gets RawDataLength, Tag 95.
func (m News) GetRawDataLength() (v int, err quickfix.MessageRejectError) {
	var f field.RawDataLengthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetRawData gets RawData, Tag 96.
func (m News) GetRawData() (v string, err quickfix.MessageRejectError) {
	var f field.RawDataField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetNoRelatedSym gets NoRelatedSym, Tag 146.
func (m News) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

// GetHeadline gets Headline, Tag 148.
func (m News) GetHeadline() (v string, err quickfix.MessageRejectError) {
	var f field.HeadlineField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetURLLink gets URLLink, Tag 149.
func (m News) GetURLLink() (v string, err quickfix.MessageRejectError) {
	var f field.URLLinkField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasLinesOfText returns true if LinesOfText is present, Tag 33.
func (m News) HasLinesOfText() bool {
	return m.Has(tag.LinesOfText)
}

// HasOrigTime returns true if OrigTime is present, Tag 42.
func (m News) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

// HasUrgency returns true if Urgency is present, Tag 61.
func (m News) HasUrgency() bool {
	return m.Has(tag.Urgency)
}

// HasRawDataLength returns true if RawDataLength is present, Tag 95.
func (m News) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

// HasRawData returns true if RawData is present, Tag 96.
func (m News) HasRawData() bool {
	return m.Has(tag.RawData)
}

// HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146.
func (m News) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

// HasHeadline returns true if Headline is present, Tag 148.
func (m News) HasHeadline() bool {
	return m.Has(tag.Headline)
}

// HasURLLink returns true if URLLink is present, Tag 149.
func (m News) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

// LinesOfText is a repeating group element, Tag 33.
type LinesOfText struct {
	*quickfix.Group
}

// SetText sets Text, Tag 58.
func (m LinesOfText) SetText(v string) {
	m.Set(field.NewText(v))
}

// GetText gets Text, Tag 58.
func (m LinesOfText) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasText returns true if Text is present, Tag 58.
func (m LinesOfText) HasText() bool {
	return m.Has(tag.Text)
}

// LinesOfTextRepeatingGroup is a repeating group, Tag 33.
type LinesOfTextRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewLinesOfTextRepeatingGroup returns an initialized, LinesOfTextRepeatingGroup.
func NewLinesOfTextRepeatingGroup() LinesOfTextRepeatingGroup {
	return LinesOfTextRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.LinesOfText,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Text)})}
}

// Add create and append a new LinesOfText to this group.
func (m LinesOfTextRepeatingGroup) Add() LinesOfText {
	g := m.RepeatingGroup.Add()
	return LinesOfText{g}
}

// Get returns the ith LinesOfText in the LinesOfTextRepeatinGroup.
func (m LinesOfTextRepeatingGroup) Get(i int) LinesOfText {
	return LinesOfText{m.RepeatingGroup.Get(i)}
}

// NoRelatedSym is a repeating group element, Tag 146.
type NoRelatedSym struct {
	*quickfix.Group
}

// SetRelatdSym sets RelatdSym, Tag 46.
func (m NoRelatedSym) SetRelatdSym(v string) {
	m.Set(field.NewRelatdSym(v))
}

// SetSymbolSfx sets SymbolSfx, Tag 65.
func (m NoRelatedSym) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

// SetSecurityID sets SecurityID, Tag 48.
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

// SetIDSource sets IDSource, Tag 22.
func (m NoRelatedSym) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

// SetSecurityType sets SecurityType, Tag 167.
func (m NoRelatedSym) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

// SetMaturityMonthYear sets MaturityMonthYear, Tag 200.
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

// SetMaturityDay sets MaturityDay, Tag 205.
func (m NoRelatedSym) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

// SetPutOrCall sets PutOrCall, Tag 201.
func (m NoRelatedSym) SetPutOrCall(v enum.PutOrCall) {
	m.Set(field.NewPutOrCall(v))
}

// SetStrikePrice sets StrikePrice, Tag 202.
func (m NoRelatedSym) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

// SetOptAttribute sets OptAttribute, Tag 206.
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

// SetSecurityExchange sets SecurityExchange, Tag 207.
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

// SetIssuer sets Issuer, Tag 106.
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

// SetSecurityDesc sets SecurityDesc, Tag 107.
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

// GetRelatdSym gets RelatdSym, Tag 46.
func (m NoRelatedSym) GetRelatdSym() (v string, err quickfix.MessageRejectError) {
	var f field.RelatdSymField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSymbolSfx gets SymbolSfx, Tag 65.
func (m NoRelatedSym) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityID gets SecurityID, Tag 48.
func (m NoRelatedSym) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIDSource gets IDSource, Tag 22.
func (m NoRelatedSym) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityType gets SecurityType, Tag 167.
func (m NoRelatedSym) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityMonthYear gets MaturityMonthYear, Tag 200.
func (m NoRelatedSym) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetMaturityDay gets MaturityDay, Tag 205.
func (m NoRelatedSym) GetMaturityDay() (v int, err quickfix.MessageRejectError) {
	var f field.MaturityDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetPutOrCall gets PutOrCall, Tag 201.
func (m NoRelatedSym) GetPutOrCall() (v enum.PutOrCall, err quickfix.MessageRejectError) {
	var f field.PutOrCallField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetStrikePrice gets StrikePrice, Tag 202.
func (m NoRelatedSym) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetOptAttribute gets OptAttribute, Tag 206.
func (m NoRelatedSym) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityExchange gets SecurityExchange, Tag 207.
func (m NoRelatedSym) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetIssuer gets Issuer, Tag 106.
func (m NoRelatedSym) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetSecurityDesc gets SecurityDesc, Tag 107.
func (m NoRelatedSym) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasRelatdSym returns true if RelatdSym is present, Tag 46.
func (m NoRelatedSym) HasRelatdSym() bool {
	return m.Has(tag.RelatdSym)
}

// HasSymbolSfx returns true if SymbolSfx is present, Tag 65.
func (m NoRelatedSym) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

// HasSecurityID returns true if SecurityID is present, Tag 48.
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

// HasIDSource returns true if IDSource is present, Tag 22.
func (m NoRelatedSym) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

// HasSecurityType returns true if SecurityType is present, Tag 167.
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

// HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200.
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

// HasMaturityDay returns true if MaturityDay is present, Tag 205.
func (m NoRelatedSym) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

// HasPutOrCall returns true if PutOrCall is present, Tag 201.
func (m NoRelatedSym) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

// HasStrikePrice returns true if StrikePrice is present, Tag 202.
func (m NoRelatedSym) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

// HasOptAttribute returns true if OptAttribute is present, Tag 206.
func (m NoRelatedSym) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

// HasSecurityExchange returns true if SecurityExchange is present, Tag 207.
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

// HasIssuer returns true if Issuer is present, Tag 106.
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

// HasSecurityDesc returns true if SecurityDesc is present, Tag 107.
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

// NoRelatedSymRepeatingGroup is a repeating group, Tag 146.
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

// NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup.
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatdSym), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.SecurityDesc)})}
}

// Add create and append a new NoRelatedSym to this group.
func (m NoRelatedSymRepeatingGroup) Add() NoRelatedSym {
	g := m.RepeatingGroup.Add()
	return NoRelatedSym{g}
}

// Get returns the ith NoRelatedSym in the NoRelatedSymRepeatinGroup.
func (m NoRelatedSymRepeatingGroup) Get(i int) NoRelatedSym {
	return NoRelatedSym{m.RepeatingGroup.Get(i)}
}
