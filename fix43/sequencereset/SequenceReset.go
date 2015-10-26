//Package sequencereset msg type = 4.
package sequencereset

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
)

//Message is a SequenceReset wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//GapFillFlag is a non-required field for SequenceReset.
func (m Message) GapFillFlag() (*field.GapFillFlagField, quickfix.MessageRejectError) {
	f := &field.GapFillFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetGapFillFlag reads a GapFillFlag from SequenceReset.
func (m Message) GetGapFillFlag(f *field.GapFillFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NewSeqNo is a required field for SequenceReset.
func (m Message) NewSeqNo() (*field.NewSeqNoField, quickfix.MessageRejectError) {
	f := &field.NewSeqNoField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewSeqNo reads a NewSeqNo from SequenceReset.
func (m Message) GetNewSeqNo(f *field.NewSeqNoField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for SequenceReset.
func New(
	newseqno *field.NewSeqNoField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIX43))
	builder.Header.Set(field.NewMsgType("4"))
	builder.Body.Set(newseqno)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.BeginStringFIX43, "4", r
}
