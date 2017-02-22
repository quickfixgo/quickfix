package quickfix

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/quickfixgo/quickfix/config"
)

func loadTLSConfig(settings *SessionSettings) (tlsConfig *tls.Config, err error) {
	insecureSkipVerify := false
	if settings.HasSetting(config.SocketInsecureSkipVerify) {
		insecureSkipVerify, err = settings.BoolSetting(config.SocketInsecureSkipVerify)
		if err != nil {
			return
		}
	}

	if !settings.HasSetting(config.SocketPrivateKeyFile) && !settings.HasSetting(config.SocketCertificateFile) {
		if insecureSkipVerify {
			tlsConfig = defaultTLSConfig()
			tlsConfig.InsecureSkipVerify = true
		}
		return
	}

	privateKeyFile, err := settings.Setting(config.SocketPrivateKeyFile)
	if err != nil {
		return
	}

	certificateFile, err := settings.Setting(config.SocketCertificateFile)
	if err != nil {
		return
	}

	tlsConfig = defaultTLSConfig()
	tlsConfig.Certificates = make([]tls.Certificate, 1)
	tlsConfig.InsecureSkipVerify = insecureSkipVerify

	minVersion := "TLS12"
	if settings.HasSetting(config.SocketMinimumTLSVersion) {
		minVersion, err = settings.Setting(config.SocketMinimumTLSVersion)
		if err != nil {
			return
		}

		switch minVersion {
		case "SSL30":
			tlsConfig.MinVersion = tls.VersionSSL30
		case "TLS10":
			tlsConfig.MinVersion = tls.VersionTLS10
		case "TLS11":
			tlsConfig.MinVersion = tls.VersionTLS11
		case "TLS12":
			tlsConfig.MinVersion = tls.VersionTLS12
		}
	}


	if tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(certificateFile, privateKeyFile); err != nil {
		return
	}

	if !settings.HasSetting(config.SocketCAFile) {
		return
	}

	caFile, err := settings.Setting(config.SocketCAFile)
	if err != nil {
		return
	}

	pem, err := ioutil.ReadFile(caFile)
	if err != nil {
		return
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pem) {
		err = fmt.Errorf("Failed to parse %v", caFile)
		return
	}

	tlsConfig.RootCAs = certPool
	tlsConfig.ClientCAs = certPool
	tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert

	return
}

//defaultTLSConfig brought to you by https://github.com/gtank/cryptopasta/
func defaultTLSConfig() *tls.Config {
	return &tls.Config{
		// Avoids most of the memorably-named TLS attacks
		MinVersion: tls.VersionTLS12,
		// Causes servers to use Go's default ciphersuite preferences,
		// which are tuned to avoid attacks. Does nothing on clients.
		PreferServerCipherSuites: true,
		// Only use curves which have constant-time implementations
		CurvePreferences: []tls.CurveID{
			tls.CurveP256,
		},
	}
}
