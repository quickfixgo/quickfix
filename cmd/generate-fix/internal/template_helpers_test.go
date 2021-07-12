package internal

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestRequiredFieldsRequiredComponent(t *testing.T) {
	ft1 := datadictionary.NewFieldType("type1", 11, "STRING")

	requiredfd1 := datadictionary.NewFieldDef(ft1, true)
	ct1 := datadictionary.NewComponentType("ct1", []datadictionary.MessagePart{requiredfd1})

	// when the component is optional the fields shouldn't be considered required
	{
		md := datadictionary.NewMessageDef("some message", "X", []datadictionary.MessagePart{
			datadictionary.NewComponent(ct1, false),
		})

		res := requiredFields(md)
		assert.Len(t, res, 0)
	}

	// when the component is required the fields should be considered required
	{
		md := datadictionary.NewMessageDef("some message", "X", []datadictionary.MessagePart{
			datadictionary.NewComponent(ct1, true),
		})

		res := requiredFields(md)
		assert.Len(t, res, 1)
	}
}
