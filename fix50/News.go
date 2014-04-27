package fix50

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

import (
	"github.com/quickfixgo/quickfix/fix/enum"
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
	nolinesoftext field.NoLinesOfText) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("B"))
	builder.Body.Set(headline)
	builder.Body.Set(nolinesoftext)
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

//EncodedHeadlineLen is a non-required field for News.
func (m News) EncodedHeadlineLen() (*field.EncodedHeadlineLen, errors.MessageRejectError) {
	f := new(field.EncodedHeadlineLen)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadlineLen reads a EncodedHeadlineLen from News.
func (m News) GetEncodedHeadlineLen(f *field.EncodedHeadlineLen) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadline is a non-required field for News.
func (m News) EncodedHeadline() (*field.EncodedHeadline, errors.MessageRejectError) {
	f := new(field.EncodedHeadline)
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadline reads a EncodedHeadline from News.
func (m News) GetEncodedHeadline(f *field.EncodedHeadline) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for News.
func (m News) NoRoutingIDs() (*field.NoRoutingIDs, errors.MessageRejectError) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from News.
func (m News) GetNoRoutingIDs(f *field.NoRoutingIDs) errors.MessageRejectError {
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

//NoLegs is a non-required field for News.
func (m News) NoLegs() (*field.NoLegs, errors.MessageRejectError) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from News.
func (m News) GetNoLegs(f *field.NoLegs) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for News.
func (m News) NoUnderlyings() (*field.NoUnderlyings, errors.MessageRejectError) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from News.
func (m News) GetNoUnderlyings(f *field.NoUnderlyings) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLinesOfText is a required field for News.
func (m News) NoLinesOfText() (*field.NoLinesOfText, errors.MessageRejectError) {
	f := new(field.NoLinesOfText)
	err := m.Body.Get(f)
	return f, err
}

//GetNoLinesOfText reads a NoLinesOfText from News.
func (m News) GetNoLinesOfText(f *field.NoLinesOfText) errors.MessageRejectError {
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
