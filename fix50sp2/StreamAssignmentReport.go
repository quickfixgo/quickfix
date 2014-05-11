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

//StreamAssignmentReport msg type = CD.
type StreamAssignmentReport struct {
	message.Message
}

//StreamAssignmentReportBuilder builds StreamAssignmentReport messages.
type StreamAssignmentReportBuilder struct {
	message.MessageBuilder
}

//CreateStreamAssignmentReportBuilder returns an initialized StreamAssignmentReportBuilder with specified required fields.
func CreateStreamAssignmentReportBuilder(
	streamasgnrptid *field.StreamAsgnRptIDField) StreamAssignmentReportBuilder {
	var builder StreamAssignmentReportBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CD"))
	builder.Body().Set(streamasgnrptid)
	return builder
}

//StreamAsgnRptID is a required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnRptID() (*field.StreamAsgnRptIDField, errors.MessageRejectError) {
	f := &field.StreamAsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnRptID(f *field.StreamAsgnRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqType() (*field.StreamAsgnReqTypeField, errors.MessageRejectError) {
	f := &field.StreamAsgnReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnReqType(f *field.StreamAsgnReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqID is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqID() (*field.StreamAsgnReqIDField, errors.MessageRejectError) {
	f := &field.StreamAsgnReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnReqID(f *field.StreamAsgnReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) NoAsgnReqs() (*field.NoAsgnReqsField, errors.MessageRejectError) {
	f := &field.NoAsgnReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentReport.
func (m StreamAssignmentReport) GetNoAsgnReqs(f *field.NoAsgnReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
