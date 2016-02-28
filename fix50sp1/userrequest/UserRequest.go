//Package userrequest msg type = BE.
package userrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a UserRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BE"`
	Header     fixt11.Header
	//UserRequestID is a required field for UserRequest.
	UserRequestID string `fix:"923"`
	//UserRequestType is a required field for UserRequest.
	UserRequestType int `fix:"924"`
	//Username is a required field for UserRequest.
	Username string `fix:"553"`
	//Password is a non-required field for UserRequest.
	Password *string `fix:"554"`
	//NewPassword is a non-required field for UserRequest.
	NewPassword *string `fix:"925"`
	//RawDataLength is a non-required field for UserRequest.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for UserRequest.
	RawData *string `fix:"96"`
	//EncryptedPasswordMethod is a non-required field for UserRequest.
	EncryptedPasswordMethod *int `fix:"1400"`
	//EncryptedPasswordLen is a non-required field for UserRequest.
	EncryptedPasswordLen *int `fix:"1401"`
	//EncryptedPassword is a non-required field for UserRequest.
	EncryptedPassword *string `fix:"1402"`
	//EncryptedNewPasswordLen is a non-required field for UserRequest.
	EncryptedNewPasswordLen *int `fix:"1403"`
	//EncryptedNewPassword is a non-required field for UserRequest.
	EncryptedNewPassword *string `fix:"1404"`
	Trailer              fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetUserRequestID(v string)        { m.UserRequestID = v }
func (m *Message) SetUserRequestType(v int)         { m.UserRequestType = v }
func (m *Message) SetUsername(v string)             { m.Username = v }
func (m *Message) SetPassword(v string)             { m.Password = &v }
func (m *Message) SetNewPassword(v string)          { m.NewPassword = &v }
func (m *Message) SetRawDataLength(v int)           { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)              { m.RawData = &v }
func (m *Message) SetEncryptedPasswordMethod(v int) { m.EncryptedPasswordMethod = &v }
func (m *Message) SetEncryptedPasswordLen(v int)    { m.EncryptedPasswordLen = &v }
func (m *Message) SetEncryptedPassword(v string)    { m.EncryptedPassword = &v }
func (m *Message) SetEncryptedNewPasswordLen(v int) { m.EncryptedNewPasswordLen = &v }
func (m *Message) SetEncryptedNewPassword(v string) { m.EncryptedNewPassword = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BE", r
}
