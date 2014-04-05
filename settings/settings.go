//Package settings for provides settings and configuration for quickfixgo.
package settings

import (
	"fmt"
)

//Setting is the type for QuickFIX/Go configuration settings.
type Setting string

func RequiredConfigurationMissing(setting Setting) error {
	return fmt.Errorf("missing configuration: %v", setting)
}

//QuickFIX/Go settings.
const (
	BeginString       Setting = "BeginString"
	SenderCompID      Setting = "SenderCompID"
	TargetCompID      Setting = "TargetCompID"
	SocketAcceptPort  Setting = "SocketAcceptPort"
	SocketConnectPort Setting = "SocketConnectPort"
	DefaultApplVerID  Setting = "DefaultApplVerID"
	DataDictionary    Setting = "DataDictionary"
	ResetOnLogon      Setting = "ResetOnLogon"
)
