package fix43

import (
	"github.com/quickfixgo/quickfix/errors"
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

//CreateRFQRequestBuilder returns an initialized RFQRequestBuilder with specified required fields.
func CreateRFQRequestBuilder(
	rfqreqid field.RFQReqID,
	norelatedsym field.NoRelatedSym) RFQRequestBuilder {
	var builder RFQRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(rfqreqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//RFQReqID is a required field for RFQRequest.
func (m RFQRequest) RFQReqID() (field.RFQReqID, errors.MessageRejectError) {
	var f field.RFQReqID
	err := m.Body.Get(&f)
	return f, err
}

//NoRelatedSym is a required field for RFQRequest.
func (m RFQRequest) NoRelatedSym() (field.NoRelatedSym, errors.MessageRejectError) {
	var f field.NoRelatedSym
	err := m.Body.Get(&f)
	return f, err
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m RFQRequest) SubscriptionRequestType() (field.SubscriptionRequestType, errors.MessageRejectError) {
	var f field.SubscriptionRequestType
	err := m.Body.Get(&f)
	return f, err
}
