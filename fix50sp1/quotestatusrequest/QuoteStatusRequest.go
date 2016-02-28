//Package quotestatusrequest msg type = a.
package quotestatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a QuoteStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"a"`
	Header     fixt11.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusRequest.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteID is a non-required field for QuoteStatusRequest.
	QuoteID *string `fix:"117"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
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
	Trailer                 fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteStatusReqID(v string)        { m.QuoteStatusReqID = &v }
func (m *Message) SetQuoteID(v string)                 { m.QuoteID = &v }
func (m *Message) SetAccount(v string)                 { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)               { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                { m.AccountType = &v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }

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
	return enum.ApplVerID_FIX50SP1, "a", r
}
