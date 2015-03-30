package quickfix

import (
	"bufio"
	"github.com/quickfixgo/quickfix/fix/field"
	"net"
)

//Picks up session from net.Conn Initiator
func handleInitiatorConnection(netConn net.Conn, log Log, sessID SessionID, quit chan bool) {
	defer func() {
		if err := recover(); err != nil {
			log.OnEventf("Connection Terminated: %v", err)
		}

		netConn.Close()
	}()

	session := activate(sessID)
	if session == nil {
		log.OnEventf("Session not found for SessionID: %v", sessID)
		return
	}
	defer func() {
		deactivate(sessID)
	}()

	var msgOut chan []byte
	var err error
	if msgOut, err = session.initiate(); err != nil {
		log.OnEventf("Session cannot initiate: %v", err)
		return
	}

	reader := bufio.NewReader(netConn)
	parser := newParser(reader)

	msgIn := make(chan fixIn)
	go writeLoop(netConn, msgOut)
	go func() {
		readLoop(parser, msgIn)
	}()

	session.run(msgIn, quit)
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

	msg, err := parseMessage(msgBytes)
	if err != nil {
		log.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
		return
	}

	beginString := &field.BeginStringField{}
	msg.Header.Get(beginString)

	senderCompID := &field.SenderCompIDField{}
	msg.Header.Get(senderCompID)

	targetCompID := &field.TargetCompIDField{}
	msg.Header.Get(targetCompID)

	sessID := SessionID{BeginString: beginString.Value, SenderCompID: targetCompID.Value, TargetCompID: senderCompID.Value}
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

	var msgOut chan []byte
	if msgOut, err = session.accept(); err != nil {
		log.OnEventf("Session cannot accept: %v", err)
		return
	}

	msgIn := make(chan fixIn)
	go writeLoop(netConn, msgOut)
	go func() {
		msgIn <- fixIn{msgBytes, parser.lastRead}
		readLoop(parser, msgIn)
	}()

	session.run(msgIn, quit)
}

func writeLoop(connection net.Conn, messageOut chan []byte) {
	defer func() {
		close(messageOut)
	}()

	var msg []byte
	for {
		if msg = <-messageOut; msg == nil {
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
			switch err.(type) {
			//ignore message parser errors
			case parseError:
				continue
			default:
				return
			}
		} else {
			msgIn <- fixIn{msg, parser.lastRead}
		}
	}
}
