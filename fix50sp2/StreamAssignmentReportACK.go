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
	streamasgnacktype field.StreamAsgnAckType,
	streamasgnrptid field.StreamAsgnRptID) StreamAssignmentReportACKBuilder {
	var builder StreamAssignmentReportACKBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.BuildMsgType("CE"))
	builder.Body.Set(streamasgnacktype)
	builder.Body.Set(streamasgnrptid)
	return builder
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnAckType() (*field.StreamAsgnAckType, errors.MessageRejectError) {
	f := new(field.StreamAsgnAckType)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnAckType reads a StreamAsgnAckType from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnAckType(f *field.StreamAsgnAckType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRptID() (*field.StreamAsgnRptID, errors.MessageRejectError) {
	f := new(field.StreamAsgnRptID)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnRptID(f *field.StreamAsgnRptID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) StreamAsgnRejReason() (*field.StreamAsgnRejReason, errors.MessageRejectError) {
	f := new(field.StreamAsgnRejReason)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRejReason reads a StreamAsgnRejReason from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetStreamAsgnRejReason(f *field.StreamAsgnRejReason) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from StreamAssignmentReportACK.
func (m StreamAssignmentReportACK) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}
