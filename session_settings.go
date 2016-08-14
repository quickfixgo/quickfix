package quickfix

import (
	"fmt"
	"strconv"
	"time"
)

//SessionSettings maps session settings to values with typed accessors.
type SessionSettings struct {
	settings map[string]string
}

//ConditionallyRequiredSetting indicates a missing setting
type ConditionallyRequiredSetting struct {
	Setting string
}

func (e ConditionallyRequiredSetting) Error() string {
	return fmt.Sprintf("Conditionally Required Setting: %v", e.Setting)
}

//IncorrectFormatForSetting indicates a setting that is incorrectly formatted
type IncorrectFormatForSetting struct {
	Setting, Value string
	Err            error
}

func (e IncorrectFormatForSetting) Error() string {
	return fmt.Sprintf("%q is invalid for %s", e.Value, e.Setting)
}

//Init initializes or resets SessionSettings
func (s *SessionSettings) Init() {
	s.settings = make(map[string]string)
}

//NewSessionSettings returns a newly initialized SessionSettings instance
func NewSessionSettings() *SessionSettings {
	s := &SessionSettings{}
	s.Init()

	return s
}

//Set assigns a value to a setting on SessionSettings.
func (s *SessionSettings) Set(setting string, val string) {
	//lazy init
	if s.settings == nil {
		s.Init()
	}

	s.settings[setting] = val
}

//HasSetting returns true if a setting is set, false if not
func (s *SessionSettings) HasSetting(setting string) bool {
	_, ok := s.settings[setting]
	return ok
}

//Setting is a settings string accessor. Returns an error if the setting is missing.
func (s *SessionSettings) Setting(setting string) (string, error) {
	val, ok := s.settings[setting]
	if !ok {
		return val, ConditionallyRequiredSetting{setting}
	}

	return val, nil
}

//IntSetting returns the requested setting parsed as an int.  Returns an errror if the setting is not set or cannot be parsed as an int.
func (s *SessionSettings) IntSetting(setting string) (val int, err error) {
	stringVal, err := s.Setting(setting)

	if err != nil {
		return
	}

	if val, err = strconv.Atoi(stringVal); err != nil {
		return val, IncorrectFormatForSetting{Setting: setting, Value: stringVal, Err: err}
	}

	return
}

//DurationSetting returns the requested setting parsed as a time.Duration.
//Returns an error if the setting is not set or cannot be parsed as a time.Duration.
func (s *SessionSettings) DurationSetting(setting string) (val time.Duration, err error) {
	stringVal, err := s.Setting(setting)

	if err != nil {
		return
	}

	if val, err = time.ParseDuration(stringVal); err != nil {
		return val, IncorrectFormatForSetting{Setting: setting, Value: stringVal, Err: err}
	}

	return
}

//BoolSetting returns the requested setting parsed as a boolean.  Returns an errror if the setting is not set or cannot be parsed as a bool.
func (s SessionSettings) BoolSetting(setting string) (bool, error) {
	stringVal, err := s.Setting(setting)

	if err != nil {
		return false, err
	}

	switch stringVal {
	case "Y", "y":
		return true, nil
	case "N", "n":
		return false, nil
	}

	return false, IncorrectFormatForSetting{Setting: setting, Value: stringVal}
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
