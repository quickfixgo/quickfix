//Package ioi msg type = 6.
package ioi

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/financingdetails"
	"github.com/quickfixgo/quickfix/fix50/instrmtlegioigrp"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/ioiqualgrp"
	"github.com/quickfixgo/quickfix/fix50/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/routinggrp"
	"github.com/quickfixgo/quickfix/fix50/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50/stipulations"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a IOI FIX Message
type Message struct {
	FIXMsgType string `fix:"6"`
	Header     fixt11.Header
	//IOIID is a required field for IOI.
	IOIID string `fix:"23"`
	//IOITransType is a required field for IOI.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IOI.
	IOIRefID *string `fix:"26"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Side is a required field for IOI.
	Side string `fix:"54"`
	//QtyType is a non-required field for IOI.
	QtyType *int `fix:"854"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//IOIQty is a required field for IOI.
	IOIQty string `fix:"27"`
	//Currency is a non-required field for IOI.
	Currency *string `fix:"15"`
	//Stipulations Component
	Stipulations stipulations.Component
	//InstrmtLegIOIGrp Component
	InstrmtLegIOIGrp instrmtlegioigrp.Component
	//PriceType is a non-required field for IOI.
	PriceType *int `fix:"423"`
	//Price is a non-required field for IOI.
	Price *float64 `fix:"44"`
	//ValidUntilTime is a non-required field for IOI.
	ValidUntilTime *time.Time `fix:"62"`
	//IOIQltyInd is a non-required field for IOI.
	IOIQltyInd *string `fix:"25"`
	//IOINaturalFlag is a non-required field for IOI.
	IOINaturalFlag *bool `fix:"130"`
	//IOIQualGrp Component
	IOIQualGrp ioiqualgrp.Component
	//Text is a non-required field for IOI.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for IOI.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for IOI.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for IOI.
	TransactTime *time.Time `fix:"60"`
	//URLLink is a non-required field for IOI.
	URLLink *string `fix:"149"`
	//RoutingGrp Component
	RoutingGrp routinggrp.Component
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//Parties Component
	Parties parties.Component
	Trailer fixt11.Trailer
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
	return enum.BeginStringFIX50, "6", r
}
