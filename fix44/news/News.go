//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
	"time"
)

//NoRoutingIDs is a repeating group in News
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

//NewNoRoutingIDs returns an initialized NoRoutingIDs instance
func NewNoRoutingIDs() *NoRoutingIDs {
	var m NoRoutingIDs
	return &m
}

func (m *NoRoutingIDs) SetRoutingType(v int)  { m.RoutingType = &v }
func (m *NoRoutingIDs) SetRoutingID(v string) { m.RoutingID = &v }

//NoRelatedSym is a repeating group in News
type NoRelatedSym struct {
	//Instrument is a non-required component for NoRelatedSym.
	Instrument *instrument.Instrument
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym() *NoRelatedSym {
	var m NoRelatedSym
	return &m
}

func (m *NoRelatedSym) SetInstrument(v instrument.Instrument) { m.Instrument = &v }

//NoLegs is a repeating group in News
type NoLegs struct {
	//InstrumentLeg is a non-required component for NoLegs.
	InstrumentLeg *instrumentleg.InstrumentLeg
}

//NewNoLegs returns an initialized NoLegs instance
func NewNoLegs() *NoLegs {
	var m NoLegs
	return &m
}

func (m *NoLegs) SetInstrumentLeg(v instrumentleg.InstrumentLeg) { m.InstrumentLeg = &v }

//NoUnderlyings is a repeating group in News
type NoUnderlyings struct {
	//UnderlyingInstrument is a non-required component for NoUnderlyings.
	UnderlyingInstrument *underlyinginstrument.UnderlyingInstrument
}

//NewNoUnderlyings returns an initialized NoUnderlyings instance
func NewNoUnderlyings() *NoUnderlyings {
	var m NoUnderlyings
	return &m
}

func (m *NoUnderlyings) SetUnderlyingInstrument(v underlyinginstrument.UnderlyingInstrument) {
	m.UnderlyingInstrument = &v
}

//NoLinesOfText is a repeating group in News
type NoLinesOfText struct {
	//Text is a required field for NoLinesOfText.
	Text string `fix:"58"`
	//EncodedTextLen is a non-required field for NoLinesOfText.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoLinesOfText.
	EncodedText *string `fix:"355"`
}

//NewNoLinesOfText returns an initialized NoLinesOfText instance
func NewNoLinesOfText(text string) *NoLinesOfText {
	var m NoLinesOfText
	m.SetText(text)
	return &m
}

func (m *NoLinesOfText) SetText(v string)        { m.Text = v }
func (m *NoLinesOfText) SetEncodedTextLen(v int) { m.EncodedTextLen = &v }
func (m *NoLinesOfText) SetEncodedText(v string) { m.EncodedText = &v }

//Message is a News FIX Message
type Message struct {
	FIXMsgType string `fix:"B"`
	fix44.Header
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
	//NoRoutingIDs is a non-required field for News.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
	//NoRelatedSym is a non-required field for News.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	//NoLegs is a non-required field for News.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoUnderlyings is a non-required field for News.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLinesOfText is a required field for News.
	NoLinesOfText []NoLinesOfText `fix:"33"`
	//URLLink is a non-required field for News.
	URLLink *string `fix:"149"`
	//RawDataLength is a non-required field for News.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for News.
	RawData *string `fix:"96"`
	fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized News instance
func New(headline string, nolinesoftext []NoLinesOfText) *Message {
	var m Message
	m.SetHeadline(headline)
	m.SetNoLinesOfText(nolinesoftext)
	return &m
}

func (m *Message) SetOrigTime(v time.Time)            { m.OrigTime = &v }
func (m *Message) SetUrgency(v string)                { m.Urgency = &v }
func (m *Message) SetHeadline(v string)               { m.Headline = v }
func (m *Message) SetEncodedHeadlineLen(v int)        { m.EncodedHeadlineLen = &v }
func (m *Message) SetEncodedHeadline(v string)        { m.EncodedHeadline = &v }
func (m *Message) SetNoRoutingIDs(v []NoRoutingIDs)   { m.NoRoutingIDs = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym)   { m.NoRelatedSym = v }
func (m *Message) SetNoLegs(v []NoLegs)               { m.NoLegs = v }
func (m *Message) SetNoUnderlyings(v []NoUnderlyings) { m.NoUnderlyings = v }
func (m *Message) SetNoLinesOfText(v []NoLinesOfText) { m.NoLinesOfText = v }
func (m *Message) SetURLLink(v string)                { m.URLLink = &v }
func (m *Message) SetRawDataLength(v int)             { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)                { m.RawData = &v }

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
	return enum.BeginStringFIX44, "B", r
}
