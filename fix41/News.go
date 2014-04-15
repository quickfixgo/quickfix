package fix41

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
	headline field.Headline,
	linesoftext field.LinesOfText) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(headline)
	builder.Body.Set(linesoftext)
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

//Headline is a required field for News.
func (m News) Headline() (field.Headline, error) {
	var f field.Headline
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for News.
func (m News) NoRelatedSym() (field.NoRelatedSym, error) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//LinesOfText is a required field for News.
func (m News) LinesOfText() (field.LinesOfText, error) {
	var f field.LinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//URLLink is a non-required field for News.
func (m News) URLLink() (field.URLLink, error) {
	var f field.URLLink
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
