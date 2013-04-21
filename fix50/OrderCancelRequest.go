package fix50

import (
	"github.com/cbusbey/quickfixgo"
)

type OrderCancelRequest struct {
	quickfixgo.Message
}
