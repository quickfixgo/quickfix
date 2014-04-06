package fix50sp1

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
)

type MarketDefinitionRequest struct {
	quickfix.Message
}

func (m *MarketDefinitionRequest) MarketReqID() (*field.MarketReqID, error) {
	f := new(field.MarketReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionRequest) MarketID() (*field.MarketID, error) {
	f := new(field.MarketID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionRequest) MarketSegmentID() (*field.MarketSegmentID, error) {
	f := new(field.MarketSegmentID)
	err := m.Body.Get(f)
	return f, err
}
func (m *MarketDefinitionRequest) ParentMktSegmID() (*field.ParentMktSegmID, error) {
	f := new(field.ParentMktSegmID)
	err := m.Body.Get(f)
	return f, err
}
