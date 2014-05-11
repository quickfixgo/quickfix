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
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("BK"))
	return builder
}

//SecurityReportID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReportID() (*field.SecurityReportIDField, errors.MessageRejectError) {
	f := &field.SecurityReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityReportID(f *field.SecurityReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReqID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResultField, errors.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityRequestResult(f *field.SecurityRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateActionField, errors.MessageRejectError) {
	f := &field.SecurityUpdateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityUpdateAction reads a SecurityUpdateAction from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityUpdateAction(f *field.SecurityUpdateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CorporateAction is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) CorporateAction() (*field.CorporateActionField, errors.MessageRejectError) {
	f := &field.CorporateActionField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCorporateAction reads a CorporateAction from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetCorporateAction(f *field.CorporateActionField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListID() (*field.SecurityListIDField, errors.MessageRejectError) {
	f := &field.SecurityListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListID reads a SecurityListID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityListID(f *field.SecurityListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListRefID is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListRefID() (*field.SecurityListRefIDField, errors.MessageRejectError) {
	f := &field.SecurityListRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListRefID reads a SecurityListRefID from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityListRefID(f *field.SecurityListRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListDesc() (*field.SecurityListDescField, errors.MessageRejectError) {
	f := &field.SecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListDesc reads a SecurityListDesc from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityListDesc(f *field.SecurityListDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDescLen is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityListDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDescLen reads a EncodedSecurityListDescLen from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetEncodedSecurityListDescLen(f *field.EncodedSecurityListDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDesc is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) EncodedSecurityListDesc() (*field.EncodedSecurityListDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDesc reads a EncodedSecurityListDesc from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetEncodedSecurityListDesc(f *field.EncodedSecurityListDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListType is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListType() (*field.SecurityListTypeField, errors.MessageRejectError) {
	f := &field.SecurityListTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListType reads a SecurityListType from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityListType(f *field.SecurityListTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListTypeSource is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) SecurityListTypeSource() (*field.SecurityListTypeSourceField, errors.MessageRejectError) {
	f := &field.SecurityListTypeSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListTypeSource reads a SecurityListTypeSource from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetSecurityListTypeSource(f *field.SecurityListTypeSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityListUpdateReport.
func (m SecurityListUpdateReport) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityListUpdateReport.
func (m SecurityListUpdateReport) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
