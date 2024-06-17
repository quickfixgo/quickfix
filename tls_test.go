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
	"os"
	"testing"

	"github.com/stretchr/testify/suite"

	"github.com/quickfixgo/quickfix/config"
)

type TLSTestSuite struct {
	suite.Suite
	settings                                   *Settings
	PrivateKeyFile, CertificateFile, CAFile    string
	PrivateKeyBytes, CertificateBytes, CABytes []byte
}

func TestTLSTestSuite(t *testing.T) {
	suite.Run(t, new(TLSTestSuite))
}

func (s *TLSTestSuite) SetupTest() {
	s.settings = NewSettings()
	s.PrivateKeyFile = "_test_data/localhost.key"
	s.CertificateFile = "_test_data/localhost.crt"
	s.CAFile = "_test_data/ca.crt"

	privateKeyBytes, err := os.ReadFile(s.PrivateKeyFile)
	s.Require().NoError(err)
	s.PrivateKeyBytes = privateKeyBytes

	certificateBytes, err := os.ReadFile(s.CertificateFile)
	s.Require().NoError(err)
	s.CertificateBytes = certificateBytes

	caBytes, err := os.ReadFile(s.CAFile)
	s.Require().NoError(err)
	s.CABytes = caBytes
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
	s.EqualError(err, "Conditionally Required Setting: SocketCertificateFile")

	s.SetupTest()
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	_, err = loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "Conditionally Required Setting: SocketPrivateKeyFile")
}

func (s *TLSTestSuite) TestLoadTLSInvalidKeyOrCert() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, "blah")
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, "foo")
	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "failed to load key pair: open foo: no such file or directory")
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
	s.Equal(tls.RequireAndVerifyClientCert, tlsConfig.ClientAuth)
}

func (s *TLSTestSuite) TestLoadTLSWithBadCA() {
	s.settings.GlobalSettings().Set(config.SocketPrivateKeyFile, s.PrivateKeyFile)
	s.settings.GlobalSettings().Set(config.SocketCertificateFile, s.CertificateFile)
	s.settings.GlobalSettings().Set(config.SocketCAFile, "bar")

	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "failed to read CA bundle: open bar: no such file or directory")
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

func (s *TLSTestSuite) TestLoadTLSWithOnlyCA() {
	s.settings.GlobalSettings().Set(config.SocketUseSSL, "Y")
	s.settings.GlobalSettings().Set(config.SocketCAFile, s.CAFile)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.NotNil(tlsConfig.RootCAs)
	s.NotNil(tlsConfig.ClientCAs)
}

func (s *TLSTestSuite) TestLoadTLSWithoutSSLWithOnlyCA() {
	s.settings.GlobalSettings().Set(config.SocketCAFile, s.CAFile)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.Nil(tlsConfig)
}

func (s *TLSTestSuite) TestLoadTLSAllowSkipClientCerts() {
	s.settings.GlobalSettings().Set(config.SocketUseSSL, "Y")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.Equal(tls.NoClientCert, tlsConfig.ClientAuth)
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

func (s *TLSTestSuite) TestLoadTLSBytesMissingKeyOrCert() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "Conditionally Required Setting: SocketCertificateBytes")

	s.SetupTest()
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)
	_, err = loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "Conditionally Required Setting: SocketPrivateKeyBytes")
}

func (s *TLSTestSuite) TestLoadTLSBytesInvalidKeyOrCert() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, []byte("not a cert"))
	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "failed to parse key pair: tls: failed to find any PEM data in certificate input")

	s.SetupTest()
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, []byte("not a key"))
	_, err = loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "failed to parse key pair: tls: failed to find any PEM data in key input")
}

func (s *TLSTestSuite) TestLoadTLSBytesNoCA() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NoError(err)
	s.NotNil(tlsConfig)

	s.Len(tlsConfig.Certificates, 1)
	s.Nil(tlsConfig.RootCAs)
	s.Nil(tlsConfig.ClientCAs)
	s.Equal(tls.RequireAndVerifyClientCert, tlsConfig.ClientAuth)
}

func (s *TLSTestSuite) TestLoadTLSBytesWithBadCA() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCABytes, []byte("bar"))

	_, err := loadTLSConfig(s.settings.GlobalSettings())
	s.NotNil(err)
	s.EqualError(err, "failed to parse CA bundle from raw bytes")
}

func (s *TLSTestSuite) TestLoadTLSBytesWithCA() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCABytes, s.CABytes)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.Len(tlsConfig.Certificates, 1)
	s.NotNil(tlsConfig.RootCAs)
	s.NotNil(tlsConfig.ClientCAs)
	s.Equal(tls.RequireAndVerifyClientCert, tlsConfig.ClientAuth)
}

func (s *TLSTestSuite) TestLoadTLSBytesWithOnlyCA() {
	s.settings.GlobalSettings().Set(config.SocketUseSSL, "Y")
	s.settings.GlobalSettings().SetRaw(config.SocketCABytes, s.CABytes)

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)

	s.NotNil(tlsConfig.RootCAs)
	s.NotNil(tlsConfig.ClientCAs)
}

func (s *TLSTestSuite) TestServerNameWithCertsFromBytes() {
	s.settings.GlobalSettings().SetRaw(config.SocketPrivateKeyBytes, s.PrivateKeyBytes)
	s.settings.GlobalSettings().SetRaw(config.SocketCertificateBytes, s.CertificateBytes)
	s.settings.GlobalSettings().Set(config.SocketServerName, "DummyServerNameWithCerts")

	tlsConfig, err := loadTLSConfig(s.settings.GlobalSettings())
	s.Nil(err)
	s.NotNil(tlsConfig)
	s.Equal("DummyServerNameWithCerts", tlsConfig.ServerName)
}
