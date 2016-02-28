//Package tradecapturereportack msg type = AR.
package tradecapturereportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/legstipulations"
	"github.com/quickfixgo/quickfix/fix44/nestedparties"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/trdregtimestamps"
	"time"
)

//NoLegs is a repeating group in TradeCaptureReportAck
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
	//LegQty is a non-required field for NoLegs.
	LegQty *float64 `fix:"687"`
	//LegSwapType is a non-required field for NoLegs.
	LegSwapType *int `fix:"690"`
	//LegStipulations Component
	LegStipulations legstipulations.Component
	//LegPositionEffect is a non-required field for NoLegs.
	LegPositionEffect *string `fix:"564"`
	//LegCoveredOrUncovered is a non-required field for NoLegs.
	LegCoveredOrUncovered *int `fix:"565"`
	//NestedParties Component
	NestedParties nestedparties.Component
	//LegRefID is a non-required field for NoLegs.
	LegRefID *string `fix:"654"`
	//LegPrice is a non-required field for NoLegs.
	LegPrice *float64 `fix:"566"`
	//LegSettlType is a non-required field for NoLegs.
	LegSettlType *string `fix:"587"`
	//LegSettlDate is a non-required field for NoLegs.
	LegSettlDate *string `fix:"588"`
	//LegLastPx is a non-required field for NoLegs.
	LegLastPx *float64 `fix:"637"`
}

func (m *NoLegs) SetLegQty(v float64)            { m.LegQty = &v }
func (m *NoLegs) SetLegSwapType(v int)           { m.LegSwapType = &v }
func (m *NoLegs) SetLegPositionEffect(v string)  { m.LegPositionEffect = &v }
func (m *NoLegs) SetLegCoveredOrUncovered(v int) { m.LegCoveredOrUncovered = &v }
func (m *NoLegs) SetLegRefID(v string)           { m.LegRefID = &v }
func (m *NoLegs) SetLegPrice(v float64)          { m.LegPrice = &v }
func (m *NoLegs) SetLegSettlType(v string)       { m.LegSettlType = &v }
func (m *NoLegs) SetLegSettlDate(v string)       { m.LegSettlDate = &v }
func (m *NoLegs) SetLegLastPx(v float64)         { m.LegLastPx = &v }

//NoAllocs is a repeating group in TradeCaptureReportAck
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//NestedParties2 Component
	NestedParties2 nestedparties2.Component
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
}

func (m *NoAllocs) SetAllocAccount(v string)       { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)     { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string) { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)  { m.IndividualAllocID = &v }
func (m *NoAllocs) SetAllocQty(v float64)          { m.AllocQty = &v }

