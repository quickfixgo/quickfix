package fix50sp1

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
	headline field.Headline,
	nolinesoftext field.NoLinesOfText) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(headline)
	builder.Body.Set(nolinesoftext)
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

//Headline is a required field for News.
func (m News) Headline() (field.Headline, errors.MessageRejectError) {
	var f field.Headline
	err := m.Body.Get(&f)
	return f, err
}

//EncodedHeadlineLen is a non-required field for News.
func (m News) EncodedHeadlineLen() (field.EncodedHeadlineLen, errors.MessageRejectError) {
	var f field.EncodedHeadlineLen
	err := m.Body.Get(&f)
	return f, err
}

//EncodedHeadline is a non-required field for News.
func (m News) EncodedHeadline() (field.EncodedHeadline, errors.MessageRejectError) {
	var f field.EncodedHeadline
	err := m.Body.Get(&f)
	return f, err
}

//NoRoutingIDs is a non-required field for News.
func (m News) NoRoutingIDs() (field.NoRoutingIDs, errors.MessageRejectError) {
	var f field.NoRoutingIDs
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a non-required field for News.
func (m News) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//NoLegs is a non-required field for News.
func (m News) NoLegs() (field.NoLegs, errors.MessageRejectError) {
	var f field.NoLegs
	err := m.Body.Get(&f)
	return f, err
}

//NoUnderlyings is a non-required field for News.
func (m News) NoUnderlyings() (field.NoUnderlyings, errors.MessageRejectError) {
	var f field.NoUnderlyings
	err := m.Body.Get(&f)
	return f, err
}

//NoLinesOfText is a required field for News.
func (m News) NoLinesOfText() (field.NoLinesOfText, errors.MessageRejectError) {
	var f field.NoLinesOfText
	err := m.Body.Get(&f)
	return f, err
}

//URLLink is a non-required field for News.
func (m News) URLLink() (field.URLLink, errors.MessageRejectError) {
	var f field.URLLink
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

//ApplID is a non-required field for News.
func (m News) ApplID() (field.ApplID, errors.MessageRejectError) {
	var f field.ApplID
	err := m.Body.Get(&f)
	return f, err
}

//ApplSeqNum is a non-required field for News.
func (m News) ApplSeqNum() (field.ApplSeqNum, errors.MessageRejectError) {
	var f field.ApplSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplLastSeqNum is a non-required field for News.
func (m News) ApplLastSeqNum() (field.ApplLastSeqNum, errors.MessageRejectError) {
	var f field.ApplLastSeqNum
	err := m.Body.Get(&f)
	return f, err
}

//ApplResendFlag is a non-required field for News.
func (m News) ApplResendFlag() (field.ApplResendFlag, errors.MessageRejectError) {
	var f field.ApplResendFlag
	err := m.Body.Get(&f)
	return f, err
}
