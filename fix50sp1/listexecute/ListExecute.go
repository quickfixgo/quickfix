//Package listexecute msg type = L.
package listexecute

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ListExecute FIX Message
type Message struct {
	FIXMsgType string `fix:"L"`
	fixt11.Header
	//ListID is a required field for ListExecute.
	ListID string `fix:"66"`
	//ClientBidID is a non-required field for ListExecute.
	ClientBidID *string `fix:"391"`
	//BidID is a non-required field for ListExecute.
	BidID *string `fix:"390"`
	//TransactTime is a required field for ListExecute.
	TransactTime time.Time `fix:"60"`
	//Text is a non-required field for ListExecute.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ListExecute.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ListExecute.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized ListExecute instance
func New(listid string, transacttime time.Time) *Message {
	var m Message
	m.SetListID(listid)
	m.SetTransactTime(transacttime)
	return &m
}

func (m *Message) SetListID(v string)          { m.ListID = v }
func (m *Message) SetClientBidID(v string)     { m.ClientBidID = &v }
func (m *Message) SetBidID(v string)           { m.BidID = &v }
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = v }
func (m *Message) SetText(v string)            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)     { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP1, "L", r
}
