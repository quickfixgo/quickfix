//Package positionmaintenancereport msg type = AM.
package positionmaintenancereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50/positionqty"
	"github.com/quickfixgo/quickfix/fix50/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a PositionMaintenanceReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AM"`
	Header     fixt11.Header
	//PosMaintRptID is a required field for PositionMaintenanceReport.
	PosMaintRptID string `fix:"721"`
	//PosTransType is a required field for PositionMaintenanceReport.
	PosTransType int `fix:"709"`
	//PosReqID is a non-required field for PositionMaintenanceReport.
	PosReqID *string `fix:"710"`
	//PosMaintAction is a required field for PositionMaintenanceReport.
	PosMaintAction int `fix:"712"`
	//OrigPosReqRefID is a non-required field for PositionMaintenanceReport.
	OrigPosReqRefID *string `fix:"713"`
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
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for PositionMaintenanceReport.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for PositionMaintenanceReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for PositionMaintenanceReport.
	AccountType *int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for PositionMaintenanceReport.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//TransactTime is a non-required field for PositionMaintenanceReport.
	TransactTime *time.Time `fix:"60"`
	//PositionQty Component
	PositionQty positionqty.Component
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
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
	//SettlCurrency is a non-required field for PositionMaintenanceReport.
	SettlCurrency *string `fix:"120"`
	//ContraryInstructionIndicator is a non-required field for PositionMaintenanceReport.
	ContraryInstructionIndicator *bool `fix:"719"`
	//PriorSpreadIndicator is a non-required field for PositionMaintenanceReport.
	PriorSpreadIndicator *bool `fix:"720"`
	//PosMaintRptRefID is a non-required field for PositionMaintenanceReport.
	PosMaintRptRefID *string `fix:"714"`
	Trailer          fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX50, "AM", r
}
