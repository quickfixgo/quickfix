//Package tradingsessionstatusrequest msg type = g.
package tradingsessionstatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
)

//Message is a TradingSessionStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"g"`
	fix42.Header
	//TradSesReqID is a required field for TradingSessionStatusRequest.
	TradSesReqID string `fix:"335"`
	//TradingSessionID is a non-required field for TradingSessionStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradSesMethod is a non-required field for TradingSessionStatusRequest.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for TradingSessionStatusRequest.
	TradSesMode *int `fix:"339"`
	//SubscriptionRequestType is a required field for TradingSessionStatusRequest.
	SubscriptionRequestType string `fix:"263"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized TradingSessionStatusRequest instance
func New(tradsesreqid string, subscriptionrequesttype string) *Message {
	var m Message
	m.SetTradSesReqID(tradsesreqid)
	m.SetSubscriptionRequestType(subscriptionrequesttype)
	return &m
}

func (m *Message) SetTradSesReqID(v string)            { m.TradSesReqID = v }
func (m *Message) SetTradingSessionID(v string)        { m.TradingSessionID = &v }
func (m *Message) SetTradSesMethod(v int)              { m.TradSesMethod = &v }
func (m *Message) SetTradSesMode(v int)                { m.TradSesMode = &v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = v }

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
	return enum.BeginStringFIX42, "g", r
}
