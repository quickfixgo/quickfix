// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"math"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIXInt_Write(t *testing.T) {
	field := FIXInt(5)

	assert.Equal(t, "5", string(field.Write()))
}

func TestFIXInt_Read(t *testing.T) {
	var field FIXInt
	err := field.Read([]byte("15"))
	assert.Nil(t, err, "Unexpected error")
	assert.Equal(t, 15, int(field))

	err = field.Read([]byte("blah"))
	assert.NotNil(t, err, "Unexpected error")
}

func TestFIXInt_ReadRange(t *testing.T) {
	var field FIXInt
	assert.Nil(t, field.Read([]byte(strconv.Itoa(math.MaxInt))))
	assert.Equal(t, math.MaxInt, int(field))

	assert.Nil(t, field.Read([]byte(strconv.Itoa(-math.MaxInt))))
	assert.Equal(t, -math.MaxInt, int(field))

	assert.Nil(t, field.Read([]byte("0000000000000000000000015")))
	assert.Equal(t, 15, int(field))

	for _, v := range []string{"9223372036854775808", "-9223372036854775809", "99999999999999999999"} {
		assert.NotNil(t, field.Read([]byte(v)), v)
	}
}

func TestFIXInt_Int(t *testing.T) {
	f := FIXInt(4)
	assert.Equal(t, 4, f.Int())
}

func BenchmarkFIXInt_Read(b *testing.B) {
	intBytes := []byte("1500")
	var field FIXInt

	for i := 0; i < b.N; i++ {
		_ = field.Read(intBytes)
	}
}
