//Package securitydefinition msg type = d.
package securitydefinition

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp2/marketsegmentgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp2/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SecurityDefinition FIX Message
type Message struct {
	FIXMsgType string `fix:"d"`
	fixt11.Header
	//SecurityReqID is a non-required field for SecurityDefinition.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityDefinition.
	SecurityResponseID *string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinition.
	SecurityResponseType *int `fix:"323"`
	//Instrument is a non-required component for SecurityDefinition.
	Instrument *instrument.Instrument
	//InstrumentExtension is a non-required component for SecurityDefinition.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//UndInstrmtGrp is a non-required component for SecurityDefinition.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for SecurityDefinition.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinition.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinition.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinition.
	EncodedText *string `fix:"355"`
	//InstrmtLegGrp is a non-required component for SecurityDefinition.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//SecurityReportID is a non-required field for SecurityDefinition.
	SecurityReportID *int `fix:"964"`
	//ClearingBusinessDate is a non-required field for SecurityDefinition.
	ClearingBusinessDate *string `fix:"715"`
	//Stipulations is a non-required component for SecurityDefinition.
	Stipulations *stipulations.Stipulations
	//SpreadOrBenchmarkCurveData is a non-required component for SecurityDefinition.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for SecurityDefinition.
	YieldData *yielddata.YieldData
	//CorporateAction is a non-required field for SecurityDefinition.
	CorporateAction *string `fix:"292"`
	//MarketSegmentGrp is a non-required component for SecurityDefinition.
	MarketSegmentGrp *marketsegmentgrp.MarketSegmentGrp
	//ApplicationSequenceControl is a non-required component for SecurityDefinition.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	//TransactTime is a non-required field for SecurityDefinition.
	TransactTime *time.Time `fix:"60"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SecurityDefinition instance
func New() *Message {
	var m Message
	return &m
}

func (m *Message) SetSecurityReqID(v string)             { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)        { m.SecurityResponseID = &v }
func (m *Message) SetSecurityResponseType(v int)         { m.SecurityResponseType = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *Message) SetCurrency(v string)                           { m.Currency = &v }
func (m *Message) SetText(v string)                               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *Message) SetSecurityReportID(v int)                      { m.SecurityReportID = &v }
func (m *Message) SetClearingBusinessDate(v string)               { m.ClearingBusinessDate = &v }
func (m *Message) SetStipulations(v stipulations.Stipulations)    { m.Stipulations = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData)                      { m.YieldData = &v }
func (m *Message) SetCorporateAction(v string)                             { m.CorporateAction = &v }
func (m *Message) SetMarketSegmentGrp(v marketsegmentgrp.MarketSegmentGrp) { m.MarketSegmentGrp = &v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = &v }

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
	return enum.ApplVerID_FIX50SP2, "d", r
}