//Message is a TradeCaptureReportAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AR"`
	Header     fix44.Header
	//TradeReportID is a required field for TradeCaptureReportAck.
	TradeReportID string `fix:"571"`
	//TradeReportTransType is a non-required field for TradeCaptureReportAck.
	TradeReportTransType *int `fix:"487"`
	//TradeReportType is a non-required field for TradeCaptureReportAck.
	TradeReportType *int `fix:"856"`
	//TrdType is a non-required field for TradeCaptureReportAck.
	TrdType *int `fix:"828"`
	//TrdSubType is a non-required field for TradeCaptureReportAck.
	TrdSubType *int `fix:"829"`
	//SecondaryTrdType is a non-required field for TradeCaptureReportAck.
	SecondaryTrdType *int `fix:"855"`
	//TransferReason is a non-required field for TradeCaptureReportAck.
	TransferReason *string `fix:"830"`
	//ExecType is a required field for TradeCaptureReportAck.
	ExecType string `fix:"150"`
	//TradeReportRefID is a non-required field for TradeCaptureReportAck.
	TradeReportRefID *string `fix:"572"`
	//SecondaryTradeReportRefID is a non-required field for TradeCaptureReportAck.
	SecondaryTradeReportRefID *string `fix:"881"`
	//TrdRptStatus is a non-required field for TradeCaptureReportAck.
	TrdRptStatus *int `fix:"939"`
	//TradeReportRejectReason is a non-required field for TradeCaptureReportAck.
	TradeReportRejectReason *int `fix:"751"`
	//SecondaryTradeReportID is a non-required field for TradeCaptureReportAck.
	SecondaryTradeReportID *string `fix:"818"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportAck.
	SubscriptionRequestType *string `fix:"263"`
	//TradeLinkID is a non-required field for TradeCaptureReportAck.
	TradeLinkID *string `fix:"820"`
	//TrdMatchID is a non-required field for TradeCaptureReportAck.
	TrdMatchID *string `fix:"880"`
	//ExecID is a non-required field for TradeCaptureReportAck.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for TradeCaptureReportAck.
	SecondaryExecID *string `fix:"527"`
	//Instrument Component
	Instrument instrument.Component
	//TransactTime is a non-required field for TradeCaptureReportAck.
	TransactTime *time.Time `fix:"60"`
	//TrdRegTimestamps Component
	TrdRegTimestamps trdregtimestamps.Component
	//ResponseTransportType is a non-required field for TradeCaptureReportAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for TradeCaptureReportAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for TradeCaptureReportAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportAck.
	EncodedText *string `fix:"355"`
	//NoLegs is a non-required field for TradeCaptureReportAck.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//ClearingFeeIndicator is a non-required field for TradeCaptureReportAck.
	ClearingFeeIndicator *string `fix:"635"`
	//OrderCapacity is a non-required field for TradeCaptureReportAck.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for TradeCaptureReportAck.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for TradeCaptureReportAck.
	CustOrderCapacity *int `fix:"582"`
	//Account is a non-required field for TradeCaptureReportAck.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for TradeCaptureReportAck.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for TradeCaptureReportAck.
	AccountType *int `fix:"581"`
	//PositionEffect is a non-required field for TradeCaptureReportAck.
	PositionEffect *string `fix:"77"`
	//PreallocMethod is a non-required field for TradeCaptureReportAck.
	PreallocMethod *string `fix:"591"`
	//NoAllocs is a non-required field for TradeCaptureReportAck.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	Trailer  fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradeReportID(v string)             { m.TradeReportID = v }
func (m *Message) SetTradeReportTransType(v int)         { m.TradeReportTransType = &v }
func (m *Message) SetTradeReportType(v int)              { m.TradeReportType = &v }
func (m *Message) SetTrdType(v int)                      { m.TrdType = &v }
func (m *Message) SetTrdSubType(v int)                   { m.TrdSubType = &v }
func (m *Message) SetSecondaryTrdType(v int)             { m.SecondaryTrdType = &v }
func (m *Message) SetTransferReason(v string)            { m.TransferReason = &v }
func (m *Message) SetExecType(v string)                  { m.ExecType = v }
func (m *Message) SetTradeReportRefID(v string)          { m.TradeReportRefID = &v }
func (m *Message) SetSecondaryTradeReportRefID(v string) { m.SecondaryTradeReportRefID = &v }
func (m *Message) SetTrdRptStatus(v int)                 { m.TrdRptStatus = &v }
func (m *Message) SetTradeReportRejectReason(v int)      { m.TradeReportRejectReason = &v }
func (m *Message) SetSecondaryTradeReportID(v string)    { m.SecondaryTradeReportID = &v }
func (m *Message) SetSubscriptionRequestType(v string)   { m.SubscriptionRequestType = &v }
func (m *Message) SetTradeLinkID(v string)               { m.TradeLinkID = &v }
func (m *Message) SetTrdMatchID(v string)                { m.TrdMatchID = &v }
func (m *Message) SetExecID(v string)                    { m.ExecID = &v }
func (m *Message) SetSecondaryExecID(v string)           { m.SecondaryExecID = &v }
func (m *Message) SetTransactTime(v time.Time)           { m.TransactTime = &v }
func (m *Message) SetResponseTransportType(v int)        { m.ResponseTransportType = &v }
func (m *Message) SetResponseDestination(v string)       { m.ResponseDestination = &v }
func (m *Message) SetText(v string)                      { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)               { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)               { m.EncodedText = &v }
func (m *Message) SetNoLegs(v []NoLegs)                  { m.NoLegs = v }
func (m *Message) SetClearingFeeIndicator(v string)      { m.ClearingFeeIndicator = &v }
func (m *Message) SetOrderCapacity(v string)             { m.OrderCapacity = &v }
func (m *Message) SetOrderRestrictions(v string)         { m.OrderRestrictions = &v }
func (m *Message) SetCustOrderCapacity(v int)            { m.CustOrderCapacity = &v }
func (m *Message) SetAccount(v string)                   { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)                 { m.AcctIDSource = &v }
func (m *Message) SetAccountType(v int)                  { m.AccountType = &v }
func (m *Message) SetPositionEffect(v string)            { m.PositionEffect = &v }
func (m *Message) SetPreallocMethod(v string)            { m.PreallocMethod = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)              { m.NoAllocs = v }

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
	return enum.BeginStringFIX44, "AR", r
}
