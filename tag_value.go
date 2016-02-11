package quickfix

import (
	"bytes"
	"fmt"
	"strconv"
)

//tagValue is a low-level FIX field abstraction
type tagValue struct {
	Tag
	Value []byte
	bytes []byte
}

//tagValues is a slice of tagValue
type tagValues []tagValue

func (tv *tagValue) init(tag Tag, value []byte) {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(tag)))
	buf.WriteString("=")
	buf.Write(value)
	buf.WriteString("")

	tv.Tag = tag
	tv.bytes = buf.Bytes()
	tv.Value = value
}

func (tv *tagValue) parse(rawFieldBytes []byte) (err error) {
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

	tv.Tag = Tag(parsedTag)
	tv.Value = rawFieldBytes[(sepIndex + 1):(len(rawFieldBytes) - 1)]
	tv.bytes = rawFieldBytes

	return
}

func (tv tagValue) String() string {
	return string(tv.bytes)
}

func (tv tagValue) total() int {
	total := 0

	for _, b := range []byte(tv.bytes) {
		total += int(b)
	}

	return total
}

func (tv tagValue) length() int {
	return len(tv.bytes)
}
