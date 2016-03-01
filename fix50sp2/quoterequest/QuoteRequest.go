//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/quotreqgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/rootparties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a QuoteRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"R"`
	Header     fixt11.Header
	//QuoteReqID is a required field for QuoteRequest.
	QuoteReqID string `fix:"131"`
	//RFQReqID is a non-required field for QuoteRequest.
	RFQReqID *string `fix:"644"`
	//ClOrdID is a non-required field for QuoteRequest.
	ClOrdID *string `fix:"11"`
	//OrderCapacity is a non-required field for QuoteRequest.
	OrderCapacity *string `fix:"528"`
	//QuotReqGrp Component
	QuotReqGrp quotreqgrp.Component
	//Text is a non-required field for QuoteRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteRequest.
	EncodedText *string `fix:"355"`
	//RootParties Component
	RootParties rootparties.Component
	//PrivateQuote is a non-required field for QuoteRequest.
	PrivateQuote *bool `fix:"1171"`
	//RespondentType is a non-required field for QuoteRequest.
	RespondentType *int `fix:"1172"`
	//PreTradeAnonymity is a non-required field for QuoteRequest.
	PreTradeAnonymity *bool `fix:"1091"`
	//BookingType is a non-required field for QuoteRequest.
	BookingType *int `fix:"775"`
	//OrderRestrictions is a non-required field for QuoteRequest.
	OrderRestrictions *string `fix:"529"`
	Trailer           fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "R", r
}
