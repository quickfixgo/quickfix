//Package quotecancel msg type = Z.
package quotecancel

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoQuoteEntries is a repeating group in QuoteCancel
type NoQuoteEntries struct {
	//Symbol is a required field for NoQuoteEntries.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NoQuoteEntries.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoQuoteEntries.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoQuoteEntries.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoQuoteEntries.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoQuoteEntries.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoQuoteEntries.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoQuoteEntries.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoQuoteEntries.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoQuoteEntries.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoQuoteEntries.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoQuoteEntries.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoQuoteEntries.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoQuoteEntries.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoQuoteEntries.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoQuoteEntries.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoQuoteEntries.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoQuoteEntries.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoQuoteEntries.
	EncodedSecurityDesc *string `fix:"351"`
	//UnderlyingSymbol is a non-required field for NoQuoteEntries.
	UnderlyingSymbol *string `fix:"311"`
}

func (m *NoQuoteEntries) SetSymbol(v string)              { m.Symbol = v }
func (m *NoQuoteEntries) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoQuoteEntries) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoQuoteEntries) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoQuoteEntries) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoQuoteEntries) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoQuoteEntries) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoQuoteEntries) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoQuoteEntries) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoQuoteEntries) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoQuoteEntries) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoQuoteEntries) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoQuoteEntries) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoQuoteEntries) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoQuoteEntries) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoQuoteEntries) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoQuoteEntries) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoQuoteEntries) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoQuoteEntries) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *NoQuoteEntries) SetUnderlyingSymbol(v string)    { m.UnderlyingSymbol = &v }

//Message is a QuoteCancel FIX Message
type Message struct {
	FIXMsgType string `fix:"Z"`
	fix42.Header
	//QuoteReqID is a non-required field for QuoteCancel.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for QuoteCancel.
	QuoteID string `fix:"117"`
	//QuoteCancelType is a required field for QuoteCancel.
	QuoteCancelType int `fix:"298"`
	//QuoteResponseLevel is a non-required field for QuoteCancel.
	QuoteResponseLevel *int `fix:"301"`
	//TradingSessionID is a non-required field for QuoteCancel.
	TradingSessionID *string `fix:"336"`
	//NoQuoteEntries is a required field for QuoteCancel.
	NoQuoteEntries []NoQuoteEntries `fix:"295"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)               { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)                  { m.QuoteID = v }
func (m *Message) SetQuoteCancelType(v int)             { m.QuoteCancelType = v }
func (m *Message) SetQuoteResponseLevel(v int)          { m.QuoteResponseLevel = &v }
func (m *Message) SetTradingSessionID(v string)         { m.TradingSessionID = &v }
func (m *Message) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }

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
	return enum.BeginStringFIX42, "Z", r
}
