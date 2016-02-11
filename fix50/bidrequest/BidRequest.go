//Package bidrequest msg type = k.
package bidrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix50/bidcompreqgrp"
	"github.com/quickfixgo/quickfix/fix50/biddescreqgrp"
	"github.com/quickfixgo/quickfix/fixt11"
	"time"
)

//Message is a BidRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"k"`
	Header     fixt11.Header
	//BidID is a non-required field for BidRequest.
	BidID *string `fix:"390"`
	//ClientBidID is a required field for BidRequest.
	ClientBidID string `fix:"391"`
	//BidRequestTransType is a required field for BidRequest.
	BidRequestTransType string `fix:"374"`
	//ListName is a non-required field for BidRequest.
	ListName *string `fix:"392"`
	//TotNoRelatedSym is a required field for BidRequest.
	TotNoRelatedSym int `fix:"393"`
	//BidType is a required field for BidRequest.
	BidType int `fix:"394"`
	//NumTickets is a non-required field for BidRequest.
	NumTickets *int `fix:"395"`
	//Currency is a non-required field for BidRequest.
	Currency *string `fix:"15"`
	//SideValue1 is a non-required field for BidRequest.
	SideValue1 *float64 `fix:"396"`
	//SideValue2 is a non-required field for BidRequest.
	SideValue2 *float64 `fix:"397"`
	//BidDescReqGrp Component
	BidDescReqGrp biddescreqgrp.Component
	//BidCompReqGrp Component
	BidCompReqGrp bidcompreqgrp.Component
	//LiquidityIndType is a non-required field for BidRequest.
	LiquidityIndType *int `fix:"409"`
	//WtAverageLiquidity is a non-required field for BidRequest.
	WtAverageLiquidity *float64 `fix:"410"`
	//ExchangeForPhysical is a non-required field for BidRequest.
	ExchangeForPhysical *bool `fix:"411"`
	//OutMainCntryUIndex is a non-required field for BidRequest.
	OutMainCntryUIndex *float64 `fix:"412"`
	//CrossPercent is a non-required field for BidRequest.
	CrossPercent *float64 `fix:"413"`
	//ProgRptReqs is a non-required field for BidRequest.
	ProgRptReqs *int `fix:"414"`
	//ProgPeriodInterval is a non-required field for BidRequest.
	ProgPeriodInterval *int `fix:"415"`
	//IncTaxInd is a non-required field for BidRequest.
	IncTaxInd *int `fix:"416"`
	//ForexReq is a non-required field for BidRequest.
	ForexReq *bool `fix:"121"`
	//NumBidders is a non-required field for BidRequest.
	NumBidders *int `fix:"417"`
	//TradeDate is a non-required field for BidRequest.
	TradeDate *string `fix:"75"`
	//BidTradeType is a required field for BidRequest.
	BidTradeType string `fix:"418"`
	//BasisPxType is a required field for BidRequest.
	BasisPxType string `fix:"419"`
	//StrikeTime is a non-required field for BidRequest.
	StrikeTime *time.Time `fix:"443"`
	//Text is a non-required field for BidRequest.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for BidRequest.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for BidRequest.
	EncodedText *string `fix:"355"`
	Trailer     fixt11.Trailer
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
	return enum.BeginStringFIX50, "k", r
}
