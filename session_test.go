package quickfix

import (
	"github.com/quickfixgo/quickfix/enum"
	"testing"
	"time"
)

func buildMessage() Message {
	builder := Message{}
	builder.Init()
	builder.Header.SetField(tagBeginString, FIXString(enum.BeginStringFIX40))
	builder.Header.SetField(tagMsgType, FIXString("D"))
	return builder
}

func TestSession_CheckCorrectCompID(t *testing.T) {
	session := session{}
	session.sessionID.TargetCompID = "TAR"
	session.sessionID.SenderCompID = "SND"

	var testCases = []struct {
		senderCompID *FIXString
		targetCompID *FIXString
		returnsError bool
		rejectReason int
	}{
		{returnsError: true, rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: NewFIXString("TAR"),
			returnsError: true,
			rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: NewFIXString("TAR"),
			targetCompID: NewFIXString("JCD"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: NewFIXString("JCD"),
			targetCompID: NewFIXString("SND"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: NewFIXString("TAR"),
			targetCompID: NewFIXString("SND"),
			returnsError: false},
	}

	for _, tc := range testCases {
		builder := buildMessage()

		if tc.senderCompID != nil {
			builder.Header.SetField(tagSenderCompID, tc.senderCompID)
		}

		if tc.targetCompID != nil {
			builder.Header.SetField(tagTargetCompID, tc.targetCompID)
		}

		msgBytes, _ := builder.Build()
		msg, _ := parseMessage(msgBytes)
		err := session.checkCompID(msg)

		if err == nil {
			if tc.returnsError {
				t.Error("expected error")
			}

			return
		}

		if !tc.returnsError {
			t.Fatal("unexpected error", err)
		}

		if err.RejectReason() != tc.rejectReason {
			t.Errorf("expected %v got %v", tc.rejectReason, err.RejectReason())
		}
	}
}

func TestSession_CheckBeginString(t *testing.T) {
	session := session{
		sessionID: SessionID{BeginString: "FIX.4.2"},
	}

	builder := buildMessage()

	//wrong value
	builder.Header.SetField(tagBeginString, FIXString("FIX.4.4"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	err := session.checkBeginString(msg)
	if err == nil {
		t.Error("Expected Error")
	}
	_ = err.(incorrectBeginString)

	builder.Header.SetField(tagBeginString, FIXString(session.sessionID.BeginString))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkBeginString(msg)

	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckTargetTooHigh(t *testing.T) {
	store := new(memoryStore)
	session := session{store: store}
	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooHigh(msg)

	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Expected required tag missing, got", err.RejectReason())
	}

	//too low
	builder.Header.SetField(tagMsgSeqNum, FIXInt(47))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)
	err = session.checkTargetTooHigh(msg)

	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooHigh)

	//spot on
	builder.Header.SetField(tagMsgSeqNum, FIXInt(45))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooHigh(msg)
	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckSendingTime(t *testing.T) {
	session := session{}
	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	//missing sending time
	err := session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason())
	}

	//sending time ok
	sendingTime = time.Now()
	builder.Header.SetField(tagSendingTime, FIXUTCTimestamp{Value: sendingTime})
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(msg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestSession_CheckTargetTooLow(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := session{store: store}

	builder := buildMessage()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooLow(msg)
	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Unexpected reject reason", err.RejectReason())
	}

	//too low
	builder.Header.SetField(tagMsgSeqNum, FIXInt(43))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooLow(msg)
	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooLow)

	//spot on
	builder.Header.SetField(tagMsgSeqNum, FIXInt(45))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooLow(msg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

type TestClient struct {
	adminCalled int
	appCalled   int
}

func (e *TestClient) OnCreate(sessionID SessionID) {
}

func (e *TestClient) OnLogon(sessionID SessionID) {
}

func (e *TestClient) OnLogout(sessionID SessionID) {
}

func (e *TestClient) FromAdmin(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

func (e *TestClient) ToAdmin(msg Message, sessionID SessionID) {
	e.adminCalled = e.adminCalled + 1
}

func (e *TestClient) ToApp(msg Message, sessionID SessionID) (err error) {
	e.appCalled = e.appCalled + 1
	return nil
}

func (e *TestClient) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

func TestSession_CheckToAdminCalled(t *testing.T) {
	app := new(TestClient)
	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	session := session{
		store:       new(memoryStore),
		application: app,
		messageOut:  otherEnd,
		log:         nullLog{},
	}
	builder := buildMessage()
	builder.Header.SetField(tagMsgType, FIXString("A"))
	session.send(builder)

	if app.adminCalled != 1 {
		t.Error("ToAdmin should have been called exactly once, instead was called", app.adminCalled, "times")
	}
	if app.appCalled != 0 {
		t.Error("ToApp should not have been called, instead was called", app.appCalled, "times")
	}
}

func TestSession_CheckToAppCalled(t *testing.T) {
	app := new(TestClient)
	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	session := session{
		store:       new(memoryStore),
		application: app,
		messageOut:  otherEnd,
		log:         nullLog{}}
	builder := buildMessage()
	session.send(builder)

	if app.appCalled != 1 {
		t.Error("Toapp should have been called exactly once, instead was called", app.appCalled, "times")
	}
	if app.adminCalled != 0 {
		t.Error("Toadmin should not have been called, instead was called", app.adminCalled, "times")
	}
}

func TestSession_sendQueued(t *testing.T) {
	app := new(TestClient)
	otherEnd := make(chan []byte)
	go func() {
		for {
			_, ok := <-otherEnd
			if !ok {
				return
			}
		}
	}()
	session := session{
		store:       new(memoryStore),
		application: app,
		messageOut:  otherEnd,
		log:         nullLog{}}
	session.queueForSend(buildMessage())
	session.queueForSend(buildMessage())
	session.queueForSend(buildMessage())

	if len(session.toSend) != 3 {
		t.Errorf("Expected %v queued messages, got %v", 3, len(session.toSend))
	}

	var tests = []struct {
		sessionState
	}{
		{logonState{}},
		{logoutState{}},
	}

	for _, test := range tests {
		session.sessionState = test.sessionState
		session.sendQueued()

		if app.appCalled != 0 {
			t.Fatalf("session state %v should not allow send but sent %v times", session.sessionState, app.appCalled)
		}
	}

	session.sessionState = inSession{}

	session.sendQueued()

	if app.appCalled != 3 {
		t.Errorf("Toapp should have been called %v times, instead was called %v times", 3, app.appCalled)
	}

	if len(session.toSend) != 0 {
		t.Errorf("Expected no queued messages, got %v", len(session.toSend))
	}
}
