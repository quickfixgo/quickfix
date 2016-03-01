//Package massquote msg type = i.
package massquote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/quotsetgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MassQuote FIX Message
type Message struct {
	FIXMsgType string `fix:"i"`
	Header     fixt11.Header
	//QuoteReqID is a non-required field for MassQuote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for MassQuote.
	QuoteID string `fix:"117"`
	//QuoteType is a non-required field for MassQuote.
	QuoteType *int `fix:"537"`
	//QuoteResponseLevel is a non-required field for MassQuote.
	QuoteResponseLevel *int `fix:"301"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for MassQuote.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for MassQuote.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for MassQuote.
	AccountType *int `fix:"581"`
	//DefBidSize is a non-required field for MassQuote.
	DefBidSize *float64 `fix:"293"`
	//DefOfferSize is a non-required field for MassQuote.
	DefOfferSize *float64 `fix:"294"`
	//QuotSetGrp Component
	QuotSetGrp quotsetgrp.Component
	Trailer    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "i", r
}
