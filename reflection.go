package quickfix

import (
	"strconv"
	"strings"
)

func parseStructTag(goTag string) (tag Tag, omitEmpty bool, defaultVal *string) {
	tagParts := strings.Split(goTag, ",")
	if fixTag, err := strconv.Atoi(tagParts[0]); err != nil {
		panic(err)
	} else {
		tag = Tag(fixTag)
	}

	if len(tagParts) == 1 {
		return
	}

	if tagParts[1] == "omitempty" {
		omitEmpty = true
	} else {
		d := strings.TrimPrefix(tagParts[1], "default=")
		if d != tagParts[1] {
			defaultVal = &d
		}
	}

	return
}
