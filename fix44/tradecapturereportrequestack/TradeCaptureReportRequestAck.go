//Package tradecapturereportrequestack msg type = AQ.
package tradecapturereportrequestack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoUnderlyings is a repeating group in TradeCaptureReportRequestAck
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in TradeCaptureReportRequestAck
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//Message is a TradeCaptureReportRequestAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AQ"`
	Header     fix44.Header
	//TradeRequestID is a required field for TradeCaptureReportRequestAck.
	TradeRequestID string `fix:"568"`
	//TradeRequestType is a required field for TradeCaptureReportRequestAck.
	TradeRequestType int `fix:"569"`
	//SubscriptionRequestType is a non-required field for TradeCaptureReportRequestAck.
	SubscriptionRequestType *string `fix:"263"`
	//TotNumTradeReports is a non-required field for TradeCaptureReportRequestAck.
	TotNumTradeReports *int `fix:"748"`
	//TradeRequestResult is a required field for TradeCaptureReportRequestAck.
	TradeRequestResult int `fix:"749"`
	//TradeRequestStatus is a required field for TradeCaptureReportRequestAck.
	TradeRequestStatus int `fix:"750"`
	//Instrument Component
	Instrument instrument.Component
	//NoUnderlyings is a non-required field for TradeCaptureReportRequestAck.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for TradeCaptureReportRequestAck.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//MultiLegReportingType is a non-required field for TradeCaptureReportRequestAck.
	MultiLegReportingType *string `fix:"442"`
	//ResponseTransportType is a non-required field for TradeCaptureReportRequestAck.
	ResponseTransportType *int `fix:"725"`
	//ResponseDestination is a non-required field for TradeCaptureReportRequestAck.
	ResponseDestination *string `fix:"726"`
	//Text is a non-required field for TradeCaptureReportRequestAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradeCaptureReportRequestAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradeCaptureReportRequestAck.
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
	return enum.BeginStringFIX44, "AQ", r
}
