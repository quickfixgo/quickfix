//Package ordermassstatusrequest msg type = AF.
package ordermassstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrument"
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a OrderMassStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AF"`
	Header     fixt11.Header
	//MassStatusReqID is a required field for OrderMassStatusRequest.
	MassStatusReqID string `fix:"584"`
	//MassStatusReqType is a required field for OrderMassStatusRequest.
	MassStatusReqType int `fix:"585"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for OrderMassStatusRequest.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for OrderMassStatusRequest.
	AcctIDSource *int `fix:"660"`
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
	Trailer fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMassStatusReqID(v string)     { m.MassStatusReqID = v }
func (m *Message) SetMassStatusReqType(v int)      { m.MassStatusReqType = v }
func (m *Message) SetAccount(v string)             { m.Account = &v }
func (m *Message) SetAcctIDSource(v int)           { m.AcctIDSource = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *Message) SetSide(v string)                { m.Side = &v }

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
	return enum.BeginStringFIX50, "AF", r
}
