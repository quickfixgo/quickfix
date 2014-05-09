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

//StreamAssignmentRequest msg type = CC.
type StreamAssignmentRequest struct {
	message.Message
}

//StreamAssignmentRequestBuilder builds StreamAssignmentRequest messages.
type StreamAssignmentRequestBuilder struct {
	message.MessageBuilder
}

//CreateStreamAssignmentRequestBuilder returns an initialized StreamAssignmentRequestBuilder with specified required fields.
func CreateStreamAssignmentRequestBuilder(
	streamasgnreqid *field.StreamAsgnReqIDField,
	streamasgnreqtype *field.StreamAsgnReqTypeField) StreamAssignmentRequestBuilder {
	var builder StreamAssignmentRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("CC"))
	builder.Body.Set(streamasgnreqid)
	builder.Body.Set(streamasgnreqtype)
	return builder
}

//StreamAsgnReqID is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqID() (*field.StreamAsgnReqIDField, errors.MessageRejectError) {
	f := &field.StreamAsgnReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetStreamAsgnReqID(f *field.StreamAsgnReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqType() (*field.StreamAsgnReqTypeField, errors.MessageRejectError) {
	f := &field.StreamAsgnReqTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetStreamAsgnReqType(f *field.StreamAsgnReqTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) NoAsgnReqs() (*field.NoAsgnReqsField, errors.MessageRejectError) {
	f := &field.NoAsgnReqsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetNoAsgnReqs(f *field.NoAsgnReqsField) errors.MessageRejectError {
	return m.Body.Get(f)
}
