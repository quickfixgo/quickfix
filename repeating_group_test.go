package quickfix

import (
	"bytes"
	"testing"
)

func TestRepeatingGroup_Write(t *testing.T) {
	var tests = []struct {
		groups   []Group
		expected []byte
	}{
		{[]Group{{RepeatingGroupField{Tag(1), NewFIXString("hello")}}},
			[]byte("11=hello")},
		{[]Group{{RepeatingGroupField{Tag(1), NewFIXString("hello")}, RepeatingGroupField{Tag(2), NewFIXString("world")}}},
			[]byte("11=hello2=world")},
		{[]Group{{RepeatingGroupField{Tag(1), NewFIXString("hello")}}, {RepeatingGroupField{Tag(1), NewFIXString("world")}}},
			[]byte("21=hello1=world")},
		{[]Group{{RepeatingGroupField{Tag(1), NewFIXString("hello")}, RepeatingGroupField{Tag(2), NewFIXString("world")}}, {RepeatingGroupField{Tag(1), NewFIXString("goodbye")}}},
			[]byte("21=hello2=world1=goodbye")},
	}

	for _, test := range tests {
		var f RepeatingGroup
		for _, group := range test.groups {
			f.AddGroup(group)
		}

		fieldBytes := f.Write()
		if !bytes.Equal(test.expected, fieldBytes) {
			t.Errorf("expected %s got %s", test.expected, fieldBytes)
		}
	}
}

func TestRepeatingGroup_Read(t *testing.T) {

	singleFieldTemplate := Group{RepeatingGroupField{Tag(1), new(FIXString)}}
	multiFieldTemplate := Group{RepeatingGroupField{Tag(1), new(FIXString)}, RepeatingGroupField{Tag(2), new(FIXString)}, RepeatingGroupField{Tag(3), new(FIXString)}}

	tests := []struct {
		groupTemplate      Group
		tv                 TagValues
		expectedGroupNum   int
		expectedGroupBytes [][][]byte
	}{
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("0")}}, 0, [][][]byte{}},
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("1")}, TagValue{Tag: Tag(1), Value: []byte("hello")}}, 1,
			[][][]byte{
				[][]byte{
					[]byte("hello")}}},
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("1")}, TagValue{Tag: Tag(1), Value: []byte("hello")}, TagValue{Tag: Tag(2), Value: []byte("not in group")}}, 1,
			[][][]byte{
				[][]byte{
					[]byte("hello")}}},
		{singleFieldTemplate, TagValues{TagValue{Value: []byte("2")}, TagValue{Tag: Tag(1), Value: []byte("hello")}, TagValue{Tag: Tag(1), Value: []byte("world")}}, 2,
			[][][]byte{
				[][]byte{[]byte("hello")},
				[][]byte{[]byte("world")}}},
		{multiFieldTemplate, TagValues{
			TagValue{Value: []byte("2")},
			TagValue{Tag: Tag(1), Value: []byte("hello")},
			TagValue{Tag: Tag(1), Value: []byte("goodbye")},
			TagValue{Tag: Tag(2), Value: []byte("cruel")},
			TagValue{Tag: Tag(3), Value: []byte("world")},
		}, 2,
			[][][]byte{
				[][]byte{[]byte("hello")},
				[][]byte{[]byte("goodbye"), []byte("cruel"), []byte("world")}}},
		{multiFieldTemplate, TagValues{
			TagValue{Value: []byte("3")},
			TagValue{Tag: Tag(1), Value: []byte("hello")},
			TagValue{Tag: Tag(1), Value: []byte("goodbye")},
			TagValue{Tag: Tag(2), Value: []byte("cruel")},
			TagValue{Tag: Tag(3), Value: []byte("world")},
			TagValue{Tag: Tag(1), Value: []byte("another")},
		}, 3,
			[][][]byte{
				[][]byte{[]byte("hello")},
				[][]byte{[]byte("goodbye"), []byte("cruel"), []byte("world")},
				[][]byte{[]byte("another")}}},
	}

	for _, test := range tests {
		f := RepeatingGroup{GroupTemplate: test.groupTemplate}

		f.Read(test.tv)
		if len(f.Groups) != test.expectedGroupNum {
			t.Errorf("expected %v groups, got %v", test.expectedGroupNum, len(f.Groups))
		}

		var allGroupBytes [][][]byte
		for _, group := range f.Groups {
			var groupBytes [][]byte
			for _, field := range group {
				groupBytes = append(groupBytes, field.Write())
			}

			allGroupBytes = append(allGroupBytes, groupBytes)
		}

		for i, groupBytes := range allGroupBytes {
			if len(groupBytes) != len(test.expectedGroupBytes[i]) {
				t.Errorf("Expected %v fields for group %v (%s) got %v (%s)", len(test.expectedGroupBytes[i]), i, test.expectedGroupBytes[i], len(groupBytes), groupBytes)
			}

			for j, fieldBytes := range groupBytes {
				if !bytes.Equal(fieldBytes, test.expectedGroupBytes[i][j]) {
					t.Errorf("Expected '%s' for field %v of group %v, got '%s'", test.expectedGroupBytes[i][j], j, i, fieldBytes)
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

	template := Group{RepeatingGroupField{Tag(269), new(FIXString)}, RepeatingGroupField{Tag(270), new(FIXString)}, RepeatingGroupField{Tag(271), new(FIXString)}, RepeatingGroupField{Tag(272), new(FIXString)}, RepeatingGroupField{Tag(273), new(FIXString)}}
	f := RepeatingGroup{GroupTemplate: template}
	err = msg.Body.GetField(Tag(268), &f)
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
	expectedGroupBytes := [][][]byte{
		[][]byte{[]byte("4"), []byte("0.07499"), []byte("20151027"), []byte("18:41:52.698")},
		[][]byte{[]byte("7"), []byte("0.07501"), []byte("20151027"), []byte("18:41:52.698")},
		[][]byte{[]byte("8"), []byte("0.07494"), []byte("20151027"), []byte("18:41:52.698")},
		[][]byte{[]byte("B"), []byte("60"), []byte("20151027"), []byte("18:41:52.698")},
	}

	for i, group := range f.Groups {
		if len(group) != len(expectedGroupTags[i]) {
			t.Errorf("expected %v tags in group %v got %v", len(expectedGroupTags[i]), i, len(group))
		}

		for j, field := range group {
			if field.Tag != expectedGroupTags[i][j] {
				t.Errorf("expected %v in group %v, field %v got %v", expectedGroupTags[i][j], i, j, field.Tag)
			}

			if !bytes.Equal(field.Write(), expectedGroupBytes[i][j]) {
				t.Errorf("Expected '%s' for field %v of group %v, got '%s'", expectedGroupBytes[i][j], j, i, field.Write())
			}

		}
	}
}
