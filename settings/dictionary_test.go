package settings

import (
	. "launchpad.net/gocheck"
	"testing"
)

func Test(t *testing.T) { TestingT(t) }

type DictionaryTests struct{ dict MutableDictionary }

var _ = Suite(&DictionaryTests{})

func (s *DictionaryTests) SetUpTest(c *C) {
	s.dict = NewDictionary()
}

func (s *DictionaryTests) TestStringSettings(c *C) {
	s.dict.SetString(BeginString, "foo")
	s.dict.SetString(SenderCompID, "blah")

	val, ok := s.dict.GetString(BeginString)
	c.Check(ok, Equals, true)
	c.Check(val, Equals, "foo")

	val, ok = s.dict.GetString(TargetCompID)
	c.Check(ok, Equals, false)
}

func (s *DictionaryTests) TestIntSettings(c *C) {
	s.dict.SetInt(SocketAcceptPort, 101)

	val, ok := s.dict.GetInt(SocketAcceptPort)
	c.Check(ok, Equals, true)
	c.Check(val, Equals, 101)

	val, ok = s.dict.GetInt(SocketConnectPort)
	c.Check(ok, Equals, false)
}

func (s *DictionaryTests) TestBoolSettings(c *C) {
	s.dict.SetBool(ResetOnLogon, true)

	val, ok := s.dict.GetBool(ResetOnLogon)
	c.Check(ok, Equals, true)
	c.Check(val, Equals, true)

	val, ok = s.dict.GetBool(Setting("RefreshOnLogon"))
	c.Check(ok, Equals, false)
}

func (s *DictionaryTests) testClone(c *C) {
	s.dict.SetInt(SocketAcceptPort, 101)
	s.dict.SetString(BeginString, "foo")
	s.dict.SetBool(ResetOnLogon, false)

	clone := CloneDictionary(s.dict)

	var intVal int
	var ok bool
	intVal, ok = clone.GetInt(SocketAcceptPort)
	c.Check(intVal, Equals, 101)
	c.Check(ok, Equals, true)

	var stringVal string
	stringVal, ok = clone.GetString(BeginString)
	c.Check(ok, Equals, true)
	c.Check(stringVal, Equals, "foo")

	var boolVal bool
	boolVal, ok = clone.GetBool(ResetOnLogon)
	c.Check(ok, Equals, true)
	c.Check(boolVal, Equals, false)
}

func (s *DictionaryTests) TestOverlay(c *C) {
	s.dict.SetInt(SocketAcceptPort, 101)
	s.dict.SetString(BeginString, "foo")

	overlay := NewDictionary()
	overlay.SetInt(SocketAcceptPort, 102)
	overlay.SetString(SenderCompID, "blah")

	s.dict.Overlay(overlay)

	if val, ok := s.dict.GetInt(SocketAcceptPort); ok {
		c.Check(val, Equals, 102)
	} else {
		c.Check(ok, Equals, true)
	}

	val, ok := s.dict.GetString(BeginString)
	c.Check(ok, Equals, true)
	c.Check(val, Equals, "foo")

	val, ok = s.dict.GetString(SenderCompID)
	c.Check(ok, Equals, true)
	c.Check(val, Equals, "blah")
}
