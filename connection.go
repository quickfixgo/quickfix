package quickfix

import (
	"context"
	"io"

	"golang.org/x/time/rate"
)

const maxBurst = 1

func writeLoop(connection io.Writer, messageOut chan []byte, log Log, maxMessagesPerSecond int) {
	limiter := rate.NewLimiter(rate.Limit(maxMessagesPerSecond), maxBurst)
	ctx := context.Background()

	for {
		msg, ok := <-messageOut
		if !ok {
			return
		}

		if maxMessagesPerSecond > 0 {
			err := limiter.Wait(ctx)
			// if we get an error from the limiter we log, but continue on to
			// attempt to write the message to the remote party
			if err != nil {
				log.OnEvent(err.Error())
			}
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
			return
		}
		msgIn <- fixIn{msg, parser.lastRead}
	}
}
