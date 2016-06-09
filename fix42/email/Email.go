package email

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//Email is the fix42 Email type, MsgType = C
type Email struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
}

//FromMessage creates a Email from a quickfix.Message instance
func FromMessage(m quickfix.Message) Email {
	return Email{
		Header:  fix42.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix42.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m Email) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a Email initialized with the required fields for Email
func New(emailthreadid field.EmailThreadIDField, emailtype field.EmailTypeField, subject field.SubjectField) (m Email) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("C"))
	m.Set(emailthreadid)
	m.Set(emailtype)
	m.Set(subject)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Email, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "C", r
}

//SetClOrdID sets ClOrdID, Tag 11
func (m Email) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetLinesOfText sets LinesOfText, Tag 33
func (m Email) SetLinesOfText(f LinesOfTextRepeatingGroup) {
	m.SetGroup(f)
}

//SetOrderID sets OrderID, Tag 37
func (m Email) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrigTime sets OrigTime, Tag 42
func (m Email) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

//SetEmailType sets EmailType, Tag 94
func (m Email) SetEmailType(v string) {
	m.Set(field.NewEmailType(v))
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m Email) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m Email) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//SetNoRelatedSym sets NoRelatedSym, Tag 146
func (m Email) SetNoRelatedSym(f NoRelatedSymRepeatingGroup) {
	m.SetGroup(f)
}

//SetSubject sets Subject, Tag 147
func (m Email) SetSubject(v string) {
	m.Set(field.NewSubject(v))
}

//SetEmailThreadID sets EmailThreadID, Tag 164
func (m Email) SetEmailThreadID(v string) {
	m.Set(field.NewEmailThreadID(v))
}

//SetNoRoutingIDs sets NoRoutingIDs, Tag 215
func (m Email) SetNoRoutingIDs(f NoRoutingIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetEncodedSubjectLen sets EncodedSubjectLen, Tag 356
func (m Email) SetEncodedSubjectLen(v int) {
	m.Set(field.NewEncodedSubjectLen(v))
}

//SetEncodedSubject sets EncodedSubject, Tag 357
func (m Email) SetEncodedSubject(v string) {
	m.Set(field.NewEncodedSubject(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m Email) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLinesOfText gets LinesOfText, Tag 33
func (m Email) GetLinesOfText() (LinesOfTextRepeatingGroup, quickfix.MessageRejectError) {
	f := NewLinesOfTextRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetOrderID gets OrderID, Tag 37
func (m Email) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTime gets OrigTime, Tag 42
func (m Email) GetOrigTime() (f field.OrigTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEmailType gets EmailType, Tag 94
func (m Email) GetEmailType() (f field.EmailTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m Email) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m Email) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRelatedSym gets NoRelatedSym, Tag 146
func (m Email) GetNoRelatedSym() (NoRelatedSymRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRelatedSymRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSubject gets Subject, Tag 147
func (m Email) GetSubject() (f field.SubjectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEmailThreadID gets EmailThreadID, Tag 164
func (m Email) GetEmailThreadID() (f field.EmailThreadIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRoutingIDs gets NoRoutingIDs, Tag 215
func (m Email) GetNoRoutingIDs() (NoRoutingIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRoutingIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetEncodedSubjectLen gets EncodedSubjectLen, Tag 356
func (m Email) GetEncodedSubjectLen() (f field.EncodedSubjectLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSubject gets EncodedSubject, Tag 357
func (m Email) GetEncodedSubject() (f field.EncodedSubjectField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m Email) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasLinesOfText returns true if LinesOfText is present, Tag 33
func (m Email) HasLinesOfText() bool {
	return m.Has(tag.LinesOfText)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m Email) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrigTime returns true if OrigTime is present, Tag 42
func (m Email) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

//HasEmailType returns true if EmailType is present, Tag 94
func (m Email) HasEmailType() bool {
	return m.Has(tag.EmailType)
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m Email) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m Email) HasRawData() bool {
	return m.Has(tag.RawData)
}

//HasNoRelatedSym returns true if NoRelatedSym is present, Tag 146
func (m Email) HasNoRelatedSym() bool {
	return m.Has(tag.NoRelatedSym)
}

//HasSubject returns true if Subject is present, Tag 147
func (m Email) HasSubject() bool {
	return m.Has(tag.Subject)
}

//HasEmailThreadID returns true if EmailThreadID is present, Tag 164
func (m Email) HasEmailThreadID() bool {
	return m.Has(tag.EmailThreadID)
}

//HasNoRoutingIDs returns true if NoRoutingIDs is present, Tag 215
func (m Email) HasNoRoutingIDs() bool {
	return m.Has(tag.NoRoutingIDs)
}

//HasEncodedSubjectLen returns true if EncodedSubjectLen is present, Tag 356
func (m Email) HasEncodedSubjectLen() bool {
	return m.Has(tag.EncodedSubjectLen)
}

//HasEncodedSubject returns true if EncodedSubject is present, Tag 357
func (m Email) HasEncodedSubject() bool {
	return m.Has(tag.EncodedSubject)
}

//LinesOfText is a repeating group element, Tag 33
type LinesOfText struct {
	quickfix.Group
}

//SetText sets Text, Tag 58
func (m LinesOfText) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m LinesOfText) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m LinesOfText) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetText gets Text, Tag 58
func (m LinesOfText) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m LinesOfText) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m LinesOfText) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m LinesOfText) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m LinesOfText) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m LinesOfText) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//LinesOfTextRepeatingGroup is a repeating group, Tag 33
type LinesOfTextRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewLinesOfTextRepeatingGroup returns an initialized, LinesOfTextRepeatingGroup
func NewLinesOfTextRepeatingGroup() LinesOfTextRepeatingGroup {
	return LinesOfTextRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.LinesOfText,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
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

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoRelatedSym) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoRelatedSym) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoRelatedSym) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoRelatedSym) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoRelatedSym) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoRelatedSym) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
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

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoRelatedSym) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoRelatedSym) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
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

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoRelatedSym) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoRelatedSym) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoRelatedSym) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoRelatedSym) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoRelatedSym) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
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

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoRelatedSym) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoRelatedSym) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoRelatedSym) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoRelatedSym) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoRelatedSym) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoRelatedSym) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoRelatedSym) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoRelatedSym) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoRelatedSym) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//NoRelatedSymRepeatingGroup is a repeating group, Tag 146
type NoRelatedSymRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRelatedSymRepeatingGroup returns an initialized, NoRelatedSymRepeatingGroup
func NewNoRelatedSymRepeatingGroup() NoRelatedSymRepeatingGroup {
	return NoRelatedSymRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRelatedSym,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RelatdSym), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc)})}
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

