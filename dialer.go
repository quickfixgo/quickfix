package quickfix

import (
	"fmt"
	"net"

	"github.com/quickfixgo/quickfix/config"
	"golang.org/x/net/proxy"
)

func loadDialerConfig(settings *SessionSettings) (dialer proxy.Dialer, err error) {
	stdDialer := &net.Dialer{}
	if settings.HasSetting(config.SocketTimeout) {
		if stdDialer.Timeout, err = settings.DurationSetting(config.SocketTimeout); err != nil {
			return
		}
	}
	dialer = stdDialer

	if !settings.HasSetting(config.ProxyType) {
		return
	}

	var proxyType string
	if proxyType, err = settings.Setting(config.ProxyType); err != nil {
		return
	}

	switch proxyType {
	case "socks":
		var proxyHost string
		var proxyPort int
		if proxyHost, err = settings.Setting(config.ProxyHost); err != nil {
			return
		} else if proxyPort, err = settings.IntSetting(config.ProxyPort); err != nil {
			return
		}

		proxyAuth := new(proxy.Auth)
		if settings.HasSetting(config.ProxyUser) {
			if proxyAuth.User, err = settings.Setting(config.ProxyUser); err != nil {
				return
			}
		}
		if settings.HasSetting(config.ProxyPassword) {
			if proxyAuth.Password, err = settings.Setting(config.ProxyPassword); err != nil {
				return
			}
		}

		dialer, err = proxy.SOCKS5("tcp", fmt.Sprintf("%s:%d", proxyHost, proxyPort), proxyAuth, dialer)
	default:
		err = fmt.Errorf("unsupported proxy type %s", proxyType)
	}
	return
}
