package fix44

import (
	"github.com/cbusbey/quickfixgo"
)

type ResendRequest struct {
	quickfixgo.Message
}
