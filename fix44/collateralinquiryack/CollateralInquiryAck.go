//Package collateralinquiryack msg type = BG.
package collateralinquiryack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/financingdetails"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/parties"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoCollInquiryQualifier is a repeating group in CollateralInquiryAck
type NoCollInquiryQualifier struct {
	//CollInquiryQualifier is a non-required field for NoCollInquiryQualifier.
	CollInquiryQualifier *int `fix:"896"`
}

//NoExecs is a repeating group in CollateralInquiryAck
type NoExecs struct {
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
}

//NoTrades is a repeating group in CollateralInquiryAck
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//NoLegs is a repeating group in CollateralInquiryAck
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoUnderlyings is a repeating group in CollateralInquiryAck
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//Message is a CollateralInquiryAck FIX Message
type Message struct {
	FIXMsgType string `fix:"BG"`
	Header     fix44.Header
	//CollInquiryID is a required field for CollateralInquiryAck.
	CollInquiryID string `fix:"909"`
	//CollInquiryStatus is a required field for CollateralInquiryAck.
	CollInquiryStatus int `fix:"945"`
	//CollInquiryResult is a non-required field for CollateralInquiryAck.
	CollInquiryResult *int `fix:"946"`
	//NoCollInquiryQualifier is a non-required field for CollateralInquiryAck.
	NoCollInquiryQualifier []NoCollInquiryQualifier `fix:"938,omitempty"`
	//TotNumReports is a non-required field for CollateralInquiryAck.
	TotNumReports *int `fix:"911"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for CollateralInquiryAck.
	Account *string `fix:"1"`
	//AccountType is a non-required field for CollateralInquiryAck.
	AccountType *int `fix:"581"`
	//ClOrdID is a non-required field for CollateralInquiryAck.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for CollateralInquiryAck.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for CollateralInquiryAck.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for CollateralInquiryAck.
	SecondaryClOrdID *string `fix:"526"`
	//NoExecs is a non-required field for CollateralInquiryAck.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//NoTrades is a non-required field for CollateralInquiryAck.
	NoTrades []NoTrades `fix:"897,omitempty"`
	//Instrument Component
	Instrument instrument.Component
	//FinancingDetails Component
	FinancingDetails financingdetails.Component
	//SettlDate is a non-required field for CollateralInquiryAck.
	SettlDate *string `fix:"64"`
	//Quantity is a non-required field for CollateralInquiryAck.
	Quantity *float64 `fix:"53"`
	//QtyType is a non-required field for CollateralInquiryAck.
	QtyType *int `fix:"854"`
	//Currency is a non-required field for CollateralInquiryAck.
	Currency *string `fix:"15"`
	//NoLegs is a non-required field for CollateralInquiryAck.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for CollateralInquiryAck.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//TradingSessionID is a non-required field for CollateralInquiryAck.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for CollateralInquiryAck.
	TradingSessionSubID *string `fix:"625"`
	//SettlSessID is a non-required field for CollateralInquiryAck.
	SettlSessID *string `fix:"716"`
	//SettlSessSubID is a non-required field for CollateralInquiryAck.
	SettlSessSubID *string `fix:"717"`
	//ClearingBusinessDate is a non-required field for CollateralInquiryAck.
	ClearingBusinessDate *string `fix:"715"`
	//ResponseTransportType is a non-required field for CollateralInquiryAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for CollateralInquiryAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for CollateralInquiryAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for CollateralInquiryAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for CollateralInquiryAck.
	EncodedText *string `fix:"355"`
	Trailer     fix44.Trailer
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
	return enum.BeginStringFIX44, "BG", r
}
