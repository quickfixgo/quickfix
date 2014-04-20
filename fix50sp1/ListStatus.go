package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateListStatusBuilder returns an initialized ListStatusBuilder with specified required fields.
func CreateListStatusBuilder(
	listid field.ListID,
	liststatustype field.ListStatusType,
	norpts field.NoRpts,
	listorderstatus field.ListOrderStatus,
	rptseq field.RptSeq,
	totnoorders field.TotNoOrders,
	noorders field.NoOrders) ListStatusBuilder {
	var builder ListStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
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
func (m ListStatus) ListID() (*field.ListID, errors.MessageRejectError) {
	f := new(field.ListID)
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatus.
func (m ListStatus) GetListID(f *field.ListID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusType is a required field for ListStatus.
func (m ListStatus) ListStatusType() (*field.ListStatusType, errors.MessageRejectError) {
	f := new(field.ListStatusType)
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusType reads a ListStatusType from ListStatus.
func (m ListStatus) GetListStatusType(f *field.ListStatusType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRpts is a required field for ListStatus.
func (m ListStatus) NoRpts() (*field.NoRpts, errors.MessageRejectError) {
	f := new(field.NoRpts)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRpts reads a NoRpts from ListStatus.
func (m ListStatus) GetNoRpts(f *field.NoRpts) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListOrderStatus is a required field for ListStatus.
func (m ListStatus) ListOrderStatus() (*field.ListOrderStatus, errors.MessageRejectError) {
	f := new(field.ListOrderStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetListOrderStatus reads a ListOrderStatus from ListStatus.
func (m ListStatus) GetListOrderStatus(f *field.ListOrderStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RptSeq is a required field for ListStatus.
func (m ListStatus) RptSeq() (*field.RptSeq, errors.MessageRejectError) {
	f := new(field.RptSeq)
	err := m.Body.Get(f)
	return f, err
}

//GetRptSeq reads a RptSeq from ListStatus.
func (m ListStatus) GetRptSeq(f *field.RptSeq) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusText is a non-required field for ListStatus.
func (m ListStatus) ListStatusText() (*field.ListStatusText, errors.MessageRejectError) {
	f := new(field.ListStatusText)
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusText reads a ListStatusText from ListStatus.
func (m ListStatus) GetListStatusText(f *field.ListStatusText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusTextLen is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusTextLen() (*field.EncodedListStatusTextLen, errors.MessageRejectError) {
	f := new(field.EncodedListStatusTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusTextLen reads a EncodedListStatusTextLen from ListStatus.
func (m ListStatus) GetEncodedListStatusTextLen(f *field.EncodedListStatusTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusText is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusText() (*field.EncodedListStatusText, errors.MessageRejectError) {
	f := new(field.EncodedListStatusText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusText reads a EncodedListStatusText from ListStatus.
func (m ListStatus) GetEncodedListStatusText(f *field.EncodedListStatusText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ListStatus.
func (m ListStatus) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListStatus.
func (m ListStatus) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for ListStatus.
func (m ListStatus) TotNoOrders() (*field.TotNoOrders, errors.MessageRejectError) {
	f := new(field.TotNoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from ListStatus.
func (m ListStatus) GetTotNoOrders(f *field.TotNoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for ListStatus.
func (m ListStatus) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from ListStatus.
func (m ListStatus) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for ListStatus.
func (m ListStatus) NoOrders() (*field.NoOrders, errors.MessageRejectError) {
	f := new(field.NoOrders)
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ListStatus.
func (m ListStatus) GetNoOrders(f *field.NoOrders) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ContingencyType is a non-required field for ListStatus.
func (m ListStatus) ContingencyType() (*field.ContingencyType, errors.MessageRejectError) {
	f := new(field.ContingencyType)
	err := m.Body.Get(f)
	return f, err
}

//GetContingencyType reads a ContingencyType from ListStatus.
func (m ListStatus) GetContingencyType(f *field.ContingencyType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListRejectReason is a non-required field for ListStatus.
func (m ListStatus) ListRejectReason() (*field.ListRejectReason, errors.MessageRejectError) {
	f := new(field.ListRejectReason)
	err := m.Body.Get(f)
	return f, err
}

//GetListRejectReason reads a ListRejectReason from ListStatus.
func (m ListStatus) GetListRejectReason(f *field.ListRejectReason) errors.MessageRejectError {
	return m.Body.Get(f)
}
