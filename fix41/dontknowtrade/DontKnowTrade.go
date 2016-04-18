//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a DontKnowTrade FIX Message
type Message struct {
	FIXMsgType string `fix:"Q"`
	fix41.Header
	//OrderID is a non-required field for DontKnowTrade.
	OrderID *string `fix:"37"`
	//ExecID is a non-required field for DontKnowTrade.
	ExecID *string `fix:"17"`
	//DKReason is a required field for DontKnowTrade.
	DKReason string `fix:"127"`
	//Symbol is a required field for DontKnowTrade.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for DontKnowTrade.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for DontKnowTrade.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for DontKnowTrade.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for DontKnowTrade.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for DontKnowTrade.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for DontKnowTrade.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for DontKnowTrade.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for DontKnowTrade.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for DontKnowTrade.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for DontKnowTrade.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for DontKnowTrade.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for DontKnowTrade.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for DontKnowTrade.
	Side string `fix:"54"`
	//OrderQty is a non-required field for DontKnowTrade.
	OrderQty *int `fix:"38"`
	//CashOrderQty is a non-required field for DontKnowTrade.
	CashOrderQty *float64 `fix:"152"`
	//LastShares is a non-required field for DontKnowTrade.
	LastShares *int `fix:"32"`
	//LastPx is a non-required field for DontKnowTrade.
	LastPx *float64 `fix:"31"`
	//Text is a non-required field for DontKnowTrade.
	Text *string `fix:"58"`
	fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized DontKnowTrade instance
func New(dkreason string, symbol string, side string) *Message {
	var m Message
	m.SetDKReason(dkreason)
	m.SetSymbol(symbol)
	m.SetSide(side)
	return &m
}

func (m *Message) SetOrderID(v string)           { m.OrderID = &v }
func (m *Message) SetExecID(v string)            { m.ExecID = &v }
func (m *Message) SetDKReason(v string)          { m.DKReason = v }
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
func (m *Message) SetLastShares(v int)           { m.LastShares = &v }
func (m *Message) SetLastPx(v float64)           { m.LastPx = &v }
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
	return enum.BeginStringFIX41, "Q", r
}
