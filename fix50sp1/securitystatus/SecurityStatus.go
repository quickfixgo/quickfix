//Package securitystatus msg type = f.
package securitystatus

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50sp1/applicationsequencecontrol"
	"github.com/quickfixgo/quickfix/fix50sp1/instrmtleggrp"
	"github.com/quickfixgo/quickfix/fix50sp1/instrument"
	"github.com/quickfixgo/quickfix/fix50sp1/instrumentextension"
	"github.com/quickfixgo/quickfix/fix50sp1/undinstrmtgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a SecurityStatus FIX Message
type Message struct {
	FIXMsgType string `fix:"f"`
	Header     fixt11.Header
	//SecurityStatusReqID is a non-required field for SecurityStatus.
	SecurityStatusReqID *string `fix:"324"`
	//Instrument Component
	Instrument instrument.Component
	//InstrumentExtension Component
	InstrumentExtension instrumentextension.Component
	//UndInstrmtGrp Component
	UndInstrmtGrp undinstrmtgrp.Component
	//InstrmtLegGrp Component
	InstrmtLegGrp instrmtleggrp.Component
	//Currency is a non-required field for SecurityStatus.
	Currency *string `fix:"15"`
	//TradingSessionID is a non-required field for SecurityStatus.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for SecurityStatus.
	TradingSessionSubID *string `fix:"625"`
	//UnsolicitedIndicator is a non-required field for SecurityStatus.
	UnsolicitedIndicator *bool `fix:"325"`
	//SecurityTradingStatus is a non-required field for SecurityStatus.
	SecurityTradingStatus *int `fix:"326"`
	//FinancialStatus is a non-required field for SecurityStatus.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for SecurityStatus.
	CorporateAction *string `fix:"292"`
	//HaltReasonChar is a non-required field for SecurityStatus.
	HaltReasonChar *string `fix:"327"`
	//InViewOfCommon is a non-required field for SecurityStatus.
	InViewOfCommon *bool `fix:"328"`
	//DueToRelated is a non-required field for SecurityStatus.
	DueToRelated *bool `fix:"329"`
	//BuyVolume is a non-required field for SecurityStatus.
	BuyVolume *float64 `fix:"330"`
	//SellVolume is a non-required field for SecurityStatus.
	SellVolume *float64 `fix:"331"`
	//HighPx is a non-required field for SecurityStatus.
	HighPx *float64 `fix:"332"`
	//LowPx is a non-required field for SecurityStatus.
	LowPx *float64 `fix:"333"`
	//LastPx is a non-required field for SecurityStatus.
	LastPx *float64 `fix:"31"`
	//TransactTime is a non-required field for SecurityStatus.
	TransactTime *time.Time `fix:"60"`
	//Adjustment is a non-required field for SecurityStatus.
	Adjustment *int `fix:"334"`
	//Text is a non-required field for SecurityStatus.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for SecurityStatus.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for SecurityStatus.
	EncodedText *string `fix:"355"`
	//FirstPx is a non-required field for SecurityStatus.
	FirstPx *float64 `fix:"1025"`
	//SecurityTradingEvent is a non-required field for SecurityStatus.
	SecurityTradingEvent *int `fix:"1174"`
	//MDBookType is a non-required field for SecurityStatus.
	MDBookType *int `fix:"1021"`
	//MarketDepth is a non-required field for SecurityStatus.
	MarketDepth *int `fix:"264"`
	//MarketID is a non-required field for SecurityStatus.
	MarketID *string `fix:"1301"`
	//MarketSegmentID is a non-required field for SecurityStatus.
	MarketSegmentID *string `fix:"1300"`
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
	return enum.ApplVerID_FIX50SP1, "f", r
}
