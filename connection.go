package quickfix

import "io"

func writeLoop(connection io.Writer, messageOut chan []byte, log Log) {
	//等待有发送的消息
	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}

		//socket写入，此时只关注返回错误
		if _, err := connection.Write(msg); err != nil {
			log.OnEvent(err.Error())
		}
	}
}

func readLoop(parser *parser, msgIn chan fixIn) {
	defer close(msgIn)

	//fixIn保存了接收到的消息
	for {
		msg, err := parser.ReadMessage()
		if err != nil {
			return
		}
		//读取消息之后，放到到MsgIn通道中
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
