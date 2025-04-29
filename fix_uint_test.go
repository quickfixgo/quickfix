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
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestFIXUInt_Write(t *testing.T) {
	field := FIXUint64(5)

	assert.Equal(t, "5", string(field.Write()))
}

func TestFIXUInt_Read(t *testing.T) {
	var field FIXUint64
	err := field.Read([]byte("15"))
	assert.Nil(t, err, "Unexpected error")
	assert.Equal(t, uint64(15), uint64(field))

	err = field.Read([]byte("blah"))
	assert.NotNil(t, err, "Unexpected error")
}

func TestFIXUInt_UInt(t *testing.T) {
	f := FIXUint64(4)
	assert.Equal(t, uint64(4), f.Uint64())
}

func BenchmarkFIXUInt_Read(b *testing.B) {
	intBytes := []byte("1500")
	var field FIXUint64

	for i := 0; i < b.N; i++ {
		_ = field.Read(intBytes)
	}
}
