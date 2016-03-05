//Package collateralassignment msg type = AY.
package collateralassignment

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtcollgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CollateralAssignment FIX Message
type Message struct {
	FIXMsgType string `fix:"AY"`
	fixt11.Header
	//CollAsgnID is a required field for CollateralAssignment.
	CollAsgnID string `fix:"902"`
	//CollReqID is a non-required field for CollateralAssignment.
	CollReqID *string `fix:"894"`
	//CollAsgnReason is a required field for CollateralAssignment.
	CollAsgnReason int `fix:"895"`
	//CollAsgnTransType is a required field for CollateralAssignment.
	CollAsgnTransType int `fix:"903"`
	//CollAsgnRefID is a non-required field for CollateralAssignment.
	CollAsgnRefID *string `fix:"907"`
	//TransactTime is a required field for CollateralAssignment.
	TransactTime time.Time `fix:"60"`
	//ExpireTime is a non-required field for CollateralAssignment.
	ExpireTime *time.Time `fix:"126"`
	//Parties Component
	parties.Parties
	//Account is a non-required field for CollateralAssignment.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralAssignment.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralAssignment.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralAssignment.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralAssignment.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralAssignment.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp Component
	execcollgrp.ExecCollGrp
	//TrdCollGrp Component
	trdcollgrp.TrdCollGrp
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralAssignment.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralAssignment.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralAssignment.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralAssignment.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	instrmtleggrp.InstrmtLegGrp
	//UndInstrmtCollGrp Component
	undinstrmtcollgrp.UndInstrmtCollGrp
	//MarginExcess is a non-required field for CollateralAssignment.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralAssignment.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralAssignment.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	trdregtimestamps.TrdRegTimestamps
	//Side is a non-required field for CollateralAssignment.
	Side *string `fix:"54"`
	//MiscFeesGrp Component
	miscfeesgrp.MiscFeesGrp
	//Price is a non-required field for CollateralAssignment.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralAssignment.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralAssignment.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralAssignment.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralAssignment.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralAssignment.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations Component
	stipulations.Stipulations
	//SettlInstructionsData Component
	settlinstructionsdata.SettlInstructionsData
	//TradingSessionID is a non-required field for CollateralAssignment.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralAssignment.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralAssignment.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralAssignment.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralAssignment.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralAssignment.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralAssignment.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralAssignment.
	EncodedText *string `fix:"355"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCollAsgnID(v string)             { m.CollAsgnID = v }
func (m *Message) SetCollReqID(v string)              { m.CollReqID = &v }
func (m *Message) SetCollAsgnReason(v int)            { m.CollAsgnReason = v }
func (m *Message) SetCollAsgnTransType(v int)         { m.CollAsgnTransType = v }
func (m *Message) SetCollAsgnRefID(v string)          { m.CollAsgnRefID = &v }
func (m *Message) SetTransactTime(v time.Time)        { m.TransactTime = v }
func (m *Message) SetExpireTime(v time.Time)          { m.ExpireTime = &v }
func (m *Message) SetAccount(v string)                { m.Account = &v }
func (m *Message) SetAccountType(v int)               { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)       { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)       { m.SecondaryClOrdID = &v }
func (m *Message) SetSettlDate(v string)              { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)              { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                   { m.QtyType = &v }
func (m *Message) SetCurrency(v string)               { m.Currency = &v }
func (m *Message) SetMarginExcess(v float64)          { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)         { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)       { m.CashOutstanding = &v }
func (m *Message) SetSide(v string)                   { m.Side = &v }
func (m *Message) SetPrice(v float64)                 { m.Price = &v }
func (m *Message) SetPriceType(v int)                 { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)    { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64) { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)             { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)               { m.EndCash = &v }
func (m *Message) SetTradingSessionID(v string)       { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)    { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)            { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)         { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string)   { m.ClearingBusinessDate = &v }
func (m *Message) SetText(v string)                   { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)            { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP2, "AY", r
}
