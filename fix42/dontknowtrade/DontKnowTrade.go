//Package dontknowtrade msg type = Q.
package dontknowtrade

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//Message is a DontKnowTrade FIX Message
type Message struct {
	FIXMsgType string `fix:"Q"`
	fix42.Header
	//OrderID is a required field for DontKnowTrade.
	OrderID string `fix:"37"`
	//ExecID is a required field for DontKnowTrade.
	ExecID string `fix:"17"`
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
	//ContractMultiplier is a non-required field for DontKnowTrade.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for DontKnowTrade.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for DontKnowTrade.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for DontKnowTrade.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for DontKnowTrade.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for DontKnowTrade.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for DontKnowTrade.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for DontKnowTrade.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for DontKnowTrade.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a required field for DontKnowTrade.
	Side string `fix:"54"`
	//OrderQty is a non-required field for DontKnowTrade.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for DontKnowTrade.
	CashOrderQty *float64 `fix:"152"`
	//LastShares is a non-required field for DontKnowTrade.
	LastShares *float64 `fix:"32"`
	//LastPx is a non-required field for DontKnowTrade.
	LastPx *float64 `fix:"31"`
	//Text is a non-required field for DontKnowTrade.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for DontKnowTrade.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for DontKnowTrade.
	EncodedText *string `fix:"355"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)             { m.OrderID = v }
func (m *Message) SetExecID(v string)              { m.ExecID = v }
func (m *Message) SetDKReason(v string)            { m.DKReason = v }
func (m *Message) SetSymbol(v string)              { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)            { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *Message) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *Message) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *Message) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)              { m.Issuer = &v }
func (m *Message) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *Message) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *Message) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *Message) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *Message) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *Message) SetSide(v string)                { m.Side = v }
func (m *Message) SetOrderQty(v float64)           { m.OrderQty = &v }
func (m *Message) SetCashOrderQty(v float64)       { m.CashOrderQty = &v }
func (m *Message) SetLastShares(v float64)         { m.LastShares = &v }
func (m *Message) SetLastPx(v float64)             { m.LastPx = &v }
func (m *Message) SetText(v string)                { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)         { m.EncodedText = &v }

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
	return enum.BeginStringFIX42, "Q", r
}
