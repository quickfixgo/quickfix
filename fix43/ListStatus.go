package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	listid *field.ListIDField,
	liststatustype *field.ListStatusTypeField,
	norpts *field.NoRptsField,
	listorderstatus *field.ListOrderStatusField,
	rptseq *field.RptSeqField,
	totnoorders *field.TotNoOrdersField,
	noorders *field.NoOrdersField) ListStatusBuilder {
	var builder ListStatusBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header().Set(field.NewMsgType("N"))
	builder.Body().Set(listid)
	builder.Body().Set(liststatustype)
	builder.Body().Set(norpts)
	builder.Body().Set(listorderstatus)
	builder.Body().Set(rptseq)
	builder.Body().Set(totnoorders)
	builder.Body().Set(noorders)
	return builder
}

//ListID is a required field for ListStatus.
func (m ListStatus) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatus.
func (m ListStatus) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusType is a required field for ListStatus.
func (m ListStatus) ListStatusType() (*field.ListStatusTypeField, errors.MessageRejectError) {
	f := &field.ListStatusTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusType reads a ListStatusType from ListStatus.
func (m ListStatus) GetListStatusType(f *field.ListStatusTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRpts is a required field for ListStatus.
func (m ListStatus) NoRpts() (*field.NoRptsField, errors.MessageRejectError) {
	f := &field.NoRptsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRpts reads a NoRpts from ListStatus.
func (m ListStatus) GetNoRpts(f *field.NoRptsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListOrderStatus is a required field for ListStatus.
func (m ListStatus) ListOrderStatus() (*field.ListOrderStatusField, errors.MessageRejectError) {
	f := &field.ListOrderStatusField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListOrderStatus reads a ListOrderStatus from ListStatus.
func (m ListStatus) GetListOrderStatus(f *field.ListOrderStatusField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RptSeq is a required field for ListStatus.
func (m ListStatus) RptSeq() (*field.RptSeqField, errors.MessageRejectError) {
	f := &field.RptSeqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRptSeq reads a RptSeq from ListStatus.
func (m ListStatus) GetRptSeq(f *field.RptSeqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ListStatusText is a non-required field for ListStatus.
func (m ListStatus) ListStatusText() (*field.ListStatusTextField, errors.MessageRejectError) {
	f := &field.ListStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListStatusText reads a ListStatusText from ListStatus.
func (m ListStatus) GetListStatusText(f *field.ListStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusTextLen is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusTextLen() (*field.EncodedListStatusTextLenField, errors.MessageRejectError) {
	f := &field.EncodedListStatusTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusTextLen reads a EncodedListStatusTextLen from ListStatus.
func (m ListStatus) GetEncodedListStatusTextLen(f *field.EncodedListStatusTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedListStatusText is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusText() (*field.EncodedListStatusTextField, errors.MessageRejectError) {
	f := &field.EncodedListStatusTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedListStatusText reads a EncodedListStatusText from ListStatus.
func (m ListStatus) GetEncodedListStatusText(f *field.EncodedListStatusTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for ListStatus.
func (m ListStatus) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from ListStatus.
func (m ListStatus) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoOrders is a required field for ListStatus.
func (m ListStatus) TotNoOrders() (*field.TotNoOrdersField, errors.MessageRejectError) {
	f := &field.TotNoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoOrders reads a TotNoOrders from ListStatus.
func (m ListStatus) GetTotNoOrders(f *field.TotNoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for ListStatus.
func (m ListStatus) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ListStatus.
func (m ListStatus) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}
