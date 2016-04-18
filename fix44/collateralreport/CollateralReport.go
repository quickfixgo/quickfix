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

//NewNoExecs returns an initialized NoExecs instance
func NewNoExecs() *NoExecs {
	var m NoExecs
	return &m
}

func (m *NoExecs) SetExecID(v string) { m.ExecID = &v }

//NoTrades is a repeating group in CollateralReport
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//NewNoTrades returns an initialized NoTrades instance
func NewNoTrades() *NoTrades {
	var m NoTrades
	return &m
}

func (m *NoTrades) SetTradeReportID(v string)          { m.TradeReportID = &v }
func (m *NoTrades) SetSecondaryTradeReportID(v string) { m.SecondaryTradeReportID = &v }

//NoLegs is a repeating group in CollateralReport
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//NoUnderlyings is a repeating group in CollateralReport
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
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

//NewNoMiscFees returns an initialized NoMiscFees instance
func NewNoMiscFees() *NoMiscFees {
	var m NoMiscFees
	return &m
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }
func (m *NoMiscFees) SetMiscFeeBasis(v int)   { m.MiscFeeBasis = &v }

//Message is a CollateralReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BA"`
	fix44.Header
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
	//Parties is a non-required component for CollateralReport.
	Parties *parties.Parties
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
	//Instrument is a non-required component for CollateralReport.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralReport.
	FinancingDetails *financingdetails.FinancingDetails
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
	//TrdRegTimestamps is a non-required component for CollateralReport.
	TrdRegTimestamps *trdregtimestamps.TrdRegTimestamps
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
	//SpreadOrBenchmarkCurveData is a non-required component for CollateralReport.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations is a non-required component for CollateralReport.
	Stipulations *stipulations.Stipulations
	//SettlInstructionsData is a non-required component for CollateralReport.
	SettlInstructionsData *settlinstructionsdata.SettlInstructionsData
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
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized CollateralReport instance
func New(collrptid string, collstatus int) *Message {
	var m Message
	m.SetCollRptID(collrptid)
	m.SetCollStatus(collstatus)
	return &m
}

func (m *Message) SetCollRptID(v string)                                   { m.CollRptID = v }
func (m *Message) SetCollInquiryID(v string)                               { m.CollInquiryID = &v }
func (m *Message) SetCollStatus(v int)                                     { m.CollStatus = v }
func (m *Message) SetTotNumReports(v int)                                  { m.TotNumReports = &v }
func (m *Message) SetLastRptRequested(v bool)                              { m.LastRptRequested = &v }
func (m *Message) SetParties(v parties.Parties)                            { m.Parties = &v }
func (m *Message) SetAccount(v string)                                     { m.Account = &v }
func (m *Message) SetAccountType(v int)                                    { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                                     { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                                     { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)                            { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                            { m.SecondaryClOrdID = &v }
func (m *Message) SetNoExecs(v []NoExecs)                                  { m.NoExecs = v }
func (m *Message) SetNoTrades(v []NoTrades)                                { m.NoTrades = v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = &v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetSettlDate(v string)                                   { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)                                   { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                                    { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                      { m.NoUnderlyings = v }
func (m *Message) SetMarginExcess(v float64)                               { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                              { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                            { m.CashOutstanding = &v }
func (m *Message) SetTrdRegTimestamps(v trdregtimestamps.TrdRegTimestamps) { m.TrdRegTimestamps = &v }
func (m *Message) SetSide(v string)                                        { m.Side = &v }
func (m *Message) SetNoMiscFees(v []NoMiscFees)                            { m.NoMiscFees = v }
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
	return enum.BeginStringFIX44, "BA", r
}
