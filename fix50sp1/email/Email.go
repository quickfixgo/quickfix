//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/linesoftextgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/routinggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	fixt11.Header
	//EmailThreadID is a required field for Email.
	EmailThreadID string `fix:"164"`
	//EmailType is a required field for Email.
	EmailType string `fix:"94"`
	//OrigTime is a non-required field for Email.
	OrigTime *time.Time `fix:"42"`
	//Subject is a required field for Email.
	Subject string `fix:"147"`
	//EncodedSubjectLen is a non-required field for Email.
	EncodedSubjectLen *int `fix:"356"`
	//EncodedSubject is a non-required field for Email.
	EncodedSubject *string `fix:"357"`
	//RoutingGrp is a non-required component for Email.
	RoutingGrp *routinggrp.RoutingGrp
	//InstrmtGrp is a non-required component for Email.
	InstrmtGrp *instrmtgrp.InstrmtGrp
	//UndInstrmtGrp is a non-required component for Email.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//InstrmtLegGrp is a non-required component for Email.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//OrderID is a non-required field for Email.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for Email.
	ClOrdID *string `fix:"11"`
	//LinesOfTextGrp is a required component for Email.
	linesoftextgrp.LinesOfTextGrp
	//RawDataLength is a non-required field for Email.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Email.
	RawData *string `fix:"96"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEmailThreadID(v string)                         { m.EmailThreadID = v }
func (m *Message) SetEmailType(v string)                             { m.EmailType = v }
func (m *Message) SetOrigTime(v time.Time)                           { m.OrigTime = &v }
func (m *Message) SetSubject(v string)                               { m.Subject = v }
func (m *Message) SetEncodedSubjectLen(v int)                        { m.EncodedSubjectLen = &v }
func (m *Message) SetEncodedSubject(v string)                        { m.EncodedSubject = &v }
func (m *Message) SetRoutingGrp(v routinggrp.RoutingGrp)             { m.RoutingGrp = &v }
func (m *Message) SetInstrmtGrp(v instrmtgrp.InstrmtGrp)             { m.InstrmtGrp = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)    { m.UndInstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)    { m.InstrmtLegGrp = &v }
func (m *Message) SetOrderID(v string)                               { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)                               { m.ClOrdID = &v }
func (m *Message) SetLinesOfTextGrp(v linesoftextgrp.LinesOfTextGrp) { m.LinesOfTextGrp = v }
func (m *Message) SetRawDataLength(v int)                            { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)                               { m.RawData = &v }

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
	return enum.ApplVerID_FIX50SP1, "C", r
}
