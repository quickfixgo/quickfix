//Package confirmation msg type = AK.
package confirmation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/commissiondata"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentextension"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/nestedparties2"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix44/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix44/stipulations"
	"github.com/quickfixgo/quickfix/fix44/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fix44/yielddata"
	"time"
)

//NoOrders is a repeating group in Confirmation
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//NestedParties2 Component
	nestedparties2.NestedParties2
	//OrderQty is a non-required field for NoOrders.
	OrderQty *float64 `fix:"38"`
	//OrderAvgPx is a non-required field for NoOrders.
	OrderAvgPx *float64 `fix:"799"`
	//OrderBookingQty is a non-required field for NoOrders.
	OrderBookingQty *float64 `fix:"800"`
}

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }
func (m *NoOrders) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListID(v string)           { m.ListID = &v }
func (m *NoOrders) SetOrderQty(v float64)        { m.OrderQty = &v }
func (m *NoOrders) SetOrderAvgPx(v float64)      { m.OrderAvgPx = &v }
func (m *NoOrders) SetOrderBookingQty(v float64) { m.OrderBookingQty = &v }

//NoUnderlyings is a repeating group in Confirmation
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
}

//NoLegs is a repeating group in Confirmation
type NoLegs struct {
	//InstrumentLeg Component
	instrumentleg.InstrumentLeg
}

//NoCapacities is a repeating group in Confirmation
type NoCapacities struct {
	//OrderCapacity is a required field for NoCapacities.
	OrderCapacity string `fix:"528"`
	//OrderRestrictions is a non-required field for NoCapacities.
	OrderRestrictions *string `fix:"529"`
	//OrderCapacityQty is a required field for NoCapacities.
	OrderCapacityQty float64 `fix:"863"`
}

func (m *NoCapacities) SetOrderCapacity(v string)     { m.OrderCapacity = v }
func (m *NoCapacities) SetOrderRestrictions(v string) { m.OrderRestrictions = &v }
func (m *NoCapacities) SetOrderCapacityQty(v float64) { m.OrderCapacityQty = v }

//NoMiscFees is a repeating group in Confirmation
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

