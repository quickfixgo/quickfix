package datadictionary_test

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestNewGroupField(t *testing.T) {
	ft := datadictionary.NewFieldType("aname", 11, "INT")
	fg := datadictionary.NewGroupFieldDef(ft, true, []datadictionary.MessagePart{})
	assert.NotNil(t, fg)
	assert.Equal(t, "aname", fg.Name())
	assert.Equal(t, true, fg.Required())
}
