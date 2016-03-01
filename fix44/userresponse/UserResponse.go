//Package userresponse msg type = BF.
package userresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//Message is a UserResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"BF"`
	Header     fix44.Header
	//UserRequestID is a required field for UserResponse.
	UserRequestID string `fix:"923"`
	//Username is a required field for UserResponse.
	Username string `fix:"553"`
	//UserStatus is a non-required field for UserResponse.
	UserStatus *int `fix:"926"`
	//UserStatusText is a non-required field for UserResponse.
	UserStatusText *string `fix:"927"`
	Trailer        fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX44, "BF", r
}
