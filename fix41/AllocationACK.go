package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//AllocationACK msg type = P.
type AllocationACK struct {
	message.Message
}

//AllocationACKBuilder builds AllocationACK messages.
type AllocationACKBuilder struct {
	message.MessageBuilder
}

//CreateAllocationACKBuilder returns an initialized AllocationACKBuilder with specified required fields.
func CreateAllocationACKBuilder(
	allocid field.AllocID,
	tradedate field.TradeDate,
	allocstatus field.AllocStatus) AllocationACKBuilder {
	var builder AllocationACKBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("P"))
	builder.Body.Set(allocid)
	builder.Body.Set(tradedate)
	builder.Body.Set(allocstatus)
	return builder
}

//ClientID is a non-required field for AllocationACK.
func (m AllocationACK) ClientID() (*field.ClientID, errors.MessageRejectError) {
	f := new(field.ClientID)
	err := m.Body.Get(f)
	return f, err
}

//GetClientID reads a ClientID from AllocationACK.
func (m AllocationACK) GetClientID(f *field.ClientID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExecBroker is a non-required field for AllocationACK.
func (m AllocationACK) ExecBroker() (*field.ExecBroker, errors.MessageRejectError) {
	f := new(field.ExecBroker)
	err := m.Body.Get(f)
	return f, err
}

//GetExecBroker reads a ExecBroker from AllocationACK.
func (m AllocationACK) GetExecBroker(f *field.ExecBroker) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocID is a required field for AllocationACK.
func (m AllocationACK) AllocID() (*field.AllocID, errors.MessageRejectError) {
	f := new(field.AllocID)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocID reads a AllocID from AllocationACK.
func (m AllocationACK) GetAllocID(f *field.AllocID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TradeDate is a required field for AllocationACK.
func (m AllocationACK) TradeDate() (*field.TradeDate, errors.MessageRejectError) {
	f := new(field.TradeDate)
	err := m.Body.Get(f)
	return f, err
}

//GetTradeDate reads a TradeDate from AllocationACK.
func (m AllocationACK) GetTradeDate(f *field.TradeDate) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a non-required field for AllocationACK.
func (m AllocationACK) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from AllocationACK.
func (m AllocationACK) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocStatus is a required field for AllocationACK.
func (m AllocationACK) AllocStatus() (*field.AllocStatus, errors.MessageRejectError) {
	f := new(field.AllocStatus)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocStatus reads a AllocStatus from AllocationACK.
func (m AllocationACK) GetAllocStatus(f *field.AllocStatus) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocRejCode is a non-required field for AllocationACK.
func (m AllocationACK) AllocRejCode() (*field.AllocRejCode, errors.MessageRejectError) {
	f := new(field.AllocRejCode)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocRejCode reads a AllocRejCode from AllocationACK.
func (m AllocationACK) GetAllocRejCode(f *field.AllocRejCode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for AllocationACK.
func (m AllocationACK) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from AllocationACK.
func (m AllocationACK) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}
