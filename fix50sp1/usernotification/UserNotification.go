//Package usernotification msg type = CB.
package usernotification

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/usernamegrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a UserNotification FIX Message
type Message struct {
	FIXMsgType string `fix:"CB"`
	Header     fixt11.Header
	//UsernameGrp Component
	UsernameGrp usernamegrp.Component
	//UserStatus is a required field for UserNotification.
	UserStatus int `fix:"926"`
	//Text is a non-required field for UserNotification.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for UserNotification.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for UserNotification.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "CB", r
}
