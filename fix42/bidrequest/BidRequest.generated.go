package bidrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//BidRequest is the fix42 BidRequest type, MsgType = k
type BidRequest struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a BidRequest from a quickfix.Message instance
func FromMessage(m *quickfix.Message) BidRequest {
	return BidRequest{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m BidRequest) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a BidRequest initialized with the required fields for BidRequest
func New(clientbidid field.ClientBidIDField, bidrequesttranstype field.BidRequestTransTypeField, totalnumsecurities field.TotalNumSecuritiesField, bidtype field.BidTypeField, tradetype field.TradeTypeField, basispxtype field.BasisPxTypeField) (m BidRequest) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("k"))
	m.Set(clientbidid)
	m.Set(bidrequesttranstype)
	m.Set(totalnumsecurities)
	m.Set(bidtype)
	m.Set(tradetype)
	m.Set(basispxtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg BidRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "k", r
}

//SetCurrency sets Currency, Tag 15
func (m BidRequest) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetText sets Text, Tag 58
func (m BidRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m BidRequest) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetForexReq sets ForexReq, Tag 121
func (m BidRequest) SetForexReq(v bool) {
	m.Set(field.NewForexReq(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m BidRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m BidRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetBidRequestTransType sets BidRequestTransType, Tag 374
func (m BidRequest) SetBidRequestTransType(v enum.BidRequestTransType) {
	m.Set(field.NewBidRequestTransType(v))
}

//SetBidID sets BidID, Tag 390
func (m BidRequest) SetBidID(v string) {
	m.Set(field.NewBidID(v))
}

//SetClientBidID sets ClientBidID, Tag 391
func (m BidRequest) SetClientBidID(v string) {
	m.Set(field.NewClientBidID(v))
}

//SetListName sets ListName, Tag 392
func (m BidRequest) SetListName(v string) {
	m.Set(field.NewListName(v))
}

//SetTotalNumSecurities sets TotalNumSecurities, Tag 393
func (m BidRequest) SetTotalNumSecurities(v int) {
	m.Set(field.NewTotalNumSecurities(v))
}

//SetBidType sets BidType, Tag 394
func (m BidRequest) SetBidType(v enum.BidType) {
	m.Set(field.NewBidType(v))
}

//SetNumTickets sets NumTickets, Tag 395
func (m BidRequest) SetNumTickets(v int) {
	m.Set(field.NewNumTickets(v))
}

//SetSideValue1 sets SideValue1, Tag 396
func (m BidRequest) SetSideValue1(value decimal.Decimal, scale int32) {
	m.Set(field.NewSideValue1(value, scale))
}

//SetSideValue2 sets SideValue2, Tag 397
func (m BidRequest) SetSideValue2(value decimal.Decimal, scale int32) {
	m.Set(field.NewSideValue2(value, scale))
}

//SetNoBidDescriptors sets NoBidDescriptors, Tag 398
func (m BidRequest) SetNoBidDescriptors(f NoBidDescriptorsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLiquidityIndType sets LiquidityIndType, Tag 409
func (m BidRequest) SetLiquidityIndType(v enum.LiquidityIndType) {
	m.Set(field.NewLiquidityIndType(v))
}

//SetWtAverageLiquidity sets WtAverageLiquidity, Tag 410
func (m BidRequest) SetWtAverageLiquidity(value decimal.Decimal, scale int32) {
	m.Set(field.NewWtAverageLiquidity(value, scale))
}

//SetExchangeForPhysical sets ExchangeForPhysical, Tag 411
func (m BidRequest) SetExchangeForPhysical(v bool) {
	m.Set(field.NewExchangeForPhysical(v))
}

//SetOutMainCntryUIndex sets OutMainCntryUIndex, Tag 412
func (m BidRequest) SetOutMainCntryUIndex(value decimal.Decimal, scale int32) {
	m.Set(field.NewOutMainCntryUIndex(value, scale))
}

//SetCrossPercent sets CrossPercent, Tag 413
func (m BidRequest) SetCrossPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewCrossPercent(value, scale))
}

//SetProgRptReqs sets ProgRptReqs, Tag 414
func (m BidRequest) SetProgRptReqs(v enum.ProgRptReqs) {
	m.Set(field.NewProgRptReqs(v))
}

//SetProgPeriodInterval sets ProgPeriodInterval, Tag 415
func (m BidRequest) SetProgPeriodInterval(v int) {
	m.Set(field.NewProgPeriodInterval(v))
}

//SetIncTaxInd sets IncTaxInd, Tag 416
func (m BidRequest) SetIncTaxInd(v enum.IncTaxInd) {
	m.Set(field.NewIncTaxInd(v))
}

//SetNumBidders sets NumBidders, Tag 417
func (m BidRequest) SetNumBidders(v int) {
	m.Set(field.NewNumBidders(v))
}

//SetTradeType sets TradeType, Tag 418
func (m BidRequest) SetTradeType(v enum.TradeType) {
	m.Set(field.NewTradeType(v))
}

//SetBasisPxType sets BasisPxType, Tag 419
func (m BidRequest) SetBasisPxType(v enum.BasisPxType) {
	m.Set(field.NewBasisPxType(v))
}

//SetNoBidComponents sets NoBidComponents, Tag 420
func (m BidRequest) SetNoBidComponents(f NoBidComponentsRepeatingGroup) {
	m.SetGroup(f)
}

//SetStrikeTime sets StrikeTime, Tag 443
func (m BidRequest) SetStrikeTime(v time.Time) {
	m.Set(field.NewStrikeTime(v))
}

//GetCurrency gets Currency, Tag 15
func (m BidRequest) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m BidRequest) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m BidRequest) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m BidRequest) GetForexReq() (v bool, err quickfix.MessageRejectError) {
	var f field.ForexReqField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m BidRequest) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m BidRequest) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidRequestTransType gets BidRequestTransType, Tag 374
