package fix42

import (
	"github.com/cbusbey/quickfixgo"
)

type QuoteRequest struct {
	quickfixgo.Message
}
