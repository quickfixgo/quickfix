//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp2/mdincgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/routinggrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MarketDataIncrementalRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"X"`
	fixt11.Header
	//MDReqID is a non-required field for MarketDataIncrementalRefresh.
	MDReqID *string `fix:"262"`
	//MDIncGrp is a required component for MarketDataIncrementalRefresh.
	mdincgrp.MDIncGrp
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
	//RoutingGrp is a non-required component for MarketDataIncrementalRefresh.
	RoutingGrp *routinggrp.RoutingGrp
	//ApplicationSequenceControl is a non-required component for MarketDataIncrementalRefresh.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMDReqID(v string)                   { m.MDReqID = &v }
func (m *Message) SetMDIncGrp(v mdincgrp.MDIncGrp)       { m.MDIncGrp = v }
func (m *Message) SetApplQueueDepth(v int)               { m.ApplQueueDepth = &v }
func (m *Message) SetApplQueueResolution(v int)          { m.ApplQueueResolution = &v }
func (m *Message) SetMDBookType(v int)                   { m.MDBookType = &v }
func (m *Message) SetMDFeedType(v string)                { m.MDFeedType = &v }
func (m *Message) SetTradeDate(v string)                 { m.TradeDate = &v }
func (m *Message) SetRoutingGrp(v routinggrp.RoutingGrp) { m.RoutingGrp = &v }
func (m *Message) SetApplicationSequenceControl(v applicationsequencecontrol.ApplicationSequenceControl) {
	m.ApplicationSequenceControl = &v
}

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
	return enum.ApplVerID_FIX50SP2, "X", r
}
