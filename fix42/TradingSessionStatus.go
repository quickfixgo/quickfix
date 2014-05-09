package fix42

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TradingSessionStatus msg type = h.
type TradingSessionStatus struct {
	message.Message
}

//TradingSessionStatusBuilder builds TradingSessionStatus messages.
type TradingSessionStatusBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionStatusBuilder returns an initialized TradingSessionStatusBuilder with specified required fields.
func CreateTradingSessionStatusBuilder(
	tradingsessionid *field.TradingSessionIDField,
	tradsesstatus *field.TradSesStatusField) TradingSessionStatusBuilder {
	var builder TradingSessionStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX42))
	builder.Header.Set(field.NewMsgType("h"))
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionID() (*field.TradingSessionIDField, errors.MessageRejectError) {
	f := &field.TradingSessionIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradingSessionID reads a TradingSessionID from TradingSessionStatus.
func (m TradingSessionStatus) GetTradingSessionID(f *field.TradingSessionIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMethod() (*field.TradSesMethodField, errors.MessageRejectError) {
	f := &field.TradSesMethodField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMethod reads a TradSesMethod from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesMethod(f *field.TradSesMethodField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMode() (*field.TradSesModeField, errors.MessageRejectError) {
	f := &field.TradSesModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesMode reads a TradSesMode from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesMode(f *field.TradSesModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnsolicitedIndicator() (*field.UnsolicitedIndicatorField, errors.MessageRejectError) {
	f := &field.UnsolicitedIndicatorField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUnsolicitedIndicator reads a UnsolicitedIndicator from TradingSessionStatus.
func (m TradingSessionStatus) GetUnsolicitedIndicator(f *field.UnsolicitedIndicatorField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatus() (*field.TradSesStatusField, errors.MessageRejectError) {
	f := &field.TradSesStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStatus reads a TradSesStatus from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesStatus(f *field.TradSesStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStartTime() (*field.TradSesStartTimeField, errors.MessageRejectError) {
	f := &field.TradSesStartTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesStartTime reads a TradSesStartTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesStartTime(f *field.TradSesStartTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesOpenTime() (*field.TradSesOpenTimeField, errors.MessageRejectError) {
	f := &field.TradSesOpenTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesOpenTime reads a TradSesOpenTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesOpenTime(f *field.TradSesOpenTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesPreCloseTime() (*field.TradSesPreCloseTimeField, errors.MessageRejectError) {
	f := &field.TradSesPreCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesPreCloseTime reads a TradSesPreCloseTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesPreCloseTime(f *field.TradSesPreCloseTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesCloseTime() (*field.TradSesCloseTimeField, errors.MessageRejectError) {
	f := &field.TradSesCloseTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesCloseTime reads a TradSesCloseTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesCloseTime(f *field.TradSesCloseTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesEndTime() (*field.TradSesEndTimeField, errors.MessageRejectError) {
	f := &field.TradSesEndTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesEndTime reads a TradSesEndTime from TradingSessionStatus.
func (m TradingSessionStatus) GetTradSesEndTime(f *field.TradSesEndTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TotalVolumeTraded() (*field.TotalVolumeTradedField, errors.MessageRejectError) {
	f := &field.TotalVolumeTradedField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalVolumeTraded reads a TotalVolumeTraded from TradingSessionStatus.
func (m TradingSessionStatus) GetTotalVolumeTraded(f *field.TotalVolumeTradedField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from TradingSessionStatus.
func (m TradingSessionStatus) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from TradingSessionStatus.
func (m TradingSessionStatus) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
