//Package quoterequest msg type = R.
package quoterequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
)

//Message is a QuoteRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"R"`
	Header     fix40.Header
	//QuoteReqID is a required field for QuoteRequest.
	QuoteReqID string `fix:"131"`
	//Symbol is a required field for QuoteRequest.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for QuoteRequest.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for QuoteRequest.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for QuoteRequest.
	IDSource *string `fix:"22"`
	//Issuer is a non-required field for QuoteRequest.
	Issuer *string `fix:"106"`
	//SecurityDesc is a non-required field for QuoteRequest.
	SecurityDesc *string `fix:"107"`
	//PrevClosePx is a non-required field for QuoteRequest.
	PrevClosePx *float64 `fix:"140"`
	//Side is a non-required field for QuoteRequest.
	Side *string `fix:"54"`
	//OrderQty is a non-required field for QuoteRequest.
	OrderQty *int `fix:"38"`
	Trailer  fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX40, "R", r
}
