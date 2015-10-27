package datadictionary

import (
	. "gopkg.in/check.v1"
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
	f, ok := s.dict.FieldTypeByTag[655]
	c.Check(ok, Equals, true)
	c.Check(f.Tag, Equals, 655)
	c.Check(f.Name, Equals, "ContraLegRefID")
	c.Check(f.Type, Equals, "STRING")
	c.Check(f.Enums, IsNil)

	f, ok = s.dict.FieldTypeByTag[658]
	c.Check(ok, Equals, true)
	c.Check(f.Tag, Equals, 658)
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

	_, required := m.RequiredTags[11]
	c.Check(required, Equals, true)
	_, required = m.RequiredTags[526]
	c.Check(required, Equals, false)

	_, required = s.dict.Header.RequiredTags[8]
	c.Check(required, Equals, true)

	_, required = s.dict.Header.RequiredTags[115]
	c.Check(required, Equals, false)

	_, required = s.dict.Trailer.RequiredTags[10]
	c.Check(required, Equals, true)

	_, required = s.dict.Trailer.RequiredTags[89]
	c.Check(required, Equals, false)

}

func (s *DataDictionaryTests) TestMessageTags(c *C) {
	m, ok := s.dict.Messages["D"]
	c.Check(ok, Equals, true)

	_, known := m.Tags[11]
	c.Check(known, Equals, true)
	_, known = m.Tags[526]
	c.Check(known, Equals, true)

	_, known = s.dict.Header.Tags[8]
	c.Check(known, Equals, true)

	_, known = s.dict.Header.Tags[115]
	c.Check(known, Equals, true)

	_, known = s.dict.Trailer.Tags[10]
	c.Check(known, Equals, true)

	_, known = s.dict.Trailer.Tags[89]
	c.Check(known, Equals, true)
}
