package quickfix

import (
	"bufio"
	"crypto/tls"
	"io"
	"net"
	"time"
)

//Picks up session from net.Conn Initiator
func handleInitiatorConnection(address string, log Log, sessID SessionID, reconnectInterval time.Duration, tlsConfig *tls.Config) {
	session := activate(sessID)
	if session == nil {
		log.OnEventf("Session not found for SessionID: %v", sessID)
		return
	}

	defer deactivate(sessID)

	for {
		msgIn := make(chan fixIn)
		msgOut := make(chan []byte)

		var netConn net.Conn
		if tlsConfig != nil {
			tlsConn, err := tls.Dial("tcp", address, tlsConfig)
			if err != nil {
				log.OnEventf("Failed to connect: %v", err)
				goto reconnect
			}

			err = tlsConn.Handshake()
			if err != nil {
				log.OnEventf("Failed handshake:%v", err)
				goto reconnect
			}
			netConn = tlsConn
		} else {
			var err error
			netConn, err = net.Dial("tcp", address)
			if err != nil {
				log.OnEventf("Failed to connect: %v", err)
				goto reconnect
			}
		}

		go readLoop(newParser(bufio.NewReader(netConn)), msgIn)
		session.initiate(msgIn, msgOut)
		writeLoop(netConn, msgOut)
		netConn.Close()

	reconnect:
		log.OnEventf("%v Reconnecting in %v", sessID, reconnectInterval)
		time.Sleep(reconnectInterval)
		continue
	}
}

//Picks up session from net.Conn Acceptor
func handleAcceptorConnection(netConn net.Conn, qualifiedSessionIDs map[SessionID]SessionID, log Log) {
	defer func() {
		if err := recover(); err != nil {
			log.OnEventf("Connection Terminated: %v", err)
		}

		netConn.Close()
	}()

	reader := bufio.NewReader(netConn)
	parser := newParser(reader)

	msgBytes, err := parser.ReadMessage()

	if err != nil {
		return
	}

	msg, err := ParseMessage(msgBytes)
	if err != nil {
		log.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
		return
	}

	var beginString FIXString
	msg.Header.GetField(tagBeginString, &beginString)

	var senderCompID FIXString
	msg.Header.GetField(tagSenderCompID, &senderCompID)

	var targetCompID FIXString
	msg.Header.GetField(tagTargetCompID, &targetCompID)

	sessID := SessionID{BeginString: string(beginString), SenderCompID: string(targetCompID), TargetCompID: string(senderCompID)}
	qualifiedSessID, validID := qualifiedSessionIDs[sessID]

	if !validID {
		log.OnEventf("Session %v not found for incoming message: %v", sessID, msg.String())
		return
	}

	session := activate(qualifiedSessID)

	if session == nil {
		log.OnEventf("Cannot activate session for incoming message: %v", msg.String())
		return
	}
	defer func() {
		deactivate(qualifiedSessID)
	}()

	msgIn := make(chan fixIn)
	msgOut := make(chan []byte)

	go func() {
		msgIn <- fixIn{msgBytes, parser.lastRead}
		readLoop(parser, msgIn)
	}()

	session.accept(msgIn, msgOut)
	writeLoop(netConn, msgOut)
}

func writeLoop(connection io.Writer, messageOut chan []byte) {
	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}

		connection.Write(msg)
	}
}

func readLoop(parser *parser, msgIn chan fixIn) {
	defer close(msgIn)

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
