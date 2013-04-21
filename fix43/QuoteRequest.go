package fix43

import (
	"github.com/cbusbey/quickfixgo"
)

type QuoteRequest struct {
	quickfixgo.Message
}
