package fix44

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
	settlinstreqid *field.SettlInstReqIDField,
	transacttime *field.TransactTimeField) SettlementInstructionRequestBuilder {
	var builder SettlementInstructionRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header.Set(field.NewMsgType("AV"))
	builder.Body.Set(settlinstreqid)
	builder.Body.Set(transacttime)
	return builder
}

//SettlInstReqID is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SettlInstReqID() (*field.SettlInstReqIDField, errors.MessageRejectError) {
	f := &field.SettlInstReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSettlInstReqID(f *field.SettlInstReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAccount() (*field.AllocAccountField, errors.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetAllocAccount(f *field.AllocAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) AllocAcctIDSource() (*field.AllocAcctIDSourceField, errors.MessageRejectError) {
	f := &field.AllocAcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetAllocAcctIDSource(f *field.AllocAcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) LastUpdateTime() (*field.LastUpdateTimeField, errors.MessageRejectError) {
	f := &field.LastUpdateTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetLastUpdateTime(f *field.LastUpdateTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbType() (*field.StandInstDbTypeField, errors.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbType(f *field.StandInstDbTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbName() (*field.StandInstDbNameField, errors.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbName(f *field.StandInstDbNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m SettlementInstructionRequest) StandInstDbID() (*field.StandInstDbIDField, errors.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructionRequest.
func (m SettlementInstructionRequest) GetStandInstDbID(f *field.StandInstDbIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
