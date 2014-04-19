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
func (m SecurityList) SecurityReqID() (field.SecurityReqID, errors.MessageRejectError) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a required field for SecurityList.
func (m SecurityList) SecurityResponseID() (field.SecurityResponseID, errors.MessageRejectError) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (field.SecurityRequestResult, errors.MessageRejectError) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotalNumSecurities is a non-required field for SecurityList.
func (m SecurityList) TotalNumSecurities() (field.TotalNumSecurities, errors.MessageRejectError) {
	var f field.TotalNumSecurities
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}
