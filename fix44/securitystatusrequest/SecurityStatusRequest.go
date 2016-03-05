//Package securitystatusrequest msg type = e.
package securitystatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoUnderlyings is a repeating group in SecurityStatusRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in SecurityStatusRequest
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//Message is a SecurityStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"e"`
	fix44.Header
	//SecurityStatusReqID is a required field for SecurityStatusRequest.
	SecurityStatusReqID string `fix:"324"`
	//Instrument Component
	instrument.Instrument
	//InstrumentExtension Component
	instrumentextension.InstrumentExtension
	//NoUnderlyings is a non-required field for SecurityStatusRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for SecurityStatusRequest.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//Currency is a non-required field for SecurityStatusRequest.
	Currency *string `fix:"15"`
	//SubscriptionRequestType is a required field for SecurityStatusRequest.
	SubscriptionRequestType string `fix:"263"`
	//TradingSessionID is a non-required field for SecurityStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityStatusReqID(v string)     { m.SecurityStatusReqID = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)  { m.NoUnderlyings = v }
func (m *Message) SetNoLegs(v []NoLegs)                { m.NoLegs = v }
func (m *Message) SetCurrency(v string)                { m.Currency = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }

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
	return enum.BeginStringFIX44, "e", r
}
