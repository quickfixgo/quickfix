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
	"context"
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/tls"
	"crypto/x509"
	"crypto/x509/pkix"
	"errors"
	"fmt"
	"io"
	"math/big"
	"net"
	"os"
	"strings"
	"sync"
	"sync/atomic"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

// stallingConn stands in for a connection to a counterparty that stops
// reading: once stalled, every write blocks until the connection is closed.
type stallingConn struct {
	net.Conn
	stalled   atomic.Bool
	closed    chan struct{}
	closeOnce sync.Once
}

func (c *stallingConn) Write(b []byte) (int, error) {
	if c.stalled.Load() {
		<-c.closed
		return 0, net.ErrClosed
	}
	return c.Conn.Write(b)
}

func (c *stallingConn) Close() error {
	c.closeOnce.Do(func() { close(c.closed) })
	return c.Conn.Close()
}

func (c *stallingConn) isClosed() bool {
	select {
	case <-c.closed:
		return true
	default:
		return false
	}
}

type stallingDialer struct {
	conns chan *stallingConn
}

func (d stallingDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	var dialer net.Dialer
	conn, err := dialer.DialContext(ctx, network, address)
	if err != nil {
		return nil, err
	}
	stalling := &stallingConn{Conn: conn, closed: make(chan struct{})}
	d.conns <- stalling
	return stalling, nil
}

type logonRecorder struct {
	loggedOn  chan SessionID
	loggedOut chan SessionID
}

func (a logonRecorder) OnCreate(SessionID) {}
func (a logonRecorder) OnLogon(sessionID SessionID) {
	select {
	case a.loggedOn <- sessionID:
	default:
	}
}
func (a logonRecorder) OnLogout(sessionID SessionID) {
	select {
	case a.loggedOut <- sessionID:
	default:
	}
}
func (a logonRecorder) ToAdmin(*Message, SessionID)                      {}
func (a logonRecorder) ToApp(*Message, SessionID) error                  { return nil }
func (a logonRecorder) FromAdmin(*Message, SessionID) MessageRejectError { return nil }
func (a logonRecorder) FromApp(*Message, SessionID) MessageRejectError   { return nil }

// A counterparty that stops reading leaves the writer blocked on the Logout
// with no deadline, while the session loop times the Logout out and exits.
// Stop must not return with that connection still open: the caller takes a
// returned Stop to mean the session is gone, and may log the same session on
// again.
func TestInitiatorStopClosesAConnectionBlockedOnLogout(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	port := listener.Addr().(*net.TCPAddr).Port
	require.NoError(t, listener.Close())

	acceptorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
SocketAcceptPort=%d
[SESSION]
BeginString=FIX.4.4
SenderCompID=STALL_ACCEPTOR
TargetCompID=STALL_INITIATOR
HeartBtInt=30
`, port)))
	require.NoError(t, err)
	acceptor, err := NewAcceptor(logonRecorder{loggedOn: make(chan SessionID, 1)}, NewMemoryStoreFactory(), acceptorSettings, NewNullLogFactory())
	require.NoError(t, err)
	require.NoError(t, acceptor.Start())
	t.Cleanup(acceptor.Stop)

	initiatorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
ReconnectInterval=1
LogoutTimeout=1
[SESSION]
BeginString=FIX.4.4
SenderCompID=STALL_INITIATOR
TargetCompID=STALL_ACCEPTOR
SocketConnectHost=127.0.0.1
SocketConnectPort=%d
HeartBtInt=30
`, port)))
	require.NoError(t, err)
	app := logonRecorder{loggedOn: make(chan SessionID, 1)}
	initiator, err := NewInitiator(app, NewMemoryStoreFactory(), initiatorSettings, NewNullLogFactory())
	require.NoError(t, err)

	// Start, with a dialer whose connection the test can stall.
	dialer := stallingDialer{conns: make(chan *stallingConn, 1)}
	initiator.stopChan = make(chan interface{})
	for sessionID := range initiator.sessionSettings {
		initiator.wg.Add(1)
		go func(sessionID SessionID) {
			defer initiator.wg.Done()
			initiator.handleConnection(initiator.sessions[sessionID], nil, dialer)
		}(sessionID)
	}

	conn := <-dialer.conns
	select {
	case <-app.loggedOn:
	case <-time.After(10 * time.Second):
		t.Fatal("the initiator did not log on")
	}
	conn.stalled.Store(true)

	stopped := make(chan struct{})
	go func() {
		initiator.Stop()
		close(stopped)
	}()
	select {
	case <-stopped:
	case <-time.After(10 * time.Second):
		conn.Close() // let the stuck Stop end with the test
		t.Fatal("Stop did not return")
	}
	require.True(t, conn.isClosed(), "Stop returned with the connection open")
}

