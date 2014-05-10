package fix50sp1

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
	headline *field.HeadlineField,
	nolinesoftext *field.NoLinesOfTextField) NewsBuilder {
	var builder NewsBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP1))
	builder.Header.Set(field.NewMsgType("B"))
	builder.Body.Set(headline)
	builder.Body.Set(nolinesoftext)
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

//Headline is a required field for News.
func (m News) Headline() (*field.HeadlineField, errors.MessageRejectError) {
	f := &field.HeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeadline reads a Headline from News.
func (m News) GetHeadline(f *field.HeadlineField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadlineLen is a non-required field for News.
func (m News) EncodedHeadlineLen() (*field.EncodedHeadlineLenField, errors.MessageRejectError) {
	f := &field.EncodedHeadlineLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadlineLen reads a EncodedHeadlineLen from News.
func (m News) GetEncodedHeadlineLen(f *field.EncodedHeadlineLenField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadline is a non-required field for News.
func (m News) EncodedHeadline() (*field.EncodedHeadlineField, errors.MessageRejectError) {
	f := &field.EncodedHeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadline reads a EncodedHeadline from News.
func (m News) GetEncodedHeadline(f *field.EncodedHeadlineField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for News.
func (m News) NoRoutingIDs() (*field.NoRoutingIDsField, errors.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from News.
func (m News) GetNoRoutingIDs(f *field.NoRoutingIDsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for News.
func (m News) NoRelatedSym() (*field.NoRelatedSymField, errors.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from News.
func (m News) GetNoRelatedSym(f *field.NoRelatedSymField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for News.
func (m News) NoLegs() (*field.NoLegsField, errors.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from News.
func (m News) GetNoLegs(f *field.NoLegsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for News.
func (m News) NoUnderlyings() (*field.NoUnderlyingsField, errors.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from News.
func (m News) GetNoUnderlyings(f *field.NoUnderlyingsField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoLinesOfText is a required field for News.
func (m News) NoLinesOfText() (*field.NoLinesOfTextField, errors.MessageRejectError) {
	f := &field.NoLinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLinesOfText reads a NoLinesOfText from News.
func (m News) GetNoLinesOfText(f *field.NoLinesOfTextField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for News.
func (m News) URLLink() (*field.URLLinkField, errors.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from News.
func (m News) GetURLLink(f *field.URLLinkField) errors.MessageRejectError {
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

//ApplID is a non-required field for News.
func (m News) ApplID() (*field.ApplIDField, errors.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from News.
func (m News) GetApplID(f *field.ApplIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for News.
func (m News) ApplSeqNum() (*field.ApplSeqNumField, errors.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from News.
func (m News) GetApplSeqNum(f *field.ApplSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for News.
func (m News) ApplLastSeqNum() (*field.ApplLastSeqNumField, errors.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from News.
func (m News) GetApplLastSeqNum(f *field.ApplLastSeqNumField) errors.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for News.
func (m News) ApplResendFlag() (*field.ApplResendFlagField, errors.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from News.
func (m News) GetApplResendFlag(f *field.ApplResendFlagField) errors.MessageRejectError {
	return m.Body.Get(f)
}
