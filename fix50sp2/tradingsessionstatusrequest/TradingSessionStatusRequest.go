//Package tradingsessionstatusrequest msg type = g.
package tradingsessionstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a TradingSessionStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"g"`
	fixt11.Header
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
	//SecurityExchange is a non-required field for TradingSessionStatusRequest.
	SecurityExchange *string `fix:"207"`
	//MarketID is a non-required field for TradingSessionStatusRequest.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for TradingSessionStatusRequest.
	MarketSegmentID *string `fix:"1300"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetTradSesReqID(v string)            { m.TradSesReqID = v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)     { m.TradingSessionSubID = &v }
func (m *Message) SetTradSesMethod(v int)              { m.TradSesMethod = &v }
func (m *Message) SetTradSesMode(v int)                { m.TradSesMode = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = v }
func (m *Message) SetSecurityExchange(v string)        { m.SecurityExchange = &v }
func (m *Message) SetMarketID(v string)                { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string)         { m.MarketSegmentID = &v }

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
	return enum.ApplVerID_FIX50SP2, "g", r
}
