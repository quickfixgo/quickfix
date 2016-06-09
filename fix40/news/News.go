package news

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//News is the fix40 News type, MsgType = B
type News struct {
	fix40.Header
	quickfix.Body
	fix40.Trailer
}

//FromMessage creates a News from a quickfix.Message instance
func FromMessage(m quickfix.Message) News {
	return News{
		Header:  fix40.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix40.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m News) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a News initialized with the required fields for News
func New(linesoftext field.LinesOfTextField, text field.TextField) (m News) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("B"))
	m.Set(linesoftext)
	m.Set(text)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg News, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
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
func (m News) SetUrgency(v string) {
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
func (m News) GetLinesOfText() (f field.LinesOfTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigTime gets OrigTime, Tag 42
func (m News) GetOrigTime() (f field.OrigTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRelatdSym gets RelatdSym, Tag 46
func (m News) GetRelatdSym() (f field.RelatdSymField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m News) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUrgency gets Urgency, Tag 61
func (m News) GetUrgency() (f field.UrgencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawDataLength gets RawDataLength, Tag 95
func (m News) GetRawDataLength() (f field.RawDataLengthField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRawData gets RawData, Tag 96
func (m News) GetRawData() (f field.RawDataField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
