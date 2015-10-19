//Package email msg type = C.
package email

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
)

//Message is a Email wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//EmailThreadID is a required field for Email.
func (m Message) EmailThreadID() (*field.EmailThreadIDField, quickfix.MessageRejectError) {
	f := &field.EmailThreadIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEmailThreadID reads a EmailThreadID from Email.
func (m Message) GetEmailThreadID(f *field.EmailThreadIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EmailType is a required field for Email.
func (m Message) EmailType() (*field.EmailTypeField, quickfix.MessageRejectError) {
	f := &field.EmailTypeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEmailType reads a EmailType from Email.
func (m Message) GetEmailType(f *field.EmailTypeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrigTime is a non-required field for Email.
func (m Message) OrigTime() (*field.OrigTimeField, quickfix.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from Email.
func (m Message) GetOrigTime(f *field.OrigTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Subject is a required field for Email.
func (m Message) Subject() (*field.SubjectField, quickfix.MessageRejectError) {
	f := &field.SubjectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSubject reads a Subject from Email.
func (m Message) GetSubject(f *field.SubjectField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSubjectLen is a non-required field for Email.
func (m Message) EncodedSubjectLen() (*field.EncodedSubjectLenField, quickfix.MessageRejectError) {
	f := &field.EncodedSubjectLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSubjectLen reads a EncodedSubjectLen from Email.
func (m Message) GetEncodedSubjectLen(f *field.EncodedSubjectLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedSubject is a non-required field for Email.
func (m Message) EncodedSubject() (*field.EncodedSubjectField, quickfix.MessageRejectError) {
	f := &field.EncodedSubjectField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedSubject reads a EncodedSubject from Email.
func (m Message) GetEncodedSubject(f *field.EncodedSubjectField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for Email.
func (m Message) NoRoutingIDs() (*field.NoRoutingIDsField, quickfix.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from Email.
func (m Message) GetNoRoutingIDs(f *field.NoRoutingIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for Email.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from Email.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for Email.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from Email.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for Email.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from Email.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//OrderID is a non-required field for Email.
func (m Message) OrderID() (*field.OrderIDField, quickfix.MessageRejectError) {
	f := &field.OrderIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrderID reads a OrderID from Email.
func (m Message) GetOrderID(f *field.OrderIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for Email.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from Email.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLinesOfText is a required field for Email.
func (m Message) NoLinesOfText() (*field.NoLinesOfTextField, quickfix.MessageRejectError) {
	f := &field.NoLinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLinesOfText reads a NoLinesOfText from Email.
func (m Message) GetNoLinesOfText(f *field.NoLinesOfTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for Email.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from Email.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for Email.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from Email.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds Email messages.
type MessageBuilder struct {
	quickfix.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for Email.
func Builder(
	emailthreadid *field.EmailThreadIDField,
	emailtype *field.EmailTypeField,
	subject *field.SubjectField,
	nolinesoftext *field.NoLinesOfTextField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = *quickfix.NewMessageBuilder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("C"))
	builder.Body.Set(emailthreadid)
	builder.Body.Set(emailtype)
	builder.Body.Set(subject)
	builder.Body.Set(nolinesoftext)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "C", r
}
