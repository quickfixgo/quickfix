package message

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/tag"
	"math"
)

//Header is a collection of fields representing the header of a FIX message.
type Header struct {
	FieldMapBuilder
}

func (header *Header) init() {
	header.FieldMapBuilder.init(headerFieldOrder)
}

//The first 3 fields in the message header must be 8,9,35
func headerFieldOrder(i, j fix.Tag) bool {
	var ordering = func(t fix.Tag) uint32 {
		switch t {
		case tag.BeginString:
			return 1
		case tag.BodyLength:
			return 2
		case tag.MsgType:
			return 3
		}

		return math.MaxUint32
	}

	orderi := ordering(i)
	orderj := ordering(j)

	switch {
	case orderi < orderj:
		return true
	case orderi > orderj:
		return false
	}

	return i < j
}
