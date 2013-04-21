package session

import (
	"bufio"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/message/basic"
	"github.com/cbusbey/quickfixgo/tag"
	"net"
)

//Picks up session from net.Conn Acceptor
func HandleAcceptorConnection(netConn net.Conn, log log.Log) {
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

	msg, err := basic.MessageFromParsedBytes(msgBytes)
	if err != nil {
		log.OnEvent("Invalid message: " + string(msgBytes) + err.Error())
		return
	}

	var sessID ID
	if beginString, ok := msg.Header().StringValue(tag.BeginString); ok {
		sessID.BeginString = beginString
	}

	if senderCompID, ok := msg.Header().StringValue(tag.SenderCompID); ok {
		sessID.TargetCompID = senderCompID
	}

	if targetCompID, ok := msg.Header().StringValue(tag.TargetCompID); ok {
		sessID.SenderCompID = targetCompID
	}

	if defaultApplVerID, ok := msg.Body().StringValue(tag.DefaultApplVerID); ok {
		sessID.DefaultApplVerID = defaultApplVerID
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

func writeLoop(connection net.Conn, messageOut chan message.Buffer) {
	for {
		msg, ok := <-messageOut
		if ok {
			connection.Write(msg.Bytes())
			msg.Free()
		} else {
			return
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
			case message.ParseError:
				continue
			default:
				return
			}
		} else {
			msgIn <- msg
		}
	}
}
