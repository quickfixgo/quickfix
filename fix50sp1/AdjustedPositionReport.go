package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AdjustedPositionReport msg type = BL.
type AdjustedPositionReport struct {
	message.Message
}

//AdjustedPositionReportBuilder builds AdjustedPositionReport messages.
type AdjustedPositionReportBuilder struct {
	message.MessageBuilder
}

//NewAdjustedPositionReportBuilder returns an initialized AdjustedPositionReportBuilder with specified required fields.
func NewAdjustedPositionReportBuilder(
	posmaintrptid field.PosMaintRptID,
	clearingbusinessdate field.ClearingBusinessDate) *AdjustedPositionReportBuilder {
	builder := new(AdjustedPositionReportBuilder)
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) PosMaintRptID() (*field.PosMaintRptID, error) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) PosReqType() (*field.PosReqType, error) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, error) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) SettlSessID() (*field.SettlSessID, error) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) NoPositions() (*field.NoPositions, error) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) SettlPrice() (*field.SettlPrice, error) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) PriorSettlPrice() (*field.PriorSettlPrice, error) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m *AdjustedPositionReport) PosMaintRptRefID() (*field.PosMaintRptRefID, error) {
	f := new(field.PosMaintRptRefID)
	err := m.Body.Get(f)
	return f, err
}
