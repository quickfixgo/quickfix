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
	"fmt"
	"log"
	"os"
	"path"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/config"
)

type fileLog struct {
	eventLogger   *log.Logger
	messageLogger *log.Logger
	eventWriter   *rollingWriter
	messageWriter *rollingWriter
}

func (l fileLog) OnIncoming(msg []byte) {
	l.messageLogger.Print(string(msg))
}

func (l fileLog) OnOutgoing(msg []byte) {
	l.messageLogger.Print(string(msg))
}

func (l fileLog) OnEvent(msg string) {
	l.eventLogger.Print(msg)
}

func (l fileLog) OnEventf(format string, v ...interface{}) {
	l.eventLogger.Printf(format, v...)
}

type rollingConfig struct {
	maxSize    int
	maxBackups int
	maxAge     int
	compress   bool
}

type fileLogFactory struct {
	globalLogPath   string
	sessionLogPaths map[quickfix.SessionID]string
	globalConfig    rollingConfig
	sessionConfigs  map[quickfix.SessionID]rollingConfig
}

// NewLogFactory creates an instance of LogFactory that writes messages and events to file.
// The location of global and session log files is configured via FileLogPath.
// Optional rolling configuration can be set via FileLogMaxSize, FileLogMaxBackups, FileLogMaxAge, and FileLogCompress.
func NewLogFactory(settings *quickfix.Settings) (quickfix.LogFactory, error) {
	logFactory := fileLogFactory{}

	var err error
	globalSettings := settings.GlobalSettings()
	if logFactory.globalLogPath, err = globalSettings.Setting(config.FileLogPath); err != nil {
		return logFactory, err
	}

	// Read global rolling configuration
	logFactory.globalConfig = readRollingConfig(globalSettings)

	logFactory.sessionLogPaths = make(map[quickfix.SessionID]string)
	logFactory.sessionConfigs = make(map[quickfix.SessionID]rollingConfig)

	// SessionSettings() already merges global settings with session-specific settings
	for sid, sessionSettings := range settings.SessionSettings() {
		logPath, err := sessionSettings.Setting(config.FileLogPath)
		if err != nil {
			return logFactory, err
		}
		logFactory.sessionLogPaths[sid] = logPath
		// Read merged rolling configuration (global + session overrides)
		logFactory.sessionConfigs[sid] = readRollingConfig(sessionSettings)
	}

	return logFactory, nil
}

// readRollingConfig reads rolling configuration from settings.
// Returns default values (all zeros/false) if settings are not present.
func readRollingConfig(settings *quickfix.SessionSettings) rollingConfig {
	cfg := rollingConfig{}

	if settings.HasSetting(config.FileLogMaxSize) {
		if maxSize, err := settings.IntSetting(config.FileLogMaxSize); err == nil {
			cfg.maxSize = maxSize
		}
	}

	if settings.HasSetting(config.FileLogMaxBackups) {
		if maxBackups, err := settings.IntSetting(config.FileLogMaxBackups); err == nil {
			cfg.maxBackups = maxBackups
		}
	}

	if settings.HasSetting(config.FileLogMaxAge) {
		if maxAge, err := settings.IntSetting(config.FileLogMaxAge); err == nil {
			cfg.maxAge = maxAge
		}
	}

	if settings.HasSetting(config.FileLogCompress) {
		if compress, err := settings.BoolSetting(config.FileLogCompress); err == nil {
			cfg.compress = compress
		}
	}

	return cfg
}

func newFileLog(prefix string, logPath string, config rollingConfig) (fileLog, error) {
	l := fileLog{}

	eventLogName := path.Join(logPath, prefix+".event.current.log")
	messageLogName := path.Join(logPath, prefix+".messages.current.log")

	// Use rolling writer if any rolling option is enabled
	useRolling := config.maxSize > 0 || config.maxAge > 0 || config.maxBackups > 0 || config.compress

	if useRolling {
		// Create rolling writers
		eventWriter, err := newRollingWriter(eventLogName, config.maxSize, config.maxBackups, config.maxAge, config.compress)
		if err != nil {
			return l, err
		}

		messageWriter, err := newRollingWriter(messageLogName, config.maxSize, config.maxBackups, config.maxAge, config.compress)
		if err != nil {
			eventWriter.Close()
			return l, err
		}

		l.eventWriter = eventWriter
		l.messageWriter = messageWriter

		logFlag := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC
		l.eventLogger = log.New(eventWriter, "", logFlag)
		l.messageLogger = log.New(messageWriter, "", logFlag)
	} else {
		// Use regular file writers (backward compatible)
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
			eventFile.Close()
			return l, err
		}

		logFlag := log.Ldate | log.Ltime | log.Lmicroseconds | log.LUTC
		l.eventLogger = log.New(eventFile, "", logFlag)
		l.messageLogger = log.New(messageFile, "", logFlag)
	}

	return l, nil
}

func (f fileLogFactory) Create() (quickfix.Log, error) {
	return newFileLog("GLOBAL", f.globalLogPath, f.globalConfig)
}

func (f fileLogFactory) CreateSessionLog(sessionID quickfix.SessionID) (quickfix.Log, error) {
	logPath, ok := f.sessionLogPaths[sessionID]

	if !ok {
		return nil, fmt.Errorf("logger not defined for %v", sessionID)
	}

	prefix := sessionIDFilenamePrefix(sessionID)
	// Use session config (already merged with global in NewLogFactory)
	config := f.globalConfig
	if sessionConfig, hasSessionConfig := f.sessionConfigs[sessionID]; hasSessionConfig {
		config = sessionConfig
	}
	return newFileLog(prefix, logPath, config)
}
