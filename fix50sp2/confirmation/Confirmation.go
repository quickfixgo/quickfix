//Package confirmation msg type = AK.
package confirmation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp2/cpctyconfgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/financingdetails"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/ordallocgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/settlinstructionsdata"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/trdregtimestamps"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a Confirmation FIX Message
type Message struct {
	FIXMsgType string `fix:"AK"`
	Header     fixt11.Header
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
	Parties parties.Component
	//OrdAllocGrp Component
	OrdAllocGrp ordallocgrp.Component
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
	TrdRegTimestamps trdregtimestamps.Component
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//YieldData Component
	YieldData yielddata.Component
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
	//CpctyConfGrp Component
	CpctyConfGrp cpctyconfgrp.Component
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
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
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
	SettlInstructionsData settlinstructionsdata.Component
	//CommissionData Component
	CommissionData commissiondata.Component
	//SharedCommission is a non-required field for Confirmation.
	SharedCommission *float64 `fix:"858"`
	//Stipulations Component
	Stipulations stipulations.Component
	//MiscFeesGrp Component
	MiscFeesGrp miscfeesgrp.Component
	Trailer     fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.ApplVerID_FIX50SP2, "AK", r
}
