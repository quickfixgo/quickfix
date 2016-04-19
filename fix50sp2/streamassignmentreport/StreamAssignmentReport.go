//Package streamassignmentreport msg type = CD.
package streamassignmentreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/strmasgnrptgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a StreamAssignmentReport FIX Message
type Message struct {
	FIXMsgType string `fix:"CD"`
	fixt11.Header
	//StreamAsgnRptID is a required field for StreamAssignmentReport.
	StreamAsgnRptID string `fix:"1501"`
	//StreamAsgnReqType is a non-required field for StreamAssignmentReport.
	StreamAsgnReqType *int `fix:"1498"`
	//StreamAsgnReqID is a non-required field for StreamAssignmentReport.
	StreamAsgnReqID *string `fix:"1497"`
	//StrmAsgnRptGrp is a non-required component for StreamAssignmentReport.
	StrmAsgnRptGrp *strmasgnrptgrp.StrmAsgnRptGrp
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized StreamAssignmentReport instance
func New(streamasgnrptid string) *Message {
	var m Message
	m.SetStreamAsgnRptID(streamasgnrptid)
	return &m
}

func (m *Message) SetStreamAsgnRptID(v string)                       { m.StreamAsgnRptID = v }
func (m *Message) SetStreamAsgnReqType(v int)                        { m.StreamAsgnReqType = &v }
func (m *Message) SetStreamAsgnReqID(v string)                       { m.StreamAsgnReqID = &v }
func (m *Message) SetStrmAsgnRptGrp(v strmasgnrptgrp.StrmAsgnRptGrp) { m.StrmAsgnRptGrp = &v }

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
	return enum.ApplVerID_FIX50SP2, "CD", r
}
