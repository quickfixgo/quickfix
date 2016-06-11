package news

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//News is the fix41 News type, MsgType = B
type News struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a News from a quickfix.Message instance
func FromMessage(m quickfix.Message) News {
	return News{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m News) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a News initialized with the required fields for News
func New(headline field.HeadlineField) (m News) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("B"))
	m.Set(headline)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg News, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "B", r
}

//SetLinesOfText sets LinesOfText, Tag 33
func (m News) SetLinesOfText(f LinesOfTextRepeatingGroup) {
	m.SetGroup(f)
}

//SetOrigTime sets OrigTime, Tag 42
func (m News) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

//SetUrgency sets Urgency, Tag 61
func (m News) SetUrgency(v string) {
	m.Set(field.NewUrgency(v))
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m News) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m News) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m News) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetHeadline sets Headline, Tag 148
func (m News) SetHeadline(v string) {
	m.Set(field.NewHeadline(v))
}

//SetURLLink sets URLLink, Tag 149
func (m News) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

//GetLinesOfText gets LinesOfText, Tag 33
func (m News) GetLinesOfText() (LinesOfTextRepeatingGroup, quickfix.MessageRejectError) {
	f := NewLinesOfTextRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetOrigTime gets OrigTime, Tag 42
func (m News) GetOrigTime() (f field.OrigTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUrgency gets Urgency, Tag 61
func (m News) GetUrgency() (f field.UrgencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m News) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m News) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m News) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetHeadline gets Headline, Tag 148
func (m News) GetHeadline() (f field.HeadlineField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetURLLink gets URLLink, Tag 149
func (m News) GetURLLink() (f field.URLLinkField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLinesOfText returns true if LinesOfText is present, Tag 33
func (m News) HasLinesOfText() bool {
	return m.Has(tag.LinesOfText)
}

//HasOrigTime returns true if OrigTime is present, Tag 42
func (m News) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

//HasUrgency returns true if Urgency is present, Tag 61
func (m News) HasUrgency() bool {
	return m.Has(tag.Urgency)
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m News) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m News) HasRawData() bool {
	return m.Has(tag.RawData)
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m News) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasHeadline returns true if Headline is present, Tag 148
func (m News) HasHeadline() bool {
	return m.Has(tag.Headline)
}

//HasURLLink returns true if URLLink is present, Tag 149
func (m News) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

//LinesOfText is a repeating group element, Tag 33
type LinesOfText struct {
	quickfix.Group
}

//SetText sets Text, Tag 58
func (m LinesOfText) SetText(v string) {
	m.Set(field.NewText(v))
}

//GetText gets Text, Tag 58
func (m LinesOfText) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m LinesOfText) HasText() bool {
	return m.Has(tag.Text)
}

//LinesOfTextRepeatingGroup is a repeating group, Tag 33
type LinesOfTextRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewLinesOfTextRepeatingGroup returns an initialized, LinesOfTextRepeatingGroup
func NewLinesOfTextRepeatingGroup() LinesOfTextRepeatingGroup {
	return LinesOfTextRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.LinesOfText,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Text)})}
}

//Add create and append a new LinesOfText to this group
func (m LinesOfTextRepeatingGroup) Add() LinesOfText {
	g := m.RepeatingGroup.Add()
	return LinesOfText{g}
}

//Get returns the ith LinesOfText in the LinesOfTextRepeatinGroup
func (m LinesOfTextRepeatingGroup) Get(i int) LinesOfText {
	return LinesOfText{m.RepeatingGroup.Get(i)}
}

//NoRelatedSym is a repeating group element, Tag 146
type NoRelatedSym struct {
	quickfix.Group
}

//SetRelatdSym sets RelatdSym, Tag 46
func (m NoRelatedSym) SetRelatdSym(v string) {
	m.Set(field.NewRelatdSym(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoRelatedSym) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoRelatedSym) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NoRelatedSym) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoRelatedSym) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoRelatedSym) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NoRelatedSym) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoRelatedSym) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoRelatedSym) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoRelatedSym) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//GetRelatdSym gets RelatdSym, Tag 46
func (m NoRelatedSym) GetRelatdSym() (f field.RelatdSymField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoRelatedSym) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoRelatedSym) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NoRelatedSym) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoRelatedSym) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoRelatedSym) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NoRelatedSym) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoRelatedSym) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoRelatedSym) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoRelatedSym) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoRelatedSym) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoRelatedSym) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoRelatedSym) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRelatdSym returns true if RelatdSym is present, Tag 46
func (m NoRelatedSym) HasRelatdSym() bool {
	return m.Has(tag.RelatdSym)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoRelatedSym) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoRelatedSym) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m NoRelatedSym) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoRelatedSym) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoRelatedSym) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NoRelatedSym) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoRelatedSym) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoRelatedSym) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoRelatedSym) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatdSym), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.SecurityDesc)})}
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
