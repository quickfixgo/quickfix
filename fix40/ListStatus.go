package fix40

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ListStatus msg type = N.
type ListStatus struct {
	message.Message
}

//ListStatusBuilder builds ListStatus messages.
type ListStatusBuilder struct {
	message.MessageBuilder
}

//NewListStatusBuilder returns an initialized ListStatusBuilder with specified required fields.
func NewListStatusBuilder(
	listid field.ListID,
	norpts field.NoRpts,
	rptseq field.RptSeq,
	noorders field.NoOrders) *ListStatusBuilder {
	builder := new(ListStatusBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(norpts)
	builder.Body.Set(rptseq)
	builder.Body.Set(noorders)
	return builder
}

//ListID is a required field for ListStatus.
func (m *ListStatus) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//WaveNo is a non-required field for ListStatus.
func (m *ListStatus) WaveNo() (*field.WaveNo, error) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//NoRpts is a required field for ListStatus.
func (m *ListStatus) NoRpts() (*field.NoRpts, error) {
	f := new(field.NoRpts)
	err := m.Body.Get(f)
	return f, err
}

//RptSeq is a required field for ListStatus.
func (m *ListStatus) RptSeq() (*field.RptSeq, error) {
	f := new(field.RptSeq)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a required field for ListStatus.
func (m *ListStatus) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
