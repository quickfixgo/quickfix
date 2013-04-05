package session

import (
	"bufio"
	"net"
	"github.com/cbusbey/quickfixgo/fix"
	"github.com/cbusbey/quickfixgo/log"
	"github.com/cbusbey/quickfixgo/message"
	"github.com/cbusbey/quickfixgo/message/basic"
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
	if beginString, err := msg.Header().StringField(fix.BeginString); err == nil {
		sessID.BeginString = beginString.Value()
	}

	if senderCompID, err := msg.Header().StringField(fix.SenderCompID); err == nil {
		sessID.TargetCompID = senderCompID.Value()
	}

	if targetCompID, err := msg.Header().StringField(fix.TargetCompID); err == nil {
		sessID.SenderCompID = targetCompID.Value()
	}

	if defaultApplVerID, err := msg.Body().StringField(fix.DefaultApplVerID); err == nil {
		sessID.DefaultApplVerID = defaultApplVerID.Value()
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
