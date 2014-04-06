package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type StreamAssignmentReportACK struct {
	quickfix.Message
}

func (m *StreamAssignmentReportACK) StreamAsgnAckType() (*field.StreamAsgnAckType, error) {
	f := new(field.StreamAsgnAckType)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReportACK) StreamAsgnRptID() (*field.StreamAsgnRptID, error) {
	f := new(field.StreamAsgnRptID)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReportACK) StreamAsgnRejReason() (*field.StreamAsgnRejReason, error) {
	f := new(field.StreamAsgnRejReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReportACK) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReportACK) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *StreamAssignmentReportACK) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
