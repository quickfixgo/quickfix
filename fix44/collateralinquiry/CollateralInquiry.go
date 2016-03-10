//Package collateralinquiry msg type = BB.
package collateralinquiry

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

//NoCollInquiryQualifier is a repeating group in CollateralInquiry
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

func (m *NoCollInquiryQualifier) SetCollInquiryQualifier(v int) { m.CollInquiryQualifier = &v }

//NoExecs is a repeating group in CollateralInquiry
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

func (m *NoExecs) SetExecID(v string) { m.ExecID = &v }

//NoTrades is a repeating group in CollateralInquiry
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

func (m *NoTrades) SetTradeReportID(v string)          { m.TradeReportID = &v }
func (m *NoTrades) SetSecondaryTradeReportID(v string) { m.SecondaryTradeReportID = &v }

//NoLegs is a repeating group in CollateralInquiry
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//NoUnderlyings is a repeating group in CollateralInquiry
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//Message is a CollateralInquiry FIX Message
type Message struct {
	FIXMsgType string `fix:"BB"`
	fix44.Header
	//CollInquiryID is a non-required field for CollateralInquiry.
	CollInquiryID *string `fix:"909"`
	//NoCollInquiryQualifier is a non-required field for CollateralInquiry.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
	//SubscriptionRequestType is a non-required field for CollateralInquiry.
	SubscriptionRequestType *string `fix:"263"`
	//ResponseTransportType is a non-required field for CollateralInquiry.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiry.
	ResponseDestination *string `fix:"726"`
	//Parties Component
	parties.Parties
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
	//NoExecs is a non-required field for CollateralInquiry.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralInquiry.
	NoTrades []NoTrades `fix:"897,omitempty"`
	//Instrument Component
	instrument.Instrument
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralInquiry.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiry.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiry.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiry.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for CollateralInquiry.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralInquiry.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//MarginExcess is a non-required field for CollateralInquiry.
	MarginExcess *float64 `fix:"899"`
	//TotalNetValue is a non-required field for CollateralInquiry.
	TotalNetValue *float64 `fix:"900"`
	//CashOutstanding is a non-required field for CollateralInquiry.
	CashOutstanding *float64 `fix:"901"`
	//TrdRegTimestamps Component
	trdregtimestamps.TrdRegTimestamps
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
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//Stipulations Component
	stipulations.Stipulations
	//SettlInstructionsData Component
	settlinstructionsdata.SettlInstructionsData
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
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCollInquiryID(v string)                            { m.CollInquiryID = &v }
func (m *Message) SetNoCollInquiryQualifier(v []NoCollInquiryQualifier) { m.NoCollInquiryQualifier = v }
func (m *Message) SetSubscriptionRequestType(v string)                  { m.SubscriptionRequestType = &v }
func (m *Message) SetResponseTransportType(v int)                       { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)                      { m.ResponseDestination = &v }
func (m *Message) SetAccount(v string)                                  { m.Account = &v }
func (m *Message) SetAccountType(v int)                                 { m.AccountType = &v }
func (m *Message) SetClOrdID(v string)                                  { m.ClOrdID = &v }
func (m *Message) SetOrderID(v string)                                  { m.OrderID = &v }
func (m *Message) SetSecondaryOrderID(v string)                         { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                         { m.SecondaryClOrdID = &v }
func (m *Message) SetNoExecs(v []NoExecs)                               { m.NoExecs = v }
func (m *Message) SetNoTrades(v []NoTrades)                             { m.NoTrades = v }
func (m *Message) SetSettlDate(v string)                                { m.SettlDate = &v }
func (m *Message) SetQuantity(v float64)                                { m.Quantity = &v }
func (m *Message) SetQtyType(v int)                                     { m.QtyType = &v }
func (m *Message) SetCurrency(v string)                                 { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                                 { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)                   { m.NoUnderlyings = v }
func (m *Message) SetMarginExcess(v float64)                            { m.MarginExcess = &v }
func (m *Message) SetTotalNetValue(v float64)                           { m.TotalNetValue = &v }
func (m *Message) SetCashOutstanding(v float64)                         { m.CashOutstanding = &v }
func (m *Message) SetSide(v string)                                     { m.Side = &v }
func (m *Message) SetPrice(v float64)                                   { m.Price = &v }
func (m *Message) SetPriceType(v int)                                   { m.PriceType = &v }
func (m *Message) SetAccruedInterestAmt(v float64)                      { m.AccruedInterestAmt = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64)                   { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)                               { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)                                 { m.EndCash = &v }
func (m *Message) SetTradingSessionID(v string)                         { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                      { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)                              { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                           { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string)                     { m.ClearingBusinessDate = &v }
func (m *Message) SetText(v string)                                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                              { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "BB", r
}
