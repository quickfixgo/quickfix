package fix42

import (
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
	tradingsessionid field.TradingSessionID,
	tradsesstatus field.TradSesStatus) TradingSessionStatusBuilder {
	var builder TradingSessionStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(tradingsessionid)
	builder.Body.Set(tradsesstatus)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesReqID() (field.TradSesReqID, error) {
	var f field.TradSesReqID
	err := m.Body.Get(&f)
	return f, err
}

//TradingSessionID is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradingSessionID() (field.TradingSessionID, error) {
	var f field.TradingSessionID
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMethod is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMethod() (field.TradSesMethod, error) {
	var f field.TradSesMethod
	err := m.Body.Get(&f)
	return f, err
}

//TradSesMode is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesMode() (field.TradSesMode, error) {
	var f field.TradSesMode
	err := m.Body.Get(&f)
	return f, err
}

//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) UnsolicitedIndicator() (field.UnsolicitedIndicator, error) {
	var f field.UnsolicitedIndicator
	err := m.Body.Get(&f)
	return f, err
}

//TradSesStatus is a required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStatus() (field.TradSesStatus, error) {
	var f field.TradSesStatus
	err := m.Body.Get(&f)
	return f, err
}

//TradSesStartTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesStartTime() (field.TradSesStartTime, error) {
	var f field.TradSesStartTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesOpenTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesOpenTime() (field.TradSesOpenTime, error) {
	var f field.TradSesOpenTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesPreCloseTime() (field.TradSesPreCloseTime, error) {
	var f field.TradSesPreCloseTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesCloseTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesCloseTime() (field.TradSesCloseTime, error) {
	var f field.TradSesCloseTime
	err := m.Body.Get(&f)
	return f, err
}

//TradSesEndTime is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TradSesEndTime() (field.TradSesEndTime, error) {
	var f field.TradSesEndTime
	err := m.Body.Get(&f)
	return f, err
}

//TotalVolumeTraded is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) TotalVolumeTraded() (field.TotalVolumeTraded, error) {
	var f field.TotalVolumeTraded
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for TradingSessionStatus.
func (m TradingSessionStatus) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
