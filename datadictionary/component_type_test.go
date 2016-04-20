package datadictionary_test

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestNewComponentType(t *testing.T) {
	ft1 := datadictionary.NewFieldType("aname1", 11, "INT")
	ft2 := datadictionary.NewFieldType("aname2", 12, "INT")

	optionalField1 := datadictionary.NewFieldDef(ft1, false)
	requiredField1 := datadictionary.NewFieldDef(ft1, true)

	optionalField2 := datadictionary.NewFieldDef(ft2, false)
	requiredField2 := datadictionary.NewFieldDef(ft2, true)

	requiredComp1 := datadictionary.NewComponent(
		datadictionary.NewComponentType("comp1", []datadictionary.MessagePart{requiredField1}),
		true)

	optionalComp1 := datadictionary.NewComponent(
		datadictionary.NewComponentType("comp1", []datadictionary.MessagePart{requiredField1}),
		false)

	var tests = []struct {
		testName               string
		parts                  []datadictionary.MessagePart
		expectedFields         []*datadictionary.FieldDef
		expectedRequiredParts  []datadictionary.MessagePart
		expectedRequiredFields []*datadictionary.FieldDef
	}{
		{
			testName:       "test1",
			parts:          []datadictionary.MessagePart{optionalField1},
			expectedFields: []*datadictionary.FieldDef{optionalField1},
		},
		{
			testName:               "test2",
			parts:                  []datadictionary.MessagePart{requiredField1},
			expectedFields:         []*datadictionary.FieldDef{requiredField1},
			expectedRequiredFields: []*datadictionary.FieldDef{requiredField1},
			expectedRequiredParts:  []datadictionary.MessagePart{requiredField1},
		},
		{
			testName:               "test3",
			parts:                  []datadictionary.MessagePart{requiredField1, optionalField2},
			expectedFields:         []*datadictionary.FieldDef{requiredField1, optionalField2},
			expectedRequiredFields: []*datadictionary.FieldDef{requiredField1},
			expectedRequiredParts:  []datadictionary.MessagePart{requiredField1},
		},
		{
			testName:               "test4",
			parts:                  []datadictionary.MessagePart{requiredField2, optionalComp1},
			expectedFields:         []*datadictionary.FieldDef{requiredField2, requiredField1},
			expectedRequiredFields: []*datadictionary.FieldDef{requiredField2},
			expectedRequiredParts:  []datadictionary.MessagePart{requiredField2},
		},
		{
			testName:               "test5",
			parts:                  []datadictionary.MessagePart{requiredField2, requiredComp1},
			expectedFields:         []*datadictionary.FieldDef{requiredField2, requiredField1},
			expectedRequiredFields: []*datadictionary.FieldDef{requiredField2, requiredField1},
			expectedRequiredParts:  []datadictionary.MessagePart{requiredField2, requiredComp1},
		},
	}

	for _, test := range tests {
		ct := datadictionary.NewComponentType("cname", test.parts)
		assert.NotNil(t, ct, test.testName)
		assert.Equal(t, "cname", ct.Name(), test.testName)
		assert.Equal(t, test.expectedFields, ct.Fields(), test.testName)
		assert.Equal(t, test.parts, ct.Parts(), test.testName)
		assert.Equal(t, test.expectedRequiredFields, ct.RequiredFields(), test.testName)
		assert.Equal(t, test.expectedRequiredParts, ct.RequiredParts(), test.testName)
	}
}