//Message is a Confirmation FIX Message
type Message struct {
	FIXMsgType string `fix:"AK"`
	fix44.Header
	//ConfirmID is a required field for Confirmation.
	ConfirmID string `fix:"664"`
	//ConfirmRefID is a non-required field for Confirmation.
	ConfirmRefID *string `fix:"772"`
	//ConfirmReqID is a non-required field for Confirmation.
	ConfirmReqID *string `fix:"859"`
	//ConfirmTransType is a required field for Confirmation.
	ConfirmTransType int `fix:"666"`
	//ConfirmType is a required field for Confirmation.
	ConfirmType int `fix:"773"`
	//CopyMsgIndicator is a non-required field for Confirmation.
	CopyMsgIndicator *bool `fix:"797"`
	//LegalConfirm is a non-required field for Confirmation.
	LegalConfirm *bool `fix:"650"`
	//ConfirmStatus is a required field for Confirmation.
	ConfirmStatus int `fix:"665"`
	//Parties Component
	parties.Parties
	//NoOrders is a non-required field for Confirmation.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//AllocID is a non-required field for Confirmation.
	AllocID *string `fix:"70"`
	//SecondaryAllocID is a non-required field for Confirmation.
	SecondaryAllocID *string `fix:"793"`
	//IndividualAllocID is a non-required field for Confirmation.
	IndividualAllocID *string `fix:"467"`
	//TransactTime is a required field for Confirmation.
	TransactTime time.Time `fix:"60"`
	//TradeDate is a required field for Confirmation.
	TradeDate string `fix:"75"`
	//TrdRegTimestamps Component
	trdregtimestamps.TrdRegTimestamps
	//Instrument Component
	instrument.Instrument
	//InstrumentExtension Component
	instrumentextension.InstrumentExtension
	//FinancingDetails Component
	financingdetails.FinancingDetails
	//NoUnderlyings is a required field for Confirmation.
	NoUnderlyings []NoUnderlyings `fix:"711"`
	//NoLegs is a required field for Confirmation.
	NoLegs []NoLegs `fix:"555"`
	//YieldData Component
	yielddata.YieldData
	//AllocQty is a required field for Confirmation.
	AllocQty float64 `fix:"80"`
	//QtyType is a non-required field for Confirmation.
	QtyType *int `fix:"854"`
	//Side is a required field for Confirmation.
	Side string `fix:"54"`
	//Currency is a non-required field for Confirmation.
	Currency *string `fix:"15"`
	//LastMkt is a non-required field for Confirmation.
	LastMkt *string `fix:"30"`
	//NoCapacities is a required field for Confirmation.
	NoCapacities []NoCapacities `fix:"862"`
	//AllocAccount is a required field for Confirmation.
	AllocAccount string `fix:"79"`
	//AllocAcctIDSource is a non-required field for Confirmation.
	AllocAcctIDSource *int `fix:"661"`
	//AllocAccountType is a non-required field for Confirmation.
	AllocAccountType *int `fix:"798"`
	//AvgPx is a required field for Confirmation.
	AvgPx float64 `fix:"6"`
	//AvgPxPrecision is a non-required field for Confirmation.
	AvgPxPrecision *int `fix:"74"`
	//PriceType is a non-required field for Confirmation.
	PriceType *int `fix:"423"`
	//AvgParPx is a non-required field for Confirmation.
	AvgParPx *float64 `fix:"860"`
	//SpreadOrBenchmarkCurveData Component
	spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//ReportedPx is a non-required field for Confirmation.
	ReportedPx *float64 `fix:"861"`
	//Text is a non-required field for Confirmation.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Confirmation.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Confirmation.
	EncodedText *string `fix:"355"`
	//ProcessCode is a non-required field for Confirmation.
	ProcessCode *string `fix:"81"`
	//GrossTradeAmt is a required field for Confirmation.
	GrossTradeAmt float64 `fix:"381"`
	//NumDaysInterest is a non-required field for Confirmation.
	NumDaysInterest *int `fix:"157"`
	//ExDate is a non-required field for Confirmation.
	ExDate *string `fix:"230"`
	//AccruedInterestRate is a non-required field for Confirmation.
	AccruedInterestRate *float64 `fix:"158"`
	//AccruedInterestAmt is a non-required field for Confirmation.
	AccruedInterestAmt *float64 `fix:"159"`
	//InterestAtMaturity is a non-required field for Confirmation.
	InterestAtMaturity *float64 `fix:"738"`
	//EndAccruedInterestAmt is a non-required field for Confirmation.
	EndAccruedInterestAmt *float64 `fix:"920"`
	//StartCash is a non-required field for Confirmation.
	StartCash *float64 `fix:"921"`
	//EndCash is a non-required field for Confirmation.
	EndCash *float64 `fix:"922"`
	//Concession is a non-required field for Confirmation.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for Confirmation.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a required field for Confirmation.
	NetMoney float64 `fix:"118"`
	//MaturityNetMoney is a non-required field for Confirmation.
	MaturityNetMoney *float64 `fix:"890"`
	//SettlCurrAmt is a non-required field for Confirmation.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for Confirmation.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for Confirmation.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for Confirmation.
	SettlCurrFxRateCalc *string `fix:"156"`
	//SettlType is a non-required field for Confirmation.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for Confirmation.
	SettlDate *string `fix:"64"`
	//SettlInstructionsData Component
	settlinstructionsdata.SettlInstructionsData
	//CommissionData Component
	commissiondata.CommissionData
	//SharedCommission is a non-required field for Confirmation.
	SharedCommission *float64 `fix:"858"`
	//Stipulations Component
	stipulations.Stipulations
	//NoMiscFees is a non-required field for Confirmation.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetConfirmID(v string)              { m.ConfirmID = v }
func (m *Message) SetConfirmRefID(v string)           { m.ConfirmRefID = &v }
func (m *Message) SetConfirmReqID(v string)           { m.ConfirmReqID = &v }
func (m *Message) SetConfirmTransType(v int)          { m.ConfirmTransType = v }
func (m *Message) SetConfirmType(v int)               { m.ConfirmType = v }
func (m *Message) SetCopyMsgIndicator(v bool)         { m.CopyMsgIndicator = &v }
func (m *Message) SetLegalConfirm(v bool)             { m.LegalConfirm = &v }
func (m *Message) SetConfirmStatus(v int)             { m.ConfirmStatus = v }
func (m *Message) SetNoOrders(v []NoOrders)           { m.NoOrders = v }
func (m *Message) SetAllocID(v string)                { m.AllocID = &v }
func (m *Message) SetSecondaryAllocID(v string)       { m.SecondaryAllocID = &v }
func (m *Message) SetIndividualAllocID(v string)      { m.IndividualAllocID = &v }
func (m *Message) SetTransactTime(v time.Time)        { m.TransactTime = v }
func (m *Message) SetTradeDate(v string)              { m.TradeDate = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *Message) SetNoLegs(v []NoLegs)               { m.NoLegs = v }
func (m *Message) SetAllocQty(v float64)              { m.AllocQty = v }
func (m *Message) SetQtyType(v int)                   { m.QtyType = &v }
func (m *Message) SetSide(v string)                   { m.Side = v }
func (m *Message) SetCurrency(v string)               { m.Currency = &v }
func (m *Message) SetLastMkt(v string)                { m.LastMkt = &v }
func (m *Message) SetNoCapacities(v []NoCapacities)   { m.NoCapacities = v }
func (m *Message) SetAllocAccount(v string)           { m.AllocAccount = v }
func (m *Message) SetAllocAcctIDSource(v int)         { m.AllocAcctIDSource = &v }
func (m *Message) SetAllocAccountType(v int)          { m.AllocAccountType = &v }
func (m *Message) SetAvgPx(v float64)                 { m.AvgPx = v }
func (m *Message) SetAvgPxPrecision(v int)            { m.AvgPxPrecision = &v }
func (m *Message) SetPriceType(v int)                 { m.PriceType = &v }
func (m *Message) SetAvgParPx(v float64)              { m.AvgParPx = &v }
func (m *Message) SetReportedPx(v float64)            { m.ReportedPx = &v }
func (m *Message) SetText(v string)                   { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)            { m.EncodedText = &v }
func (m *Message) SetProcessCode(v string)            { m.ProcessCode = &v }
func (m *Message) SetGrossTradeAmt(v float64)         { m.GrossTradeAmt = v }
func (m *Message) SetNumDaysInterest(v int)           { m.NumDaysInterest = &v }
func (m *Message) SetExDate(v string)                 { m.ExDate = &v }
func (m *Message) SetAccruedInterestRate(v float64)   { m.AccruedInterestRate = &v }
func (m *Message) SetAccruedInterestAmt(v float64)    { m.AccruedInterestAmt = &v }
func (m *Message) SetInterestAtMaturity(v float64)    { m.InterestAtMaturity = &v }
func (m *Message) SetEndAccruedInterestAmt(v float64) { m.EndAccruedInterestAmt = &v }
func (m *Message) SetStartCash(v float64)             { m.StartCash = &v }
func (m *Message) SetEndCash(v float64)               { m.EndCash = &v }
func (m *Message) SetConcession(v float64)            { m.Concession = &v }
func (m *Message) SetTotalTakedown(v float64)         { m.TotalTakedown = &v }
func (m *Message) SetNetMoney(v float64)              { m.NetMoney = v }
func (m *Message) SetMaturityNetMoney(v float64)      { m.MaturityNetMoney = &v }
func (m *Message) SetSettlCurrAmt(v float64)          { m.SettlCurrAmt = &v }
func (m *Message) SetSettlCurrency(v string)          { m.SettlCurrency = &v }
func (m *Message) SetSettlCurrFxRate(v float64)       { m.SettlCurrFxRate = &v }
func (m *Message) SetSettlCurrFxRateCalc(v string)    { m.SettlCurrFxRateCalc = &v }
func (m *Message) SetSettlType(v string)              { m.SettlType = &v }
func (m *Message) SetSettlDate(v string)              { m.SettlDate = &v }
func (m *Message) SetSharedCommission(v float64)      { m.SharedCommission = &v }
func (m *Message) SetNoMiscFees(v []NoMiscFees)       { m.NoMiscFees = v }

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
	return enum.BeginStringFIX44, "AK", r
}
