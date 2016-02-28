//Package orderstatusrequest msg type = H.
package orderstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix41"
)

//Message is a OrderStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"H"`
	Header     fix41.Header
	//OrderID is a non-required field for OrderStatusRequest.
	OrderID *string `fix:"37"`
	//ClOrdID is a required field for OrderStatusRequest.
	ClOrdID string `fix:"11"`
	//ClientID is a non-required field for OrderStatusRequest.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for OrderStatusRequest.
	ExecBroker *string `fix:"76"`
	//Symbol is a required field for OrderStatusRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for OrderStatusRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for OrderStatusRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for OrderStatusRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for OrderStatusRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for OrderStatusRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for OrderStatusRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for OrderStatusRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for OrderStatusRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for OrderStatusRequest.
	OptAttribute *string `fix:"206"`
	//SecurityExchange is a non-required field for OrderStatusRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for OrderStatusRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for OrderStatusRequest.
	SecurityDesc *string `fix:"107"`
	//Side is a required field for OrderStatusRequest.
	Side    string `fix:"54"`
	Trailer fix41.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)           { m.OrderID = &v }
func (m *Message) SetClOrdID(v string)           { m.ClOrdID = v }
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
	return enum.BeginStringFIX41, "H", r
}
