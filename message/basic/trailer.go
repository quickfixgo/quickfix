package basic

import (
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/message"
)

//Message Trailer type. TrailerOrder FieldOrder.
type Trailer struct {
	FieldMap
}

//CheckSum is a required field of the trailer
func (t *Trailer) setCheckSum(checkSum message.StringField) {
	t.Set(checkSum)
}

//Must be called before use
func (t *Trailer) init() {
	t.FieldMap.init(trailerFieldOrder)
}

//Field 10 must be last in the trailer
func trailerFieldOrder(i, j message.Tag) bool {
	switch {
	case i == fix.CheckSum:
		return false
	case j == fix.CheckSum:
		return true
	}

	return i < j
}
