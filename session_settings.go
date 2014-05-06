package quickfix

import (
	"errors"
	"fmt"
	"strconv"
)

//SessionSettings maps session settings to values with typed accessors.
type SessionSettings struct {
	settings map[string]string
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
		return val, errors.New("setting not found")
	}

	return val, nil
}

//IntSetting returns the requested setting parsed as an int.  Returns an errror if the setting is not set or cannot be parsed as an int.
func (s *SessionSettings) IntSetting(setting string) (int, error) {
	stringVal, err := s.Setting(setting)

	if err != nil {
		return 0, err
	}

	return strconv.Atoi(stringVal)
}

//BoolSetting returns the requested setting parsed as a boolean.  Returns an errror if the setting is not set or cannot be parsed as a bool.
func (s SessionSettings) BoolSetting(setting string) (bool, error) {
	stringVal, err := s.Setting(setting)

	if err != nil {
		return false, err
	}

	if stringVal == "Y" || stringVal == "y" {
		return true, nil
	}

	if stringVal == "N" || stringVal == "n" {
		return false, nil
	}

	return false, fmt.Errorf("%v cannot be parsed as a bool", stringVal)
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
