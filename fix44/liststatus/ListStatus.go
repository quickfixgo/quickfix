//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"time"
)

//NoOrders is a repeating group in ListStatus
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//CumQty is a required field for NoOrders.
	CumQty float64 `fix:"14"`
	//OrdStatus is a required field for NoOrders.
	OrdStatus string `fix:"39"`
	//WorkingIndicator is a non-required field for NoOrders.
	WorkingIndicator *bool `fix:"636"`
	//LeavesQty is a required field for NoOrders.
	LeavesQty float64 `fix:"151"`
	//CxlQty is a required field for NoOrders.
	CxlQty float64 `fix:"84"`
	//AvgPx is a required field for NoOrders.
	AvgPx float64 `fix:"6"`
	//OrdRejReason is a non-required field for NoOrders.
	OrdRejReason *int `fix:"103"`
	//Text is a non-required field for NoOrders.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoOrders.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoOrders.
	EncodedText *string `fix:"355"`
}

//Message is a ListStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"N"`
	Header     fix44.Header
	//ListID is a required field for ListStatus.
	ListID string `fix:"66"`
	//ListStatusType is a required field for ListStatus.
	ListStatusType int `fix:"429"`
	//NoRpts is a required field for ListStatus.
	NoRpts int `fix:"82"`
	//ListOrderStatus is a required field for ListStatus.
	ListOrderStatus int `fix:"431"`
	//RptSeq is a required field for ListStatus.
	RptSeq int `fix:"83"`
	//ListStatusText is a non-required field for ListStatus.
	ListStatusText *string `fix:"444"`
	//EncodedListStatusTextLen is a non-required field for ListStatus.
	EncodedListStatusTextLen *int `fix:"445"`
	//EncodedListStatusText is a non-required field for ListStatus.
	EncodedListStatusText *string `fix:"446"`
	//TransactTime is a non-required field for ListStatus.
	TransactTime *time.Time `fix:"60"`
	//TotNoOrders is a required field for ListStatus.
	TotNoOrders int `fix:"68"`
	//LastFragment is a non-required field for ListStatus.
	LastFragment *bool `fix:"893"`
	//NoOrders is a required field for ListStatus.
	NoOrders []NoOrders `fix:"73"`
	Trailer  fix44.Trailer
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
	return enum.BeginStringFIX44, "N", r
}
