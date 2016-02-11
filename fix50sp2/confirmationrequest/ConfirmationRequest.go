//Package confirmationrequest msg type = BH.
package confirmationrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/ordallocgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a ConfirmationRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"BH"`
	Header     fixt11.Header
	//ConfirmReqID is a required field for ConfirmationRequest.
	ConfirmReqID string `fix:"859"`
	//ConfirmType is a required field for ConfirmationRequest.
	ConfirmType int `fix:"773"`
	//OrdAllocGrp Component
	OrdAllocGrp ordallocgrp.Component
	//AllocID is a non-required field for ConfirmationRequest.
	AllocID *string `fix:"70"`
	//SecondaryAllocID is a non-required field for ConfirmationRequest.
	SecondaryAllocID *string `fix:"793"`
	//IndividualAllocID is a non-required field for ConfirmationRequest.
	IndividualAllocID *string `fix:"467"`
	//TransactTime is a required field for ConfirmationRequest.
	TransactTime time.Time `fix:"60"`
	//AllocAccount is a non-required field for ConfirmationRequest.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for ConfirmationRequest.
	AllocAcctIDSource *int `fix:"661"`
	//AllocAccountType is a non-required field for ConfirmationRequest.
	AllocAccountType *int `fix:"798"`
	//Text is a non-required field for ConfirmationRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for ConfirmationRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for ConfirmationRequest.
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
	return enum.ApplVerID_FIX50SP2, "BH", r
}
