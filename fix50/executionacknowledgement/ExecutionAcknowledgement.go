//Package executionacknowledgement msg type = BN.
package executionacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ExecutionAcknowledgement FIX Message
type Message struct {
	FIXMsgType string `fix:"BN"`
	fixt11.Header
	//OrderID is a required field for ExecutionAcknowledgement.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for ExecutionAcknowledgement.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for ExecutionAcknowledgement.
	ClOrdID *string `fix:"11"`
	//ExecAckStatus is a required field for ExecutionAcknowledgement.
	ExecAckStatus string `fix:"1036"`
	//ExecID is a required field for ExecutionAcknowledgement.
	ExecID string `fix:"17"`
	//DKReason is a non-required field for ExecutionAcknowledgement.
	DKReason *string `fix:"127"`
	//Instrument Component
	instrument.Instrument
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//Side is a required field for ExecutionAcknowledgement.
	Side string `fix:"54"`
	//OrderQtyData Component
	orderqtydata.OrderQtyData
	//LastQty is a non-required field for ExecutionAcknowledgement.
	LastQty *float64 `fix:"32"`
	//LastPx is a non-required field for ExecutionAcknowledgement.
	LastPx *float64 `fix:"31"`
	//PriceType is a non-required field for ExecutionAcknowledgement.
	PriceType *int `fix:"423"`
	//LastParPx is a non-required field for ExecutionAcknowledgement.
	LastParPx *float64 `fix:"669"`
	//CumQty is a non-required field for ExecutionAcknowledgement.
	CumQty *float64 `fix:"14"`
	//AvgPx is a non-required field for ExecutionAcknowledgement.
	AvgPx *float64 `fix:"6"`
	//Text is a non-required field for ExecutionAcknowledgement.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ExecutionAcknowledgement.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ExecutionAcknowledgement.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = v }
func (m *Message) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *Message) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *Message) SetExecAckStatus(v string)    { m.ExecAckStatus = v }
func (m *Message) SetExecID(v string)           { m.ExecID = v }
func (m *Message) SetDKReason(v string)         { m.DKReason = &v }
func (m *Message) SetSide(v string)             { m.Side = v }
func (m *Message) SetLastQty(v float64)         { m.LastQty = &v }
func (m *Message) SetLastPx(v float64)          { m.LastPx = &v }
func (m *Message) SetPriceType(v int)           { m.PriceType = &v }
func (m *Message) SetLastParPx(v float64)       { m.LastParPx = &v }
func (m *Message) SetCumQty(v float64)          { m.CumQty = &v }
func (m *Message) SetAvgPx(v float64)           { m.AvgPx = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }

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
	return enum.BeginStringFIX50, "BN", r
}
