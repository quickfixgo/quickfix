package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	notradingsessions field.NoTradingSessions) TradingSessionListUpdateReportBuilder {
	var builder TradingSessionListUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("BS"))
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) TradSesReqID() (*field.TradSesReqID, errors.MessageRejectError) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesReqID reads a TradSesReqID from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetTradSesReqID(f *field.TradSesReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradSesUpdateAction is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) TradSesUpdateAction() (*field.TradSesUpdateAction, errors.MessageRejectError) {
	f := new(field.TradSesUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetTradSesUpdateAction reads a TradSesUpdateAction from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetTradSesUpdateAction(f *field.TradSesUpdateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoTradingSessions is a required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) NoTradingSessions() (*field.NoTradingSessions, errors.MessageRejectError) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoTradingSessions reads a NoTradingSessions from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetNoTradingSessions(f *field.NoTradingSessions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from TradingSessionListUpdateReport.
func (m TradingSessionListUpdateReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
