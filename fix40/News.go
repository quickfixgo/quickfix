package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//News msg type = B.
type News struct {
	message.Message
}

//NewsBuilder builds News messages.
type NewsBuilder struct {
	message.MessageBuilder
}

//CreateNewsBuilder returns an initialized NewsBuilder with specified required fields.
func CreateNewsBuilder(
	linesoftext *field.LinesOfTextField,
	text *field.TextField) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.NewMsgType("B"))
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//OrigTime is a non-required field for News.
func (m News) OrigTime() (*field.OrigTimeField, errors.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m News) GetOrigTime(f *field.OrigTimeField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m News) Urgency() (*field.UrgencyField, errors.MessageRejectError) {
	f := &field.UrgencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m News) GetUrgency(f *field.UrgencyField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RelatdSym is a non-required field for News.
func (m News) RelatdSym() (*field.RelatdSymField, errors.MessageRejectError) {
	f := &field.RelatdSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRelatdSym reads a RelatdSym from News.
func (m News) GetRelatdSym(f *field.RelatdSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for News.
func (m News) LinesOfText() (*field.LinesOfTextField, errors.MessageRejectError) {
	f := &field.LinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from News.
func (m News) GetLinesOfText(f *field.LinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a required field for News.
func (m News) Text() (*field.TextField, errors.MessageRejectError) {
	f := &field.TextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from News.
func (m News) GetText(f *field.TextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m News) RawDataLength() (*field.RawDataLengthField, errors.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m News) GetRawDataLength(f *field.RawDataLengthField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m News) RawData() (*field.RawDataField, errors.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m News) GetRawData(f *field.RawDataField) errors.MessageRejectError {
	return m.Body.Get(f)
}
