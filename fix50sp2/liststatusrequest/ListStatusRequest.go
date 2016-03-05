//Package liststatusrequest msg type = M.
package liststatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ListStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"M"`
	fixt11.Header
	//ListID is a required field for ListStatusRequest.
	ListID string `fix:"66"`
	//Text is a non-required field for ListStatusRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ListStatusRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ListStatusRequest.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)      { m.ListID = v }
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
	return enum.ApplVerID_FIX50SP2, "M", r
}
