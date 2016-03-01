//Package ordermassactionreport msg type = BZ.
package ordermassactionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/affectedordgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/notaffectedordersgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a OrderMassActionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BZ"`
	Header     fixt11.Header
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
	//AffectedOrdGrp Component
	AffectedOrdGrp affectedordgrp.Component
	//MarketID is a non-required field for OrderMassActionReport.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for OrderMassActionReport.
	MarketSegmentID *string `fix:"1300"`
	//TradingSessionID is a non-required field for OrderMassActionReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassActionReport.
	TradingSessionSubID *string `fix:"625"`
	//Parties Component
	Parties parties.Component
	//Instrument Component
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
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
	//NotAffectedOrdersGrp Component
	NotAffectedOrdersGrp notaffectedordersgrp.Component
	Trailer              fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "BZ", r
}
