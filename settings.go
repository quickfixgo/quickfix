package quickfix

import (
	"bufio"
	"container/list"
	"fmt"
	"io"
	"regexp"
)

//The Settings type represents a collection of global and session settings.
type Settings struct {
	globalSettings  *SessionSettings
	sessionSettings *list.List
}

//Init initializes or resets a Settings instance
func (s *Settings) Init() {
	s.globalSettings = NewSessionSettings()
	s.sessionSettings = list.New()
}

func (s *Settings) lazyInit() {
	if s.globalSettings == nil {
		s.Init()
	}
}

//NewSettings creates a Settings instance
func NewSettings() *Settings {
	s := &Settings{}
	s.Init()
	return s
}

//ParseSettings creates and initializes a Settings instance with config parsed from a Reader.
//Returns error if the config is has parse errors
func ParseSettings(reader io.Reader) (*Settings, error) {
	s := NewSettings()

	scanner := bufio.NewScanner(reader)
	blankRegEx := regexp.MustCompile(`^\s*$`)
	commentRegEx := regexp.MustCompile(`^#.*`)
	defaultRegEx := regexp.MustCompile(`^\[DEFAULT\]\s*$`)
	sessionRegEx := regexp.MustCompile(`^\[SESSION\]\s*$`)
	settingRegEx := regexp.MustCompile(`^(.*)=(.*)$`)

	var settings *SessionSettings

	lineNumber := 0
	for scanner.Scan() {
		lineNumber++
		line := scanner.Text()

		switch {
		case commentRegEx.MatchString(line) || blankRegEx.MatchString(line):
			continue

		case defaultRegEx.MatchString(line):
			settings = s.GlobalSettings()

		case sessionRegEx.MatchString(line):
			if settings != nil && settings != s.GlobalSettings() {
				s.AddSession(settings)
			}
			settings = NewSessionSettings()

		case settingRegEx.MatchString(line):
			parts := settingRegEx.FindStringSubmatch(line)
			settings.Set(parts[1], parts[2])

		default:
			return s, fmt.Errorf("error parsing line %v", lineNumber)
		}
	}

	if err := scanner.Err(); err != nil {
		return s, err
	}

	if settings == nil || settings == s.GlobalSettings() {
		return s, fmt.Errorf("no sessions declared")
	}

	s.AddSession(settings)

	return s, nil
}

//GlobalSettings are default setting inherited by all session settings.
func (s *Settings) GlobalSettings() *SessionSettings {
	s.lazyInit()
	return s.globalSettings
}

//SessionSettings return all session settings overlaying globalsettings.
func (s *Settings) SessionSettings() []*SessionSettings {
	allSessionSettings := make([]*SessionSettings, s.sessionSettings.Len())

	i := 0
	for e := s.sessionSettings.Front(); e != nil; e = e.Next() {
		cloneSettings := s.globalSettings.clone()
		cloneSettings.overlay(e.Value.(*SessionSettings))
		allSessionSettings[i] = cloneSettings
		i++
	}

	return allSessionSettings
}

//AddSession adds Session Settings to Settings instance.
func (s *Settings) AddSession(sessionSettings *SessionSettings) {
	s.lazyInit()

	s.sessionSettings.PushBack(sessionSettings)
}
