//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//NoOrders is a repeating group in ListStatus
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//CumQty is a required field for NoOrders.
	CumQty int `fix:"14"`
	//LeavesQty is a required field for NoOrders.
	LeavesQty int `fix:"151"`
	//CxlQty is a required field for NoOrders.
	CxlQty int `fix:"84"`
	//AvgPx is a required field for NoOrders.
	AvgPx float64 `fix:"6"`
}

//Message is a ListStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"N"`
	Header     fix41.Header
	//ListID is a required field for ListStatus.
	ListID string `fix:"66"`
	//WaveNo is a non-required field for ListStatus.
	WaveNo *string `fix:"105"`
	//NoRpts is a required field for ListStatus.
	NoRpts int `fix:"82"`
	//RptSeq is a required field for ListStatus.
	RptSeq int `fix:"83"`
	//NoOrders is a required field for ListStatus.
	NoOrders []NoOrders `fix:"73"`
	Trailer  fix41.Trailer
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
	return enum.BeginStringFIX41, "N", r
}
