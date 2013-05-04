package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type StreamAssignmentReport struct {
	quickfixgo.Message
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
