package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//StreamAssignmentReportACK msg type = CE.
type StreamAssignmentReportACK struct {
	message.Message
}

//StreamAssignmentReportACKBuilder builds StreamAssignmentReportACK messages.
type StreamAssignmentReportACKBuilder struct {
	message.MessageBuilder
}

//CreateStreamAssignmentReportACKBuilder returns an initialized StreamAssignmentReportACKBuilder with specified required fields.
func CreateStreamAssignmentReportACKBuilder(
	streamasgnacktype field.StreamAsgnAckType,
	streamasgnrptid field.StreamAsgnRptID) StreamAssignmentReportACKBuilder {
	var builder StreamAssignmentReportACKBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(streamasgnacktype)
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnAckType() (field.StreamAsgnAckType, error) {
	var f field.StreamAsgnAckType
	err := m.Body.Get(&f)
	return f, err
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRptID() (field.StreamAsgnRptID, error) {
	var f field.StreamAsgnRptID
	err := m.Body.Get(&f)
	return f, err
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRejReason() (field.StreamAsgnRejReason, error) {
	var f field.StreamAsgnRejReason
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedTextLen() (field.EncodedTextLen, error) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedText() (field.EncodedText, error) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
