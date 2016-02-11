//Package ordermasscancelreport msg type = r.
package ordermasscancelreport

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

//Message is a OrderMassCancelReport FIX Message
type Message struct {
	FIXMsgType string `fix:"r"`
	Header     fixt11.Header
	//ClOrdID is a non-required field for OrderMassCancelReport.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassCancelReport.
	SecondaryClOrdID *string `fix:"526"`
	//OrderID is a required field for OrderMassCancelReport.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for OrderMassCancelReport.
	SecondaryOrderID *string `fix:"198"`
	//MassCancelRequestType is a required field for OrderMassCancelReport.
	MassCancelRequestType string `fix:"530"`
	//MassCancelResponse is a required field for OrderMassCancelReport.
	MassCancelResponse string `fix:"531"`
	//MassCancelRejectReason is a non-required field for OrderMassCancelReport.
	MassCancelRejectReason *int `fix:"532"`
	//TotalAffectedOrders is a non-required field for OrderMassCancelReport.
	TotalAffectedOrders *int `fix:"533"`
	//AffectedOrdGrp Component
	AffectedOrdGrp affectedordgrp.Component
	//TradingSessionID is a non-required field for OrderMassCancelReport.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassCancelReport.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//Side is a non-required field for OrderMassCancelReport.
	Side *string `fix:"54"`
	//TransactTime is a non-required field for OrderMassCancelReport.
	TransactTime *time.Time `fix:"60"`
	//Text is a non-required field for OrderMassCancelReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassCancelReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassCancelReport.
	EncodedText *string `fix:"355"`
	//Parties Component
	Parties parties.Component
	//MassActionReportID is a required field for OrderMassCancelReport.
	MassActionReportID string `fix:"1369"`
	//NotAffectedOrdersGrp Component
	NotAffectedOrdersGrp notaffectedordersgrp.Component
	//MarketID is a non-required field for OrderMassCancelReport.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for OrderMassCancelReport.
	MarketSegmentID *string `fix:"1300"`
	//TargetParties Component
	TargetParties targetparties.Component
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
	return enum.ApplVerID_FIX50SP2, "r", r
}
