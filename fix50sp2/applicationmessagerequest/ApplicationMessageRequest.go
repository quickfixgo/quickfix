//Package applicationmessagerequest msg type = BW.
package applicationmessagerequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applidrequestgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ApplicationMessageRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BW"`
	fixt11.Header
	//ApplReqID is a required field for ApplicationMessageRequest.
	ApplReqID string `fix:"1346"`
	//ApplReqType is a required field for ApplicationMessageRequest.
	ApplReqType int `fix:"1347"`
	//ApplIDRequestGrp Component
	applidrequestgrp.ApplIDRequestGrp
	//Text is a non-required field for ApplicationMessageRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ApplicationMessageRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ApplicationMessageRequest.
	EncodedText *string `fix:"355"`
	//Parties Component
	parties.Parties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetApplReqID(v string)   { m.ApplReqID = v }
func (m *Message) SetApplReqType(v int)    { m.ApplReqType = v }
func (m *Message) SetText(v string)        { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int) { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string) { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP2, "BW", r
}
