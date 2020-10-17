package quickfix

import (
	"crypto/tls"
	"testing"

	"github.com/quickfixgo/quickfix/config"
	"github.com/stretchr/testify/suite"
)

type TLSTestSuite struct {
	suite.Suite
	settings                                *Settings
	PrivateKeyFile, CertificateFile, CAFile string
}

func TestTLSTestSuite(t *testing.T) {
	suite.Run(t, new(TLSTestSuite))
}

func (s *TLSTestSuite) SetupTest() {
	s.settings = NewSettings()
	s.PrivateKeyFile = "_test_data/localhost.key"
	s.CertificateFile = "_test_data/localhost.crt"
	s.CAFile = "_test_data/ca.crt"
}

func (s *TLSTestSuite) TestLoadTLSNoSettings() {
	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(tlsConfig)
	s.Nil(err)
}

func (s *TLSTestSuite) TestLoadTLSMissingKeyOrCert() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)

	s.SetupTest()
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	_, err = loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
}

func (s *TLSTestSuite) TestLoadTLSInvalidKeyOrCert() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, "blah")
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, "foo")
	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
}

func (s *TLSTestSuite) TestLoadTLSNoCA() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.Len(tlsConfig.Certificates, 1)
	s.Nil(tlsConfig.RootCAs)
	s.Nil(tlsConfig.ClientCAs)
	s.Equal(tls.NoClientCert, tlsConfig.ClientAuth)
}

func (s *TLSTestSuite) TestLoadTLSWithBadCA() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketCAFile, "bar")

	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
}

func (s *TLSTestSuite) TestLoadTLSWithCA() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketCAFile, s.CAFile)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.Len(tlsConfig.Certificates, 1)
	s.NotNil(tlsConfig.RootCAs)
	s.NotNil(tlsConfig.ClientCAs)
	s.Equal(tls.RequireAndVerifyClientCert, tlsConfig.ClientAuth)
}

func (s *TLSTestSuite) TestServerNameUseSSL() {
	s.settings.GlobalSettings().Set(config.SocketUseSSL, "Y")
	s.settings.GlobalSettings().Set(config.SocketServerName, "DummyServerNameUseSSL")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal("DummyServerNameUseSSL", tlsConfig.ServerName)
}

func (s *TLSTestSuite) TestServerNameWithCerts() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketServerName, "DummyServerNameWithCerts")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal("DummyServerNameWithCerts", tlsConfig.ServerName)
}

func (s *TLSTestSuite) TestInsecureSkipVerify() {
	s.settings.GlobalSettings().Set(config.SocketInsecureSkipVerify, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.Nil(tlsConfig)
}

func (s *TLSTestSuite) TestInsecureSkipVerifyWithUseSSL() {
	s.settings.GlobalSettings().Set(config.SocketUseSSL, "Y")
	s.settings.GlobalSettings().Set(config.SocketInsecureSkipVerify, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.True(tlsConfig.InsecureSkipVerify)
}

func (s *TLSTestSuite) TestInsecureSkipVerifyAndCerts() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketInsecureSkipVerify, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.True(tlsConfig.InsecureSkipVerify)
	s.Len(tlsConfig.Certificates, 1)
}

func (s *TLSTestSuite) TestMinimumTLSVersion() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)

	// SSL30
	s.settings.GlobalSettings().Set(config.SocketMinimumTLSVersion, "SSL30")
	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())

	s.Nil(err)
	s.NotNil(tlsConfig)
	//nolint:staticcheck
	s.Equal(tlsConfig.MinVersion, uint16(tls.VersionSSL30))

	// TLS10
	s.settings.GlobalSettings().Set(config.SocketMinimumTLSVersion, "TLS10")
	tlsConfig, err = loadTLSConfig(s.settings.GlobalSettings())

	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal(tlsConfig.MinVersion, uint16(tls.VersionTLS10))

	// TLS11
	s.settings.GlobalSettings().Set(config.SocketMinimumTLSVersion, "TLS11")
	tlsConfig, err = loadTLSConfig(s.settings.GlobalSettings())

	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal(tlsConfig.MinVersion, uint16(tls.VersionTLS11))

	// TLS12
	s.settings.GlobalSettings().Set(config.SocketMinimumTLSVersion, "TLS12")
	tlsConfig, err = loadTLSConfig(s.settings.GlobalSettings())

	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal(tlsConfig.MinVersion, uint16(tls.VersionTLS12))
}
