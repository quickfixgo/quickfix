package logout

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//Logout is the fixt11 Logout type, MsgType = 5
type Logout struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
}

//FromMessage creates a Logout from a quickfix.Message instance
func FromMessage(m quickfix.Message) Logout {
	return Logout{
		Header:  fixt11.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fixt11.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m Logout) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a Logout initialized with the required fields for Logout
func New() (m Logout) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("5"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Logout, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIXT.1.1", "5", r
}

//SetText sets Text, Tag 58
func (m Logout) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m Logout) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m Logout) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetText gets Text, Tag 58
func (m Logout) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m Logout) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m Logout) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasText returns true if Text is present, Tag 58
func (m Logout) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m Logout) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m Logout) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}
