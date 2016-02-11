package quickfix_test

import (
	"bytes"
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"testing"
	"time"
)

func TestMarshal_FIXMsgType(t *testing.T) {
	type Message struct {
		FIXMsgType string `fix:"0"`
	}
	var m Message
	fixMsg := quickfix.Marshal(m)
	fixMsg.Build()

	var msgType field.MsgTypeField
	fixMsg.Header.Get(&msgType)
	if msgType.FIXString != "0" {
		t.Errorf("Expected %v got %v", "1", msgType)
	}
}

func TestMarshal_Literals(t *testing.T) {
	type Message struct {
		FIXMsgType        string    `fix:"0"`
		StringField       string    `fix:"114"`
		IntField          int       `fix:"115"`
		UnsetIntField     int       `fix:"116"`
		BoolField         bool      `fix:"117"`
		FloatField        float64   `fix:"118"`
		UTCTimestampField time.Time `fix:"119"`
	}

	tVal, _ := time.Parse("2006-Jan-02", "2014-Jun-16")

	m := Message{
		StringField:       "world",
		IntField:          5,
		BoolField:         true,
		FloatField:        100.02,
		UTCTimestampField: tVal,
	}
	fixMsg := quickfix.Marshal(m)

	var tests = []struct {
		tag        quickfix.Tag
		fieldValue quickfix.FieldValue
		expected   []byte
	}{
		{quickfix.Tag(114), new(quickfix.FIXString), []byte("world")},
		{quickfix.Tag(115), new(quickfix.FIXInt), []byte("5")},
		{quickfix.Tag(116), new(quickfix.FIXInt), []byte("0")},
		{quickfix.Tag(117), new(quickfix.FIXBoolean), []byte("Y")},
		{quickfix.Tag(118), new(quickfix.FIXFloat), []byte("100.02")},
		{quickfix.Tag(119), new(quickfix.FIXUTCTimestamp), []byte("20140616-00:00:00.000")},
	}

	for _, test := range tests {
		if err := fixMsg.Body.GetField(test.tag, test.fieldValue); err != nil {
			t.Error("Unexpected error", err)
		}

		if !bytes.Equal(test.expected, test.fieldValue.Write()) {
			t.Errorf("Expected %s got %s", test.expected, test.fieldValue.Write())
		}
	}
}

func TestMarshal_LiteralsOptional(t *testing.T) {
	type Message struct {
		FIXMsgType          string   `fix:"0"`
		StringPtrField      *string  `fix:"112"`
		UnsetStringPtrField *string  `fix:"113"`
		IntPtrField         *int     `fix:"115"`
		BoolPtrField        *bool    `fix:"117"`
		FloatPtrField       *float64 `fix:"119"`
	}

	strVal := "hello"
	intVal := 42
	boolVal := false
	floatVal := 456.123
	m := Message{
		StringPtrField: &strVal,
		IntPtrField:    &intVal,
		BoolPtrField:   &boolVal,
		FloatPtrField:  &floatVal,
	}
	fixMsg := quickfix.Marshal(m)

	var tests = []struct {
		tag        quickfix.Tag
		fieldValue quickfix.FieldValue
		fieldSet   bool
		expected   []byte
	}{
		{quickfix.Tag(112), new(quickfix.FIXString), true, []byte("hello")},
		{quickfix.Tag(113), new(quickfix.FIXString), false, []byte{}},
		{quickfix.Tag(115), new(quickfix.FIXInt), true, []byte("42")},
		{quickfix.Tag(117), new(quickfix.FIXBoolean), true, []byte("N")},
		{quickfix.Tag(119), new(quickfix.FIXFloat), true, []byte("456.123")},
	}

	for _, test := range tests {
		if !test.fieldSet {
			if fixMsg.Body.Has(test.tag) {
				t.Error("Unexpected field for tag", test.tag)
			}
			continue
		}

		if err := fixMsg.Body.GetField(test.tag, test.fieldValue); err != nil {
			t.Error("Unexpected error", err)
		}

		if !bytes.Equal(test.expected, test.fieldValue.Write()) {
			t.Errorf("Expected %s got %s", test.expected, test.fieldValue.Write())
		}
	}
}

func TestMarshal_Components(t *testing.T) {
	type Component1 struct {
		StringField string `fix:"8"`
	}
	type Component2 struct {
		IntField int `fix:"1"`
	}

	type Message struct {
		FIXMsgType string `fix:"0"`
		Component1
		StringField string `fix:"112"`
		*Component2
	}

	m := Message{
		Component1:  Component1{StringField: "hello"},
		StringField: "world",
	}
	fixMsg := quickfix.Marshal(m)

	var tests = []struct {
		tag        quickfix.Tag
		fieldValue quickfix.FieldValue
		fieldSet   bool
		expected   []byte
	}{
		{quickfix.Tag(8), new(quickfix.FIXString), true, []byte("hello")},
		{quickfix.Tag(1), new(quickfix.FIXInt), false, []byte{}},
		{quickfix.Tag(112), new(quickfix.FIXString), true, []byte("world")},
	}

	for _, test := range tests {
		if !test.fieldSet {
			if fixMsg.Body.Has(test.tag) {
				t.Error("Unexpected field for tag", test.tag)
			}
			continue
		}

		if err := fixMsg.Body.GetField(test.tag, test.fieldValue); err != nil {
			t.Error("Unexpected error", err)
			continue
		}

		if !bytes.Equal(test.expected, test.fieldValue.Write()) {
			t.Errorf("Expected %s got %s", test.expected, test.fieldValue.Write())
		}
	}
}

