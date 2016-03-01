//Package allocationinstructionack msg type = P.
package allocationinstructionack

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/allocackgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a AllocationInstructionAck FIX Message
type Message struct {
	FIXMsgType string `fix:"P"`
	Header     fixt11.Header
	//AllocID is a required field for AllocationInstructionAck.
	AllocID string `fix:"70"`
	//Parties Component
	Parties parties.Component
	//SecondaryAllocID is a non-required field for AllocationInstructionAck.
	SecondaryAllocID *string `fix:"793"`
	//TradeDate is a non-required field for AllocationInstructionAck.
	TradeDate *string `fix:"75"`
	//TransactTime is a non-required field for AllocationInstructionAck.
	TransactTime *time.Time `fix:"60"`
	//AllocStatus is a required field for AllocationInstructionAck.
	AllocStatus int `fix:"87"`
	//AllocRejCode is a non-required field for AllocationInstructionAck.
	AllocRejCode *int `fix:"88"`
	//AllocType is a non-required field for AllocationInstructionAck.
	AllocType *int `fix:"626"`
	//AllocIntermedReqType is a non-required field for AllocationInstructionAck.
	AllocIntermedReqType *int `fix:"808"`
	//MatchStatus is a non-required field for AllocationInstructionAck.
	MatchStatus *string `fix:"573"`
	//Product is a non-required field for AllocationInstructionAck.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for AllocationInstructionAck.
	SecurityType *string `fix:"167"`
	//Text is a non-required field for AllocationInstructionAck.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for AllocationInstructionAck.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for AllocationInstructionAck.
	EncodedText *string `fix:"355"`
	//AllocAckGrp Component
	AllocAckGrp allocackgrp.Component
	Trailer     fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "P", r
}
