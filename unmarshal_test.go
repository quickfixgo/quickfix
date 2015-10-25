package quickfix_test

import (
	"github.com/quickfixgo/quickfix"
	"testing"
	"time"
)

func TestUnmarshal_Literals(t *testing.T) {
	type Message struct {
		FIXMsgType        string    `fix:"0"`
		StringField       string    `fix:"114"`
		IntField          int       `fix:"115"`
		UnsetIntField     int       `fix:"116"`
		BoolField         bool      `fix:"117"`
		FloatField        float64   `fix:"118"`
		UTCTimestampField time.Time `fix:"119"`
	}

	fixMsg := quickfix.NewMessage()
	fixMsg.Body.SetField(quickfix.Tag(114), quickfix.FIXString("world"))
	fixMsg.Body.SetField(quickfix.Tag(115), quickfix.FIXInt(3))
	fixMsg.Body.SetField(quickfix.Tag(117), quickfix.FIXBoolean(true))
	fixMsg.Body.SetField(quickfix.Tag(118), quickfix.FIXFloat(500.123))
	tVal, _ := time.Parse("2006-Jan-02", "2014-Jun-16")
	fixMsg.Body.SetField(quickfix.Tag(119), quickfix.FIXUTCTimestamp{Value: tVal})

	var msgOut Message
	quickfix.Unmarshal(fixMsg, &msgOut)

	var tests = []struct {
		expected interface{}
		actual   interface{}
	}{
		{"world", msgOut.StringField},
		{3, msgOut.IntField},
		{0, msgOut.UnsetIntField},
		{true, msgOut.BoolField},
		{500.123, msgOut.FloatField},
		{tVal, msgOut.UTCTimestampField},
	}

	for _, test := range tests {
		if test.expected != test.actual {
			t.Errorf("Expected %v got %v", test.expected, test.actual)
		}
	}
}

func TestUnmarshal_LiteralsOptional(t *testing.T) {
	type Message struct {
		FIXMsgType           string     `fix:"0"`
		StringPtrField       *string    `fix:"112"`
		UnsetStringPtrField  *string    `fix:"113"`
		IntPtrField          *int       `fix:"114"`
		BoolPtrField         *bool      `fix:"115"`
		FloatPtrField        *float64   `fix:"116"`
		UTCTimestampPtrField *time.Time `fix:"117"`
	}

	fixMsg := quickfix.NewMessage()
	fixMsg.Body.SetField(quickfix.Tag(112), quickfix.FIXString("world"))
	fixMsg.Body.SetField(quickfix.Tag(114), quickfix.FIXInt(3))
	fixMsg.Body.SetField(quickfix.Tag(115), quickfix.FIXBoolean(false))
	fixMsg.Body.SetField(quickfix.Tag(116), quickfix.FIXFloat(500.123))
	tVal, _ := time.Parse("2006-Jan-02", "2014-Jun-16")
	fixMsg.Body.SetField(quickfix.Tag(117), quickfix.FIXUTCTimestamp{Value: tVal})

	var msgOut Message
	quickfix.Unmarshal(fixMsg, &msgOut)

	var tests = []struct {
		testName      string
		expectNil     bool
		actualIsNil   bool
		expectedDeref interface{}
		actualDeref   interface{}
	}{
		{"set string ptr", false, msgOut.StringPtrField == nil, "world", *msgOut.StringPtrField},
		{"unset string ptr", true, msgOut.UnsetStringPtrField == nil, "", "unset"},
		{"int ptr", false, msgOut.IntPtrField == nil, 3, *msgOut.IntPtrField},
		{"bool ptr", false, msgOut.BoolPtrField == nil, false, *msgOut.BoolPtrField},
		{"float ptr", false, msgOut.FloatPtrField == nil, 500.123, *msgOut.FloatPtrField},
		{"timestamp ptr", false, msgOut.UTCTimestampPtrField == nil, tVal, *msgOut.UTCTimestampPtrField},
	}

	for _, test := range tests {
		if test.expectNil {
			if !test.actualIsNil {
				t.Errorf("%v: Expected nil not %v", test.testName, test.actualDeref)
			}
			continue
		}

		if test.actualIsNil {
			t.Errorf("%v: Did not expect nil", test.testName)
			continue
		}

		if test.expectedDeref != test.actualDeref {
			t.Errorf("Expected '%v' got '%v'", test.expectedDeref, test.actualDeref)
		}
	}
}

