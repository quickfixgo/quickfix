package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SettlementInstructionRequest msg type = AV.
type SettlementInstructionRequest struct {
	message.Message
}

//SettlementInstructionRequestBuilder builds SettlementInstructionRequest messages.
type SettlementInstructionRequestBuilder struct {
	message.MessageBuilder
}

//CreateSettlementInstructionRequestBuilder returns an initialized SettlementInstructionRequestBuilder with specified required fields.
func CreateSettlementInstructionRequestBuilder(
	settlinstreqid field.SettlInstReqID,
	transacttime field.TransactTime) SettlementInstructionRequestBuilder {
	var builder SettlementInstructionRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(settlinstreqid)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstReqID is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlInstReqID() (field.SettlInstReqID, errors.MessageRejectError) {
	var f field.SettlInstReqID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) TransactTime() (field.TransactTime, errors.MessageRejectError) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) NoPartyIDs() (field.NoPartyIDs, errors.MessageRejectError) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAccount() (field.AllocAccount, errors.MessageRejectError) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAcctIDSource() (field.AllocAcctIDSource, errors.MessageRejectError) {
	var f field.AllocAcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Side() (field.Side, errors.MessageRejectError) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Product() (field.Product, errors.MessageRejectError) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SecurityType() (field.SecurityType, errors.MessageRejectError) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) CFICode() (field.CFICode, errors.MessageRejectError) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) EffectiveTime() (field.EffectiveTime, errors.MessageRejectError) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) ExpireTime() (field.ExpireTime, errors.MessageRejectError) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) LastUpdateTime() (field.LastUpdateTime, errors.MessageRejectError) {
	var f field.LastUpdateTime
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbType() (field.StandInstDbType, errors.MessageRejectError) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbName() (field.StandInstDbName, errors.MessageRejectError) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbID() (field.StandInstDbID, errors.MessageRejectError) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlCurrency() (field.SettlCurrency, errors.MessageRejectError) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}
