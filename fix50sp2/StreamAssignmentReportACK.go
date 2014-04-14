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

//NewStreamAssignmentReportACKBuilder returns an initialized StreamAssignmentReportACKBuilder with specified required fields.
func NewStreamAssignmentReportACKBuilder(
	streamasgnacktype field.StreamAsgnAckType,
	streamasgnrptid field.StreamAsgnRptID) *StreamAssignmentReportACKBuilder {
	builder := new(StreamAssignmentReportACKBuilder)
	builder.Body.Set(streamasgnacktype)
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) StreamAsgnAckType() (*field.StreamAsgnAckType, error) {
	f := new(field.StreamAsgnAckType)
	err := m.Body.Get(f)
	return f, err
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) StreamAsgnRptID() (*field.StreamAsgnRptID, error) {
	f := new(field.StreamAsgnRptID)
	err := m.Body.Get(f)
	return f, err
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) StreamAsgnRejReason() (*field.StreamAsgnRejReason, error) {
	f := new(field.StreamAsgnRejReason)
	err := m.Body.Get(f)
	return f, err
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m *StreamAssignmentReportACK) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
