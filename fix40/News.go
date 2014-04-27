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
	linesoftext field.LinesOfText,
	text field.Text) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.BuildMsgType("B"))
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//OrigTime is a non-required field for News.
func (m News) OrigTime() (*field.OrigTime, errors.MessageRejectError) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m News) GetOrigTime(f *field.OrigTime) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m News) Urgency() (*field.Urgency, errors.MessageRejectError) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m News) GetUrgency(f *field.Urgency) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RelatdSym is a non-required field for News.
func (m News) RelatdSym() (*field.RelatdSym, errors.MessageRejectError) {
	f := new(field.RelatdSym)
	err := m.Body.Get(f)
	return f, err
}

//GetRelatdSym reads a RelatdSym from News.
func (m News) GetRelatdSym(f *field.RelatdSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//LinesOfText is a required field for News.
func (m News) LinesOfText() (*field.LinesOfText, errors.MessageRejectError) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//GetLinesOfText reads a LinesOfText from News.
func (m News) GetLinesOfText(f *field.LinesOfText) errors.MessageRejectError {
	return m.Body.Get(f)
}

//Text is a required field for News.
func (m News) Text() (*field.Text, errors.MessageRejectError) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//GetText reads a Text from News.
func (m News) GetText(f *field.Text) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m News) RawDataLength() (*field.RawDataLength, errors.MessageRejectError) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m News) GetRawDataLength(f *field.RawDataLength) errors.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m News) RawData() (*field.RawData, errors.MessageRejectError) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m News) GetRawData(f *field.RawData) errors.MessageRejectError {
	return m.Body.Get(f)
}
