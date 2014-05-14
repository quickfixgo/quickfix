//Package settlementinstructionrequest msg type = AV.
package settlementinstructionrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a SettlementInstructionRequest wrapper for the generic Message type
type Message struct {
	message.Message
}

//SettlInstReqID is a required field for SettlementInstructionRequest.
func (m Message) SettlInstReqID() (*field.SettlInstReqIDField, errors.MessageRejectError) {
	f := &field.SettlInstReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructionRequest.
func (m Message) GetSettlInstReqID(f *field.SettlInstReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructionRequest.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructionRequest.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for SettlementInstructionRequest.
func (m Message) NoPartyIDs() (*field.NoPartyIDsField, errors.MessageRejectError) {
	f := &field.NoPartyIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from SettlementInstructionRequest.
func (m Message) GetNoPartyIDs(f *field.NoPartyIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAccount is a non-required field for SettlementInstructionRequest.
func (m Message) AllocAccount() (*field.AllocAccountField, errors.MessageRejectError) {
	f := &field.AllocAccountField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAccount reads a AllocAccount from SettlementInstructionRequest.
func (m Message) GetAllocAccount(f *field.AllocAccountField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//AllocAcctIDSource is a non-required field for SettlementInstructionRequest.
func (m Message) AllocAcctIDSource() (*field.AllocAcctIDSourceField, errors.MessageRejectError) {
	f := &field.AllocAcctIDSourceField{}
	err := m.Body.Get(f)
	return f, err
}

//GetAllocAcctIDSource reads a AllocAcctIDSource from SettlementInstructionRequest.
func (m Message) GetAllocAcctIDSource(f *field.AllocAcctIDSourceField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Side is a non-required field for SettlementInstructionRequest.
func (m Message) Side() (*field.SideField, errors.MessageRejectError) {
	f := &field.SideField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSide reads a Side from SettlementInstructionRequest.
func (m Message) GetSide(f *field.SideField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Product is a non-required field for SettlementInstructionRequest.
func (m Message) Product() (*field.ProductField, errors.MessageRejectError) {
	f := &field.ProductField{}
	err := m.Body.Get(f)
	return f, err
}

//GetProduct reads a Product from SettlementInstructionRequest.
func (m Message) GetProduct(f *field.ProductField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityType is a non-required field for SettlementInstructionRequest.
func (m Message) SecurityType() (*field.SecurityTypeField, errors.MessageRejectError) {
	f := &field.SecurityTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityType reads a SecurityType from SettlementInstructionRequest.
func (m Message) GetSecurityType(f *field.SecurityTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//CFICode is a non-required field for SettlementInstructionRequest.
func (m Message) CFICode() (*field.CFICodeField, errors.MessageRejectError) {
	f := &field.CFICodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetCFICode reads a CFICode from SettlementInstructionRequest.
func (m Message) GetCFICode(f *field.CFICodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EffectiveTime is a non-required field for SettlementInstructionRequest.
func (m Message) EffectiveTime() (*field.EffectiveTimeField, errors.MessageRejectError) {
	f := &field.EffectiveTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEffectiveTime reads a EffectiveTime from SettlementInstructionRequest.
func (m Message) GetEffectiveTime(f *field.EffectiveTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ExpireTime is a non-required field for SettlementInstructionRequest.
func (m Message) ExpireTime() (*field.ExpireTimeField, errors.MessageRejectError) {
	f := &field.ExpireTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetExpireTime reads a ExpireTime from SettlementInstructionRequest.
func (m Message) GetExpireTime(f *field.ExpireTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastUpdateTime is a non-required field for SettlementInstructionRequest.
func (m Message) LastUpdateTime() (*field.LastUpdateTimeField, errors.MessageRejectError) {
	f := &field.LastUpdateTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLastUpdateTime reads a LastUpdateTime from SettlementInstructionRequest.
func (m Message) GetLastUpdateTime(f *field.LastUpdateTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbType is a non-required field for SettlementInstructionRequest.
func (m Message) StandInstDbType() (*field.StandInstDbTypeField, errors.MessageRejectError) {
	f := &field.StandInstDbTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbType reads a StandInstDbType from SettlementInstructionRequest.
func (m Message) GetStandInstDbType(f *field.StandInstDbTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbName is a non-required field for SettlementInstructionRequest.
func (m Message) StandInstDbName() (*field.StandInstDbNameField, errors.MessageRejectError) {
	f := &field.StandInstDbNameField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbName reads a StandInstDbName from SettlementInstructionRequest.
func (m Message) GetStandInstDbName(f *field.StandInstDbNameField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//StandInstDbID is a non-required field for SettlementInstructionRequest.
func (m Message) StandInstDbID() (*field.StandInstDbIDField, errors.MessageRejectError) {
	f := &field.StandInstDbIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetStandInstDbID reads a StandInstDbID from SettlementInstructionRequest.
func (m Message) GetStandInstDbID(f *field.StandInstDbIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlCurrency is a non-required field for SettlementInstructionRequest.
func (m Message) SettlCurrency() (*field.SettlCurrencyField, errors.MessageRejectError) {
	f := &field.SettlCurrencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlCurrency reads a SettlCurrency from SettlementInstructionRequest.
func (m Message) GetSettlCurrency(f *field.SettlCurrencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SettlementInstructionRequest messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SettlementInstructionRequest.
func Builder(
	settlinstreqid *field.SettlInstReqIDField,
	transacttime *field.TransactTimeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header().Set(field.NewMsgType("AV"))
	builder.Body().Set(settlinstreqid)
	builder.Body().Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX50, "AV", r
}
