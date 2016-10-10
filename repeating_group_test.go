package quickfix

import (
	"bytes"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestRepeatingGroup_Add(t *testing.T) {
	f := RepeatingGroup{template: GroupTemplate{GroupElement(1)}}

	var tests = []struct {
		expectedLen int
	}{
		{1},
		{2},
	}

	for _, test := range tests {
		g := f.Add()
		if f.Len() != test.expectedLen {
			t.Errorf("Expected %v groups, got %v", test.expectedLen, f.Len())
		}

		g.SetField(Tag(1), FIXString("hello"))

		if !f.groups[test.expectedLen-1].Has(Tag(1)) {
			t.Errorf("expected tag in group %v", test.expectedLen)
		}

		var v FIXString
		require.Nil(t, f.groups[test.expectedLen-1].GetField(Tag(1), &v))

		if string(v) != "hello" {
			t.Errorf("expected hello in group %v", test.expectedLen)
		}
	}
}

func TestRepeatingGroup_Write(t *testing.T) {
	f1 := RepeatingGroup{tag: 10, template: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}

	f1.Add().SetField(Tag(1), FIXString("hello"))

	f2 := RepeatingGroup{tag: 11, template: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f2.Add().SetField(Tag(1), FIXString("hello")).SetField(Tag(2), FIXString("world"))

	f3 := RepeatingGroup{tag: 12, template: GroupTemplate{
		GroupElement(1),
		GroupElement(2),
	}}
	f3.Add().SetField(Tag(1), FIXString("hello"))
	f3.Add().SetField(Tag(1), FIXString("world"))

	f4 := RepeatingGroup{tag: 13, template: GroupTemplate{
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
		for _, tv := range test.f.Write() {
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
		tv               []TagValue
		expectedGroupNum int
	}{
		{
			[]TagValue{
				TagValue{value: []byte("1")},
				TagValue{tag: Tag(2), value: []byte("not in template")},
				TagValue{tag: Tag(1), value: []byte("hello")},
			}, 0},
		{
			[]TagValue{
				TagValue{value: []byte("2")},
				TagValue{tag: Tag(1), value: []byte("hello")},
				TagValue{tag: Tag(2), value: []byte("not in template")},
				TagValue{tag: Tag(1), value: []byte("hello")},
			}, 1}}

	for _, s := range tests {
		f := RepeatingGroup{template: singleFieldTemplate}
		_, err := f.Read(s.tv)
		if err == nil || len(f.groups) != s.expectedGroupNum {
			t.Errorf("Should have raised an error because expected group number is wrong: %v instead of %v", len(f.groups), s.expectedGroupNum)
		}
	}
}

func TestRepeatingGroup_Read(t *testing.T) {

	singleFieldTemplate := GroupTemplate{GroupElement(1)}
	multiFieldTemplate := GroupTemplate{GroupElement(1), GroupElement(2), GroupElement(3)}

	tests := []struct {
		groupTemplate    GroupTemplate
		tv               []TagValue
		expectedGroupTvs [][]TagValue
	}{
		{singleFieldTemplate, []TagValue{TagValue{value: []byte("0")}},
			[][]TagValue{}},
		{singleFieldTemplate, []TagValue{TagValue{value: []byte("1")}, TagValue{tag: Tag(1), value: []byte("hello")}},
			[][]TagValue{{TagValue{tag: Tag(1), value: []byte("hello")}}}},
		{singleFieldTemplate,
			[]TagValue{TagValue{value: []byte("1")},
				TagValue{tag: Tag(1), value: []byte("hello")},
				TagValue{tag: Tag(2), value: []byte("not in group")}},
			[][]TagValue{
				{TagValue{tag: Tag(1), value: []byte("hello")}}}},
		{singleFieldTemplate,
			[]TagValue{TagValue{value: []byte("2")},
				TagValue{tag: Tag(1), value: []byte("hello")},
				TagValue{tag: Tag(1), value: []byte("world")}},
			[][]TagValue{
				{TagValue{tag: Tag(1), value: []byte("hello")}},
				{TagValue{tag: Tag(1), value: []byte("world")}},
			}},
		{multiFieldTemplate,
			[]TagValue{
				TagValue{value: []byte("2")},
				TagValue{tag: Tag(1), value: []byte("hello")},
				TagValue{tag: Tag(1), value: []byte("goodbye")}, TagValue{tag: Tag(2), value: []byte("cruel")}, TagValue{tag: Tag(3), value: []byte("world")},
			},
			[][]TagValue{
				{TagValue{tag: Tag(1), value: []byte("hello")}},
				{TagValue{tag: Tag(1), value: []byte("goodbye")}, TagValue{tag: Tag(2), value: []byte("cruel")}, TagValue{tag: Tag(3), value: []byte("world")}},
			}},
		{multiFieldTemplate,
			[]TagValue{
				TagValue{value: []byte("3")},
				TagValue{tag: Tag(1), value: []byte("hello")},
				TagValue{tag: Tag(1), value: []byte("goodbye")}, TagValue{tag: Tag(2), value: []byte("cruel")}, TagValue{tag: Tag(3), value: []byte("world")},
				TagValue{tag: Tag(1), value: []byte("another")},
			},
			[][]TagValue{
				{TagValue{tag: Tag(1), value: []byte("hello")}},
				{TagValue{tag: Tag(1), value: []byte("goodbye")}, TagValue{tag: Tag(2), value: []byte("cruel")}, TagValue{tag: Tag(3), value: []byte("world")}},
				{TagValue{tag: Tag(1), value: []byte("another")}},
			}},
	}

	for _, test := range tests {
		f := RepeatingGroup{template: test.groupTemplate}

		_, err := f.Read(test.tv)
		if err != nil {
			t.Error(err)
		}
		if len(f.groups) != len(test.expectedGroupTvs) {
			t.Errorf("expected %v groups, got %v", len(test.expectedGroupTvs), len(f.groups))
		}

		for g, group := range f.groups {
			for _, expected := range test.expectedGroupTvs[g] {
				var actual FIXString
				require.Nil(t, group.GetField(expected.tag, &actual))

				if !bytes.Equal(expected.value, []byte(actual)) {
					t.Errorf("%v, %v: expected %s, got %s", g, expected.tag, expected.value, actual)
				}
			}
		}
	}
}

func TestRepeatingGroup_ReadRecursive(t *testing.T) {
	singleFieldTemplate := GroupTemplate{GroupElement(4)}
	parentTemplate := GroupTemplate{GroupElement(2), NewRepeatingGroup(Tag(3), singleFieldTemplate), GroupElement(5)}

	f := NewRepeatingGroup(Tag(1), parentTemplate)
	_, err := f.Read([]TagValue{
		TagValue{value: []byte("2")},
		TagValue{tag: Tag(2), value: []byte("hello")},
		TagValue{tag: 3, value: []byte("1")}, TagValue{tag: 4, value: []byte("foo")},
		TagValue{tag: Tag(2), value: []byte("world")},
		TagValue{tag: 3, value: []byte("2")}, TagValue{tag: 4, value: []byte("foo")}, TagValue{tag: 4, value: []byte("bar")}, TagValue{tag: 5, value: []byte("fubar")},
	})
	require.Nil(t, err)

	assert.Equal(t, 2, f.Len())
}

func TestRepeatingGroup_ReadComplete(t *testing.T) {

	rawMsg := bytes.NewBufferString("8=FIXT.1.19=26835=W34=711849=TEST52=20151027-18:41:52.69856=TST22=9948=TSTX15262=7268=4269=4270=0.07499272=20151027273=18:41:52.698269=7270=0.07501272=20151027273=18:41:52.698269=8270=0.07494272=20151027273=18:41:52.698269=B271=60272=20151027273=18:41:52.69810=163")

	msg := NewMessage()
	err := ParseMessage(msg, rawMsg)

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

	f := RepeatingGroup{tag: 268, template: template}
	err = msg.Body.GetGroup(&f)
	if err != nil {
		t.Error("Unexpected error, ", err)
	}

	if len(f.groups) != 4 {
		t.Errorf("expected %v groups, got %v", 4, len(f.groups))
	}

	expectedGroupTags := [][]Tag{
		{Tag(269), Tag(270), Tag(272), Tag(273)},
		{Tag(269), Tag(270), Tag(272), Tag(273)},
		{Tag(269), Tag(270), Tag(272), Tag(273)},
		{Tag(269), Tag(271), Tag(272), Tag(273)},
	}

	expectedGroupValues := [][]FIXString{
		{FIXString("4"), FIXString("0.07499"), FIXString("20151027"), FIXString("18:41:52.698")},
		{FIXString("7"), FIXString("0.07501"), FIXString("20151027"), FIXString("18:41:52.698")},
		{FIXString("8"), FIXString("0.07494"), FIXString("20151027"), FIXString("18:41:52.698")},
		{FIXString("B"), FIXString("60"), FIXString("20151027"), FIXString("18:41:52.698")},
	}

	for i, group := range f.groups {

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
