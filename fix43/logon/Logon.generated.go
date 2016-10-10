package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//Logon is the fix43 Logon type, MsgType = A
type Logon struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Logon from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Logon {
	return Logon{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Logon) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Logon initialized with the required fields for Logon
func New(encryptmethod field.EncryptMethodField, heartbtint field.HeartBtIntField) (m Logon) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("A"))
	m.Set(encryptmethod)
	m.Set(heartbtint)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Logon, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "A", r
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m Logon) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m Logon) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//SetEncryptMethod sets EncryptMethod, Tag 98
func (m Logon) SetEncryptMethod(v enum.EncryptMethod) {
	m.Set(field.NewEncryptMethod(v))
}

//SetHeartBtInt sets HeartBtInt, Tag 108
func (m Logon) SetHeartBtInt(v int) {
	m.Set(field.NewHeartBtInt(v))
}

//SetResetSeqNumFlag sets ResetSeqNumFlag, Tag 141
func (m Logon) SetResetSeqNumFlag(v bool) {
	m.Set(field.NewResetSeqNumFlag(v))
}

//SetMaxMessageSize sets MaxMessageSize, Tag 383
func (m Logon) SetMaxMessageSize(v int) {
	m.Set(field.NewMaxMessageSize(v))
}

//SetNoMsgTypes sets NoMsgTypes, Tag 384
func (m Logon) SetNoMsgTypes(f NoMsgTypesRepeatingGroup) {
	m.SetGroup(f)
}

//SetTestMessageIndicator sets TestMessageIndicator, Tag 464
func (m Logon) SetTestMessageIndicator(v bool) {
	m.Set(field.NewTestMessageIndicator(v))
}

//SetUsername sets Username, Tag 553
func (m Logon) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

//SetPassword sets Password, Tag 554
func (m Logon) SetPassword(v string) {
	m.Set(field.NewPassword(v))
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m Logon) GetRawDataLength() (v int, err quickfix.MessageRejectError) {
	var f field.RawDataLengthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRawData gets RawData, Tag 96
func (m Logon) GetRawData() (v string, err quickfix.MessageRejectError) {
	var f field.RawDataField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncryptMethod gets EncryptMethod, Tag 98
func (m Logon) GetEncryptMethod() (v enum.EncryptMethod, err quickfix.MessageRejectError) {
	var f field.EncryptMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHeartBtInt gets HeartBtInt, Tag 108
func (m Logon) GetHeartBtInt() (v int, err quickfix.MessageRejectError) {
	var f field.HeartBtIntField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetResetSeqNumFlag gets ResetSeqNumFlag, Tag 141
func (m Logon) GetResetSeqNumFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ResetSeqNumFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxMessageSize gets MaxMessageSize, Tag 383
func (m Logon) GetMaxMessageSize() (v int, err quickfix.MessageRejectError) {
	var f field.MaxMessageSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMsgTypes gets NoMsgTypes, Tag 384
func (m Logon) GetNoMsgTypes() (NoMsgTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMsgTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTestMessageIndicator gets TestMessageIndicator, Tag 464
func (m Logon) GetTestMessageIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.TestMessageIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUsername gets Username, Tag 553
func (m Logon) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPassword gets Password, Tag 554
func (m Logon) GetPassword() (v string, err quickfix.MessageRejectError) {
	var f field.PasswordField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m Logon) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m Logon) HasRawData() bool {
	return m.Has(tag.RawData)
}

//HasEncryptMethod returns true if EncryptMethod is present, Tag 98
func (m Logon) HasEncryptMethod() bool {
	return m.Has(tag.EncryptMethod)
}

//HasHeartBtInt returns true if HeartBtInt is present, Tag 108
func (m Logon) HasHeartBtInt() bool {
	return m.Has(tag.HeartBtInt)
}

//HasResetSeqNumFlag returns true if ResetSeqNumFlag is present, Tag 141
func (m Logon) HasResetSeqNumFlag() bool {
	return m.Has(tag.ResetSeqNumFlag)
}

//HasMaxMessageSize returns true if MaxMessageSize is present, Tag 383
func (m Logon) HasMaxMessageSize() bool {
	return m.Has(tag.MaxMessageSize)
}

//HasNoMsgTypes returns true if NoMsgTypes is present, Tag 384
func (m Logon) HasNoMsgTypes() bool {
	return m.Has(tag.NoMsgTypes)
}

//HasTestMessageIndicator returns true if TestMessageIndicator is present, Tag 464
func (m Logon) HasTestMessageIndicator() bool {
	return m.Has(tag.TestMessageIndicator)
}

//HasUsername returns true if Username is present, Tag 553
func (m Logon) HasUsername() bool {
	return m.Has(tag.Username)
}

//HasPassword returns true if Password is present, Tag 554
func (m Logon) HasPassword() bool {
	return m.Has(tag.Password)
}

//NoMsgTypes is a repeating group element, Tag 384
type NoMsgTypes struct {
	*quickfix.Group
}

//SetRefMsgType sets RefMsgType, Tag 372
func (m NoMsgTypes) SetRefMsgType(v string) {
	m.Set(field.NewRefMsgType(v))
}

//SetMsgDirection sets MsgDirection, Tag 385
func (m NoMsgTypes) SetMsgDirection(v enum.MsgDirection) {
	m.Set(field.NewMsgDirection(v))
}

//GetRefMsgType gets RefMsgType, Tag 372
func (m NoMsgTypes) GetRefMsgType() (v string, err quickfix.MessageRejectError) {
	var f field.RefMsgTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMsgDirection gets MsgDirection, Tag 385
func (m NoMsgTypes) GetMsgDirection() (v enum.MsgDirection, err quickfix.MessageRejectError) {
	var f field.MsgDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRefMsgType returns true if RefMsgType is present, Tag 372
func (m NoMsgTypes) HasRefMsgType() bool {
	return m.Has(tag.RefMsgType)
}

//HasMsgDirection returns true if MsgDirection is present, Tag 385
func (m NoMsgTypes) HasMsgDirection() bool {
	return m.Has(tag.MsgDirection)
}

//NoMsgTypesRepeatingGroup is a repeating group, Tag 384
type NoMsgTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMsgTypesRepeatingGroup returns an initialized, NoMsgTypesRepeatingGroup
func NewNoMsgTypesRepeatingGroup() NoMsgTypesRepeatingGroup {
	return NoMsgTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMsgTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RefMsgType), quickfix.GroupElement(tag.MsgDirection)})}
}

//Add create and append a new NoMsgTypes to this group
func (m NoMsgTypesRepeatingGroup) Add() NoMsgTypes {
	g := m.RepeatingGroup.Add()
	return NoMsgTypes{g}
}

//Get returns the ith NoMsgTypes in the NoMsgTypesRepeatinGroup
func (m NoMsgTypesRepeatingGroup) Get(i int) NoMsgTypes {
	return NoMsgTypes{m.RepeatingGroup.Get(i)}
}
