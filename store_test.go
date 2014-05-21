package quickfix

import (
	"bytes"
	"testing"
)

func TestMemoryStore_GetMessages(t *testing.T) {
	store := newMemoryStore()

	messages := store.GetMessages(1, 2)
	msg, ok := <-messages

	if ok != false {
		t.Error("Did not expect messages from empty store", msg)
	}

	store.SaveMessage(1, []byte("hello"))
	store.SaveMessage(2, []byte("cruel"))
	store.SaveMessage(3, []byte("world"))

	var testCases = []struct {
		beginSeqNo, endSeqNo int
		expectedBytes        [][]byte
	}{
		{beginSeqNo: 1, endSeqNo: 1, expectedBytes: [][]byte{[]byte("hello")}},
		{beginSeqNo: 1, endSeqNo: 2, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel")}},
		{beginSeqNo: 1, endSeqNo: 3, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel"), []byte("world")}},
		{beginSeqNo: 1, endSeqNo: 4, expectedBytes: [][]byte{[]byte("hello"), []byte("cruel"), []byte("world")}},
		{beginSeqNo: 2, endSeqNo: 3, expectedBytes: [][]byte{[]byte("cruel"), []byte("world")}},
		{beginSeqNo: 3, endSeqNo: 3, expectedBytes: [][]byte{[]byte("world")}},
		{beginSeqNo: 3, endSeqNo: 4, expectedBytes: [][]byte{[]byte("world")}},
		{beginSeqNo: 4, endSeqNo: 4, expectedBytes: [][]byte{}},
		{beginSeqNo: 4, endSeqNo: 10, expectedBytes: [][]byte{}},
	}

	for _, tc := range testCases {
		messages = store.GetMessages(tc.beginSeqNo, tc.endSeqNo)

		expected := tc.expectedBytes
		for {
			msg, ok = <-messages

			if len(expected) == 0 {
				if ok == true {
					t.Error("Did not expect additional messages", msg)
				}
				break
			}

			if ok != true {
				t.Error("Did not get messages, expected ", expected[0])
			}

			if !bytes.Equal(msg, expected[0]) {
				t.Error("Expected ", expected[0], " got ", msg)
			}

			expected = expected[1:]
		}
	}
}
