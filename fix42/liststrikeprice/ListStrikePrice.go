//Package liststrikeprice msg type = m.
package liststrikeprice

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoStrikes is a repeating group in ListStrikePrice
type NoStrikes struct {
	//Symbol is a required field for NoStrikes.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NoStrikes.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoStrikes.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoStrikes.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoStrikes.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoStrikes.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoStrikes.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoStrikes.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoStrikes.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoStrikes.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoStrikes.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoStrikes.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoStrikes.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoStrikes.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoStrikes.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoStrikes.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoStrikes.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoStrikes.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoStrikes.
	EncodedSecurityDesc *string `fix:"351"`
	//PrevClosePx is a non-required field for NoStrikes.
	PrevClosePx *float64 `fix:"140"`
	//ClOrdID is a non-required field for NoStrikes.
	ClOrdID *string `fix:"11"`
	//Side is a non-required field for NoStrikes.
	Side *string `fix:"54"`
	//Price is a required field for NoStrikes.
	Price float64 `fix:"44"`
	//Currency is a non-required field for NoStrikes.
	Currency *string `fix:"15"`
	//Text is a non-required field for NoStrikes.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoStrikes.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoStrikes.
	EncodedText *string `fix:"355"`
}

//NewNoStrikes returns an initialized NoStrikes instance
func NewNoStrikes(symbol string, price float64) *NoStrikes {
	var m NoStrikes
	m.SetSymbol(symbol)
	m.SetPrice(price)
	return &m
}

func (m *NoStrikes) SetSymbol(v string)              { m.Symbol = v }
func (m *NoStrikes) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoStrikes) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoStrikes) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoStrikes) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoStrikes) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoStrikes) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoStrikes) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoStrikes) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoStrikes) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoStrikes) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoStrikes) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoStrikes) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoStrikes) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoStrikes) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoStrikes) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoStrikes) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoStrikes) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoStrikes) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *NoStrikes) SetPrevClosePx(v float64)        { m.PrevClosePx = &v }
func (m *NoStrikes) SetClOrdID(v string)             { m.ClOrdID = &v }
func (m *NoStrikes) SetSide(v string)                { m.Side = &v }
func (m *NoStrikes) SetPrice(v float64)              { m.Price = v }
func (m *NoStrikes) SetCurrency(v string)            { m.Currency = &v }
func (m *NoStrikes) SetText(v string)                { m.Text = &v }
func (m *NoStrikes) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *NoStrikes) SetEncodedText(v string)         { m.EncodedText = &v }

//Message is a ListStrikePrice FIX Message
type Message struct {
	FIXMsgType string `fix:"m"`
	fix42.Header
	//ListID is a required field for ListStrikePrice.
	ListID string `fix:"66"`
	//TotNoStrikes is a required field for ListStrikePrice.
	TotNoStrikes int `fix:"422"`
	//NoStrikes is a required field for ListStrikePrice.
	NoStrikes []NoStrikes `fix:"428"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)         { m.ListID = v }
func (m *Message) SetTotNoStrikes(v int)      { m.TotNoStrikes = v }
func (m *Message) SetNoStrikes(v []NoStrikes) { m.NoStrikes = v }

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
	return enum.BeginStringFIX42, "m", r
}
