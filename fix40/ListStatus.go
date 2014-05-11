package fix40

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
	norpts *field.NoRptsField,
	rptseq *field.RptSeqField,
	noorders *field.NoOrdersField) ListStatusBuilder {
	var builder ListStatusBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("N"))
	builder.Body().Set(listid)
	builder.Body().Set(norpts)
	builder.Body().Set(rptseq)
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

//WaveNo is a non-required field for ListStatus.
func (m ListStatus) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListStatus.
func (m ListStatus) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
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
