package datadictionary_test

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestNewMessageDef(t *testing.T) {

	ft1 := datadictionary.NewFieldType("type1", 11, "STRING")
	ft2 := datadictionary.NewFieldType("type2", 12, "STRING")
	ft3 := datadictionary.NewFieldType("type3", 13, "INT")

	optionalfd1 := datadictionary.NewFieldDef(ft1, false)
	requiredfd1 := datadictionary.NewFieldDef(ft1, true)

	optionalfd2 := datadictionary.NewFieldDef(ft2, false)
	//	requiredfd2 := datadictionary.NewFieldDef(ft2, true)

	optionalGroup1 := datadictionary.NewGroupFieldDef(ft3, false, []datadictionary.MessagePart{requiredfd1, optionalfd2})
	requiredGroup1 := datadictionary.NewGroupFieldDef(ft3, true, []datadictionary.MessagePart{requiredfd1, optionalfd2})

	ct1 := datadictionary.NewComponentType("ct1", []datadictionary.MessagePart{requiredGroup1})

	optionalComp1 := datadictionary.NewComponent(ct1, false)

	var tests = []struct {
		parts                 []datadictionary.MessagePart
		expectedTags          datadictionary.TagSet
		expectedRequiredTags  datadictionary.TagSet
		expectedRequiredParts []datadictionary.MessagePart
	}{
		{
			parts:                 []datadictionary.MessagePart{},
			expectedTags:          datadictionary.TagSet{},
			expectedRequiredTags:  datadictionary.TagSet{},
			expectedRequiredParts: []datadictionary.MessagePart(nil),
		},
		{
			parts:                 []datadictionary.MessagePart{optionalfd1},
			expectedTags:          datadictionary.TagSet{11: struct{}{}},
			expectedRequiredTags:  datadictionary.TagSet{},
			expectedRequiredParts: []datadictionary.MessagePart(nil),
		},
		{
			parts:                 []datadictionary.MessagePart{requiredfd1, optionalfd2},
			expectedTags:          datadictionary.TagSet{11: struct{}{}, 12: struct{}{}},
			expectedRequiredTags:  datadictionary.TagSet{11: struct{}{}},
			expectedRequiredParts: []datadictionary.MessagePart{requiredfd1},
		},
		{
			parts:                 []datadictionary.MessagePart{optionalGroup1},
			expectedTags:          datadictionary.TagSet{11: struct{}{}, 12: struct{}{}, 13: struct{}{}},
			expectedRequiredTags:  datadictionary.TagSet{},
			expectedRequiredParts: []datadictionary.MessagePart(nil),
		},
		{
			parts:                 []datadictionary.MessagePart{requiredGroup1},
			expectedTags:          datadictionary.TagSet{11: struct{}{}, 12: struct{}{}, 13: struct{}{}},
			expectedRequiredTags:  datadictionary.TagSet{13: struct{}{}},
			expectedRequiredParts: []datadictionary.MessagePart{requiredGroup1},
		},
		{
			parts:                 []datadictionary.MessagePart{optionalComp1},
			expectedTags:          datadictionary.TagSet{11: struct{}{}, 12: struct{}{}, 13: struct{}{}},
			expectedRequiredTags:  datadictionary.TagSet{},
			expectedRequiredParts: []datadictionary.MessagePart(nil),
		},
	}

	for _, test := range tests {
		md := datadictionary.NewMessageDef("some message", "X", test.parts)

		assert.NotNil(t, md)
		assert.Equal(t, "some message", md.Name)
		assert.Equal(t, "X", md.MsgType)
		assert.Equal(t, test.expectedTags, md.Tags)
		assert.Equal(t, test.expectedRequiredTags, md.RequiredTags)
		assert.Equal(t, test.parts, md.Parts)
		assert.Equal(t, test.expectedRequiredParts, md.RequiredParts())
	}
}
