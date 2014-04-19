package fix50sp1

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(applresponseid)
	return builder
}

//ApplResponseID is a required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseID() (field.ApplResponseID, errors.MessageRejectError) {
	var f field.ApplResponseID
	err := m.Body.Get(&f)
	return f, err
}

//ApplReqID is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqID() (field.ApplReqID, errors.MessageRejectError) {
	var f field.ApplReqID
	err := m.Body.Get(&f)
	return f, err
}

//ApplReqType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplReqType() (field.ApplReqType, errors.MessageRejectError) {
	var f field.ApplReqType
	err := m.Body.Get(&f)
	return f, err
}

//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplResponseType() (field.ApplResponseType, errors.MessageRejectError) {
	var f field.ApplResponseType
	err := m.Body.Get(&f)
	return f, err
}

//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) ApplTotalMessageCount() (field.ApplTotalMessageCount, errors.MessageRejectError) {
	var f field.ApplTotalMessageCount
	err := m.Body.Get(&f)
	return f, err
}

//NoApplIDs is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) NoApplIDs() (field.NoApplIDs, errors.MessageRejectError) {
	var f field.NoApplIDs
	err := m.Body.Get(&f)
	return f, err
}

//Text is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedTextLen() (field.EncodedTextLen, errors.MessageRejectError) {
	var f field.EncodedTextLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedText is a non-required field for ApplicationMessageRequestAck.
func (m ApplicationMessageRequestAck) EncodedText() (field.EncodedText, errors.MessageRejectError) {
	var f field.EncodedText
	err := m.Body.Get(&f)
	return f, err
}
