package fix50sp1

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

type ListStatus struct {
	message.Message
}

func (m *ListStatus) ListStatusType() (*field.ListStatusType, error) {
	f := new(field.ListStatusType)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) NoRpts() (*field.NoRpts, error) {
	f := new(field.NoRpts)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) ListStatusText() (*field.ListStatusText, error) {
	f := new(field.ListStatusText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) RptSeq() (*field.RptSeq, error) {
	f := new(field.RptSeq)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) ListOrderStatus() (*field.ListOrderStatus, error) {
	f := new(field.ListOrderStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) EncodedListStatusTextLen() (*field.EncodedListStatusTextLen, error) {
	f := new(field.EncodedListStatusTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) EncodedListStatusText() (*field.EncodedListStatusText, error) {
	f := new(field.EncodedListStatusText)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) ListRejectReason() (*field.ListRejectReason, error) {
	f := new(field.ListRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}
func (m *ListStatus) ContingencyType() (*field.ContingencyType, error) {
	f := new(field.ContingencyType)
	err := m.Body.Get(f)
	return f, err
}
