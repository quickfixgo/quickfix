package quickfix

import (
	"bytes"
	"fmt"
	"strconv"
)

//TagValue is a low-level FIX field abstraction
type TagValue struct {
	Tag
	Value []byte
	bytes []byte
}

//TagValues is a slice of TagValue
type TagValues []TagValue

func (tv *TagValue) init(tag Tag, value []byte) {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(tag)))
	buf.WriteString("=")
	buf.Write(value)
	buf.WriteString("")

	tv.Tag = tag
	tv.bytes = buf.Bytes()
	tv.Value = value
}

func (tv *TagValue) parse(rawFieldBytes []byte) (err error) {
	sepIndex := bytes.IndexByte(rawFieldBytes, '=')

	if sepIndex == -1 {
		err = fmt.Errorf("TagValue.Parse: No '=' in '%s'", rawFieldBytes)
		return
	}

	parsedTag, err := atoi(rawFieldBytes[:sepIndex])

	if err != nil {
		err = fmt.Errorf("TagValue.Parse: %s", err.Error())
		return
	}

	tv.Tag = Tag(parsedTag)
	tv.Value = rawFieldBytes[(sepIndex + 1):(len(rawFieldBytes) - 1)]
	tv.bytes = rawFieldBytes

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
