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
	streamasgnacktype *field.StreamAsgnAckTypeField,
	streamasgnrptid *field.StreamAsgnRptIDField) StreamAssignmentReportACKBuilder {
	var builder StreamAssignmentReportACKBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CE"))
	builder.Body.Set(streamasgnacktype)
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnAckType() (*field.StreamAsgnAckTypeField, errors.MessageRejectError) {
	f := &field.StreamAsgnAckTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnAckType reads a StreamAsgnAckType from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnAckType(f *field.StreamAsgnAckTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRptID() (*field.StreamAsgnRptIDField, errors.MessageRejectError) {
	f := &field.StreamAsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnRptID(f *field.StreamAsgnRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRejReason() (*field.StreamAsgnRejReasonField, errors.MessageRejectError) {
	f := &field.StreamAsgnRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRejReason reads a StreamAsgnRejReason from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnRejReason(f *field.StreamAsgnRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}
