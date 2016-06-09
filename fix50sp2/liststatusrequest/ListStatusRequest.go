package liststatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//ListStatusRequest is the fix50sp2 ListStatusRequest type, MsgType = M
type ListStatusRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a ListStatusRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) ListStatusRequest {
	return ListStatusRequest{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m ListStatusRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a ListStatusRequest initialized with the required fields for ListStatusRequest
func New(listid field.ListIDField) (m ListStatusRequest) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("M"))
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
	return "9", "M", r
}

//SetText sets Text, Tag 58
func (m ListStatusRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetListID sets ListID, Tag 66
func (m ListStatusRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ListStatusRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ListStatusRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
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

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ListStatusRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ListStatusRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
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

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ListStatusRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ListStatusRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}
