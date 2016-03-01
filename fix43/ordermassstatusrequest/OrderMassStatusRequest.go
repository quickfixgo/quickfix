//Package ordermassstatusrequest msg type = AF.
package ordermassstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"github.com/quickfixgo/quickfix/fix43/underlyinginstrument"
)

//Message is a OrderMassStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AF"`
	Header     fix43.Header
	//MassStatusReqID is a required field for OrderMassStatusRequest.
	MassStatusReqID string `fix:"584"`
	//MassStatusReqType is a required field for OrderMassStatusRequest.
	MassStatusReqType int `fix:"585"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for OrderMassStatusRequest.
	Account *string `fix:"1"`
	//TradingSessionID is a non-required field for OrderMassStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//Side is a non-required field for OrderMassStatusRequest.
	Side    *string `fix:"54"`
	Trailer fix43.Trailer
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
	return enum.BeginStringFIX43, "AF", r
}
