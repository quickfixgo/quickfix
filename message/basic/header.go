package basic

import (
	"github.com/cbusbey/quickfixgo/tag"
	"math"
)

//Header implements the MutableFieldMap interface and correctly orders fields as specified by FIX header order
type Header struct {
	FieldMap
}

func (header *Header) init() {
	header.FieldMap.init(headerFieldOrder)
}

//The first 3 fields in the message header must be 8,9,35
func headerFieldOrder(i, j tag.Tag) bool {
	var ordering = func(t tag.Tag) uint32 {
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
