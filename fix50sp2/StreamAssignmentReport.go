package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type StreamAssignmentReport struct {
	quickfix.Message
}

func (m *StreamAssignmentReport) StreamAsgnRptID() (*field.StreamAsgnRptID, error) {
	f := new(field.StreamAsgnRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReport) StreamAsgnReqType() (*field.StreamAsgnReqType, error) {
	f := new(field.StreamAsgnReqType)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReport) StreamAsgnReqID() (*field.StreamAsgnReqID, error) {
	f := new(field.StreamAsgnReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReport) NoAsgnReqs() (*field.NoAsgnReqs, error) {
	f := new(field.NoAsgnReqs)
	err := m.Body.Get(f)
	return f, err
}
