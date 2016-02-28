//Package userrequest msg type = BE.
package userrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//Message is a UserRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BE"`
	Header     fix44.Header
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
	Trailer fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetUserRequestID(v string) { m.UserRequestID = v }
func (m *Message) SetUserRequestType(v int)  { m.UserRequestType = v }
func (m *Message) SetUsername(v string)      { m.Username = v }
func (m *Message) SetPassword(v string)      { m.Password = &v }
func (m *Message) SetNewPassword(v string)   { m.NewPassword = &v }
func (m *Message) SetRawDataLength(v int)    { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)       { m.RawData = &v }

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
	return enum.BeginStringFIX44, "BE", r
}
