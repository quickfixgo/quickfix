//Package securitylistupdatereport msg type = BK.
package securitylistupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/seclstupdrelsymgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a SecurityListUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BK"`
	Header     fixt11.Header
	//SecurityReportID is a non-required field for SecurityListUpdateReport.
	SecurityReportID *int `fix:"964"`
	//SecurityReqID is a non-required field for SecurityListUpdateReport.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityListUpdateReport.
	SecurityResponseID *string `fix:"322"`
	//SecurityRequestResult is a non-required field for SecurityListUpdateReport.
	SecurityRequestResult *int `fix:"560"`
	//TotNoRelatedSym is a non-required field for SecurityListUpdateReport.
	TotNoRelatedSym *int `fix:"393"`
	//ClearingBusinessDate is a non-required field for SecurityListUpdateReport.
	ClearingBusinessDate *string `fix:"715"`
	//SecurityUpdateAction is a non-required field for SecurityListUpdateReport.
	SecurityUpdateAction *string `fix:"980"`
	//CorporateAction is a non-required field for SecurityListUpdateReport.
	CorporateAction *string `fix:"292"`
	//LastFragment is a non-required field for SecurityListUpdateReport.
	LastFragment *bool `fix:"893"`
	//SecLstUpdRelSymGrp Component
	SecLstUpdRelSymGrp seclstupdrelsymgrp.Component
	//MarketID is a non-required field for SecurityListUpdateReport.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityListUpdateReport.
	MarketSegmentID *string `fix:"1300"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

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
	return enum.ApplVerID_FIX50SP1, "BK", r
}