//NoRoutingIDs is a repeating group element, Tag 215
type NoRoutingIDs struct {
	quickfix.Group
}

//SetRoutingType sets RoutingType, Tag 216
func (m NoRoutingIDs) SetRoutingType(v int) {
	m.Set(field.NewRoutingType(v))
}

//SetRoutingID sets RoutingID, Tag 217
func (m NoRoutingIDs) SetRoutingID(v string) {
	m.Set(field.NewRoutingID(v))
}

//GetRoutingType gets RoutingType, Tag 216
func (m NoRoutingIDs) GetRoutingType() (f field.RoutingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoutingID gets RoutingID, Tag 217
func (m NoRoutingIDs) GetRoutingID() (f field.RoutingIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRoutingType returns true if RoutingType is present, Tag 216
func (m NoRoutingIDs) HasRoutingType() bool {
	return m.Has(tag.RoutingType)
}

//HasRoutingID returns true if RoutingID is present, Tag 217
func (m NoRoutingIDs) HasRoutingID() bool {
	return m.Has(tag.RoutingID)
}

//NoRoutingIDsRepeatingGroup is a repeating group, Tag 215
type NoRoutingIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRoutingIDsRepeatingGroup returns an initialized, NoRoutingIDsRepeatingGroup
func NewNoRoutingIDsRepeatingGroup() NoRoutingIDsRepeatingGroup {
	return NoRoutingIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRoutingIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RoutingType), quickfix.GroupElement(tag.RoutingID)})}
}

//Add create and append a new NoRoutingIDs to this group
func (m NoRoutingIDsRepeatingGroup) Add() NoRoutingIDs {
	g := m.RepeatingGroup.Add()
	return NoRoutingIDs{g}
}

//Get returns the ith NoRoutingIDs in the NoRoutingIDsRepeatinGroup
func (m NoRoutingIDsRepeatingGroup) Get(i int) NoRoutingIDs {
	return NoRoutingIDs{m.RepeatingGroup.Get(i)}
}