// stallingProxy forwards between an initiator and an acceptor, terminating
// TLS on the initiator's side when it has a certificate. Stalled, it stops
// reading what the initiator sends, so the initiator's writes block once the
// socket buffers fill.
type stallingProxy struct {
	listener  net.Listener
	initiator chan net.Conn
	stalled   atomic.Bool
}

func startStallingProxy(t *testing.T, acceptorPort int, cert *tls.Certificate) *stallingProxy {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	p := &stallingProxy{listener: listener, initiator: make(chan net.Conn, 1)}
	t.Cleanup(func() { _ = listener.Close() })
	go func() {
		raw, err := listener.Accept()
		if err != nil {
			return
		}
		initiator := raw
		if cert != nil {
			initiator = tls.Server(raw, &tls.Config{Certificates: []tls.Certificate{*cert}})
		}
		acceptor, err := net.Dial("tcp", fmt.Sprintf("127.0.0.1:%d", acceptorPort))
		if err != nil {
			_ = initiator.Close()
			return
		}
		t.Cleanup(func() {
			_ = initiator.Close()
			_ = acceptor.Close()
		})
		p.initiator <- initiator
		go func() { _, _ = io.Copy(initiator, acceptor) }()
		buf := make([]byte, 32*1024)
		for !p.stalled.Load() {
			_ = initiator.SetReadDeadline(time.Now().Add(10 * time.Millisecond))
			n, err := initiator.Read(buf)
			if n > 0 {
				if _, werr := acceptor.Write(buf[:n]); werr != nil {
					return
				}
			}
			if err != nil && !errors.Is(err, os.ErrDeadlineExceeded) {
				return
			}
		}
	}()
	return p
}

func selfSignedCertificate(t *testing.T) (*tls.Certificate, *x509.CertPool) {
	t.Helper()
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	require.NoError(t, err)
	template := &x509.Certificate{
		SerialNumber: big.NewInt(1),
		Subject:      pkix.Name{CommonName: "127.0.0.1"},
		IPAddresses:  []net.IP{net.ParseIP("127.0.0.1")},
		NotBefore:    time.Now().Add(-time.Hour),
		NotAfter:     time.Now().Add(time.Hour),
		KeyUsage:     x509.KeyUsageDigitalSignature,
		ExtKeyUsage:  []x509.ExtKeyUsage{x509.ExtKeyUsageServerAuth},
	}
	der, err := x509.CreateCertificate(rand.Reader, template, template, &key.PublicKey, key)
	require.NoError(t, err)
	parsed, err := x509.ParseCertificate(der)
	require.NoError(t, err)
	roots := x509.NewCertPool()
	roots.AddCert(parsed)
	return &tls.Certificate{Certificate: [][]byte{der}, PrivateKey: key}, roots
}

type plainDialer struct{}

func (plainDialer) DialContext(ctx context.Context, network, address string) (net.Conn, error) {
	var dialer net.Dialer
	return dialer.DialContext(ctx, network, address)
}

