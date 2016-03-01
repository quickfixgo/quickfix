//Package ordercancelrequest msg type = F.
package ordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a OrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"F"`
	Header     fix40.Header
	//OrigClOrdID is a required field for OrderCancelRequest.
	OrigClOrdID string `fix:"41"`
	//OrderID is a non-required field for OrderCancelRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderCancelRequest.
	ClOrdID string `fix:"11"`
	//ListID is a non-required field for OrderCancelRequest.
	ListID *string `fix:"66"`
	//CxlType is a required field for OrderCancelRequest.
	CxlType string `fix:"125"`
	//ClientID is a non-required field for OrderCancelRequest.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderCancelRequest.
	ExecBroker *string `fix:"76"`
	//Symbol is a required field for OrderCancelRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for OrderCancelRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for OrderCancelRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for OrderCancelRequest.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for OrderCancelRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for OrderCancelRequest.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for OrderCancelRequest.
	Side string `fix:"54"`
	//OrderQty is a required field for OrderCancelRequest.
	OrderQty int `fix:"38"`
	//Text is a non-required field for OrderCancelRequest.
	Text    *string `fix:"58"`
	Trailer fix40.Trailer
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
	return enum.BeginStringFIX40, "F", r
}
