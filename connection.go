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
	"errors"
	"io"
	"net"
	"sync/atomic"
	"time"
)

// stopWriter is the initiator's connection as its writer sees it. Once the
// initiator stops, every write, including one already in progress, must
// finish within timeout. A write that does not leaves the stream with a
// partial message on it and the counterparty not reading, so it closes the
// connection: the writer then fails at once, and the session can stop.
type stopWriter struct {
	conn     net.Conn
	timeout  time.Duration
	stopping atomic.Bool
}

// stop bounds the write in progress and every later one by timeout.
func (w *stopWriter) stop() {
	w.stopping.Store(true)
	_ = w.conn.SetWriteDeadline(time.Now().Add(w.timeout))
}

func (w *stopWriter) Write(b []byte) (int, error) {
	if w.stopping.Load() {
		_ = w.conn.SetWriteDeadline(time.Now().Add(w.timeout))
	}
	n, err := w.conn.Write(b)
	if err != nil && w.stopping.Load() {
		_ = transport(w.conn).Close()
	}
	return n, err
}

// transport is the connection under TLS, if any. Closing it ends a blocked
// write or close notification at once; closing a tls.Conn already being
// closed only reports net.ErrClosed while that close notification waits.
func transport(conn net.Conn) net.Conn {
	if tlsConn, ok := conn.(*tls.Conn); ok {
		return tlsConn.NetConn()
	}
	return conn
}

// closeTransport closes the connection's transport. It is closed once this
// returns, whoever closed it first.
func closeTransport(conn net.Conn, log Log) {
	if err := transport(conn).Close(); err != nil && !errors.Is(err, net.ErrClosed) {
		log.OnEvent(err.Error())
	}
}

func writeLoop(connection io.Writer, messageOut chan []byte, log Log) {
	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}

		if _, err := connection.Write(msg); err != nil {
			log.OnEvent(err.Error())
		}
	}
}

func readLoop(parser *parser, msgIn chan fixIn, log Log) {
	defer close(msgIn)

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			log.OnEvent(err.Error())
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