// A counterparty that stops reading while the session still has messages
// queued: the writer is already blocked when Stop is called, so the session
// loop is blocked behind it and cannot take the stop. Stop must still return
// within its timeouts, with the connection closed, over plain TCP and TLS.
func TestInitiatorStopEndsAWriteAlreadyBlocked(t *testing.T) {
	for _, useTLS := range []bool{false, true} {
		name := "tcp"
		if useTLS {
			name = "tls"
		}
		t.Run(name, func(t *testing.T) {
			listener, err := net.Listen("tcp", "127.0.0.1:0")
			require.NoError(t, err)
			acceptorPort := listener.Addr().(*net.TCPAddr).Port
			require.NoError(t, listener.Close())
			acceptorID, initiatorID := "FULL_ACCEPTOR_"+name, "FULL_INITIATOR_"+name

			acceptorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
SocketAcceptPort=%d
[SESSION]
BeginString=FIX.4.4
SenderCompID=%s
TargetCompID=%s
HeartBtInt=30
`, acceptorPort, acceptorID, initiatorID)))
			require.NoError(t, err)
			acceptor, err := NewAcceptor(logonRecorder{loggedOn: make(chan SessionID, 1)}, NewMemoryStoreFactory(), acceptorSettings, NewNullLogFactory())
			require.NoError(t, err)
			require.NoError(t, acceptor.Start())
			t.Cleanup(acceptor.Stop)

			var cert *tls.Certificate
			var tlsConfig *tls.Config
			if useTLS {
				var roots *x509.CertPool
				cert, roots = selfSignedCertificate(t)
				tlsConfig = &tls.Config{RootCAs: roots, MinVersion: tls.VersionTLS12}
			}
			proxy := startStallingProxy(t, acceptorPort, cert)

			initiatorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
ReconnectInterval=1
LogoutTimeout=1
[SESSION]
BeginString=FIX.4.4
SenderCompID=%s
TargetCompID=%s
SocketConnectHost=127.0.0.1
SocketConnectPort=%d
HeartBtInt=30
`, initiatorID, acceptorID, proxy.listener.Addr().(*net.TCPAddr).Port)))
			require.NoError(t, err)
			app := logonRecorder{loggedOn: make(chan SessionID, 1)}
			initiator, err := NewInitiator(app, NewMemoryStoreFactory(), initiatorSettings, NewNullLogFactory())
			require.NoError(t, err)
			initiator.stopChan = make(chan interface{})
			var sessionID SessionID
			for id := range initiator.sessionSettings {
				sessionID = id
				initiator.wg.Add(1)
				go func() {
					defer initiator.wg.Done()
					initiator.handleConnection(initiator.sessions[id], tlsConfig, plainDialer{})
				}()
			}

			var proxied net.Conn
			select {
			case proxied = <-proxy.initiator:
			case <-time.After(10 * time.Second):
				t.Fatal("the initiator did not connect")
			}
			select {
			case <-app.loggedOn:
			case <-time.After(10 * time.Second):
				t.Fatal("the initiator did not log on")
			}

			proxy.stalled.Store(true)
			filler := strings.Repeat("x", 64*1024)
			for range 512 { // 32 MiB, past any loopback socket buffers
				msg := NewMessage()
				msg.Header.SetString(tagMsgType, "1") // TestRequest
				msg.Body.SetString(Tag(112), filler)
				require.NoError(t, SendToTarget(msg, sessionID))
			}
			time.Sleep(500 * time.Millisecond) // the writer blocks on a full socket

			stopped := make(chan struct{})
			go func() {
				initiator.Stop()
				close(stopped)
			}()
			select {
			case <-stopped:
			case <-time.After(6 * time.Second): // LogoutTimeout is 1 s
				_ = proxied.Close() // let the stuck Stop end with the test
				t.Fatal("Stop did not return")
			}

			// The proxy drains what was sent and then reaches the end of
			// the connection: the initiator closed it.
			done := make(chan error, 1)
			go func() {
				_ = proxied.SetReadDeadline(time.Now().Add(10 * time.Second))
				_, err := io.Copy(io.Discard, proxied)
				done <- err
			}()
			// A close without TLS close notification reads as an
			// unexpected EOF; only a deadline means the connection is open.
			if err := <-done; errors.Is(err, os.ErrDeadlineExceeded) {
				t.Fatalf("the initiator's connection was not closed: %v", err)
			}
		})
	}
}

