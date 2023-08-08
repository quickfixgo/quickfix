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
	"bytes"
	"crypto/tls"
	"io"
	"net"
	"runtime/debug"
	"strconv"
	"sync"

	proxyproto "github.com/armon/go-proxyproto"

	"github.com/quickfixgo/quickfix/config"
)

// Acceptor accepts connections from FIX clients and manages the associated sessions.
type Acceptor struct {
	app                   Application
	settings              *Settings
	logFactory            LogFactory
	storeFactory          MessageStoreFactory
	globalLog             Log
	sessions              map[SessionID]*session
	sessionGroup          sync.WaitGroup
	listenerShutdown      sync.WaitGroup
	dynamicSessions       bool
	dynamicQualifier      bool
	dynamicQualifierCount int
	dynamicSessionChan    chan *session
	sessionAddr           sync.Map
	sessionHostPort       map[SessionID]int
	listeners             map[string]net.Listener
	connectionValidator   ConnectionValidator
	sessionFactory
}

// ConnectionValidator is an interface allowing to implement a custom authentication logic.
type ConnectionValidator interface {
	// Validate the connection for validity. This can be a part of authentication process.
	// For example, you may tie up a SenderCompID to an IP range, or to a specific TLS certificate as a part of mTLS.
	Validate(netConn net.Conn, session SessionID) error
}

// Start accepting connections.
func (a *Acceptor) Start() (err error) {
	socketAcceptHost := ""
	if a.settings.GlobalSettings().HasSetting(config.SocketAcceptHost) {
		if socketAcceptHost, err = a.settings.GlobalSettings().Setting(config.SocketAcceptHost); err != nil {
			return
		}
	}

	a.sessionHostPort = make(map[SessionID]int)
	a.listeners = make(map[string]net.Listener)
	for sessionID, sessionSettings := range a.settings.SessionSettings() {
		if sessionSettings.HasSetting(config.SocketAcceptPort) {
			if a.sessionHostPort[sessionID], err = sessionSettings.IntSetting(config.SocketAcceptPort); err != nil {
				return
			}
		} else if a.sessionHostPort[sessionID], err = a.settings.GlobalSettings().IntSetting(config.SocketAcceptPort); err != nil {
			return
		}
		address := net.JoinHostPort(socketAcceptHost, strconv.Itoa(a.sessionHostPort[sessionID]))
		a.listeners[address] = nil
	}

	var tlsConfig *tls.Config
	if tlsConfig, err = loadTLSConfig(a.settings.GlobalSettings()); err != nil {
		return
	}

	var useTCPProxy bool
	if a.settings.GlobalSettings().HasSetting(config.UseTCPProxy) {
		if useTCPProxy, err = a.settings.GlobalSettings().BoolSetting(config.UseTCPProxy); err != nil {
			return
		}
	}

	for address := range a.listeners {
		if tlsConfig != nil {
			if a.listeners[address], err = tls.Listen("tcp", address, tlsConfig); err != nil {
				return
			}
		} else if a.listeners[address], err = net.Listen("tcp", address); err != nil {
			return
		} else if useTCPProxy {
			a.listeners[address] = &proxyproto.Listener{Listener: a.listeners[address]}
		}
	}

	for _, s := range a.sessions {
		a.sessionGroup.Add(1)
		go func(s *session) {
			s.run()
			a.sessionGroup.Done()
		}(s)
	}
	if a.dynamicSessions {
		a.dynamicSessionChan = make(chan *session)
		a.sessionGroup.Add(1)
		go func() {
			a.dynamicSessionsLoop()
			a.sessionGroup.Done()
		}()
	}
	a.listenerShutdown.Add(len(a.listeners))
	for _, listener := range a.listeners {
		go a.listenForConnections(listener)
	}
	return
}

// Stop logs out existing sessions, close their connections, and stop accepting new connections.
func (a *Acceptor) Stop() {
	defer func() {
		_ = recover() // suppress sending on closed channel error
	}()

	for _, listener := range a.listeners {
		listener.Close()
	}
	a.listenerShutdown.Wait()
	if a.dynamicSessions {
		close(a.dynamicSessionChan)
	}
	for _, session := range a.sessions {
		session.stop()
	}
	a.sessionGroup.Wait()
}

