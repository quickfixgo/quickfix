package fix50

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

//NewSettlementInstructionRequestBuilder returns an initialized SettlementInstructionRequestBuilder with specified required fields.
func NewSettlementInstructionRequestBuilder(
	settlinstreqid field.SettlInstReqID,
	transacttime field.TransactTime) *SettlementInstructionRequestBuilder {
	builder := new(SettlementInstructionRequestBuilder)
	builder.Body.Set(settlinstreqid)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstReqID is a required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) SettlInstReqID() (*field.SettlInstReqID, error) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) AllocAccount() (*field.AllocAccount, error) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, error) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//Side is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//Product is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) Product() (*field.Product, error) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) SecurityType() (*field.SecurityType, error) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) CFICode() (*field.CFICode, error) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) EffectiveTime() (*field.EffectiveTime, error) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) LastUpdateTime() (*field.LastUpdateTime, error) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) StandInstDbType() (*field.StandInstDbType, error) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) StandInstDbName() (*field.StandInstDbName, error) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) StandInstDbID() (*field.StandInstDbID, error) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//SettlCurrency is a non-required field for SettlementInstructionRequest.
func (m *SettlementInstructionRequest) SettlCurrency() (*field.SettlCurrency, error) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}