// With a counterparty that reads and answers, Stop still logs out
// gracefully: the counterparty receives the Logout and answers it before the
// connection closes, over plain TCP and TLS.
func TestInitiatorStopLogsOutGracefully(t *testing.T) {
	for _, useTLS := range []bool{false, true} {
		name := "tcp"
		if useTLS {
			name = "tls"
		}
		t.Run(name, func(t *testing.T) {
			listener, err := net.Listen("tcp", "127.0.0.1:0")
			require.NoError(t, err)
			acceptorPort := listener.Addr().(*net.TCPAddr).Port
			require.NoError(t, listener.Close())
			acceptorID, initiatorID := "GRACE_ACCEPTOR_"+name, "GRACE_INITIATOR_"+name

			acceptorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
SocketAcceptPort=%d
[SESSION]
BeginString=FIX.4.4
SenderCompID=%s
TargetCompID=%s
HeartBtInt=30
`, acceptorPort, acceptorID, initiatorID)))
			require.NoError(t, err)
			counterparty := logonRecorder{loggedOn: make(chan SessionID, 1), loggedOut: make(chan SessionID, 1)}
			acceptor, err := NewAcceptor(counterparty, NewMemoryStoreFactory(), acceptorSettings, NewNullLogFactory())
			require.NoError(t, err)
			require.NoError(t, acceptor.Start())
			t.Cleanup(acceptor.Stop)

			var cert *tls.Certificate
			var tlsConfig *tls.Config
			if useTLS {
				var roots *x509.CertPool
				cert, roots = selfSignedCertificate(t)
				tlsConfig = &tls.Config{RootCAs: roots, MinVersion: tls.VersionTLS12}
			}
			proxy := startStallingProxy(t, acceptorPort, cert)

			initiatorSettings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
ReconnectInterval=1
LogoutTimeout=5
[SESSION]
BeginString=FIX.4.4
SenderCompID=%s
TargetCompID=%s
SocketConnectHost=127.0.0.1
SocketConnectPort=%d
HeartBtInt=30
`, initiatorID, acceptorID, proxy.listener.Addr().(*net.TCPAddr).Port)))
			require.NoError(t, err)
			app := logonRecorder{loggedOn: make(chan SessionID, 1), loggedOut: make(chan SessionID, 1)}
			initiator, err := NewInitiator(app, NewMemoryStoreFactory(), initiatorSettings, NewNullLogFactory())
			require.NoError(t, err)
			initiator.stopChan = make(chan interface{})
			for id := range initiator.sessionSettings {
				initiator.wg.Add(1)
				go func() {
					defer initiator.wg.Done()
					initiator.handleConnection(initiator.sessions[id], tlsConfig, plainDialer{})
				}()
			}
			select {
			case <-app.loggedOn:
			case <-time.After(10 * time.Second):
				t.Fatal("the initiator did not log on")
			}

			began := time.Now()
			stopped := make(chan struct{})
			go func() {
				initiator.Stop()
				close(stopped)
			}()
			select {
			case <-stopped:
			case <-time.After(10 * time.Second):
				t.Fatal("Stop did not return")
			}
			// Answered, not timed out: well inside the 5 s LogoutTimeout.
			require.Less(t, time.Since(began), 4*time.Second, "the Logout was not answered")
			select {
			case <-counterparty.loggedOut:
			case <-time.After(5 * time.Second):
				t.Fatal("the counterparty never received the Logout")
			}
		})
	}
}

// silentTLSPeer accepts TCP connections and never answers the TLS
// handshake. Each accepted connection is handed to the test.
func silentTLSPeer(t *testing.T) (int, chan net.Conn) {
	t.Helper()
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	accepted := make(chan net.Conn, 16)
	t.Cleanup(func() { _ = listener.Close() })
	go func() {
		for {
			conn, err := listener.Accept()
			if err != nil {
				return
			}
			t.Cleanup(func() { _ = conn.Close() })
			accepted <- conn
		}
	}()
	return listener.Addr().(*net.TCPAddr).Port, accepted
}

// closedByPeer reports whether the other side closed conn within timeout:
// it drains what was sent and waits for the end of the stream.
func closedByPeer(conn net.Conn, timeout time.Duration) bool {
	_ = conn.SetReadDeadline(time.Now().Add(timeout))
	_, err := io.Copy(io.Discard, conn)
	return !errors.Is(err, os.ErrDeadlineExceeded)
}

// startTLSInitiator runs one initiator session against port over TLS, the
// way Start would, and returns the initiator.
func startTLSInitiator(t *testing.T, compIDs string, port int, extra string, tlsConfig *tls.Config) *Initiator {
	t.Helper()
	settings, err := ParseSettings(strings.NewReader(fmt.Sprintf(`[DEFAULT]
ReconnectInterval=1
%s
[SESSION]
BeginString=FIX.4.4
SenderCompID=%s_INITIATOR
TargetCompID=%s_ACCEPTOR
SocketConnectHost=127.0.0.1
SocketConnectPort=%d
HeartBtInt=30
`, extra, compIDs, compIDs, port)))
	require.NoError(t, err)
	initiator, err := NewInitiator(logonRecorder{loggedOn: make(chan SessionID, 1), loggedOut: make(chan SessionID, 1)}, NewMemoryStoreFactory(), settings, NewNullLogFactory())
	require.NoError(t, err)
	initiator.stopChan = make(chan interface{})
	for id := range initiator.sessionSettings {
		initiator.wg.Add(1)
		go func() {
			defer initiator.wg.Done()
			initiator.handleConnection(initiator.sessions[id], tlsConfig, plainDialer{})
		}()
	}
	return initiator
}

