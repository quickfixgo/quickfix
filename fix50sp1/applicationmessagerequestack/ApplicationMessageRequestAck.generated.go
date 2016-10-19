package applicationmessagerequestack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ApplicationMessageRequestAck is the fix50sp1 ApplicationMessageRequestAck type, MsgType = BX
type ApplicationMessageRequestAck struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a ApplicationMessageRequestAck from a quickfix.Message instance
func FromMessage(m *quickfix.Message) ApplicationMessageRequestAck {
	return ApplicationMessageRequestAck{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ApplicationMessageRequestAck) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a ApplicationMessageRequestAck initialized with the required fields for ApplicationMessageRequestAck
func New(applresponseid field.ApplResponseIDField) (m ApplicationMessageRequestAck) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BX"))
	m.Set(applresponseid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ApplicationMessageRequestAck, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "BX", r
}

//SetText sets Text, Tag 58
func (m ApplicationMessageRequestAck) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ApplicationMessageRequestAck) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ApplicationMessageRequestAck) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetApplReqID sets ApplReqID, Tag 1346
func (m ApplicationMessageRequestAck) SetApplReqID(v string) {
	m.Set(field.NewApplReqID(v))
}

//SetApplReqType sets ApplReqType, Tag 1347
func (m ApplicationMessageRequestAck) SetApplReqType(v enum.ApplReqType) {
	m.Set(field.NewApplReqType(v))
}

//SetApplResponseType sets ApplResponseType, Tag 1348
func (m ApplicationMessageRequestAck) SetApplResponseType(v enum.ApplResponseType) {
	m.Set(field.NewApplResponseType(v))
}

//SetApplTotalMessageCount sets ApplTotalMessageCount, Tag 1349
func (m ApplicationMessageRequestAck) SetApplTotalMessageCount(v int) {
	m.Set(field.NewApplTotalMessageCount(v))
}

//SetNoApplIDs sets NoApplIDs, Tag 1351
func (m ApplicationMessageRequestAck) SetNoApplIDs(f NoApplIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetApplResponseID sets ApplResponseID, Tag 1353
func (m ApplicationMessageRequestAck) SetApplResponseID(v string) {
	m.Set(field.NewApplResponseID(v))
}

//GetText gets Text, Tag 58
func (m ApplicationMessageRequestAck) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ApplicationMessageRequestAck) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ApplicationMessageRequestAck) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplReqID gets ApplReqID, Tag 1346
func (m ApplicationMessageRequestAck) GetApplReqID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplReqType gets ApplReqType, Tag 1347
func (m ApplicationMessageRequestAck) GetApplReqType() (v enum.ApplReqType, err quickfix.MessageRejectError) {
	var f field.ApplReqTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResponseType gets ApplResponseType, Tag 1348
func (m ApplicationMessageRequestAck) GetApplResponseType() (v enum.ApplResponseType, err quickfix.MessageRejectError) {
	var f field.ApplResponseTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplTotalMessageCount gets ApplTotalMessageCount, Tag 1349
func (m ApplicationMessageRequestAck) GetApplTotalMessageCount() (v int, err quickfix.MessageRejectError) {
	var f field.ApplTotalMessageCountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoApplIDs gets NoApplIDs, Tag 1351
func (m ApplicationMessageRequestAck) GetNoApplIDs() (NoApplIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoApplIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetApplResponseID gets ApplResponseID, Tag 1353
func (m ApplicationMessageRequestAck) GetApplResponseID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplResponseIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m ApplicationMessageRequestAck) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ApplicationMessageRequestAck) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ApplicationMessageRequestAck) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasApplReqID returns true if ApplReqID is present, Tag 1346
func (m ApplicationMessageRequestAck) HasApplReqID() bool {
	return m.Has(tag.ApplReqID)
}

//HasApplReqType returns true if ApplReqType is present, Tag 1347
func (m ApplicationMessageRequestAck) HasApplReqType() bool {
	return m.Has(tag.ApplReqType)
}

//HasApplResponseType returns true if ApplResponseType is present, Tag 1348
func (m ApplicationMessageRequestAck) HasApplResponseType() bool {
	return m.Has(tag.ApplResponseType)
}

//HasApplTotalMessageCount returns true if ApplTotalMessageCount is present, Tag 1349
func (m ApplicationMessageRequestAck) HasApplTotalMessageCount() bool {
	return m.Has(tag.ApplTotalMessageCount)
}

//HasNoApplIDs returns true if NoApplIDs is present, Tag 1351
func (m ApplicationMessageRequestAck) HasNoApplIDs() bool {
	return m.Has(tag.NoApplIDs)
}

//HasApplResponseID returns true if ApplResponseID is present, Tag 1353
func (m ApplicationMessageRequestAck) HasApplResponseID() bool {
	return m.Has(tag.ApplResponseID)
}

//NoApplIDs is a repeating group element, Tag 1351
type NoApplIDs struct {
	*quickfix.Group
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

//SetRefApplLastSeqNum sets RefApplLastSeqNum, Tag 1357
func (m NoApplIDs) SetRefApplLastSeqNum(v int) {
	m.Set(field.NewRefApplLastSeqNum(v))
}

//SetApplResponseError sets ApplResponseError, Tag 1354
func (m NoApplIDs) SetApplResponseError(v enum.ApplResponseError) {
	m.Set(field.NewApplResponseError(v))
}

//GetRefApplID gets RefApplID, Tag 1355
func (m NoApplIDs) GetRefApplID() (v string, err quickfix.MessageRejectError) {
	var f field.RefApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplBegSeqNum gets ApplBegSeqNum, Tag 1182
func (m NoApplIDs) GetApplBegSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplBegSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplEndSeqNum gets ApplEndSeqNum, Tag 1183
func (m NoApplIDs) GetApplEndSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplEndSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefApplLastSeqNum gets RefApplLastSeqNum, Tag 1357
func (m NoApplIDs) GetRefApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.RefApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResponseError gets ApplResponseError, Tag 1354
func (m NoApplIDs) GetApplResponseError() (v enum.ApplResponseError, err quickfix.MessageRejectError) {
	var f field.ApplResponseErrorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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

//HasRefApplLastSeqNum returns true if RefApplLastSeqNum is present, Tag 1357
func (m NoApplIDs) HasRefApplLastSeqNum() bool {
	return m.Has(tag.RefApplLastSeqNum)
}

//HasApplResponseError returns true if ApplResponseError is present, Tag 1354
func (m NoApplIDs) HasApplResponseError() bool {
	return m.Has(tag.ApplResponseError)
}

//NoApplIDsRepeatingGroup is a repeating group, Tag 1351
type NoApplIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoApplIDsRepeatingGroup returns an initialized, NoApplIDsRepeatingGroup
func NewNoApplIDsRepeatingGroup() NoApplIDsRepeatingGroup {
	return NoApplIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoApplIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefApplID), quickfix.GroupElement(tag.ApplBegSeqNum), quickfix.GroupElement(tag.ApplEndSeqNum), quickfix.GroupElement(tag.RefApplLastSeqNum), quickfix.GroupElement(tag.ApplResponseError)})}
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
