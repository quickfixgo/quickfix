package quickfix

import (
	"bytes"
	"testing"
	"time"
)

func TestMemoryStore_IncMsgSeqNum(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})

	var testCases = []struct {
		expectedNextSeqNum int
	}{
		{1}, {2}, {3},
	}

	for _, tc := range testCases {

		if store.NextSenderMsgSeqNum() != tc.expectedNextSeqNum {
			t.Errorf("Did not get expected sender seq num %v, got %v", tc.expectedNextSeqNum, store.NextSenderMsgSeqNum())
		}

		store.IncrNextSenderMsgSeqNum()

		if store.NextTargetMsgSeqNum() != tc.expectedNextSeqNum {
			t.Errorf("Did not get expected target seq num %v, got %v", tc.expectedNextSeqNum, store.NextTargetMsgSeqNum())
		}

		store.IncrNextTargetMsgSeqNum()
	}
}

func TestMemoryStore_SetMsgSeqNum(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	store.SetNextSenderMsgSeqNum(50)
	store.SetNextTargetMsgSeqNum(30)

	if store.NextSenderMsgSeqNum() != 50 {
		t.Errorf("Did not get expected sender seq num %v, got %v", 50, store.NextSenderMsgSeqNum())
	}

	if store.NextTargetMsgSeqNum() != 30 {
		t.Errorf("Did not get expected target seq num %v, got %v", 30, store.NextTargetMsgSeqNum())
	}
}

func TestMemoryStore_GetMessages(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})

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

func TestMemoryStoreFactory_Create(t *testing.T) {
	t0 := time.Now()
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	t1 := time.Now()

	if store.CreationTime().Before(t0) {
		t.Errorf("Expected %v to be before %v", t0, store.CreationTime())
	}

	if store.CreationTime().After(t1) {
		t.Errorf("Expected %v to be after %v", t1, store.CreationTime())
	}
}

func TestMemoryStore_Reset(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	store.SetNextSenderMsgSeqNum(50)
	store.SetNextTargetMsgSeqNum(30)

	store.SaveMessage(1, []byte("hello"))
	store.SaveMessage(2, []byte("cruel"))
	store.SaveMessage(3, []byte("world"))

	store.Reset()

	messages := store.GetMessages(1, 3)
	msg, ok := <-messages

	if ok {
		t.Error("Did not expect messages, got ", string(msg))
	}

	if store.NextSenderMsgSeqNum() != 1 {
		t.Error("SenderMsgSeqNum should reset")
	}

	if store.NextTargetMsgSeqNum() != 1 {
		t.Error("TargetMsgSeqNum should reset")
	}
}
