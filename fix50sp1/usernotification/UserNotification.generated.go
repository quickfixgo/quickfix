package usernotification

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//UserNotification is the fix50sp1 UserNotification type, MsgType = CB
type UserNotification struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a UserNotification from a quickfix.Message instance
func FromMessage(m *quickfix.Message) UserNotification {
	return UserNotification{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m UserNotification) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a UserNotification initialized with the required fields for UserNotification
func New(userstatus field.UserStatusField) (m UserNotification) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("CB"))
	m.Set(userstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg UserNotification, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "CB", r
}

//SetText sets Text, Tag 58
func (m UserNotification) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m UserNotification) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m UserNotification) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetUsername sets Username, Tag 553
func (m UserNotification) SetUsername(v string) {
	m.Set(field.NewUsername(v))
}

//SetUserStatus sets UserStatus, Tag 926
func (m UserNotification) SetUserStatus(v enum.UserStatus) {
	m.Set(field.NewUserStatus(v))
}

//GetText gets Text, Tag 58
func (m UserNotification) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m UserNotification) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m UserNotification) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUsername gets Username, Tag 553
func (m UserNotification) GetUsername() (v string, err quickfix.MessageRejectError) {
	var f field.UsernameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUserStatus gets UserStatus, Tag 926
func (m UserNotification) GetUserStatus() (v enum.UserStatus, err quickfix.MessageRejectError) {
	var f field.UserStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m UserNotification) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m UserNotification) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m UserNotification) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasUsername returns true if Username is present, Tag 553
func (m UserNotification) HasUsername() bool {
	return m.Has(tag.Username)
}

//HasUserStatus returns true if UserStatus is present, Tag 926
func (m UserNotification) HasUserStatus() bool {
	return m.Has(tag.UserStatus)
}
