package fix50sp2

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type News struct {
	quickfix.Message
}

func (m *News) Urgency() (*field.Urgency, error) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) EncodedHeadlineLen() (*field.EncodedHeadlineLen, error) {
	f := new(field.EncodedHeadlineLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoUnderlyings() (*field.NoUnderlyings, error) {
	f := new(field.NoUnderlyings)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoNewsRefIDs() (*field.NoNewsRefIDs, error) {
	f := new(field.NoNewsRefIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NewsCategory() (*field.NewsCategory, error) {
	f := new(field.NewsCategory)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoLegs() (*field.NoLegs, error) {
	f := new(field.NoLegs)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) Headline() (*field.Headline, error) {
	f := new(field.Headline)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) ApplID() (*field.ApplID, error) {
	f := new(field.ApplID)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) ApplLastSeqNum() (*field.ApplLastSeqNum, error) {
	f := new(field.ApplLastSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) ApplResendFlag() (*field.ApplResendFlag, error) {
	f := new(field.ApplResendFlag)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NewsID() (*field.NewsID, error) {
	f := new(field.NewsID)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) EncodedHeadline() (*field.EncodedHeadline, error) {
	f := new(field.EncodedHeadline)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoLinesOfText() (*field.NoLinesOfText, error) {
	f := new(field.NoLinesOfText)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) ApplSeqNum() (*field.ApplSeqNum, error) {
	f := new(field.ApplSeqNum)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) LanguageCode() (*field.LanguageCode, error) {
	f := new(field.LanguageCode)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
