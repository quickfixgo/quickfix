//Package quoterequestreject msg type = AG.
package quoterequestreject

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/quotreqrjctgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/rootparties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a QuoteRequestReject FIX Message
type Message struct {
	FIXMsgType string `fix:"AG"`
	Header     fixt11.Header
	//QuoteReqID is a required field for QuoteRequestReject.
	QuoteReqID string `fix:"131"`
	//RFQReqID is a non-required field for QuoteRequestReject.
	RFQReqID *string `fix:"644"`
	//QuoteRequestRejectReason is a required field for QuoteRequestReject.
	QuoteRequestRejectReason int `fix:"658"`
	//QuotReqRjctGrp Component
	QuotReqRjctGrp quotreqrjctgrp.Component
	//Text is a non-required field for QuoteRequestReject.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for QuoteRequestReject.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for QuoteRequestReject.
	EncodedText *string `fix:"355"`
	//RootParties Component
	RootParties rootparties.Component
	//PrivateQuote is a non-required field for QuoteRequestReject.
	PrivateQuote *bool `fix:"1171"`
	//RespondentType is a non-required field for QuoteRequestReject.
	RespondentType *int `fix:"1172"`
	//PreTradeAnonymity is a non-required field for QuoteRequestReject.
	PreTradeAnonymity *bool `fix:"1091"`
	Trailer           fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetQuoteReqID(v string)            { m.QuoteReqID = v }
func (m *Message) SetRFQReqID(v string)              { m.RFQReqID = &v }
func (m *Message) SetQuoteRequestRejectReason(v int) { m.QuoteRequestRejectReason = v }
func (m *Message) SetText(v string)                  { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)           { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)           { m.EncodedText = &v }
func (m *Message) SetPrivateQuote(v bool)            { m.PrivateQuote = &v }
func (m *Message) SetRespondentType(v int)           { m.RespondentType = &v }
func (m *Message) SetPreTradeAnonymity(v bool)       { m.PreTradeAnonymity = &v }

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
	return enum.ApplVerID_FIX50SP2, "AG", r
}
