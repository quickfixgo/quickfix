package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
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
	streamasgnreqid field.StreamAsgnReqID,
	streamasgnreqtype field.StreamAsgnReqType) StreamAssignmentRequestBuilder {
	var builder StreamAssignmentRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("CC"))
	builder.Body.Set(streamasgnreqid)
	builder.Body.Set(streamasgnreqtype)
	return builder
}

//StreamAsgnReqID is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqID() (*field.StreamAsgnReqID, errors.MessageRejectError) {
	f := new(field.StreamAsgnReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqID reads a StreamAsgnReqID from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetStreamAsgnReqID(f *field.StreamAsgnReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StreamAsgnReqType is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqType() (*field.StreamAsgnReqType, errors.MessageRejectError) {
	f := new(field.StreamAsgnReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetStreamAsgnReqType reads a StreamAsgnReqType from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetStreamAsgnReqType(f *field.StreamAsgnReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoAsgnReqs is a non-required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) NoAsgnReqs() (*field.NoAsgnReqs, errors.MessageRejectError) {
	f := new(field.NoAsgnReqs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoAsgnReqs reads a NoAsgnReqs from StreamAssignmentRequest.
func (m StreamAssignmentRequest) GetNoAsgnReqs(f *field.NoAsgnReqs) errors.MessageRejectError {
	return m.Body.Get(f)
}
