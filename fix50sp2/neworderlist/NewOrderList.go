//Package neworderlist msg type = E.
package neworderlist

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/listordgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/rootparties"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a NewOrderList FIX Message
type Message struct {
	FIXMsgType string `fix:"E"`
	fixt11.Header
	//ListID is a required field for NewOrderList.
	ListID string `fix:"66"`
	//BidID is a non-required field for NewOrderList.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for NewOrderList.
	ClientBidID *string `fix:"391"`
	//ProgRptReqs is a non-required field for NewOrderList.
	ProgRptReqs *int `fix:"414"`
	//BidType is a required field for NewOrderList.
	BidType int `fix:"394"`
	//ProgPeriodInterval is a non-required field for NewOrderList.
	ProgPeriodInterval *int `fix:"415"`
	//CancellationRights is a non-required field for NewOrderList.
	CancellationRights *string `fix:"480"`
	//MoneyLaunderingStatus is a non-required field for NewOrderList.
	MoneyLaunderingStatus *string `fix:"481"`
	//RegistID is a non-required field for NewOrderList.
	RegistID *string `fix:"513"`
	//ListExecInstType is a non-required field for NewOrderList.
	ListExecInstType *string `fix:"433"`
	//ListExecInst is a non-required field for NewOrderList.
	ListExecInst *string `fix:"69"`
	//EncodedListExecInstLen is a non-required field for NewOrderList.
	EncodedListExecInstLen *int `fix:"352"`
	//EncodedListExecInst is a non-required field for NewOrderList.
	EncodedListExecInst *string `fix:"353"`
	//AllowableOneSidednessPct is a non-required field for NewOrderList.
	AllowableOneSidednessPct *float64 `fix:"765"`
	//AllowableOneSidednessValue is a non-required field for NewOrderList.
	AllowableOneSidednessValue *float64 `fix:"766"`
	//AllowableOneSidednessCurr is a non-required field for NewOrderList.
	AllowableOneSidednessCurr *string `fix:"767"`
	//TotNoOrders is a required field for NewOrderList.
	TotNoOrders int `fix:"68"`
	//LastFragment is a non-required field for NewOrderList.
	LastFragment *bool `fix:"893"`
	//ListOrdGrp is a required component for NewOrderList.
	listordgrp.ListOrdGrp
	//RootParties is a non-required component for NewOrderList.
	RootParties *rootparties.RootParties
	//ContingencyType is a non-required field for NewOrderList.
	ContingencyType *int `fix:"1385"`
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetListID(v string)                       { m.ListID = v }
func (m *Message) SetBidID(v string)                        { m.BidID = &v }
func (m *Message) SetClientBidID(v string)                  { m.ClientBidID = &v }
func (m *Message) SetProgRptReqs(v int)                     { m.ProgRptReqs = &v }
func (m *Message) SetBidType(v int)                         { m.BidType = v }
func (m *Message) SetProgPeriodInterval(v int)              { m.ProgPeriodInterval = &v }
func (m *Message) SetCancellationRights(v string)           { m.CancellationRights = &v }
func (m *Message) SetMoneyLaunderingStatus(v string)        { m.MoneyLaunderingStatus = &v }
func (m *Message) SetRegistID(v string)                     { m.RegistID = &v }
func (m *Message) SetListExecInstType(v string)             { m.ListExecInstType = &v }
func (m *Message) SetListExecInst(v string)                 { m.ListExecInst = &v }
func (m *Message) SetEncodedListExecInstLen(v int)          { m.EncodedListExecInstLen = &v }
func (m *Message) SetEncodedListExecInst(v string)          { m.EncodedListExecInst = &v }
func (m *Message) SetAllowableOneSidednessPct(v float64)    { m.AllowableOneSidednessPct = &v }
func (m *Message) SetAllowableOneSidednessValue(v float64)  { m.AllowableOneSidednessValue = &v }
func (m *Message) SetAllowableOneSidednessCurr(v string)    { m.AllowableOneSidednessCurr = &v }
func (m *Message) SetTotNoOrders(v int)                     { m.TotNoOrders = v }
func (m *Message) SetLastFragment(v bool)                   { m.LastFragment = &v }
func (m *Message) SetListOrdGrp(v listordgrp.ListOrdGrp)    { m.ListOrdGrp = v }
func (m *Message) SetRootParties(v rootparties.RootParties) { m.RootParties = &v }
func (m *Message) SetContingencyType(v int)                 { m.ContingencyType = &v }

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
	return enum.ApplVerID_FIX50SP2, "E", r
}
