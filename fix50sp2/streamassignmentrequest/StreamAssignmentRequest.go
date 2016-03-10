//Package streamassignmentrequest msg type = CC.
package streamassignmentrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnreqgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a StreamAssignmentRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"CC"`
	fixt11.Header
	//StreamAsgnReqID is a required field for StreamAssignmentRequest.
	StreamAsgnReqID string `fix:"1497"`
	//StreamAsgnReqType is a required field for StreamAssignmentRequest.
	StreamAsgnReqType int `fix:"1498"`
	//StrmAsgnReqGrp Component
	strmasgnreqgrp.StrmAsgnReqGrp
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetStreamAsgnReqID(v string) { m.StreamAsgnReqID = v }
func (m *Message) SetStreamAsgnReqType(v int)  { m.StreamAsgnReqType = v }

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
	return enum.ApplVerID_FIX50SP2, "CC", r
}
