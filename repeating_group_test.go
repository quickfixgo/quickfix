package quickfix

import (
	"bytes"
	"testing"
)

func TestRepeatingGroup_Add(t *testing.T) {
	f := RepeatingGroup{GroupTemplate: GroupTemplate{GroupElement(1)}}

	var tests = []struct {
		expectedLen int
	}{
		{1},
		{2},
	}

	for _, test := range tests {
		g := f.Add()
		if len(f.Groups) != test.expectedLen {
			t.Errorf("Expected %v groups, got %v", test.expectedLen, len(f.Groups))
		}

		g.SetField(Tag(1), FIXString("hello"))

		if !f.Groups[test.expectedLen-1].Has(Tag(1)) {
			t.Errorf("expected tag in group %v", test.expectedLen)
		}

		var v FIXString
		f.Groups[test.expectedLen-1].GetField(Tag(1), &v)

		if string(v) != "hello" {
			t.Errorf("expected hello in group %v", test.expectedLen)
		}
	}
}

func TestRepeatingGroup_Write(t *testing.T) {
	f1 := RepeatingGroup{GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}

	f1.Add().SetField(Tag(1), FIXString("hello"))

	f2 := RepeatingGroup{GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f2.Add().SetField(Tag(1), FIXString("hello")).SetField(Tag(2), FIXString("world"))

	f3 := RepeatingGroup{GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f3.Add().SetField(Tag(1), FIXString("hello"))
	f3.Add().SetField(Tag(1), FIXString("world"))

	f4 := RepeatingGroup{GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f4.Add().SetField(Tag(1), FIXString("hello")).SetField(Tag(2), FIXString("world"))
	f4.Add().SetField(Tag(1), FIXString("goodbye"))

	var tests = []struct {
		f        RepeatingGroup
		expected []byte
	}{
		{f1, []byte("11=hello")},
		{f2, []byte("11=hello2=world")},
		{f3, []byte("21=hello1=world")},
		{f4, []byte("21=hello2=world1=goodbye")},
	}

	for _, test := range tests {
		fieldBytes := test.f.Write()
		if !bytes.Equal(test.expected, fieldBytes) {
			t.Errorf("expected %s got %s", test.expected, fieldBytes)
		}
	}
}

func TestRepeatingGroup_ReadError(t *testing.T) {
	singleFieldTemplate := GroupTemplate{GroupElement(1)}
	tests := []struct {
		tv               TagValues
		expectedGroupNum int
	}{
		{
			TagValues{
				TagValue{Value: []byte("1")},
				TagValue{Tag: Tag(2), Value: []byte("not in template")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
			}, 0},
		{
			TagValues{
				TagValue{Value: []byte("2")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
				TagValue{Tag: Tag(2), Value: []byte("not in template")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
			}, 1}}

	for _, s := range tests {
		f := RepeatingGroup{GroupTemplate: singleFieldTemplate}
		_, err := f.Read(s.tv)
		if err == nil || len(f.Groups) != s.expectedGroupNum {
			t.Errorf("Should have raised an error because expected group number is wrong: %v instead of %v", len(f.Groups), s.expectedGroupNum)
		}
	}
}

func TestRepeatingGroup_Read(t *testing.T) {

	singleFieldTemplate := GroupTemplate{GroupElement(1)}
	multiFieldTemplate := GroupTemplate{GroupElement(1), GroupElement(2), GroupElement(3)}

	tests := []struct {
		groupTemplate    GroupTemplate
		tv               TagValues
		expectedGroupTvs []TagValues
	}{
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("0")}},
			[]TagValues{}},
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("1")}, TagValue{Tag: Tag(1), Value: []byte("hello")}},
			[]TagValues{TagValues{TagValue{Tag: Tag(1), Value: []byte("hello")}}}},
		{singleFieldTemplate,
			TagValues{TagValue{Value: []byte("1")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
				TagValue{Tag: Tag(2), Value: []byte("not in group")}},
			[]TagValues{
				TagValues{TagValue{Tag: Tag(1), Value: []byte("hello")}}}},
		{singleFieldTemplate,
			TagValues{TagValue{Value: []byte("2")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
				TagValue{Tag: Tag(1), Value: []byte("world")}},
			[]TagValues{
				TagValues{TagValue{Tag: Tag(1), Value: []byte("hello")}},
				TagValues{TagValue{Tag: Tag(1), Value: []byte("world")}},
			}},
		{multiFieldTemplate,
			TagValues{
				TagValue{Value: []byte("2")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
				TagValue{Tag: Tag(1), Value: []byte("goodbye")}, TagValue{Tag: Tag(2), Value: []byte("cruel")}, TagValue{Tag: Tag(3), Value: []byte("world")},
			},
			[]TagValues{
				TagValues{TagValue{Tag: Tag(1), Value: []byte("hello")}},
				TagValues{TagValue{Tag: Tag(1), Value: []byte("goodbye")}, TagValue{Tag: Tag(2), Value: []byte("cruel")}, TagValue{Tag: Tag(3), Value: []byte("world")}},
			}},
		{multiFieldTemplate,
			TagValues{
				TagValue{Value: []byte("3")},
				TagValue{Tag: Tag(1), Value: []byte("hello")},
				TagValue{Tag: Tag(1), Value: []byte("goodbye")}, TagValue{Tag: Tag(2), Value: []byte("cruel")}, TagValue{Tag: Tag(3), Value: []byte("world")},
				TagValue{Tag: Tag(1), Value: []byte("another")},
			},
			[]TagValues{
				TagValues{TagValue{Tag: Tag(1), Value: []byte("hello")}},
				TagValues{TagValue{Tag: Tag(1), Value: []byte("goodbye")}, TagValue{Tag: Tag(2), Value: []byte("cruel")}, TagValue{Tag: Tag(3), Value: []byte("world")}},
				TagValues{TagValue{Tag: Tag(1), Value: []byte("another")}},
			}},
	}

	for _, test := range tests {
		f := RepeatingGroup{GroupTemplate: test.groupTemplate}

		_, err := f.Read(test.tv)
		if err != nil {
			t.Error(err)
		}
		if len(f.Groups) != len(test.expectedGroupTvs) {
			t.Errorf("expected %v groups, got %v", len(test.expectedGroupTvs), len(f.Groups))
		}

		for g, group := range f.Groups {
			for _, expected := range test.expectedGroupTvs[g] {
				var actual FIXString
				group.GetField(expected.Tag, &actual)

				if !bytes.Equal(expected.Value, []byte(actual)) {
					t.Errorf("%v, %v: expected %s, got %s", g, expected.Tag, expected.Value, actual)
				}
			}
		}
	}
}

func TestRepeatingGroup_ReadComplete(t *testing.T) {

	rawMsg := []byte("8=FIXT.1.19=26835=W34=711849=TEST52=20151027-18:41:52.69856=TST22=9948=TSTX15262=7268=4269=4270=0.07499272=20151027273=18:41:52.698269=7270=0.07501272=20151027273=18:41:52.698269=8270=0.07494272=20151027273=18:41:52.698269=B271=60272=20151027273=18:41:52.69810=163")

	msg, err := parseMessage(rawMsg)

	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	template := GroupTemplate{
		GroupElement(269),
		GroupElement(270),
		GroupElement(271),
		GroupElement(272),
		GroupElement(273),
	}

	f := RepeatingGroup{GroupTemplate: template}
	err = msg.Body.GetGroupField(Tag(268), &f)
	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	if len(f.Groups) != 4 {
		t.Errorf("expected %v groups, got %v", 4, len(f.Groups))
	}

	expectedGroupTags := [][]Tag{
		[]Tag{Tag(269), Tag(270), Tag(272), Tag(273)},
		[]Tag{Tag(269), Tag(270), Tag(272), Tag(273)},
		[]Tag{Tag(269), Tag(270), Tag(272), Tag(273)},
		[]Tag{Tag(269), Tag(271), Tag(272), Tag(273)},
	}

	expectedGroupValues := [][]FIXString{
		[]FIXString{FIXString("4"), FIXString("0.07499"), FIXString("20151027"), FIXString("18:41:52.698")},
		[]FIXString{FIXString("7"), FIXString("0.07501"), FIXString("20151027"), FIXString("18:41:52.698")},
		[]FIXString{FIXString("8"), FIXString("0.07494"), FIXString("20151027"), FIXString("18:41:52.698")},
		[]FIXString{FIXString("B"), FIXString("60"), FIXString("20151027"), FIXString("18:41:52.698")},
	}

	for i, group := range f.Groups {

		for j, tag := range expectedGroupTags[i] {
			if !group.Has(tag) {
				t.Errorf("expected %v in group %v", expectedGroupTags[i][j], i)
				continue
			}

			var actual FIXString
			if err := group.GetField(tag, &actual); err != nil {
				t.Errorf("error getting field %v from group %v", tag, i)
				continue
			}

			if expectedGroupValues[i][j] != actual {
				t.Errorf("Expected %v got %v", expectedGroupTags[i][j], actual)
			}
		}
	}
}