func (m BidRequest) GetBidRequestTransType() (v enum.BidRequestTransType, err quickfix.MessageRejectError) {
	var f field.BidRequestTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidID gets BidID, Tag 390
func (m BidRequest) GetBidID() (v string, err quickfix.MessageRejectError) {
	var f field.BidIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClientBidID gets ClientBidID, Tag 391
func (m BidRequest) GetClientBidID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientBidIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListName gets ListName, Tag 392
func (m BidRequest) GetListName() (v string, err quickfix.MessageRejectError) {
	var f field.ListNameField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalNumSecurities gets TotalNumSecurities, Tag 393
func (m BidRequest) GetTotalNumSecurities() (v int, err quickfix.MessageRejectError) {
	var f field.TotalNumSecuritiesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidType gets BidType, Tag 394
func (m BidRequest) GetBidType() (v enum.BidType, err quickfix.MessageRejectError) {
	var f field.BidTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNumTickets gets NumTickets, Tag 395
func (m BidRequest) GetNumTickets() (v int, err quickfix.MessageRejectError) {
	var f field.NumTicketsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSideValue1 gets SideValue1, Tag 396
func (m BidRequest) GetSideValue1() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SideValue1Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSideValue2 gets SideValue2, Tag 397
func (m BidRequest) GetSideValue2() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SideValue2Field
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoBidDescriptors gets NoBidDescriptors, Tag 398
func (m BidRequest) GetNoBidDescriptors() (NoBidDescriptorsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoBidDescriptorsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLiquidityIndType gets LiquidityIndType, Tag 409
func (m BidRequest) GetLiquidityIndType() (v enum.LiquidityIndType, err quickfix.MessageRejectError) {
	var f field.LiquidityIndTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetWtAverageLiquidity gets WtAverageLiquidity, Tag 410
func (m BidRequest) GetWtAverageLiquidity() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.WtAverageLiquidityField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExchangeForPhysical gets ExchangeForPhysical, Tag 411
func (m BidRequest) GetExchangeForPhysical() (v bool, err quickfix.MessageRejectError) {
	var f field.ExchangeForPhysicalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOutMainCntryUIndex gets OutMainCntryUIndex, Tag 412
func (m BidRequest) GetOutMainCntryUIndex() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OutMainCntryUIndexField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCrossPercent gets CrossPercent, Tag 413
func (m BidRequest) GetCrossPercent() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CrossPercentField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProgRptReqs gets ProgRptReqs, Tag 414
func (m BidRequest) GetProgRptReqs() (v enum.ProgRptReqs, err quickfix.MessageRejectError) {
	var f field.ProgRptReqsField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProgPeriodInterval gets ProgPeriodInterval, Tag 415
func (m BidRequest) GetProgPeriodInterval() (v int, err quickfix.MessageRejectError) {
	var f field.ProgPeriodIntervalField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIncTaxInd gets IncTaxInd, Tag 416
func (m BidRequest) GetIncTaxInd() (v enum.IncTaxInd, err quickfix.MessageRejectError) {
	var f field.IncTaxIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNumBidders gets NumBidders, Tag 417
func (m BidRequest) GetNumBidders() (v int, err quickfix.MessageRejectError) {
	var f field.NumBiddersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeType gets TradeType, Tag 418
func (m BidRequest) GetTradeType() (v enum.TradeType, err quickfix.MessageRejectError) {
	var f field.TradeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBasisPxType gets BasisPxType, Tag 419
func (m BidRequest) GetBasisPxType() (v enum.BasisPxType, err quickfix.MessageRejectError) {
	var f field.BasisPxTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoBidComponents gets NoBidComponents, Tag 420
func (m BidRequest) GetNoBidComponents() (NoBidComponentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoBidComponentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetStrikeTime gets StrikeTime, Tag 443
func (m BidRequest) GetStrikeTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.StrikeTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m BidRequest) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasText returns true if Text is present, Tag 58
func (m BidRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m BidRequest) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasForexReq returns true if ForexReq is present, Tag 121
func (m BidRequest) HasForexReq() bool {
	return m.Has(tag.ForexReq)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m BidRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m BidRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasBidRequestTransType returns true if BidRequestTransType is present, Tag 374
func (m BidRequest) HasBidRequestTransType() bool {
	return m.Has(tag.BidRequestTransType)
}

//HasBidID returns true if BidID is present, Tag 390
func (m BidRequest) HasBidID() bool {
	return m.Has(tag.BidID)
}

//HasClientBidID returns true if ClientBidID is present, Tag 391
func (m BidRequest) HasClientBidID() bool {
	return m.Has(tag.ClientBidID)
}

//HasListName returns true if ListName is present, Tag 392
func (m BidRequest) HasListName() bool {
	return m.Has(tag.ListName)
}

//HasTotalNumSecurities returns true if TotalNumSecurities is present, Tag 393
func (m BidRequest) HasTotalNumSecurities() bool {
	return m.Has(tag.TotalNumSecurities)
}

//HasBidType returns true if BidType is present, Tag 394
func (m BidRequest) HasBidType() bool {
	return m.Has(tag.BidType)
}

//HasNumTickets returns true if NumTickets is present, Tag 395
func (m BidRequest) HasNumTickets() bool {
	return m.Has(tag.NumTickets)
}

//HasSideValue1 returns true if SideValue1 is present, Tag 396
func (m BidRequest) HasSideValue1() bool {
	return m.Has(tag.SideValue1)
}

//HasSideValue2 returns true if SideValue2 is present, Tag 397
func (m BidRequest) HasSideValue2() bool {
	return m.Has(tag.SideValue2)
}

//HasNoBidDescriptors returns true if NoBidDescriptors is present, Tag 398
func (m BidRequest) HasNoBidDescriptors() bool {
	return m.Has(tag.NoBidDescriptors)
}

//HasLiquidityIndType returns true if LiquidityIndType is present, Tag 409
func (m BidRequest) HasLiquidityIndType() bool {
	return m.Has(tag.LiquidityIndType)
}

//HasWtAverageLiquidity returns true if WtAverageLiquidity is present, Tag 410
func (m BidRequest) HasWtAverageLiquidity() bool {
	return m.Has(tag.WtAverageLiquidity)
}

//HasExchangeForPhysical returns true if ExchangeForPhysical is present, Tag 411
func (m BidRequest) HasExchangeForPhysical() bool {
	return m.Has(tag.ExchangeForPhysical)
}

//HasOutMainCntryUIndex returns true if OutMainCntryUIndex is present, Tag 412
func (m BidRequest) HasOutMainCntryUIndex() bool {
	return m.Has(tag.OutMainCntryUIndex)
}

//HasCrossPercent returns true if CrossPercent is present, Tag 413
func (m BidRequest) HasCrossPercent() bool {
	return m.Has(tag.CrossPercent)
}

//HasProgRptReqs returns true if ProgRptReqs is present, Tag 414
func (m BidRequest) HasProgRptReqs() bool {
	return m.Has(tag.ProgRptReqs)
}

//HasProgPeriodInterval returns true if ProgPeriodInterval is present, Tag 415
func (m BidRequest) HasProgPeriodInterval() bool {
	return m.Has(tag.ProgPeriodInterval)
}

//HasIncTaxInd returns true if IncTaxInd is present, Tag 416
func (m BidRequest) HasIncTaxInd() bool {
	return m.Has(tag.IncTaxInd)
}

//HasNumBidders returns true if NumBidders is present, Tag 417
func (m BidRequest) HasNumBidders() bool {
	return m.Has(tag.NumBidders)
}

//HasTradeType returns true if TradeType is present, Tag 418
func (m BidRequest) HasTradeType() bool {
	return m.Has(tag.TradeType)
}

//HasBasisPxType returns true if BasisPxType is present, Tag 419
func (m BidRequest) HasBasisPxType() bool {
	return m.Has(tag.BasisPxType)
}

//HasNoBidComponents returns true if NoBidComponents is present, Tag 420
func (m BidRequest) HasNoBidComponents() bool {
	return m.Has(tag.NoBidComponents)
}

//HasStrikeTime returns true if StrikeTime is present, Tag 443
func (m BidRequest) HasStrikeTime() bool {
	return m.Has(tag.StrikeTime)
}

//NoBidDescriptors is a repeating group element, Tag 398
type NoBidDescriptors struct {
	*quickfix.Group
}

//SetBidDescriptorType sets BidDescriptorType, Tag 399
func (m NoBidDescriptors) SetBidDescriptorType(v enum.BidDescriptorType) {
	m.Set(field.NewBidDescriptorType(v))
}

//SetBidDescriptor sets BidDescriptor, Tag 400
func (m NoBidDescriptors) SetBidDescriptor(v string) {
	m.Set(field.NewBidDescriptor(v))
}

//SetSideValueInd sets SideValueInd, Tag 401
func (m NoBidDescriptors) SetSideValueInd(v enum.SideValueInd) {
	m.Set(field.NewSideValueInd(v))
}

//SetLiquidityValue sets LiquidityValue, Tag 404
func (m NoBidDescriptors) SetLiquidityValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewLiquidityValue(value, scale))
}

//SetLiquidityNumSecurities sets LiquidityNumSecurities, Tag 441
func (m NoBidDescriptors) SetLiquidityNumSecurities(v int) {
	m.Set(field.NewLiquidityNumSecurities(v))
}

//SetLiquidityPctLow sets LiquidityPctLow, Tag 402
func (m NoBidDescriptors) SetLiquidityPctLow(value decimal.Decimal, scale int32) {
	m.Set(field.NewLiquidityPctLow(value, scale))
}

//SetLiquidityPctHigh sets LiquidityPctHigh, Tag 403
func (m NoBidDescriptors) SetLiquidityPctHigh(value decimal.Decimal, scale int32) {
	m.Set(field.NewLiquidityPctHigh(value, scale))
}

//SetEFPTrackingError sets EFPTrackingError, Tag 405
func (m NoBidDescriptors) SetEFPTrackingError(value decimal.Decimal, scale int32) {
	m.Set(field.NewEFPTrackingError(value, scale))
}

//SetFairValue sets FairValue, Tag 406
func (m NoBidDescriptors) SetFairValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewFairValue(value, scale))
}

//SetOutsideIndexPct sets OutsideIndexPct, Tag 407
func (m NoBidDescriptors) SetOutsideIndexPct(value decimal.Decimal, scale int32) {
	m.Set(field.NewOutsideIndexPct(value, scale))
}

//SetValueOfFutures sets ValueOfFutures, Tag 408
func (m NoBidDescriptors) SetValueOfFutures(value decimal.Decimal, scale int32) {
	m.Set(field.NewValueOfFutures(value, scale))
}

//GetBidDescriptorType gets BidDescriptorType, Tag 399
func (m NoBidDescriptors) GetBidDescriptorType() (v enum.BidDescriptorType, err quickfix.MessageRejectError) {
	var f field.BidDescriptorTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBidDescriptor gets BidDescriptor, Tag 400
func (m NoBidDescriptors) GetBidDescriptor() (v string, err quickfix.MessageRejectError) {
	var f field.BidDescriptorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSideValueInd gets SideValueInd, Tag 401
func (m NoBidDescriptors) GetSideValueInd() (v enum.SideValueInd, err quickfix.MessageRejectError) {
	var f field.SideValueIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLiquidityValue gets LiquidityValue, Tag 404
func (m NoBidDescriptors) GetLiquidityValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LiquidityValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLiquidityNumSecurities gets LiquidityNumSecurities, Tag 441
func (m NoBidDescriptors) GetLiquidityNumSecurities() (v int, err quickfix.MessageRejectError) {
	var f field.LiquidityNumSecuritiesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLiquidityPctLow gets LiquidityPctLow, Tag 402
func (m NoBidDescriptors) GetLiquidityPctLow() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LiquidityPctLowField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLiquidityPctHigh gets LiquidityPctHigh, Tag 403
func (m NoBidDescriptors) GetLiquidityPctHigh() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LiquidityPctHighField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEFPTrackingError gets EFPTrackingError, Tag 405
func (m NoBidDescriptors) GetEFPTrackingError() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EFPTrackingErrorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFairValue gets FairValue, Tag 406
func (m NoBidDescriptors) GetFairValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FairValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOutsideIndexPct gets OutsideIndexPct, Tag 407
func (m NoBidDescriptors) GetOutsideIndexPct() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.OutsideIndexPctField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetValueOfFutures gets ValueOfFutures, Tag 408
func (m NoBidDescriptors) GetValueOfFutures() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ValueOfFuturesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasBidDescriptorType returns true if BidDescriptorType is present, Tag 399
func (m NoBidDescriptors) HasBidDescriptorType() bool {
	return m.Has(tag.BidDescriptorType)
}

//HasBidDescriptor returns true if BidDescriptor is present, Tag 400
func (m NoBidDescriptors) HasBidDescriptor() bool {
	return m.Has(tag.BidDescriptor)
}

//HasSideValueInd returns true if SideValueInd is present, Tag 401
func (m NoBidDescriptors) HasSideValueInd() bool {
	return m.Has(tag.SideValueInd)
}

//HasLiquidityValue returns true if LiquidityValue is present, Tag 404
func (m NoBidDescriptors) HasLiquidityValue() bool {
	return m.Has(tag.LiquidityValue)
}

//HasLiquidityNumSecurities returns true if LiquidityNumSecurities is present, Tag 441
func (m NoBidDescriptors) HasLiquidityNumSecurities() bool {
	return m.Has(tag.LiquidityNumSecurities)
}

//HasLiquidityPctLow returns true if LiquidityPctLow is present, Tag 402
func (m NoBidDescriptors) HasLiquidityPctLow() bool {
	return m.Has(tag.LiquidityPctLow)
}

//HasLiquidityPctHigh returns true if LiquidityPctHigh is present, Tag 403
func (m NoBidDescriptors) HasLiquidityPctHigh() bool {
	return m.Has(tag.LiquidityPctHigh)
}

//HasEFPTrackingError returns true if EFPTrackingError is present, Tag 405
func (m NoBidDescriptors) HasEFPTrackingError() bool {
	return m.Has(tag.EFPTrackingError)
}

//HasFairValue returns true if FairValue is present, Tag 406
func (m NoBidDescriptors) HasFairValue() bool {
	return m.Has(tag.FairValue)
}

//HasOutsideIndexPct returns true if OutsideIndexPct is present, Tag 407
func (m NoBidDescriptors) HasOutsideIndexPct() bool {
	return m.Has(tag.OutsideIndexPct)
}

//HasValueOfFutures returns true if ValueOfFutures is present, Tag 408
func (m NoBidDescriptors) HasValueOfFutures() bool {
	return m.Has(tag.ValueOfFutures)
}

//NoBidDescriptorsRepeatingGroup is a repeating group, Tag 398
type NoBidDescriptorsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoBidDescriptorsRepeatingGroup returns an initialized, NoBidDescriptorsRepeatingGroup
func NewNoBidDescriptorsRepeatingGroup() NoBidDescriptorsRepeatingGroup {
	return NoBidDescriptorsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoBidDescriptors,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.BidDescriptorType), quickfix.GroupElement(tag.BidDescriptor), quickfix.GroupElement(tag.SideValueInd), quickfix.GroupElement(tag.LiquidityValue), quickfix.GroupElement(tag.LiquidityNumSecurities), quickfix.GroupElement(tag.LiquidityPctLow), quickfix.GroupElement(tag.LiquidityPctHigh), quickfix.GroupElement(tag.EFPTrackingError), quickfix.GroupElement(tag.FairValue), quickfix.GroupElement(tag.OutsideIndexPct), quickfix.GroupElement(tag.ValueOfFutures)})}
}

//Add create and append a new NoBidDescriptors to this group
func (m NoBidDescriptorsRepeatingGroup) Add() NoBidDescriptors {
	g := m.RepeatingGroup.Add()
	return NoBidDescriptors{g}
}

//Get returns the ith NoBidDescriptors in the NoBidDescriptorsRepeatinGroup
func (m NoBidDescriptorsRepeatingGroup) Get(i int) NoBidDescriptors {
	return NoBidDescriptors{m.RepeatingGroup.Get(i)}
}

//NoBidComponents is a repeating group element, Tag 420
type NoBidComponents struct {
	*quickfix.Group
}

//SetListID sets ListID, Tag 66
func (m NoBidComponents) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetSide sets Side, Tag 54
func (m NoBidComponents) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoBidComponents) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetNetGrossInd sets NetGrossInd, Tag 430
func (m NoBidComponents) SetNetGrossInd(v enum.NetGrossInd) {
	m.Set(field.NewNetGrossInd(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m NoBidComponents) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m NoBidComponents) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetAccount sets Account, Tag 1
func (m NoBidComponents) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//GetListID gets ListID, Tag 66
func (m NoBidComponents) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m NoBidComponents) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoBidComponents) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetGrossInd gets NetGrossInd, Tag 430
func (m NoBidComponents) GetNetGrossInd() (v enum.NetGrossInd, err quickfix.MessageRejectError) {
	var f field.NetGrossIndField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m NoBidComponents) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m NoBidComponents) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAccount gets Account, Tag 1
func (m NoBidComponents) GetAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasListID returns true if ListID is present, Tag 66
func (m NoBidComponents) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasSide returns true if Side is present, Tag 54
func (m NoBidComponents) HasSide() bool {
	return m.Has(tag.Side)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoBidComponents) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasNetGrossInd returns true if NetGrossInd is present, Tag 430
func (m NoBidComponents) HasNetGrossInd() bool {
	return m.Has(tag.NetGrossInd)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m NoBidComponents) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m NoBidComponents) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasAccount returns true if Account is present, Tag 1
func (m NoBidComponents) HasAccount() bool {
	return m.Has(tag.Account)
}

//NoBidComponentsRepeatingGroup is a repeating group, Tag 420
type NoBidComponentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoBidComponentsRepeatingGroup returns an initialized, NoBidComponentsRepeatingGroup
func NewNoBidComponentsRepeatingGroup() NoBidComponentsRepeatingGroup {
	return NoBidComponentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoBidComponents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.NetGrossInd), quickfix.GroupElement(tag.SettlmntTyp), quickfix.GroupElement(tag.FutSettDate), quickfix.GroupElement(tag.Account)})}
}

//Add create and append a new NoBidComponents to this group
func (m NoBidComponentsRepeatingGroup) Add() NoBidComponents {
	g := m.RepeatingGroup.Add()
	return NoBidComponents{g}
}

//Get returns the ith NoBidComponents in the NoBidComponentsRepeatinGroup
func (m NoBidComponentsRepeatingGroup) Get(i int) NoBidComponents {
	return NoBidComponents{m.RepeatingGroup.Get(i)}
}
