//Package reject msg type = 3.
package reject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a Reject FIX Message
type Message struct {
	FIXMsgType string `fix:"3"`
	Header     fix40.Header
	//RefSeqNum is a required field for Reject.
	RefSeqNum int `fix:"45"`
	//Text is a non-required field for Reject.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetRefSeqNum(v int) { m.RefSeqNum = v }
func (m *Message) SetText(v string)   { m.Text = &v }

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
	return enum.BeginStringFIX40, "3", r
}
