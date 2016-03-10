//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
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

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetCumQty(v float64)          { m.CumQty = v }
func (m *NoOrders) SetOrdStatus(v string)        { m.OrdStatus = v }
func (m *NoOrders) SetWorkingIndicator(v bool)   { m.WorkingIndicator = &v }
func (m *NoOrders) SetLeavesQty(v float64)       { m.LeavesQty = v }
func (m *NoOrders) SetCxlQty(v float64)          { m.CxlQty = v }
func (m *NoOrders) SetAvgPx(v float64)           { m.AvgPx = v }
func (m *NoOrders) SetOrdRejReason(v int)        { m.OrdRejReason = &v }
func (m *NoOrders) SetText(v string)             { m.Text = &v }
func (m *NoOrders) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *NoOrders) SetEncodedText(v string)      { m.EncodedText = &v }

//Message is a ListStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"N"`
	fix43.Header
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
	//NoOrders is a required field for ListStatus.
	NoOrders []NoOrders `fix:"73"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)                { m.ListID = v }
func (m *Message) SetListStatusType(v int)           { m.ListStatusType = v }
func (m *Message) SetNoRpts(v int)                   { m.NoRpts = v }
func (m *Message) SetListOrderStatus(v int)          { m.ListOrderStatus = v }
func (m *Message) SetRptSeq(v int)                   { m.RptSeq = v }
func (m *Message) SetListStatusText(v string)        { m.ListStatusText = &v }
func (m *Message) SetEncodedListStatusTextLen(v int) { m.EncodedListStatusTextLen = &v }
func (m *Message) SetEncodedListStatusText(v string) { m.EncodedListStatusText = &v }
func (m *Message) SetTransactTime(v time.Time)       { m.TransactTime = &v }
func (m *Message) SetTotNoOrders(v int)              { m.TotNoOrders = v }
func (m *Message) SetNoOrders(v []NoOrders)          { m.NoOrders = v }

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
	return enum.BeginStringFIX43, "N", r
}
