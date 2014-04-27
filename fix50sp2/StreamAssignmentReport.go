package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	streamasgnrptid field.StreamAsgnRptID) StreamAssignmentReportBuilder {
	var builder StreamAssignmentReportBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("CD"))
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnRptID is a required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnRptID() (*field.StreamAsgnRptID, errors.MessageRejectError) {
	f := new(field.StreamAsgnRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnRptID(f *field.StreamAsgnRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqType() (*field.StreamAsgnReqType, errors.MessageRejectError) {
	f := new(field.StreamAsgnReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnReqType(f *field.StreamAsgnReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqID is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqID() (*field.StreamAsgnReqID, errors.MessageRejectError) {
	f := new(field.StreamAsgnReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentReport.
func (m StreamAssignmentReport) GetStreamAsgnReqID(f *field.StreamAsgnReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) NoAsgnReqs() (*field.NoAsgnReqs, errors.MessageRejectError) {
	f := new(field.NoAsgnReqs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentReport.
func (m StreamAssignmentReport) GetNoAsgnReqs(f *field.NoAsgnReqs) errors.MessageRejectError {
	return m.Body.Get(f)
}
