//Package ordercancelrequest msg type = F.
package ordercancelrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a OrderCancelRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"F"`
	fix41.Header
	//OrigClOrdID is a required field for OrderCancelRequest.
	OrigClOrdID string `fix:"41"`
	//OrderID is a non-required field for OrderCancelRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderCancelRequest.
	ClOrdID string `fix:"11"`
	//ListID is a non-required field for OrderCancelRequest.
	ListID *string `fix:"66"`
	//ClientID is a non-required field for OrderCancelRequest.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderCancelRequest.
	ExecBroker *string `fix:"76"`
	//Symbol is a required field for OrderCancelRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for OrderCancelRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for OrderCancelRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for OrderCancelRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for OrderCancelRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for OrderCancelRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for OrderCancelRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for OrderCancelRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for OrderCancelRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for OrderCancelRequest.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for OrderCancelRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for OrderCancelRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for OrderCancelRequest.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for OrderCancelRequest.
	Side string `fix:"54"`
	//OrderQty is a non-required field for OrderCancelRequest.
	OrderQty *int `fix:"38"`
	//CashOrderQty is a non-required field for OrderCancelRequest.
	CashOrderQty *float64 `fix:"152"`
	//Text is a non-required field for OrderCancelRequest.
	Text *string `fix:"58"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrigClOrdID(v string)       { m.OrigClOrdID = v }
func (m *Message) SetOrderID(v string)           { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)           { m.ClOrdID = v }
func (m *Message) SetListID(v string)            { m.ListID = &v }
func (m *Message) SetClientID(v string)          { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)        { m.ExecBroker = &v }
func (m *Message) SetSymbol(v string)            { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)         { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)        { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)          { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)      { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string) { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)          { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)            { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)      { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)      { m.OptAttribute = &v }
func (m *Message) SetSecurityExchange(v string)  { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)            { m.Issuer = &v }
func (m *Message) SetSecurityDesc(v string)      { m.SecurityDesc = &v }
func (m *Message) SetSide(v string)              { m.Side = v }
func (m *Message) SetOrderQty(v int)             { m.OrderQty = &v }
func (m *Message) SetCashOrderQty(v float64)     { m.CashOrderQty = &v }
func (m *Message) SetText(v string)              { m.Text = &v }

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
	return enum.BeginStringFIX41, "F", r
}
