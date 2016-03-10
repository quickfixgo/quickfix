//Package quotestatusrequest msg type = a.
package quotestatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//Message is a QuoteStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"a"`
	fix42.Header
	//QuoteID is a non-required field for QuoteStatusRequest.
	QuoteID *string `fix:"117"`
	//Symbol is a required field for QuoteStatusRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for QuoteStatusRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for QuoteStatusRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for QuoteStatusRequest.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for QuoteStatusRequest.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for QuoteStatusRequest.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for QuoteStatusRequest.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for QuoteStatusRequest.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for QuoteStatusRequest.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for QuoteStatusRequest.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for QuoteStatusRequest.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for QuoteStatusRequest.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for QuoteStatusRequest.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for QuoteStatusRequest.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for QuoteStatusRequest.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for QuoteStatusRequest.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for QuoteStatusRequest.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for QuoteStatusRequest.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for QuoteStatusRequest.
	EncodedSecurityDesc *string `fix:"351"`
	//Side is a non-required field for QuoteStatusRequest.
	Side *string `fix:"54"`
	//TradingSessionID is a non-required field for QuoteStatusRequest.
	TradingSessionID *string `fix:"336"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteID(v string)             { m.QuoteID = &v }
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
func (m *Message) SetSide(v string)                { m.Side = &v }
func (m *Message) SetTradingSessionID(v string)    { m.TradingSessionID = &v }

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
	return enum.BeginStringFIX42, "a", r
}
