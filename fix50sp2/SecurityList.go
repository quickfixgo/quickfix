package fix50sp2

import (
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

//NewSecurityListBuilder returns an initialized SecurityListBuilder with specified required fields.
func NewSecurityListBuilder() *SecurityListBuilder {
	builder := new(SecurityListBuilder)
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m *SecurityList) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityList.
func (m *SecurityList) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m *SecurityList) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m *SecurityList) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for SecurityList.
func (m *SecurityList) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityList.
func (m *SecurityList) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//SecurityReportID is a non-required field for SecurityList.
func (m *SecurityList) SecurityReportID() (*field.SecurityReportID, error) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m *SecurityList) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//MarketID is a non-required field for SecurityList.
func (m *SecurityList) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}

//MarketSegmentID is a non-required field for SecurityList.
func (m *SecurityList) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}

//ApplID is a non-required field for SecurityList.
func (m *SecurityList) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}

//ApplSeqNum is a non-required field for SecurityList.
func (m *SecurityList) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplLastSeqNum is a non-required field for SecurityList.
func (m *SecurityList) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}

//ApplResendFlag is a non-required field for SecurityList.
func (m *SecurityList) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListID is a non-required field for SecurityList.
func (m *SecurityList) SecurityListID() (*field.SecurityListID, error) {
	f := new(field.SecurityListID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListRefID is a non-required field for SecurityList.
func (m *SecurityList) SecurityListRefID() (*field.SecurityListRefID, error) {
	f := new(field.SecurityListRefID)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListDesc is a non-required field for SecurityList.
func (m *SecurityList) SecurityListDesc() (*field.SecurityListDesc, error) {
	f := new(field.SecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityListDescLen is a non-required field for SecurityList.
func (m *SecurityList) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLen, error) {
	f := new(field.EncodedSecurityListDescLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedSecurityListDesc is a non-required field for SecurityList.
func (m *SecurityList) EncodedSecurityListDesc() (*field.EncodedSecurityListDesc, error) {
	f := new(field.EncodedSecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListType is a non-required field for SecurityList.
func (m *SecurityList) SecurityListType() (*field.SecurityListType, error) {
	f := new(field.SecurityListType)
	err := m.Body.Get(f)
	return f, err
}

//SecurityListTypeSource is a non-required field for SecurityList.
func (m *SecurityList) SecurityListTypeSource() (*field.SecurityListTypeSource, error) {
	f := new(field.SecurityListTypeSource)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for SecurityList.
func (m *SecurityList) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
