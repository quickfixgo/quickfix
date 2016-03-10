//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"time"
)

//NoRoutingIDs is a repeating group in Email
type NoRoutingIDs struct {
	//RoutingType is a non-required field for NoRoutingIDs.
	RoutingType *int `fix:"216"`
	//RoutingID is a non-required field for NoRoutingIDs.
	RoutingID *string `fix:"217"`
}

func (m *NoRoutingIDs) SetRoutingType(v int)  { m.RoutingType = &v }
func (m *NoRoutingIDs) SetRoutingID(v string) { m.RoutingID = &v }

//NoRelatedSym is a repeating group in Email
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
}

//LinesOfText is a repeating group in Email
type LinesOfText struct {
	//Text is a required field for LinesOfText.
	Text string `fix:"58"`
	//EncodedTextLen is a non-required field for LinesOfText.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for LinesOfText.
	EncodedText *string `fix:"355"`
}

func (m *LinesOfText) SetText(v string)        { m.Text = v }
func (m *LinesOfText) SetEncodedTextLen(v int) { m.EncodedTextLen = &v }
func (m *LinesOfText) SetEncodedText(v string) { m.EncodedText = &v }

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	fix43.Header
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
	//NoRoutingIDs is a non-required field for Email.
	NoRoutingIDs []NoRoutingIDs `fix:"215,omitempty"`
	//NoRelatedSym is a non-required field for Email.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	//OrderID is a non-required field for Email.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for Email.
	ClOrdID *string `fix:"11"`
	//LinesOfText is a required field for Email.
	LinesOfText []LinesOfText `fix:"33"`
	//RawDataLength is a non-required field for Email.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Email.
	RawData *string `fix:"96"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEmailThreadID(v string)        { m.EmailThreadID = v }
func (m *Message) SetEmailType(v string)            { m.EmailType = v }
func (m *Message) SetOrigTime(v time.Time)          { m.OrigTime = &v }
func (m *Message) SetSubject(v string)              { m.Subject = v }
func (m *Message) SetEncodedSubjectLen(v int)       { m.EncodedSubjectLen = &v }
func (m *Message) SetEncodedSubject(v string)       { m.EncodedSubject = &v }
func (m *Message) SetNoRoutingIDs(v []NoRoutingIDs) { m.NoRoutingIDs = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
func (m *Message) SetOrderID(v string)              { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)              { m.ClOrdID = &v }
func (m *Message) SetLinesOfText(v []LinesOfText)   { m.LinesOfText = v }
func (m *Message) SetRawDataLength(v int)           { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)              { m.RawData = &v }

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
	return enum.BeginStringFIX43, "C", r
}
