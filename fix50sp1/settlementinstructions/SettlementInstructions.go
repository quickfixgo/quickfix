//Package settlementinstructions msg type = T.
package settlementinstructions

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a SettlementInstructions wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//SettlInstMsgID is a required field for SettlementInstructions.
func (m Message) SettlInstMsgID() (*field.SettlInstMsgIDField, quickfix.MessageRejectError) {
	f := &field.SettlInstMsgIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMsgID reads a SettlInstMsgID from SettlementInstructions.
func (m Message) GetSettlInstMsgID(f *field.SettlInstMsgIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqID is a non-required field for SettlementInstructions.
func (m Message) SettlInstReqID() (*field.SettlInstReqIDField, quickfix.MessageRejectError) {
	f := &field.SettlInstReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqID reads a SettlInstReqID from SettlementInstructions.
func (m Message) GetSettlInstReqID(f *field.SettlInstReqIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstMode is a required field for SettlementInstructions.
func (m Message) SettlInstMode() (*field.SettlInstModeField, quickfix.MessageRejectError) {
	f := &field.SettlInstModeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstMode reads a SettlInstMode from SettlementInstructions.
func (m Message) GetSettlInstMode(f *field.SettlInstModeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//SettlInstReqRejCode is a non-required field for SettlementInstructions.
func (m Message) SettlInstReqRejCode() (*field.SettlInstReqRejCodeField, quickfix.MessageRejectError) {
	f := &field.SettlInstReqRejCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetSettlInstReqRejCode reads a SettlInstReqRejCode from SettlementInstructions.
func (m Message) GetSettlInstReqRejCode(f *field.SettlInstReqRejCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a non-required field for SettlementInstructions.
func (m Message) Text() (*field.TextField, quickfix.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from SettlementInstructions.
func (m Message) GetText(f *field.TextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedTextLen is a non-required field for SettlementInstructions.
func (m Message) EncodedTextLen() (*field.EncodedTextLenField, quickfix.MessageRejectError) {
	f := &field.EncodedTextLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedTextLen reads a EncodedTextLen from SettlementInstructions.
func (m Message) GetEncodedTextLen(f *field.EncodedTextLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedText is a non-required field for SettlementInstructions.
func (m Message) EncodedText() (*field.EncodedTextField, quickfix.MessageRejectError) {
	f := &field.EncodedTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedText reads a EncodedText from SettlementInstructions.
func (m Message) GetEncodedText(f *field.EncodedTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ClOrdID is a non-required field for SettlementInstructions.
func (m Message) ClOrdID() (*field.ClOrdIDField, quickfix.MessageRejectError) {
	f := &field.ClOrdIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetClOrdID reads a ClOrdID from SettlementInstructions.
func (m Message) GetClOrdID(f *field.ClOrdIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//TransactTime is a required field for SettlementInstructions.
func (m Message) TransactTime() (*field.TransactTimeField, quickfix.MessageRejectError) {
	f := &field.TransactTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTransactTime reads a TransactTime from SettlementInstructions.
func (m Message) GetTransactTime(f *field.TransactTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoSettlInst is a non-required field for SettlementInstructions.
func (m Message) NoSettlInst() (*field.NoSettlInstField, quickfix.MessageRejectError) {
	f := &field.NoSettlInstField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoSettlInst reads a NoSettlInst from SettlementInstructions.
func (m Message) GetNoSettlInst(f *field.NoSettlInstField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for SettlementInstructions.
func New(
	settlinstmsgid *field.SettlInstMsgIDField,
	settlinstmode *field.SettlInstModeField,
	transacttime *field.TransactTimeField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("T"))
	builder.Body.Set(settlinstmsgid)
	builder.Body.Set(settlinstmode)
	builder.Body.Set(transacttime)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP1, "T", r
}
