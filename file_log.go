package quickfix

import (
	"fmt"
	"log"
	"os"
	"path"

	"github.com/omni3x/quickfix/config"
)

type fileLog struct {
	eventLogger   *log.Logger
	messageLogger *log.Logger
}

const delim byte = 1

var (
	redacted = map[string]bool{"467": true, "2001": true, "2002": true, "554": true}
)

func (l fileLog) OnIncoming(msg []byte) {
	msgType := getMsgType(msg)
	if msgType == "W" || msgType == "X" {
		return // don't save price data
	}
	replaceDelimiter(msg)
	l.messageLogger.Print(string(msg))
}

func (l fileLog) OnOutgoing(msg []byte) {
	msgType := getMsgType(msg)
	if msgType == "W" || msgType == "X" {
		return // don't save price data
	} else if msgType == "D" || msgType == "A" { // NewOrderSingle: API KEY (467), SECRET (2001), PASS (2002) | LOGON: Password (554)
		redactTags(redacted, msg)
	}
	replaceDelimiter(msg)
	l.messageLogger.Print(string(msg))
}

// replaceDelimiter iterates through message and replaces every instance of the delimiter with pipe character: |
func replaceDelimiter(msg []byte) {
	for i, c := range msg {
		if c == delim {
			msg[i] = '|'
		}
	}
}

// getMsgType returns the Message Type of the inputted FIX Message as a string
func getMsgType(msg []byte) string {
	for i, c := range msg {
		if c != delim {
			continue // Always start parsing at a delimiter
		}
		for j, t := range msg[i:] {
			if t == '=' {
				newIdx := i + j
				parsedTag := string(msg[i+1 : newIdx])
				if parsedTag == "35" {
					for k, v := range msg[newIdx:] {
						if v == delim {
							return string(msg[newIdx+1 : i+j+k])
						}
					}
				}
				break // If tag does not match continue processing
			}
		}
	}
	return ""
}

// redactTags modifies the message to remove the FIX Values of the tags that exists as keys in the inputted tags
func redactTags(tags map[string]bool, msg []byte) {
	for i, c := range msg {
		if c != delim {
			continue // Always start parsing at a delimiter
		}
		for j, t := range msg[i:] {
			if t == '=' {
				newIdx := i + j
				parsedTag := string(msg[i+1 : newIdx])
				if _, ok := tags[parsedTag]; ok {
					newIdx++ // skip past = sign
					for v := msg[newIdx]; v != delim; {
						msg[newIdx] = '*'
						newIdx++
						v = msg[newIdx]
					}
					break // break instead of return to replace duplicate tags
				}
			}
		}
	}
}

func (l fileLog) OnEvent(msg string) {
	l.eventLogger.Print(msg)
}

func (l fileLog) OnEventf(format string, v ...interface{}) {
	l.eventLogger.Printf(format, v...)
}

type fileLogFactory struct {
	globalLogPath   string
	sessionLogPaths map[SessionID]string
}

//NewFileLogFactory creates an instance of LogFactory that writes messages and events to file.
//The location of global and session log files is configured via FileLogPath.
func NewFileLogFactory(settings *Settings) (LogFactory, error) {
	logFactory := fileLogFactory{}

	var err error
	if logFactory.globalLogPath, err = settings.GlobalSettings().Setting(config.FileLogPath); err != nil {
		return logFactory, err
	}

	logFactory.sessionLogPaths = make(map[SessionID]string)

	for sid, sessionSettings := range settings.SessionSettings() {
		logPath, err := sessionSettings.Setting(config.FileLogPath)
		if err != nil {
			return logFactory, err
		}
		logFactory.sessionLogPaths[sid] = logPath
	}

	return logFactory, nil
}

func newFileLog(prefix string, logPath string) (fileLog, error) {
	l := fileLog{}

	eventLogName := path.Join(logPath, prefix+".event.current.log")
	messageLogName := path.Join(logPath, prefix+".messages.current.log")

	if err := os.MkdirAll(logPath, os.ModePerm); err != nil {
		return l, err
	}

	fileFlags := os.O_RDWR | os.O_CREATE | os.O_APPEND
	eventFile, err := os.OpenFile(eventLogName, fileFlags, os.ModePerm)
	if err != nil {
		return l, err
	}

	messageFile, err := os.OpenFile(messageLogName, fileFlags, os.ModePerm)
	if err != nil {
		return l, err
	}

	logFlag := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC
	l.eventLogger = log.New(eventFile, "", logFlag)
	l.messageLogger = log.New(messageFile, "", logFlag)

	return l, nil
}

func (f fileLogFactory) Create() (Log, error) {
	return newFileLog("GLOBAL", f.globalLogPath)
}

func (f fileLogFactory) CreateSessionLog(sessionID SessionID) (Log, error) {
	logPath, ok := f.sessionLogPaths[sessionID]

	if !ok {
		return nil, fmt.Errorf("logger not defined for %v", sessionID)
	}

	prefix := sessionIDFilenamePrefix(sessionID)
	return newFileLog(prefix, logPath)
}
