//Package streamassignmentreportack msg type = CE.
package streamassignmentreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a StreamAssignmentReportACK FIX Message
type Message struct {
	FIXMsgType string `fix:"CE"`
	Header     fixt11.Header
	//StreamAsgnAckType is a required field for StreamAssignmentReportACK.
	StreamAsgnAckType int `fix:"1503"`
	//StreamAsgnRptID is a required field for StreamAssignmentReportACK.
	StreamAsgnRptID string `fix:"1501"`
	//StreamAsgnRejReason is a non-required field for StreamAssignmentReportACK.
	StreamAsgnRejReason *int `fix:"1502"`
	//Text is a non-required field for StreamAssignmentReportACK.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for StreamAssignmentReportACK.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for StreamAssignmentReportACK.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetStreamAsgnAckType(v int)   { m.StreamAsgnAckType = v }
func (m *Message) SetStreamAsgnRptID(v string)  { m.StreamAsgnRptID = v }
func (m *Message) SetStreamAsgnRejReason(v int) { m.StreamAsgnRejReason = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "CE", r
}
