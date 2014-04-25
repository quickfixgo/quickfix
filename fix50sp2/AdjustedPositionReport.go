package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Header.Set(field.BuildMsgType("BL"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptID() (*field.PosMaintRptID, errors.MessageRejectError) {
	f := new(field.PosMaintRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosMaintRptID(f *field.PosMaintRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosReqType() (*field.PosReqType, errors.MessageRejectError) {
	f := new(field.PosReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosReqType(f *field.PosReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AdjustedPositionReport.
func (m AdjustedPositionReport) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlSessID() (*field.SettlSessID, errors.MessageRejectError) {
	f := new(field.SettlSessID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetSettlSessID(f *field.SettlSessID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPositions() (*field.NoPositions, errors.MessageRejectError) {
	f := new(field.NoPositions)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoPositions(f *field.NoPositions) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlPrice() (*field.SettlPrice, errors.MessageRejectError) {
	f := new(field.SettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AdjustedPositionReport.
func (m AdjustedPositionReport) GetSettlPrice(f *field.SettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PriorSettlPrice() (*field.PriorSettlPrice, errors.MessageRejectError) {
	f := new(field.PriorSettlPrice)
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPriorSettlPrice(f *field.PriorSettlPrice) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptRefID() (*field.PosMaintRptRefID, errors.MessageRejectError) {
	f := new(field.PosMaintRptRefID)
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptRefID reads a PosMaintRptRefID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosMaintRptRefID(f *field.PosMaintRptRefID) errors.MessageRejectError {
	return m.Body.Get(f)
}
