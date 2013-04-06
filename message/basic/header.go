package basic

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
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
func headerFieldOrder(i, j message.Tag) bool {
	var ordering = func(t message.Tag) uint32 {
		switch t {
		case fix.BeginString:
			return 1
		case fix.BodyLength:
			return 2
		case fix.MsgType:
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
