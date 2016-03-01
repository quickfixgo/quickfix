//Package tradingsessionlistrequest msg type = BI.
package tradingsessionlistrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradingSessionListRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BI"`
	Header     fixt11.Header
	//TradSesReqID is a required field for TradingSessionListRequest.
	TradSesReqID string `fix:"335"`
	//TradingSessionID is a non-required field for TradingSessionListRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for TradingSessionListRequest.
	TradingSessionSubID *string `fix:"625"`
	//SecurityExchange is a non-required field for TradingSessionListRequest.
	SecurityExchange *string `fix:"207"`
	//TradSesMethod is a non-required field for TradingSessionListRequest.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for TradingSessionListRequest.
	TradSesMode *int `fix:"339"`
	//SubscriptionRequestType is a required field for TradingSessionListRequest.
	SubscriptionRequestType string `fix:"263"`
	//MarketID is a non-required field for TradingSessionListRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for TradingSessionListRequest.
	MarketSegmentID *string `fix:"1300"`
	Trailer         fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "BI", r
}
