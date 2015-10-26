//Package streamassignmentreportack msg type = CE.
package streamassignmentreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a StreamAssignmentReportACK wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnAckType() (*field.StreamAsgnAckTypeField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnAckTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnAckType reads a StreamAsgnAckType from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnAckType(f *field.StreamAsgnAckTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnRptID() (*field.StreamAsgnRptIDField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnRptIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRptID reads a StreamAsgnRptID from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnRptID(f *field.StreamAsgnRptIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
func (m Message) StreamAsgnRejReason() (*field.StreamAsgnRejReasonField, quickfix.MessageRejectError) {
	f := &field.StreamAsgnRejReasonField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnRejReason reads a StreamAsgnRejReason from StreamAssignmentReportACK.
func (m Message) GetStreamAsgnRejReason(f *field.StreamAsgnRejReasonField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for StreamAssignmentReportACK.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from StreamAssignmentReportACK.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from StreamAssignmentReportACK.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for StreamAssignmentReportACK.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from StreamAssignmentReportACK.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for StreamAssignmentReportACK.
func New(
	streamasgnacktype *field.StreamAsgnAckTypeField,
	streamasgnrptid *field.StreamAsgnRptIDField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CE"))
	builder.Body.Set(streamasgnacktype)
	builder.Body.Set(streamasgnrptid)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CE", r
}
