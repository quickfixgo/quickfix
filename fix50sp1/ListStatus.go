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
func (m ListStatus) ListID() (field.ListID, errors.MessageRejectError) {
	var f field.ListID
	err := m.Body.Get(&f)
	return f, err
}

//ListStatusType is a required field for ListStatus.
func (m ListStatus) ListStatusType() (field.ListStatusType, errors.MessageRejectError) {
	var f field.ListStatusType
	err := m.Body.Get(&f)
	return f, err
}

//NoRpts is a required field for ListStatus.
func (m ListStatus) NoRpts() (field.NoRpts, errors.MessageRejectError) {
	var f field.NoRpts
	err := m.Body.Get(&f)
	return f, err
}

//ListOrderStatus is a required field for ListStatus.
func (m ListStatus) ListOrderStatus() (field.ListOrderStatus, errors.MessageRejectError) {
	var f field.ListOrderStatus
	err := m.Body.Get(&f)
	return f, err
}

//RptSeq is a required field for ListStatus.
func (m ListStatus) RptSeq() (field.RptSeq, errors.MessageRejectError) {
	var f field.RptSeq
	err := m.Body.Get(&f)
	return f, err
}

//ListStatusText is a non-required field for ListStatus.
func (m ListStatus) ListStatusText() (field.ListStatusText, errors.MessageRejectError) {
	var f field.ListStatusText
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListStatusTextLen is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusTextLen() (field.EncodedListStatusTextLen, errors.MessageRejectError) {
	var f field.EncodedListStatusTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedListStatusText is a non-required field for ListStatus.
func (m ListStatus) EncodedListStatusText() (field.EncodedListStatusText, errors.MessageRejectError) {
	var f field.EncodedListStatusText
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for ListStatus.
func (m ListStatus) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//TotNoOrders is a required field for ListStatus.
func (m ListStatus) TotNoOrders() (field.TotNoOrders, errors.MessageRejectError) {
	var f field.TotNoOrders
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for ListStatus.
func (m ListStatus) LastFragment() (field.LastFragment, errors.MessageRejectError) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoOrders is a required field for ListStatus.
func (m ListStatus) NoOrders() (field.NoOrders, errors.MessageRejectError) {
	var f field.NoOrders
	err := m.Body.Get(&f)
	return f, err
}

//ContingencyType is a non-required field for ListStatus.
func (m ListStatus) ContingencyType() (field.ContingencyType, errors.MessageRejectError) {
	var f field.ContingencyType
	err := m.Body.Get(&f)
	return f, err
}

//ListRejectReason is a non-required field for ListStatus.
func (m ListStatus) ListRejectReason() (field.ListRejectReason, errors.MessageRejectError) {
	var f field.ListRejectReason
	err := m.Body.Get(&f)
	return f, err
}
