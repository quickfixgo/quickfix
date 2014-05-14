//Package liststatus msg type = N.
package liststatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a ListStatus wrapper for the generic Message type
type Message struct {
	message.Message
}

//ListID is a required field for ListStatus.
func (m Message) ListID() (*field.ListIDField, errors.MessageRejectError) {
	f := &field.ListIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetListID reads a ListID from ListStatus.
func (m Message) GetListID(f *field.ListIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//WaveNo is a non-required field for ListStatus.
func (m Message) WaveNo() (*field.WaveNoField, errors.MessageRejectError) {
	f := &field.WaveNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetWaveNo reads a WaveNo from ListStatus.
func (m Message) GetWaveNo(f *field.WaveNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRpts is a required field for ListStatus.
func (m Message) NoRpts() (*field.NoRptsField, errors.MessageRejectError) {
	f := &field.NoRptsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRpts reads a NoRpts from ListStatus.
func (m Message) GetNoRpts(f *field.NoRptsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RptSeq is a required field for ListStatus.
func (m Message) RptSeq() (*field.RptSeqField, errors.MessageRejectError) {
	f := &field.RptSeqField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRptSeq reads a RptSeq from ListStatus.
func (m Message) GetRptSeq(f *field.RptSeqField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoOrders is a required field for ListStatus.
func (m Message) NoOrders() (*field.NoOrdersField, errors.MessageRejectError) {
	f := &field.NoOrdersField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoOrders reads a NoOrders from ListStatus.
func (m Message) GetNoOrders(f *field.NoOrdersField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds ListStatus messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for ListStatus.
func Builder(
	listid *field.ListIDField,
	norpts *field.NoRptsField,
	rptseq *field.RptSeqField,
	noorders *field.NoOrdersField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX41))
	builder.Header().Set(field.NewMsgType("N"))
	builder.Body().Set(listid)
	builder.Body().Set(norpts)
	builder.Body().Set(rptseq)
	builder.Body().Set(noorders)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX41, "N", r
}
