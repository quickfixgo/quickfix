//Package allocationack msg type = P.
package allocationack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/parties"
	"time"
)

//Message is a AllocationAck FIX Message
type Message struct {
	FIXMsgType string `fix:"P"`
	fix43.Header
	//Parties is a non-required component for AllocationAck.
	Parties *parties.Parties
	//AllocID is a required field for AllocationAck.
	AllocID string `fix:"70"`
	//TradeDate is a required field for AllocationAck.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationAck.
	TransactTime *time.Time `fix:"60"`
	//AllocStatus is a required field for AllocationAck.
	AllocStatus int `fix:"87"`
	//AllocRejCode is a non-required field for AllocationAck.
	AllocRejCode *int `fix:"88"`
	//Text is a non-required field for AllocationAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationAck.
	EncodedText *string `fix:"355"`
	//LegalConfirm is a non-required field for AllocationAck.
	LegalConfirm *bool `fix:"650"`
	fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized AllocationAck instance
func New(allocid string, tradedate string, allocstatus int) *Message {
	var m Message
	m.SetAllocID(allocid)
	m.SetTradeDate(tradedate)
	m.SetAllocStatus(allocstatus)
	return &m
}

func (m *Message) SetParties(v parties.Parties) { m.Parties = &v }
func (m *Message) SetAllocID(v string)          { m.AllocID = v }
func (m *Message) SetTradeDate(v string)        { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)  { m.TransactTime = &v }
func (m *Message) SetAllocStatus(v int)         { m.AllocStatus = v }
func (m *Message) SetAllocRejCode(v int)        { m.AllocRejCode = &v }
func (m *Message) SetText(v string)             { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)      { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)      { m.EncodedText = &v }
func (m *Message) SetLegalConfirm(v bool)       { m.LegalConfirm = &v }

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
	return enum.BeginStringFIX43, "P", r
}
