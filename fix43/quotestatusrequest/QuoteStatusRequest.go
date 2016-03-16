//Package quotestatusrequest msg type = a.
package quotestatusrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/parties"
)

//Message is a QuoteStatusRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"a"`
	fix43.Header
	//QuoteStatusReqID is a non-required field for QuoteStatusRequest.
	QuoteStatusReqID *string `fix:"649"`
	//QuoteID is a non-required field for QuoteStatusRequest.
	QuoteID *string `fix:"117"`
	//Instrument is a required component for QuoteStatusRequest.
	instrument.Instrument
	//Parties is a non-required component for QuoteStatusRequest.
	Parties *parties.Parties
	//Account is a non-required field for QuoteStatusRequest.
	Account *string `fix:"1"`
	//AccountType is a non-required field for QuoteStatusRequest.
	AccountType *int `fix:"581"`
	//TradingSessionID is a non-required field for QuoteStatusRequest.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for QuoteStatusRequest.
	TradingSessionSubID *string `fix:"625"`
	//SubscriptionRequestType is a non-required field for QuoteStatusRequest.
	SubscriptionRequestType *string `fix:"263"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteStatusReqID(v string)          { m.QuoteStatusReqID = &v }
func (m *Message) SetQuoteID(v string)                   { m.QuoteID = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = v }
func (m *Message) SetParties(v parties.Parties)          { m.Parties = &v }
func (m *Message) SetAccount(v string)                   { m.Account = &v }
func (m *Message) SetAccountType(v int)                  { m.AccountType = &v }
func (m *Message) SetTradingSessionID(v string)          { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)       { m.TradingSessionSubID = &v }
func (m *Message) SetSubscriptionRequestType(v string)   { m.SubscriptionRequestType = &v }

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
	return enum.BeginStringFIX43, "a", r
}
