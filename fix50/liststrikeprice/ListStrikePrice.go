//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtstrkpxgrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtstrkpxgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ListStrikePrice FIX Message
type Message struct {
	FIXMsgType string `fix:"m"`
	Header     fixt11.Header
	//ListID is a required field for ListStrikePrice.
	ListID string `fix:"66"`
	//TotNoStrikes is a required field for ListStrikePrice.
	TotNoStrikes int `fix:"422"`
	//LastFragment is a non-required field for ListStrikePrice.
	LastFragment *bool `fix:"893"`
	//InstrmtStrkPxGrp Component
	InstrmtStrkPxGrp instrmtstrkpxgrp.Component
	//UndInstrmtStrkPxGrp Component
	UndInstrmtStrkPxGrp undinstrmtstrkpxgrp.Component
	Trailer             fixt11.Trailer
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
	return enum.BeginStringFIX50, "m", r
}
