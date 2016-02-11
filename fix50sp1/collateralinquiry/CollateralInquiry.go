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
	Header     fixt11.Header
	//CollInquiryID is a non-required field for CollateralInquiry.
	CollInquiryID *string `fix:"909"`
	//CollInqQualGrp Component
	CollInqQualGrp collinqqualgrp.Component
	//SubscriptionRequestType is a non-required field for CollateralInquiry.
	SubscriptionRequestType *string `fix:"263"`
	//ResponseTransportType is a non-required field for CollateralInquiry.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiry.
	ResponseDestination *string `fix:"726"`
	//Parties Component
	Parties parties.Component
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
	//ExecCollGrp Component
	ExecCollGrp execcollgrp.Component
	//TrdCollGrp Component
	TrdCollGrp trdcollgrp.Component
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralInquiry.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiry.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiry.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiry.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//MarginExcess is a non-required field for CollateralInquiry.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralInquiry.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralInquiry.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
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
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Stipulations Component
	Stipulations stipulations.Component
	//SettlInstructionsData Component
	SettlInstructionsData settlinstructionsdata.Component
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
	return enum.ApplVerID_FIX50SP1, "BB", r
}
