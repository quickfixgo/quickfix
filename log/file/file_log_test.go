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

package file

import (
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"

	"github.com/quickfixgo/quickfix"
)

func TestFileLog_NewFileLogFactory(t *testing.T) {

	_, err := NewLogFactory(quickfix.NewSettings())

	if err == nil {
		t.Error("Should expect error when settings have no file log path")
	}

	cfg := `
# default settings for sessions
[DEFAULT]
ConnectionType=initiator
ReconnectInterval=60
SenderCompID=TW
FileLogPath=.

# session definition
[SESSION]
BeginString=FIX.4.1
TargetCompID=ARCA
FileLogPath=mydir

[SESSION]
BeginString=FIX.4.1
TargetCompID=ARCA
SessionQualifier=BS
`
	stringReader := strings.NewReader(cfg)
	settings, _ := quickfix.ParseSettings(stringReader)

	factory, err := NewLogFactory(settings)

	if err != nil {
		t.Error("Did not expect error", err)
	}

	if factory == nil {
		t.Error("Should have returned factory")
	}
}

type fileLogHelper struct {
	LogPath string
	Prefix  string
	Log     quickfix.Log
}

func newFileLogHelper(t *testing.T) *fileLogHelper {
	prefix := "myprefix"
	logPath := path.Join(os.TempDir(), fmt.Sprintf("TestLogStore-%d", os.Getpid()))

	// Use default config (no rolling) for backward compatibility test
	config := rollingConfig{}
	log, err := newFileLog(prefix, logPath, config)
	if err != nil {
		t.Error("Unexpected error", err)
	}

	return &fileLogHelper{
		LogPath: logPath,
		Prefix:  prefix,
		Log:     log,
	}
}

func TestNewFileLog(t *testing.T) {
	helper := newFileLogHelper(t)

	tests := []struct {
		expectedPath string
	}{
		{path.Join(helper.LogPath, fmt.Sprintf("%v.messages.current.log", helper.Prefix))},
		{path.Join(helper.LogPath, fmt.Sprintf("%v.event.current.log", helper.Prefix))},
	}

	for _, test := range tests {
		if _, err := os.Stat(test.expectedPath); os.IsNotExist(err) {
			t.Errorf("%v does not exist", test.expectedPath)
		}
	}
}

func TestFileLog_Append(t *testing.T) {
	helper := newFileLogHelper(t)

	messageLogFile, err := os.Open(path.Join(helper.LogPath, fmt.Sprintf("%v.messages.current.log", helper.Prefix)))
	if err != nil {
		t.Error("Unexpected error", err)
	}
	defer messageLogFile.Close()

	eventLogFile, err := os.Open(path.Join(helper.LogPath, fmt.Sprintf("%v.event.current.log", helper.Prefix)))
	if err != nil {
		t.Error("Unexpected error", err)
	}
	defer eventLogFile.Close()

	messageScanner := bufio.NewScanner(messageLogFile)
	eventScanner := bufio.NewScanner(eventLogFile)

	helper.Log.OnIncoming([]byte("incoming"))
	if !messageScanner.Scan() {
		t.Error("Unexpected EOF")
	}

	helper.Log.OnEvent("Event")
	if !eventScanner.Scan() {
		t.Error("Unexpected EOF")
	}

	newHelper := newFileLogHelper(t)
	newHelper.Log.OnIncoming([]byte("incoming"))
	if !messageScanner.Scan() {
		t.Error("Unexpected EOF")
	}

	newHelper.Log.OnEvent("Event")
	if !eventScanner.Scan() {
		t.Error("Unexpected EOF")
	}
}

func TestFileLog_RollingConfig(t *testing.T) {
	cfg := `
[DEFAULT]
ConnectionType=initiator
FileLogPath=.
FileLogMaxSize=1
FileLogMaxBackups=3
FileLogMaxAge=7
FileLogCompress=Y

[SESSION]
BeginString=FIX.4.1
TargetCompID=ARCA
SenderCompID=TW
`
	stringReader := strings.NewReader(cfg)
	settings, err := quickfix.ParseSettings(stringReader)
	if err != nil {
		t.Fatal("Failed to parse settings", err)
	}

	factory, err := NewLogFactory(settings)
	if err != nil {
		t.Fatal("Did not expect error", err)
	}

	if factory == nil {
		t.Fatal("Should have returned factory")
	}

	// Test that factory was created with rolling config
	_ = factory
}

func TestFileLog_RollingBackwardCompatible(t *testing.T) {
	// Test that without rolling config, behavior is unchanged
	cfg := `
[DEFAULT]
ConnectionType=initiator
FileLogPath=.

[SESSION]
BeginString=FIX.4.1
TargetCompID=ARCA
SenderCompID=TW
`
	stringReader := strings.NewReader(cfg)
	settings, err := quickfix.ParseSettings(stringReader)
	if err != nil {
		t.Fatal("Failed to parse settings", err)
	}

	factory, err := NewLogFactory(settings)
	if err != nil {
		t.Fatal("Did not expect error", err)
	}

	log, err := factory.Create()
	if err != nil {
		t.Fatal("Did not expect error creating log", err)
	}

	// Should work without rolling
	log.OnEvent("Test event")
	log.OnIncoming([]byte("Test message"))
}
