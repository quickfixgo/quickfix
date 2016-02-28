//Package positionmaintenancerequest msg type = AL.
package positionmaintenancerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/positionqty"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoLegs is a repeating group in PositionMaintenanceRequest
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in PositionMaintenanceRequest
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoTradingSessions is a repeating group in PositionMaintenanceRequest
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a PositionMaintenanceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AL"`
	Header     fix44.Header
	//PosReqID is a required field for PositionMaintenanceRequest.
	PosReqID string `fix:"710"`
	//PosTransType is a required field for PositionMaintenanceRequest.
	PosTransType int `fix:"709"`
	//PosMaintAction is a required field for PositionMaintenanceRequest.
	PosMaintAction int `fix:"712"`
	//OrigPosReqRefID is a non-required field for PositionMaintenanceRequest.
	OrigPosReqRefID *string `fix:"713"`
	//PosMaintRptRefID is a non-required field for PositionMaintenanceRequest.
	PosMaintRptRefID *string `fix:"714"`
	//ClearingBusinessDate is a required field for PositionMaintenanceRequest.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for PositionMaintenanceRequest.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for PositionMaintenanceRequest.
	SettlSessSubID *string `fix:"717"`
	//Parties Component
	Parties parties.Component
	//Account is a required field for PositionMaintenanceRequest.
	Account string `fix:"1"`
	//AcctIDSource is a non-required field for PositionMaintenanceRequest.
	AcctIDSource *int `fix:"660"`
	//AccountType is a required field for PositionMaintenanceRequest.
	AccountType int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for PositionMaintenanceRequest.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for PositionMaintenanceRequest.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for PositionMaintenanceRequest.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoTradingSessions is a non-required field for PositionMaintenanceRequest.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//TransactTime is a required field for PositionMaintenanceRequest.
	TransactTime time.Time `fix:"60"`
	//PositionQty Component
	PositionQty positionqty.Component
	//AdjustmentType is a non-required field for PositionMaintenanceRequest.
	AdjustmentType *int `fix:"718"`
	//ContraryInstructionIndicator is a non-required field for PositionMaintenanceRequest.
	ContraryInstructionIndicator *bool `fix:"719"`
	//PriorSpreadIndicator is a non-required field for PositionMaintenanceRequest.
	PriorSpreadIndicator *bool `fix:"720"`
	//ThresholdAmount is a non-required field for PositionMaintenanceRequest.
	ThresholdAmount *float64 `fix:"834"`
	//Text is a non-required field for PositionMaintenanceRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PositionMaintenanceRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PositionMaintenanceRequest.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetPosReqID(v string)                       { m.PosReqID = v }
func (m *Message) SetPosTransType(v int)                      { m.PosTransType = v }
func (m *Message) SetPosMaintAction(v int)                    { m.PosMaintAction = v }
func (m *Message) SetOrigPosReqRefID(v string)                { m.OrigPosReqRefID = &v }
func (m *Message) SetPosMaintRptRefID(v string)               { m.PosMaintRptRefID = &v }
func (m *Message) SetClearingBusinessDate(v string)           { m.ClearingBusinessDate = v }
func (m *Message) SetSettlSessID(v string)                    { m.SettlSessID = &v }
func (m *Message) SetSettlSessSubID(v string)                 { m.SettlSessSubID = &v }
func (m *Message) SetAccount(v string)                        { m.Account = v }
func (m *Message) SetAcctIDSource(v int)                      { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                       { m.AccountType = v }
func (m *Message) SetCurrency(v string)                       { m.Currency = &v }
func (m *Message) SetNoLegs(v []NoLegs)                       { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetAdjustmentType(v int)                    { m.AdjustmentType = &v }
func (m *Message) SetContraryInstructionIndicator(v bool)     { m.ContraryInstructionIndicator = &v }
func (m *Message) SetPriorSpreadIndicator(v bool)             { m.PriorSpreadIndicator = &v }
func (m *Message) SetThresholdAmount(v float64)               { m.ThresholdAmount = &v }
func (m *Message) SetText(v string)                           { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                    { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                    { m.EncodedText = &v }

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
	return enum.BeginStringFIX44, "AL", r
}
