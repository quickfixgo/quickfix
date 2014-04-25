package fix50

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
func CreateSecurityListBuilder() SecurityListBuilder {
	var builder SecurityListBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildMsgType("y"))
	return builder
}

//SecurityReqID is a non-required field for SecurityList.
func (m SecurityList) SecurityReqID() (*field.SecurityReqID, errors.MessageRejectError) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReqID reads a SecurityReqID from SecurityList.
func (m SecurityList) GetSecurityReqID(f *field.SecurityReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityResponseID is a non-required field for SecurityList.
func (m SecurityList) SecurityResponseID() (*field.SecurityResponseID, errors.MessageRejectError) {
	f := new(field.SecurityResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityResponseID reads a SecurityResponseID from SecurityList.
func (m SecurityList) GetSecurityResponseID(f *field.SecurityResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SecurityRequestResult is a non-required field for SecurityList.
func (m SecurityList) SecurityRequestResult() (*field.SecurityRequestResult, errors.MessageRejectError) {
	f := new(field.SecurityRequestResult)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityRequestResult reads a SecurityRequestResult from SecurityList.
func (m SecurityList) GetSecurityRequestResult(f *field.SecurityRequestResult) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TotNoRelatedSym is a non-required field for SecurityList.
func (m SecurityList) TotNoRelatedSym() (*field.TotNoRelatedSym, errors.MessageRejectError) {
	f := new(field.TotNoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetTotNoRelatedSym reads a TotNoRelatedSym from SecurityList.
func (m SecurityList) GetTotNoRelatedSym(f *field.TotNoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LastFragment is a non-required field for SecurityList.
func (m SecurityList) LastFragment() (*field.LastFragment, errors.MessageRejectError) {
	f := new(field.LastFragment)
	err := m.Body.Get(f)
	return f, err
}

//GetLastFragment reads a LastFragment from SecurityList.
func (m SecurityList) GetLastFragment(f *field.LastFragment) errors.MessageRejectError {
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

//SecurityReportID is a non-required field for SecurityList.
func (m SecurityList) SecurityReportID() (*field.SecurityReportID, errors.MessageRejectError) {
	f := new(field.SecurityReportID)
	err := m.Body.Get(f)
	return f, err
}

//GetSecurityReportID reads a SecurityReportID from SecurityList.
func (m SecurityList) GetSecurityReportID(f *field.SecurityReportID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClearingBusinessDate is a non-required field for SecurityList.
func (m SecurityList) ClearingBusinessDate() (*field.ClearingBusinessDate, errors.MessageRejectError) {
	f := new(field.ClearingBusinessDate)
	err := m.Body.Get(f)
	return f, err
}

//GetClearingBusinessDate reads a ClearingBusinessDate from SecurityList.
func (m SecurityList) GetClearingBusinessDate(f *field.ClearingBusinessDate) errors.MessageRejectError {
	return m.Body.Get(f)
}
