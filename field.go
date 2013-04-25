package quickfixgo

import (
	"bytes"
	"fmt"
	"github.com/cbusbey/quickfixgo/tag"
	"strconv"
)

type field struct {
	Data  []byte
	Value []byte
}

func newField(tag tag.Tag, value []byte) *field {
	var buf bytes.Buffer
	buf.WriteString(strconv.Itoa(int(tag)))
	buf.WriteString("=")
	buf.Write(value)
	buf.WriteString("")

	f := new(field)
	f.Data = buf.Bytes()
	f.Value = value

	return f
}

func parseField(rawFieldBytes []byte) (t tag.Tag, f *field, err error) {
	sepIndex := bytes.IndexByte(rawFieldBytes, '=')

	if sepIndex == -1 {
		err = fmt.Errorf("field.Parse: No '=' in '%s'", rawFieldBytes)
		return
	}

	parsed_tag, err := strconv.Atoi(string(rawFieldBytes[:sepIndex]))

	if err != nil {
		err = fmt.Errorf("field.Parse: %s", err.Error())
		return
	}

	t = tag.Tag(parsed_tag)
	f = new(field)
	f.Value = rawFieldBytes[(sepIndex + 1):(len(rawFieldBytes) - 1)]
	f.Data = rawFieldBytes

	return
}

func (f *field) String() string {
	return string(f.Data)
}

func (f *field) Total() int {
	total := 0

	for _, b := range []byte(f.Data) {
		total += int(b)
	}

	return total
}

func (f *field) Length() int {
	return len(f.Data)
}
