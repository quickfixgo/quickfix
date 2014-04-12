package message

import (
	"github.com/quickfixgo/quickfix/fix"
)

type tagContainer struct {
	tag fix.Tag
}

func (c tagContainer) Tag() fix.Tag {
	return c.tag
}
