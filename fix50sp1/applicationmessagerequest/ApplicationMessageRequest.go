package applicationmessagerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ApplicationMessageRequest is the fix50sp1 ApplicationMessageRequest type, MsgType = BW
type ApplicationMessageRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a ApplicationMessageRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ApplicationMessageRequest {
	return ApplicationMessageRequest{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a ApplicationMessageRequest initialized with the required fields for ApplicationMessageRequest
func New(applreqid field.ApplReqIDField, applreqtype field.ApplReqTypeField) (m ApplicationMessageRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BW"))
	m.Set(applreqid)
	m.Set(applreqtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "BW", r
}

//SetText sets Text, Tag 58
func (m ApplicationMessageRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ApplicationMessageRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ApplicationMessageRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetApplReqID sets ApplReqID, Tag 1346
func (m ApplicationMessageRequest) SetApplReqID(v string) {
	m.Set(field.NewApplReqID(v))
}

//SetApplReqType sets ApplReqType, Tag 1347
func (m ApplicationMessageRequest) SetApplReqType(v int) {
	m.Set(field.NewApplReqType(v))
}

//SetNoApplIDs sets NoApplIDs, Tag 1351
func (m ApplicationMessageRequest) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
	m.SetGroup(f)
}

//GetText gets Text, Tag 58
func (m ApplicationMessageRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ApplicationMessageRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ApplicationMessageRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplReqID gets ApplReqID, Tag 1346
func (m ApplicationMessageRequest) GetApplReqID() (f field.ApplReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplReqType gets ApplReqType, Tag 1347
func (m ApplicationMessageRequest) GetApplReqType() (f field.ApplReqTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoApplIDs gets NoApplIDs, Tag 1351
func (m ApplicationMessageRequest) GetNoApplIDs() (NoApplIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoApplIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasText returns true if Text is present, Tag 58
func (m ApplicationMessageRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ApplicationMessageRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ApplicationMessageRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasApplReqID returns true if ApplReqID is present, Tag 1346
func (m ApplicationMessageRequest) HasApplReqID() bool {
	return m.Has(tag.ApplReqID)
}

//HasApplReqType returns true if ApplReqType is present, Tag 1347
func (m ApplicationMessageRequest) HasApplReqType() bool {
	return m.Has(tag.ApplReqType)
}

//HasNoApplIDs returns true if NoApplIDs is present, Tag 1351
func (m ApplicationMessageRequest) HasNoApplIDs() bool {
	return m.Has(tag.NoApplIDs)
}

//NoApplIDs is a repeating group element, Tag 1351
type NoApplIDs struct {
	quickfix.Group
}

//SetRefApplID sets RefApplID, Tag 1355
func (m NoApplIDs) SetRefApplID(v string) {
	m.Set(field.NewRefApplID(v))
}

//SetApplBegSeqNum sets ApplBegSeqNum, Tag 1182
func (m NoApplIDs) SetApplBegSeqNum(v int) {
	m.Set(field.NewApplBegSeqNum(v))
}

//SetApplEndSeqNum sets ApplEndSeqNum, Tag 1183
func (m NoApplIDs) SetApplEndSeqNum(v int) {
	m.Set(field.NewApplEndSeqNum(v))
}

//GetRefApplID gets RefApplID, Tag 1355
func (m NoApplIDs) GetRefApplID() (f field.RefApplIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplBegSeqNum gets ApplBegSeqNum, Tag 1182
func (m NoApplIDs) GetApplBegSeqNum() (f field.ApplBegSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetApplEndSeqNum gets ApplEndSeqNum, Tag 1183
func (m NoApplIDs) GetApplEndSeqNum() (f field.ApplEndSeqNumField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRefApplID returns true if RefApplID is present, Tag 1355
func (m NoApplIDs) HasRefApplID() bool {
	return m.Has(tag.RefApplID)
}

//HasApplBegSeqNum returns true if ApplBegSeqNum is present, Tag 1182
func (m NoApplIDs) HasApplBegSeqNum() bool {
	return m.Has(tag.ApplBegSeqNum)
}

//HasApplEndSeqNum returns true if ApplEndSeqNum is present, Tag 1183
func (m NoApplIDs) HasApplEndSeqNum() bool {
	return m.Has(tag.ApplEndSeqNum)
}

//NoApplIDsRepeatingGroup is a repeating group, Tag 1351
type NoApplIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoApplIDsRepeatingGroup returns an initialized, NoApplIDsRepeatingGroup
func NewNoApplIDsRepeatingGroup() NoApplIDsRepeatingGroup {
	return NoApplIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoApplIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.ApplBegSeqNum), quickfix.GroupElement(tag.ApplEndSeqNum)})}
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
