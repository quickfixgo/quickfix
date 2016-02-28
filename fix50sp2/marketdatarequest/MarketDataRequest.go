//Package marketdatarequest msg type = V.
package marketdatarequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp2/instrmtmdreqgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/mdreqgrp"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/trdgsesgrp"
	"github.com/quickfixgo/quickfix/fixt11"
)

//Message is a MarketDataRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"V"`
	Header     fixt11.Header
	//MDReqID is a required field for MarketDataRequest.
	MDReqID string `fix:"262"`
	//SubscriptionRequestType is a required field for MarketDataRequest.
	SubscriptionRequestType string `fix:"263"`
	//MarketDepth is a required field for MarketDataRequest.
	MarketDepth int `fix:"264"`
	//MDUpdateType is a non-required field for MarketDataRequest.
	MDUpdateType *int `fix:"265"`
	//AggregatedBook is a non-required field for MarketDataRequest.
	AggregatedBook *bool `fix:"266"`
	//OpenCloseSettlFlag is a non-required field for MarketDataRequest.
	OpenCloseSettlFlag *string `fix:"286"`
	//Scope is a non-required field for MarketDataRequest.
	Scope *string `fix:"546"`
	//MDImplicitDelete is a non-required field for MarketDataRequest.
	MDImplicitDelete *bool `fix:"547"`
	//MDReqGrp Component
	MDReqGrp mdreqgrp.Component
	//InstrmtMDReqGrp Component
	InstrmtMDReqGrp instrmtmdreqgrp.Component
	//TrdgSesGrp Component
	TrdgSesGrp trdgsesgrp.Component
	//ApplQueueAction is a non-required field for MarketDataRequest.
	ApplQueueAction *int `fix:"815"`
	//ApplQueueMax is a non-required field for MarketDataRequest.
	ApplQueueMax *int `fix:"812"`
	//MDQuoteType is a non-required field for MarketDataRequest.
	MDQuoteType *int `fix:"1070"`
	//Parties Component
	Parties parties.Component
	Trailer fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMDReqID(v string)                 { m.MDReqID = v }
func (m *Message) SetSubscriptionRequestType(v string) { m.SubscriptionRequestType = v }
func (m *Message) SetMarketDepth(v int)                { m.MarketDepth = v }
func (m *Message) SetMDUpdateType(v int)               { m.MDUpdateType = &v }
func (m *Message) SetAggregatedBook(v bool)            { m.AggregatedBook = &v }
func (m *Message) SetOpenCloseSettlFlag(v string)      { m.OpenCloseSettlFlag = &v }
func (m *Message) SetScope(v string)                   { m.Scope = &v }
func (m *Message) SetMDImplicitDelete(v bool)          { m.MDImplicitDelete = &v }
func (m *Message) SetApplQueueAction(v int)            { m.ApplQueueAction = &v }
func (m *Message) SetApplQueueMax(v int)               { m.ApplQueueMax = &v }
func (m *Message) SetMDQuoteType(v int)                { m.MDQuoteType = &v }

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
	return enum.ApplVerID_FIX50SP2, "V", r
}
