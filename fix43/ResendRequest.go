package fix43

import (
	"github.com/cbusbey/quickfixgo"
)

type ResendRequest struct {
	quickfixgo.Message
}
