package datadictionary

import (
	"github.com/cbusbey/quickfixgo/tag"
	. "launchpad.net/gocheck"
)

var _ = Suite(&DataDictionaryTests{})

type DataDictionaryTests struct {
	dict *DataDictionary
}

func (s *DataDictionaryTests) SetUpTest(c *C) {
	dict, err := Parse("../spec/FIX43.xml")
	c.Check(err, IsNil)
	c.Check(dict, NotNil)

	s.dict = dict
}

func (s *DataDictionaryTests) TestParseBadPath(c *C) {
	_, err := Parse("../spec/bogus.xml")
	c.Check(err, NotNil)
}

func (s *DataDictionaryTests) TestParseRecursiveComponents(c *C) {
	dict, err := Parse("../spec/FIX44.xml")
	c.Check(err, IsNil)
	c.Check(dict, NotNil)
}

func (s *DataDictionaryTests) TestComponents(c *C) {
	_, ok := s.dict.Components["SpreadOrBenchmarkCurveData"]
	c.Check(ok, Equals, true)
}

func (s *DataDictionaryTests) TestFieldsByTag(c *C) {
	f, ok := s.dict.FieldTypeByTag[tag.ContraLegRefID]
	c.Check(ok, Equals, true)
	c.Check(f.Tag, Equals, tag.ContraLegRefID)
	c.Check(f.Name, Equals, "ContraLegRefID")
	c.Check(f.Type, Equals, "STRING")
	c.Check(f.Enums, IsNil)

	f, ok = s.dict.FieldTypeByTag[tag.QuoteRequestRejectReason]
	c.Check(ok, Equals, true)
	c.Check(f.Tag, Equals, tag.QuoteRequestRejectReason)
	c.Check(f.Name, Equals, "QuoteRequestRejectReason")
	c.Check(f.Type, Equals, "INT")

	c.Check(f.Enums, NotNil)
	c.Check(len(f.Enums), Equals, 6)
	c.Check(f.Enums["1"].Value, Equals, "1")
	c.Check(f.Enums["1"].Description, Equals, "UNKNOWN_SYMBOL")
}

func (s *DataDictionaryTests) TestMessages(c *C) {
	_, ok := s.dict.Messages["D"]
	c.Check(ok, Equals, true)
}

func (s *DataDictionaryTests) TestHeader(c *C) {
	c.Check(s.dict.Header, NotNil)
}

func (s *DataDictionaryTests) TestTrailer(c *C) {
	c.Check(s.dict.Trailer, NotNil)
}

func (s *DataDictionaryTests) TestMessageRequiredTags(c *C) {
	m, ok := s.dict.Messages["D"]
	c.Check(ok, Equals, true)

	_, required := m.RequiredTags[tag.ClOrdID]
	c.Check(required, Equals, true)
	_, required = m.RequiredTags[tag.SecondaryClOrdID]
	c.Check(required, Equals, false)

	_, required = s.dict.Header.RequiredTags[tag.BeginString]
	c.Check(required, Equals, true)

	_, required = s.dict.Header.RequiredTags[tag.OnBehalfOfCompID]
	c.Check(required, Equals, false)

	_, required = s.dict.Trailer.RequiredTags[tag.CheckSum]
	c.Check(required, Equals, true)

	_, required = s.dict.Trailer.RequiredTags[tag.Signature]
	c.Check(required, Equals, false)

}

func (s *DataDictionaryTests) TestMessageTags(c *C) {
	m, ok := s.dict.Messages["D"]
	c.Check(ok, Equals, true)

	_, known := m.Tags[tag.ClOrdID]
	c.Check(known, Equals, true)
	_, known = m.Tags[tag.SecondaryClOrdID]
	c.Check(known, Equals, true)

	_, known = s.dict.Header.Tags[tag.BeginString]
	c.Check(known, Equals, true)

	_, known = s.dict.Header.Tags[tag.OnBehalfOfCompID]
	c.Check(known, Equals, true)

	_, known = s.dict.Trailer.Tags[tag.CheckSum]
	c.Check(known, Equals, true)

	_, known = s.dict.Trailer.Tags[tag.Signature]
	c.Check(known, Equals, true)
}
