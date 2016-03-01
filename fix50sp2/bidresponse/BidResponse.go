//Package bidresponse msg type = l.
package bidresponse

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/bidcomprspgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a BidResponse FIX Message
type Message struct {
	FIXMsgType string `fix:"l"`
	Header     fixt11.Header
	//BidID is a non-required field for BidResponse.
	BidID *string `fix:"390"`
	//ClientBidID is a non-required field for BidResponse.
	ClientBidID *string `fix:"391"`
	//BidCompRspGrp Component
	BidCompRspGrp bidcomprspgrp.Component
	Trailer       fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP2, "l", r
}
