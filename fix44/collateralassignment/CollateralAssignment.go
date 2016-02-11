//Package collateralassignment msg type = AY.
package collateralassignment

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoExecs is a repeating group in CollateralAssignment
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//NoTrades is a repeating group in CollateralAssignment
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//NoLegs is a repeating group in CollateralAssignment
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in CollateralAssignment
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//CollAction is a non-required field for NoUnderlyings.
	CollAction *int `fix:"944"`
}

//NoMiscFees is a repeating group in CollateralAssignment
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

//Message is a CollateralAssignment FIX Message
type Message struct {
	FIXMsgType string `fix:"AY"`
	Header     fix44.Header
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
	Parties parties.Component
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
	//NoExecs is a non-required field for CollateralAssignment.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralAssignment.
	NoTrades []NoTrades `fix:"897,omitempty"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralAssignment.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralAssignment.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralAssignment.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralAssignment.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for CollateralAssignment.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralAssignment.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//MarginExcess is a non-required field for CollateralAssignment.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralAssignment.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralAssignment.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//Side is a non-required field for CollateralAssignment.
	Side *string `fix:"54"`
	//NoMiscFees is a non-required field for CollateralAssignment.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
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
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Stipulations Component
	Stipulations stipulations.Component
	//SettlInstructionsData Component
	SettlInstructionsData settlinstructionsdata.Component
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
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "AY", r
}
