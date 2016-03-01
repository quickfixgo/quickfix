//Package collateralresponse msg type = AZ.
package collateralresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoExecs is a repeating group in CollateralResponse
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//NoTrades is a repeating group in CollateralResponse
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//NoLegs is a repeating group in CollateralResponse
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in CollateralResponse
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//CollAction is a non-required field for NoUnderlyings.
	CollAction *int `fix:"944"`
}

//NoMiscFees is a repeating group in CollateralResponse
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
	//MiscFeeBasis is a non-required field for NoMiscFees.
	MiscFeeBasis *int `fix:"891"`
}

//Message is a CollateralResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AZ"`
	Header     fix44.Header
	//CollRespID is a required field for CollateralResponse.
	CollRespID string `fix:"904"`
	//CollAsgnID is a required field for CollateralResponse.
	CollAsgnID string `fix:"902"`
	//CollReqID is a non-required field for CollateralResponse.
	CollReqID *string `fix:"894"`
	//CollAsgnReason is a required field for CollateralResponse.
	CollAsgnReason int `fix:"895"`
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
	//NoExecs is a non-required field for CollateralResponse.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralResponse.
	NoTrades []NoTrades `fix:"897,omitempty"`
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
	//NoLegs is a non-required field for CollateralResponse.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralResponse.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
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
	//NoMiscFees is a non-required field for CollateralResponse.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
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
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "AZ", r
}
