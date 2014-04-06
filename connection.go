package quickfix

import (
	"bufio"
	"github.com/quickfixgo/quickfix/log"
	"github.com/quickfixgo/quickfix/tag"
	"net"
)

//Picks up session from net.Conn Acceptor
func handleAcceptorConnection(netConn net.Conn, log log.Log) {
	defer func() {
		if err := recover(); err != nil {
			log.OnEventf("Connection Terminated: %v", err)
		}

		netConn.Close()
	}()

	reader := bufio.NewReader(netConn)
	parser := newParser(reader)

	msgBytes, err := parser.readMessage()
	if err != nil {
		return
	}

	msg, err := MessageFromParsedBytes(msgBytes)
	if err != nil {
		log.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
		return
	}

	var sessID SessionID
	beginString := new(StringValue)
	if err := msg.Header.GetField(tag.BeginString, beginString); err == nil {
		sessID.BeginString = beginString.Value
	}

	senderCompID := new(StringValue)
	if err := msg.Header.GetField(tag.SenderCompID, senderCompID); err == nil {
		sessID.TargetCompID = senderCompID.Value
	}

	targetCompID := new(StringValue)
	if err := msg.Header.GetField(tag.TargetCompID, targetCompID); err == nil {
		sessID.SenderCompID = targetCompID.Value
	}

	defaultApplVerID := new(StringValue)
	if err := msg.Body.GetField(tag.DefaultApplVerID, defaultApplVerID); err == nil {
		sessID.DefaultApplVerID = defaultApplVerID.Value
	}

	session := activate(sessID)

	if session == nil {
		log.OnEventf("Session not found for incoming message: %v", msg.String())
		return
	} else {
		defer func() {
			deactivate(sessID)
		}()
	}

	if msgOut, err := session.accept(); err != nil {
		log.OnEventf("Session cannot accept: %v", err)
		return
	} else {

		msgIn := make(chan []byte)
		go writeLoop(netConn, msgOut)
		go func() {
			msgIn <- msgBytes
			readLoop(parser, msgIn)
		}()

		session.run(msgIn)
	}
}

func writeLoop(connection net.Conn, messageOut chan Buffer) {
	defer func() {
		close(messageOut)
	}()

	for {
		if msg := <-messageOut; msg == nil {
			return
		} else {
			connection.Write(msg.Bytes())
			msg.Free()
		}
	}
}

func readLoop(parser *parser, msgIn chan []byte) {
	defer func() {
		close(msgIn)
	}()

	for {
		if msg, err := parser.readMessage(); err != nil {
			switch err.(type) {
			//ignore message parser errors
			case ParseError:
				continue
			default:
				return
			}
		} else {
			msgIn <- msg
		}
	}
}
