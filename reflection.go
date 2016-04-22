package quickfix

import (
	"reflect"
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

func buildGroupTemplate(elem reflect.Type) GroupTemplate {
	template := make(GroupTemplate, 0)

	var walkFunc func(t reflect.Type)

	walkFunc = func(t reflect.Type) {
		for i := 0; i < t.NumField(); i++ {
			sf := t.Field(i)

			elementType := sf.Type
			if elementType.Kind() == reflect.Ptr {
				elementType = elementType.Elem()
			}

			// recurse if item is a component, optional or not
			structTag := sf.Tag.Get("fix")
			if structTag == "" && elementType.Kind() == reflect.Struct {
				walkFunc(elementType)
				continue
			}

			afixTag, _, _ := parseStructTag(structTag)
			template = append(template, GroupElement(Tag(afixTag)))
		}
	}

	walkFunc(elem)
	return template
}
