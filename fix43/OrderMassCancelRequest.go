package fix43

import (
	"github.com/cbusbey/quickfixgo"
)

type OrderMassCancelRequest struct {
	quickfixgo.Message
}
