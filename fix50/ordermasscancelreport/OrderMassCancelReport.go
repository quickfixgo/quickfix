//Package ordermasscancelreport msg type = r.
package ordermasscancelreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/affectedordgrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a OrderMassCancelReport FIX Message
type Message struct {
	FIXMsgType string `fix:"r"`
	fixt11.Header
	//ClOrdID is a non-required field for OrderMassCancelReport.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
	SecondaryClOrdID *string `fix:"526"`
	//OrderID is a required field for OrderMassCancelReport.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for OrderMassCancelReport.
	SecondaryOrderID *string `fix:"198"`
	//MassCancelRequestType is a required field for OrderMassCancelReport.
	MassCancelRequestType string `fix:"530"`
	//MassCancelResponse is a required field for OrderMassCancelReport.
	MassCancelResponse string `fix:"531"`
	//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
	MassCancelRejectReason *int `fix:"532"`
	//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
	TotalAffectedOrders *int `fix:"533"`
	//AffectedOrdGrp is a non-required component for OrderMassCancelReport.
	AffectedOrdGrp *affectedordgrp.AffectedOrdGrp
	//TradingSessionID is a non-required field for OrderMassCancelReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassCancelReport.
	TradingSessionSubID *string `fix:"625"`
	//Instrument is a non-required component for OrderMassCancelReport.
	Instrument *instrument.Instrument
	//UnderlyingInstrument is a non-required component for OrderMassCancelReport.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//Side is a non-required field for OrderMassCancelReport.
	Side *string `fix:"54"`
	//TransactTime is a non-required field for OrderMassCancelReport.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for OrderMassCancelReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassCancelReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassCancelReport.
	EncodedText *string `fix:"355"`
	//Parties is a non-required component for OrderMassCancelReport.
	Parties *parties.Parties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized OrderMassCancelReport instance
func New(orderid string, masscancelrequesttype string, masscancelresponse string) *Message {
	var m Message
	m.SetOrderID(orderid)
	m.SetMassCancelRequestType(masscancelrequesttype)
	m.SetMassCancelResponse(masscancelresponse)
	return &m
}

func (m *Message) SetClOrdID(v string)                               { m.ClOrdID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                      { m.SecondaryClOrdID = &v }
func (m *Message) SetOrderID(v string)                               { m.OrderID = v }
func (m *Message) SetSecondaryOrderID(v string)                      { m.SecondaryOrderID = &v }
func (m *Message) SetMassCancelRequestType(v string)                 { m.MassCancelRequestType = v }
func (m *Message) SetMassCancelResponse(v string)                    { m.MassCancelResponse = v }
func (m *Message) SetMassCancelRejectReason(v int)                   { m.MassCancelRejectReason = &v }
func (m *Message) SetTotalAffectedOrders(v int)                      { m.TotalAffectedOrders = &v }
func (m *Message) SetAffectedOrdGrp(v affectedordgrp.AffectedOrdGrp) { m.AffectedOrdGrp = &v }
func (m *Message) SetTradingSessionID(v string)                      { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                   { m.TradingSessionSubID = &v }
func (m *Message) SetInstrument(v instrument.Instrument)             { m.Instrument = &v }
func (m *Message) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *Message) SetSide(v string)             { m.Side = &v }
func (m *Message) SetTransactTime(v time.Time)  { m.TransactTime = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }
func (m *Message) SetParties(v parties.Parties) { m.Parties = &v }

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
	return enum.BeginStringFIX50, "r", r
}
