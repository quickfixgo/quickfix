//Package allocationreportack msg type = AT.
package allocationreportack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/allocackgrp"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a AllocationReportAck FIX Message
type Message struct {
	FIXMsgType string `fix:"AT"`
	fixt11.Header
	//AllocReportID is a required field for AllocationReportAck.
	AllocReportID string `fix:"755"`
	//AllocID is a required field for AllocationReportAck.
	AllocID string `fix:"70"`
	//Parties Component
	parties.Parties
	//SecondaryAllocID is a non-required field for AllocationReportAck.
	SecondaryAllocID *string `fix:"793"`
	//TradeDate is a non-required field for AllocationReportAck.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for AllocationReportAck.
	TransactTime *time.Time `fix:"60"`
	//AllocStatus is a non-required field for AllocationReportAck.
	AllocStatus *int `fix:"87"`
	//AllocRejCode is a non-required field for AllocationReportAck.
	AllocRejCode *int `fix:"88"`
	//AllocReportType is a non-required field for AllocationReportAck.
	AllocReportType *int `fix:"794"`
	//AllocIntermedReqType is a non-required field for AllocationReportAck.
	AllocIntermedReqType *int `fix:"808"`
	//MatchStatus is a non-required field for AllocationReportAck.
	MatchStatus *string `fix:"573"`
	//Product is a non-required field for AllocationReportAck.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for AllocationReportAck.
	SecurityType *string `fix:"167"`
	//Text is a non-required field for AllocationReportAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationReportAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationReportAck.
	EncodedText *string `fix:"355"`
	//AllocAckGrp Component
	allocackgrp.AllocAckGrp
	//ClearingBusinessDate is a non-required field for AllocationReportAck.
	ClearingBusinessDate *string `fix:"715"`
	//AvgPxIndicator is a non-required field for AllocationReportAck.
	AvgPxIndicator *int `fix:"819"`
	//Quantity is a non-required field for AllocationReportAck.
	Quantity *float64 `fix:"53"`
	//AllocTransType is a non-required field for AllocationReportAck.
	AllocTransType *string `fix:"71"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAllocReportID(v string)        { m.AllocReportID = v }
func (m *Message) SetAllocID(v string)              { m.AllocID = v }
func (m *Message) SetSecondaryAllocID(v string)     { m.SecondaryAllocID = &v }
func (m *Message) SetTradeDate(v string)            { m.TradeDate = &v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *Message) SetAllocStatus(v int)             { m.AllocStatus = &v }
func (m *Message) SetAllocRejCode(v int)            { m.AllocRejCode = &v }
func (m *Message) SetAllocReportType(v int)         { m.AllocReportType = &v }
func (m *Message) SetAllocIntermedReqType(v int)    { m.AllocIntermedReqType = &v }
func (m *Message) SetMatchStatus(v string)          { m.MatchStatus = &v }
func (m *Message) SetProduct(v int)                 { m.Product = &v }
func (m *Message) SetSecurityType(v string)         { m.SecurityType = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *Message) SetClearingBusinessDate(v string) { m.ClearingBusinessDate = &v }
func (m *Message) SetAvgPxIndicator(v int)          { m.AvgPxIndicator = &v }
func (m *Message) SetQuantity(v float64)            { m.Quantity = &v }
func (m *Message) SetAllocTransType(v string)       { m.AllocTransType = &v }

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
	return enum.ApplVerID_FIX50SP1, "AT", r
}
