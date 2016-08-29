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

func (s *TLSTestSuite) TestTLSForTesting() {
	s.settings.GlobalSettings().Set(config.SocketTLSForTesting, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.True(tlsConfig.InsecureSkipVerify)
}

func (s *TLSTestSuite) TestTLSForTestingDisableIfCertPresented() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketTLSForTesting, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.False(tlsConfig.InsecureSkipVerify)
	s.Len(tlsConfig.Certificates, 1)
}
