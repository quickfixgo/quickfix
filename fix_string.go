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

// FIXString is a FIX String Value, implements FieldValue.
type FIXString string

func (f FIXString) String() string {
	return string(f)
}

func (f *FIXString) Read(bytes []byte) (err error) {
	*f = FIXString(bytes)
	return
}

func (f FIXString) Write() []byte {
	return []byte(f)
}
