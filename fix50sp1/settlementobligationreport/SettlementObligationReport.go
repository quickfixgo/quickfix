//Package settlementobligationreport msg type = BQ.
package settlementobligationreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/settlobligationinstructions"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SettlementObligationReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BQ"`
	fixt11.Header
	//ClearingBusinessDate is a non-required field for SettlementObligationReport.
	ClearingBusinessDate *string `fix:"715"`
	//SettlementCycleNo is a non-required field for SettlementObligationReport.
	SettlementCycleNo *int `fix:"1153"`
	//SettlObligMsgID is a required field for SettlementObligationReport.
	SettlObligMsgID string `fix:"1160"`
	//SettlObligMode is a required field for SettlementObligationReport.
	SettlObligMode int `fix:"1159"`
	//Text is a non-required field for SettlementObligationReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SettlementObligationReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SettlementObligationReport.
	EncodedText *string `fix:"355"`
	//TransactTime is a non-required field for SettlementObligationReport.
	TransactTime *time.Time `fix:"60"`
	//SettlObligationInstructions is a required component for SettlementObligationReport.
	settlobligationinstructions.SettlObligationInstructions
	//ApplicationSequenceControl is a non-required component for SettlementObligationReport.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SettlementObligationReport instance
func New(settlobligmsgid string, settlobligmode int, settlobligationinstructions settlobligationinstructions.SettlObligationInstructions) *Message {
	var m Message
	m.SetSettlObligMsgID(settlobligmsgid)
	m.SetSettlObligMode(settlobligmode)
	m.SetSettlObligationInstructions(settlobligationinstructions)
	return &m
}

func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetSettlementCycleNo(v int)       { m.SettlementCycleNo = &v }
func (m *Message) SetSettlObligMsgID(v string)      { m.SettlObligMsgID = v }
func (m *Message) SetSettlObligMode(v int)          { m.SettlObligMode = v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *Message) SetSettlObligationInstructions(v settlobligationinstructions.SettlObligationInstructions) {
	m.SettlObligationInstructions = v
}
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}

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
	return enum.ApplVerID_FIX50SP1, "BQ", r
}
