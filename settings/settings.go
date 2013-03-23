//settings for quickfixgo
package settings

import (
	"fmt"
)

type Setting string

func RequiredConfigurationMissing(setting Setting) error {
	return fmt.Errorf("Missing Configuration: %v", setting)
}

const (
	BeginString       Setting = "BeginString"
	SenderCompID      Setting = "SenderCompID"
	TargetCompID      Setting = "TargetCompID"
	SocketAcceptPort  Setting = "SocketAcceptPort"
	SocketConnectPort Setting = "SocketConnectPort"
	DefaultApplVerID  Setting = "DefaultApplVerID"
)
