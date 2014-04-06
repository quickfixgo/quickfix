package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type SecurityList struct {
	quickfix.Message
}

func (m *SecurityList) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityReportID() (*field.SecurityReportID, error) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityListID() (*field.SecurityListID, error) {
	f := new(field.SecurityListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityListDesc() (*field.SecurityListDesc, error) {
	f := new(field.SecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLen, error) {
	f := new(field.EncodedSecurityListDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) EncodedSecurityListDesc() (*field.EncodedSecurityListDesc, error) {
	f := new(field.EncodedSecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityListRefID() (*field.SecurityListRefID, error) {
	f := new(field.SecurityListRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityListType() (*field.SecurityListType, error) {
	f := new(field.SecurityListType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) SecurityListTypeSource() (*field.SecurityListTypeSource, error) {
	f := new(field.SecurityListTypeSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityList) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
