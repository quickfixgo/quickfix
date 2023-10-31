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

// FieldValueWriter is an interface for writing field values.
type FieldValueWriter interface {
	// Write writes out the contents of the FieldValue to a `[]byte`.
	Write() []byte
}

// FieldValueReader is an interface for reading field values.
type FieldValueReader interface {
	// Read reads the contents of the `[]byte` into FieldValue.
	// Returns an error if there are issues in the data processing.
	Read([]byte) error
}

// The FieldValue interface is used to write/extract typed field values to/from raw bytes.
type FieldValue interface {
	FieldValueWriter
	FieldValueReader
}

// FieldWriter is an interface for a writing a field.
type FieldWriter interface {
	Tag() Tag
	FieldValueWriter
}

// Field is the interface implemented by all typed Fields in a Message.
type Field interface {
	FieldWriter
	FieldValueReader
}

// FieldGroupWriter is an interface for writing a FieldGroup.
type FieldGroupWriter interface {
	Tag() Tag
	Write() []TagValue
}

// FieldGroupReader is an interface for reading a FieldGroup.
type FieldGroupReader interface {
	Tag() Tag
	Read([]TagValue) ([]TagValue, error)
}

// FieldGroup is the interface implemented by all typed Groups in a Message.
type FieldGroup interface {
	Tag() Tag
	Write() []TagValue
	Read([]TagValue) ([]TagValue, error)
}
