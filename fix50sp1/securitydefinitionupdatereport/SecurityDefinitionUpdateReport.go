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
	Header     fixt11.Header
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
	//Instrument Component
	Instrument instrument.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//Currency is a non-required field for SecurityDefinitionUpdateReport.
	Currency *string `fix:"15"`
	//Text is a non-required field for SecurityDefinitionUpdateReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinitionUpdateReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinitionUpdateReport.
	EncodedText *string `fix:"355"`
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//Stipulations Component
	Stipulations stipulations.Component
	//SpreadOrBenchmarkCurveData Component
	SpreadOrBenchmarkCurveData spreadorbenchmarkcurvedata.Component
	//YieldData Component
	YieldData yielddata.Component
	//MarketSegmentGrp Component
	MarketSegmentGrp marketsegmentgrp.Component
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "BP", r
}
