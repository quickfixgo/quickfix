//Package marketdatarequest msg type = V.
package marketdatarequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix44"
	"github.com/quickfixgo/quickfix/fix44/instrument"
	"github.com/quickfixgo/quickfix/fix44/instrumentleg"
	"github.com/quickfixgo/quickfix/fix44/underlyinginstrument"
)

//NoMDEntryTypes is a repeating group in MarketDataRequest
type NoMDEntryTypes struct {
	//MDEntryType is a required field for NoMDEntryTypes.
	MDEntryType string `fix:"269"`
}

func (m *NoMDEntryTypes) SetMDEntryType(v string) { m.MDEntryType = v }

//NoRelatedSym is a repeating group in MarketDataRequest
type NoRelatedSym struct {
	//Instrument Component
	Instrument instrument.Component
	//NoUnderlyings is a non-required field for NoRelatedSym.
	NoUnderlyings []NoUnderlyings `fix:"711,omitempty"`
	//NoLegs is a non-required field for NoRelatedSym.
	NoLegs []NoLegs `fix:"555,omitempty"`
	//NoTradingSessions is a non-required field for NoRelatedSym.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ApplQueueAction is a non-required field for NoRelatedSym.
	ApplQueueAction *int `fix:"815"`
	//ApplQueueMax is a non-required field for NoRelatedSym.
	ApplQueueMax *int `fix:"812"`
}

func (m *NoRelatedSym) SetNoUnderlyings(v []NoUnderlyings)         { m.NoUnderlyings = v }
func (m *NoRelatedSym) SetNoLegs(v []NoLegs)                       { m.NoLegs = v }
func (m *NoRelatedSym) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *NoRelatedSym) SetApplQueueAction(v int)                   { m.ApplQueueAction = &v }
func (m *NoRelatedSym) SetApplQueueMax(v int)                      { m.ApplQueueMax = &v }

//NoUnderlyings is a repeating group in NoRelatedSym
type NoUnderlyings struct {
	//UnderlyingInstrument Component
	UnderlyingInstrument underlyinginstrument.Component
}

//NoLegs is a repeating group in NoRelatedSym
type NoLegs struct {
	//InstrumentLeg Component
	InstrumentLeg instrumentleg.Component
}

//NoTradingSessions is a repeating group in NoRelatedSym
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

func (m *NoTradingSessions) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }

//Message is a MarketDataRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"V"`
	Header     fix44.Header
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
	//NoMDEntryTypes is a required field for MarketDataRequest.
	NoMDEntryTypes []NoMDEntryTypes `fix:"267"`
	//NoRelatedSym is a required field for MarketDataRequest.
	NoRelatedSym []NoRelatedSym `fix:"146"`
	Trailer      fix44.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetMDReqID(v string)                  { m.MDReqID = v }
func (m *Message) SetSubscriptionRequestType(v string)  { m.SubscriptionRequestType = v }
func (m *Message) SetMarketDepth(v int)                 { m.MarketDepth = v }
func (m *Message) SetMDUpdateType(v int)                { m.MDUpdateType = &v }
func (m *Message) SetAggregatedBook(v bool)             { m.AggregatedBook = &v }
func (m *Message) SetOpenCloseSettlFlag(v string)       { m.OpenCloseSettlFlag = &v }
func (m *Message) SetScope(v string)                    { m.Scope = &v }
func (m *Message) SetMDImplicitDelete(v bool)           { m.MDImplicitDelete = &v }
func (m *Message) SetNoMDEntryTypes(v []NoMDEntryTypes) { m.NoMDEntryTypes = v }
func (m *Message) SetNoRelatedSym(v []NoRelatedSym)     { m.NoRelatedSym = v }

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
	return enum.BeginStringFIX44, "V", r
}
