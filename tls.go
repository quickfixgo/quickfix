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
	"crypto/x509"
	"fmt"
	"io/ioutil"

	"github.com/quickfixgo/quickfix/config"
)

func loadTLSConfig(settings *SessionSettings) (tlsConfig *tls.Config, err error) {
	allowSkipClientCerts := false
	if settings.HasSetting(config.SocketUseSSL) {
		allowSkipClientCerts, err = settings.BoolSetting(config.SocketUseSSL)
		if err != nil {
			return
		}
	}

	var serverName string
	if settings.HasSetting(config.SocketServerName) {
		serverName, err = settings.Setting(config.SocketServerName)
		if err != nil {
			return
		}
	}

	insecureSkipVerify := false
	if settings.HasSetting(config.SocketInsecureSkipVerify) {
		insecureSkipVerify, err = settings.BoolSetting(config.SocketInsecureSkipVerify)
		if err != nil {
			return
		}
	}

	if !settings.HasSetting(config.SocketPrivateKeyFile) && !settings.HasSetting(config.SocketCertificateFile) {
		if !allowSkipClientCerts {
			return
		}
	}

	tlsConfig = defaultTLSConfig()
	tlsConfig.ServerName = serverName
	tlsConfig.InsecureSkipVerify = insecureSkipVerify
	setMinVersionExplicit(settings, tlsConfig)

	if settings.HasSetting(config.SocketPrivateKeyFile) || settings.HasSetting(config.SocketCertificateFile) {

		var privateKeyFile string
		var certificateFile string

		privateKeyFile, err = settings.Setting(config.SocketPrivateKeyFile)
		if err != nil {
			return
		}

		certificateFile, err = settings.Setting(config.SocketCertificateFile)
		if err != nil {
			return
		}

		tlsConfig.Certificates = make([]tls.Certificate, 1)

		if tlsConfig.Certificates[0], err = tls.LoadX509KeyPair(certificateFile, privateKeyFile); err != nil {
			return
		}
	}

	if !allowSkipClientCerts {
		tlsConfig.ClientAuth = tls.RequireAndVerifyClientCert
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

	return
}

// defaultTLSConfig brought to you by https://github.com/gtank/cryptopasta/
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

func setMinVersionExplicit(settings *SessionSettings, tlsConfig *tls.Config) {
	if settings.HasSetting(config.SocketMinimumTLSVersion) {
		minVersion, err := settings.Setting(config.SocketMinimumTLSVersion)
		if err != nil {
			return
		}

		switch minVersion {
		case "SSL30":
			//nolint:staticcheck // SA1019 min version ok
			tlsConfig.MinVersion = tls.VersionSSL30
		case "TLS10":
			tlsConfig.MinVersion = tls.VersionTLS10
		case "TLS11":
			tlsConfig.MinVersion = tls.VersionTLS11
		case "TLS12":
			tlsConfig.MinVersion = tls.VersionTLS12
		}
	}
}
