//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/mdincgrp"
	"github.com/quickfixgo/quickfix/fix50/routinggrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MarketDataIncrementalRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"X"`
	Header     fixt11.Header
	//MDReqID is a non-required field for MarketDataIncrementalRefresh.
	MDReqID *string `fix:"262"`
	//MDIncGrp Component
	MDIncGrp mdincgrp.Component
	//ApplQueueDepth is a non-required field for MarketDataIncrementalRefresh.
	ApplQueueDepth *int `fix:"813"`
	//ApplQueueResolution is a non-required field for MarketDataIncrementalRefresh.
	ApplQueueResolution *int `fix:"814"`
	//MDBookType is a non-required field for MarketDataIncrementalRefresh.
	MDBookType *int `fix:"1021"`
	//MDFeedType is a non-required field for MarketDataIncrementalRefresh.
	MDFeedType *string `fix:"1022"`
	//TradeDate is a non-required field for MarketDataIncrementalRefresh.
	TradeDate *string `fix:"75"`
	//RoutingGrp Component
	RoutingGrp routinggrp.Component
	Trailer    fixt11.Trailer
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
	return enum.BeginStringFIX50, "X", r
}
