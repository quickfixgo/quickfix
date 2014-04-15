package fix50sp1

import (
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
func (m SettlementInstructionRequest) SettlInstReqID() (field.SettlInstReqID, error) {
	var f field.SettlInstReqID
	err := m.Body.Get(&f)
	return f, err
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) TransactTime() (field.TransactTime, error) {
	var f field.TransactTime
	err := m.Body.Get(&f)
	return f, err
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) NoPartyIDs() (field.NoPartyIDs, error) {
	var f field.NoPartyIDs
	err := m.Body.Get(&f)
	return f, err
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAccount() (field.AllocAccount, error) {
	var f field.AllocAccount
	err := m.Body.Get(&f)
	return f, err
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAcctIDSource() (field.AllocAcctIDSource, error) {
	var f field.AllocAcctIDSource
	err := m.Body.Get(&f)
	return f, err
}

//Side is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Side() (field.Side, error) {
	var f field.Side
	err := m.Body.Get(&f)
	return f, err
}

//Product is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Product() (field.Product, error) {
	var f field.Product
	err := m.Body.Get(&f)
	return f, err
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SecurityType() (field.SecurityType, error) {
	var f field.SecurityType
	err := m.Body.Get(&f)
	return f, err
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) CFICode() (field.CFICode, error) {
	var f field.CFICode
	err := m.Body.Get(&f)
	return f, err
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) EffectiveTime() (field.EffectiveTime, error) {
	var f field.EffectiveTime
	err := m.Body.Get(&f)
	return f, err
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) ExpireTime() (field.ExpireTime, error) {
	var f field.ExpireTime
	err := m.Body.Get(&f)
	return f, err
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) LastUpdateTime() (field.LastUpdateTime, error) {
	var f field.LastUpdateTime
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbType() (field.StandInstDbType, error) {
	var f field.StandInstDbType
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbName() (field.StandInstDbName, error) {
	var f field.StandInstDbName
	err := m.Body.Get(&f)
	return f, err
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbID() (field.StandInstDbID, error) {
	var f field.StandInstDbID
	err := m.Body.Get(&f)
	return f, err
}

//SettlCurrency is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlCurrency() (field.SettlCurrency, error) {
	var f field.SettlCurrency
	err := m.Body.Get(&f)
	return f, err
}
