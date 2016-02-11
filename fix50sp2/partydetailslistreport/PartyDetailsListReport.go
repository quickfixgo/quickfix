//Package partydetailslistreport msg type = CG.
package partydetailslistreport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/partylistgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a PartyDetailsListReport FIX Message
type Message struct {
	FIXMsgType string `fix:"CG"`
	Header     fixt11.Header
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	//PartyDetailsListReportID is a required field for PartyDetailsListReport.
	PartyDetailsListReportID string `fix:"1510"`
	//PartyDetailsListRequestID is a non-required field for PartyDetailsListReport.
	PartyDetailsListRequestID *string `fix:"1505"`
	//PartyDetailsRequestResult is a non-required field for PartyDetailsListReport.
	PartyDetailsRequestResult *int `fix:"1511"`
	//TotNoPartyList is a non-required field for PartyDetailsListReport.
	TotNoPartyList *int `fix:"1512"`
	//LastFragment is a non-required field for PartyDetailsListReport.
	LastFragment *bool `fix:"893"`
	//PartyListGrp Component
	PartyListGrp partylistgrp.Component
	//Text is a non-required field for PartyDetailsListReport.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for PartyDetailsListReport.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for PartyDetailsListReport.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.ApplVerID_FIX50SP2, "CG", r
}
