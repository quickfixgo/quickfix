//Package rfqrequest msg type = AH.
package rfqrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/rfqreqgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a RFQRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AH"`
	Header     fixt11.Header
	//RFQReqID is a required field for RFQRequest.
	RFQReqID string `fix:"644"`
	//RFQReqGrp Component
	RFQReqGrp rfqreqgrp.Component
	//SubscriptionRequestType is a non-required field for RFQRequest.
	SubscriptionRequestType *string `fix:"263"`
	Trailer                 fixt11.Trailer
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
	return enum.BeginStringFIX50, "AH", r
}
