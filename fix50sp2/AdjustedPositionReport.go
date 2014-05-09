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
	posmaintrptid *field.PosMaintRptIDField,
	clearingbusinessdate *field.ClearingBusinessDateField) AdjustedPositionReportBuilder {
	var builder AdjustedPositionReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("BL"))
	builder.Body.Set(posmaintrptid)
	builder.Body.Set(clearingbusinessdate)
	return builder
}

//PosMaintRptID is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptID() (*field.PosMaintRptIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptID reads a PosMaintRptID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosMaintRptID(f *field.PosMaintRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosReqType is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosReqType() (*field.PosReqTypeField, errors.MessageRejectError) {
	f := &field.PosReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosReqType reads a PosReqType from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosReqType(f *field.PosReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a required field for AdjustedPositionReport.
func (m AdjustedPositionReport) ClearingBusinessDate() (*field.ClearingBusinessDateField, errors.MessageRejectError) {
	f := &field.ClearingBusinessDateField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from AdjustedPositionReport.
func (m AdjustedPositionReport) GetClearingBusinessDate(f *field.ClearingBusinessDateField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlSessID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlSessID() (*field.SettlSessIDField, errors.MessageRejectError) {
	f := &field.SettlSessIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlSessID reads a SettlSessID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetSettlSessID(f *field.SettlSessIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPositions is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoPositions() (*field.NoPositionsField, errors.MessageRejectError) {
	f := &field.NoPositionsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPositions reads a NoPositions from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoPositions(f *field.NoPositionsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from AdjustedPositionReport.
func (m AdjustedPositionReport) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) SettlPrice() (*field.SettlPriceField, errors.MessageRejectError) {
	f := &field.SettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlPrice reads a SettlPrice from AdjustedPositionReport.
func (m AdjustedPositionReport) GetSettlPrice(f *field.SettlPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PriorSettlPrice is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PriorSettlPrice() (*field.PriorSettlPriceField, errors.MessageRejectError) {
	f := &field.PriorSettlPriceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPriorSettlPrice reads a PriorSettlPrice from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPriorSettlPrice(f *field.PriorSettlPriceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//PosMaintRptRefID is a non-required field for AdjustedPositionReport.
func (m AdjustedPositionReport) PosMaintRptRefID() (*field.PosMaintRptRefIDField, errors.MessageRejectError) {
	f := &field.PosMaintRptRefIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetPosMaintRptRefID reads a PosMaintRptRefID from AdjustedPositionReport.
func (m AdjustedPositionReport) GetPosMaintRptRefID(f *field.PosMaintRptRefIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
