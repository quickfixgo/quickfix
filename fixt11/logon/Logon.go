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
	Header     fixt11.Header
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
	//MsgTypeGrp Component
	MsgTypeGrp msgtypegrp.Component
	Trailer    fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIXT11, "A", r
}
