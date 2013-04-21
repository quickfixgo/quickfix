package fix43

import (
	"github.com/cbusbey/quickfixgo"
)

type CrossOrderCancelRequest struct {
	quickfixgo.Message
}
