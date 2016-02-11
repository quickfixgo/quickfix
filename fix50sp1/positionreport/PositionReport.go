//Package positionreport msg type = AP.
package positionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/positionamountdata"
	"github.com/quickfixgo/quickfix/fix50sp1/positionqty"
	"github.com/quickfixgo/quickfix/fix50sp1/posundinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a PositionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"AP"`
	Header     fixt11.Header
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
	//PosReqResult is a non-required field for PositionReport.
	PosReqResult *int `fix:"728"`
	//ClearingBusinessDate is a required field for PositionReport.
	ClearingBusinessDate string `fix:"715"`
	//SettlSessID is a non-required field for PositionReport.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for PositionReport.
	SettlSessSubID *string `fix:"717"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for PositionReport.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for PositionReport.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for PositionReport.
	AccountType *int `fix:"581"`
	//Instrument Component
	Instrument instrument.Component
	//Currency is a non-required field for PositionReport.
	Currency *string `fix:"15"`
	//SettlPrice is a non-required field for PositionReport.
	SettlPrice *float64 `fix:"730"`
	//SettlPriceType is a non-required field for PositionReport.
	SettlPriceType *int `fix:"731"`
	//PriorSettlPrice is a non-required field for PositionReport.
	PriorSettlPrice *float64 `fix:"734"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//PosUndInstrmtGrp Component
	PosUndInstrmtGrp posundinstrmtgrp.Component
	//PositionQty Component
	PositionQty positionqty.Component
	//PositionAmountData Component
	PositionAmountData positionamountdata.Component
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
	//MatchStatus is a non-required field for PositionReport.
	MatchStatus *string `fix:"573"`
	//PriceType is a non-required field for PositionReport.
	PriceType *int `fix:"423"`
	//SettlCurrency is a non-required field for PositionReport.
	SettlCurrency *string `fix:"120"`
	//MessageEventSource is a non-required field for PositionReport.
	MessageEventSource *string `fix:"1011"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "AP", r
}
