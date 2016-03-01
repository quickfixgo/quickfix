//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	Header     fix42.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//QuoteResponseLevel is a non-required field for Quote.
	QuoteResponseLevel *int `fix:"301"`
	//TradingSessionID is a non-required field for Quote.
	TradingSessionID *string `fix:"336"`
	//Symbol is a required field for Quote.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Quote.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Quote.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Quote.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for Quote.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for Quote.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for Quote.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for Quote.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for Quote.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for Quote.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for Quote.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for Quote.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for Quote.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Quote.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for Quote.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for Quote.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for Quote.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for Quote.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for Quote.
	EncodedSecurityDesc *string `fix:"351"`
	//BidPx is a non-required field for Quote.
	BidPx *float64 `fix:"132"`
	//OfferPx is a non-required field for Quote.
	OfferPx *float64 `fix:"133"`
	//BidSize is a non-required field for Quote.
	BidSize *float64 `fix:"134"`
	//OfferSize is a non-required field for Quote.
	OfferSize *float64 `fix:"135"`
	//ValidUntilTime is a non-required field for Quote.
	ValidUntilTime *time.Time `fix:"62"`
	//BidSpotRate is a non-required field for Quote.
	BidSpotRate *float64 `fix:"188"`
	//OfferSpotRate is a non-required field for Quote.
	OfferSpotRate *float64 `fix:"190"`
	//BidForwardPoints is a non-required field for Quote.
	BidForwardPoints *float64 `fix:"189"`
	//OfferForwardPoints is a non-required field for Quote.
	OfferForwardPoints *float64 `fix:"191"`
	//TransactTime is a non-required field for Quote.
	TransactTime *time.Time `fix:"60"`
	//FutSettDate is a non-required field for Quote.
	FutSettDate *string `fix:"64"`
	//OrdType is a non-required field for Quote.
	OrdType *string `fix:"40"`
	//FutSettDate2 is a non-required field for Quote.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for Quote.
	OrderQty2 *float64 `fix:"192"`
	//Currency is a non-required field for Quote.
	Currency *string `fix:"15"`
	Trailer  fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.BeginStringFIX42, "S", r
}
