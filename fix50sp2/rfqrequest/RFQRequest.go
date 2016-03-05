//Package rfqrequest msg type = AH.
package rfqrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/rfqreqgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a RFQRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"AH"`
	fixt11.Header
	//RFQReqID is a required field for RFQRequest.
	RFQReqID string `fix:"644"`
	//RFQReqGrp Component
	rfqreqgrp.RFQReqGrp
	//SubscriptionRequestType is a non-required field for RFQRequest.
	SubscriptionRequestType *string `fix:"263"`
	//PrivateQuote is a non-required field for RFQRequest.
	PrivateQuote *bool `fix:"1171"`
	//Parties Component
	parties.Parties
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetRFQReqID(v string)                { m.RFQReqID = v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = &v }
func (m *Message) SetPrivateQuote(v bool)              { m.PrivateQuote = &v }

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
	return enum.ApplVerID_FIX50SP2, "AH", r
}
