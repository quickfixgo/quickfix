package userrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//UserRequest is the fix50 UserRequest type, MsgType = BE
type UserRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a UserRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) UserRequest {
	return UserRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m UserRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a UserRequest initialized with the required fields for UserRequest
func New(userrequestid field.UserRequestIDField, userrequesttype field.UserRequestTypeField, username field.UsernameField) (m UserRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

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
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m UserRequest) SetUserRequestType(v int) {
	m.Set(field.NewUserRequestType(v))
}

//SetNewPassword sets NewPassword, Tag 925
func (m UserRequest) SetNewPassword(v string) {
	m.Set(field.NewNewPassword(v))
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m UserRequest) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m UserRequest) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUsername gets Username, Tag 553
func (m UserRequest) GetUsername() (f field.UsernameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPassword gets Password, Tag 554
func (m UserRequest) GetPassword() (f field.PasswordField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUserRequestID gets UserRequestID, Tag 923
func (m UserRequest) GetUserRequestID() (f field.UserRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUserRequestType gets UserRequestType, Tag 924
func (m UserRequest) GetUserRequestType() (f field.UserRequestTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNewPassword gets NewPassword, Tag 925
func (m UserRequest) GetNewPassword() (f field.NewPasswordField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
