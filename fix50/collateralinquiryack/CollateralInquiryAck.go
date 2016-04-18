//Package collateralinquiryack msg type = BG.
package collateralinquiryack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/collinqqualgrp"
	"github.com/quickfixgo/quickfix/fix50/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a CollateralInquiryAck FIX Message
type Message struct {
	FIXMsgType string `fix:"BG"`
	fixt11.Header
	//CollInquiryID is a required field for CollateralInquiryAck.
	CollInquiryID string `fix:"909"`
	//CollInquiryStatus is a required field for CollateralInquiryAck.
	CollInquiryStatus int `fix:"945"`
	//CollInquiryResult is a non-required field for CollateralInquiryAck.
	CollInquiryResult *int `fix:"946"`
	//CollInqQualGrp is a non-required component for CollateralInquiryAck.
	CollInqQualGrp *collinqqualgrp.CollInqQualGrp
	//TotNumReports is a non-required field for CollateralInquiryAck.
	TotNumReports *int `fix:"911"`
	//Parties is a non-required component for CollateralInquiryAck.
	Parties *parties.Parties
	//Account is a non-required field for CollateralInquiryAck.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralInquiryAck.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralInquiryAck.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralInquiryAck.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralInquiryAck.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp is a non-required component for CollateralInquiryAck.
	ExecCollGrp *execcollgrp.ExecCollGrp
	//TrdCollGrp is a non-required component for CollateralInquiryAck.
	TrdCollGrp *trdcollgrp.TrdCollGrp
	//Instrument is a non-required component for CollateralInquiryAck.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralInquiryAck.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralInquiryAck.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiryAck.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiryAck.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiryAck.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for CollateralInquiryAck.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp is a non-required component for CollateralInquiryAck.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//TradingSessionID is a non-required field for CollateralInquiryAck.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralInquiryAck.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralInquiryAck.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralInquiryAck.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
	ClearingBusinessDate *string `fix:"715"`
	//ResponseTransportType is a non-required field for CollateralInquiryAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiryAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for CollateralInquiryAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralInquiryAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralInquiryAck.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized CollateralInquiryAck instance
func New(collinquiryid string, collinquirystatus int) *Message {
	var m Message
	m.SetCollInquiryID(collinquiryid)
	m.SetCollInquiryStatus(collinquirystatus)
	return &m
}

func (m *Message) SetCollInquiryID(v string)                               { m.CollInquiryID = v }
func (m *Message) SetCollInquiryStatus(v int)                              { m.CollInquiryStatus = v }
func (m *Message) SetCollInquiryResult(v int)                              { m.CollInquiryResult = &v }
func (m *Message) SetCollInqQualGrp(v collinqqualgrp.CollInqQualGrp)       { m.CollInqQualGrp = &v }
func (m *Message) SetTotNumReports(v int)                                  { m.TotNumReports = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                                     { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)                            { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                            { m.SecondaryClOrdID = &v }
func (m *Message) SetExecCollGrp(v execcollgrp.ExecCollGrp)                { m.ExecCollGrp = &v }
func (m *Message) SetTrdCollGrp(v trdcollgrp.TrdCollGrp)                   { m.TrdCollGrp = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = &v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)                                   { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)          { m.InstrmtLegGrp = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetTradingSessionID(v string)                            { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                         { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)                                 { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                              { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = &v }
func (m *Message) SetResponseTransportType(v int)                          { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)                         { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                                        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                                 { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                                 { m.EncodedText = &v }

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
	return enum.BeginStringFIX50, "BG", r
}
