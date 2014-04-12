package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type StreamAssignmentRequest struct {
	message.Message
}

func (m *StreamAssignmentRequest) StreamAsgnReqID() (*field.StreamAsgnReqID, error) {
	f := new(field.StreamAsgnReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentRequest) StreamAsgnReqType() (*field.StreamAsgnReqType, error) {
	f := new(field.StreamAsgnReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentRequest) NoAsgnReqs() (*field.NoAsgnReqs, error) {
	f := new(field.NoAsgnReqs)
	err := m.Body.Get(f)
	return f, err
}
