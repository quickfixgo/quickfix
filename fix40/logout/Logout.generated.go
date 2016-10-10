package logout

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Logout is the fix40 Logout type, MsgType = 5
type Logout struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Logout from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Logout {
	return Logout{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Logout) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Logout initialized with the required fields for Logout
func New() (m Logout) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("5"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Logout, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "5", r
}

//SetText sets Text, Tag 58
func (m Logout) SetText(v string) {
	m.Set(field.NewText(v))
}

//GetText gets Text, Tag 58
func (m Logout) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m Logout) HasText() bool {
	return m.Has(tag.Text)
}
