package fixt11

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
	testreqid *field.TestReqIDField) TestRequestBuilder {
	var builder TestRequestBuilder
	builder.MessageBuilder = message.Builder()
	builder.Header.Set(field.NewBeginString(fix.BeginString_FIXT11))
	builder.Header.Set(field.NewMsgType("1"))
	builder.Body.Set(testreqid)
	return builder
}

//TestReqID is a required field for TestRequest.
func (m TestRequest) TestReqID() (*field.TestReqIDField, errors.MessageRejectError) {
	f := &field.TestReqIDField{}
	err := m.Body.Get(f)
	return f, err
}

//GetTestReqID reads a TestReqID from TestRequest.
func (m TestRequest) GetTestReqID(f *field.TestReqIDField) errors.MessageRejectError {
	return m.Body.Get(f)
}
