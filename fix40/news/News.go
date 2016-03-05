//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a News FIX Message
type Message struct {
	FIXMsgType string `fix:"B"`
	fix40.Header
	//OrigTime is a non-required field for News.
	OrigTime *time.Time `fix:"42"`
	//Urgency is a non-required field for News.
	Urgency *string `fix:"61"`
	//RelatdSym is a non-required field for News.
	RelatdSym *string `fix:"46"`
	//LinesOfText is a required field for News.
	LinesOfText int `fix:"33"`
	//Text is a required field for News.
	Text string `fix:"58"`
	//RawDataLength is a non-required field for News.
	RawDataLength *int `fix:"95"`
	//RawData is a non-required field for News.
	RawData *string `fix:"96"`
	fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrigTime(v time.Time) { m.OrigTime = &v }
func (m *Message) SetUrgency(v string)     { m.Urgency = &v }
func (m *Message) SetRelatdSym(v string)   { m.RelatdSym = &v }
func (m *Message) SetLinesOfText(v int)    { m.LinesOfText = v }
func (m *Message) SetText(v string)        { m.Text = v }
func (m *Message) SetRawDataLength(v int)  { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)     { m.RawData = &v }

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
	return enum.BeginStringFIX40, "B", r
}
