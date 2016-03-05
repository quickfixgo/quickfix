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

func (m *NoExecs) SetExecID(v string) { m.ExecID = &v }

//NoTrades is a repeating group in CollateralResponse
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

func (m *NoTrades) SetTradeReportID(v string)          { m.TradeReportID = &v }
func (m *NoTrades) SetSecondaryTradeReportID(v string) { m.SecondaryTradeReportID = &v }

//NoLegs is a repeating group in CollateralResponse
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//NoUnderlyings is a repeating group in CollateralResponse
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//CollAction is a non-required field for NoUnderlyings.
	CollAction *int `fix:"944"`
}

func (m *NoUnderlyings) SetCollAction(v int) { m.CollAction = &v }

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

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }
func (m *NoMiscFees) SetMiscFeeBasis(v int)   { m.MiscFeeBasis = &v }

//Message is a CollateralResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"AZ"`
	fix44.Header
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
	parties.Parties
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
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
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
	trdregtimestamps.TrdRegTimestamps
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
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations Component
	stipulations.Stipulations
	//Text is a non-required field for CollateralResponse.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralResponse.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralResponse.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCollRespID(v string)             { m.CollRespID = v }
func (m *Message) SetCollAsgnID(v string)             { m.CollAsgnID = v }
func (m *Message) SetCollReqID(v string)              { m.CollReqID = &v }
func (m *Message) SetCollAsgnReason(v int)            { m.CollAsgnReason = v }
func (m *Message) SetCollAsgnTransType(v int)         { m.CollAsgnTransType = &v }
func (m *Message) SetCollAsgnRespType(v int)          { m.CollAsgnRespType = v }
func (m *Message) SetCollAsgnRejectReason(v int)      { m.CollAsgnRejectReason = &v }
func (m *Message) SetTransactTime(v time.Time)        { m.TransactTime = v }
func (m *Message) SetAccount(v string)                { m.Account = &v }
func (m *Message) SetAccountType(v int)               { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)       { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)       { m.SecondaryClOrdID = &v }
func (m *Message) SetNoExecs(v []NoExecs)             { m.NoExecs = v }
func (m *Message) SetNoTrades(v []NoTrades)           { m.NoTrades = v }
func (m *Message) SetSettlDate(v string)              { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)              { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                   { m.QtyType = &v }
func (m *Message) SetCurrency(v string)               { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)               { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *Message) SetMarginExcess(v float64)          { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)         { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)       { m.CashOutstanding = &v }
func (m *Message) SetSide(v string)                   { m.Side = &v }
func (m *Message) SetNoMiscFees(v []NoMiscFees)       { m.NoMiscFees = v }
func (m *Message) SetPrice(v float64)                 { m.Price = &v }
func (m *Message) SetPriceType(v int)                 { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)    { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64) { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)             { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)               { m.EndCash = &v }
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
	return enum.BeginStringFIX44, "AZ", r
}
