package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityListUpdateReport msg type = BK.
type SecurityListUpdateReport struct {
	message.Message
}

//SecurityListUpdateReportBuilder builds SecurityListUpdateReport messages.
type SecurityListUpdateReportBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListUpdateReportBuilder returns an initialized SecurityListUpdateReportBuilder with specified required fields.
func CreateSecurityListUpdateReportBuilder() SecurityListUpdateReportBuilder {
	var builder SecurityListUpdateReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReportID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReportID() (*field.SecurityReportID, errors.MessageRejectError) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityReportID(f *field.SecurityReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReqID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, errors.MessageRejectError) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityRequestResult(f *field.SecurityRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, errors.MessageRejectError) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetTotNoRelatedSym(f *field.TotNoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, errors.MessageRejectError) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityUpdateAction(f *field.SecurityUpdateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) CorporateAction() (*field.CorporateAction, errors.MessageRejectError) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetCorporateAction(f *field.CorporateAction) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketID() (*field.MarketID, errors.MessageRejectError) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetMarketID(f *field.MarketID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketSegmentID() (*field.MarketSegmentID, errors.MessageRejectError) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetMarketSegmentID(f *field.MarketSegmentID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplID() (*field.ApplID, errors.MessageRejectError) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplID(f *field.ApplID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, errors.MessageRejectError) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplSeqNum(f *field.ApplSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, errors.MessageRejectError) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNum) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, errors.MessageRejectError) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplResendFlag(f *field.ApplResendFlag) errors.MessageRejectError {
	return m.Body.Get(f)
}
