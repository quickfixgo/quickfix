//Package collateralresponse msg type = AZ.
package collateralresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtcollgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CollateralResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AZ"`
	fixt11.Header
	//CollRespID is a required field for CollateralResponse.
	CollRespID string `fix:"904"`
	//CollAsgnID is a non-required field for CollateralResponse.
	CollAsgnID *string `fix:"902"`
	//CollReqID is a non-required field for CollateralResponse.
	CollReqID *string `fix:"894"`
	//CollAsgnReason is a non-required field for CollateralResponse.
	CollAsgnReason *int `fix:"895"`
	//CollAsgnTransType is a non-required field for CollateralResponse.
	CollAsgnTransType *int `fix:"903"`
	//CollAsgnRespType is a required field for CollateralResponse.
	CollAsgnRespType int `fix:"905"`
	//CollAsgnRejectReason is a non-required field for CollateralResponse.
	CollAsgnRejectReason *int `fix:"906"`
	//TransactTime is a required field for CollateralResponse.
	TransactTime time.Time `fix:"60"`
	//Parties is a non-required component for CollateralResponse.
	Parties *parties.Parties
	//Account is a non-required field for CollateralResponse.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralResponse.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralResponse.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralResponse.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralResponse.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralResponse.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp is a non-required component for CollateralResponse.
	ExecCollGrp *execcollgrp.ExecCollGrp
	//TrdCollGrp is a non-required component for CollateralResponse.
	TrdCollGrp *trdcollgrp.TrdCollGrp
	//Instrument is a non-required component for CollateralResponse.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralResponse.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralResponse.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralResponse.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralResponse.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralResponse.
	Currency *string `fix:"15"`
	//InstrmtLegGrp is a non-required component for CollateralResponse.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtCollGrp is a non-required component for CollateralResponse.
	UndInstrmtCollGrp *undinstrmtcollgrp.UndInstrmtCollGrp
	//MarginExcess is a non-required field for CollateralResponse.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralResponse.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralResponse.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps is a non-required component for CollateralResponse.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
	//Side is a non-required field for CollateralResponse.
	Side *string `fix:"54"`
	//MiscFeesGrp is a non-required component for CollateralResponse.
	MiscFeesGrp *miscfeesgrp.MiscFeesGrp
	//Price is a non-required field for CollateralResponse.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralResponse.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralResponse.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralResponse.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralResponse.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralResponse.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData is a non-required component for CollateralResponse.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations is a non-required component for CollateralResponse.
	Stipulations *stipulations.Stipulations
	//Text is a non-required field for CollateralResponse.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralResponse.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralResponse.
	EncodedText *string `fix:"355"`
	//CollApplType is a non-required field for CollateralResponse.
	CollApplType *int `fix:"1043"`
	//FinancialStatus is a non-required field for CollateralResponse.
	FinancialStatus *string `fix:"291"`
	//ClearingBusinessDate is a non-required field for CollateralResponse.
	ClearingBusinessDate *string `fix:"715"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized CollateralResponse instance
func New(collrespid string, collasgnresptype int, transacttime time.Time) *Message {
	var m Message
	m.SetCollRespID(collrespid)
	m.SetCollAsgnRespType(collasgnresptype)
	m.SetTransactTime(transacttime)
	return &m
}

func (m *Message) SetCollRespID(v string)                                  { m.CollRespID = v }
func (m *Message) SetCollAsgnID(v string)                                  { m.CollAsgnID = &v }
func (m *Message) SetCollReqID(v string)                                   { m.CollReqID = &v }
func (m *Message) SetCollAsgnReason(v int)                                 { m.CollAsgnReason = &v }
func (m *Message) SetCollAsgnTransType(v int)                              { m.CollAsgnTransType = &v }
func (m *Message) SetCollAsgnRespType(v int)                               { m.CollAsgnRespType = v }
func (m *Message) SetCollAsgnRejectReason(v int)                           { m.CollAsgnRejectReason = &v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = v }
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
func (m *Message) SetUndInstrmtCollGrp(v undinstrmtcollgrp.UndInstrmtCollGrp) {
	m.UndInstrmtCollGrp = &v
}
func (m *Message) SetMarginExcess(v float64)                               { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                              { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                            { m.CashOutstanding = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetMiscFeesGrp(v miscfeesgrp.MiscFeesGrp)                { m.MiscFeesGrp = &v }
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
func (m *Message) SetText(v string)                            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                     { m.EncodedText = &v }
func (m *Message) SetCollApplType(v int)                       { m.CollApplType = &v }
func (m *Message) SetFinancialStatus(v string)                 { m.FinancialStatus = &v }
func (m *Message) SetClearingBusinessDate(v string)            { m.ClearingBusinessDate = &v }

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
	return enum.BeginStringFIX50, "AZ", r
}
