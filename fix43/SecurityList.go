package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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
	securityreqid field.SecurityReqID,
	securityresponseid field.SecurityResponseID,
	securityrequestresult field.SecurityRequestResult) SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(securityreqid)
	builder.Body.Set(securityresponseid)
	builder.Body.Set(securityrequestresult)
	return builder
}

//SecurityReqID is a required field for SecurityList.
func (m SecurityList) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m SecurityList) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a required field for SecurityList.
func (m SecurityList) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m SecurityList) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (*field.SecurityRequestResult, errors.MessageRejectError) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m SecurityList) GetSecurityRequestResult(f *field.SecurityRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotalNumSecurities is a non-required field for SecurityList.
func (m SecurityList) TotalNumSecurities() (*field.TotalNumSecurities, errors.MessageRejectError) {
	f := new(field.TotalNumSecurities)
	err := m.Body.Get(f)
	return f, err
}

//GetTotalNumSecurities reads a TotalNumSecurities from SecurityList.
func (m SecurityList) GetTotalNumSecurities(f *field.TotalNumSecurities) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from SecurityList.
func (m SecurityList) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}
