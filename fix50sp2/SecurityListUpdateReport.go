package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type SecurityListUpdateReport struct {
	quickfix.Message
}

func (m *SecurityListUpdateReport) SecurityResponseID() (*field.SecurityResponseID, error) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityRequestResult() (*field.SecurityRequestResult, error) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) CorporateAction() (*field.CorporateAction, error) {
	f := new(field.CorporateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) EncodedSecurityListDescLen() (*field.EncodedSecurityListDescLen, error) {
	f := new(field.EncodedSecurityListDescLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) TotNoRelatedSym() (*field.TotNoRelatedSym, error) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityReportID() (*field.SecurityReportID, error) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityUpdateAction() (*field.SecurityUpdateAction, error) {
	f := new(field.SecurityUpdateAction)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityListID() (*field.SecurityListID, error) {
	f := new(field.SecurityListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityListRefID() (*field.SecurityListRefID, error) {
	f := new(field.SecurityListRefID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityListDesc() (*field.SecurityListDesc, error) {
	f := new(field.SecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) EncodedSecurityListDesc() (*field.EncodedSecurityListDesc, error) {
	f := new(field.EncodedSecurityListDesc)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityListType() (*field.SecurityListType, error) {
	f := new(field.SecurityListType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListUpdateReport) SecurityListTypeSource() (*field.SecurityListTypeSource, error) {
	f := new(field.SecurityListTypeSource)
	err := m.Body.Get(f)
	return f, err
}
