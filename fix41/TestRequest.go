package fix41

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/message"
)

//TestRequest msg type = 1.
type TestRequest struct {
	message.Message
}

//TestRequestBuilder builds TestRequest messages.
type TestRequestBuilder struct {
	message.MessageBuilder
}

//CreateTestRequestBuilder returns an initialized TestRequestBuilder with specified required fields.
func CreateTestRequestBuilder(
	testreqid field.TestReqID) TestRequestBuilder {
	var builder TestRequestBuilder
	builder.MessageBuilder = message.CreateMessageBuilder()
	builder.Body.Set(testreqid)
	return builder
}

//TestReqID is a required field for TestRequest.
func (m TestRequest) TestReqID() (field.TestReqID, errors.MessageRejectError) {
	var f field.TestReqID
	err := m.Body.Get(&f)
	return f, err
}
