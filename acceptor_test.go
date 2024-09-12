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
	"crypto/tls"
	"net"
	"testing"

	"github.com/quickfixgo/quickfix/config"

	proxyproto "github.com/pires/go-proxyproto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
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

func TestAcceptor_SetTLSConfig(t *testing.T) {
	sessionSettings := NewSessionSettings()
	sessionSettings.Set(config.BeginString, BeginStringFIX42)
	sessionSettings.Set(config.SenderCompID, "sender")
	sessionSettings.Set(config.TargetCompID, "target")

	genericSettings := NewSettings()

	genericSettings.GlobalSettings().Set("SocketAcceptPort", "5001")
	_, err := genericSettings.AddSession(sessionSettings)
	require.NoError(t, err)

	logger, err := NewNullLogFactory().Create()
	require.NoError(t, err)
	acceptor := &Acceptor{settings: genericSettings, globalLog: logger}
	defer acceptor.Stop()
	// example of a customized tls.Config that loads the certificates dynamically by the `GetCertificate` function
	// as opposed to the Certificates slice, that is static in nature, and is only populated once and needs application restart to reload the certs.
	customizedTLSConfig := tls.Config{
		Certificates: []tls.Certificate{},
		GetCertificate: func(*tls.ClientHelloInfo) (*tls.Certificate, error) {
			cert, err := tls.LoadX509KeyPair("_test_data/localhost.crt", "_test_data/localhost.key")
			if err != nil {
				return nil, err
			}
			return &cert, nil
		},
	}

	acceptor.SetTLSConfig(&customizedTLSConfig)
	assert.NoError(t, acceptor.Start())
	assert.Len(t, acceptor.listeners, 1)

	conn, err := tls.Dial("tcp", "localhost:5001", &tls.Config{
		InsecureSkipVerify: true,
	})
	require.NoError(t, err)
	assert.NotNil(t, conn)
	defer conn.Close()
}
