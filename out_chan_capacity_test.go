// Copyright (c) quickfixengine.org  All rights reserved.
//
// This file may be distributed under the terms of the quickfixengine.org
// license as defined by quickfixengine.org and appearing in the file
// LICENSE included in the packaging of this file.
//
// This file is provided AS IS with NO WARRANTY OF ANY KIND, INCLUDING
// THE WARRANTY OF DESIGN, MERCHANTABILITY AND FITNESS FOR A
// PARTICULAR PURPOSE.
//
// See http://www.quickfixengine.org/LICENSE for licensing information.
//
// Contact ask@quickfixengine.org if any conditions of this licensing
// are not clear to you.

package quickfix

import (
	"bytes"
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"

	"github.com/quickfixgo/quickfix/config"
	"github.com/quickfixgo/quickfix/internal"
)

func newOutChanTestSession(capacity int) (*session, chan []byte) {
	out := make(chan []byte, capacity)
	s := &session{
		log:          nullLog{},
		messageOut:   out,
		messageEvent: make(chan bool, 1),
		stateTimer:   internal.NewEventTimer(func() {}),
	}
	s.OutChanCapacity = capacity
	s.HeartBtInt = time.Minute
	return s, out
}

// woken reports whether the session has been told to come round, and consumes the notice.
func woken(s *session) bool {
	select {
	case <-s.messageEvent:
		return true
	default:
		return false
	}
}

func TestSendQueuedWithAnUnbufferedChannelComesRoundAtOnce(t *testing.T) {
	s, _ := newOutChanTestSession(0)
	defer s.stateTimer.Stop()

	s.toSend = [][]byte{[]byte("1")}
	s.sendQueued(false)

	assert.Len(t, s.toSend, 1, "no writer stands at the channel, so the message stays queued")
	assert.True(t, woken(s), "an unbuffered channel keeps today's behaviour: the session comes round at once")
}

func TestSendQueuedWaitsForTheWriterWhenTheBufferIsFull(t *testing.T) {
	s, out := newOutChanTestSession(2)
	defer s.stateTimer.Stop()

	s.toSend = [][]byte{[]byte("1"), []byte("2"), []byte("3")}
	s.sendQueued(false)
	require.Len(t, s.toSend, 1, "two messages fit in the buffer and the third waits")
	assert.False(t, woken(s), "nothing the session does makes room, so it must not come round")

	s.sendQueued(false)
	assert.False(t, woken(s), "coming round again before the writer took anything changes nothing")

	<-out
	s.messageTaken()
	assert.True(t, woken(s), "the writer's take wakes the session")
	s.sendQueued(false)
	assert.Empty(t, s.toSend)

	<-out
	s.messageTaken()
	assert.False(t, woken(s), "a take that no one waits for wakes nothing")
}

func TestWriteLoopTellsEachTake(t *testing.T) {
	out := make(chan []byte, 3)
	out <- []byte("a")
	out <- []byte("b")
	out <- []byte("c")
	close(out)

	takes := 0
	var written bytes.Buffer
	writeLoop(&written, out, nullLog{}, func() { takes++ })

	assert.Equal(t, 3, takes)
	assert.Equal(t, "abc", written.String())
}

type slowWriter struct {
	delay   time.Duration
	written [][]byte
}

func (w *slowWriter) Write(p []byte) (int, error) {
	time.Sleep(w.delay)
	w.written = append(w.written, append([]byte(nil), p...))
	return len(p), nil
}

func TestABufferedChannelDeliversEveryMessageInOrderToASlowWriter(t *testing.T) {
	const messages = 2000
	s, out := newOutChanTestSession(4)
	defer s.stateTimer.Stop()

	w := &slowWriter{delay: 50 * time.Microsecond}
	done := make(chan struct{})
	go func() {
		writeLoop(w, out, nullLog{}, s.messageTaken)
		close(done)
	}()

	for i := 0; i < messages; i++ {
		s.toSend = append(s.toSend, []byte(fmt.Sprintf("%05d", i)))
	}

	// The session goroutine's part of the run loop: come round when told to.
	rounds := 0
	s.sendQueued(false)
	deadline := time.After(30 * time.Second)
	for len(s.toSend) > 0 {
		select {
		case <-s.messageEvent:
			rounds++
			s.sendQueued(false)
		case <-deadline:
			t.Fatalf("%d messages still queued after %d rounds", len(s.toSend), rounds)
		}
	}
	close(out)
	<-done

	require.Len(t, w.written, messages)
	for i, msg := range w.written {
		require.Equal(t, fmt.Sprintf("%05d", i), string(msg))
	}
	assert.LessOrEqual(t, rounds, messages, "each round follows a take, so there are no more rounds than messages")
}

func TestOutChanCapacitySetting(t *testing.T) {
	build := func(value string) (*session, error) {
		settings := NewSessionSettings()
		settings.Set(config.BeginString, BeginStringFIX42)
		settings.Set(config.SenderCompID, "S")
		settings.Set(config.TargetCompID, "T")
		settings.Set(config.HeartBtInt, "30")
		if value != "" {
			settings.Set(config.OutChanCapacity, value)
		}
		return sessionFactory{}.newSession(SessionID{BeginString: BeginStringFIX42, SenderCompID: "S", TargetCompID: "T"},
			NewMemoryStoreFactory(), settings, NewNullLogFactory(), nil)
	}

	s, err := build("")
	require.NoError(t, err)
	assert.Equal(t, 0, s.OutChanCapacity, "unset, the channel stays unbuffered")

	s, err = build("256")
	require.NoError(t, err)
	assert.Equal(t, 256, s.OutChanCapacity)

	_, err = build("-1")
	assert.Error(t, err, "a negative capacity is refused")
}
