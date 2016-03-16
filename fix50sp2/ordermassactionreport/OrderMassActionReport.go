//Package ordermassactionreport msg type = BZ.
package ordermassactionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/affectedordgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/notaffectedordersgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/targetparties"
	"github.com/quickfixgo/quickfix/fix50sp2/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a OrderMassActionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BZ"`
	fixt11.Header
	//ClOrdID is a non-required field for OrderMassActionReport.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassActionReport.
	SecondaryClOrdID *string `fix:"526"`
	//MassActionReportID is a required field for OrderMassActionReport.
	MassActionReportID string `fix:"1369"`
	//MassActionType is a required field for OrderMassActionReport.
	MassActionType int `fix:"1373"`
	//MassActionScope is a required field for OrderMassActionReport.
	MassActionScope int `fix:"1374"`
	//MassActionResponse is a required field for OrderMassActionReport.
	MassActionResponse int `fix:"1375"`
	//MassActionRejectReason is a non-required field for OrderMassActionReport.
	MassActionRejectReason *int `fix:"1376"`
	//TotalAffectedOrders is a non-required field for OrderMassActionReport.
	TotalAffectedOrders *int `fix:"533"`
	//AffectedOrdGrp is a non-required component for OrderMassActionReport.
	AffectedOrdGrp *affectedordgrp.AffectedOrdGrp
	//MarketID is a non-required field for OrderMassActionReport.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for OrderMassActionReport.
	MarketSegmentID *string `fix:"1300"`
	//TradingSessionID is a non-required field for OrderMassActionReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassActionReport.
	TradingSessionSubID *string `fix:"625"`
	//Parties is a non-required component for OrderMassActionReport.
	Parties *parties.Parties
	//Instrument is a non-required component for OrderMassActionReport.
	Instrument *instrument.Instrument
	//UnderlyingInstrument is a non-required component for OrderMassActionReport.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
	//Side is a non-required field for OrderMassActionReport.
	Side *string `fix:"54"`
	//TransactTime is a non-required field for OrderMassActionReport.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for OrderMassActionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassActionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassActionReport.
	EncodedText *string `fix:"355"`
	//NotAffectedOrdersGrp is a non-required component for OrderMassActionReport.
	NotAffectedOrdersGrp *notaffectedordersgrp.NotAffectedOrdersGrp
	//TargetParties is a non-required component for OrderMassActionReport.
	TargetParties *targetparties.TargetParties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)                               { m.ClOrdID = &v }
func (m *Message) SetSecondaryClOrdID(v string)                      { m.SecondaryClOrdID = &v }
func (m *Message) SetMassActionReportID(v string)                    { m.MassActionReportID = v }
func (m *Message) SetMassActionType(v int)                           { m.MassActionType = v }
func (m *Message) SetMassActionScope(v int)                          { m.MassActionScope = v }
func (m *Message) SetMassActionResponse(v int)                       { m.MassActionResponse = v }
func (m *Message) SetMassActionRejectReason(v int)                   { m.MassActionRejectReason = &v }
func (m *Message) SetTotalAffectedOrders(v int)                      { m.TotalAffectedOrders = &v }
func (m *Message) SetAffectedOrdGrp(v affectedordgrp.AffectedOrdGrp) { m.AffectedOrdGrp = &v }
func (m *Message) SetMarketID(v string)                              { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string)                       { m.MarketSegmentID = &v }
func (m *Message) SetTradingSessionID(v string)                      { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)                   { m.TradingSessionSubID = &v }
func (m *Message) SetParties(v parties.Parties)                      { m.Parties = &v }
func (m *Message) SetInstrument(v instrument.Instrument)             { m.Instrument = &v }
func (m *Message) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}
func (m *Message) SetSide(v string)            { m.Side = &v }
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = &v }
func (m *Message) SetText(v string)            { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)     { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)     { m.EncodedText = &v }
func (m *Message) SetNotAffectedOrdersGrp(v notaffectedordersgrp.NotAffectedOrdersGrp) {
	m.NotAffectedOrdersGrp = &v
}
func (m *Message) SetTargetParties(v targetparties.TargetParties) { m.TargetParties = &v }

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
	return enum.ApplVerID_FIX50SP2, "BZ", r
}
