package fix40

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type Reject struct {
	quickfixgo.Message
}

func (m *Reject) RefSeqNum() (*field.RefSeqNum, error) {
	f := new(field.RefSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *Reject) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
