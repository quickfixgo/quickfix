package fix50sp2

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type SecurityListRequest struct {
	quickfixgo.Message
}

func (m *SecurityListRequest) SecurityReqID() (*field.SecurityReqID, error) {
	f := new(field.SecurityReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) SecurityListRequestType() (*field.SecurityListRequestType, error) {
	f := new(field.SecurityListRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) SecurityListID() (*field.SecurityListID, error) {
	f := new(field.SecurityListID)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) SecurityListType() (*field.SecurityListType, error) {
	f := new(field.SecurityListType)
	err := m.Body.Get(f)
	return f, err
}
func (m *SecurityListRequest) SecurityListTypeSource() (*field.SecurityListTypeSource, error) {
	f := new(field.SecurityListTypeSource)
	err := m.Body.Get(f)
	return f, err
}
