//Package tradingsessionlist msg type = BJ.
package tradingsessionlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/trdsesslstgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradingSessionList FIX Message
type Message struct {
	FIXMsgType string `fix:"BJ"`
	fixt11.Header
	//TradSesReqID is a non-required field for TradingSessionList.
	TradSesReqID *string `fix:"335"`
	//TrdSessLstGrp Component
	trdsesslstgrp.TrdSessLstGrp
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradSesReqID(v string) { m.TradSesReqID = &v }

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
	return enum.BeginStringFIX50, "BJ", r
}
