//Package ordercancelreject msg type = 9.
package ordercancelreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"time"
)

//Message is a OrderCancelReject FIX Message
type Message struct {
	FIXMsgType string `fix:"9"`
	fix43.Header
	//OrderID is a required field for OrderCancelReject.
	OrderID string `fix:"37"`
	//SecondaryOrderID is a non-required field for OrderCancelReject.
	SecondaryOrderID *string `fix:"198"`
	//SecondaryClOrdID is a non-required field for OrderCancelReject.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdID is a required field for OrderCancelReject.
	ClOrdID string `fix:"11"`
	//ClOrdLinkID is a non-required field for OrderCancelReject.
	ClOrdLinkID *string `fix:"583"`
	//OrigClOrdID is a required field for OrderCancelReject.
	OrigClOrdID string `fix:"41"`
	//OrdStatus is a required field for OrderCancelReject.
	OrdStatus string `fix:"39"`
	//WorkingIndicator is a non-required field for OrderCancelReject.
	WorkingIndicator *bool `fix:"636"`
	//OrigOrdModTime is a non-required field for OrderCancelReject.
	OrigOrdModTime *time.Time `fix:"586"`
	//ListID is a non-required field for OrderCancelReject.
	ListID *string `fix:"66"`
	//Account is a non-required field for OrderCancelReject.
	Account *string `fix:"1"`
	//AccountType is a non-required field for OrderCancelReject.
	AccountType *int `fix:"581"`
	//TradeOriginationDate is a non-required field for OrderCancelReject.
	TradeOriginationDate *string `fix:"229"`
	//TransactTime is a non-required field for OrderCancelReject.
	TransactTime *time.Time `fix:"60"`
	//CxlRejResponseTo is a required field for OrderCancelReject.
	CxlRejResponseTo string `fix:"434"`
	//CxlRejReason is a non-required field for OrderCancelReject.
	CxlRejReason *int `fix:"102"`
	//Text is a non-required field for OrderCancelReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for OrderCancelReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for OrderCancelReject.
	EncodedText *string `fix:"355"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetOrderID(v string)              { m.OrderID = v }
func (m *Message) SetSecondaryOrderID(v string)     { m.SecondaryOrderID = &v }
func (m *Message) SetSecondaryClOrdID(v string)     { m.SecondaryClOrdID = &v }
func (m *Message) SetClOrdID(v string)              { m.ClOrdID = v }
func (m *Message) SetClOrdLinkID(v string)          { m.ClOrdLinkID = &v }
func (m *Message) SetOrigClOrdID(v string)          { m.OrigClOrdID = v }
func (m *Message) SetOrdStatus(v string)            { m.OrdStatus = v }
func (m *Message) SetWorkingIndicator(v bool)       { m.WorkingIndicator = &v }
func (m *Message) SetOrigOrdModTime(v time.Time)    { m.OrigOrdModTime = &v }
func (m *Message) SetListID(v string)               { m.ListID = &v }
func (m *Message) SetAccount(v string)              { m.Account = &v }
func (m *Message) SetAccountType(v int)             { m.AccountType = &v }
func (m *Message) SetTradeOriginationDate(v string) { m.TradeOriginationDate = &v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *Message) SetCxlRejResponseTo(v string)     { m.CxlRejResponseTo = v }
func (m *Message) SetCxlRejReason(v int)            { m.CxlRejReason = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }

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
	return enum.BeginStringFIX43, "9", r
}
