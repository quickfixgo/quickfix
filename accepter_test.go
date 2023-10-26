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
	"net"
	"testing"

	"github.com/quickfixgo/quickfix/config"

	proxyproto "github.com/armon/go-proxyproto"
	"github.com/stretchr/testify/assert"
)

func TestAcceptor_Start(t *testing.T) {
	sessionSettings := NewSessionSettings()
	sessionSettings.Set(config.BeginString, BeginStringFIX42)
	sessionSettings.Set(config.SenderCompID, "sender")
	sessionSettings.Set(config.TargetCompID, "target")

	settingsWithTCPProxy := NewSettings()
	settingsWithTCPProxy.GlobalSettings().Set("UseTCPProxy", "Y")

	settingsWithNoTCPProxy := NewSettings()
	settingsWithNoTCPProxy.GlobalSettings().Set("UseTCPProxy", "N")

	genericSettings := NewSettings()

	const (
		GenericListener = iota
		ProxyListener
	)

	acceptorStartTests := []struct {
		name         string
		settings     *Settings
		listenerType int
	}{
		{"with TCP proxy set", settingsWithTCPProxy, ProxyListener},
		{"with no TCP proxy set", settingsWithNoTCPProxy, GenericListener},
		{"no TCP proxy configuration set", genericSettings, GenericListener},
	}

	for _, tt := range acceptorStartTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.settings.GlobalSettings().Set("SocketAcceptPort", "5001")
			if _, err := tt.settings.AddSession(sessionSettings); err != nil {
				assert.Nil(t, err)
			}

			acceptor := &Acceptor{settings: tt.settings}
			if err := acceptor.Start(); err != nil {
				assert.NotNil(t, err)
			}
			assert.Len(t, acceptor.listeners, 1)

			for _, listener := range acceptor.listeners {
				if tt.listenerType == ProxyListener {
					_, ok := listener.(*proxyproto.Listener)
					assert.True(t, ok)
				}

				if tt.listenerType == GenericListener {
					_, ok := listener.(*net.TCPListener)
					assert.True(t, ok)
				}
			}

			acceptor.Stop()
		})
	}
}
