package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//SecurityList msg type = y.
type SecurityList struct {
	message.Message
}

//SecurityListBuilder builds SecurityList messages.
type SecurityListBuilder struct {
	message.MessageBuilder
}

//CreateSecurityListBuilder returns an initialized SecurityListBuilder with specified required fields.
func CreateSecurityListBuilder(
	securityreqid *field.SecurityReqIDField,
	securityresponseid *field.SecurityResponseIDField,
	securityrequestresult *field.SecurityRequestResultField) SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX43))
	builder.Header.Set(field.NewMsgType("y"))
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityrequestresult)
	return builder
}

//SecurityReqID is a required field for SecurityList.
func (m SecurityList) SecurityReqID() (*field.SecurityReqIDField, errors.MessageRejectError) {
	f := &field.SecurityReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m SecurityList) GetSecurityReqID(f *field.SecurityReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityList.
func (m SecurityList) SecurityResponseID() (*field.SecurityResponseIDField, errors.MessageRejectError) {
	f := &field.SecurityResponseIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m SecurityList) GetSecurityResponseID(f *field.SecurityResponseIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (*field.SecurityRequestResultField, errors.MessageRejectError) {
	f := &field.SecurityRequestResultField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m SecurityList) GetSecurityRequestResult(f *field.SecurityRequestResultField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumSecurities is a non-required field for SecurityList.
func (m SecurityList) TotalNumSecurities() (*field.TotalNumSecuritiesField, errors.MessageRejectError) {
	f := &field.TotalNumSecuritiesField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumSecurities reads a TotalNumSecurities from SecurityList.
func (m SecurityList) GetTotalNumSecurities(f *field.TotalNumSecuritiesField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityList.
func (m SecurityList) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}
