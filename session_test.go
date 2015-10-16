package quickfix

import (
	"github.com/quickfixgo/quickfix/fix"
	"github.com/quickfixgo/quickfix/fix/field"
	"github.com/quickfixgo/quickfix/fix/tag"
	"os"
	"testing"
	"time"
)

func getBuilder() MessageBuilder {
	builder := NewMessageBuilder()
	builder.Header().Set(fix.NewStringField(tag.BeginString, fix.BeginString_FIX40))
	builder.Header().Set(fix.NewStringField(tag.MsgType, "D"))
	return builder
}

func TestSession_CheckCorrectCompID(t *testing.T) {
	session := Session{}
	session.sessionID.TargetCompID = "TAR"
	session.sessionID.SenderCompID = "SND"

	var testCases = []struct {
		senderCompID *field.SenderCompIDField
		targetCompID *field.TargetCompIDField
		returnsError bool
		rejectReason int
	}{
		{returnsError: true, rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: field.NewSenderCompID("TAR"),
			returnsError: true,
			rejectReason: rejectReasonRequiredTagMissing},
		{senderCompID: field.NewSenderCompID("TAR"),
			targetCompID: field.NewTargetCompID("JCD"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: field.NewSenderCompID("JCD"),
			targetCompID: field.NewTargetCompID("SND"),
			returnsError: true,
			rejectReason: rejectReasonCompIDProblem},
		{senderCompID: field.NewSenderCompID("TAR"),
			targetCompID: field.NewTargetCompID("SND"),
			returnsError: false},
	}

	for _, tc := range testCases {
		builder := getBuilder()

		if tc.senderCompID != nil {
			builder.Header().Set(tc.senderCompID)
		}

		if tc.targetCompID != nil {
			builder.Header().Set(tc.targetCompID)
		}

		msgBytes, _ := builder.Build()
		msg, _ := parseMessage(msgBytes)
		err := session.checkCompID(*msg)

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
	session := Session{
		sessionID: SessionID{BeginString: "FIX.4.2"},
	}

	builder := getBuilder()

	//wrong value
	builder.Header().Set(fix.NewStringField(tag.BeginString, "FIX.4.4"))
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	err := session.checkBeginString(*msg)
	if err == nil {
		t.Error("Expected Error")
	}
	_ = err.(incorrectBeginString)

	builder.Header().Set(fix.NewStringField(tag.BeginString, session.sessionID.BeginString))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkBeginString(*msg)

	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckTargetTooHigh(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := Session{store: store}
	builder := getBuilder()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooHigh(*msg)

	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Expected required tag missing, got", err.RejectReason())
	}

	//too low
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 47))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)
	err = session.checkTargetTooHigh(*msg)

	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooHigh)

	//spot on
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 45))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooHigh(*msg)
	if err != nil {
		t.Error("Unexpected error", err)
	}
}

func TestSession_CheckSendingTime(t *testing.T) {
	session := Session{}
	builder := getBuilder()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	//missing sending time
	err := session.checkSendingTime(*msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Reject reason not expected, got ", err.RejectReason)
	}

	//sending time too late
	sendingTime := time.Now().Add(time.Duration(-200) * time.Second)
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(*msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason)
	}

	//future sending time
	sendingTime = time.Now().Add(time.Duration(200) * time.Second)
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(*msg)
	if err == nil {
		t.Error("Expected error")
	}
	if err.RejectReason() != rejectReasonSendingTimeAccuracyProblem {
		t.Error("Reject reason not expected, got ", err.RejectReason)
	}

	//sending time ok
	sendingTime = time.Now()
	builder.Header().Set(fix.NewUTCTimestampField(tag.SendingTime, sendingTime))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkSendingTime(*msg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}
}

func TestSession_CheckTargetTooLow(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	session := Session{store: store}

	builder := getBuilder()
	msgBytes, _ := builder.Build()
	msg, _ := parseMessage(msgBytes)

	store.SetNextTargetMsgSeqNum(45)

	//missing seq number
	err := session.checkTargetTooLow(*msg)
	if err == nil {
		t.Error("Expected error")
	}

	if err.RejectReason() != rejectReasonRequiredTagMissing {
		t.Error("Unexpected reject reason", err.RejectReason())
	}

	//too low
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 43))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooLow(*msg)
	if err == nil {
		t.Error("Expected error")
	}
	_ = err.(targetTooLow)

	//spot on
	builder.Header().Set(fix.NewIntField(tag.MsgSeqNum, 45))
	msgBytes, _ = builder.Build()
	msg, _ = parseMessage(msgBytes)

	err = session.checkTargetTooLow(*msg)
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

func (e *TestClient) ToAdmin(msg MessageBuilder, sessionID SessionID) {
	e.adminCalled = e.adminCalled + 1
}

func (e *TestClient) ToApp(msg MessageBuilder, sessionID SessionID) (err error) {
	e.appCalled = e.appCalled + 1
	return nil
}

func (e *TestClient) FromApp(msg Message, sessionID SessionID) (reject MessageRejectError) {
	return nil
}

func TestSession_CheckToAdminCalled(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	app := &TestClient{0, 0}
	cfg, err := os.Open("_test/fix44.cfg")
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	appSettings, err := ParseSettings(cfg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	session := Session{store: store, application: app, messageOut: otherEnd}
	session.toSend = make(chan MessageBuilder)
	session.sessionEvent = make(chan event)
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	logFactory, err := NewFileLogFactory(appSettings)
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	session.log, err = logFactory.Create()
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	builder := getBuilder()
	builder.Header().Set(field.NewMsgType("A"))
	session.send(builder)

	if app.adminCalled != 1 {
		t.Error("ToAdmin should have been called exactly once, instead was called", app.adminCalled, "times")
	}
	if app.appCalled != 0 {
		t.Error("ToApp should not have been called, instead was called", app.appCalled, "times")
	}
}

func TestSession_CheckToAppCalled(t *testing.T) {
	store, _ := NewMemoryStoreFactory().Create(SessionID{})
	app := &TestClient{0, 0}
	cfg, err := os.Open("_test/fix44.cfg")
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	appSettings, err := ParseSettings(cfg)
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	otherEnd := make(chan []byte)
	go func() {
		<-otherEnd
	}()

	session := Session{store: store, application: app, messageOut: otherEnd}
	session.toSend = make(chan MessageBuilder)
	session.sessionEvent = make(chan event)
	session.stateTimer = eventTimer{Task: func() { session.sessionEvent <- needHeartbeat }}
	session.peerTimer = eventTimer{Task: func() { session.sessionEvent <- peerTimeout }}

	logFactory, err := NewFileLogFactory(appSettings)
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	session.log, err = logFactory.Create()
	if err != nil {
		t.Error("Unexpected error ", err)
	}

	builder := getBuilder()
	session.send(builder)

	if app.appCalled != 1 {
		t.Error("Toapp should have been called exactly once, instead was called", app.appCalled, "times")
	}
	if app.adminCalled != 0 {
		t.Error("Toadmin should not have been called, instead was called", app.adminCalled, "times")
	}
}
