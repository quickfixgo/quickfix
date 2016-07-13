package applicationmessagereport

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ApplicationMessageReport is the fix50sp1 ApplicationMessageReport type, MsgType = BY
type ApplicationMessageReport struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ApplicationMessageReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) ApplicationMessageReport {
	return ApplicationMessageReport{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ApplicationMessageReport initialized with the required fields for ApplicationMessageReport
func New(applreportid field.ApplReportIDField, applreporttype field.ApplReportTypeField) (m ApplicationMessageReport) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BY"))
	m.Header.Set(field.NewBeginString("8"))
	m.Set(applreportid)
	m.Set(applreporttype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "BY", r
}

//SetText sets Text, Tag 58
func (m ApplicationMessageReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ApplicationMessageReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ApplicationMessageReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetNoApplIDs sets NoApplIDs, Tag 1351
func (m ApplicationMessageReport) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetApplReportID sets ApplReportID, Tag 1356
func (m ApplicationMessageReport) SetApplReportID(v string) {
	m.Set(field.NewApplReportID(v))
}

//SetApplReportType sets ApplReportType, Tag 1426
func (m ApplicationMessageReport) SetApplReportType(v int) {
	m.Set(field.NewApplReportType(v))
}

//GetText gets Text, Tag 58
func (m ApplicationMessageReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ApplicationMessageReport) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ApplicationMessageReport) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoApplIDs gets NoApplIDs, Tag 1351
func (m ApplicationMessageReport) GetNoApplIDs() (NoApplIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoApplIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetApplReportID gets ApplReportID, Tag 1356
func (m ApplicationMessageReport) GetApplReportID() (f field.ApplReportIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplReportType gets ApplReportType, Tag 1426
func (m ApplicationMessageReport) GetApplReportType() (f field.ApplReportTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ApplicationMessageReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ApplicationMessageReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ApplicationMessageReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasNoApplIDs returns true if NoApplIDs is present, Tag 1351
func (m ApplicationMessageReport) HasNoApplIDs() bool {
	return m.Has(tag.NoApplIDs)
}

//HasApplReportID returns true if ApplReportID is present, Tag 1356
func (m ApplicationMessageReport) HasApplReportID() bool {
	return m.Has(tag.ApplReportID)
}

//HasApplReportType returns true if ApplReportType is present, Tag 1426
func (m ApplicationMessageReport) HasApplReportType() bool {
	return m.Has(tag.ApplReportType)
}

//NoApplIDs is a repeating group element, Tag 1351
type NoApplIDs struct {
	quickfix.Group
}

//SetRefApplID sets RefApplID, Tag 1355
func (m NoApplIDs) SetRefApplID(v string) {
	m.Set(field.NewRefApplID(v))
}

//SetApplNewSeqNum sets ApplNewSeqNum, Tag 1399
func (m NoApplIDs) SetApplNewSeqNum(v int) {
	m.Set(field.NewApplNewSeqNum(v))
}

//SetRefApplLastSeqNum sets RefApplLastSeqNum, Tag 1357
func (m NoApplIDs) SetRefApplLastSeqNum(v int) {
	m.Set(field.NewRefApplLastSeqNum(v))
}

//GetRefApplID gets RefApplID, Tag 1355
func (m NoApplIDs) GetRefApplID() (f field.RefApplIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplNewSeqNum gets ApplNewSeqNum, Tag 1399
func (m NoApplIDs) GetApplNewSeqNum() (f field.ApplNewSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefApplLastSeqNum gets RefApplLastSeqNum, Tag 1357
func (m NoApplIDs) GetRefApplLastSeqNum() (f field.RefApplLastSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRefApplID returns true if RefApplID is present, Tag 1355
func (m NoApplIDs) HasRefApplID() bool {
	return m.Has(tag.RefApplID)
}

//HasApplNewSeqNum returns true if ApplNewSeqNum is present, Tag 1399
func (m NoApplIDs) HasApplNewSeqNum() bool {
	return m.Has(tag.ApplNewSeqNum)
}

//HasRefApplLastSeqNum returns true if RefApplLastSeqNum is present, Tag 1357
func (m NoApplIDs) HasRefApplLastSeqNum() bool {
	return m.Has(tag.RefApplLastSeqNum)
}

//NoApplIDsRepeatingGroup is a repeating group, Tag 1351
type NoApplIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoApplIDsRepeatingGroup returns an initialized, NoApplIDsRepeatingGroup
func NewNoApplIDsRepeatingGroup() NoApplIDsRepeatingGroup {
	return NoApplIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoApplIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.ApplNewSeqNum), quickfix.GroupElement(tag.RefApplLastSeqNum)})}
}

//Add create and append a new NoApplIDs to this group
func (m NoApplIDsRepeatingGroup) Add() NoApplIDs {
	g := m.RepeatingGroup.Add()
	return NoApplIDs{g}
}

//Get returns the ith NoApplIDs in the NoApplIDsRepeatinGroup
func (m NoApplIDsRepeatingGroup) Get(i int) NoApplIDs {
	return NoApplIDs{m.RepeatingGroup.Get(i)}
}
