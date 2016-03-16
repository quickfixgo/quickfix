//Package requestforpositions msg type = AN.
package requestforpositions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoLegs is a repeating group in RequestForPositions
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//NoUnderlyings is a repeating group in RequestForPositions
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//NoTradingSessions is a repeating group in RequestForPositions
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a RequestForPositions FIX Message
type Message struct {
	FIXMsgType string `fix:"AN"`
	fix44.Header
	//PosReqID is a required field for RequestForPositions.
	PosReqID string `fix:"710"`
	//PosReqType is a required field for RequestForPositions.
	PosReqType int `fix:"724"`
	//MatchStatus is a non-required field for RequestForPositions.
	MatchStatus *string `fix:"573"`
	//SubscriptionRequestType is a non-required field for RequestForPositions.
	SubscriptionRequestType *string `fix:"263"`
	//Parties is a required component for RequestForPositions.
	parties.Parties
	//Account is a required field for RequestForPositions.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for RequestForPositions.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for RequestForPositions.
	AccountType int `fix:"581"`
	//Instrument is a non-required component for RequestForPositions.
	Instrument *instrument.Instrument
	//Currency is a non-required field for RequestForPositions.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for RequestForPositions.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for RequestForPositions.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//ClearingBusinessDate is a required field for RequestForPositions.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for RequestForPositions.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for RequestForPositions.
	SettlSessSubID *string `fix:"717"`
	//NoTradingSessions is a non-required field for RequestForPositions.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//TransactTime is a required field for RequestForPositions.
	TransactTime time.Time `fix:"60"`
	//ResponseTransportType is a non-required field for RequestForPositions.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for RequestForPositions.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for RequestForPositions.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for RequestForPositions.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for RequestForPositions.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosReqID(v string)                       { m.PosReqID = v }
func (m *Message) SetPosReqType(v int)                        { m.PosReqType = v }
func (m *Message) SetMatchStatus(v string)                    { m.MatchStatus = &v }
func (m *Message) SetSubscriptionRequestType(v string)        { m.SubscriptionRequestType = &v }
func (m *Message) SetParties(v parties.Parties)               { m.Parties = v }
func (m *Message) SetAccount(v string)                        { m.Account = v }
func (m *Message) SetAcctIDSource(v int)                      { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                       { m.AccountType = v }
func (m *Message) SetInstrument(v instrument.Instrument)      { m.Instrument = &v }
func (m *Message) SetCurrency(v string)                       { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                       { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *Message) SetClearingBusinessDate(v string)           { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                    { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                 { m.SettlSessSubID = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetResponseTransportType(v int)             { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)            { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                           { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                    { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                    { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AN", r
}
