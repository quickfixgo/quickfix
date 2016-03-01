//Package ordercancelreject msg type = 9.
package ordercancelreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a OrderCancelReject FIX Message
type Message struct {
	FIXMsgType string `fix:"9"`
	Header     fix42.Header
	//OrderID is a required field for OrderCancelReject.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for OrderCancelReject.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a required field for OrderCancelReject.
	ClOrdID string `fix:"11"`
	//OrigClOrdID is a required field for OrderCancelReject.
	OrigClOrdID string `fix:"41"`
	//OrdStatus is a required field for OrderCancelReject.
	OrdStatus string `fix:"39"`
	//ClientID is a non-required field for OrderCancelReject.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderCancelReject.
	ExecBroker *string `fix:"76"`
	//ListID is a non-required field for OrderCancelReject.
	ListID *string `fix:"66"`
	//Account is a non-required field for OrderCancelReject.
	Account *string `fix:"1"`
	//TransactTime is a non-required field for OrderCancelReject.
	TransactTime *time.Time `fix:"60"`
	//CxlRejResponseTo is a required field for OrderCancelReject.
	CxlRejResponseTo string `fix:"434"`
	//CxlRejReason is a non-required field for OrderCancelReject.
	CxlRejReason *int `fix:"102"`
	//Text is a non-required field for OrderCancelReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReject.
	EncodedText *string `fix:"355"`
	Trailer     fix42.Trailer
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
	return enum.BeginStringFIX42, "9", r
}
