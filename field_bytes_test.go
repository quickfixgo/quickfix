package quickfixgo

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

var _ = Suite(&FieldBytesTests{})

type FieldBytesTests struct{}

func (s *FieldBytesTests) TestNewField(c *C) {
	f := newFieldBytes(tag.Tag(8), []byte("blahblah"))
	c.Check(f, NotNil)
	c.Check(f.Data, DeepEquals, []byte("8=blahblah"))
	c.Check(f.Value, DeepEquals, []byte("blahblah"))
}

func (s *FieldBytesTests) TestParseField(c *C) {
	stringField := "8=FIX.4.0"

	t, f, err := parseField([]byte(stringField))
	c.Check(err, IsNil)
	c.Check(t, Equals, tag.Tag(8))
	c.Check(f, NotNil)
	c.Check(f.Data, DeepEquals, []byte(stringField))
	c.Check(f.Value, DeepEquals, []byte("FIX.4.0"))
}

func (s *FieldBytesTests) TestParseFieldFail(c *C) {
	stringField := "not_tag_equal_value"

	_, _, err := parseField([]byte(stringField))
	c.Check(err, NotNil)

	stringField = "tag_not_an_int=uhoh"
	_, _, err = parseField([]byte(stringField))
	c.Check(err, NotNil)
}

func (s *FieldBytesTests) TestString(c *C) {
	stringField := "8=FIX.4.0"
	_, f, _ := parseField([]byte(stringField))

	c.Check(f.String(), Equals, "8=FIX.4.0")
}

func (s *FieldBytesTests) TestLength(c *C) {
	stringField := "8=FIX.4.0"
	_, f, _ := parseField([]byte(stringField))

	c.Check(f.Length(), Equals, len(stringField))
}

func (s *FieldBytesTests) TestTotal(c *C) {
	stringField := "1=hello"
	_, f, _ := parseField([]byte(stringField))
	c.Check(f.Total(), Equals, 643, Commentf("Total is the summation of the ascii byte values of the field string"))
}
