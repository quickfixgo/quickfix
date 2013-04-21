package fix40

import (
	"github.com/cbusbey/quickfixgo"
)

type ResendRequest struct {
	quickfixgo.Message
}
