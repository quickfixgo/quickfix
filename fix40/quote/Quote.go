//Package quote msg type = S.
package quote

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a Quote FIX Message
type Message struct {
	FIXMsgType string `fix:"S"`
	Header     fix40.Header
	//QuoteReqID is a non-required field for Quote.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for Quote.
	QuoteID string `fix:"117"`
	//Symbol is a required field for Quote.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Quote.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Quote.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Quote.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for Quote.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for Quote.
	SecurityDesc *string `fix:"107"`
	//BidPx is a required field for Quote.
	BidPx float64 `fix:"132"`
	//OfferPx is a non-required field for Quote.
	OfferPx *float64 `fix:"133"`
	//BidSize is a non-required field for Quote.
	BidSize *int `fix:"134"`
	//OfferSize is a non-required field for Quote.
	OfferSize *int `fix:"135"`
	//ValidUntilTime is a non-required field for Quote.
	ValidUntilTime *time.Time `fix:"62"`
	Trailer        fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX40, "S", r
}
