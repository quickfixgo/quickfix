package fix50sp2

import (
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//RFQRequest msg type = AH.
type RFQRequest struct {
	message.Message
}

//RFQRequestBuilder builds RFQRequest messages.
type RFQRequestBuilder struct {
	message.MessageBuilder
}

//NewRFQRequestBuilder returns an initialized RFQRequestBuilder with specified required fields.
func NewRFQRequestBuilder(
	rfqreqid field.RFQReqID,
	norelatedsym field.NoRelatedSym) *RFQRequestBuilder {
	builder := new(RFQRequestBuilder)
	builder.Body.Set(rfqreqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//RFQReqID is a required field for RFQRequest.
func (m *RFQRequest) RFQReqID() (*field.RFQReqID, error) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//NoRelatedSym is a required field for RFQRequest.
func (m *RFQRequest) NoRelatedSym() (*field.NoRelatedSym, error) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m *RFQRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, error) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//PrivateQuote is a non-required field for RFQRequest.
func (m *RFQRequest) PrivateQuote() (*field.PrivateQuote, error) {
	f := new(field.PrivateQuote)
	err := m.Body.Get(f)
	return f, err
}

//NoPartyIDs is a non-required field for RFQRequest.
func (m *RFQRequest) NoPartyIDs() (*field.NoPartyIDs, error) {
	f := new(field.NoPartyIDs)
	err := m.Body.Get(f)
	return f, err
}
