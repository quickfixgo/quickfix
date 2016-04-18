//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoRelatedSym is a repeating group in QuoteRequest
type NoRelatedSym struct {
	//Symbol is a required field for NoRelatedSym.
	Symbol string `fix:"55"`
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
	//ContractMultiplier is a non-required field for NoRelatedSym.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoRelatedSym.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoRelatedSym.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoRelatedSym.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoRelatedSym.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoRelatedSym.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoRelatedSym.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoRelatedSym.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoRelatedSym.
	EncodedSecurityDesc *string `fix:"351"`
	//PrevClosePx is a non-required field for NoRelatedSym.
	PrevClosePx *float64 `fix:"140"`
	//QuoteRequestType is a non-required field for NoRelatedSym.
	QuoteRequestType *int `fix:"303"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//Side is a non-required field for NoRelatedSym.
	Side *string `fix:"54"`
	//OrderQty is a non-required field for NoRelatedSym.
	OrderQty *float64 `fix:"38"`
	//FutSettDate is a non-required field for NoRelatedSym.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for NoRelatedSym.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for NoRelatedSym.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NoRelatedSym.
	OrderQty2 *float64 `fix:"192"`
	//ExpireTime is a non-required field for NoRelatedSym.
	ExpireTime *time.Time `fix:"126"`
	//TransactTime is a non-required field for NoRelatedSym.
	TransactTime *time.Time `fix:"60"`
	//Currency is a non-required field for NoRelatedSym.
	Currency *string `fix:"15"`
}

//NewNoRelatedSym returns an initialized NoRelatedSym instance
func NewNoRelatedSym(symbol string) *NoRelatedSym {
	var m NoRelatedSym
	m.SetSymbol(symbol)
	return &m
}

func (m *NoRelatedSym) SetSymbol(v string)              { m.Symbol = v }
func (m *NoRelatedSym) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoRelatedSym) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoRelatedSym) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoRelatedSym) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoRelatedSym) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoRelatedSym) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoRelatedSym) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoRelatedSym) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoRelatedSym) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoRelatedSym) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoRelatedSym) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoRelatedSym) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoRelatedSym) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoRelatedSym) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoRelatedSym) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoRelatedSym) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoRelatedSym) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoRelatedSym) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *NoRelatedSym) SetPrevClosePx(v float64)        { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)       { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetSide(v string)                { m.Side = &v }
func (m *NoRelatedSym) SetOrderQty(v float64)           { m.OrderQty = &v }
func (m *NoRelatedSym) SetFutSettDate(v string)         { m.FutSettDate = &v }
func (m *NoRelatedSym) SetOrdType(v string)             { m.OrdType = &v }
func (m *NoRelatedSym) SetFutSettDate2(v string)        { m.FutSettDate2 = &v }
func (m *NoRelatedSym) SetOrderQty2(v float64)          { m.OrderQty2 = &v }
func (m *NoRelatedSym) SetExpireTime(v time.Time)       { m.ExpireTime = &v }
func (m *NoRelatedSym) SetTransactTime(v time.Time)     { m.TransactTime = &v }
func (m *NoRelatedSym) SetCurrency(v string)            { m.Currency = &v }

//Message is a QuoteRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"R"`
	fix42.Header
	//QuoteReqID is a required field for QuoteRequest.
	QuoteReqID string `fix:"131"`
	//NoRelatedSym is a required field for QuoteRequest.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)           { m.QuoteReqID = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym) { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX42, "R", r
}
