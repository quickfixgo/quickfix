package fix50sp2

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//ApplicationMessageRequestAck msg type = BX.
type ApplicationMessageRequestAck struct {
	message.Message
}

//ApplicationMessageRequestAckBuilder builds ApplicationMessageRequestAck messages.
type ApplicationMessageRequestAckBuilder struct {
	message.MessageBuilder
}

//CreateApplicationMessageRequestAckBuilder returns an initialized ApplicationMessageRequestAckBuilder with specified required fields.
func CreateApplicationMessageRequestAckBuilder(
	applresponseid field.ApplResponseID) ApplicationMessageRequestAckBuilder {
	var builder ApplicationMessageRequestAckBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildMsgType("BX"))
	builder.Body.Set(applresponseid)
	return builder
}

//ApplResponseID is a required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseID() (*field.ApplResponseID, errors.MessageRejectError) {
	f := new(field.ApplResponseID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseID reads a ApplResponseID from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplResponseID(f *field.ApplResponseID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqID() (*field.ApplReqID, errors.MessageRejectError) {
	f := new(field.ApplReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqID reads a ApplReqID from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplReqID(f *field.ApplReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqType() (*field.ApplReqType, errors.MessageRejectError) {
	f := new(field.ApplReqType)
	err := m.Body.Get(f)
	return f, err
}

//GetApplReqType reads a ApplReqType from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplReqType(f *field.ApplReqType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseType() (*field.ApplResponseType, errors.MessageRejectError) {
	f := new(field.ApplResponseType)
	err := m.Body.Get(f)
	return f, err
}

//GetApplResponseType reads a ApplResponseType from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplResponseType(f *field.ApplResponseType) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplTotalMessageCount() (*field.ApplTotalMessageCount, errors.MessageRejectError) {
	f := new(field.ApplTotalMessageCount)
	err := m.Body.Get(f)
	return f, err
}

//GetApplTotalMessageCount reads a ApplTotalMessageCount from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetApplTotalMessageCount(f *field.ApplTotalMessageCount) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) NoApplIDs() (*field.NoApplIDs, errors.MessageRejectError) {
	f := new(field.NoApplIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoApplIDs reads a NoApplIDs from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetNoApplIDs(f *field.NoApplIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedTextLen() (*field.EncodedTextLen, errors.MessageRejectError) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetEncodedTextLen(f *field.EncodedTextLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedText() (*field.EncodedText, errors.MessageRejectError) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetEncodedText(f *field.EncodedText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoPartyIDs is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) NoPartyIDs() (*field.NoPartyIDs, errors.MessageRejectError) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoPartyIDs reads a NoPartyIDs from ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) GetNoPartyIDs(f *field.NoPartyIDs) errors.MessageRejectError {
	return m.Body.Get(f)
}
