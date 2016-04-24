//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/instrmtgrp"
	"github.com/quickfixgo/quickfix/fix50/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50/linesoftextgrp"
	"github.com/quickfixgo/quickfix/fix50/routinggrp"
	"github.com/quickfixgo/quickfix/fix50/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a News FIX Message
type Message struct {
	FIXMsgType string `fix:"B"`
	fixt11.Header
	//OrigTime is a non-required field for News.
	OrigTime *time.Time `fix:"42"`
	//Urgency is a non-required field for News.
	Urgency *string `fix:"61"`
	//Headline is a required field for News.
	Headline string `fix:"148"`
	//EncodedHeadlineLen is a non-required field for News.
	EncodedHeadlineLen *int `fix:"358"`
	//EncodedHeadline is a non-required field for News.
	EncodedHeadline *string `fix:"359"`
	//RoutingGrp is a non-required component for News.
	RoutingGrp *routinggrp.RoutingGrp
	//InstrmtGrp is a non-required component for News.
	InstrmtGrp *instrmtgrp.InstrmtGrp
	//InstrmtLegGrp is a non-required component for News.
	InstrmtLegGrp *instrmtleggrp.InstrmtLegGrp
	//UndInstrmtGrp is a non-required component for News.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//LinesOfTextGrp is a required component for News.
	linesoftextgrp.LinesOfTextGrp
	//URLLink is a non-required field for News.
	URLLink *string `fix:"149"`
	//RawDataLength is a non-required field for News.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for News.
	RawData *string `fix:"96"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized News instance
func New(headline string, linesoftextgrp linesoftextgrp.LinesOfTextGrp) *Message {
	var m Message
	m.SetHeadline(headline)
	m.SetLinesOfTextGrp(linesoftextgrp)
	return &m
}

func (m *Message) SetOrigTime(v time.Time)                           { m.OrigTime = &v }
func (m *Message) SetUrgency(v string)                               { m.Urgency = &v }
func (m *Message) SetHeadline(v string)                              { m.Headline = v }
func (m *Message) SetEncodedHeadlineLen(v int)                       { m.EncodedHeadlineLen = &v }
func (m *Message) SetEncodedHeadline(v string)                       { m.EncodedHeadline = &v }
func (m *Message) SetRoutingGrp(v routinggrp.RoutingGrp)             { m.RoutingGrp = &v }
func (m *Message) SetInstrmtGrp(v instrmtgrp.InstrmtGrp)             { m.InstrmtGrp = &v }
func (m *Message) SetInstrmtLegGrp(v instrmtleggrp.InstrmtLegGrp)    { m.InstrmtLegGrp = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp)    { m.UndInstrmtGrp = &v }
func (m *Message) SetLinesOfTextGrp(v linesoftextgrp.LinesOfTextGrp) { m.LinesOfTextGrp = v }
func (m *Message) SetURLLink(v string)                               { m.URLLink = &v }
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
	return enum.ApplVerID_FIX50, "B", r
}
