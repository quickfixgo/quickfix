//Package rfqrequest msg type = AH.
package rfqrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
)

//NoRelatedSym is a repeating group in RFQRequest
type NoRelatedSym struct {
	//Instrument Component
	instrument.Instrument
	//PrevClosePx is a non-required field for NoRelatedSym.
	PrevClosePx *float64 `fix:"140"`
	//QuoteRequestType is a non-required field for NoRelatedSym.
	QuoteRequestType *int `fix:"303"`
	//QuoteType is a non-required field for NoRelatedSym.
	QuoteType *int `fix:"537"`
	//TradingSessionID is a non-required field for NoRelatedSym.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoRelatedSym.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoRelatedSym) SetPrevClosePx(v float64)        { m.PrevClosePx = &v }
func (m *NoRelatedSym) SetQuoteRequestType(v int)       { m.QuoteRequestType = &v }
func (m *NoRelatedSym) SetQuoteType(v int)              { m.QuoteType = &v }
func (m *NoRelatedSym) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoRelatedSym) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a RFQRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AH"`
	fix43.Header
	//RFQReqID is a required field for RFQRequest.
	RFQReqID string `fix:"644"`
	//NoRelatedSym is a required field for RFQRequest.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	//SubscriptionRequestType is a non-required field for RFQRequest.
	SubscriptionRequestType *string `fix:"263"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetRFQReqID(v string)                { m.RFQReqID = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym)    { m.NoRelatedSym = v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }

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
	return enum.BeginStringFIX43, "AH", r
}
