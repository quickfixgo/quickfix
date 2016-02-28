//Package crossordercancelrequest msg type = u.
package crossordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/orderqtydata"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//NoSides is a repeating group in CrossOrderCancelRequest
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrigClOrdID is a required field for NoSides.
	OrigClOrdID string `fix:"41"`
	//ClOrdID is a required field for NoSides.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NoSides.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for NoSides.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties Component
	Parties parties.Component
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//OrderQtyData Component
	OrderQtyData orderqtydata.Component
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
}

func (m *NoSides) SetSide(v string)                 { m.Side = v }
func (m *NoSides) SetOrigClOrdID(v string)          { m.OrigClOrdID = v }
func (m *NoSides) SetClOrdID(v string)              { m.ClOrdID = v }
func (m *NoSides) SetSecondaryClOrdID(v string)     { m.SecondaryClOrdID = &v }
func (m *NoSides) SetClOrdLinkID(v string)          { m.ClOrdLinkID = &v }
func (m *NoSides) SetOrigOrdModTime(v time.Time)    { m.OrigOrdModTime = &v }
func (m *NoSides) SetTradeOriginationDate(v string) { m.TradeOriginationDate = &v }
func (m *NoSides) SetComplianceID(v string)         { m.ComplianceID = &v }
func (m *NoSides) SetText(v string)                 { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)          { m.EncodedText = &v }

//Message is a CrossOrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"u"`
	Header     fix43.Header
	//OrderID is a non-required field for CrossOrderCancelRequest.
	OrderID *string `fix:"37"`
	//CrossID is a required field for CrossOrderCancelRequest.
	CrossID string `fix:"548"`
	//OrigCrossID is a required field for CrossOrderCancelRequest.
	OrigCrossID string `fix:"551"`
	//CrossType is a required field for CrossOrderCancelRequest.
	CrossType int `fix:"549"`
	//CrossPrioritization is a required field for CrossOrderCancelRequest.
	CrossPrioritization int `fix:"550"`
	//NoSides is a required field for CrossOrderCancelRequest.
	NoSides []NoSides `fix:"552"`
	//Instrument Component
	Instrument instrument.Component
	//TransactTime is a required field for CrossOrderCancelRequest.
	TransactTime time.Time `fix:"60"`
	Trailer      fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)          { m.OrderID = &v }
func (m *Message) SetCrossID(v string)          { m.CrossID = v }
func (m *Message) SetOrigCrossID(v string)      { m.OrigCrossID = v }
func (m *Message) SetCrossType(v int)           { m.CrossType = v }
func (m *Message) SetCrossPrioritization(v int) { m.CrossPrioritization = v }
func (m *Message) SetNoSides(v []NoSides)       { m.NoSides = v }
func (m *Message) SetTransactTime(v time.Time)  { m.TransactTime = v }

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
	return enum.BeginStringFIX43, "u", r
}
