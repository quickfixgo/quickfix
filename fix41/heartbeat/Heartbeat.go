package heartbeat

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//Heartbeat is the fix41 Heartbeat type, MsgType = 0
type Heartbeat struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a Heartbeat from a quickfix.Message instance
func FromMessage(m quickfix.Message) Heartbeat {
	return Heartbeat{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Heartbeat) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a Heartbeat initialized with the required fields for Heartbeat
func New() (m Heartbeat) {
	m.Header = fix41.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("0"))
	m.Header.Set(field.NewBeginString("FIX.4.1"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Heartbeat, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "0", r
}

//SetTestReqID sets TestReqID, Tag 112
func (m Heartbeat) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

//GetTestReqID gets TestReqID, Tag 112
func (m Heartbeat) GetTestReqID() (f field.TestReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTestReqID returns true if TestReqID is present, Tag 112
func (m Heartbeat) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}
