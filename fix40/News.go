package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
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
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//OrigTime is a non-required field for News.
func (m News) OrigTime() (field.OrigTime, errors.MessageRejectError) {
	var f field.OrigTime
	err := m.Body.Get(&f)
	return f, err
}

//Urgency is a non-required field for News.
func (m News) Urgency() (field.Urgency, errors.MessageRejectError) {
	var f field.Urgency
	err := m.Body.Get(&f)
	return f, err
}

//RelatdSym is a non-required field for News.
func (m News) RelatdSym() (field.RelatdSym, errors.MessageRejectError) {
	var f field.RelatdSym
	err := m.Body.Get(&f)
	return f, err
}

//LinesOfText is a required field for News.
func (m News) LinesOfText() (field.LinesOfText, errors.MessageRejectError) {
	var f field.LinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//Text is a required field for News.
func (m News) Text() (field.Text, errors.MessageRejectError) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for News.
func (m News) RawDataLength() (field.RawDataLength, errors.MessageRejectError) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for News.
func (m News) RawData() (field.RawData, errors.MessageRejectError) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}
