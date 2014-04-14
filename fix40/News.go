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

//NewNewsBuilder returns an initialized NewsBuilder with specified required fields.
func NewNewsBuilder(
	linesoftext field.LinesOfText,
	text field.Text) *NewsBuilder {
	builder := new(NewsBuilder)
	builder.Body.Set(linesoftext)
	builder.Body.Set(text)
	return builder
}

//OrigTime is a non-required field for News.
func (m *News) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}

//Urgency is a non-required field for News.
func (m *News) Urgency() (*field.Urgency, error) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}

//RelatdSym is a non-required field for News.
func (m *News) RelatdSym() (*field.RelatdSym, error) {
	f := new(field.RelatdSym)
	err := m.Body.Get(f)
	return f, err
}

//LinesOfText is a required field for News.
func (m *News) LinesOfText() (*field.LinesOfText, error) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//Text is a required field for News.
func (m *News) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}

//RawDataLength is a non-required field for News.
func (m *News) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}

//RawData is a non-required field for News.
func (m *News) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
