package fix40

import (
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
func (m News) OrigTime() (field.OrigTime, error) {
	var f field.OrigTime
	err := m.Body.Get(&f)
	return f, err
}

//Urgency is a non-required field for News.
func (m News) Urgency() (field.Urgency, error) {
	var f field.Urgency
	err := m.Body.Get(&f)
	return f, err
}

//RelatdSym is a non-required field for News.
func (m News) RelatdSym() (field.RelatdSym, error) {
	var f field.RelatdSym
	err := m.Body.Get(&f)
	return f, err
}

//LinesOfText is a required field for News.
func (m News) LinesOfText() (field.LinesOfText, error) {
	var f field.LinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//Text is a required field for News.
func (m News) Text() (field.Text, error) {
	var f field.Text
	err := m.Body.Get(&f)
	return f, err
}

//RawDataLength is a non-required field for News.
func (m News) RawDataLength() (field.RawDataLength, error) {
	var f field.RawDataLength
	err := m.Body.Get(&f)
	return f, err
}

//RawData is a non-required field for News.
func (m News) RawData() (field.RawData, error) {
	var f field.RawData
	err := m.Body.Get(&f)
	return f, err
}
