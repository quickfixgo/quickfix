//Package sequencereset msg type = 4.
package sequencereset

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//Message is a SequenceReset wrapper for the generic Message type
type Message struct {
	message.Message
}

//GapFillFlag is a non-required field for SequenceReset.
func (m Message) GapFillFlag() (*field.GapFillFlagField, errors.MessageRejectError) {
	f := &field.GapFillFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGapFillFlag reads a GapFillFlag from SequenceReset.
func (m Message) GetGapFillFlag(f *field.GapFillFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NewSeqNo is a required field for SequenceReset.
func (m Message) NewSeqNo() (*field.NewSeqNoField, errors.MessageRejectError) {
	f := &field.NewSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewSeqNo reads a NewSeqNo from SequenceReset.
func (m Message) GetNewSeqNo(f *field.NewSeqNoField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//MessageBuilder builds SequenceReset messages.
type MessageBuilder struct {
	message.MessageBuilder
}

//Builder returns an initialized MessageBuilder with specified required fields for SequenceReset.
func Builder(
	newseqno *field.NewSeqNoField) MessageBuilder {
	var builder MessageBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header().Set(field.NewMsgType("4"))
	builder.Body().Set(newseqno)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) errors.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg message.Message, sessionID quickfix.SessionID) errors.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return fix.BeginString_FIX40, "4", r
}
