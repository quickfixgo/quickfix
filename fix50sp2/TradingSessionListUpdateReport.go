package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//TradingSessionListUpdateReport msg type = BS.
type TradingSessionListUpdateReport struct {
	message.Message
}

//TradingSessionListUpdateReportBuilder builds TradingSessionListUpdateReport messages.
type TradingSessionListUpdateReportBuilder struct {
	message.MessageBuilder
}

//CreateTradingSessionListUpdateReportBuilder returns an initialized TradingSessionListUpdateReportBuilder with specified required fields.
func CreateTradingSessionListUpdateReportBuilder(
	notradingsessions *field.NoTradingSessionsField) TradingSessionListUpdateReportBuilder {
	var builder TradingSessionListUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BS"))
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) TradSesReqID() (*field.TradSesReqIDField, errors.MessageRejectError) {
	f := &field.TradSesReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetTradSesReqID(f *field.TradSesReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) NoTradingSessions() (*field.NoTradingSessionsField, errors.MessageRejectError) {
	f := &field.NoTradingSessionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetNoTradingSessions(f *field.NoTradingSessionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
