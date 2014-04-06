package quickfix

import (
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/tag"
)

//Trailer is a collection of fields representing the trailer of a FIX message.
type Trailer struct {
	FieldMapBuilder
}

//CheckSum is a required field of the trailer.
func (t *Trailer) setCheckSum(checkSum *field.StringField) {
	t.Set(checkSum)
}

//Must be called before use
func (t *Trailer) init() {
	t.FieldMapBuilder.init(trailerFieldOrder)
}

//Field 10 must be last in the trailer
func trailerFieldOrder(i, j tag.Tag) bool {
	switch {
	case i == tag.CheckSum:
		return false
	case j == tag.CheckSum:
		return true
	}

	return i < j
}
