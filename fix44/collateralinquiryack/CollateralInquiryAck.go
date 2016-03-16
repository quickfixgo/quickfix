//Package collateralinquiryack msg type = BG.
package collateralinquiryack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoCollInquiryQualifier is a repeating group in CollateralInquiryAck
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

func (m *NoCollInquiryQualifier) SetCollInquiryQualifier(v int) { m.CollInquiryQualifier = &v }

//NoExecs is a repeating group in CollateralInquiryAck
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

func (m *NoExecs) SetExecID(v string) { m.ExecID = &v }

//NoTrades is a repeating group in CollateralInquiryAck
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

func (m *NoTrades) SetTradeReportID(v string)          { m.TradeReportID = &v }
func (m *NoTrades) SetSecondaryTradeReportID(v string) { m.SecondaryTradeReportID = &v }

//NoLegs is a repeating group in CollateralInquiryAck
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//NoUnderlyings is a repeating group in CollateralInquiryAck
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//Message is a CollateralInquiryAck FIX Message
type Message struct {
	FIXMsgType string `fix:"BG"`
	fix44.Header
	//CollInquiryID is a required field for CollateralInquiryAck.
	CollInquiryID string `fix:"909"`
	//CollInquiryStatus is a required field for CollateralInquiryAck.
	CollInquiryStatus int `fix:"945"`
	//CollInquiryResult is a non-required field for CollateralInquiryAck.
	CollInquiryResult *int `fix:"946"`
	//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
	//TotNumReports is a non-required field for CollateralInquiryAck.
	TotNumReports *int `fix:"911"`
	//Parties is a non-required component for CollateralInquiryAck.
	Parties *parties.Parties
	//Account is a non-required field for CollateralInquiryAck.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralInquiryAck.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralInquiryAck.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralInquiryAck.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralInquiryAck.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
	SecondaryClOrdID *string `fix:"526"`
	//NoExecs is a non-required field for CollateralInquiryAck.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralInquiryAck.
	NoTrades []NoTrades `fix:"897,omitempty"`
	//Instrument is a non-required component for CollateralInquiryAck.
	Instrument *instrument.Instrument
	//FinancingDetails is a non-required component for CollateralInquiryAck.
	FinancingDetails *financingdetails.FinancingDetails
	//SettlDate is a non-required field for CollateralInquiryAck.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiryAck.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiryAck.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiryAck.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for CollateralInquiryAck.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralInquiryAck.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//TradingSessionID is a non-required field for CollateralInquiryAck.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralInquiryAck.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralInquiryAck.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralInquiryAck.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
	ClearingBusinessDate *string `fix:"715"`
	//ResponseTransportType is a non-required field for CollateralInquiryAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiryAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for CollateralInquiryAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralInquiryAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralInquiryAck.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetCollInquiryID(v string)                               { m.CollInquiryID = v }
func (m *Message) SetCollInquiryStatus(v int)                              { m.CollInquiryStatus = v }
func (m *Message) SetCollInquiryResult(v int)                              { m.CollInquiryResult = &v }
func (m *Message) SetNoCollInquiryQualifier(v []NoCollInquiryQualifier)    { m.NoCollInquiryQualifier = v }
func (m *Message) SetTotNumReports(v int)                                  { m.TotNumReports = &v }
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
func (m *Message) SetTradingSessionID(v string)                            { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                         { m.TradingSessionSubID = &v }
func (m *Message) SetSettlSessID(v string)                                 { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                              { m.SettlSessSubID = &v }
func (m *Message) SetClearingBusinessDate(v string)                        { m.ClearingBusinessDate = &v }
func (m *Message) SetResponseTransportType(v int)                          { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)                         { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                                        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                                 { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                                 { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "BG", r
}
