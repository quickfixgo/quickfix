//Package ordermasscancelrequest msg type = q.
package ordermasscancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/underlyinginstrument"
	"time"
)

//Message is a OrderMassCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"q"`
	fix43.Header
	//ClOrdID is a required field for OrderMassCancelRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassCancelRequest.
	SecondaryClOrdID *string `fix:"526"`
	//MassCancelRequestType is a required field for OrderMassCancelRequest.
	MassCancelRequestType string `fix:"530"`
	//TradingSessionID is a non-required field for OrderMassCancelRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassCancelRequest.
	TradingSessionSubID *string `fix:"625"`
	//Instrument Component
	instrument.Instrument
	//UnderlyingInstrument Component
	underlyinginstrument.UnderlyingInstrument
	//Side is a non-required field for OrderMassCancelRequest.
	Side *string `fix:"54"`
	//TransactTime is a required field for OrderMassCancelRequest.
	TransactTime time.Time `fix:"60"`
	//Text is a non-required field for OrderMassCancelRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassCancelRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassCancelRequest.
	EncodedText *string `fix:"355"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)               { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)      { m.SecondaryClOrdID = &v }
func (m *Message) SetMassCancelRequestType(v string) { m.MassCancelRequestType = v }
func (m *Message) SetTradingSessionID(v string)      { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)   { m.TradingSessionSubID = &v }
func (m *Message) SetSide(v string)                  { m.Side = &v }
func (m *Message) SetTransactTime(v time.Time)       { m.TransactTime = v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }

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
	return enum.BeginStringFIX43, "q", r
}
