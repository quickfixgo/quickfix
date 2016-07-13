package liststatusrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//ListStatusRequest is the fix40 ListStatusRequest type, MsgType = M
type ListStatusRequest struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListStatusRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListStatusRequest {
	return ListStatusRequest{
		Header:      fix40.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix40.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListStatusRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListStatusRequest initialized with the required fields for ListStatusRequest
func New(listid field.ListIDField) (m ListStatusRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("M"))
	m.Header.Set(field.NewBeginString("FIX.4.0"))
	m.Set(listid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListStatusRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "M", r
}

//SetText sets Text, Tag 58
func (m ListStatusRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetListID sets ListID, Tag 66
func (m ListStatusRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetWaveNo sets WaveNo, Tag 105
func (m ListStatusRequest) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//GetText gets Text, Tag 58
func (m ListStatusRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ListStatusRequest) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m ListStatusRequest) GetWaveNo() (f field.WaveNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ListStatusRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasListID returns true if ListID is present, Tag 66
func (m ListStatusRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m ListStatusRequest) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}
