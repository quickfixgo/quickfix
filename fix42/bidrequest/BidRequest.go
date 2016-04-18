//Package bidrequest msg type = k.
package bidrequest

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoBidDescriptors is a repeating group in BidRequest
type NoBidDescriptors struct {
	//BidDescriptorType is a non-required field for NoBidDescriptors.
	BidDescriptorType *int `fix:"399"`
	//BidDescriptor is a non-required field for NoBidDescriptors.
	BidDescriptor *string `fix:"400"`
	//SideValueInd is a non-required field for NoBidDescriptors.
	SideValueInd *int `fix:"401"`
	//LiquidityValue is a non-required field for NoBidDescriptors.
	LiquidityValue *float64 `fix:"404"`
	//LiquidityNumSecurities is a non-required field for NoBidDescriptors.
	LiquidityNumSecurities *int `fix:"441"`
	//LiquidityPctLow is a non-required field for NoBidDescriptors.
	LiquidityPctLow *float64 `fix:"402"`
	//LiquidityPctHigh is a non-required field for NoBidDescriptors.
	LiquidityPctHigh *float64 `fix:"403"`
	//EFPTrackingError is a non-required field for NoBidDescriptors.
	EFPTrackingError *float64 `fix:"405"`
	//FairValue is a non-required field for NoBidDescriptors.
	FairValue *float64 `fix:"406"`
	//OutsideIndexPct is a non-required field for NoBidDescriptors.
	OutsideIndexPct *float64 `fix:"407"`
	//ValueOfFutures is a non-required field for NoBidDescriptors.
	ValueOfFutures *float64 `fix:"408"`
}

//NewNoBidDescriptors returns an initialized NoBidDescriptors instance
func NewNoBidDescriptors() *NoBidDescriptors {
	var m NoBidDescriptors
	return &m
}

func (m *NoBidDescriptors) SetBidDescriptorType(v int)      { m.BidDescriptorType = &v }
func (m *NoBidDescriptors) SetBidDescriptor(v string)       { m.BidDescriptor = &v }
func (m *NoBidDescriptors) SetSideValueInd(v int)           { m.SideValueInd = &v }
func (m *NoBidDescriptors) SetLiquidityValue(v float64)     { m.LiquidityValue = &v }
func (m *NoBidDescriptors) SetLiquidityNumSecurities(v int) { m.LiquidityNumSecurities = &v }
func (m *NoBidDescriptors) SetLiquidityPctLow(v float64)    { m.LiquidityPctLow = &v }
func (m *NoBidDescriptors) SetLiquidityPctHigh(v float64)   { m.LiquidityPctHigh = &v }
func (m *NoBidDescriptors) SetEFPTrackingError(v float64)   { m.EFPTrackingError = &v }
func (m *NoBidDescriptors) SetFairValue(v float64)          { m.FairValue = &v }
func (m *NoBidDescriptors) SetOutsideIndexPct(v float64)    { m.OutsideIndexPct = &v }
func (m *NoBidDescriptors) SetValueOfFutures(v float64)     { m.ValueOfFutures = &v }

//NoBidComponents is a repeating group in BidRequest
type NoBidComponents struct {
	//ListID is a non-required field for NoBidComponents.
	ListID *string `fix:"66"`
	//Side is a non-required field for NoBidComponents.
	Side *string `fix:"54"`
	//TradingSessionID is a non-required field for NoBidComponents.
	TradingSessionID *string `fix:"336"`
	//NetGrossInd is a non-required field for NoBidComponents.
	NetGrossInd *int `fix:"430"`
	//SettlmntTyp is a non-required field for NoBidComponents.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NoBidComponents.
	FutSettDate *string `fix:"64"`
	//Account is a non-required field for NoBidComponents.
	Account *string `fix:"1"`
}

//NewNoBidComponents returns an initialized NoBidComponents instance
func NewNoBidComponents() *NoBidComponents {
	var m NoBidComponents
	return &m
}

