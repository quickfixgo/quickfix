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
	"fmt"
	"strconv"
	"time"
)

// SessionSettings maps session settings to values with typed accessors.
type SessionSettings struct {
	settings map[string][]byte
}

// ConditionallyRequiredSetting indicates a missing setting.
type ConditionallyRequiredSetting struct {
	Setting string
}

func (e ConditionallyRequiredSetting) Error() string {
	return fmt.Sprintf("Conditionally Required Setting: %v", e.Setting)
}

// IncorrectFormatForSetting indicates a setting that is incorrectly formatted.
type IncorrectFormatForSetting struct {
	Setting string
	Value   []byte
	Err     error
}

func (e IncorrectFormatForSetting) Error() string {
	return fmt.Sprintf("%q is invalid for %s", e.Value, e.Setting)
}

// Init initializes or resets SessionSettings.
func (s *SessionSettings) Init() {
	s.settings = make(map[string][]byte)
}

// NewSessionSettings returns a newly initialized SessionSettings instance.
func NewSessionSettings() *SessionSettings {
	s := &SessionSettings{}
	s.Init()

	return s
}

// SetRaw assigns a value to a setting on SessionSettings.
func (s *SessionSettings) SetRaw(setting string, val []byte) {
	// Lazy init.
	if s.settings == nil {
		s.Init()
	}

	s.settings[setting] = val
}

// Set assigns a string value to a setting on SessionSettings.
func (s *SessionSettings) Set(setting string, val string) {
	// Lazy init
	if s.settings == nil {
		s.Init()
	}

	s.settings[setting] = []byte(val)
}

// HasSetting returns true if a setting is set, false if not.
func (s *SessionSettings) HasSetting(setting string) bool {
	_, ok := s.settings[setting]
	return ok
}

// RawSetting is a settings accessor that returns the raw byte slice value of
// the setting. Returns an error if the setting is missing.
func (s *SessionSettings) RawSetting(setting string) ([]byte, error) {
	val, ok := s.settings[setting]
	if !ok {
		return nil, ConditionallyRequiredSetting{Setting: setting}
	}

	return val, nil
}

// Setting is a settings string accessor. Returns an error if the setting is missing.
func (s *SessionSettings) Setting(setting string) (string, error) {
	val, err := s.RawSetting(setting)
	if err != nil {
		return "", err
	}

	return string(val), nil
}

// IntSetting returns the requested setting parsed as an int.  Returns an errror if the setting is not set or cannot be parsed as an int.
func (s *SessionSettings) IntSetting(setting string) (int, error) {
	rawVal, err := s.RawSetting(setting)
	if err != nil {
		return 0, err
	}

	if val, err := strconv.Atoi(string(rawVal)); err == nil {
		return val, nil
	}

	return 0, IncorrectFormatForSetting{Setting: setting, Value: rawVal, Err: err}
}

// DurationSetting returns the requested setting parsed as a time.Duration.
// Returns an error if the setting is not set or cannot be parsed as a time.Duration.
func (s *SessionSettings) DurationSetting(setting string) (time.Duration, error) {
	rawVal, err := s.RawSetting(setting)
	if err != nil {
		return 0, err
	}

	if val, err := time.ParseDuration(string(rawVal)); err == nil {
		return val, nil
	}

	return 0, IncorrectFormatForSetting{Setting: setting, Value: rawVal, Err: err}
}

// BoolSetting returns the requested setting parsed as a boolean.  Returns an error if the setting is not set or cannot be parsed as a bool.
func (s SessionSettings) BoolSetting(setting string) (bool, error) {
	rawVal, err := s.RawSetting(setting)
	if err != nil {
		return false, err
	}

	switch string(rawVal) {
	case "Y", "y":
		return true, nil
	case "N", "n":
		return false, nil
	}

	return false, IncorrectFormatForSetting{Setting: setting, Value: rawVal}
}

func (s *SessionSettings) overlay(overlay *SessionSettings) {
	for key, val := range overlay.settings {
		s.settings[key] = val
	}
}

func (s *SessionSettings) clone() *SessionSettings {
	sClone := NewSessionSettings()

	for k, v := range s.settings {
		sClone.settings[k] = v
	}

	return sClone
}
