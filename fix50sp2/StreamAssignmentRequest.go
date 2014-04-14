package fix50sp2

import (
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

//NewStreamAssignmentRequestBuilder returns an initialized StreamAssignmentRequestBuilder with specified required fields.
func NewStreamAssignmentRequestBuilder(
	streamasgnreqid field.StreamAsgnReqID,
	streamasgnreqtype field.StreamAsgnReqType) *StreamAssignmentRequestBuilder {
	builder := new(StreamAssignmentRequestBuilder)
	builder.Body.Set(streamasgnreqid)
	builder.Body.Set(streamasgnreqtype)
	return builder
}

//StreamAsgnReqID is a required field for StreamAssignmentRequest.
func (m *StreamAssignmentRequest) StreamAsgnReqID() (*field.StreamAsgnReqID, error) {
	f := new(field.StreamAsgnReqID)
	err := m.Body.Get(f)
	return f, err
}

//StreamAsgnReqType is a required field for StreamAssignmentRequest.
func (m *StreamAssignmentRequest) StreamAsgnReqType() (*field.StreamAsgnReqType, error) {
	f := new(field.StreamAsgnReqType)
	err := m.Body.Get(f)
	return f, err
}

//NoAsgnReqs is a non-required field for StreamAssignmentRequest.
func (m *StreamAssignmentRequest) NoAsgnReqs() (*field.NoAsgnReqs, error) {
	f := new(field.NoAsgnReqs)
	err := m.Body.Get(f)
	return f, err
}
