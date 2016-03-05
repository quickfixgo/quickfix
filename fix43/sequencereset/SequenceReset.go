//Package sequencereset msg type = 4.
package sequencereset

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
)

//Message is a SequenceReset FIX Message
type Message struct {
	FIXMsgType string `fix:"4"`
	fix43.Header
	//GapFillFlag is a non-required field for SequenceReset.
	GapFillFlag *bool `fix:"123"`
	//NewSeqNo is a required field for SequenceReset.
	NewSeqNo int `fix:"36"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetGapFillFlag(v bool) { m.GapFillFlag = &v }
func (m *Message) SetNewSeqNo(v int)     { m.NewSeqNo = v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX43, "4", r
}