func TestUnmarshal_Components(t *testing.T) {
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

	fixMsg := quickfix.NewMessage()
	fixMsg.Body.SetField(quickfix.Tag(8), quickfix.FIXString("hello"))
	fixMsg.Body.SetField(quickfix.Tag(112), quickfix.FIXString("world"))

	var msgOut Message
	quickfix.Unmarshal(fixMsg, &msgOut)

	if msgOut.Component1.StringField != "hello" {
		t.Errorf("Expected '%v' got '%v'", "hello", msgOut.Component1.StringField)
	}

	if msgOut.StringField != "world" {
		t.Errorf("Expected '%v' got '%v'", "world", msgOut.StringField)
	}
}

func TestUnmarshal_RepeatingGroups(t *testing.T) {
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

	fixMsg := quickfix.NewMessage()
	fixMsg.Body.SetField(quickfix.Tag(112), quickfix.FIXString("world"))

	group1Template := quickfix.GroupTemplate{
		quickfix.GroupElement(quickfix.Tag(8)),
		quickfix.GroupElement(quickfix.Tag(9)),
	}

	group1 := quickfix.RepeatingGroup{Tag: quickfix.Tag(100), GroupTemplate: group1Template}
	group1.Add().SetField(quickfix.Tag(8), quickfix.FIXString("hello")).SetField(quickfix.Tag(9), quickfix.FIXString("world"))
	group1.Add().SetField(quickfix.Tag(8), quickfix.FIXString("goodbye"))
	group1.Add().SetField(quickfix.Tag(8), quickfix.FIXString("OHHAI")).SetField(quickfix.Tag(9), quickfix.FIXString("world"))
	fixMsg.Body.SetGroup(group1)

	var msgOut Message
	quickfix.Unmarshal(fixMsg, &msgOut)

	if msgOut.StringField != "world" {
		t.Errorf("Expected '%v' got '%v'", "world", msgOut.StringField)
	}

	if len(msgOut.GroupField1) != 3 {
		t.Errorf("Expected group size '%v' got '%v'", 3, len(msgOut.GroupField1))
	}

	if msgOut.GroupField1[0].StringField1 != "hello" {
		t.Errorf("Expected %v got %v", "hello", msgOut.GroupField1[0].StringField1)
	}

	if *(msgOut.GroupField1[0].StringField2) != "world" {
		t.Errorf("Expected %v got %v", "world", *(msgOut.GroupField1[0].StringField2))
	}

	if msgOut.GroupField1[1].StringField1 != "goodbye" {
		t.Errorf("Expected %v got %v", "goodbye", msgOut.GroupField1[1].StringField1)
	}

	if msgOut.GroupField1[1].StringField2 != nil {
		t.Errorf("Expected stringfield 2 to be nil, got %v", *(msgOut.GroupField1[1].StringField2))
	}
}

func TestUnmarshal_HeaderTrailer(t *testing.T) {
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

	fixMsg := quickfix.NewMessage()
	fixMsg.Header.SetField(quickfix.Tag(35), quickfix.FIXString("0"))
	fixMsg.Header.SetField(quickfix.Tag(49), quickfix.FIXString("hello"))
	fixMsg.Body.SetField(quickfix.Tag(112), quickfix.FIXString("world"))
	fixMsg.Trailer.SetField(quickfix.Tag(89), quickfix.FIXInt(3))

	var msgOut Message
	quickfix.Unmarshal(fixMsg, &msgOut)

	if msgOut.Header.StringField != "hello" {
		t.Errorf("Expected '%v' got '%v'", "hello", msgOut.Header.StringField)
	}
	if msgOut.Trailer.IntField != 3 {
		t.Errorf("Expected '%v' got '%v'", 3, msgOut.Trailer.IntField)
	}

}
