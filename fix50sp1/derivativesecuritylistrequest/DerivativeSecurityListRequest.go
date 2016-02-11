//Package derivativesecuritylistrequest msg type = z.
package derivativesecuritylistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/derivativeinstrument"
	"github.com/quickfixgo/quickfix/fix50sp1/underlyinginstrument"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a DerivativeSecurityListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"z"`
	Header     fixt11.Header
	//SecurityReqID is a required field for DerivativeSecurityListRequest.
	SecurityReqID string `fix:"320"`
	//SecurityListRequestType is a required field for DerivativeSecurityListRequest.
	SecurityListRequestType int `fix:"559"`
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
	//SecuritySubType is a non-required field for DerivativeSecurityListRequest.
	SecuritySubType *string `fix:"762"`
	//Currency is a non-required field for DerivativeSecurityListRequest.
	Currency *string `fix:"15"`
	//Text is a non-required field for DerivativeSecurityListRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for DerivativeSecurityListRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for DerivativeSecurityListRequest.
	EncodedText *string `fix:"355"`
	//TradingSessionID is a non-required field for DerivativeSecurityListRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for DerivativeSecurityListRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for DerivativeSecurityListRequest.
	SubscriptionRequestType *string `fix:"263"`
	//MarketID is a non-required field for DerivativeSecurityListRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for DerivativeSecurityListRequest.
	MarketSegmentID *string `fix:"1300"`
	//DerivativeInstrument Component
	DerivativeInstrument derivativeinstrument.Component
	Trailer              fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.ApplVerID_FIX50SP1, "z", r
}
