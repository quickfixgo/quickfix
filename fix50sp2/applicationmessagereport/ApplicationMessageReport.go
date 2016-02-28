//Package applicationmessagereport msg type = BY.
package applicationmessagereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applidreportgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ApplicationMessageReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BY"`
	Header     fixt11.Header
	//ApplReportID is a required field for ApplicationMessageReport.
	ApplReportID string `fix:"1356"`
	//ApplReportType is a required field for ApplicationMessageReport.
	ApplReportType int `fix:"1426"`
	//ApplIDReportGrp Component
	ApplIDReportGrp applidreportgrp.Component
	//Text is a non-required field for ApplicationMessageReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ApplicationMessageReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ApplicationMessageReport.
	EncodedText *string `fix:"355"`
	//ApplReqID is a non-required field for ApplicationMessageReport.
	ApplReqID *string `fix:"1346"`
	Trailer   fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetApplReportID(v string) { m.ApplReportID = v }
func (m *Message) SetApplReportType(v int)  { m.ApplReportType = v }
func (m *Message) SetText(v string)         { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)  { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)  { m.EncodedText = &v }
func (m *Message) SetApplReqID(v string)    { m.ApplReqID = &v }

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
	return enum.ApplVerID_FIX50SP2, "BY", r
}
