package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

type fieldBytes struct {
	Data  []byte
	Value []byte
}

func newFieldBytes(tag tag.Tag, value []byte) *fieldBytes {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(tag)))
	buf.WriteString("=")
	buf.Write(value)
	buf.WriteString("")

	f := new(fieldBytes)
	f.Data = buf.Bytes()
	f.Value = value

	return f
}

func parseField(rawFieldBytes []byte) (t tag.Tag, f *fieldBytes, err error) {
	sepIndex := bytes.IndexByte(rawFieldBytes, '=')

	if sepIndex == -1 {
		err = fmt.Errorf("fieldBytes.Parse: No '=' in '%s'", rawFieldBytes)
		return
	}

	parsed_tag, err := strconv.Atoi(string(rawFieldBytes[:sepIndex]))

	if err != nil {
		err = fmt.Errorf("fieldBytes.Parse: %s", err.Error())
		return
	}

	t = tag.Tag(parsed_tag)
	f = new(fieldBytes)
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
