package fix50sp2

import (
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
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnRptID is a required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnRptID() (field.StreamAsgnRptID, error) {
	var f field.StreamAsgnRptID
	err := m.Body.Get(&f)
	return f, err
}

//StreamAsgnReqType is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqType() (field.StreamAsgnReqType, error) {
	var f field.StreamAsgnReqType
	err := m.Body.Get(&f)
	return f, err
}

//StreamAsgnReqID is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) StreamAsgnReqID() (field.StreamAsgnReqID, error) {
	var f field.StreamAsgnReqID
	err := m.Body.Get(&f)
	return f, err
}

//NoAsgnReqs is a non-required field for StreamAssignmentReport.
func (m StreamAssignmentReport) NoAsgnReqs() (field.NoAsgnReqs, error) {
	var f field.NoAsgnReqs
	err := m.Body.Get(&f)
	return f, err
}
