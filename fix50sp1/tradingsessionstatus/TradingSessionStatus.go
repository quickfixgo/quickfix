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
	fixt11.Header
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
	//Instrument is a non-required component for TradingSessionStatus.
	Instrument *instrument.Instrument
	//MarketID is a non-required field for TradingSessionStatus.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for TradingSessionStatus.
	MarketSegmentID *string `fix:"1300"`
	//TradSesEvent is a non-required field for TradingSessionStatus.
	TradSesEvent *int `fix:"1368"`
	//ApplicationSequenceControl is a non-required component for TradingSessionStatus.
	ApplicationSequenceControl *applicationsequencecontrol.ApplicationSequenceControl
	fixt11.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized TradingSessionStatus instance
func New(tradingsessionid string, tradsesstatus int) *Message {
	var m Message
	m.SetTradingSessionID(tradingsessionid)
	m.SetTradSesStatus(tradsesstatus)
	return &m
}

func (m *Message) SetTradSesReqID(v string)              { m.TradSesReqID = &v }
func (m *Message) SetTradingSessionID(v string)          { m.TradingSessionID = v }
func (m *Message) SetTradingSessionSubID(v string)       { m.TradingSessionSubID = &v }
func (m *Message) SetTradSesMethod(v int)                { m.TradSesMethod = &v }
func (m *Message) SetTradSesMode(v int)                  { m.TradSesMode = &v }
func (m *Message) SetUnsolicitedIndicator(v bool)        { m.UnsolicitedIndicator = &v }
func (m *Message) SetTradSesStatus(v int)                { m.TradSesStatus = v }
func (m *Message) SetTradSesStatusRejReason(v int)       { m.TradSesStatusRejReason = &v }
func (m *Message) SetTradSesStartTime(v time.Time)       { m.TradSesStartTime = &v }
func (m *Message) SetTradSesOpenTime(v time.Time)        { m.TradSesOpenTime = &v }
func (m *Message) SetTradSesPreCloseTime(v time.Time)    { m.TradSesPreCloseTime = &v }
func (m *Message) SetTradSesCloseTime(v time.Time)       { m.TradSesCloseTime = &v }
func (m *Message) SetTradSesEndTime(v time.Time)         { m.TradSesEndTime = &v }
func (m *Message) SetTotalVolumeTraded(v float64)        { m.TotalVolumeTraded = &v }
func (m *Message) SetText(v string)                      { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)               { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)               { m.EncodedText = &v }
func (m *Message) SetInstrument(v instrument.Instrument) { m.Instrument = &v }
func (m *Message) SetMarketID(v string)                  { m.MarketID = &v }
func (m *Message) SetMarketSegmentID(v string)           { m.MarketSegmentID = &v }
func (m *Message) SetTradSesEvent(v int)                 { m.TradSesEvent = &v }
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
	return enum.ApplVerID_FIX50SP1, "h", r
}
