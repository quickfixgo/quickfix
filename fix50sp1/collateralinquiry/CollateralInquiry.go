//Package collateralinquiry msg type = BB.
package collateralinquiry

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/collinqqualgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a CollateralInquiry FIX Message
type Message struct {
	FIXMsgType string `fix:"BB"`
	fixt11.Header
	//CollInquiryID is a non-required field for CollateralInquiry.
	CollInquiryID *string `fix:"909"`
	//CollInqQualGrp is a non-required component for CollateralInquiry.
	CollInqQualGrp *collinqqualgrp.CollInqQualGrp
	//SubscriptionRequestType is a non-required field for CollateralInquiry.
	SubscriptionRequestType *string `fix:"263"`
	//ResponseTransportType is a non-required field for CollateralInquiry.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiry.
	ResponseDestination *string `fix:"726"`
	//Parties is a non-required component for CollateralInquiry.
	Parties *parties.Parties
	//Account is a non-required field for CollateralInquiry.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralInquiry.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralInquiry.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralInquiry.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralInquiry.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralInquiry.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp is a non-required component for CollateralInquiry.
	ExecCollGrp *execcollgrp.ExecCollGrp
	//TrdCollGrp is a non-required component for CollateralInquiry.
	TrdCollGrp *trdcollgrp.TrdCollGrp
	//Instrument is a non-required component for CollateralInquiry.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralInquiry.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralInquiry.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiry.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiry.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiry.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for CollateralInquiry.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp is a non-required component for CollateralInquiry.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//MarginExcess is a non-required field for CollateralInquiry.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralInquiry.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralInquiry.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps is a non-required component for CollateralInquiry.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
	//Side is a non-required field for CollateralInquiry.
	Side *string `fix:"54"`
	//Price is a non-required field for CollateralInquiry.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralInquiry.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralInquiry.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralInquiry.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralInquiry.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralInquiry.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData is a non-required component for CollateralInquiry.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations is a non-required component for CollateralInquiry.
	Stipulations *stipulations.Stipulations
	//SettlInstructionsData is a non-required component for CollateralInquiry.
	SettlInstructionsData *settlinstructionsdata.SettlInstructionsData
	//TradingSessionID is a non-required field for CollateralInquiry.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralInquiry.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralInquiry.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralInquiry.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralInquiry.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralInquiry.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralInquiry.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralInquiry.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCollInquiryID(v string)                               { m.CollInquiryID = &v }
func (m *Message) SetCollInqQualGrp(v collinqqualgrp.CollInqQualGrp)       { m.CollInqQualGrp = &v }
func (m *Message) SetSubscriptionRequestType(v string)                     { m.SubscriptionRequestType = &v }
func (m *Message) SetResponseTransportType(v int)                          { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)                         { m.ResponseDestination = &v }
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
func (m *Message) SetMarginExcess(v float64)                               { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                              { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                            { m.CashOutstanding = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)                         { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)                      { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)                                  { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                                    { m.EndCash = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *Message) SetSettlInstructionsData(v settlinstructionsdata.SettlInstructionsData) {
	m.SettlInstructionsData = &v
}
func (m *Message) SetTradingSessionID(v string)     { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)  { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)          { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)       { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP1, "BB", r
}
