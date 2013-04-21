package fix43

import (
	"github.com/cbusbey/quickfixgo"
)

type Heartbeat struct {
	quickfixgo.Message
}
