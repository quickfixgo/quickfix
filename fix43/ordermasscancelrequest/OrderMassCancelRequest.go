//Package ordermasscancelrequest msg type = q.
package ordermasscancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/underlyinginstrument"
	"time"
)

//Message is a OrderMassCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"q"`
	Header     fix43.Header
	//ClOrdID is a required field for OrderMassCancelRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassCancelRequest.
	SecondaryClOrdID *string `fix:"526"`
	//MassCancelRequestType is a required field for OrderMassCancelRequest.
	MassCancelRequestType string `fix:"530"`
	//TradingSessionID is a non-required field for OrderMassCancelRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassCancelRequest.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//Side is a non-required field for OrderMassCancelRequest.
	Side *string `fix:"54"`
	//TransactTime is a required field for OrderMassCancelRequest.
	TransactTime time.Time `fix:"60"`
	//Text is a non-required field for OrderMassCancelRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassCancelRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassCancelRequest.
	EncodedText *string `fix:"355"`
	Trailer     fix43.Trailer
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
	return enum.BeginStringFIX43, "q", r
}
