//Package requestforpositionsack msg type = AO.
package requestforpositionsack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoLegs is a repeating group in RequestForPositionsAck
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in RequestForPositionsAck
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Message is a RequestForPositionsAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AO"`
	Header     fix44.Header
	//PosMaintRptID is a required field for RequestForPositionsAck.
	PosMaintRptID string `fix:"721"`
	//PosReqID is a non-required field for RequestForPositionsAck.
	PosReqID *string `fix:"710"`
	//TotalNumPosReports is a non-required field for RequestForPositionsAck.
	TotalNumPosReports *int `fix:"727"`
	//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
	UnsolicitedIndicator *bool `fix:"325"`
	//PosReqResult is a required field for RequestForPositionsAck.
	PosReqResult int `fix:"728"`
	//PosReqStatus is a required field for RequestForPositionsAck.
	PosReqStatus int `fix:"729"`
	//Parties Component
	Parties parties.Component
	//Account is a required field for RequestForPositionsAck.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for RequestForPositionsAck.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for RequestForPositionsAck.
	AccountType int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for RequestForPositionsAck.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for RequestForPositionsAck.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for RequestForPositionsAck.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//ResponseTransportType is a non-required field for RequestForPositionsAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for RequestForPositionsAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for RequestForPositionsAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for RequestForPositionsAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for RequestForPositionsAck.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "AO", r
}