func (m *NoBidComponents) SetListID(v string)           { m.ListID = &v }
func (m *NoBidComponents) SetSide(v string)             { m.Side = &v }
func (m *NoBidComponents) SetTradingSessionID(v string) { m.TradingSessionID = &v }
func (m *NoBidComponents) SetNetGrossInd(v int)         { m.NetGrossInd = &v }
func (m *NoBidComponents) SetSettlmntTyp(v string)      { m.SettlmntTyp = &v }
func (m *NoBidComponents) SetFutSettDate(v string)      { m.FutSettDate = &v }
func (m *NoBidComponents) SetAccount(v string)          { m.Account = &v }

//Message is a BidRequest FIX Message
type Message struct {
	FIXMsgType string `fix:"k"`
	fix42.Header
	//BidID is a non-required field for BidRequest.
	BidID *string `fix:"390"`
	//ClientBidID is a required field for BidRequest.
	ClientBidID string `fix:"391"`
	//BidRequestTransType is a required field for BidRequest.
	BidRequestTransType string `fix:"374"`
	//ListName is a non-required field for BidRequest.
	ListName *string `fix:"392"`
	//TotalNumSecurities is a required field for BidRequest.
	TotalNumSecurities int `fix:"393"`
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
	//NoBidDescriptors is a non-required field for BidRequest.
	NoBidDescriptors []NoBidDescriptors `fix:"398,omitempty"`
	//NoBidComponents is a non-required field for BidRequest.
	NoBidComponents []NoBidComponents `fix:"420,omitempty"`
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
	//TradeType is a required field for BidRequest.
	TradeType string `fix:"418"`
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
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetBidID(v string)                        { m.BidID = &v }
func (m *Message) SetClientBidID(v string)                  { m.ClientBidID = v }
func (m *Message) SetBidRequestTransType(v string)          { m.BidRequestTransType = v }
func (m *Message) SetListName(v string)                     { m.ListName = &v }
func (m *Message) SetTotalNumSecurities(v int)              { m.TotalNumSecurities = v }
func (m *Message) SetBidType(v int)                         { m.BidType = v }
func (m *Message) SetNumTickets(v int)                      { m.NumTickets = &v }
func (m *Message) SetCurrency(v string)                     { m.Currency = &v }
func (m *Message) SetSideValue1(v float64)                  { m.SideValue1 = &v }
func (m *Message) SetSideValue2(v float64)                  { m.SideValue2 = &v }
func (m *Message) SetNoBidDescriptors(v []NoBidDescriptors) { m.NoBidDescriptors = v }
func (m *Message) SetNoBidComponents(v []NoBidComponents)   { m.NoBidComponents = v }
func (m *Message) SetLiquidityIndType(v int)                { m.LiquidityIndType = &v }
func (m *Message) SetWtAverageLiquidity(v float64)          { m.WtAverageLiquidity = &v }
func (m *Message) SetExchangeForPhysical(v bool)            { m.ExchangeForPhysical = &v }
func (m *Message) SetOutMainCntryUIndex(v float64)          { m.OutMainCntryUIndex = &v }
func (m *Message) SetCrossPercent(v float64)                { m.CrossPercent = &v }
func (m *Message) SetProgRptReqs(v int)                     { m.ProgRptReqs = &v }
func (m *Message) SetProgPeriodInterval(v int)              { m.ProgPeriodInterval = &v }
func (m *Message) SetIncTaxInd(v int)                       { m.IncTaxInd = &v }
func (m *Message) SetForexReq(v bool)                       { m.ForexReq = &v }
func (m *Message) SetNumBidders(v int)                      { m.NumBidders = &v }
func (m *Message) SetTradeDate(v string)                    { m.TradeDate = &v }
func (m *Message) SetTradeType(v string)                    { m.TradeType = v }
func (m *Message) SetBasisPxType(v string)                  { m.BasisPxType = v }
func (m *Message) SetStrikeTime(v time.Time)                { m.StrikeTime = &v }
func (m *Message) SetText(v string)                         { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                  { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                  { m.EncodedText = &v }

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
	return enum.BeginStringFIX42, "k", r
}
