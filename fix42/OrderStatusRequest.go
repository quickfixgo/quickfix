package fix42

import (
	"github.com/cbusbey/quickfixgo"
)

type OrderStatusRequest struct {
	quickfixgo.Message
}
