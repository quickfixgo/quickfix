package settings

import(
    "fmt"
    )

type ApplicationSettings interface {
  AddSession(name string, dict Dictionary ) error
  GetSessionSettings() map[string] Dictionary

  GetGlobalSettings() Dictionary
}

type applicationSettings struct {
  globalSettings Dictionary
  sessionSettings map[string] Dictionary
}

func (s * applicationSettings) GetGlobalSettings() Dictionary {
  return s.globalSettings
}

func (s * applicationSettings) GetSessionSettings() map[string] Dictionary {
  allSessionSettings:=make(map[string] Dictionary)

  for name,sessionSettings:=range s.sessionSettings {
    cloneSettings:=CloneDictionary(s.globalSettings)
    cloneSettings.Overlay(sessionSettings)
    allSessionSettings[name]=cloneSettings
  }

  return allSessionSettings
}

func (s * applicationSettings) AddSession(name string, sessionSettings Dictionary) error {
  if _, sessionExists:=s.sessionSettings[name]; sessionExists {
    return fmt.Errorf("Duplicate Session %v ", name)
  }

  s.sessionSettings[name]=CloneDictionary(sessionSettings)
  return nil
}


func NewApplicationSettings(globalSettings Dictionary) ApplicationSettings {
  s:=applicationSettings{globalSettings: CloneDictionary(globalSettings),
      sessionSettings: make(map[string] Dictionary)}

  return &s
}
