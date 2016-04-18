//Package positionreport msg type = AP.
package positionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/positionamountdata"
	"github.com/quickfixgo/quickfix/fix44/positionqty"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoLegs is a repeating group in PositionReport
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

//NoUnderlyings is a repeating group in PositionReport
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//UnderlyingSettlPrice is a required field for NoUnderlyings.
	UnderlyingSettlPrice float64 `fix:"732"`
	//UnderlyingSettlPriceType is a required field for NoUnderlyings.
	UnderlyingSettlPriceType int `fix:"733"`
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings(underlyingsettlprice float64, underlyingsettlpricetype int) *NoUnderlyings {
	var m NoUnderlyings
	m.SetUnderlyingSettlPrice(underlyingsettlprice)
	m.SetUnderlyingSettlPriceType(underlyingsettlpricetype)
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *NoUnderlyings) SetUnderlyingSettlPrice(v float64) { m.UnderlyingSettlPrice = v }
func (m *NoUnderlyings) SetUnderlyingSettlPriceType(v int) { m.UnderlyingSettlPriceType = v }

//Message is a PositionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AP"`
	fix44.Header
	//PosMaintRptID is a required field for PositionReport.
	PosMaintRptID string `fix:"721"`
	//PosReqID is a non-required field for PositionReport.
	PosReqID *string `fix:"710"`
	//PosReqType is a non-required field for PositionReport.
	PosReqType *int `fix:"724"`
	//SubscriptionRequestType is a non-required field for PositionReport.
	SubscriptionRequestType *string `fix:"263"`
	//TotalNumPosReports is a non-required field for PositionReport.
	TotalNumPosReports *int `fix:"727"`
	//UnsolicitedIndicator is a non-required field for PositionReport.
	UnsolicitedIndicator *bool `fix:"325"`
	//PosReqResult is a required field for PositionReport.
	PosReqResult int `fix:"728"`
	//ClearingBusinessDate is a required field for PositionReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for PositionReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for PositionReport.
	SettlSessSubID *string `fix:"717"`
	//Parties is a required component for PositionReport.
	parties.Parties
	//Account is a required field for PositionReport.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for PositionReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for PositionReport.
	AccountType int `fix:"581"`
	//Instrument is a non-required component for PositionReport.
	Instrument *instrument.Instrument
	//Currency is a non-required field for PositionReport.
	Currency *string `fix:"15"`
	//SettlPrice is a required field for PositionReport.
	SettlPrice float64 `fix:"730"`
	//SettlPriceType is a required field for PositionReport.
	SettlPriceType int `fix:"731"`
	//PriorSettlPrice is a required field for PositionReport.
	PriorSettlPrice float64 `fix:"734"`
	//NoLegs is a non-required field for PositionReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for PositionReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PositionQty is a required component for PositionReport.
	positionqty.PositionQty
	//PositionAmountData is a required component for PositionReport.
	positionamountdata.PositionAmountData
	//RegistStatus is a non-required field for PositionReport.
	RegistStatus *string `fix:"506"`
	//DeliveryDate is a non-required field for PositionReport.
	DeliveryDate *string `fix:"743"`
	//Text is a non-required field for PositionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PositionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PositionReport.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized PositionReport instance
func New(posmaintrptid string, posreqresult int, clearingbusinessdate string, parties parties.Parties, account string, accounttype int, settlprice float64, settlpricetype int, priorsettlprice float64, positionqty positionqty.PositionQty, positionamountdata positionamountdata.PositionAmountData) *Message {
	var m Message
	m.SetPosMaintRptID(posmaintrptid)
	m.SetPosReqResult(posreqresult)
	m.SetClearingBusinessDate(clearingbusinessdate)
	m.SetParties(parties)
	m.SetAccount(account)
	m.SetAccountType(accounttype)
	m.SetSettlPrice(settlprice)
	m.SetSettlPriceType(settlpricetype)
	m.SetPriorSettlPrice(priorsettlprice)
	m.SetPositionQty(positionqty)
	m.SetPositionAmountData(positionamountdata)
	return &m
}

func (m *Message) SetPosMaintRptID(v string)                { m.PosMaintRptID = v }
func (m *Message) SetPosReqID(v string)                     { m.PosReqID = &v }
func (m *Message) SetPosReqType(v int)                      { m.PosReqType = &v }
func (m *Message) SetSubscriptionRequestType(v string)      { m.SubscriptionRequestType = &v }
func (m *Message) SetTotalNumPosReports(v int)              { m.TotalNumPosReports = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)           { m.UnsolicitedIndicator = &v }
func (m *Message) SetPosReqResult(v int)                    { m.PosReqResult = v }
func (m *Message) SetClearingBusinessDate(v string)         { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                  { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)               { m.SettlSessSubID = &v }
func (m *Message) SetParties(v parties.Parties)             { m.Parties = v }
func (m *Message) SetAccount(v string)                      { m.Account = v }
func (m *Message) SetAcctIDSource(v int)                    { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                     { m.AccountType = v }
func (m *Message) SetInstrument(v instrument.Instrument)    { m.Instrument = &v }
func (m *Message) SetCurrency(v string)                     { m.Currency = &v }
func (m *Message) SetSettlPrice(v float64)                  { m.SettlPrice = v }
func (m *Message) SetSettlPriceType(v int)                  { m.SettlPriceType = v }
func (m *Message) SetPriorSettlPrice(v float64)             { m.PriorSettlPrice = v }
func (m *Message) SetNoLegs(v []NoLegs)                     { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)       { m.NoUnderlyings = v }
func (m *Message) SetPositionQty(v positionqty.PositionQty) { m.PositionQty = v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = v
}
func (m *Message) SetRegistStatus(v string) { m.RegistStatus = &v }
func (m *Message) SetDeliveryDate(v string) { m.DeliveryDate = &v }
func (m *Message) SetText(v string)         { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)  { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)  { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AP", r
}
