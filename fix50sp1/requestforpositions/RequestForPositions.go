//Package requestforpositions msg type = AN.
package requestforpositions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a RequestForPositions FIX Message
type Message struct {
	FIXMsgType string `fix:"AN"`
	Header     fixt11.Header
	//PosReqID is a required field for RequestForPositions.
	PosReqID string `fix:"710"`
	//PosReqType is a required field for RequestForPositions.
	PosReqType int `fix:"724"`
	//MatchStatus is a non-required field for RequestForPositions.
	MatchStatus *string `fix:"573"`
	//SubscriptionRequestType is a non-required field for RequestForPositions.
	SubscriptionRequestType *string `fix:"263"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for RequestForPositions.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for RequestForPositions.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for RequestForPositions.
	AccountType *int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for RequestForPositions.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//ClearingBusinessDate is a required field for RequestForPositions.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for RequestForPositions.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for RequestForPositions.
	SettlSessSubID *string `fix:"717"`
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
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
	//SettlCurrency is a non-required field for RequestForPositions.
	SettlCurrency *string `fix:"120"`
	Trailer       fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "AN", r
}
