//Package securitytypes msg type = w.
package securitytypes

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
)

//NoSecurityTypes is a repeating group in SecurityTypes
type NoSecurityTypes struct {
	//SecurityType is a non-required field for NoSecurityTypes.
	SecurityType *string `fix:"167"`
	//Product is a non-required field for NoSecurityTypes.
	Product *int `fix:"460"`
	//CFICode is a non-required field for NoSecurityTypes.
	CFICode *string `fix:"461"`
}

//Message is a SecurityTypes FIX Message
type Message struct {
	FIXMsgType string `fix:"w"`
	Header     fix43.Header
	//SecurityReqID is a required field for SecurityTypes.
	SecurityReqID string `fix:"320"`
	//SecurityResponseID is a required field for SecurityTypes.
	SecurityResponseID string `fix:"322"`
	//SecurityResponseType is a required field for SecurityTypes.
	SecurityResponseType int `fix:"323"`
	//TotalNumSecurityTypes is a non-required field for SecurityTypes.
	TotalNumSecurityTypes *int `fix:"557"`
	//NoSecurityTypes is a non-required field for SecurityTypes.
	NoSecurityTypes []NoSecurityTypes `fix:"558,omitempty"`
	//Text is a non-required field for SecurityTypes.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityTypes.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityTypes.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for SecurityTypes.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityTypes.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for SecurityTypes.
	SubscriptionRequestType *string `fix:"263"`
	Trailer                 fix43.Trailer
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
	return enum.BeginStringFIX43, "w", r
}
