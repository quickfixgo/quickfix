//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/linesoftextgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/routinggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a News FIX Message
type Message struct {
	FIXMsgType string `fix:"B"`
	Header     fixt11.Header
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
	//RoutingGrp Component
	RoutingGrp routinggrp.Component
	//InstrmtGrp Component
	InstrmtGrp instrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//LinesOfTextGrp Component
	LinesOfTextGrp linesoftextgrp.Component
	//URLLink is a non-required field for News.
	URLLink *string `fix:"149"`
	//RawDataLength is a non-required field for News.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for News.
	RawData *string `fix:"96"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrigTime(v time.Time)     { m.OrigTime = &v }
func (m *Message) SetUrgency(v string)         { m.Urgency = &v }
func (m *Message) SetHeadline(v string)        { m.Headline = v }
func (m *Message) SetEncodedHeadlineLen(v int) { m.EncodedHeadlineLen = &v }
func (m *Message) SetEncodedHeadline(v string) { m.EncodedHeadline = &v }
func (m *Message) SetURLLink(v string)         { m.URLLink = &v }
func (m *Message) SetRawDataLength(v int)      { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)         { m.RawData = &v }

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
	return enum.ApplVerID_FIX50SP1, "B", r
}
