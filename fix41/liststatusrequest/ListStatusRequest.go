//Package liststatusrequest msg type = M.
package liststatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a ListStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"M"`
	fix41.Header
	//ListID is a required field for ListStatusRequest.
	ListID string `fix:"66"`
	//WaveNo is a non-required field for ListStatusRequest.
	WaveNo *string `fix:"105"`
	//Text is a non-required field for ListStatusRequest.
	Text *string `fix:"58"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string) { m.ListID = v }
func (m *Message) SetWaveNo(v string) { m.WaveNo = &v }
func (m *Message) SetText(v string)   { m.Text = &v }

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
	return enum.BeginStringFIX41, "M", r
}
