//Package collateralrequest msg type = AX.
package collateralrequest

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

//Message is a CollateralRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AX"`
	Header     fixt11.Header
	//CollReqID is a required field for CollateralRequest.
	CollReqID string `fix:"894"`
	//CollAsgnReason is a required field for CollateralRequest.
	CollAsgnReason int `fix:"895"`
	//TransactTime is a required field for CollateralRequest.
	TransactTime time.Time `fix:"60"`
	//ExpireTime is a non-required field for CollateralRequest.
	ExpireTime *time.Time `fix:"126"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for CollateralRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralRequest.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralRequest.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralRequest.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralRequest.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralRequest.
	SecondaryClOrdID *string `fix:"526"`
	//ExecCollGrp Component
	ExecCollGrp execcollgrp.Component
	//TrdCollGrp Component
	TrdCollGrp trdcollgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralRequest.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralRequest.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralRequest.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralRequest.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtCollGrp Component
	UndInstrmtCollGrp undinstrmtcollgrp.Component
	//MarginExcess is a non-required field for CollateralRequest.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralRequest.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralRequest.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//Side is a non-required field for CollateralRequest.
	Side *string `fix:"54"`
	//MiscFeesGrp Component
	MiscFeesGrp miscfeesgrp.Component
	//Price is a non-required field for CollateralRequest.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralRequest.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralRequest.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralRequest.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralRequest.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralRequest.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Stipulations Component
	Stipulations stipulations.Component
	//TradingSessionID is a non-required field for CollateralRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralRequest.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralRequest.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralRequest.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralRequest.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralRequest.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
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
	return enum.BeginStringFIX50, "AX", r
}
