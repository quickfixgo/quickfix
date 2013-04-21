package fix50

import (
	"github.com/cbusbey/quickfixgo"
)

type CrossOrderCancelRequest struct {
	quickfixgo.Message
}
