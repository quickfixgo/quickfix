package resendrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//ResendRequest is the fix42 ResendRequest type, MsgType = 2
type ResendRequest struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ResendRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ResendRequest {
	return ResendRequest{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ResendRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ResendRequest initialized with the required fields for ResendRequest
func New(beginseqno field.BeginSeqNoField, endseqno field.EndSeqNoField) (m ResendRequest) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("2"))
	m.Set(beginseqno)
	m.Set(endseqno)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ResendRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "2", r
}

//SetBeginSeqNo sets BeginSeqNo, Tag 7
func (m ResendRequest) SetBeginSeqNo(v int) {
	m.Set(field.NewBeginSeqNo(v))
}

//SetEndSeqNo sets EndSeqNo, Tag 16
func (m ResendRequest) SetEndSeqNo(v int) {
	m.Set(field.NewEndSeqNo(v))
}

//GetBeginSeqNo gets BeginSeqNo, Tag 7
func (m ResendRequest) GetBeginSeqNo() (f field.BeginSeqNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEndSeqNo gets EndSeqNo, Tag 16
func (m ResendRequest) GetEndSeqNo() (f field.EndSeqNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasBeginSeqNo returns true if BeginSeqNo is present, Tag 7
func (m ResendRequest) HasBeginSeqNo() bool {
	return m.Has(tag.BeginSeqNo)
}

//HasEndSeqNo returns true if EndSeqNo is present, Tag 16
func (m ResendRequest) HasEndSeqNo() bool {
	return m.Has(tag.EndSeqNo)
}
