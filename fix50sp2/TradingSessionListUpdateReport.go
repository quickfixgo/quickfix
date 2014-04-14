package fix50sp2

import (
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

//NewTradingSessionListUpdateReportBuilder returns an initialized TradingSessionListUpdateReportBuilder with specified required fields.
func NewTradingSessionListUpdateReportBuilder(
	notradingsessions field.NoTradingSessions) *TradingSessionListUpdateReportBuilder {
	builder := new(TradingSessionListUpdateReportBuilder)
	builder.Body.Set(notradingsessions)
	return builder
}

//TradSesReqID is a non-required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) TradSesReqID() (*field.TradSesReqID, error) {
	f := new(field.TradSesReqID)
	err := m.Body.Get(f)
	return f, err
}

//NoTradingSessions is a required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) NoTradingSessions() (*field.NoTradingSessions, error) {
	f := new(field.NoTradingSessions)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for TradingSessionListUpdateReport.
func (m *TradingSessionListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
