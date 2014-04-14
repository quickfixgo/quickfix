package fix50sp1

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
	liststatustype field.ListStatusType,
	norpts field.NoRpts,
	listorderstatus field.ListOrderStatus,
	rptseq field.RptSeq,
	totnoorders field.TotNoOrders,
	noorders field.NoOrders) *ListStatusBuilder {
	builder := new(ListStatusBuilder)
	builder.Body.Set(listid)
	builder.Body.Set(liststatustype)
	builder.Body.Set(norpts)
	builder.Body.Set(listorderstatus)
	builder.Body.Set(rptseq)
	builder.Body.Set(totnoorders)
	builder.Body.Set(noorders)
	return builder
}

//ListID is a required field for ListStatus.
func (m *ListStatus) ListID() (*field.ListID, error) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//ListStatusType is a required field for ListStatus.
func (m *ListStatus) ListStatusType() (*field.ListStatusType, error) {
	f := new(field.ListStatusType)
	err := m.Body.Get(f)
	return f, err
}

//NoRpts is a required field for ListStatus.
func (m *ListStatus) NoRpts() (*field.NoRpts, error) {
	f := new(field.NoRpts)
	err := m.Body.Get(f)
	return f, err
}

//ListOrderStatus is a required field for ListStatus.
func (m *ListStatus) ListOrderStatus() (*field.ListOrderStatus, error) {
	f := new(field.ListOrderStatus)
	err := m.Body.Get(f)
	return f, err
}

//RptSeq is a required field for ListStatus.
func (m *ListStatus) RptSeq() (*field.RptSeq, error) {
	f := new(field.RptSeq)
	err := m.Body.Get(f)
	return f, err
}

//ListStatusText is a non-required field for ListStatus.
func (m *ListStatus) ListStatusText() (*field.ListStatusText, error) {
	f := new(field.ListStatusText)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListStatusTextLen is a non-required field for ListStatus.
func (m *ListStatus) EncodedListStatusTextLen() (*field.EncodedListStatusTextLen, error) {
	f := new(field.EncodedListStatusTextLen)
	err := m.Body.Get(f)
	return f, err
}

//EncodedListStatusText is a non-required field for ListStatus.
func (m *ListStatus) EncodedListStatusText() (*field.EncodedListStatusText, error) {
	f := new(field.EncodedListStatusText)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a non-required field for ListStatus.
func (m *ListStatus) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//TotNoOrders is a required field for ListStatus.
func (m *ListStatus) TotNoOrders() (*field.TotNoOrders, error) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}

//LastFragment is a non-required field for ListStatus.
func (m *ListStatus) LastFragment() (*field.LastFragment, error) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//NoOrders is a required field for ListStatus.
func (m *ListStatus) NoOrders() (*field.NoOrders, error) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//ContingencyType is a non-required field for ListStatus.
func (m *ListStatus) ContingencyType() (*field.ContingencyType, error) {
	f := new(field.ContingencyType)
	err := m.Body.Get(f)
	return f, err
}

//ListRejectReason is a non-required field for ListStatus.
func (m *ListStatus) ListRejectReason() (*field.ListRejectReason, error) {
	f := new(field.ListRejectReason)
	err := m.Body.Get(f)
	return f, err
}
