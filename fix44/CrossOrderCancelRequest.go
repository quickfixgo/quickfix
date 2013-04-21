package fix44

import (
	"github.com/cbusbey/quickfixgo"
)

type CrossOrderCancelRequest struct {
	quickfixgo.Message
}
