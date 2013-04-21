package fix42

import (
	"github.com/cbusbey/quickfixgo"
)

type BidRequest struct {
	quickfixgo.Message
}
