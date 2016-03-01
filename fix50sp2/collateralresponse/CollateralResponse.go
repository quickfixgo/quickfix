//Package collateralresponse msg type = AZ.
package collateralresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/execcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/trdcollgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtcollgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a CollateralResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AZ"`
	Header     fixt11.Header
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
	//Parties Component
	Parties parties.Component
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
	//ExecCollGrp Component
	ExecCollGrp execcollgrp.Component
	//TrdCollGrp Component
	TrdCollGrp trdcollgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralResponse.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralResponse.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralResponse.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralResponse.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtCollGrp Component
	UndInstrmtCollGrp undinstrmtcollgrp.Component
	//MarginExcess is a non-required field for CollateralResponse.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralResponse.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralResponse.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//Side is a non-required field for CollateralResponse.
	Side *string `fix:"54"`
	//MiscFeesGrp Component
	MiscFeesGrp miscfeesgrp.Component
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
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Stipulations Component
	Stipulations stipulations.Component
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
	Trailer              fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "AZ", r
}
