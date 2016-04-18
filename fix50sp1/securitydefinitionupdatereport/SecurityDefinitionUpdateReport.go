//Package securitydefinitionupdatereport msg type = BP.
package securitydefinitionupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/marketsegmentgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/spreadorbenchmarkcurvedata"
	"github.com/quickfixgo/quickfix/fix50sp1/stipulations"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/yielddata"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityDefinitionUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BP"`
	fixt11.Header
	//SecurityReportID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityReportID *int `fix:"964"`
	//SecurityReqID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityDefinitionUpdateReport.
	SecurityResponseID *string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinitionUpdateReport.
	SecurityResponseType *int `fix:"323"`
	//ClearingBusinessDate is a non-required field for SecurityDefinitionUpdateReport.
	ClearingBusinessDate *string `fix:"715"`
	//SecurityUpdateAction is a non-required field for SecurityDefinitionUpdateReport.
	SecurityUpdateAction *string `fix:"980"`
	//CorporateAction is a non-required field for SecurityDefinitionUpdateReport.
	CorporateAction *string `fix:"292"`
	//Instrument is a non-required component for SecurityDefinitionUpdateReport.
	Instrument *instrument.Instrument
	//UndInstrmtGrp is a non-required component for SecurityDefinitionUpdateReport.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//Currency is a non-required field for SecurityDefinitionUpdateReport.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinitionUpdateReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionUpdateReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionUpdateReport.
	EncodedText *string `fix:"355"`
	//InstrmtLegGrp is a non-required component for SecurityDefinitionUpdateReport.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//InstrumentExtension is a non-required component for SecurityDefinitionUpdateReport.
	InstrumentExtension *instrumentextension.InstrumentExtension
	//Stipulations is a non-required component for SecurityDefinitionUpdateReport.
	Stipulations *stipulations.Stipulations
	//SpreadOrBenchmarkCurveData is a non-required component for SecurityDefinitionUpdateReport.
	SpreadOrBenchmarkCurveData *spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData
	//YieldData is a non-required component for SecurityDefinitionUpdateReport.
	YieldData *yielddata.YieldData
	//MarketSegmentGrp is a non-required component for SecurityDefinitionUpdateReport.
	MarketSegmentGrp *marketsegmentgrp.MarketSegmentGrp
	//ApplicationSequenceControl is a non-required component for SecurityDefinitionUpdateReport.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SecurityDefinitionUpdateReport instance
func New() *Message {
	var m Message
	return &m
}

func (m *Message) SetSecurityReportID(v int)                      { m.SecurityReportID = &v }
func (m *Message) SetSecurityReqID(v string)                      { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)                 { m.SecurityResponseID = &v }
func (m *Message) SetSecurityResponseType(v int)                  { m.SecurityResponseType = &v }
func (m *Message) SetClearingBusinessDate(v string)               { m.ClearingBusinessDate = &v }
func (m *Message) SetSecurityUpdateAction(v string)               { m.SecurityUpdateAction = &v }
func (m *Message) SetCorporateAction(v string)                    { m.CorporateAction = &v }
func (m *Message) SetInstrument(v instrument.Instrument)          { m.Instrument = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
func (m *Message) SetCurrency(v string)                           { m.Currency = &v }
func (m *Message) SetText(v string)                               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp) { m.InstrmtLegGrp = &v }
func (m *Message) SetInstrumentExtension(v instrumentextension.InstrumentExtension) {
	m.InstrumentExtension = &v
}
func (m *Message) SetStipulations(v stipulations.Stipulations) { m.Stipulations = &v }
func (m *Message) SetSpreadOrBenchmarkCurveData(v spreadorbenchmarkcurvedata.SpreadOrBenchmarkCurveData) {
	m.SpreadOrBenchmarkCurveData = &v
}
func (m *Message) SetYieldData(v yielddata.YieldData)                      { m.YieldData = &v }
func (m *Message) SetMarketSegmentGrp(v marketsegmentgrp.MarketSegmentGrp) { m.MarketSegmentGrp = &v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}

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
	return enum.ApplVerID_FIX50SP1, "BP", r
}
