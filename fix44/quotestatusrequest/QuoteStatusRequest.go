//Package quotestatusrequest msg type = a.
package quotestatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoUnderlyings is a repeating group in QuoteStatusRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in QuoteStatusRequest
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//Message is a QuoteStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"a"`
	Header     fix44.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusRequest.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteID is a non-required field for QuoteStatusRequest.
	QuoteID *string `fix:"117"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//NoUnderlyings is a non-required field for QuoteStatusRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for QuoteStatusRequest.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for QuoteStatusRequest.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for QuoteStatusRequest.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for QuoteStatusRequest.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for QuoteStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for QuoteStatusRequest.
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
	return enum.BeginStringFIX44, "a", r
}
