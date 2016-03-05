//Package orderstatusrequest msg type = H.
package orderstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a OrderStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"H"`
	fixt11.Header
	//OrderID is a non-required field for OrderStatusRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for OrderStatusRequest.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderStatusRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for OrderStatusRequest.
	ClOrdLinkID *string `fix:"583"`
	//Parties Component
	parties.Parties
	//OrdStatusReqID is a non-required field for OrderStatusRequest.
	OrdStatusReqID *string `fix:"790"`
	//Account is a non-required field for OrderStatusRequest.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for OrderStatusRequest.
	AcctIDSource *int `fix:"660"`
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//Side is a required field for OrderStatusRequest.
	Side string `fix:"54"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *Message) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdLinkID(v string)      { m.ClOrdLinkID = &v }
func (m *Message) SetOrdStatusReqID(v string)   { m.OrdStatusReqID = &v }
func (m *Message) SetAccount(v string)          { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)        { m.AcctIDSource = &v }
func (m *Message) SetSide(v string)             { m.Side = v }

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
	return enum.ApplVerID_FIX50SP1, "H", r
}
