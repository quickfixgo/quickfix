//Package streamassignmentreportack msg type = CE.
package streamassignmentreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a StreamAssignmentReportACK wrapper for the generic Message type
type Message struct {
	message.Message
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnAckType() (*field.StreamAsgnAckTypeField, errors.MessageRejectError) {
	f := &field.StreamAsgnAckTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnAckType reads a StreamAsgnAckType from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnAckType(f *field.StreamAsgnAckTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnRptID() (*field.StreamAsgnRptIDField, errors.MessageRejectError) {
	f := &field.StreamAsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnRptID(f *field.StreamAsgnRptIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnRejReason() (*field.StreamAsgnRejReasonField, errors.MessageRejectError) {
	f := &field.StreamAsgnRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRejReason reads a StreamAsgnRejReason from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnRejReason(f *field.StreamAsgnRejReasonField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from StreamAssignmentReportACK.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from StreamAssignmentReportACK.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from StreamAssignmentReportACK.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds StreamAssignmentReportACK messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for StreamAssignmentReportACK.
func Builder(
	streamasgnacktype *field.StreamAsgnAckTypeField,
	streamasgnrptid *field.StreamAsgnRptIDField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header().Set(field.NewMsgType("CE"))
	builder.Body().Set(streamasgnacktype)
	builder.Body().Set(streamasgnrptid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CE", r
}
