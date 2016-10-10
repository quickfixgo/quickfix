package userrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//UserRequest is the fix50 UserRequest type, MsgType = BE
type UserRequest struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a UserRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) UserRequest {
	return UserRequest{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m UserRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a UserRequest initialized with the required fields for UserRequest
func New(userrequestid field.UserRequestIDField, userrequesttype field.UserRequestTypeField, username field.UsernameField) (m UserRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BE"))
	m.Set(userrequestid)
	m.Set(userrequesttype)
	m.Set(username)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg UserRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "BE", r
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m UserRequest) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m UserRequest) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//SetUsername sets Username, Tag 553
func (m UserRequest) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

//SetPassword sets Password, Tag 554
func (m UserRequest) SetPassword(v string) {
	m.Set(field.NewPassword(v))
}

//SetUserRequestID sets UserRequestID, Tag 923
func (m UserRequest) SetUserRequestID(v string) {
	m.Set(field.NewUserRequestID(v))
}

//SetUserRequestType sets UserRequestType, Tag 924
func (m UserRequest) SetUserRequestType(v enum.UserRequestType) {
	m.Set(field.NewUserRequestType(v))
}

//SetNewPassword sets NewPassword, Tag 925
func (m UserRequest) SetNewPassword(v string) {
	m.Set(field.NewNewPassword(v))
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m UserRequest) GetRawDataLength() (v int, err quickfix.MessageRejectError) {
	var f field.RawDataLengthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRawData gets RawData, Tag 96
func (m UserRequest) GetRawData() (v string, err quickfix.MessageRejectError) {
	var f field.RawDataField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUsername gets Username, Tag 553
func (m UserRequest) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPassword gets Password, Tag 554
func (m UserRequest) GetPassword() (v string, err quickfix.MessageRejectError) {
	var f field.PasswordField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUserRequestID gets UserRequestID, Tag 923
func (m UserRequest) GetUserRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.UserRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUserRequestType gets UserRequestType, Tag 924
func (m UserRequest) GetUserRequestType() (v enum.UserRequestType, err quickfix.MessageRejectError) {
	var f field.UserRequestTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNewPassword gets NewPassword, Tag 925
func (m UserRequest) GetNewPassword() (v string, err quickfix.MessageRejectError) {
	var f field.NewPasswordField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m UserRequest) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m UserRequest) HasRawData() bool {
	return m.Has(tag.RawData)
}

//HasUsername returns true if Username is present, Tag 553
func (m UserRequest) HasUsername() bool {
	return m.Has(tag.Username)
}

//HasPassword returns true if Password is present, Tag 554
func (m UserRequest) HasPassword() bool {
	return m.Has(tag.Password)
}

//HasUserRequestID returns true if UserRequestID is present, Tag 923
func (m UserRequest) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

//HasUserRequestType returns true if UserRequestType is present, Tag 924
func (m UserRequest) HasUserRequestType() bool {
	return m.Has(tag.UserRequestType)
}

//HasNewPassword returns true if NewPassword is present, Tag 925
func (m UserRequest) HasNewPassword() bool {
	return m.Has(tag.NewPassword)
}
