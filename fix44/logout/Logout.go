//Package logout msg type = 5.
package logout

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
)

//Message is a Logout FIX Message
type Message struct {
	FIXMsgType string `fix:"5"`
	Header     fix44.Header
	//Text is a non-required field for Logout.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Logout.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Logout.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetText(v string)        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int) { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string) { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "5", r
}
