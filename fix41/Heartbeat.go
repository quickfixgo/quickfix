package fix41

import (
	"github.com/cbusbey/quickfixgo/message"
)

type Heartbeat struct {
	message.Message
}
