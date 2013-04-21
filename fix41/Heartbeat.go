package fix41

import (
	"github.com/cbusbey/quickfixgo"
)

type Heartbeat struct {
	quickfixgo.Message
}
