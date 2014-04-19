package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m SecurityList) SecurityReqID() (field.SecurityReqID, errors.MessageRejectError) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityList.
func (m SecurityList) SecurityResponseID() (field.SecurityResponseID, errors.MessageRejectError) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (field.SecurityRequestResult, errors.MessageRejectError) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) TotNoRelatedSym() (field.TotNoRelatedSym, errors.MessageRejectError) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for SecurityList.
func (m SecurityList) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//SecurityReportID is a non-required field for SecurityList.
func (m SecurityList) SecurityReportID() (field.SecurityReportID, errors.MessageRejectError) {
	var f field.SecurityReportID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m SecurityList) ClearingBusinessDate() (field.ClearingBusinessDate, errors.MessageRejectError) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//MarketID is a non-required field for SecurityList.
func (m SecurityList) MarketID() (field.MarketID, errors.MessageRejectError) {
	var f field.MarketID
	err := m.Body.Get(&f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityList.
func (m SecurityList) MarketSegmentID() (field.MarketSegmentID, errors.MessageRejectError) {
	var f field.MarketSegmentID
	err := m.Body.Get(&f)
	return f, err
}

//ApplID is a non-required field for SecurityList.
func (m SecurityList) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityList.
func (m SecurityList) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityList.
func (m SecurityList) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListID is a non-required field for SecurityList.
func (m SecurityList) SecurityListID() (field.SecurityListID, errors.MessageRejectError) {
	var f field.SecurityListID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListRefID is a non-required field for SecurityList.
func (m SecurityList) SecurityListRefID() (field.SecurityListRefID, errors.MessageRejectError) {
	var f field.SecurityListRefID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) SecurityListDesc() (field.SecurityListDesc, errors.MessageRejectError) {
	var f field.SecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDescLen is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDescLen() (field.EncodedSecurityListDescLen, errors.MessageRejectError) {
	var f field.EncodedSecurityListDescLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedSecurityListDesc is a non-required field for SecurityList.
func (m SecurityList) EncodedSecurityListDesc() (field.EncodedSecurityListDesc, errors.MessageRejectError) {
	var f field.EncodedSecurityListDesc
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListType is a non-required field for SecurityList.
func (m SecurityList) SecurityListType() (field.SecurityListType, errors.MessageRejectError) {
	var f field.SecurityListType
	err := m.Body.Get(&f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityList.
func (m SecurityList) SecurityListTypeSource() (field.SecurityListTypeSource, errors.MessageRejectError) {
	var f field.SecurityListTypeSource
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for SecurityList.
func (m SecurityList) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}
