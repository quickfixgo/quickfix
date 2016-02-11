//Package collateralreport msg type = BA.
package collateralreport

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
)

//NoExecs is a repeating group in CollateralReport
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//NoTrades is a repeating group in CollateralReport
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//NoLegs is a repeating group in CollateralReport
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in CollateralReport
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoMiscFees is a repeating group in CollateralReport
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

//Message is a CollateralReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BA"`
	Header     fix44.Header
	//CollRptID is a required field for CollateralReport.
	CollRptID string `fix:"908"`
	//CollInquiryID is a non-required field for CollateralReport.
	CollInquiryID *string `fix:"909"`
	//CollStatus is a required field for CollateralReport.
	CollStatus int `fix:"910"`
	//TotNumReports is a non-required field for CollateralReport.
	TotNumReports *int `fix:"911"`
	//LastRptRequested is a non-required field for CollateralReport.
	LastRptRequested *bool `fix:"912"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for CollateralReport.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralReport.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralReport.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralReport.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralReport.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralReport.
	SecondaryClOrdID *string `fix:"526"`
	//NoExecs is a non-required field for CollateralReport.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralReport.
	NoTrades []NoTrades `fix:"897,omitempty"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralReport.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralReport.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralReport.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralReport.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for CollateralReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//MarginExcess is a non-required field for CollateralReport.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralReport.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralReport.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//Side is a non-required field for CollateralReport.
	Side *string `fix:"54"`
	//NoMiscFees is a non-required field for CollateralReport.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	//Price is a non-required field for CollateralReport.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for CollateralReport.
	PriceType *int `fix:"423"`
	//AccruedInterestAmt is a non-required field for CollateralReport.
	AccruedInterestAmt *float64 `fix:"159"`
	//EndAccruedInterestAmt is a non-required field for CollateralReport.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for CollateralReport.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for CollateralReport.
	EndCash *float64 `fix:"922"`
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//Stipulations Component
	Stipulations stipulations.Component
	//SettlInstructionsData Component
	SettlInstructionsData settlinstructionsdata.Component
	//TradingSessionID is a non-required field for CollateralReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralReport.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralReport.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralReport.
	ClearingBusinessDate *string `fix:"715"`
	//Text is a non-required field for CollateralReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralReport.
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
	return enum.BeginStringFIX44, "BA", r
}
