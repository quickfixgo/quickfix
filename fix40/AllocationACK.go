package fix40

import (
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
	builder.Body.Set(allocid)
	builder.Body.Set(tradedate)
	builder.Body.Set(allocstatus)
	return builder
}

//ClientID is a non-required field for AllocationACK.
func (m AllocationACK) ClientID() (field.ClientID, error) {
	var f field.ClientID
	err := m.Body.Get(&f)
	return f, err
}

//ExecBroker is a non-required field for AllocationACK.
func (m AllocationACK) ExecBroker() (field.ExecBroker, error) {
	var f field.ExecBroker
	err := m.Body.Get(&f)
	return f, err
}

//AllocID is a required field for AllocationACK.
func (m AllocationACK) AllocID() (field.AllocID, error) {
	var f field.AllocID
	err := m.Body.Get(&f)
	return f, err
}

//TradeDate is a required field for AllocationACK.
func (m AllocationACK) TradeDate() (field.TradeDate, error) {
	var f field.TradeDate
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a non-required field for AllocationACK.
func (m AllocationACK) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//AllocStatus is a required field for AllocationACK.
func (m AllocationACK) AllocStatus() (field.AllocStatus, error) {
	var f field.AllocStatus
	err := m.Body.Get(&f)
	return f, err
}

//AllocRejCode is a non-required field for AllocationACK.
func (m AllocationACK) AllocRejCode() (field.AllocRejCode, error) {
	var f field.AllocRejCode
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for AllocationACK.
func (m AllocationACK) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}
