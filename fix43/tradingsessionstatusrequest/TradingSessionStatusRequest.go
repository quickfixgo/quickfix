//Package tradingsessionstatusrequest msg type = g.
package tradingsessionstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
)

//Message is a TradingSessionStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"g"`
	Header     fix43.Header
	//TradSesReqID is a required field for TradingSessionStatusRequest.
	TradSesReqID string `fix:"335"`
	//TradingSessionID is a non-required field for TradingSessionStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for TradingSessionStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	//TradSesMethod is a non-required field for TradingSessionStatusRequest.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for TradingSessionStatusRequest.
	TradSesMode *int `fix:"339"`
	//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
	SubscriptionRequestType string `fix:"263"`
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
	return enum.BeginStringFIX43, "g", r
}
