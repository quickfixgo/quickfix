package fix41

import (
	"github.com/cbusbey/quickfixgo"
)

type OrderStatusRequest struct {
	quickfixgo.Message
}
