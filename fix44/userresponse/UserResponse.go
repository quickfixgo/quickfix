package userresponse

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/tag"
)

//UserResponse is the fix44 UserResponse type, MsgType = BF
type UserResponse struct {
	fix44.Header
	quickfix.Body
	fix44.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a UserResponse from a quickfix.Message instance
func FromMessage(m quickfix.Message) UserResponse {
	return UserResponse{
		Header:      fix44.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix44.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m UserResponse) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a UserResponse initialized with the required fields for UserResponse
func New(userrequestid field.UserRequestIDField, username field.UsernameField) (m UserResponse) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BF"))
	m.Header.Set(field.NewBeginString("FIX.4.4"))
	m.Set(userrequestid)
	m.Set(username)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BF", r
}

//SetUsername sets Username, Tag 553
func (m UserResponse) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

//SetUserRequestID sets UserRequestID, Tag 923
func (m UserResponse) SetUserRequestID(v string) {
	m.Set(field.NewUserRequestID(v))
}

//SetUserStatus sets UserStatus, Tag 926
func (m UserResponse) SetUserStatus(v int) {
	m.Set(field.NewUserStatus(v))
}

//SetUserStatusText sets UserStatusText, Tag 927
func (m UserResponse) SetUserStatusText(v string) {
	m.Set(field.NewUserStatusText(v))
}

//GetUsername gets Username, Tag 553
func (m UserResponse) GetUsername() (f field.UsernameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUserRequestID gets UserRequestID, Tag 923
func (m UserResponse) GetUserRequestID() (f field.UserRequestIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUserStatus gets UserStatus, Tag 926
func (m UserResponse) GetUserStatus() (f field.UserStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUserStatusText gets UserStatusText, Tag 927
func (m UserResponse) GetUserStatusText() (f field.UserStatusTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasUsername returns true if Username is present, Tag 553
func (m UserResponse) HasUsername() bool {
	return m.Has(tag.Username)
}

//HasUserRequestID returns true if UserRequestID is present, Tag 923
func (m UserResponse) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

//HasUserStatus returns true if UserStatus is present, Tag 926
func (m UserResponse) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

//HasUserStatusText returns true if UserStatusText is present, Tag 927
func (m UserResponse) HasUserStatusText() bool {
	return m.Has(tag.UserStatusText)
}
