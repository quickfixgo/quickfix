//Package applicationmessagereport msg type = BY.
package applicationmessagereport

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

//Message is a ApplicationMessageReport wrapper for the generic Message type
type Message struct {
	message.Message
}

//ApplReportID is a required field for ApplicationMessageReport.
func (m Message) ApplReportID() (*field.ApplReportIDField, errors.MessageRejectError) {
	f := &field.ApplReportIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportID reads a ApplReportID from ApplicationMessageReport.
func (m Message) GetApplReportID(f *field.ApplReportIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReportType is a required field for ApplicationMessageReport.
func (m Message) ApplReportType() (*field.ApplReportTypeField, errors.MessageRejectError) {
	f := &field.ApplReportTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplReportType reads a ApplReportType from ApplicationMessageReport.
func (m Message) GetApplReportType(f *field.ApplReportTypeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageReport.
func (m Message) NoApplIDs() (*field.NoApplIDsField, errors.MessageRejectError) {
	f := &field.NoApplIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageReport.
func (m Message) GetNoApplIDs(f *field.NoApplIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageReport.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageReport.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageReport.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageReport.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageReport.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageReport.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ApplicationMessageReport messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ApplicationMessageReport.
func Builder(
	applreportid *field.ApplReportIDField,
	applreporttype *field.ApplReportTypeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header().Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header().Set(field.NewMsgType("BY"))
	builder.Body().Set(applreportid)
	builder.Body().Set(applreporttype)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "BY", r
}
