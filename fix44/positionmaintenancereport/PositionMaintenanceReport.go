//Package positionmaintenancereport msg type = AM.
package positionmaintenancereport

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
	"time"
)

//NoLegs is a repeating group in PositionMaintenanceReport
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

//NoUnderlyings is a repeating group in PositionMaintenanceReport
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

//NoTradingSessions is a repeating group in PositionMaintenanceReport
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//NewNoTradingSessions returns an initialized NoTradingSessions instance
func NewNoTradingSessions() *NoTradingSessions {
	var m NoTradingSessions
	return &m
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a PositionMaintenanceReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AM"`
	fix44.Header
	//PosMaintRptID is a required field for PositionMaintenanceReport.
	PosMaintRptID string `fix:"721"`
	//PosTransType is a required field for PositionMaintenanceReport.
	PosTransType int `fix:"709"`
	//PosReqID is a non-required field for PositionMaintenanceReport.
	PosReqID *string `fix:"710"`
	//PosMaintAction is a required field for PositionMaintenanceReport.
	PosMaintAction int `fix:"712"`
	//OrigPosReqRefID is a required field for PositionMaintenanceReport.
	OrigPosReqRefID string `fix:"713"`
	//PosMaintStatus is a required field for PositionMaintenanceReport.
	PosMaintStatus int `fix:"722"`
	//PosMaintResult is a non-required field for PositionMaintenanceReport.
	PosMaintResult *int `fix:"723"`
	//ClearingBusinessDate is a required field for PositionMaintenanceReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for PositionMaintenanceReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for PositionMaintenanceReport.
	SettlSessSubID *string `fix:"717"`
	//Parties is a non-required component for PositionMaintenanceReport.
	Parties *parties.Parties
	//Account is a required field for PositionMaintenanceReport.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for PositionMaintenanceReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for PositionMaintenanceReport.
	AccountType int `fix:"581"`
	//Instrument is a required component for PositionMaintenanceReport.
	instrument.Instrument
	//Currency is a non-required field for PositionMaintenanceReport.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for PositionMaintenanceReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for PositionMaintenanceReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoTradingSessions is a non-required field for PositionMaintenanceReport.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//TransactTime is a required field for PositionMaintenanceReport.
	TransactTime time.Time `fix:"60"`
	//PositionQty is a required component for PositionMaintenanceReport.
	positionqty.PositionQty
	//PositionAmountData is a required component for PositionMaintenanceReport.
	positionamountdata.PositionAmountData
	//AdjustmentType is a non-required field for PositionMaintenanceReport.
	AdjustmentType *int `fix:"718"`
	//ThresholdAmount is a non-required field for PositionMaintenanceReport.
	ThresholdAmount *float64 `fix:"834"`
	//Text is a non-required field for PositionMaintenanceReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PositionMaintenanceReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PositionMaintenanceReport.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosMaintRptID(v string)                  { m.PosMaintRptID = v }
func (m *Message) SetPosTransType(v int)                      { m.PosTransType = v }
func (m *Message) SetPosReqID(v string)                       { m.PosReqID = &v }
func (m *Message) SetPosMaintAction(v int)                    { m.PosMaintAction = v }
func (m *Message) SetOrigPosReqRefID(v string)                { m.OrigPosReqRefID = v }
func (m *Message) SetPosMaintStatus(v int)                    { m.PosMaintStatus = v }
func (m *Message) SetPosMaintResult(v int)                    { m.PosMaintResult = &v }
func (m *Message) SetClearingBusinessDate(v string)           { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                    { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                 { m.SettlSessSubID = &v }
func (m *Message) SetParties(v parties.Parties)               { m.Parties = &v }
func (m *Message) SetAccount(v string)                        { m.Account = v }
func (m *Message) SetAcctIDSource(v int)                      { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                       { m.AccountType = v }
func (m *Message) SetInstrument(v instrument.Instrument)      { m.Instrument = v }
func (m *Message) SetCurrency(v string)                       { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                       { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetPositionQty(v positionqty.PositionQty)   { m.PositionQty = v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = v
}
func (m *Message) SetAdjustmentType(v int)      { m.AdjustmentType = &v }
func (m *Message) SetThresholdAmount(v float64) { m.ThresholdAmount = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AM", r
}
