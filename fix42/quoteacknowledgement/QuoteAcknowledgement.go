//Package quoteacknowledgement msg type = b.
package quoteacknowledgement

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoQuoteSets is a repeating group in QuoteAcknowledgement
type NoQuoteSets struct {
	//QuoteSetID is a non-required field for NoQuoteSets.
	QuoteSetID *string `fix:"302"`
	//UnderlyingSymbol is a non-required field for NoQuoteSets.
	UnderlyingSymbol *string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for NoQuoteSets.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for NoQuoteSets.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingIDSource is a non-required field for NoQuoteSets.
	UnderlyingIDSource *string `fix:"305"`
	//UnderlyingSecurityType is a non-required field for NoQuoteSets.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingMaturityMonthYear is a non-required field for NoQuoteSets.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDay is a non-required field for NoQuoteSets.
	UnderlyingMaturityDay *int `fix:"314"`
	//UnderlyingPutOrCall is a non-required field for NoQuoteSets.
	UnderlyingPutOrCall *int `fix:"315"`
	//UnderlyingStrikePrice is a non-required field for NoQuoteSets.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingOptAttribute is a non-required field for NoQuoteSets.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for NoQuoteSets.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for NoQuoteSets.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for NoQuoteSets.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for NoQuoteSets.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for NoQuoteSets.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for NoQuoteSets.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for NoQuoteSets.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for NoQuoteSets.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for NoQuoteSets.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
	//TotQuoteEntries is a non-required field for NoQuoteSets.
	TotQuoteEntries *int `fix:"304"`
	//NoQuoteEntries is a non-required field for NoQuoteSets.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
}

//NoQuoteEntries is a repeating group in NoQuoteSets
type NoQuoteEntries struct {
	//QuoteEntryID is a non-required field for NoQuoteEntries.
	QuoteEntryID *string `fix:"299"`
	//Symbol is a non-required field for NoQuoteEntries.
	Symbol *string `fix:"55"`
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
	//QuoteEntryRejectReason is a non-required field for NoQuoteEntries.
	QuoteEntryRejectReason *int `fix:"368"`
}

//Message is a QuoteAcknowledgement FIX Message
type Message struct {
	FIXMsgType string `fix:"b"`
	Header     fix42.Header
	//QuoteReqID is a non-required field for QuoteAcknowledgement.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a non-required field for QuoteAcknowledgement.
	QuoteID *string `fix:"117"`
	//QuoteAckStatus is a required field for QuoteAcknowledgement.
	QuoteAckStatus int `fix:"297"`
	//QuoteRejectReason is a non-required field for QuoteAcknowledgement.
	QuoteRejectReason *int `fix:"300"`
	//QuoteResponseLevel is a non-required field for QuoteAcknowledgement.
	QuoteResponseLevel *int `fix:"301"`
	//TradingSessionID is a non-required field for QuoteAcknowledgement.
	TradingSessionID *string `fix:"336"`
	//Text is a non-required field for QuoteAcknowledgement.
	Text *string `fix:"58"`
	//NoQuoteSets is a non-required field for QuoteAcknowledgement.
	NoQuoteSets []NoQuoteSets `fix:"296,omitempty"`
	Trailer     fix42.Trailer
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
	return enum.BeginStringFIX42, "b", r
}
