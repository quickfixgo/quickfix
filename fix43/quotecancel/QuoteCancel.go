//Package quotecancel msg type = Z.
package quotecancel

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
)

//NoQuoteEntries is a repeating group in QuoteCancel
type NoQuoteEntries struct {
	//Instrument Component
	Instrument instrument.Component
}

//Message is a QuoteCancel FIX Message
type Message struct {
	FIXMsgType string `fix:"Z"`
	Header     fix43.Header
	//QuoteReqID is a non-required field for QuoteCancel.
	QuoteReqID *string `fix:"131"`
	//QuoteID is a required field for QuoteCancel.
	QuoteID string `fix:"117"`
	//QuoteCancelType is a required field for QuoteCancel.
	QuoteCancelType int `fix:"298"`
	//QuoteResponseLevel is a non-required field for QuoteCancel.
	QuoteResponseLevel *int `fix:"301"`
	//Parties Component
	Parties parties.Component
	//Account is a non-required field for QuoteCancel.
	Account *string `fix:"1"`
	//AccountType is a non-required field for QuoteCancel.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for QuoteCancel.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteCancel.
	TradingSessionSubID *string `fix:"625"`
	//NoQuoteEntries is a non-required field for QuoteCancel.
	NoQuoteEntries []NoQuoteEntries `fix:"295,omitempty"`
	Trailer        fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)               { m.QuoteReqID = &v }
func (m *Message) SetQuoteID(v string)                  { m.QuoteID = v }
func (m *Message) SetQuoteCancelType(v int)             { m.QuoteCancelType = v }
func (m *Message) SetQuoteResponseLevel(v int)          { m.QuoteResponseLevel = &v }
func (m *Message) SetAccount(v string)                  { m.Account = &v }
func (m *Message) SetAccountType(v int)                 { m.AccountType = &v }
func (m *Message) SetTradingSessionID(v string)         { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)      { m.TradingSessionSubID = &v }
func (m *Message) SetNoQuoteEntries(v []NoQuoteEntries) { m.NoQuoteEntries = v }

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
	return enum.BeginStringFIX43, "Z", r
}
