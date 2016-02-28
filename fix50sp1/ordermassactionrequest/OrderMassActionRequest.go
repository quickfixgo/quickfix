//Package ordermassactionrequest msg type = CA.
package ordermassactionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a OrderMassActionRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"CA"`
	Header     fixt11.Header
	//ClOrdID is a required field for OrderMassActionRequest.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for OrderMassActionRequest.
	SecondaryClOrdID *string `fix:"526"`
	//MassActionType is a required field for OrderMassActionRequest.
	MassActionType int `fix:"1373"`
	//MassActionScope is a required field for OrderMassActionRequest.
	MassActionScope int `fix:"1374"`
	//MarketID is a non-required field for OrderMassActionRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for OrderMassActionRequest.
	MarketSegmentID *string `fix:"1300"`
	//TradingSessionID is a non-required field for OrderMassActionRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for OrderMassActionRequest.
	TradingSessionSubID *string `fix:"625"`
	//Parties Component
	Parties parties.Component
	//Instrument Component
	Instrument instrument.Component
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//Side is a non-required field for OrderMassActionRequest.
	Side *string `fix:"54"`
	//TransactTime is a required field for OrderMassActionRequest.
	TransactTime time.Time `fix:"60"`
	//Text is a non-required field for OrderMassActionRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderMassActionRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderMassActionRequest.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClOrdID(v string)             { m.ClOrdID = v }
func (m *Message) SetSecondaryClOrdID(v string)    { m.SecondaryClOrdID = &v }
func (m *Message) SetMassActionType(v int)         { m.MassActionType = v }
func (m *Message) SetMassActionScope(v int)        { m.MassActionScope = v }
func (m *Message) SetMarketID(v string)            { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string)     { m.MarketSegmentID = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *Message) SetSide(v string)                { m.Side = &v }
func (m *Message) SetTransactTime(v time.Time)     { m.TransactTime = v }
func (m *Message) SetText(v string)                { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)         { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP1, "CA", r
}
