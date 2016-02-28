//Package networkcounterpartysystemstatusresponse msg type = BD.
package networkcounterpartysystemstatusresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/compidstatgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a NetworkCounterpartySystemStatusResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"BD"`
	Header     fixt11.Header
	//NetworkStatusResponseType is a required field for NetworkCounterpartySystemStatusResponse.
	NetworkStatusResponseType int `fix:"937"`
	//NetworkRequestID is a non-required field for NetworkCounterpartySystemStatusResponse.
	NetworkRequestID *string `fix:"933"`
	//NetworkResponseID is a required field for NetworkCounterpartySystemStatusResponse.
	NetworkResponseID string `fix:"932"`
	//LastNetworkResponseID is a non-required field for NetworkCounterpartySystemStatusResponse.
	LastNetworkResponseID *string `fix:"934"`
	//CompIDStatGrp Component
	CompIDStatGrp compidstatgrp.Component
	Trailer       fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetNetworkStatusResponseType(v int) { m.NetworkStatusResponseType = v }
func (m *Message) SetNetworkRequestID(v string)       { m.NetworkRequestID = &v }
func (m *Message) SetNetworkResponseID(v string)      { m.NetworkResponseID = v }
func (m *Message) SetLastNetworkResponseID(v string)  { m.LastNetworkResponseID = &v }

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
	return enum.BeginStringFIX50, "BD", r
}
