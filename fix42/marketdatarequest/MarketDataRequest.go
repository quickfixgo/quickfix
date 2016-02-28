//Package marketdatarequest msg type = V.
package marketdatarequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoMDEntryTypes is a repeating group in MarketDataRequest
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

func (m *NoMDEntryTypes) SetMDEntryType(v string) { m.MDEntryType = v }

//NoRelatedSym is a repeating group in MarketDataRequest
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
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
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
func (m *NoRelatedSym) SetTradingSessionID(v string)    { m.TradingSessionID = &v }

//Message is a MarketDataRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"V"`
	Header     fix42.Header
	//MDReqID is a required field for MarketDataRequest.
	MDReqID string `fix:"262"`
	//SubscriptionRequestType is a required field for MarketDataRequest.
	SubscriptionRequestType string `fix:"263"`
	//MarketDepth is a required field for MarketDataRequest.
	MarketDepth int `fix:"264"`
	//MDUpdateType is a non-required field for MarketDataRequest.
	MDUpdateType *int `fix:"265"`
	//AggregatedBook is a non-required field for MarketDataRequest.
	AggregatedBook *bool `fix:"266"`
	//NoMDEntryTypes is a required field for MarketDataRequest.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
	//NoRelatedSym is a required field for MarketDataRequest.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	Trailer      fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMDReqID(v string)                  { m.MDReqID = v }
func (m *Message) SetSubscriptionRequestType(v string)  { m.SubscriptionRequestType = v }
func (m *Message) SetMarketDepth(v int)                 { m.MarketDepth = v }
func (m *Message) SetMDUpdateType(v int)                { m.MDUpdateType = &v }
func (m *Message) SetAggregatedBook(v bool)             { m.AggregatedBook = &v }
func (m *Message) SetNoMDEntryTypes(v []NoMDEntryTypes) { m.NoMDEntryTypes = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym)     { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX42, "V", r
}
