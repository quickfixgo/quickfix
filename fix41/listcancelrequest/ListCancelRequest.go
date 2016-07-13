package listcancelrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//ListCancelRequest is the fix41 ListCancelRequest type, MsgType = K
type ListCancelRequest struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ListCancelRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListCancelRequest {
	return ListCancelRequest{
		Header:      fix41.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix41.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListCancelRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ListCancelRequest initialized with the required fields for ListCancelRequest
func New(listid field.ListIDField) (m ListCancelRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("K"))
	m.Header.Set(field.NewBeginString("FIX.4.1"))
	m.Set(listid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ListCancelRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "K", r
}

//SetText sets Text, Tag 58
func (m ListCancelRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetListID sets ListID, Tag 66
func (m ListCancelRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetWaveNo sets WaveNo, Tag 105
func (m ListCancelRequest) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//GetText gets Text, Tag 58
func (m ListCancelRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ListCancelRequest) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m ListCancelRequest) GetWaveNo() (f field.WaveNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m ListCancelRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasListID returns true if ListID is present, Tag 66
func (m ListCancelRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m ListCancelRequest) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}
