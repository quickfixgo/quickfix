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

//CreateAdjustedPositionReportBuilder returns an initialized AdjustedPositionReportBuilder with specified required fields.
func CreateAdjustedPositionReportBuilder(
	posmaintrptid field.PosMaintRptID,
	clearingbusinessdate field.ClearingBusinessDate) AdjustedPositionReportBuilder {
	var builder AdjustedPositionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptID() (field.PosMaintRptID, error) {
	var f field.PosMaintRptID
	err := m.Body.Get(&f)
	return f, err
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosReqType() (field.PosReqType, error) {
	var f field.PosReqType
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlSessID() (field.SettlSessID, error) {
	var f field.SettlSessID
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPositions() (field.NoPositions, error) {
	var f field.NoPositions
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlPrice() (field.SettlPrice, error) {
	var f field.SettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PriorSettlPrice() (field.PriorSettlPrice, error) {
	var f field.PriorSettlPrice
	err := m.Body.Get(&f)
	return f, err
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptRefID() (field.PosMaintRptRefID, error) {
	var f field.PosMaintRptRefID
	err := m.Body.Get(&f)
	return f, err
}
