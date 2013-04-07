package store

import (
	"github.com/cbusbey/quickfixgo/message/basic"
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type StoreTests struct {
	store MemoryStore
}

var _ = Suite(&StoreTests{})

func (s *StoreTests) SetUpTest(c *C) {
	s.store = NewMemoryStore()
}

func (s *StoreTests) TestGetMessages(c *C) {
	messages := s.store.GetMessages(1, 2)
	msg, ok := <-messages
	c.Assert(ok, Equals, false)

	buf1 := basic.Buffer("hello")
	s.store.SaveMessage(1, buf1)

	messages = s.store.GetMessages(1, 2)
	msg, ok = <-messages
	c.Assert(ok, Equals, true)
	c.Assert(msg.Bytes(), DeepEquals, buf1.Bytes())
	msg, ok = <-messages
	c.Assert(ok, Equals, false)

	buf2 := basic.Buffer("cruel")
	s.store.SaveMessage(2, buf2)

	buf3 := basic.Buffer("world")
	s.store.SaveMessage(3, buf3)

	messages = s.store.GetMessages(1, 2)
	msg, ok = <-messages
	c.Assert(ok, Equals, true)
	c.Assert(msg.Bytes(), DeepEquals, buf1.Bytes())

	msg, ok = <-messages
	c.Assert(ok, Equals, true)
	c.Assert(msg.Bytes(), DeepEquals, buf2.Bytes())

	msg, ok = <-messages
	c.Assert(ok, Equals, false)
}
