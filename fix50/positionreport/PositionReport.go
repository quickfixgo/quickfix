//Package positionreport msg type = AP.
package positionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50/positionqty"
	"github.com/quickfixgo/quickfix/fix50/posundinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a PositionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AP"`
	fixt11.Header
	//PosMaintRptID is a required field for PositionReport.
	PosMaintRptID string `fix:"721"`
	//PosReqID is a non-required field for PositionReport.
	PosReqID *string `fix:"710"`
	//PosReqType is a non-required field for PositionReport.
	PosReqType *int `fix:"724"`
	//SubscriptionRequestType is a non-required field for PositionReport.
	SubscriptionRequestType *string `fix:"263"`
	//TotalNumPosReports is a non-required field for PositionReport.
	TotalNumPosReports *int `fix:"727"`
	//UnsolicitedIndicator is a non-required field for PositionReport.
	UnsolicitedIndicator *bool `fix:"325"`
	//PosReqResult is a non-required field for PositionReport.
	PosReqResult *int `fix:"728"`
	//ClearingBusinessDate is a required field for PositionReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for PositionReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for PositionReport.
	SettlSessSubID *string `fix:"717"`
	//Parties is a required component for PositionReport.
	parties.Parties
	//Account is a non-required field for PositionReport.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for PositionReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for PositionReport.
	AccountType *int `fix:"581"`
	//Instrument is a non-required component for PositionReport.
	Instrument *instrument.Instrument
	//Currency is a non-required field for PositionReport.
	Currency *string `fix:"15"`
	//SettlPrice is a non-required field for PositionReport.
	SettlPrice *float64 `fix:"730"`
	//SettlPriceType is a non-required field for PositionReport.
	SettlPriceType *int `fix:"731"`
	//PriorSettlPrice is a non-required field for PositionReport.
	PriorSettlPrice *float64 `fix:"734"`
	//InstrmtLegGrp is a non-required component for PositionReport.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//PosUndInstrmtGrp is a non-required component for PositionReport.
	PosUndInstrmtGrp *posundinstrmtgrp.PosUndInstrmtGrp
	//PositionQty is a non-required component for PositionReport.
	PositionQty *positionqty.PositionQty
	//PositionAmountData is a non-required component for PositionReport.
	PositionAmountData *positionamountdata.PositionAmountData
	//RegistStatus is a non-required field for PositionReport.
	RegistStatus *string `fix:"506"`
	//DeliveryDate is a non-required field for PositionReport.
	DeliveryDate *string `fix:"743"`
	//Text is a non-required field for PositionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PositionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PositionReport.
	EncodedText *string `fix:"355"`
	//MatchStatus is a non-required field for PositionReport.
	MatchStatus *string `fix:"573"`
	//PriceType is a non-required field for PositionReport.
	PriceType *int `fix:"423"`
	//SettlCurrency is a non-required field for PositionReport.
	SettlCurrency *string `fix:"120"`
	//MessageEventSource is a non-required field for PositionReport.
	MessageEventSource *string `fix:"1011"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized PositionReport instance
func New(posmaintrptid string, clearingbusinessdate string, parties parties.Parties) *Message {
	var m Message
	m.SetPosMaintRptID(posmaintrptid)
	m.SetClearingBusinessDate(clearingbusinessdate)
	m.SetParties(parties)
	return &m
}

func (m *Message) SetPosMaintRptID(v string)                               { m.PosMaintRptID = v }
func (m *Message) SetPosReqID(v string)                                    { m.PosReqID = &v }
func (m *Message) SetPosReqType(v int)                                     { m.PosReqType = &v }
func (m *Message) SetSubscriptionRequestType(v string)                     { m.SubscriptionRequestType = &v }
func (m *Message) SetTotalNumPosReports(v int)                             { m.TotalNumPosReports = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)                          { m.UnsolicitedIndicator = &v }
func (m *Message) SetPosReqResult(v int)                                   { m.PosReqResult = &v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                                 { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                              { m.SettlSessSubID = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                                   { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetSettlPrice(v float64)                                 { m.SettlPrice = &v }
func (m *Message) SetSettlPriceType(v int)                                 { m.SettlPriceType = &v }
func (m *Message) SetPriorSettlPrice(v float64)                            { m.PriorSettlPrice = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetPosUndInstrmtGrp(v posundinstrmtgrp.PosUndInstrmtGrp) { m.PosUndInstrmtGrp = &v }
func (m *Message) SetPositionQty(v positionqty.PositionQty)                { m.PositionQty = &v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = &v
}
func (m *Message) SetRegistStatus(v string)       { m.RegistStatus = &v }
func (m *Message) SetDeliveryDate(v string)       { m.DeliveryDate = &v }
func (m *Message) SetText(v string)               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)        { m.EncodedText = &v }
func (m *Message) SetMatchStatus(v string)        { m.MatchStatus = &v }
func (m *Message) SetPriceType(v int)             { m.PriceType = &v }
func (m *Message) SetSettlCurrency(v string)      { m.SettlCurrency = &v }
func (m *Message) SetMessageEventSource(v string) { m.MessageEventSource = &v }

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
	return enum.BeginStringFIX50, "AP", r
}
