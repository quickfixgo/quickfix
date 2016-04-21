package datadictionary_test

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestNewFieldDef(t *testing.T) {
	ft := datadictionary.NewFieldType("aname", 11, "INT")

	var tests = []struct {
		required bool
	}{
		{false},
	}

	for _, test := range tests {
		fd := datadictionary.NewFieldDef(ft, test.required)
		assert.False(t, fd.IsGroup(), "field def is not a group")
		assert.Equal(t, "aname", fd.Name())
		assert.Equal(t, test.required, fd.Required())
	}
}
