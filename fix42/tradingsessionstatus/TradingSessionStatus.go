//Package tradingsessionstatus msg type = h.
package tradingsessionstatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a TradingSessionStatus wrapper for the generic Message type
type Message struct {
	message.Message
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m Message) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatus.
func (m Message) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatus.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m Message) TradSesMethod() (*field.TradSesMethodField, errors.MessageRejectError) {
	f := &field.TradSesMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatus.
func (m Message) GetTradSesMethod(f *field.TradSesMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m Message) TradSesMode() (*field.TradSesModeField, errors.MessageRejectError) {
	f := &field.TradSesModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatus.
func (m Message) GetTradSesMode(f *field.TradSesModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m Message) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, errors.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradingSessionStatus.
func (m Message) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m Message) TradSesStatus() (*field.TradSesStatusField, errors.MessageRejectError) {
	f := &field.TradSesStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatus reads a TradSesStatus from TradingSessionStatus.
func (m Message) GetTradSesStatus(f *field.TradSesStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesStartTime() (*field.TradSesStartTimeField, errors.MessageRejectError) {
	f := &field.TradSesStartTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStartTime reads a TradSesStartTime from TradingSessionStatus.
func (m Message) GetTradSesStartTime(f *field.TradSesStartTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesOpenTime() (*field.TradSesOpenTimeField, errors.MessageRejectError) {
	f := &field.TradSesOpenTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesOpenTime reads a TradSesOpenTime from TradingSessionStatus.
func (m Message) GetTradSesOpenTime(f *field.TradSesOpenTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesPreCloseTime() (*field.TradSesPreCloseTimeField, errors.MessageRejectError) {
	f := &field.TradSesPreCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesPreCloseTime reads a TradSesPreCloseTime from TradingSessionStatus.
func (m Message) GetTradSesPreCloseTime(f *field.TradSesPreCloseTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesCloseTime() (*field.TradSesCloseTimeField, errors.MessageRejectError) {
	f := &field.TradSesCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesCloseTime reads a TradSesCloseTime from TradingSessionStatus.
func (m Message) GetTradSesCloseTime(f *field.TradSesCloseTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesEndTime() (*field.TradSesEndTimeField, errors.MessageRejectError) {
	f := &field.TradSesEndTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesEndTime reads a TradSesEndTime from TradingSessionStatus.
func (m Message) GetTradSesEndTime(f *field.TradSesEndTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m Message) TotalVolumeTraded() (*field.TotalVolumeTradedField, errors.MessageRejectError) {
	f := &field.TotalVolumeTradedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalVolumeTraded reads a TotalVolumeTraded from TradingSessionStatus.
func (m Message) GetTotalVolumeTraded(f *field.TotalVolumeTradedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradingSessionStatus.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradingSessionStatus.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradingSessionStatus.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradingSessionStatus.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds TradingSessionStatus messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for TradingSessionStatus.
func Builder(
	tradingsessionid *field.TradingSessionIDField,
	tradsesstatus *field.TradSesStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header().Set(field.NewMsgType("h"))
	builder.Body().Set(tradingsessionid)
	builder.Body().Set(tradsesstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX42, "h", r
}
