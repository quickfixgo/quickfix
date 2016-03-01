//Package networkcounterpartysystemstatusrequest msg type = BC.
package networkcounterpartysystemstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/compidreqgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a NetworkCounterpartySystemStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BC"`
	Header     fixt11.Header
	//NetworkRequestType is a required field for NetworkCounterpartySystemStatusRequest.
	NetworkRequestType int `fix:"935"`
	//NetworkRequestID is a required field for NetworkCounterpartySystemStatusRequest.
	NetworkRequestID string `fix:"933"`
	//CompIDReqGrp Component
	CompIDReqGrp compidreqgrp.Component
	Trailer      fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "BC", r
}
