//Package tradingsessionlistupdatereport msg type = BS.
package tradingsessionlistupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/trdsesslstgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradingSessionListUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BS"`
	fixt11.Header
	//TradSesReqID is a non-required field for TradingSessionListUpdateReport.
	TradSesReqID *string `fix:"335"`
	//TradSesUpdateAction is a non-required field for TradingSessionListUpdateReport.
	TradSesUpdateAction *string `fix:"1327"`
	//TrdSessLstGrp is a required component for TradingSessionListUpdateReport.
	trdsesslstgrp.TrdSessLstGrp
	//ApplicationSequenceControl is a non-required component for TradingSessionListUpdateReport.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradSesReqID(v string)                       { m.TradSesReqID = &v }
func (m *Message) SetTradSesUpdateAction(v string)                { m.TradSesUpdateAction = &v }
func (m *Message) SetTrdSessLstGrp(v trdsesslstgrp.TrdSessLstGrp) { m.TrdSessLstGrp = v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}

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
	return enum.ApplVerID_FIX50SP1, "BS", r
}
