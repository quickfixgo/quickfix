package fix43

import (
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

//NewTestRequestBuilder returns an initialized TestRequestBuilder with specified required fields.
func NewTestRequestBuilder(
	testreqid field.TestReqID) *TestRequestBuilder {
	builder := new(TestRequestBuilder)
	builder.Body.Set(testreqid)
	return builder
}

//TestReqID is a required field for TestRequest.
func (m *TestRequest) TestReqID() (*field.TestReqID, error) {
	f := new(field.TestReqID)
	err := m.Body.Get(f)
	return f, err
}
