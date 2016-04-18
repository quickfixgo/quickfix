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
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//NoUnderlyings is a repeating group in RequestForPositionsAck
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//Message is a RequestForPositionsAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AO"`
	fix44.Header
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
	//Parties is a required component for RequestForPositionsAck.
	parties.Parties
	//Account is a required field for RequestForPositionsAck.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for RequestForPositionsAck.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for RequestForPositionsAck.
	AccountType int `fix:"581"`
	//Instrument is a non-required component for RequestForPositionsAck.
	Instrument *instrument.Instrument
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
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized RequestForPositionsAck instance
func New(posmaintrptid string, posreqresult int, posreqstatus int, parties parties.Parties, account string, accounttype int) *Message {
	var m Message
	m.SetPosMaintRptID(posmaintrptid)
	m.SetPosReqResult(posreqresult)
	m.SetPosReqStatus(posreqstatus)
	m.SetParties(parties)
	m.SetAccount(account)
	m.SetAccountType(accounttype)
	return &m
}

func (m *Message) SetPosMaintRptID(v string)             { m.PosMaintRptID = v }
func (m *Message) SetPosReqID(v string)                  { m.PosReqID = &v }
func (m *Message) SetTotalNumPosReports(v int)           { m.TotalNumPosReports = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)        { m.UnsolicitedIndicator = &v }
func (m *Message) SetPosReqResult(v int)                 { m.PosReqResult = v }
func (m *Message) SetPosReqStatus(v int)                 { m.PosReqStatus = v }
func (m *Message) SetParties(v parties.Parties)          { m.Parties = v }
func (m *Message) SetAccount(v string)                   { m.Account = v }
func (m *Message) SetAcctIDSource(v int)                 { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                  { m.AccountType = v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetCurrency(v string)                  { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                  { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)    { m.NoUnderlyings = v }
func (m *Message) SetResponseTransportType(v int)        { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)       { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                      { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)               { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)               { m.EncodedText = &v }

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
