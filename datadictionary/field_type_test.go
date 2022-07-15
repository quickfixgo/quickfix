package datadictionary_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/quickfixgo/quickfix/datadictionary"
)

func TestNewFieldType(t *testing.T) {
	ft := datadictionary.NewFieldType("myname", 10, "STRING")
	assert.NotNil(t, ft)
	assert.Equal(t, "myname", ft.Name())
	assert.Equal(t, 10, ft.Tag())
	assert.Equal(t, "STRING", ft.Type)
}
