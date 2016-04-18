//Package assignmentreport msg type = AW.
package assignmentreport

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

//NoLegs is a repeating group in AssignmentReport
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

//NoUnderlyings is a repeating group in AssignmentReport
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

//Message is a AssignmentReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AW"`
	fix44.Header
	//AsgnRptID is a required field for AssignmentReport.
	AsgnRptID string `fix:"833"`
	//TotNumAssignmentReports is a non-required field for AssignmentReport.
	TotNumAssignmentReports *int `fix:"832"`
	//LastRptRequested is a non-required field for AssignmentReport.
	LastRptRequested *bool `fix:"912"`
	//Parties is a required component for AssignmentReport.
	parties.Parties
	//Account is a non-required field for AssignmentReport.
	Account *string `fix:"1"`
	//AccountType is a required field for AssignmentReport.
	AccountType int `fix:"581"`
	//Instrument is a non-required component for AssignmentReport.
	Instrument *instrument.Instrument
	//Currency is a non-required field for AssignmentReport.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for AssignmentReport.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for AssignmentReport.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//PositionQty is a required component for AssignmentReport.
	positionqty.PositionQty
	//PositionAmountData is a required component for AssignmentReport.
	positionamountdata.PositionAmountData
	//ThresholdAmount is a non-required field for AssignmentReport.
	ThresholdAmount *float64 `fix:"834"`
	//SettlPrice is a required field for AssignmentReport.
	SettlPrice float64 `fix:"730"`
	//SettlPriceType is a required field for AssignmentReport.
	SettlPriceType int `fix:"731"`
	//UnderlyingSettlPrice is a required field for AssignmentReport.
	UnderlyingSettlPrice float64 `fix:"732"`
	//ExpireDate is a non-required field for AssignmentReport.
	ExpireDate *string `fix:"432"`
	//AssignmentMethod is a required field for AssignmentReport.
	AssignmentMethod string `fix:"744"`
	//AssignmentUnit is a non-required field for AssignmentReport.
	AssignmentUnit *float64 `fix:"745"`
	//OpenInterest is a required field for AssignmentReport.
	OpenInterest float64 `fix:"746"`
	//ExerciseMethod is a required field for AssignmentReport.
	ExerciseMethod string `fix:"747"`
	//SettlSessID is a required field for AssignmentReport.
	SettlSessID string `fix:"716"`
	//SettlSessSubID is a required field for AssignmentReport.
	SettlSessSubID string `fix:"717"`
	//ClearingBusinessDate is a required field for AssignmentReport.
	ClearingBusinessDate string `fix:"715"`
	//Text is a non-required field for AssignmentReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AssignmentReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AssignmentReport.
	EncodedText *string `fix:"355"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAsgnRptID(v string)                    { m.AsgnRptID = v }
func (m *Message) SetTotNumAssignmentReports(v int)         { m.TotNumAssignmentReports = &v }
func (m *Message) SetLastRptRequested(v bool)               { m.LastRptRequested = &v }
func (m *Message) SetParties(v parties.Parties)             { m.Parties = v }
func (m *Message) SetAccount(v string)                      { m.Account = &v }
func (m *Message) SetAccountType(v int)                     { m.AccountType = v }
func (m *Message) SetInstrument(v instrument.Instrument)    { m.Instrument = &v }
func (m *Message) SetCurrency(v string)                     { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                     { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)       { m.NoUnderlyings = v }
func (m *Message) SetPositionQty(v positionqty.PositionQty) { m.PositionQty = v }
func (m *Message) SetPositionAmountData(v positionamountdata.PositionAmountData) {
	m.PositionAmountData = v
}
func (m *Message) SetThresholdAmount(v float64)      { m.ThresholdAmount = &v }
func (m *Message) SetSettlPrice(v float64)           { m.SettlPrice = v }
func (m *Message) SetSettlPriceType(v int)           { m.SettlPriceType = v }
func (m *Message) SetUnderlyingSettlPrice(v float64) { m.UnderlyingSettlPrice = v }
func (m *Message) SetExpireDate(v string)            { m.ExpireDate = &v }
func (m *Message) SetAssignmentMethod(v string)      { m.AssignmentMethod = v }
func (m *Message) SetAssignmentUnit(v float64)       { m.AssignmentUnit = &v }
func (m *Message) SetOpenInterest(v float64)         { m.OpenInterest = v }
func (m *Message) SetExerciseMethod(v string)        { m.ExerciseMethod = v }
func (m *Message) SetSettlSessID(v string)           { m.SettlSessID = v }
func (m *Message) SetSettlSessSubID(v string)        { m.SettlSessSubID = v }
func (m *Message) SetClearingBusinessDate(v string)  { m.ClearingBusinessDate = v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AW", r
}
