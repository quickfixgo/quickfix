//Package securitydefinitionrequest msg type = c.
package securitydefinitionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoUnderlyings is a repeating group in SecurityDefinitionRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in SecurityDefinitionRequest
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//Message is a SecurityDefinitionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"c"`
	Header     fix44.Header
	//SecurityReqID is a required field for SecurityDefinitionRequest.
	SecurityReqID string `fix:"320"`
	//SecurityRequestType is a required field for SecurityDefinitionRequest.
	SecurityRequestType int `fix:"321"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//NoUnderlyings is a non-required field for SecurityDefinitionRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Currency is a non-required field for SecurityDefinitionRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinitionRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityDefinitionRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityDefinitionRequest.
	TradingSessionSubID *string `fix:"625"`
	//NoLegs is a non-required field for SecurityDefinitionRequest.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//ExpirationCycle is a non-required field for SecurityDefinitionRequest.
	ExpirationCycle *int `fix:"827"`
	//SubscriptionRequestType is a non-required field for SecurityDefinitionRequest.
	SubscriptionRequestType *string `fix:"263"`
	Trailer                 fix44.Trailer
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
	return enum.BeginStringFIX44, "c", r
}
