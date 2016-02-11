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
	Header     fixt11.Header
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
	//OrdListStatGrp Component
	OrdListStatGrp ordliststatgrp.Component
	//ContingencyType is a non-required field for ListStatus.
	ContingencyType *int `fix:"1385"`
	//ListRejectReason is a non-required field for ListStatus.
	ListRejectReason *int `fix:"1386"`
	Trailer          fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "N", r
}
