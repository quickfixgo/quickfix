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
	fixt11.Header
	//IOIID is a required field for IOI.
	IOIID string `fix:"23"`
	//IOITransType is a required field for IOI.
	IOITransType string `fix:"28"`
	//IOIRefID is a non-required field for IOI.
	IOIRefID *string `fix:"26"`
	//Instrument is a required component for IOI.
	instrument.Instrument
	//FinancingDetails is a non-required component for IOI.
	FinancingDetails *financingdetails.FinancingDetails
	//UndInstrmtGrp is a non-required component for IOI.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Side is a required field for IOI.
	Side string `fix:"54"`
	//QtyType is a non-required field for IOI.
	QtyType *int `fix:"854"`
	//OrderQtyData is a non-required component for IOI.
	OrderQtyData *orderqtydata.OrderQtyData
	//IOIQty is a required field for IOI.
	IOIQty string `fix:"27"`
	//Currency is a non-required field for IOI.
	Currency *string `fix:"15"`
	//Stipulations is a non-required component for IOI.
	Stipulations *stipulations.Stipulations
	//InstrmtLegIOIGrp is a non-required component for IOI.
	InstrmtLegIOIGrp *instrmtlegioigrp.InstrmtLegIOIGrp
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
	//IOIQualGrp is a non-required component for IOI.
	IOIQualGrp *ioiqualgrp.IOIQualGrp
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
	//RoutingGrp is a non-required component for IOI.
	RoutingGrp *routinggrp.RoutingGrp
	//SpreadOrBenchmarkCurveData is a non-required component for IOI.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for IOI.
	YieldData *yielddata.YieldData
	//Parties is a non-required component for IOI.
	Parties *parties.Parties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized IOI instance
func New(ioiid string, ioitranstype string, instrument instrument.Instrument, side string, ioiqty string) *Message {
	var m Message
	m.SetIOIID(ioiid)
	m.SetIOITransType(ioitranstype)
	m.SetInstrument(instrument)
	m.SetSide(side)
	m.SetIOIQty(ioiqty)
	return &m
}

func (m *Message) SetIOIID(v string)                                       { m.IOIID = v }
func (m *Message) SetIOITransType(v string)                                { m.IOITransType = v }
func (m *Message) SetIOIRefID(v string)                                    { m.IOIRefID = &v }
func (m *Message) SetInstrument(v instrument.Instrument)                   { m.Instrument = v }
func (m *Message) SetFinancingDetails(v financingdetails.FinancingDetails) { m.FinancingDetails = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)          { m.UndInstrmtGrp = &v }
func (m *Message) SetSide(v string)                                        { m.Side = v }
func (m *Message) SetQtyType(v int)                                        { m.QtyType = &v }
func (m *Message) SetOrderQtyData(v orderqtydata.OrderQtyData)             { m.OrderQtyData = &v }
func (m *Message) SetIOIQty(v string)                                      { m.IOIQty = v }
func (m *Message) SetCurrency(v string)                                    { m.Currency = &v }
func (m *Message) SetStipulations(v stipulations.Stipulations)             { m.Stipulations = &v }
func (m *Message) SetInstrmtLegIOIGrp(v instrmtlegioigrp.InstrmtLegIOIGrp) { m.InstrmtLegIOIGrp = &v }
func (m *Message) SetPriceType(v int)                                      { m.PriceType = &v }
func (m *Message) SetPrice(v float64)                                      { m.Price = &v }
func (m *Message) SetValidUntilTime(v time.Time)                           { m.ValidUntilTime = &v }
func (m *Message) SetIOIQltyInd(v string)                                  { m.IOIQltyInd = &v }
func (m *Message) SetIOINaturalFlag(v bool)                                { m.IOINaturalFlag = &v }
func (m *Message) SetIOIQualGrp(v ioiqualgrp.IOIQualGrp)                   { m.IOIQualGrp = &v }
func (m *Message) SetText(v string)                                        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                                 { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                                 { m.EncodedText = &v }
func (m *Message) SetTransactTime(v time.Time)                             { m.TransactTime = &v }
func (m *Message) SetURLLink(v string)                                     { m.URLLink = &v }
func (m *Message) SetRoutingGrp(v routinggrp.RoutingGrp)                   { m.RoutingGrp = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData) { m.YieldData = &v }
func (m *Message) SetParties(v parties.Parties)       { m.Parties = &v }

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
	return enum.ApplVerID_FIX50, "6", r
}
