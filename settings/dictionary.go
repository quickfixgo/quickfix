package settings

import (
	"fmt"
)

type Dictionary interface {
	GetString(setting Setting) (string, bool)
	GetInt(setting Setting) (int, bool)
	GetBool(setting Setting) (bool, bool)

	GetStringKeys() []Setting
	GetIntKeys() []Setting
	GetBoolKeys() []Setting
}

type MutableDictionary interface {
	Dictionary
	SetString(setting Setting, val string) MutableDictionary
	SetInt(setting Setting, val int) MutableDictionary
	SetBool(setting Setting, val bool) MutableDictionary

	Overlay(dict Dictionary)
}

type dictionary struct {
	stringSettings map[Setting]interface{}
	intSettings    map[Setting]interface{}
	boolSettings   map[Setting]interface{}
}

func (d *dictionary) GetString(setting Setting) (string, bool) {
	if obj, ok := d.stringSettings[setting]; ok {
		switch typedVal := obj.(type) {
		case string:
			return typedVal, true
		default:
			panic(fmt.Sprintf("Expected to find string for %v, got %v", setting, obj))
		}
	}

	return "", false
}

func getKeys(settings map[Setting]interface{}) []Setting {
	keys := make([]Setting, 0, len(settings))

	for k := range settings {
		keys = append(keys, k)
	}

	return keys
}

func (d *dictionary) GetStringKeys() []Setting {
	return getKeys(d.stringSettings)
}

func (d *dictionary) SetString(setting Setting, val string) MutableDictionary {
	d.stringSettings[setting] = val
	return d
}

func (d *dictionary) GetInt(setting Setting) (int, bool) {
	if obj, ok := d.intSettings[setting]; ok {
		switch typedVal := obj.(type) {
		case int:
			return typedVal, true
		default:
			panic(fmt.Sprintf("expected to find int for %v, got %v", setting, obj))
		}
	}

	return 0, false
}

func (d *dictionary) GetBool(setting Setting) (bool, bool) {
	if obj, ok := d.boolSettings[setting]; ok {
		switch typedVal := obj.(type) {
		case bool:
			return typedVal, true
		default:
			panic(fmt.Sprintf("expected to find bool for %v, got %v", setting, obj))
		}
	}

	return false, false
}

func (d *dictionary) GetIntKeys() []Setting {
	return getKeys(d.intSettings)
}

func (d *dictionary) GetBoolKeys() []Setting {
	return getKeys(d.boolSettings)
}

func (d *dictionary) SetInt(setting Setting, val int) MutableDictionary {
	d.intSettings[setting] = val
	return d
}

func (d *dictionary) SetBool(setting Setting, val bool) MutableDictionary {
	d.boolSettings[setting] = val
	return d
}

func (d *dictionary) Overlay(overlay Dictionary) {
	for _, key := range overlay.GetStringKeys() {
		val, _ := overlay.GetString(key)
		d.SetString(key, val)
	}

	for _, key := range overlay.GetIntKeys() {
		val, _ := overlay.GetInt(key)
		d.SetInt(key, val)
	}

	for _, key := range overlay.GetBoolKeys() {
		val, _ := overlay.GetBool(key)
		d.SetBool(key, val)
	}
}

func CloneDictionary(d Dictionary) MutableDictionary {
	clone := NewDictionary()
	clone.Overlay(d)
	return clone
}

func NewDictionary() MutableDictionary {
	d := dictionary{}
	d.stringSettings = make(map[Setting]interface{})
	d.intSettings = make(map[Setting]interface{})
	d.boolSettings = make(map[Setting]interface{})

	return &d
}
