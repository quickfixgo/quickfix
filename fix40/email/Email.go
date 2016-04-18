//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	fix40.Header
	//EmailType is a required field for Email.
	EmailType string `fix:"94"`
	//OrigTime is a non-required field for Email.
	OrigTime *time.Time `fix:"42"`
	//RelatdSym is a non-required field for Email.
	RelatdSym *string `fix:"46"`
	//OrderID is a non-required field for Email.
	OrderID *string `fix:"37"`
	//ClOrdID is a non-required field for Email.
	ClOrdID *string `fix:"11"`
	//LinesOfText is a required field for Email.
	LinesOfText int `fix:"33"`
	//Text is a required field for Email.
	Text string `fix:"58"`
	//RawDataLength is a non-required field for Email.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for Email.
	RawData *string `fix:"96"`
	fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized Email instance
func New(emailtype string, linesoftext int, text string) *Message {
	var m Message
	m.SetEmailType(emailtype)
	m.SetLinesOfText(linesoftext)
	m.SetText(text)
	return &m
}

func (m *Message) SetEmailType(v string)   { m.EmailType = v }
func (m *Message) SetOrigTime(v time.Time) { m.OrigTime = &v }
func (m *Message) SetRelatdSym(v string)   { m.RelatdSym = &v }
func (m *Message) SetOrderID(v string)     { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)     { m.ClOrdID = &v }
func (m *Message) SetLinesOfText(v int)    { m.LinesOfText = v }
func (m *Message) SetText(v string)        { m.Text = v }
func (m *Message) SetRawDataLength(v int)  { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)     { m.RawData = &v }

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
	return enum.BeginStringFIX40, "C", r
}
