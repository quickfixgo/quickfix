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
