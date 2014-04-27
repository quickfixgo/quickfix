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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.BuildDefaultApplVerID(enum.ApplVerID_FIX50))
	builder.Header.Set(field.BuildMsgType("AH"))
	builder.Body.Set(rfqreqid)
	builder.Body.Set(norelatedsym)
	return builder
}

//RFQReqID is a required field for RFQRequest.
func (m RFQRequest) RFQReqID() (*field.RFQReqID, errors.MessageRejectError) {
	f := new(field.RFQReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetRFQReqID reads a RFQReqID from RFQRequest.
func (m RFQRequest) GetRFQReqID(f *field.RFQReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}

//NoRelatedSym is a required field for RFQRequest.
func (m RFQRequest) NoRelatedSym() (*field.NoRelatedSym, errors.MessageRejectError) {
	f := new(field.NoRelatedSym)
	err := m.Body.Get(f)
	return f, err
}

//GetNoRelatedSym reads a NoRelatedSym from RFQRequest.
func (m RFQRequest) GetNoRelatedSym(f *field.NoRelatedSym) errors.MessageRejectError {
	return m.Body.Get(f)
}

//SubscriptionRequestType is a non-required field for RFQRequest.
func (m RFQRequest) SubscriptionRequestType() (*field.SubscriptionRequestType, errors.MessageRejectError) {
	f := new(field.SubscriptionRequestType)
	err := m.Body.Get(f)
	return f, err
}

//GetSubscriptionRequestType reads a SubscriptionRequestType from RFQRequest.
func (m RFQRequest) GetSubscriptionRequestType(f *field.SubscriptionRequestType) errors.MessageRejectError {
	return m.Body.Get(f)
}
