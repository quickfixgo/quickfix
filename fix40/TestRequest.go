package fix40

import (
	"github.com/quickfixgo/quickfix/errors"
	"github.com/quickfixgo/quickfix/fix"
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
	builder.Header.Set(field.BuildBeginString(fix.BeginString_FIX40))
	builder.Header.Set(field.BuildMsgType("1"))
	builder.Body.Set(testreqid)
	return builder
}

//TestReqID is a required field for TestRequest.
func (m TestRequest) TestReqID() (*field.TestReqID, errors.MessageRejectError) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from TestRequest.
func (m TestRequest) GetTestReqID(f *field.TestReqID) errors.MessageRejectError {
	return m.Body.Get(f)
}
