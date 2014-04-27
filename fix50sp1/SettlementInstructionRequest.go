package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("AV"))
	builder.Body.Set(settlinstreqid)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstReqID is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlInstReqID() (*field.SettlInstReqID, errors.MessageRejectError) {
	f := new(field.SettlInstReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSettlInstReqID(f *field.SettlInstReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) TransactTime() (*field.TransactTime, errors.MessageRejectError) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetTransactTime(f *field.TransactTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAccount() (*field.AllocAccount, errors.MessageRejectError) {
	f := new(field.AllocAccount)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetAllocAccount(f *field.AllocAccount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAcctIDSource() (*field.AllocAcctIDSource, errors.MessageRejectError) {
	f := new(field.AllocAcctIDSource)
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetAllocAcctIDSource(f *field.AllocAcctIDSource) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Side() (*field.Side, errors.MessageRejectError) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSide(f *field.Side) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Product() (*field.Product, errors.MessageRejectError) {
	f := new(field.Product)
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetProduct(f *field.Product) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SecurityType() (*field.SecurityType, errors.MessageRejectError) {
	f := new(field.SecurityType)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSecurityType(f *field.SecurityType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) CFICode() (*field.CFICode, errors.MessageRejectError) {
	f := new(field.CFICode)
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetCFICode(f *field.CFICode) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) EffectiveTime() (*field.EffectiveTime, errors.MessageRejectError) {
	f := new(field.EffectiveTime)
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetEffectiveTime(f *field.EffectiveTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) ExpireTime() (*field.ExpireTime, errors.MessageRejectError) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetExpireTime(f *field.ExpireTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) LastUpdateTime() (*field.LastUpdateTime, errors.MessageRejectError) {
	f := new(field.LastUpdateTime)
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetLastUpdateTime(f *field.LastUpdateTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbType() (*field.StandInstDbType, errors.MessageRejectError) {
	f := new(field.StandInstDbType)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbType(f *field.StandInstDbType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbName() (*field.StandInstDbName, errors.MessageRejectError) {
	f := new(field.StandInstDbName)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbName(f *field.StandInstDbName) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbID() (*field.StandInstDbID, errors.MessageRejectError) {
	f := new(field.StandInstDbID)
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbID(f *field.StandInstDbID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlCurrency() (*field.SettlCurrency, errors.MessageRejectError) {
	f := new(field.SettlCurrency)
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSettlCurrency(f *field.SettlCurrency) errors.MessageRejectError {
	return m.Body.Get(f)
}
