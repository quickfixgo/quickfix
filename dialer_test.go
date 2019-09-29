package quickfix

import (
	"net"
	"testing"
	"time"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/suite"
)

type DialerTestSuite struct {
	suite.Suite
	settings *Settings
}

func TestDialerTestSuite(t *testing.T) {
	suite.Run(t, new(DialerTestSuite))
}

func (s *DialerTestSuite) SetupTest() {
	s.settings = NewSettings()
}

func (s *DialerTestSuite) TestLoadDialerNoSettings() {
	dialer, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().Nil(err)

	stdDialer, ok := dialer.(*net.Dialer)
	s.Require().True(ok)
	s.Require().NotNil(stdDialer)
	s.Zero(stdDialer.Timeout)
}

func (s *DialerTestSuite) TestLoadDialerWithTimeout() {
	s.settings.GlobalSettings().Set(config.SocketTimeout, "10s")
	dialer, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().Nil(err)

	stdDialer, ok := dialer.(*net.Dialer)
	s.Require().True(ok)
	s.Require().NotNil(stdDialer)
	s.EqualValues(10*time.Second, stdDialer.Timeout)
}

func (s *DialerTestSuite) TestLoadDialerInvalidProxy() {
	s.settings.GlobalSettings().Set(config.ProxyType, "totallyinvalidproxytype")
	_, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().NotNil(err)
}

func (s *DialerTestSuite) TestLoadDialerSocksProxy() {
	s.settings.GlobalSettings().Set(config.ProxyType, "socks")
	s.settings.GlobalSettings().Set(config.ProxyHost, "localhost")
	s.settings.GlobalSettings().Set(config.ProxyPort, "31337")
	dialer, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().Nil(err)
	s.Require().NotNil(dialer)

	_, ok := dialer.(*net.Dialer)
	s.Require().False(ok)
}

func (s *DialerTestSuite) TestLoadDialerSocksProxyInvalidHost() {
	s.settings.GlobalSettings().Set(config.ProxyType, "socks")
	s.settings.GlobalSettings().Set(config.ProxyPort, "31337")
	_, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().NotNil(err)
}

func (s *DialerTestSuite) TestLoadDialerSocksProxyInvalidPort() {
	s.settings.GlobalSettings().Set(config.ProxyType, "socks")
	s.settings.GlobalSettings().Set(config.ProxyHost, "localhost")
	_, err := loadDialerConfig(s.settings.GlobalSettings())
	s.Require().NotNil(err)
}
