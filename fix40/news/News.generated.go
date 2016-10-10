package news

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//News is the fix40 News type, MsgType = B
type News struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a News from a quickfix.Message instance
func FromMessage(m *quickfix.Message) News {
	return News{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m News) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a News initialized with the required fields for News
func New(linesoftext field.LinesOfTextField, text field.TextField) (m News) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("B"))
	m.Set(linesoftext)
	m.Set(text)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg News, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "B", r
}

//SetLinesOfText sets LinesOfText, Tag 33
func (m News) SetLinesOfText(v int) {
	m.Set(field.NewLinesOfText(v))
}

//SetOrigTime sets OrigTime, Tag 42
func (m News) SetOrigTime(v time.Time) {
	m.Set(field.NewOrigTime(v))
}

//SetRelatdSym sets RelatdSym, Tag 46
func (m News) SetRelatdSym(v string) {
	m.Set(field.NewRelatdSym(v))
}

//SetText sets Text, Tag 58
func (m News) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetUrgency sets Urgency, Tag 61
func (m News) SetUrgency(v enum.Urgency) {
	m.Set(field.NewUrgency(v))
}

//SetRawDataLength sets RawDataLength, Tag 95
func (m News) SetRawDataLength(v int) {
	m.Set(field.NewRawDataLength(v))
}

//SetRawData sets RawData, Tag 96
func (m News) SetRawData(v string) {
	m.Set(field.NewRawData(v))
}

//GetLinesOfText gets LinesOfText, Tag 33
func (m News) GetLinesOfText() (v int, err quickfix.MessageRejectError) {
	var f field.LinesOfTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrigTime gets OrigTime, Tag 42
func (m News) GetOrigTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.OrigTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRelatdSym gets RelatdSym, Tag 46
func (m News) GetRelatdSym() (v string, err quickfix.MessageRejectError) {
	var f field.RelatdSymField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m News) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUrgency gets Urgency, Tag 61
func (m News) GetUrgency() (v enum.Urgency, err quickfix.MessageRejectError) {
	var f field.UrgencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m News) GetRawDataLength() (v int, err quickfix.MessageRejectError) {
	var f field.RawDataLengthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRawData gets RawData, Tag 96
func (m News) GetRawData() (v string, err quickfix.MessageRejectError) {
	var f field.RawDataField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLinesOfText returns true if LinesOfText is present, Tag 33
func (m News) HasLinesOfText() bool {
	return m.Has(tag.LinesOfText)
}

//HasOrigTime returns true if OrigTime is present, Tag 42
func (m News) HasOrigTime() bool {
	return m.Has(tag.OrigTime)
}

//HasRelatdSym returns true if RelatdSym is present, Tag 46
func (m News) HasRelatdSym() bool {
	return m.Has(tag.RelatdSym)
}

//HasText returns true if Text is present, Tag 58
func (m News) HasText() bool {
	return m.Has(tag.Text)
}

//HasUrgency returns true if Urgency is present, Tag 61
func (m News) HasUrgency() bool {
	return m.Has(tag.Urgency)
}

//HasRawDataLength returns true if RawDataLength is present, Tag 95
func (m News) HasRawDataLength() bool {
	return m.Has(tag.RawDataLength)
}

//HasRawData returns true if RawData is present, Tag 96
func (m News) HasRawData() bool {
	return m.Has(tag.RawData)
}
