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
			log.Println(`parser read message failed,conection readLoop just quit, error_info->`, err.Error())
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