// RemoteAddr gets remote IP address for a given session.
func (a *Acceptor) RemoteAddr(sessionID SessionID) (net.Addr, bool) {
	addr, ok := a.sessionAddr.Load(sessionID)
	if !ok || addr == nil {
		return nil, false
	}
	val, ok := addr.(net.Addr)
	return val, ok
}

// NewAcceptor creates and initializes a new Acceptor.
func NewAcceptor(app Application, storeFactory MessageStoreFactory, settings *Settings, logFactory LogFactory) (a *Acceptor, err error) {
	a = &Acceptor{
		app:             app,
		storeFactory:    storeFactory,
		settings:        settings,
		logFactory:      logFactory,
		sessions:        make(map[SessionID]*session),
		sessionHostPort: make(map[SessionID]int),
		listeners:       make(map[string]net.Listener),
	}
	if a.settings.GlobalSettings().HasSetting(config.DynamicSessions) {
		if a.dynamicSessions, err = settings.globalSettings.BoolSetting(config.DynamicSessions); err != nil {
			return
		}

		if a.settings.GlobalSettings().HasSetting(config.DynamicQualifier) {
			if a.dynamicQualifier, err = settings.globalSettings.BoolSetting(config.DynamicQualifier); err != nil {
				return
			}
		}
	}

	if a.globalLog, err = logFactory.Create(); err != nil {
		return
	}

	for sessionID, sessionSettings := range settings.SessionSettings() {
		sessID := sessionID
		sessID.Qualifier = ""

		if _, dup := a.sessions[sessID]; dup {
			return a, errDuplicateSessionID
		}

		if a.sessions[sessID], err = a.createSession(sessionID, storeFactory, sessionSettings, logFactory, app); err != nil {
			return
		}
	}

	return
}

func (a *Acceptor) listenForConnections(listener net.Listener) {
	defer a.listenerShutdown.Done()

	for {
		netConn, err := listener.Accept()
		if err != nil {
			return
		}

		go func() {
			a.handleConnection(netConn)
		}()
	}
}

func (a *Acceptor) invalidMessage(msg *bytes.Buffer, err error) {
	a.globalLog.OnEventf("Invalid Message: %s, %v", msg.Bytes(), err.Error())
}

