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

//SecurityList msg type = y.
type SecurityList struct {
	message.Message
}

//SecurityListBuilder builds SecurityList messages.
type SecurityListBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListBuilder returns an initialized SecurityListBuilder with specified required fields.
func CreateSecurityListBuilder() SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("y"))
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m SecurityList) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m SecurityList) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityList.
func (m SecurityList) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m SecurityList) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (*field.SecurityRequestResultField, errors.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m SecurityList) GetSecurityRequestResult(f *field.SecurityRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) TotNoRelatedSym() (*field.TotNoRelatedSymField, errors.MessageRejectError) {
	f := &field.TotNoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityList.
func (m SecurityList) GetTotNoRelatedSym(f *field.TotNoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityList.
func (m SecurityList) LastFragment() (*field.LastFragmentField, errors.MessageRejectError) {
	f := &field.LastFragmentField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityList.
func (m SecurityList) GetLastFragment(f *field.LastFragmentField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityList.
func (m SecurityList) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityReportID is a non-required field for SecurityList.
func (m SecurityList) SecurityReportID() (*field.SecurityReportIDField, errors.MessageRejectError) {
	f := &field.SecurityReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityList.
func (m SecurityList) GetSecurityReportID(f *field.SecurityReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m SecurityList) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityList.
func (m SecurityList) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for SecurityList.
func (m SecurityList) MarketID() (*field.MarketIDField, errors.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from SecurityList.
func (m SecurityList) GetMarketID(f *field.MarketIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for SecurityList.
func (m SecurityList) MarketSegmentID() (*field.MarketSegmentIDField, errors.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from SecurityList.
func (m SecurityList) GetMarketSegmentID(f *field.MarketSegmentIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for SecurityList.
func (m SecurityList) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from SecurityList.
func (m SecurityList) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from SecurityList.
func (m SecurityList) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from SecurityList.
func (m SecurityList) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for SecurityList.
func (m SecurityList) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from SecurityList.
func (m SecurityList) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListID is a non-required field for SecurityList.
func (m SecurityList) SecurityListID() (*field.SecurityListIDField, errors.MessageRejectError) {
	f := &field.SecurityListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListID reads a SecurityListID from SecurityList.
func (m SecurityList) GetSecurityListID(f *field.SecurityListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListRefID is a non-required field for SecurityList.
func (m SecurityList) SecurityListRefID() (*field.SecurityListRefIDField, errors.MessageRejectError) {
	f := &field.SecurityListRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListRefID reads a SecurityListRefID from SecurityList.
func (m SecurityList) GetSecurityListRefID(f *field.SecurityListRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) SecurityListDesc() (*field.SecurityListDescField, errors.MessageRejectError) {
	f := &field.SecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListDesc reads a SecurityListDesc from SecurityList.
func (m SecurityList) GetSecurityListDesc(f *field.SecurityListDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDescLen is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLenField, errors.MessageRejectError) {
	f := &field.EncodedSecurityListDescLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDescLen reads a EncodedSecurityListDescLen from SecurityList.
func (m SecurityList) GetEncodedSecurityListDescLen(f *field.EncodedSecurityListDescLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDesc() (*field.EncodedSecurityListDescField, errors.MessageRejectError) {
	f := &field.EncodedSecurityListDescField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSecurityListDesc reads a EncodedSecurityListDesc from SecurityList.
func (m SecurityList) GetEncodedSecurityListDesc(f *field.EncodedSecurityListDescField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListType is a non-required field for SecurityList.
func (m SecurityList) SecurityListType() (*field.SecurityListTypeField, errors.MessageRejectError) {
	f := &field.SecurityListTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListType reads a SecurityListType from SecurityList.
func (m SecurityList) GetSecurityListType(f *field.SecurityListTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityListTypeSource is a non-required field for SecurityList.
func (m SecurityList) SecurityListTypeSource() (*field.SecurityListTypeSourceField, errors.MessageRejectError) {
	f := &field.SecurityListTypeSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityListTypeSource reads a SecurityListTypeSource from SecurityList.
func (m SecurityList) GetSecurityListTypeSource(f *field.SecurityListTypeSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for SecurityList.
func (m SecurityList) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SecurityList.
func (m SecurityList) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}
