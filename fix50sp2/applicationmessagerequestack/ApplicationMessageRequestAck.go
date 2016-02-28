//Package applicationmessagerequestack msg type = BX.
package applicationmessagerequestack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applidrequestackgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a ApplicationMessageRequestAck FIX Message
type Message struct {
	FIXMsgType string `fix:"BX"`
	Header     fixt11.Header
	//ApplResponseID is a required field for ApplicationMessageRequestAck.
	ApplResponseID string `fix:"1353"`
	//ApplReqID is a non-required field for ApplicationMessageRequestAck.
	ApplReqID *string `fix:"1346"`
	//ApplReqType is a non-required field for ApplicationMessageRequestAck.
	ApplReqType *int `fix:"1347"`
	//ApplResponseType is a non-required field for ApplicationMessageRequestAck.
	ApplResponseType *int `fix:"1348"`
	//ApplTotalMessageCount is a non-required field for ApplicationMessageRequestAck.
	ApplTotalMessageCount *int `fix:"1349"`
	//ApplIDRequestAckGrp Component
	ApplIDRequestAckGrp applidrequestackgrp.Component
	//Text is a non-required field for ApplicationMessageRequestAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ApplicationMessageRequestAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ApplicationMessageRequestAck.
	EncodedText *string `fix:"355"`
	//Parties Component
	Parties parties.Component
	Trailer fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetApplResponseID(v string)     { m.ApplResponseID = v }
func (m *Message) SetApplReqID(v string)          { m.ApplReqID = &v }
func (m *Message) SetApplReqType(v int)           { m.ApplReqType = &v }
func (m *Message) SetApplResponseType(v int)      { m.ApplResponseType = &v }
func (m *Message) SetApplTotalMessageCount(v int) { m.ApplTotalMessageCount = &v }
func (m *Message) SetText(v string)               { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)        { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)        { m.EncodedText = &v }

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
	return enum.ApplVerID_FIX50SP2, "BX", r
}
