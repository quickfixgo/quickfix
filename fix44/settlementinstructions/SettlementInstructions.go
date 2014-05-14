//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a SettlementInstructions wrapper for the generic Message type
type Message struct {
	message.Message
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m Message) SettlInstMsgID() (*field.SettlInstMsgIDField, errors.MessageRejectError) {
	f := &field.SettlInstMsgIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMsgID reads a SettlInstMsgID from SettlementInstructions.
func (m Message) GetSettlInstMsgID(f *field.SettlInstMsgIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m Message) SettlInstReqID() (*field.SettlInstReqIDField, errors.MessageRejectError) {
	f := &field.SettlInstReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructions.
func (m Message) GetSettlInstReqID(f *field.SettlInstReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m Message) SettlInstMode() (*field.SettlInstModeField, errors.MessageRejectError) {
	f := &field.SettlInstModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m Message) GetSettlInstMode(f *field.SettlInstModeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m Message) SettlInstReqRejCode() (*field.SettlInstReqRejCodeField, errors.MessageRejectError) {
	f := &field.SettlInstReqRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqRejCode reads a SettlInstReqRejCode from SettlementInstructions.
func (m Message) GetSettlInstReqRejCode(f *field.SettlInstReqRejCodeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementInstructions.
func (m Message) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementInstructions.
func (m Message) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, errors.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementInstructions.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementInstructions.
func (m Message) EncodedText() (*field.EncodedTextField, errors.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementInstructions.
func (m Message) GetEncodedText(f *field.EncodedTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m Message) ClOrdID() (*field.ClOrdIDField, errors.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m Message) GetClOrdID(f *field.ClOrdIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m Message) TransactTime() (*field.TransactTimeField, errors.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m Message) GetTransactTime(f *field.TransactTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m Message) NoSettlInst() (*field.NoSettlInstField, errors.MessageRejectError) {
	f := &field.NoSettlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlInst reads a NoSettlInst from SettlementInstructions.
func (m Message) GetNoSettlInst(f *field.NoSettlInstField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SettlementInstructions messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SettlementInstructions.
func Builder(
	settlinstmsgid *field.SettlInstMsgIDField,
	settlinstmode *field.SettlInstModeField,
	transacttime *field.TransactTimeField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(field.NewMsgType("T"))
	builder.Body().Set(settlinstmsgid)
	builder.Body().Set(settlinstmode)
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
	return fix.BeginString_FIX44, "T", r
}
