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

func TestRepeatingGroup_TagValues(t *testing.T) {
	f1 := RepeatingGroup{Tag: 10, GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}

	f1.Add().SetField(Tag(1), FIXString("hello"))

	f2 := RepeatingGroup{Tag: 11, GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f2.Add().SetField(Tag(1), FIXString("hello")).SetField(Tag(2), FIXString("world"))

	f3 := RepeatingGroup{Tag: 12, GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f3.Add().SetField(Tag(1), FIXString("hello"))
	f3.Add().SetField(Tag(1), FIXString("world"))

	f4 := RepeatingGroup{Tag: 13, GroupTemplate: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f4.Add().SetField(Tag(1), FIXString("hello")).SetField(Tag(2), FIXString("world"))
	f4.Add().SetField(Tag(1), FIXString("goodbye"))

	var tests = []struct {
		f        RepeatingGroup
		expected []byte
	}{
		{f1, []byte("10=11=hello")},
		{f2, []byte("11=11=hello2=world")},
		{f3, []byte("12=21=hello1=world")},
		{f4, []byte("13=21=hello2=world1=goodbye")},
	}

	for _, test := range tests {
		tvbytes := []byte{}
		for _, tv := range test.f.tagValues() {
			tvbytes = append(tvbytes, tv.bytes...)
		}

		if !bytes.Equal(test.expected, tvbytes) {
			t.Errorf("expected %s got %s", test.expected, tvbytes)
		}
	}
}

func TestRepeatingGroup_ReadError(t *testing.T) {
	singleFieldTemplate := GroupTemplate{GroupElement(1)}
	tests := []struct {
		tv               tagValues
		expectedGroupNum int
	}{
		{
			tagValues{
				tagValue{Value: []byte("1")},
				tagValue{Tag: Tag(2), Value: []byte("not in template")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
			}, 0},
		{
			tagValues{
				tagValue{Value: []byte("2")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
				tagValue{Tag: Tag(2), Value: []byte("not in template")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
			}, 1}}

	for _, s := range tests {
		f := RepeatingGroup{GroupTemplate: singleFieldTemplate}
		_, err := f.read(s.tv)
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
		tv               tagValues
		expectedGroupTvs []tagValues
	}{
		{singleFieldTemplate, tagValues{tagValue{Value: []byte("0")}},
			[]tagValues{}},
		{singleFieldTemplate, tagValues{tagValue{Value: []byte("1")}, tagValue{Tag: Tag(1), Value: []byte("hello")}},
			[]tagValues{tagValues{tagValue{Tag: Tag(1), Value: []byte("hello")}}}},
		{singleFieldTemplate,
			tagValues{tagValue{Value: []byte("1")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
				tagValue{Tag: Tag(2), Value: []byte("not in group")}},
			[]tagValues{
				tagValues{tagValue{Tag: Tag(1), Value: []byte("hello")}}}},
		{singleFieldTemplate,
			tagValues{tagValue{Value: []byte("2")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
				tagValue{Tag: Tag(1), Value: []byte("world")}},
			[]tagValues{
				tagValues{tagValue{Tag: Tag(1), Value: []byte("hello")}},
				tagValues{tagValue{Tag: Tag(1), Value: []byte("world")}},
			}},
		{multiFieldTemplate,
			tagValues{
				tagValue{Value: []byte("2")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
				tagValue{Tag: Tag(1), Value: []byte("goodbye")}, tagValue{Tag: Tag(2), Value: []byte("cruel")}, tagValue{Tag: Tag(3), Value: []byte("world")},
			},
			[]tagValues{
				tagValues{tagValue{Tag: Tag(1), Value: []byte("hello")}},
				tagValues{tagValue{Tag: Tag(1), Value: []byte("goodbye")}, tagValue{Tag: Tag(2), Value: []byte("cruel")}, tagValue{Tag: Tag(3), Value: []byte("world")}},
			}},
		{multiFieldTemplate,
			tagValues{
				tagValue{Value: []byte("3")},
				tagValue{Tag: Tag(1), Value: []byte("hello")},
				tagValue{Tag: Tag(1), Value: []byte("goodbye")}, tagValue{Tag: Tag(2), Value: []byte("cruel")}, tagValue{Tag: Tag(3), Value: []byte("world")},
				tagValue{Tag: Tag(1), Value: []byte("another")},
			},
			[]tagValues{
				tagValues{tagValue{Tag: Tag(1), Value: []byte("hello")}},
				tagValues{tagValue{Tag: Tag(1), Value: []byte("goodbye")}, tagValue{Tag: Tag(2), Value: []byte("cruel")}, tagValue{Tag: Tag(3), Value: []byte("world")}},
				tagValues{tagValue{Tag: Tag(1), Value: []byte("another")}},
			}},
	}

	for _, test := range tests {
		f := RepeatingGroup{GroupTemplate: test.groupTemplate}

		_, err := f.read(test.tv)
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

	f := RepeatingGroup{Tag: 268, GroupTemplate: template}
	err = msg.Body.GetGroup(&f)
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
