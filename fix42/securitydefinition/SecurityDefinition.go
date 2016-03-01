//Package securitydefinition msg type = d.
package securitydefinition

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//NoRelatedSym is a repeating group in SecurityDefinition
type NoRelatedSym struct {
	//UnderlyingSymbol is a non-required field for NoRelatedSym.
	UnderlyingSymbol *string `fix:"311"`
	//UnderlyingSymbolSfx is a non-required field for NoRelatedSym.
	UnderlyingSymbolSfx *string `fix:"312"`
	//UnderlyingSecurityID is a non-required field for NoRelatedSym.
	UnderlyingSecurityID *string `fix:"309"`
	//UnderlyingIDSource is a non-required field for NoRelatedSym.
	UnderlyingIDSource *string `fix:"305"`
	//UnderlyingSecurityType is a non-required field for NoRelatedSym.
	UnderlyingSecurityType *string `fix:"310"`
	//UnderlyingMaturityMonthYear is a non-required field for NoRelatedSym.
	UnderlyingMaturityMonthYear *string `fix:"313"`
	//UnderlyingMaturityDay is a non-required field for NoRelatedSym.
	UnderlyingMaturityDay *int `fix:"314"`
	//UnderlyingPutOrCall is a non-required field for NoRelatedSym.
	UnderlyingPutOrCall *int `fix:"315"`
	//UnderlyingStrikePrice is a non-required field for NoRelatedSym.
	UnderlyingStrikePrice *float64 `fix:"316"`
	//UnderlyingOptAttribute is a non-required field for NoRelatedSym.
	UnderlyingOptAttribute *string `fix:"317"`
	//UnderlyingContractMultiplier is a non-required field for NoRelatedSym.
	UnderlyingContractMultiplier *float64 `fix:"436"`
	//UnderlyingCouponRate is a non-required field for NoRelatedSym.
	UnderlyingCouponRate *float64 `fix:"435"`
	//UnderlyingSecurityExchange is a non-required field for NoRelatedSym.
	UnderlyingSecurityExchange *string `fix:"308"`
	//UnderlyingIssuer is a non-required field for NoRelatedSym.
	UnderlyingIssuer *string `fix:"306"`
	//EncodedUnderlyingIssuerLen is a non-required field for NoRelatedSym.
	EncodedUnderlyingIssuerLen *int `fix:"362"`
	//EncodedUnderlyingIssuer is a non-required field for NoRelatedSym.
	EncodedUnderlyingIssuer *string `fix:"363"`
	//UnderlyingSecurityDesc is a non-required field for NoRelatedSym.
	UnderlyingSecurityDesc *string `fix:"307"`
	//EncodedUnderlyingSecurityDescLen is a non-required field for NoRelatedSym.
	EncodedUnderlyingSecurityDescLen *int `fix:"364"`
	//EncodedUnderlyingSecurityDesc is a non-required field for NoRelatedSym.
	EncodedUnderlyingSecurityDesc *string `fix:"365"`
	//RatioQty is a non-required field for NoRelatedSym.
	RatioQty *float64 `fix:"319"`
	//Side is a non-required field for NoRelatedSym.
	Side *string `fix:"54"`
	//UnderlyingCurrency is a non-required field for NoRelatedSym.
	UnderlyingCurrency *string `fix:"318"`
}

//Message is a SecurityDefinition FIX Message
type Message struct {
	FIXMsgType string `fix:"d"`
	Header     fix42.Header
	//SecurityReqID is a required field for SecurityDefinition.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for SecurityDefinition.
	SecurityResponseID string `fix:"322"`
	//SecurityResponseType is a non-required field for SecurityDefinition.
	SecurityResponseType *int `fix:"323"`
	//TotalNumSecurities is a required field for SecurityDefinition.
	TotalNumSecurities int `fix:"393"`
	//Symbol is a non-required field for SecurityDefinition.
	Symbol *string `fix:"55"`
	//SymbolSfx is a non-required field for SecurityDefinition.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for SecurityDefinition.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for SecurityDefinition.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for SecurityDefinition.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for SecurityDefinition.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for SecurityDefinition.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for SecurityDefinition.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for SecurityDefinition.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for SecurityDefinition.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for SecurityDefinition.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for SecurityDefinition.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for SecurityDefinition.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for SecurityDefinition.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for SecurityDefinition.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for SecurityDefinition.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for SecurityDefinition.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for SecurityDefinition.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for SecurityDefinition.
	EncodedSecurityDesc *string `fix:"351"`
	//Currency is a non-required field for SecurityDefinition.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityDefinition.
	TradingSessionID *string `fix:"336"`
	//Text is a non-required field for SecurityDefinition.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityDefinition.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityDefinition.
	EncodedText *string `fix:"355"`
	//NoRelatedSym is a non-required field for SecurityDefinition.
	NoRelatedSym []NoRelatedSym `fix:"146,omitempty"`
	Trailer      fix42.Trailer
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
	return enum.BeginStringFIX42, "d", r
}
