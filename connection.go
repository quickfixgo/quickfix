package quickfix

import (
	"bufio"
	"io"
	"net"
	"time"
)

//Picks up session from net.Conn Initiator
func handleInitiatorConnection(address string, log Log, sessID SessionID, quit chan bool) {
	reconnectInterval := 30 * time.Second
	session := activate(sessID)
	if session == nil {
		log.OnEventf("Session not found for SessionID: %v", sessID)
		return
	}

	defer deactivate(sessID)

	for {
		msgIn := make(chan fixIn)
		msgOut := make(chan []byte)

		netConn, err := net.Dial("tcp", address)
		if err != nil {
			goto reconnect
		}

		go readLoop(newParser(bufio.NewReader(netConn)), msgIn)
		go func() {
			writeLoop(netConn, msgOut)
			netConn.Close()
		}()
		session.initiate(msgIn, msgOut, quit)

	reconnect:
		log.OnEventf("%v Reconnecting in %v", sessID, reconnectInterval)
		time.Sleep(reconnectInterval)
		continue
	}
}

//Picks up session from net.Conn Acceptor
func handleAcceptorConnection(netConn net.Conn, qualifiedSessionIDs map[SessionID]SessionID, log Log, quit chan bool, done chan int, cnxNumber int) {
	defer func() {
		if err := recover(); err != nil {
			log.OnEventf("Connection Terminated: %v", err)
		}

		netConn.Close()
		done <- cnxNumber
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

	go session.accept(msgIn, msgOut, quit)
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
	defer func() {
		close(msgIn)
	}()

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
