//Package contraryintentionreport msg type = BO.
package contraryintentionreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/expirationqty"
	"github.com/quickfixgo/quickfix/fix50sp2/instrument"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ContraryIntentionReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BO"`
	fixt11.Header
	//ContIntRptID is a required field for ContraryIntentionReport.
	ContIntRptID string `fix:"977"`
	//TransactTime is a non-required field for ContraryIntentionReport.
	TransactTime *time.Time `fix:"60"`
	//LateIndicator is a non-required field for ContraryIntentionReport.
	LateIndicator *bool `fix:"978"`
	//InputSource is a non-required field for ContraryIntentionReport.
	InputSource *string `fix:"979"`
	//ClearingBusinessDate is a required field for ContraryIntentionReport.
	ClearingBusinessDate string `fix:"715"`
	//Parties is a required component for ContraryIntentionReport.
	parties.Parties
	//ExpirationQty is a required component for ContraryIntentionReport.
	expirationqty.ExpirationQty
	//Instrument is a required component for ContraryIntentionReport.
	instrument.Instrument
	//Text is a non-required field for ContraryIntentionReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ContraryIntentionReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ContraryIntentionReport.
	EncodedText *string `fix:"355"`
	//UndInstrmtGrp is a non-required component for ContraryIntentionReport.
	UndInstrmtGrp *undinstrmtgrp.UndInstrmtGrp
	//ApplicationSequenceControl is a non-required component for ContraryIntentionReport.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized ContraryIntentionReport instance
func New(contintrptid string, clearingbusinessdate string, parties parties.Parties, expirationqty expirationqty.ExpirationQty, instrument instrument.Instrument) *Message {
	var m Message
	m.SetContIntRptID(contintrptid)
	m.SetClearingBusinessDate(clearingbusinessdate)
	m.SetParties(parties)
	m.SetExpirationQty(expirationqty)
	m.SetInstrument(instrument)
	return &m
}

func (m *Message) SetContIntRptID(v string)                       { m.ContIntRptID = v }
func (m *Message) SetTransactTime(v time.Time)                    { m.TransactTime = &v }
func (m *Message) SetLateIndicator(v bool)                        { m.LateIndicator = &v }
func (m *Message) SetInputSource(v string)                        { m.InputSource = &v }
func (m *Message) SetClearingBusinessDate(v string)               { m.ClearingBusinessDate = v }
func (m *Message) SetParties(v parties.Parties)                   { m.Parties = v }
func (m *Message) SetExpirationQty(v expirationqty.ExpirationQty) { m.ExpirationQty = v }
func (m *Message) SetInstrument(v instrument.Instrument)          { m.Instrument = v }
func (m *Message) SetText(v string)                               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                        { m.EncodedText = &v }
func (m *Message) SetUndInstrmtGrp(v undinstrmtgrp.UndInstrmtGrp) { m.UndInstrmtGrp = &v }
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
	return enum.ApplVerID_FIX50SP2, "BO", r
}