// Stop does not wait on a TLS handshake the counterparty never answers, and
// closes that connection.
func TestInitiatorStopDuringAStalledTLSHandshake(t *testing.T) {
	port, accepted := silentTLSPeer(t)
	initiator := startTLSInitiator(t, "HANDSHAKE_STOP", port, "LogonTimeout=60", &tls.Config{InsecureSkipVerify: true}) //nolint:gosec // the peer never answers
	var conn net.Conn
	select {
	case conn = <-accepted:
	case <-time.After(5 * time.Second):
		t.Fatal("the initiator did not connect")
	}

	stopped := make(chan struct{})
	go func() {
		initiator.Stop()
		close(stopped)
	}()
	select {
	case <-stopped:
	case <-time.After(3 * time.Second):
		_ = conn.Close() // let the stuck Stop end with the test
		t.Fatal("Stop waited on the handshake")
	}
	require.True(t, closedByPeer(conn, 2*time.Second), "the handshake's connection was left open")
}

// A TLS handshake the counterparty never answers fails after LogonTimeout,
// closes its connection, and the initiator connects again.
func TestInitiatorRetriesAStalledTLSHandshake(t *testing.T) {
	port, accepted := silentTLSPeer(t)
	initiator := startTLSInitiator(t, "HANDSHAKE_RETRY", port, "LogonTimeout=1", &tls.Config{InsecureSkipVerify: true}) //nolint:gosec // the peer never answers
	t.Cleanup(initiator.Stop)
	var first net.Conn
	select {
	case first = <-accepted:
	case <-time.After(5 * time.Second):
		t.Fatal("the initiator did not connect")
	}
	select {
	case <-accepted:
	case <-time.After(6 * time.Second): // LogonTimeout 1 s, then ReconnectInterval 1 s
		t.Fatal("the initiator did not connect again after the stalled handshake")
	}
	require.True(t, closedByPeer(first, 2*time.Second), "the stalled handshake's connection was left open")
}

// A failed TLS handshake closes the connection it was made on.
func TestInitiatorClosesAConnectionWhoseHandshakeFailed(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })
	cert, _ := selfSignedCertificate(t)
	failed := make(chan net.Conn, 1)
	go func() {
		raw, err := listener.Accept()
		if err != nil {
			return
		}
		t.Cleanup(func() { _ = raw.Close() })
		server := tls.Server(raw, &tls.Config{Certificates: []tls.Certificate{*cert}})
		_ = server.Handshake() // fails: the initiator does not trust this certificate
		failed <- raw
	}()
	// An empty root pool trusts nothing.
	initiator := startTLSInitiator(t, "HANDSHAKE_FAIL", listener.Addr().(*net.TCPAddr).Port, "LogonTimeout=60", &tls.Config{RootCAs: x509.NewCertPool(), MinVersion: tls.VersionTLS12})
	t.Cleanup(initiator.Stop)
	var raw net.Conn
	select {
	case raw = <-failed:
	case <-time.After(5 * time.Second):
		t.Fatal("the handshake did not fail")
	}
	require.True(t, closedByPeer(raw, 2*time.Second), "the failed handshake's connection was left open")
}

// SocketTimeout, when set, bounds the TLS handshake instead of LogonTimeout:
// a counterparty whose handshake takes longer than LogonTimeout but less
// than SocketTimeout still completes it, on the first connection.
func TestInitiatorHandshakeUsesSocketTimeout(t *testing.T) {
	listener, err := net.Listen("tcp", "127.0.0.1:0")
	require.NoError(t, err)
	t.Cleanup(func() { _ = listener.Close() })
	cert, roots := selfSignedCertificate(t)
	accepts := make(chan error, 4)
	go func() {
		for {
			raw, err := listener.Accept()
			if err != nil {
				return
			}
			t.Cleanup(func() { _ = raw.Close() })
			go func() {
				time.Sleep(2 * time.Second) // a slow handshake
				accepts <- tls.Server(raw, &tls.Config{Certificates: []tls.Certificate{*cert}}).Handshake()
			}()
		}
	}()
	initiator := startTLSInitiator(t, "HANDSHAKE_SLOW", listener.Addr().(*net.TCPAddr).Port, "LogonTimeout=1\nSocketTimeout=30s", &tls.Config{RootCAs: roots, MinVersion: tls.VersionTLS12})
	t.Cleanup(initiator.Stop)
	select {
	case err := <-accepts:
		require.NoError(t, err, "the slow handshake did not complete")
	case <-time.After(6 * time.Second):
		t.Fatal("the handshake never finished")
	}
	select {
	case err := <-accepts:
		t.Fatalf("the initiator connected again after a handshake that succeeded: %v", err)
	case <-time.After(2 * time.Second):
	}
}
