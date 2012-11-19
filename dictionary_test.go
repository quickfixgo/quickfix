package quickfixgo

import(
    . "launchpad.net/gocheck"
  "testing"
)

func Test(t *testing.T) {TestingT(t)}

type DictionaryTests struct {dict Dictionary}

var _ = Suite(&DictionaryTests{})

func (s *DictionaryTests) SetUpTest(c *C) {
  s.dict=NewDictionary()
}

func (s * DictionaryTests) TestStringSettings(c *C) {
  s.dict.SetString(BeginString, "foo")
  s.dict.SetString(SenderCompID, "blah")

  val,ok:=s.dict.GetString(BeginString)
  c.Check(ok, Equals, true)
  c.Check(val, Equals, "foo")

  val,ok=s.dict.GetString(TargetCompID)
  c.Check(ok, Equals, false)
}

func (s * DictionaryTests) TestIntSettings(c *C) {
  s.dict.SetInt(SocketAcceptPort, 101)

  val,ok:=s.dict.GetInt(SocketAcceptPort)
  c.Check(ok, Equals, true)
  c.Check(val, Equals, 101)

  val,ok=s.dict.GetInt(SocketConnectPort)
  c.Check(ok, Equals, false)
}
