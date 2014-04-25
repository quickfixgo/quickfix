package fix40

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
	norpts field.NoRpts,
	rptseq field.RptSeq,
	noorders field.NoOrders) ListStatusBuilder {
	var builder ListStatusBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("N"))
	builder.Body.Set(listid)
	builder.Body.Set(norpts)
	builder.Body.Set(rptseq)
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

//WaveNo is a non-required field for ListStatus.
func (m ListStatus) WaveNo() (*field.WaveNo, errors.MessageRejectError) {
	f := new(field.WaveNo)
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListStatus.
func (m ListStatus) GetWaveNo(f *field.WaveNo) errors.MessageRejectError {
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
