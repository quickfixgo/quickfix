package fix41

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
	headline field.Headline,
	linesoftext field.LinesOfText) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX41))
	builder.Header.Set(field.BuildMsgType("B"))
	builder.Body.Set(headline)
	builder.Body.Set(linesoftext)
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

//Headline is a required field for News.
func (m News) Headline() (*field.Headline, errors.MessageRejectError) {
	f := new(field.Headline)
	err := m.Body.Get(f)
	return f, err
}

//GetHeadline reads a Headline from News.
func (m News) GetHeadline(f *field.Headline) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for News.
func (m News) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from News.
func (m News) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
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

//URLLink is a non-required field for News.
func (m News) URLLink() (*field.URLLink, errors.MessageRejectError) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from News.
func (m News) GetURLLink(f *field.URLLink) errors.MessageRejectError {
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
