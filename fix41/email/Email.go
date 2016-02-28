//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
	"time"
)

//NoRelatedSym is a repeating group in Email
type NoRelatedSym struct {
	//RelatdSym is a non-required field for NoRelatedSym.
	RelatdSym *string `fix:"46"`
	//SymbolSfx is a non-required field for NoRelatedSym.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoRelatedSym.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoRelatedSym.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoRelatedSym.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoRelatedSym.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoRelatedSym.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoRelatedSym.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoRelatedSym.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoRelatedSym.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for NoRelatedSym.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoRelatedSym.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for NoRelatedSym.
	SecurityDesc *string `fix:"107"`
}

func (m *NoRelatedSym) SetRelatdSym(v string)         { m.RelatdSym = &v }
func (m *NoRelatedSym) SetSymbolSfx(v string)         { m.SymbolSfx = &v }
func (m *NoRelatedSym) SetSecurityID(v string)        { m.SecurityID = &v }
func (m *NoRelatedSym) SetIDSource(v string)          { m.IDSource = &v }
func (m *NoRelatedSym) SetSecurityType(v string)      { m.SecurityType = &v }
func (m *NoRelatedSym) SetMaturityMonthYear(v string) { m.MaturityMonthYear = &v }
func (m *NoRelatedSym) SetMaturityDay(v int)          { m.MaturityDay = &v }
func (m *NoRelatedSym) SetPutOrCall(v int)            { m.PutOrCall = &v }
func (m *NoRelatedSym) SetStrikePrice(v float64)      { m.StrikePrice = &v }
func (m *NoRelatedSym) SetOptAttribute(v string)      { m.OptAttribute = &v }
func (m *NoRelatedSym) SetSecurityExchange(v string)  { m.SecurityExchange = &v }
func (m *NoRelatedSym) SetIssuer(v string)            { m.Issuer = &v }
func (m *NoRelatedSym) SetSecurityDesc(v string)      { m.SecurityDesc = &v }

//LinesOfText is a repeating group in Email
type LinesOfText struct {
	//Text is a required field for LinesOfText.
	Text string `fix:"58"`
}

func (m *LinesOfText) SetText(v string) { m.Text = v }

//Message is a Email FIX Message
type Message struct {
	FIXMsgType string `fix:"C"`
	Header     fix41.Header
	//EmailThreadID is a required field for Email.
	EmailThreadID string `fix:"164"`
	//EmailType is a required field for Email.
	EmailType string `fix:"94"`
	//OrigTime is a non-required field for Email.
	OrigTime *time.Time `fix:"42"`
	//Subject is a required field for Email.
	Subject string `fix:"147"`
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
	Trailer fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetEmailThreadID(v string)        { m.EmailThreadID = v }
func (m *Message) SetEmailType(v string)            { m.EmailType = v }
func (m *Message) SetOrigTime(v time.Time)          { m.OrigTime = &v }
func (m *Message) SetSubject(v string)              { m.Subject = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }
func (m *Message) SetOrderID(v string)              { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)              { m.ClOrdID = &v }
func (m *Message) SetLinesOfText(v []LinesOfText)   { m.LinesOfText = v }
func (m *Message) SetRawDataLength(v int)           { m.RawDataLength = &v }
func (m *Message) SetRawData(v string)              { m.RawData = &v }

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
	return enum.BeginStringFIX41, "C", r
}
