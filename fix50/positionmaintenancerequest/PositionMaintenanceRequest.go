//Package positionmaintenancerequest msg type = AL.
package positionmaintenancerequest

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

//Message is a PositionMaintenanceRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AL"`
	Header     fixt11.Header
	//PosReqID is a non-required field for PositionMaintenanceRequest.
	PosReqID *string `fix:"710"`
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
	//Account is a non-required field for PositionMaintenanceRequest.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for PositionMaintenanceRequest.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for PositionMaintenanceRequest.
	AccountType *int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for PositionMaintenanceRequest.
	Currency *string `fix:"15"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//TransactTime is a non-required field for PositionMaintenanceRequest.
	TransactTime *time.Time `fix:"60"`
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
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
	//SettlCurrency is a non-required field for PositionMaintenanceRequest.
	SettlCurrency *string `fix:"120"`
	Trailer       fixt11.Trailer
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
	return enum.BeginStringFIX50, "AL", r
}
