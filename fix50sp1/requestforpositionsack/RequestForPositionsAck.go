//Package requestforpositionsack msg type = AO.
package requestforpositionsack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a RequestForPositionsAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AO"`
	fixt11.Header
	//PosMaintRptID is a required field for RequestForPositionsAck.
	PosMaintRptID string `fix:"721"`
	//PosReqID is a non-required field for RequestForPositionsAck.
	PosReqID *string `fix:"710"`
	//TotalNumPosReports is a non-required field for RequestForPositionsAck.
	TotalNumPosReports *int `fix:"727"`
	//UnsolicitedIndicator is a non-required field for RequestForPositionsAck.
	UnsolicitedIndicator *bool `fix:"325"`
	//PosReqResult is a required field for RequestForPositionsAck.
	PosReqResult int `fix:"728"`
	//PosReqStatus is a required field for RequestForPositionsAck.
	PosReqStatus int `fix:"729"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for RequestForPositionsAck.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for RequestForPositionsAck.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for RequestForPositionsAck.
	AccountType *int `fix:"581"`
	//Instrument Component
	instrument.Instrument
	//Currency is a non-required field for RequestForPositionsAck.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp Component
	undinstrmtgrp.UndInstrmtGrp
	//ResponseTransportType is a non-required field for RequestForPositionsAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for RequestForPositionsAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for RequestForPositionsAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for RequestForPositionsAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for RequestForPositionsAck.
	EncodedText *string `fix:"355"`
	//PosReqType is a non-required field for RequestForPositionsAck.
	PosReqType *int `fix:"724"`
	//MatchStatus is a non-required field for RequestForPositionsAck.
	MatchStatus *string `fix:"573"`
	//ClearingBusinessDate is a non-required field for RequestForPositionsAck.
	ClearingBusinessDate *string `fix:"715"`
	//SubscriptionRequestType is a non-required field for RequestForPositionsAck.
	SubscriptionRequestType *string `fix:"263"`
	//SettlSessID is a non-required field for RequestForPositionsAck.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for RequestForPositionsAck.
	SettlSessSubID *string `fix:"717"`
	//SettlCurrency is a non-required field for RequestForPositionsAck.
	SettlCurrency *string `fix:"120"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosMaintRptID(v string)           { m.PosMaintRptID = v }
func (m *Message) SetPosReqID(v string)                { m.PosReqID = &v }
func (m *Message) SetTotalNumPosReports(v int)         { m.TotalNumPosReports = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)      { m.UnsolicitedIndicator = &v }
func (m *Message) SetPosReqResult(v int)               { m.PosReqResult = v }
func (m *Message) SetPosReqStatus(v int)               { m.PosReqStatus = v }
func (m *Message) SetAccount(v string)                 { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)               { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                { m.AccountType = &v }
func (m *Message) SetCurrency(v string)                { m.Currency = &v }
func (m *Message) SetResponseTransportType(v int)      { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)     { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                    { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)             { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)             { m.EncodedText = &v }
func (m *Message) SetPosReqType(v int)                 { m.PosReqType = &v }
func (m *Message) SetMatchStatus(v string)             { m.MatchStatus = &v }
func (m *Message) SetClearingBusinessDate(v string)    { m.ClearingBusinessDate = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }
func (m *Message) SetSettlSessID(v string)             { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)          { m.SettlSessSubID = &v }
func (m *Message) SetSettlCurrency(v string)           { m.SettlCurrency = &v }

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
	return enum.ApplVerID_FIX50SP1, "AO", r
}
