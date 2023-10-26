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
	"bufio"
	"fmt"
	"os"
	"path"
	"strings"
	"testing"
)

func TestFileLog_NewFileLogFactory(t *testing.T) {

	_, err := NewFileLogFactory(NewSettings())

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
	settings, _ := ParseSettings(stringReader)

	factory, err := NewFileLogFactory(settings)

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
	Log     Log
}

func newFileLogHelper(t *testing.T) *fileLogHelper {
	prefix := "myprefix"
	logPath := path.Join(os.TempDir(), fmt.Sprintf("TestLogStore-%d", os.Getpid()))

	log, err := newFileLog(prefix, logPath)
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
