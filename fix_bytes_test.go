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

func TestFIXBytesWrite(t *testing.T) {
	val := []byte("hello")
	field := FIXBytes(val)
	b := field.Write()

	assert.Equal(t, val, b)
}

func TestFIXBytesRead(t *testing.T) {
	val := []byte("world")
	var field FIXBytes
	assert.Nil(t, field.Read(val))
	assert.Equal(t, val, []byte(field))
}
