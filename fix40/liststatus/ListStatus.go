//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//NoOrders is a repeating group in ListStatus
type NoOrders struct {
	//ClOrdID is a required field for NoOrders.
	ClOrdID string `fix:"11"`
	//CumQty is a required field for NoOrders.
	CumQty int `fix:"14"`
	//CxlQty is a required field for NoOrders.
	CxlQty int `fix:"84"`
	//AvgPx is a required field for NoOrders.
	AvgPx float64 `fix:"6"`
}

func (m *NoOrders) SetClOrdID(v string) { m.ClOrdID = v }
func (m *NoOrders) SetCumQty(v int)     { m.CumQty = v }
func (m *NoOrders) SetCxlQty(v int)     { m.CxlQty = v }
func (m *NoOrders) SetAvgPx(v float64)  { m.AvgPx = v }

//Message is a ListStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"N"`
	fix40.Header
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
	fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)       { m.ListID = v }
func (m *Message) SetWaveNo(v string)       { m.WaveNo = &v }
func (m *Message) SetNoRpts(v int)          { m.NoRpts = v }
func (m *Message) SetRptSeq(v int)          { m.RptSeq = v }
func (m *Message) SetNoOrders(v []NoOrders) { m.NoOrders = v }

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
	return enum.BeginStringFIX40, "N", r
}
