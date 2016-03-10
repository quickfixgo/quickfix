//Package allocationack msg type = P.
package allocationack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix40"
	"time"
)

//Message is a AllocationACK FIX Message
type Message struct {
	FIXMsgType string `fix:"P"`
	fix40.Header
	//ClientID is a non-required field for AllocationACK.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for AllocationACK.
	ExecBroker *string `fix:"76"`
	//AllocID is a required field for AllocationACK.
	AllocID int `fix:"70"`
	//TradeDate is a required field for AllocationACK.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for AllocationACK.
	TransactTime *time.Time `fix:"60"`
	//AllocStatus is a required field for AllocationACK.
	AllocStatus int `fix:"87"`
	//AllocRejCode is a non-required field for AllocationACK.
	AllocRejCode *int `fix:"88"`
	//Text is a non-required field for AllocationACK.
	Text *string `fix:"58"`
	fix40.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetClientID(v string)        { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)      { m.ExecBroker = &v }
func (m *Message) SetAllocID(v int)            { m.AllocID = v }
func (m *Message) SetTradeDate(v string)       { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time) { m.TransactTime = &v }
func (m *Message) SetAllocStatus(v int)        { m.AllocStatus = v }
func (m *Message) SetAllocRejCode(v int)       { m.AllocRejCode = &v }
func (m *Message) SetText(v string)            { m.Text = &v }

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
	return enum.BeginStringFIX40, "P", r
}
