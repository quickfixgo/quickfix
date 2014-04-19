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
	builder.Body.Set(streamasgnreqid)
	builder.Body.Set(streamasgnreqtype)
	return builder
}

//StreamAsgnReqID is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqID() (field.StreamAsgnReqID, errors.MessageRejectError) {
	var f field.StreamAsgnReqID
	err := m.Body.Get(&f)
	return f, err
}

//StreamAsgnReqType is a required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) StreamAsgnReqType() (field.StreamAsgnReqType, errors.MessageRejectError) {
	var f field.StreamAsgnReqType
	err := m.Body.Get(&f)
	return f, err
}

//NoAsgnReqs is a non-required field for StreamAssignmentRequest.
func (m StreamAssignmentRequest) NoAsgnReqs() (field.NoAsgnReqs, errors.MessageRejectError) {
	var f field.NoAsgnReqs
	err := m.Body.Get(&f)
	return f, err
}
