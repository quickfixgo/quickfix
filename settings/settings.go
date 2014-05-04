//Package settings provides settings and configuration for QuickFIX.
package settings

import (
	"container/list"
	"fmt"
)

func RequiredConfigurationMissing(setting string) error {
	return fmt.Errorf("missing configuration: %v", setting)
}

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

//New creates a Settings instance
func New() *Settings {
	s := &Settings{}
	s.Init()
	return s
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
