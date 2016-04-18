//Package logon msg type = A.
package logon

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/fixt11/msgtypegrp"
)

//Message is a Logon FIX Message
type Message struct {
	FIXMsgType string `fix:"A"`
	fixt11.Header
	//EncryptMethod is a required field for Logon.
	EncryptMethod int `fix:"98"`
	//HeartBtInt is a required field for Logon.
	HeartBtInt int `fix:"108"`
	//RawDataLength is a non-required field for Logon.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Logon.
	RawData *string `fix:"96"`
	//ResetSeqNumFlag is a non-required field for Logon.
	ResetSeqNumFlag *bool `fix:"141"`
	//NextExpectedMsgSeqNum is a non-required field for Logon.
	NextExpectedMsgSeqNum *int `fix:"789"`
	//MaxMessageSize is a non-required field for Logon.
	MaxMessageSize *int `fix:"383"`
	//TestMessageIndicator is a non-required field for Logon.
	TestMessageIndicator *bool `fix:"464"`
	//Username is a non-required field for Logon.
	Username *string `fix:"553"`
	//Password is a non-required field for Logon.
	Password *string `fix:"554"`
	//DefaultApplVerID is a required field for Logon.
	DefaultApplVerID string `fix:"1137"`
	//MsgTypeGrp is a non-required component for Logon.
	MsgTypeGrp *msgtypegrp.MsgTypeGrp
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized Logon instance
func New(encryptmethod int, heartbtint int, defaultapplverid string) *Message {
	var m Message
	m.SetEncryptMethod(encryptmethod)
	m.SetHeartBtInt(heartbtint)
	m.SetDefaultApplVerID(defaultapplverid)
	return &m
}

func (m *Message) SetEncryptMethod(v int)                { m.EncryptMethod = v }
func (m *Message) SetHeartBtInt(v int)                   { m.HeartBtInt = v }
func (m *Message) SetRawDataLength(v int)                { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)                   { m.RawData = &v }
func (m *Message) SetResetSeqNumFlag(v bool)             { m.ResetSeqNumFlag = &v }
func (m *Message) SetNextExpectedMsgSeqNum(v int)        { m.NextExpectedMsgSeqNum = &v }
func (m *Message) SetMaxMessageSize(v int)               { m.MaxMessageSize = &v }
func (m *Message) SetTestMessageIndicator(v bool)        { m.TestMessageIndicator = &v }
func (m *Message) SetUsername(v string)                  { m.Username = &v }
func (m *Message) SetPassword(v string)                  { m.Password = &v }
func (m *Message) SetDefaultApplVerID(v string)          { m.DefaultApplVerID = v }
func (m *Message) SetMsgTypeGrp(v msgtypegrp.MsgTypeGrp) { m.MsgTypeGrp = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIXT11, "A", r
}
