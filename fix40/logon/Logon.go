//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a Logon FIX Message
type Message struct {
	FIXMsgType string `fix:"A"`
	Header     fix40.Header
	//EncryptMethod is a required field for Logon.
	EncryptMethod int `fix:"98"`
	//HeartBtInt is a required field for Logon.
	HeartBtInt int `fix:"108"`
	//RawDataLength is a non-required field for Logon.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Logon.
	RawData *string `fix:"96"`
	Trailer fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEncryptMethod(v int) { m.EncryptMethod = v }
func (m *Message) SetHeartBtInt(v int)    { m.HeartBtInt = v }
func (m *Message) SetRawDataLength(v int) { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)    { m.RawData = &v }

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
	return enum.BeginStringFIX40, "A", r
}
