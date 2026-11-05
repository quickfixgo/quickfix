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
	"bufio"
	"context"
	"crypto/tls"
	"strings"
	"sync"
	"time"

	"golang.org/x/net/proxy"
)

// Initiator initiates connections and processes messages for all sessions.
type Initiator struct {
	app             Application
	settings        *Settings
	sessionSettings map[SessionID]*SessionSettings
	storeFactory    MessageStoreFactory
	logFactory      LogFactory
	globalLog       Log
	stopChan        chan interface{}
	wg              sync.WaitGroup
	sessions        map[SessionID]*session
	sessionFactory
}

// Start Initiator.
func (i *Initiator) Start() (err error) {
	i.stopChan = make(chan interface{})

	for sessionID, settings := range i.sessionSettings {
		// TODO: move into session factory.
		var tlsConfig *tls.Config
		if tlsConfig, err = loadTLSConfig(settings); err != nil {
			return
		}

		var dialer proxy.ContextDialer
		if dialer, err = loadDialerConfig(settings); err != nil {
			return
		}

		i.wg.Add(1)
		go func(sessID SessionID) {
			i.handleConnection(i.sessions[sessID], tlsConfig, dialer)
			i.wg.Done()
		}(sessionID)
	}
	return
}

// Stop Initiator.
func (i *Initiator) Stop() {
	select {
	case <-i.stopChan:
		// Closed already.
		return
	default:
	}
	close(i.stopChan)

	i.wg.Wait()

	for sessionID := range i.sessionSettings {
		err := UnregisterSession(sessionID)
		if err != nil {
			return
		}
	}
}

// NewInitiator creates and initializes a new Initiator.
func NewInitiator(app Application, storeFactory MessageStoreFactory, appSettings *Settings, logFactory LogFactory) (*Initiator, error) {
	i := &Initiator{
		app:             app,
		storeFactory:    storeFactory,
		settings:        appSettings,
		sessionSettings: appSettings.SessionSettings(),
		logFactory:      logFactory,
		sessions:        make(map[SessionID]*session),
		sessionFactory:  sessionFactory{true},
	}

	var err error
	i.globalLog, err = logFactory.Create()
	if err != nil {
		return i, err
	}

	for sessionID, s := range i.sessionSettings {
		session, err := i.createSession(sessionID, storeFactory, s, logFactory, app)
		if err != nil {
			return nil, err
		}

		i.sessions[sessionID] = session
	}

	return i, nil
}

// waitForInSessionTime returns true if the session is in session, false if the handler should stop.
func (i *Initiator) waitForInSessionTime(session *session) bool {
	inSessionTime := make(chan interface{})
	go func() {
		session.waitForInSessionTime()
		close(inSessionTime)
	}()

	select {
	case <-inSessionTime:
	case <-i.stopChan:
		return false
	}

	return true
}

// waitForReconnectInterval returns true if a reconnect should be re-attempted, false if handler should stop.
func (i *Initiator) waitForReconnectInterval(reconnectInterval time.Duration) bool {
	select {
	case <-time.After(reconnectInterval):
	case <-i.stopChan:
		return false
	}

	return true
}

func (i *Initiator) handleConnection(session *session, tlsConfig *tls.Config, dialer proxy.ContextDialer) {
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		session.run()
		wg.Done()
	}()

	defer func() {
		session.stop()
		wg.Wait()
	}()

	// SocketTimeout bounds the TLS handshake. Without it LogonTimeout does,
	// so a counterparty that accepts the connection and never answers cannot
	// hold a connection attempt open.
	handshakeTimeout := session.LogonTimeout
	if timeout, err := loadSocketTimeout(i.sessionSettings[session.sessionID]); err == nil && timeout > 0 {
		handshakeTimeout = timeout
	}

	connectionAttempt := 0

	for {
		if !i.waitForInSessionTime(session) {
			return
		}

		ctx, cancel := context.WithCancel(context.Background())

		// We start a goroutine in order to be able to cancel the dialer mid-connection
		// on receiving a stop signal to stop the initiator.
		go func() {
			select {
			case <-i.stopChan:
				cancel()
			case <-ctx.Done():
				return
			}
		}()

		var disconnected chan interface{}
		var msgIn chan fixIn
		var msgOut chan []byte
		var writer *stopWriter

		address := session.SocketConnectAddress[connectionAttempt%len(session.SocketConnectAddress)]
		session.log.OnEventf("Connecting to: %v", address)

		netConn, err := dialer.DialContext(ctx, "tcp", address)
		if err != nil {
			session.log.OnEventf("Failed to connect: %v", err)
			goto reconnect
		} else if tlsConfig != nil {
			// Unless InsecureSkipVerify is true, server name config is required for TLS
			// to verify the received certificate
			if !tlsConfig.InsecureSkipVerify && len(tlsConfig.ServerName) == 0 {
				serverName := address
				if c := strings.LastIndex(serverName, ":"); c > 0 {
					serverName = serverName[:c]
				}
				tlsConfig.ServerName = serverName
			}
			// The handshake ends with the dial context, so Stop cancels it.
			tlsConn := tls.Client(netConn, tlsConfig)
			handshakeCtx, cancelHandshake := context.WithTimeout(ctx, handshakeTimeout)
			err = tlsConn.HandshakeContext(handshakeCtx)
			cancelHandshake()
			if err != nil {
				session.log.OnEventf("Failed handshake: %v", err)
				closeTransport(netConn, session.log)
				goto reconnect
			}
			netConn = tlsConn
		}

		msgIn = make(chan fixIn, session.InChanCapacity)
		msgOut = make(chan []byte)
		if err := session.connect(msgIn, msgOut); err != nil {
			session.log.OnEventf("Failed to initiate: %v", err)
			closeTransport(netConn, session.log)
			goto reconnect
		}

		go readLoop(newParser(bufio.NewReader(netConn)), msgIn, session.log)
		disconnected = make(chan interface{})
		writer = &stopWriter{conn: netConn, timeout: session.LogoutTimeout}
		go func() {
			writeLoop(writer, msgOut, session.log)
			if err := netConn.Close(); err != nil {
				session.log.OnEvent(err.Error())
			}
			close(disconnected)
		}()

		// This ensures we properly cleanup the goroutine and context used for
		// dial cancelation after successful connection.
		cancel()

		select {
		case <-disconnected:
		case <-i.stopChan:
			// A write has no deadline of its own, and the session loop
			// waits for the writer to take each message: a counterparty
			// that stops reading would hold a write, the session loop
			// behind it, and so this stop, with the connection open. From
			// here every write gets LogoutTimeout. A write still not taken
			// then closes the connection, and the session can stop.
			writer.stop()
			session.stop()
			wg.Wait()
			// The writer may still be on its last message (a Logout, or
			// the answer to the counterparty's) or closing the connection.
			// Give it LogoutTimeout, then close the transport under it, so
			// Stop never returns while the connection is open.
			timer := time.NewTimer(session.LogoutTimeout)
			select {
			case <-disconnected:
			case <-timer.C:
			}
			timer.Stop()
			closeTransport(netConn, session.log)
			return
		}

	reconnect:
		cancel()

		connectionAttempt++
		session.log.OnEventf("Reconnecting in %v", session.ReconnectInterval)
		if !i.waitForReconnectInterval(session.ReconnectInterval) {
			return
		}
	}
}
