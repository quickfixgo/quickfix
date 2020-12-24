package quickfix

import (
	"io"
	"log"
)

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

func readLoop(parser *parser, msgIn chan fixIn) {
	defer close(msgIn)

	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			log.Printf(`readLoop.parser.ReadMessage error,error_info[%v]  \r\n`, err.Error())
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
