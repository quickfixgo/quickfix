package fix50

import (
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
func CreateSecurityListBuilder() SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m SecurityList) SecurityReqID() (field.SecurityReqID, error) {
	var f field.SecurityReqID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityResponseID is a non-required field for SecurityList.
func (m SecurityList) SecurityResponseID() (field.SecurityResponseID, error) {
	var f field.SecurityResponseID
	err := m.Body.Get(&f)
	return f, err
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (field.SecurityRequestResult, error) {
	var f field.SecurityRequestResult
	err := m.Body.Get(&f)
	return f, err
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) TotNoRelatedSym() (field.TotNoRelatedSym, error) {
	var f field.TotNoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LastFragment is a non-required field for SecurityList.
func (m SecurityList) LastFragment() (field.LastFragment, error) {
	var f field.LastFragment
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//SecurityReportID is a non-required field for SecurityList.
func (m SecurityList) SecurityReportID() (field.SecurityReportID, error) {
	var f field.SecurityReportID
	err := m.Body.Get(&f)
	return f, err
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m SecurityList) ClearingBusinessDate() (field.ClearingBusinessDate, error) {
	var f field.ClearingBusinessDate
	err := m.Body.Get(&f)
	return f, err
}
