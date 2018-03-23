package testrequest

import (
	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fix42"
	"github.com/alpacahq/quickfix/tag"
)

//TestRequest is the fix42 TestRequest type, MsgType = 1
type TestRequest struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a TestRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TestRequest {
	return TestRequest{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TestRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a TestRequest initialized with the required fields for TestRequest
func New(testreqid field.TestReqIDField) (m TestRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("1"))
	m.Set(testreqid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TestRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "1", r
}

//SetTestReqID sets TestReqID, Tag 112
func (m TestRequest) SetTestReqID(v string) {
	m.Set(field.NewTestReqID(v))
}

//GetTestReqID gets TestReqID, Tag 112
func (m TestRequest) GetTestReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TestReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasTestReqID returns true if TestReqID is present, Tag 112
func (m TestRequest) HasTestReqID() bool {
	return m.Has(tag.TestReqID)
}
