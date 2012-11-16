package quickfixgo

import(
    . "launchpad.net/gocheck"
  "testing"
)

func Test(t *testing.T) {TestingT(t)}

var _ = Suite(&IntFieldTests{})
var _ = Suite(&StringFieldTests{})

type IntFieldTests struct {field IntField}
type StringFieldTests struct {field StringField}

func (s *IntFieldTests) SetUpTest(c *C) {
  s.field=NewIntField(tag.Tag(1),5)
}

func (s *IntFieldTests) TestNewField(c *C) {
  c.Check(s.field.Tag(), Equals, tag.Tag(1))
  c.Check(s.field.Value(), Equals, "5")
  c.Check(s.field.IntValue(), Equals, 5)
}

func (s *StringFieldTests) SetUpTest(c *C) {
  s.field=NewStringField(1,"CWB")
}

func (s *StringFieldTests) TestNewField(c *C) {
  c.Check(s.field.Tag(), Equals, tag.Tag(1))
  c.Check(s.field.Value(), Equals, "CWB")
}
