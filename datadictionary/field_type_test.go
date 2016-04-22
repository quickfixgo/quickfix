package datadictionary_test

import (
	"testing"

	"github.com/quickfixgo/quickfix/datadictionary"
	"github.com/stretchr/testify/assert"
)

func TestNewFieldType(t *testing.T) {
	ft := datadictionary.NewFieldType("myname", 10, "STRING")
	assert.NotNil(t, ft)
	assert.Equal(t, "myname", ft.Name())
	assert.Equal(t, 10, ft.Tag())
	assert.Equal(t, "STRING", ft.Type)
}
