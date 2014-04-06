package field

import (
	"github.com/quickfixgo/quickfix/tag"
)

type tagContainer struct {
	tag tag.Tag
}

func (c tagContainer) Tag() tag.Tag {
	return c.tag
}
