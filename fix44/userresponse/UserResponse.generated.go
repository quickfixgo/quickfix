package userresponse

import (
	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/enum"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fix44"
	"github.com/alpacahq/quickfix/tag"
)

// UserResponse is the fix44 UserResponse type, MsgType = BF
type UserResponse struct {
	fix44.Header
	*quickfix.Body
	fix44.Trailer
	Message *quickfix.Message
}

// FromMessage creates a UserResponse from a quickfix.Message instance
func FromMessage(m *quickfix.Message) UserResponse {
	return UserResponse{
		Header:  fix44.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix44.Trailer{&m.Trailer},
		Message: m,
	}
}

// ToMessage returns a quickfix.Message instance
func (m UserResponse) ToMessage() *quickfix.Message {
	return m.Message
}

// New returns a UserResponse initialized with the required fields for UserResponse
func New(userrequestid field.UserRequestIDField, username field.UsernameField) (m UserResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix44.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BF"))
	m.Set(userrequestid)
	m.Set(username)

	return
}

// A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg UserResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

// Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.4", "BF", r
}

// SetUsername sets Username, Tag 553
func (m UserResponse) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

// SetUserRequestID sets UserRequestID, Tag 923
func (m UserResponse) SetUserRequestID(v string) {
	m.Set(field.NewUserRequestID(v))
}

// SetUserStatus sets UserStatus, Tag 926
func (m UserResponse) SetUserStatus(v enum.UserStatus) {
	m.Set(field.NewUserStatus(v))
}

// SetUserStatusText sets UserStatusText, Tag 927
func (m UserResponse) SetUserStatusText(v string) {
	m.Set(field.NewUserStatusText(v))
}

// GetUsername gets Username, Tag 553
func (m UserResponse) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserRequestID gets UserRequestID, Tag 923
func (m UserResponse) GetUserRequestID() (v string, err quickfix.MessageRejectError) {
	var f field.UserRequestIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatus gets UserStatus, Tag 926
func (m UserResponse) GetUserStatus() (v enum.UserStatus, err quickfix.MessageRejectError) {
	var f field.UserStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// GetUserStatusText gets UserStatusText, Tag 927
func (m UserResponse) GetUserStatusText() (v string, err quickfix.MessageRejectError) {
	var f field.UserStatusTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

// HasUsername returns true if Username is present, Tag 553
func (m UserResponse) HasUsername() bool {
	return m.Has(tag.Username)
}

// HasUserRequestID returns true if UserRequestID is present, Tag 923
func (m UserResponse) HasUserRequestID() bool {
	return m.Has(tag.UserRequestID)
}

// HasUserStatus returns true if UserStatus is present, Tag 926
func (m UserResponse) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}

// HasUserStatusText returns true if UserStatusText is present, Tag 927
func (m UserResponse) HasUserStatusText() bool {
	return m.Has(tag.UserStatusText)
}
