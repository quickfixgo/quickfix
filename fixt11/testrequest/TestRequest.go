package testrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//TestRequest is the fixt11 TestRequest type, MsgType = 1
type TestRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TestRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) TestRequest {
	return TestRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TestRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TestRequest initialized with the required fields for TestRequest
func New(testreqid field.TestReqIDField) (m TestRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("1"))
	m.Header.Set(field.NewBeginString("FIXT.1.1"))
	m.Set(testreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TestRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIXT.1.1", "1", r
}

//SetTestReqID sets TestReqID, Tag 112
func (m TestRequest) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

//GetTestReqID gets TestReqID, Tag 112
func (m TestRequest) GetTestReqID() (f field.TestReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTestReqID returns true if TestReqID is present, Tag 112
func (m TestRequest) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}
