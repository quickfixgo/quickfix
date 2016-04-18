//Package securitylistupdatereport msg type = BK.
package securitylistupdatereport

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/seclstupdrelsymgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SecurityListUpdateReport FIX Message
type Message struct {
	FIXMsgType string `fix:"BK"`
	fixt11.Header
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
	//SecLstUpdRelSymGrp is a non-required component for SecurityListUpdateReport.
	SecLstUpdRelSymGrp *seclstupdrelsymgrp.SecLstUpdRelSymGrp
	//MarketID is a non-required field for SecurityListUpdateReport.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityListUpdateReport.
	MarketSegmentID *string `fix:"1300"`
	//ApplicationSequenceControl is a non-required component for SecurityListUpdateReport.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	//SecurityListID is a non-required field for SecurityListUpdateReport.
	SecurityListID *string `fix:"1465"`
	//SecurityListRefID is a non-required field for SecurityListUpdateReport.
	SecurityListRefID *string `fix:"1466"`
	//SecurityListDesc is a non-required field for SecurityListUpdateReport.
	SecurityListDesc *string `fix:"1467"`
	//EncodedSecurityListDescLen is a non-required field for SecurityListUpdateReport.
	EncodedSecurityListDescLen *int `fix:"1468"`
	//EncodedSecurityListDesc is a non-required field for SecurityListUpdateReport.
	EncodedSecurityListDesc *string `fix:"1469"`
	//SecurityListType is a non-required field for SecurityListUpdateReport.
	SecurityListType *int `fix:"1470"`
	//SecurityListTypeSource is a non-required field for SecurityListUpdateReport.
	SecurityListTypeSource *int `fix:"1471"`
	//TransactTime is a non-required field for SecurityListUpdateReport.
	TransactTime *time.Time `fix:"60"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized SecurityListUpdateReport instance
func New() *Message {
	var m Message
	return &m
}

func (m *Message) SetSecurityReportID(v int)        { m.SecurityReportID = &v }
func (m *Message) SetSecurityReqID(v string)        { m.SecurityReqID = &v }
func (m *Message) SetSecurityResponseID(v string)   { m.SecurityResponseID = &v }
func (m *Message) SetSecurityRequestResult(v int)   { m.SecurityRequestResult = &v }
func (m *Message) SetTotNoRelatedSym(v int)         { m.TotNoRelatedSym = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetSecurityUpdateAction(v string) { m.SecurityUpdateAction = &v }
func (m *Message) SetCorporateAction(v string)      { m.CorporateAction = &v }
func (m *Message) SetLastFragment(v bool)           { m.LastFragment = &v }
func (m *Message) SetSecLstUpdRelSymGrp(v seclstupdrelsymgrp.SecLstUpdRelSymGrp) {
	m.SecLstUpdRelSymGrp = &v
}
func (m *Message) SetMarketID(v string)        { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string) { m.MarketSegmentID = &v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}
func (m *Message) SetSecurityListID(v string)          { m.SecurityListID = &v }
func (m *Message) SetSecurityListRefID(v string)       { m.SecurityListRefID = &v }
func (m *Message) SetSecurityListDesc(v string)        { m.SecurityListDesc = &v }
func (m *Message) SetEncodedSecurityListDescLen(v int) { m.EncodedSecurityListDescLen = &v }
func (m *Message) SetEncodedSecurityListDesc(v string) { m.EncodedSecurityListDesc = &v }
func (m *Message) SetSecurityListType(v int)           { m.SecurityListType = &v }
func (m *Message) SetSecurityListTypeSource(v int)     { m.SecurityListTypeSource = &v }
func (m *Message) SetTransactTime(v time.Time)         { m.TransactTime = &v }

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
	return enum.ApplVerID_FIX50SP2, "BK", r
}
