package fix41

import (
	"github.com/cbusbey/quickfixgo"
)

type ResendRequest struct {
	quickfixgo.Message
}
