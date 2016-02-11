//Package securitylist msg type = y.
package securitylist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/seclistgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SecurityList FIX Message
type Message struct {
	FIXMsgType string `fix:"y"`
	Header     fixt11.Header
	//SecurityReqID is a non-required field for SecurityList.
	SecurityReqID *string `fix:"320"`
	//SecurityResponseID is a non-required field for SecurityList.
	SecurityResponseID *string `fix:"322"`
	//SecurityRequestResult is a non-required field for SecurityList.
	SecurityRequestResult *int `fix:"560"`
	//TotNoRelatedSym is a non-required field for SecurityList.
	TotNoRelatedSym *int `fix:"393"`
	//LastFragment is a non-required field for SecurityList.
	LastFragment *bool `fix:"893"`
	//SecListGrp Component
	SecListGrp seclistgrp.Component
	//SecurityReportID is a non-required field for SecurityList.
	SecurityReportID *int `fix:"964"`
	//ClearingBusinessDate is a non-required field for SecurityList.
	ClearingBusinessDate *string `fix:"715"`
	//MarketID is a non-required field for SecurityList.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityList.
	MarketSegmentID *string `fix:"1300"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	//SecurityListID is a non-required field for SecurityList.
	SecurityListID *string `fix:"1465"`
	//SecurityListRefID is a non-required field for SecurityList.
	SecurityListRefID *string `fix:"1466"`
	//SecurityListDesc is a non-required field for SecurityList.
	SecurityListDesc *string `fix:"1467"`
	//EncodedSecurityListDescLen is a non-required field for SecurityList.
	EncodedSecurityListDescLen *int `fix:"1468"`
	//EncodedSecurityListDesc is a non-required field for SecurityList.
	EncodedSecurityListDesc *string `fix:"1469"`
	//SecurityListType is a non-required field for SecurityList.
	SecurityListType *int `fix:"1470"`
	//SecurityListTypeSource is a non-required field for SecurityList.
	SecurityListTypeSource *int `fix:"1471"`
	//TransactTime is a non-required field for SecurityList.
	TransactTime *time.Time `fix:"60"`
	Trailer      fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "y", r
}