func (a *Acceptor) handleConnection(netConn net.Conn) {
	defer func() {
		if err := recover(); err != nil {
			a.globalLog.OnEventf("Connection Terminated with Panic: %s", debug.Stack())
		}

		if err := netConn.Close(); err != nil {
			a.globalLog.OnEvent(err.Error())
		}
	}()

	reader := bufio.NewReader(netConn)
	parser := newParser(reader)

	msgBytes, err := parser.ReadMessage()
	if err != nil {
		if err == io.EOF {
			a.globalLog.OnEvent("Connection Terminated")
		} else {
			a.globalLog.OnEvent(err.Error())
		}
		return
	}

	msg := NewMessage()
	err = ParseMessage(msg, msgBytes)
	if err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var beginString FIXString
	if err := msg.Header.GetField(tagBeginString, &beginString); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var senderCompID FIXString
	if err := msg.Header.GetField(tagSenderCompID, &senderCompID); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var senderSubID FIXString
	if msg.Header.Has(tagSenderSubID) {
		if err := msg.Header.GetField(tagSenderSubID, &senderSubID); err != nil {
			a.invalidMessage(msgBytes, err)
			return
		}
	}

	var senderLocationID FIXString
	if msg.Header.Has(tagSenderLocationID) {
		if err := msg.Header.GetField(tagSenderLocationID, &senderLocationID); err != nil {
			a.invalidMessage(msgBytes, err)
			return
		}
	}

	var targetCompID FIXString
	if err := msg.Header.GetField(tagTargetCompID, &targetCompID); err != nil {
		a.invalidMessage(msgBytes, err)
		return
	}

	var targetSubID FIXString
	if msg.Header.Has(tagTargetSubID) {
		if err := msg.Header.GetField(tagTargetSubID, &targetSubID); err != nil {
			a.invalidMessage(msgBytes, err)
			return
		}
	}

	var targetLocationID FIXString
	if msg.Header.Has(tagTargetLocationID) {
		if err := msg.Header.GetField(tagTargetLocationID, &targetLocationID); err != nil {
			a.invalidMessage(msgBytes, err)
			return
		}
	}

	sessID := SessionID{BeginString: string(beginString),
		SenderCompID: string(targetCompID), SenderSubID: string(targetSubID), SenderLocationID: string(targetLocationID),
		TargetCompID: string(senderCompID), TargetSubID: string(senderSubID), TargetLocationID: string(senderLocationID),
	}

	localConnectionPort := netConn.LocalAddr().(*net.TCPAddr).Port
	if expectedPort, ok := a.sessionHostPort[sessID]; ok && expectedPort != localConnectionPort {
		a.globalLog.OnEventf("Session %v not found for incoming message: %s", sessID, msgBytes)
		return
	}

	// We have a session ID and a network connection. This seems to be a good place for any custom authentication logic.
	if a.connectionValidator != nil {
		if err := a.connectionValidator.Validate(netConn, sessID); err != nil {
			a.globalLog.OnEventf("Unable to validate a connection for session %v: %v", sessID, err.Error())
			return
		}
	}

	if a.dynamicQualifier {
		a.dynamicQualifierCount++
		sessID.Qualifier = strconv.Itoa(a.dynamicQualifierCount)
	}
	session, ok := a.sessions[sessID]
	if !ok {
		if !a.dynamicSessions {
			a.globalLog.OnEventf("Session %v not found for incoming message: %s", sessID, msgBytes)
			return
		}
		dynamicSession, err := a.sessionFactory.createSession(sessID, a.storeFactory, a.settings.globalSettings.clone(), a.logFactory, a.app)
		if err != nil {
			a.globalLog.OnEventf("Dynamic session %v failed to create: %v", sessID, err)
			return
		}
		a.dynamicSessionChan <- dynamicSession
		session = dynamicSession
		defer session.stop()
	}

	a.sessionAddr.Store(sessID, netConn.RemoteAddr())
	msgIn := make(chan fixIn)
	msgOut := make(chan []byte)

	if err := session.connect(msgIn, msgOut); err != nil {
		a.globalLog.OnEventf("Unable to accept session %v connection: %v", sessID, err.Error())
		return
	}

	go func() {
		msgIn <- fixIn{msgBytes, parser.lastRead}
		readLoop(parser, msgIn)
	}()

	writeLoop(netConn, msgOut, a.globalLog)
}

func (a *Acceptor) dynamicSessionsLoop() {
	var id int
	var sessions = map[int]*session{}
	var complete = make(chan int)
	defer close(complete)
LOOP:
	for {
		select {
		case session, ok := <-a.dynamicSessionChan:
			if !ok {
				for _, oldSession := range sessions {
					oldSession.stop()
				}
				break LOOP
			}
			id++
			sessionID := id
			sessions[sessionID] = session
			go func() {
				session.run()
				err := UnregisterSession(session.sessionID)
				if err != nil {
					a.globalLog.OnEventf("Unregister dynamic session %v failed: %v", session.sessionID, err)
					return
				}
				complete <- sessionID
			}()
		case id := <-complete:
			session, ok := sessions[id]
			if ok {
				a.sessionAddr.Delete(session.sessionID)
				delete(sessions, id)
			} else {
				a.globalLog.OnEventf("Missing dynamic session %v!", id)
			}
		}
	}

	if len(sessions) == 0 {
		return
	}

	for id := range complete {
		delete(sessions, id)
		if len(sessions) == 0 {
			return
		}
	}
}

// SetConnectionValidator sets an optional connection validator.
// Use it when you need a custom authentication logic that includes lower level interactions,
// like mTLS auth or IP whitelistening.
// To remove a previously set validator call it with a nil value:
//
//	a.SetConnectionValidator(nil)
func (a *Acceptor) SetConnectionValidator(validator ConnectionValidator) {
	a.connectionValidator = validator
}