func TestMarshal_RepeatingGroups(t *testing.T) {
	type Group1 struct {
		StringField1 string  `fix:"8"`
		StringField2 *string `fix:"9"`
	}

	type Group2 struct {
		IntField1 int `fix:"1"`
		IntField2 int `fix:"2"`
	}

	type Message struct {
		FIXMsgType  string   `fix:"0"`
		GroupField1 []Group1 `fix:"100"`
		StringField string   `fix:"112"`
		GroupField2 []Group2 `fix:"101"`
	}

	s := "world"
	m := Message{
		GroupField1: []Group1{Group1{StringField1: "hello", StringField2: &s}, Group1{StringField1: "goodbye"}, Group1{StringField1: "OHHAI", StringField2: &s}},
		StringField: "world",
		GroupField2: []Group2{Group2{IntField1: 1, IntField2: 42}},
	}
	fixMsg := quickfix.Marshal(m)

	group1Template := quickfix.GroupTemplate{
		quickfix.GroupElement(quickfix.Tag(8)),
		quickfix.GroupElement(quickfix.Tag(9)),
	}

	group2Template := quickfix.GroupTemplate{
		quickfix.GroupElement(quickfix.Tag(1)),
		quickfix.GroupElement(quickfix.Tag(2)),
	}

	var tests = []struct {
		groupTag          quickfix.Tag
		expectedGroupSize int
		template          quickfix.GroupTemplate
	}{
		{quickfix.Tag(100), 3, group1Template},
		{quickfix.Tag(101), 1, group2Template},
	}

	for _, test := range tests {
		if !fixMsg.Body.Has(test.groupTag) {
			t.Errorf("Expected tag %v", test.groupTag)
			continue
		}

		groupField := quickfix.RepeatingGroup{Tag: test.groupTag, GroupTemplate: test.template}
		if err := fixMsg.Body.GetGroup(&groupField); err != nil {
			t.Error("Unexpected error", err)
			continue
		}

		if len(groupField.Groups) != test.expectedGroupSize {
			t.Errorf("Expected group %v to have size %v, got %v", test.groupTag, test.expectedGroupSize, len(groupField.Groups))
		}
	}
}

func TestMarshal_Default(t *testing.T) {
	type Message struct {
		FIXMsgType   string  `fix:"0"`
		StringField1 string  `fix:"10,default=hello"`
		StringField2 *string `fix:"11,default=world"`
	}

	m := Message{}
	fixMsg := quickfix.Marshal(m)

	var tests = []struct {
		quickfix.Tag
		expected string
	}{
		{quickfix.Tag(10), "hello"},
		{quickfix.Tag(11), "world"},
	}

	for _, test := range tests {
		var f quickfix.FIXString
		if err := fixMsg.Body.GetField(test.Tag, &f); err != nil {
			t.Errorf("Unexpected Error for tag %v: %v", test.Tag, err)
		}

		if string(f) != test.expected {
			t.Errorf("Expected default value %v got %v", test.expected, f)
		}

	}

}

func TestMarshal_OmitEmpty(t *testing.T) {
	type Group1 struct {
		StringField1 string  `fix:"8"`
		StringField2 *string `fix:"9"`
	}

	type Group2 struct {
		IntField1 int `fix:"1"`
		IntField2 int `fix:"2"`
	}

	type Message struct {
		FIXMsgType  string   `fix:"0"`
		GroupField1 []Group1 `fix:"100"`
		StringField string   `fix:"112"`
		GroupField2 []Group2 `fix:"101,omitempty"`
	}

	m := Message{}
	fixMsg := quickfix.Marshal(m)

	if fixMsg.Body.Has(quickfix.Tag(101)) {
		t.Error("Unexpected tag for empty group")
	}

	f := new(quickfix.FIXInt)

	if err := fixMsg.Body.GetField(quickfix.Tag(100), f); err != nil {
		t.Errorf("Unexpected Error %v", err)
	}

	if *f != 0 {
		t.Errorf("Expected 0 got %v", *f)
	}
}

func TestMarshal_HeaderTrailer(t *testing.T) {
	type Header struct {
		StringField string `fix:"49"`
	}
	type Trailer struct {
		IntField int `fix:"89"`
	}

	type Message struct {
		FIXMsgType  string `fix:"0"`
		StringField string `fix:"112"`
		Header
		Trailer
	}

	m := Message{
		Header:      Header{StringField: "hello"},
		StringField: "world",
		Trailer:     Trailer{IntField: 3},
	}
	fixMsg := quickfix.Marshal(m)

	var tests = []struct {
		tag        quickfix.Tag
		fieldValue quickfix.FieldValue
		fixMsgPart quickfix.FieldMap
		expected   []byte
	}{
		{quickfix.Tag(35), new(quickfix.FIXString), fixMsg.Header, []byte("0")},
		{quickfix.Tag(49), new(quickfix.FIXString), fixMsg.Header, []byte("hello")},
		{quickfix.Tag(112), new(quickfix.FIXString), fixMsg.Body, []byte("world")},
		{quickfix.Tag(89), new(quickfix.FIXInt), fixMsg.Trailer, []byte("3")},
	}

	for _, test := range tests {
		if err := test.fixMsgPart.GetField(test.tag, test.fieldValue); err != nil {
			t.Error("Unexpected error", err)
		}

		if !bytes.Equal(test.expected, test.fieldValue.Write()) {
			t.Errorf("Expected '%s' got '%s'", test.expected, test.fieldValue.Write())
		}
	}
}
