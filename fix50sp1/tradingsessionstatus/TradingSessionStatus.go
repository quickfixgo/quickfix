//Package tradingsessionstatus msg type = h.
package tradingsessionstatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a TradingSessionStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"h"`
	Header     fixt11.Header
	//TradSesReqID is a non-required field for TradingSessionStatus.
	TradSesReqID *string `fix:"335"`
	//TradingSessionID is a required field for TradingSessionStatus.
	TradingSessionID string `fix:"336"`
	//TradingSessionSubID is a non-required field for TradingSessionStatus.
	TradingSessionSubID *string `fix:"625"`
	//TradSesMethod is a non-required field for TradingSessionStatus.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for TradingSessionStatus.
	TradSesMode *int `fix:"339"`
	//UnsolicitedIndicator is a non-required field for TradingSessionStatus.
	UnsolicitedIndicator *bool `fix:"325"`
	//TradSesStatus is a required field for TradingSessionStatus.
	TradSesStatus int `fix:"340"`
	//TradSesStatusRejReason is a non-required field for TradingSessionStatus.
	TradSesStatusRejReason *int `fix:"567"`
	//TradSesStartTime is a non-required field for TradingSessionStatus.
	TradSesStartTime *time.Time `fix:"341"`
	//TradSesOpenTime is a non-required field for TradingSessionStatus.
	TradSesOpenTime *time.Time `fix:"342"`
	//TradSesPreCloseTime is a non-required field for TradingSessionStatus.
	TradSesPreCloseTime *time.Time `fix:"343"`
	//TradSesCloseTime is a non-required field for TradingSessionStatus.
	TradSesCloseTime *time.Time `fix:"344"`
	//TradSesEndTime is a non-required field for TradingSessionStatus.
	TradSesEndTime *time.Time `fix:"345"`
	//TotalVolumeTraded is a non-required field for TradingSessionStatus.
	TotalVolumeTraded *float64 `fix:"387"`
	//Text is a non-required field for TradingSessionStatus.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for TradingSessionStatus.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for TradingSessionStatus.
	EncodedText *string `fix:"355"`
	//Instrument Component
	Instrument instrument.Component
	//MarketID is a non-required field for TradingSessionStatus.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for TradingSessionStatus.
	MarketSegmentID *string `fix:"1300"`
	//TradSesEvent is a non-required field for TradingSessionStatus.
	TradSesEvent *int `fix:"1368"`
	//ApplicationSequenceControl Component
	ApplicationSequenceControl applicationsequencecontrol.Component
	Trailer                    fixt11.Trailer
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
	return enum.ApplVerID_FIX50SP1, "h", r
}
