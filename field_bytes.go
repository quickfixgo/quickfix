package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

type fieldBytes struct {
	tag.Tag
	Data  []byte
	Value []byte
}

func newFieldBytes(tag tag.Tag, value []byte) *fieldBytes {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(tag)))
	buf.WriteString("=")
	buf.Write(value)
	buf.WriteString("")

	return &fieldBytes{Tag: tag, Data: buf.Bytes(), Value: value}
}

func parseField(rawFieldBytes []byte) (f *fieldBytes, err error) {
	sepIndex := bytes.IndexByte(rawFieldBytes, '=')

	if sepIndex == -1 {
		err = fmt.Errorf("fieldBytes.Parse: No '=' in '%s'", rawFieldBytes)
		return
	}

	parsedTag, err := strconv.Atoi(string(rawFieldBytes[:sepIndex]))

	if err != nil {
		err = fmt.Errorf("fieldBytes.Parse: %s", err.Error())
		return
	}

	f = new(fieldBytes)
	f.Tag = tag.Tag(parsedTag)
	f.Value = rawFieldBytes[(sepIndex + 1):(len(rawFieldBytes) - 1)]
	f.Data = rawFieldBytes

	return
}

func (f *fieldBytes) String() string {
	return string(f.Data)
}

func (f *fieldBytes) Total() int {
	total := 0

	for _, b := range []byte(f.Data) {
		total += int(b)
	}

	return total
}

func (f *fieldBytes) Length() int {
	return len(f.Data)
}
