package fix42

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type News struct {
	quickfix.Message
}

func (m *News) Headline() (*field.Headline, error) {
	f := new(field.Headline)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) EncodedHeadlineLen() (*field.EncodedHeadlineLen, error) {
	f := new(field.EncodedHeadlineLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoRoutingIDs() (*field.NoRoutingIDs, error) {
	f := new(field.NoRoutingIDs)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) URLLink() (*field.URLLink, error) {
	f := new(field.URLLink)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawDataLength() (*field.RawDataLength, error) {
	f := new(field.RawDataLength)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) OrigTime() (*field.OrigTime, error) {
	f := new(field.OrigTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) Urgency() (*field.Urgency, error) {
	f := new(field.Urgency)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) EncodedHeadline() (*field.EncodedHeadline, error) {
	f := new(field.EncodedHeadline)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) LinesOfText() (*field.LinesOfText, error) {
	f := new(field.LinesOfText)
	err := m.Body.Get(f)
	return f, err
}
func (m *News) RawData() (*field.RawData, error) {
	f := new(field.RawData)
	err := m.Body.Get(f)
	return f, err
}
