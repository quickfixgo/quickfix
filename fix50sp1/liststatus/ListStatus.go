//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/ordliststatgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ListStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"N"`
	fixt11.Header
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
	//OrdListStatGrp is a required component for ListStatus.
	ordliststatgrp.OrdListStatGrp
	//ContingencyType is a non-required field for ListStatus.
	ContingencyType *int `fix:"1385"`
	//ListRejectReason is a non-required field for ListStatus.
	ListRejectReason *int `fix:"1386"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)                                { m.ListID = v }
func (m *Message) SetListStatusType(v int)                           { m.ListStatusType = v }
func (m *Message) SetNoRpts(v int)                                   { m.NoRpts = v }
func (m *Message) SetListOrderStatus(v int)                          { m.ListOrderStatus = v }
func (m *Message) SetRptSeq(v int)                                   { m.RptSeq = v }
func (m *Message) SetListStatusText(v string)                        { m.ListStatusText = &v }
func (m *Message) SetEncodedListStatusTextLen(v int)                 { m.EncodedListStatusTextLen = &v }
func (m *Message) SetEncodedListStatusText(v string)                 { m.EncodedListStatusText = &v }
func (m *Message) SetTransactTime(v time.Time)                       { m.TransactTime = &v }
func (m *Message) SetTotNoOrders(v int)                              { m.TotNoOrders = v }
func (m *Message) SetLastFragment(v bool)                            { m.LastFragment = &v }
func (m *Message) SetOrdListStatGrp(v ordliststatgrp.OrdListStatGrp) { m.OrdListStatGrp = v }
func (m *Message) SetContingencyType(v int)                          { m.ContingencyType = &v }
func (m *Message) SetListRejectReason(v int)                         { m.ListRejectReason = &v }

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
	return enum.ApplVerID_FIX50SP1, "N", r
}
