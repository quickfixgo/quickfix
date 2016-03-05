//Package securitylist msg type = y.
package securitylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
)

//NoRelatedSym is a repeating group in SecurityList
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
	//InstrumentExtension Component
	instrumentextension.InstrumentExtension
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//NoUnderlyings is a non-required field for NoRelatedSym.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
	//Stipulations Component
	stipulations.Stipulations
	//NoLegs is a non-required field for NoRelatedSym.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData Component
	yielddata.YieldData
	//RoundLot is a non-required field for NoRelatedSym.
	RoundLot *float64 `fix:"561"`
	//MinTradeVol is a non-required field for NoRelatedSym.
	MinTradeVol *float64 `fix:"562"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
	//ExpirationCycle is a non-required field for NoRelatedSym.
	ExpirationCycle *int `fix:"827"`
	//Text is a non-required field for NoRelatedSym.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoRelatedSym.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoRelatedSym.
	EncodedText *string `fix:"355"`
}

func (m *NoRelatedSym) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *NoRelatedSym) SetCurrency(v string)               { m.Currency = &v }
func (m *NoRelatedSym) SetNoLegs(v []NoLegs)               { m.NoLegs = v }
func (m *NoRelatedSym) SetRoundLot(v float64)              { m.RoundLot = &v }
func (m *NoRelatedSym) SetMinTradeVol(v float64)           { m.MinTradeVol = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)       { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string)    { m.TradingSessionSubID = &v }
func (m *NoRelatedSym) SetExpirationCycle(v int)           { m.ExpirationCycle = &v }
func (m *NoRelatedSym) SetText(v string)                   { m.Text = &v }
func (m *NoRelatedSym) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *NoRelatedSym) SetEncodedText(v string)            { m.EncodedText = &v }

//NoUnderlyings is a repeating group in NoRelatedSym
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in NoRelatedSym
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegStipulations Component
	legstipulations.LegStipulations
	//LegBenchmarkCurveData Component
	legbenchmarkcurvedata.LegBenchmarkCurveData
}

func (m *NoLegs) SetLegSwapType(v int)     { m.LegSwapType = &v }
func (m *NoLegs) SetLegSettlType(v string) { m.LegSettlType = &v }

//Message is a SecurityList FIX Message
type Message struct {
	FIXMsgType string `fix:"y"`
	fix44.Header
	//SecurityReqID is a required field for SecurityList.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for SecurityList.
	SecurityResponseID string `fix:"322"`
	//SecurityRequestResult is a required field for SecurityList.
	SecurityRequestResult int `fix:"560"`
	//TotNoRelatedSym is a non-required field for SecurityList.
	TotNoRelatedSym *int `fix:"393"`
	//LastFragment is a non-required field for SecurityList.
	LastFragment *bool `fix:"893"`
	//NoRelatedSym is a non-required field for SecurityList.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = v }
func (m *Message) SetSecurityResponseID(v string)   { m.SecurityResponseID = v }
func (m *Message) SetSecurityRequestResult(v int)   { m.SecurityRequestResult = v }
func (m *Message) SetTotNoRelatedSym(v int)         { m.TotNoRelatedSym = &v }
func (m *Message) SetLastFragment(v bool)           { m.LastFragment = &v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX44, "y", r
}
