//Package news msg type = B.
package news

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/fix/enum"
	"github.com/quickfixgo/quickfix/fix/field"
)

//Message is a News wrapper for the generic Message type
type Message struct {
	quickfix.Message
}

//OrigTime is a non-required field for News.
func (m Message) OrigTime() (*field.OrigTimeField, quickfix.MessageRejectError) {
	f := &field.OrigTimeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetOrigTime reads a OrigTime from News.
func (m Message) GetOrigTime(f *field.OrigTimeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Urgency is a non-required field for News.
func (m Message) Urgency() (*field.UrgencyField, quickfix.MessageRejectError) {
	f := &field.UrgencyField{}
	err := m.Body.Get(f)
	return f, err
}

//GetUrgency reads a Urgency from News.
func (m Message) GetUrgency(f *field.UrgencyField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//Headline is a required field for News.
func (m Message) Headline() (*field.HeadlineField, quickfix.MessageRejectError) {
	f := &field.HeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetHeadline reads a Headline from News.
func (m Message) GetHeadline(f *field.HeadlineField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadlineLen is a non-required field for News.
func (m Message) EncodedHeadlineLen() (*field.EncodedHeadlineLenField, quickfix.MessageRejectError) {
	f := &field.EncodedHeadlineLenField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadlineLen reads a EncodedHeadlineLen from News.
func (m Message) GetEncodedHeadlineLen(f *field.EncodedHeadlineLenField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//EncodedHeadline is a non-required field for News.
func (m Message) EncodedHeadline() (*field.EncodedHeadlineField, quickfix.MessageRejectError) {
	f := &field.EncodedHeadlineField{}
	err := m.Body.Get(f)
	return f, err
}

//GetEncodedHeadline reads a EncodedHeadline from News.
func (m Message) GetEncodedHeadline(f *field.EncodedHeadlineField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRoutingIDs is a non-required field for News.
func (m Message) NoRoutingIDs() (*field.NoRoutingIDsField, quickfix.MessageRejectError) {
	f := &field.NoRoutingIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRoutingIDs reads a NoRoutingIDs from News.
func (m Message) GetNoRoutingIDs(f *field.NoRoutingIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a non-required field for News.
func (m Message) NoRelatedSym() (*field.NoRelatedSymField, quickfix.MessageRejectError) {
	f := &field.NoRelatedSymField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from News.
func (m Message) GetNoRelatedSym(f *field.NoRelatedSymField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLegs is a non-required field for News.
func (m Message) NoLegs() (*field.NoLegsField, quickfix.MessageRejectError) {
	f := &field.NoLegsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLegs reads a NoLegs from News.
func (m Message) GetNoLegs(f *field.NoLegsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoUnderlyings is a non-required field for News.
func (m Message) NoUnderlyings() (*field.NoUnderlyingsField, quickfix.MessageRejectError) {
	f := &field.NoUnderlyingsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoUnderlyings reads a NoUnderlyings from News.
func (m Message) GetNoUnderlyings(f *field.NoUnderlyingsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoLinesOfText is a required field for News.
func (m Message) NoLinesOfText() (*field.NoLinesOfTextField, quickfix.MessageRejectError) {
	f := &field.NoLinesOfTextField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoLinesOfText reads a NoLinesOfText from News.
func (m Message) GetNoLinesOfText(f *field.NoLinesOfTextField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//URLLink is a non-required field for News.
func (m Message) URLLink() (*field.URLLinkField, quickfix.MessageRejectError) {
	f := &field.URLLinkField{}
	err := m.Body.Get(f)
	return f, err
}

//GetURLLink reads a URLLink from News.
func (m Message) GetURLLink(f *field.URLLinkField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawDataLength is a non-required field for News.
func (m Message) RawDataLength() (*field.RawDataLengthField, quickfix.MessageRejectError) {
	f := &field.RawDataLengthField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawDataLength reads a RawDataLength from News.
func (m Message) GetRawDataLength(f *field.RawDataLengthField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//RawData is a non-required field for News.
func (m Message) RawData() (*field.RawDataField, quickfix.MessageRejectError) {
	f := &field.RawDataField{}
	err := m.Body.Get(f)
	return f, err
}

//GetRawData reads a RawData from News.
func (m Message) GetRawData(f *field.RawDataField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplID is a non-required field for News.
func (m Message) ApplID() (*field.ApplIDField, quickfix.MessageRejectError) {
	f := &field.ApplIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplID reads a ApplID from News.
func (m Message) GetApplID(f *field.ApplIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplSeqNum is a non-required field for News.
func (m Message) ApplSeqNum() (*field.ApplSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplSeqNum reads a ApplSeqNum from News.
func (m Message) GetApplSeqNum(f *field.ApplSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplLastSeqNum is a non-required field for News.
func (m Message) ApplLastSeqNum() (*field.ApplLastSeqNumField, quickfix.MessageRejectError) {
	f := &field.ApplLastSeqNumField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplLastSeqNum reads a ApplLastSeqNum from News.
func (m Message) GetApplLastSeqNum(f *field.ApplLastSeqNumField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//ApplResendFlag is a non-required field for News.
func (m Message) ApplResendFlag() (*field.ApplResendFlagField, quickfix.MessageRejectError) {
	f := &field.ApplResendFlagField{}
	err := m.Body.Get(f)
	return f, err
}

//GetApplResendFlag reads a ApplResendFlag from News.
func (m Message) GetApplResendFlag(f *field.ApplResendFlagField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NewsID is a non-required field for News.
func (m Message) NewsID() (*field.NewsIDField, quickfix.MessageRejectError) {
	f := &field.NewsIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewsID reads a NewsID from News.
func (m Message) GetNewsID(f *field.NewsIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NoNewsRefIDs is a non-required field for News.
func (m Message) NoNewsRefIDs() (*field.NoNewsRefIDsField, quickfix.MessageRejectError) {
	f := &field.NoNewsRefIDsField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNoNewsRefIDs reads a NoNewsRefIDs from News.
func (m Message) GetNoNewsRefIDs(f *field.NoNewsRefIDsField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//NewsCategory is a non-required field for News.
func (m Message) NewsCategory() (*field.NewsCategoryField, quickfix.MessageRejectError) {
	f := &field.NewsCategoryField{}
	err := m.Body.Get(f)
	return f, err
}

//GetNewsCategory reads a NewsCategory from News.
func (m Message) GetNewsCategory(f *field.NewsCategoryField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//LanguageCode is a non-required field for News.
func (m Message) LanguageCode() (*field.LanguageCodeField, quickfix.MessageRejectError) {
	f := &field.LanguageCodeField{}
	err := m.Body.Get(f)
	return f, err
}

//GetLanguageCode reads a LanguageCode from News.
func (m Message) GetLanguageCode(f *field.LanguageCodeField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketID is a non-required field for News.
func (m Message) MarketID() (*field.MarketIDField, quickfix.MessageRejectError) {
	f := &field.MarketIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketID reads a MarketID from News.
func (m Message) GetMarketID(f *field.MarketIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//MarketSegmentID is a non-required field for News.
func (m Message) MarketSegmentID() (*field.MarketSegmentIDField, quickfix.MessageRejectError) {
	f := &field.MarketSegmentIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetMarketSegmentID reads a MarketSegmentID from News.
func (m Message) GetMarketSegmentID(f *field.MarketSegmentIDField) quickfix.MessageRejectError {
	return m.Body.Get(f)
}

//New returns an initialized Message with specified required fields for News.
func New(
	headline *field.HeadlineField,
	nolinesoftext *field.NoLinesOfTextField) Message {
	builder := Message{Message: quickfix.NewMessage()}
	builder.Header.Set(field.NewBeginString(enum.BeginStringFIXT11))
	builder.Header.Set(field.NewDefaultApplVerID(enum.ApplVerID_FIX50SP2))
	builder.Header.Set(field.NewMsgType("B"))
	builder.Body.Set(headline)
	builder.Body.Set(nolinesoftext)
	return builder
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(Message{msg}, sessionID)
	}
	return enum.ApplVerID_FIX50SP2, "B", r
}
