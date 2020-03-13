package quickfix

import (
	"net"
	"testing"

	"github.com/armon/go-proxyproto"
	"github.com/stretchr/testify/assert"
)

func TestAcceptor_Start(t *testing.T) {
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
		name string
		settings *Settings
		listenerType int
	}{
		{"with TCP proxy set", settingsWithTCPProxy, ProxyListener},
		{"with no TCP proxy set", settingsWithNoTCPProxy, GenericListener},
		{"no TCP proxy configuration set", genericSettings, GenericListener},
	}

	for _, tt := range acceptorStartTests {
		t.Run(tt.name, func(t *testing.T) {
			tt.settings.GlobalSettings().Set("SocketAcceptPort", "5001")

			acceptor := &Acceptor{settings: tt.settings}
			if err := acceptor.Start(); err != nil {
				assert.NotNil(t, err)
			}
			if tt.listenerType == ProxyListener {
				_, ok := acceptor.listener.(*proxyproto.Listener)
				assert.True(t, ok)
			}

			if tt.listenerType == GenericListener {
				_, ok := acceptor.listener.(*net.TCPListener)
				assert.True(t, ok)
			}

			acceptor.Stop()
		})
	}
}