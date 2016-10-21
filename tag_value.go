package quickfix

import (
	"bytes"
	"fmt"
	"strconv"
)

//TagValue is a low-level FIX field abstraction
type TagValue struct {
	tag   Tag
	value []byte
	bytes []byte
}

func (tv *TagValue) init(tag Tag, value []byte) {
	tv.bytes = strconv.AppendInt(nil, int64(tag), 10)
	tv.bytes = append(tv.bytes, []byte("=")...)
	tv.bytes = append(tv.bytes, value...)
	tv.bytes = append(tv.bytes, []byte("")...)

	tv.tag = tag
	tv.value = value
}

func (tv *TagValue) parse(rawFieldBytes []byte) (err error) {
	sepIndex := bytes.IndexByte(rawFieldBytes, '=')

	if sepIndex == -1 {
		err = fmt.Errorf("tagValue.Parse: No '=' in '%s'", rawFieldBytes)
		return
	}

	parsedTag, err := atoi(rawFieldBytes[:sepIndex])

	if err != nil {
		err = fmt.Errorf("tagValue.Parse: %s", err.Error())
		return
	}

	tv.tag = Tag(parsedTag)
	n := len(rawFieldBytes)
	tv.value = rawFieldBytes[(sepIndex + 1):(n - 1):(n - 1)]
	tv.bytes = rawFieldBytes[:n:n]

	return
}

func (tv TagValue) String() string {
	return string(tv.bytes)
}

func (tv TagValue) total() int {
	total := 0

	for _, b := range []byte(tv.bytes) {
		total += int(b)
	}

	return total
}

func (tv TagValue) length() int {
	return len(tv.bytes)
}
