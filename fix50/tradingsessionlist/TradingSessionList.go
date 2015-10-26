//Package tradingsessionlist msg type = BJ.
package tradingsessionlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a TradingSessionList wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//TradSesReqID is a non-required field for TradingSessionList.
func (m Message) TradSesReqID() (*field.TradSesReqIDField, quickfix.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionList.
func (m Message) GetTradSesReqID(f *field.TradSesReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionList.
func (m Message) NoTradingSessions() (*field.NoTradingSessionsField, quickfix.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionList.
func (m Message) GetNoTradingSessions(f *field.NoTradingSessionsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for TradingSessionList.
func New(
	notradingsessions *field.NoTradingSessionsField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.NewMsgType("BJ"))
	builder.Body.Set(notradingsessions)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX50, "BJ", r
}
