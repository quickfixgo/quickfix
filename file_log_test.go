package quickfix

import (
	"strings"
	"testing"
)

func TestFileLog_NewFileLogFactory(t *testing.T) {

	factory, err := NewFileLogFactory(NewSettings())

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

	factory, err = NewFileLogFactory(settings)

	if err != nil {
		t.Error("Did not expect error", err)
	}

	if factory == nil {
		t.Error("Should have returned factory")
	}
}
