package quickfix

import (
	"bytes"
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"testing"
)

var builder MessageBuilder

func TestMessageBuilder_checkBuild(t *testing.T) {
	builder = NewMessageBuilder()

	builder.Header().Set(field.NewBeginString(fix.BeginString_FIX44))
	builder.Header().Set(fix.NewStringField(tag.MsgType, "A"))
	builder.Header().Set(fix.NewStringField(tag.SendingTime, "20140615-19:49:56"))

	builder.Body().Set(field.NewUsername("my_user"))
	builder.Body().Set(field.NewPassword("secret"))

	expectedBytes := []byte("8=FIX.4.49=4935=A52=20140615-19:49:56553=my_user554=secret10=072")
	result, err := builder.Build()
	if err != nil {
		t.Error("Unexpected error", err)
	}

	if !bytes.Equal(expectedBytes, result) {
		t.Error("Unexpected bytes, got ", string(result))
	}

}
