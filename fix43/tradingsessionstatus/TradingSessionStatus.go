//Package tradingsessionstatus msg type = h.
package tradingsessionstatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a TradingSessionStatus wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m Message) TradSesReqID() (*field.TradSesReqIDField, quickfix.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatus.
func (m Message) GetTradSesReqID(f *field.TradSesReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m Message) TradingSessionID() (*field.TradingSessionIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatus.
func (m Message) GetTradingSessionID(f *field.TradingSessionIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionSubID is a non-required field for TradingSessionStatus.
func (m Message) TradingSessionSubID() (*field.TradingSessionSubIDField, quickfix.MessageRejectError) {
	f := &field.TradingSessionSubIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionSubID reads a TradingSessionSubID from TradingSessionStatus.
func (m Message) GetTradingSessionSubID(f *field.TradingSessionSubIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m Message) TradSesMethod() (*field.TradSesMethodField, quickfix.MessageRejectError) {
	f := &field.TradSesMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatus.
func (m Message) GetTradSesMethod(f *field.TradSesMethodField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m Message) TradSesMode() (*field.TradSesModeField, quickfix.MessageRejectError) {
	f := &field.TradSesModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatus.
func (m Message) GetTradSesMode(f *field.TradSesModeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m Message) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, quickfix.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradingSessionStatus.
func (m Message) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m Message) TradSesStatus() (*field.TradSesStatusField, quickfix.MessageRejectError) {
	f := &field.TradSesStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatus reads a TradSesStatus from TradingSessionStatus.
func (m Message) GetTradSesStatus(f *field.TradSesStatusField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatusRejReason is a non-required field for TradingSessionStatus.
func (m Message) TradSesStatusRejReason() (*field.TradSesStatusRejReasonField, quickfix.MessageRejectError) {
	f := &field.TradSesStatusRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatusRejReason reads a TradSesStatusRejReason from TradingSessionStatus.
func (m Message) GetTradSesStatusRejReason(f *field.TradSesStatusRejReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesStartTime() (*field.TradSesStartTimeField, quickfix.MessageRejectError) {
	f := &field.TradSesStartTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStartTime reads a TradSesStartTime from TradingSessionStatus.
func (m Message) GetTradSesStartTime(f *field.TradSesStartTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesOpenTime() (*field.TradSesOpenTimeField, quickfix.MessageRejectError) {
	f := &field.TradSesOpenTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesOpenTime reads a TradSesOpenTime from TradingSessionStatus.
func (m Message) GetTradSesOpenTime(f *field.TradSesOpenTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesPreCloseTime() (*field.TradSesPreCloseTimeField, quickfix.MessageRejectError) {
	f := &field.TradSesPreCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesPreCloseTime reads a TradSesPreCloseTime from TradingSessionStatus.
func (m Message) GetTradSesPreCloseTime(f *field.TradSesPreCloseTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesCloseTime() (*field.TradSesCloseTimeField, quickfix.MessageRejectError) {
	f := &field.TradSesCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesCloseTime reads a TradSesCloseTime from TradingSessionStatus.
func (m Message) GetTradSesCloseTime(f *field.TradSesCloseTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m Message) TradSesEndTime() (*field.TradSesEndTimeField, quickfix.MessageRejectError) {
	f := &field.TradSesEndTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesEndTime reads a TradSesEndTime from TradingSessionStatus.
func (m Message) GetTradSesEndTime(f *field.TradSesEndTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m Message) TotalVolumeTraded() (*field.TotalVolumeTradedField, quickfix.MessageRejectError) {
	f := &field.TotalVolumeTradedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalVolumeTraded reads a TotalVolumeTraded from TradingSessionStatus.
func (m Message) GetTotalVolumeTraded(f *field.TotalVolumeTradedField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradingSessionStatus.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradingSessionStatus.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradingSessionStatus.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradingSessionStatus.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds TradingSessionStatus messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for TradingSessionStatus.
func Builder(
	tradingsessionid *field.TradingSessionIDField,
	tradsesstatus *field.TradSesStatusField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("h"))
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX43, "h", r
}
