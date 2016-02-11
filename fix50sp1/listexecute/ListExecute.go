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
	Header     fixt11.Header
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
	return enum.ApplVerID_FIX50SP1, "L", r
}
