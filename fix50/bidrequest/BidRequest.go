package bidrequest

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//BidRequest is the fix50 BidRequest type, MsgType = k
type BidRequest struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a BidRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) BidRequest {
	return BidRequest{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m BidRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a BidRequest initialized with the required fields for BidRequest
func New(clientbidid field.ClientBidIDField, bidrequesttranstype field.BidRequestTransTypeField, totnorelatedsym field.TotNoRelatedSymField, bidtype field.BidTypeField, bidtradetype field.BidTradeTypeField, basispxtype field.BasisPxTypeField) (m BidRequest) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("k"))
	m.Set(clientbidid)
	m.Set(bidrequesttranstype)
	m.Set(totnorelatedsym)
	m.Set(bidtype)
	m.Set(bidtradetype)
	m.Set(basispxtype)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg BidRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "k", r
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
func (m BidRequest) SetBidRequestTransType(v string) {
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

//SetTotNoRelatedSym sets TotNoRelatedSym, Tag 393
func (m BidRequest) SetTotNoRelatedSym(v int) {
	m.Set(field.NewTotNoRelatedSym(v))
}

//SetBidType sets BidType, Tag 394
func (m BidRequest) SetBidType(v int) {
	m.Set(field.NewBidType(v))
}

//SetNumTickets sets NumTickets, Tag 395
func (m BidRequest) SetNumTickets(v int) {
	m.Set(field.NewNumTickets(v))
}

//SetSideValue1 sets SideValue1, Tag 396
func (m BidRequest) SetSideValue1(v float64) {
	m.Set(field.NewSideValue1(v))
}

//SetSideValue2 sets SideValue2, Tag 397
func (m BidRequest) SetSideValue2(v float64) {
	m.Set(field.NewSideValue2(v))
}

//SetNoBidDescriptors sets NoBidDescriptors, Tag 398
func (m BidRequest) SetNoBidDescriptors(f NoBidDescriptorsRepeatingGroup) {
	m.SetGroup(f)
}

//SetLiquidityIndType sets LiquidityIndType, Tag 409
func (m BidRequest) SetLiquidityIndType(v int) {
	m.Set(field.NewLiquidityIndType(v))
}

//SetWtAverageLiquidity sets WtAverageLiquidity, Tag 410
func (m BidRequest) SetWtAverageLiquidity(v float64) {
	m.Set(field.NewWtAverageLiquidity(v))
}

//SetExchangeForPhysical sets ExchangeForPhysical, Tag 411
func (m BidRequest) SetExchangeForPhysical(v bool) {
	m.Set(field.NewExchangeForPhysical(v))
}

//SetOutMainCntryUIndex sets OutMainCntryUIndex, Tag 412
func (m BidRequest) SetOutMainCntryUIndex(v float64) {
	m.Set(field.NewOutMainCntryUIndex(v))
}

//SetCrossPercent sets CrossPercent, Tag 413
func (m BidRequest) SetCrossPercent(v float64) {
	m.Set(field.NewCrossPercent(v))
}

//SetProgRptReqs sets ProgRptReqs, Tag 414
func (m BidRequest) SetProgRptReqs(v int) {
	m.Set(field.NewProgRptReqs(v))
}

//SetProgPeriodInterval sets ProgPeriodInterval, Tag 415
func (m BidRequest) SetProgPeriodInterval(v int) {
	m.Set(field.NewProgPeriodInterval(v))
}

//SetIncTaxInd sets IncTaxInd, Tag 416
func (m BidRequest) SetIncTaxInd(v int) {
	m.Set(field.NewIncTaxInd(v))
}

//SetNumBidders sets NumBidders, Tag 417
func (m BidRequest) SetNumBidders(v int) {
	m.Set(field.NewNumBidders(v))
}

//SetBidTradeType sets BidTradeType, Tag 418
func (m BidRequest) SetBidTradeType(v string) {
	m.Set(field.NewBidTradeType(v))
}

//SetBasisPxType sets BasisPxType, Tag 419
func (m BidRequest) SetBasisPxType(v string) {
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
func (m BidRequest) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m BidRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m BidRequest) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetForexReq gets ForexReq, Tag 121
func (m BidRequest) GetForexReq() (f field.ForexReqField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m BidRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m BidRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidRequestTransType gets BidRequestTransType, Tag 374
func (m BidRequest) GetBidRequestTransType() (f field.BidRequestTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidID gets BidID, Tag 390
func (m BidRequest) GetBidID() (f field.BidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientBidID gets ClientBidID, Tag 391
func (m BidRequest) GetClientBidID() (f field.ClientBidIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListName gets ListName, Tag 392
func (m BidRequest) GetListName() (f field.ListNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotNoRelatedSym gets TotNoRelatedSym, Tag 393
func (m BidRequest) GetTotNoRelatedSym() (f field.TotNoRelatedSymField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidType gets BidType, Tag 394
func (m BidRequest) GetBidType() (f field.BidTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNumTickets gets NumTickets, Tag 395
func (m BidRequest) GetNumTickets() (f field.NumTicketsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideValue1 gets SideValue1, Tag 396
func (m BidRequest) GetSideValue1() (f field.SideValue1Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideValue2 gets SideValue2, Tag 397
func (m BidRequest) GetSideValue2() (f field.SideValue2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoBidDescriptors gets NoBidDescriptors, Tag 398
func (m BidRequest) GetNoBidDescriptors() (NoBidDescriptorsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoBidDescriptorsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetLiquidityIndType gets LiquidityIndType, Tag 409
func (m BidRequest) GetLiquidityIndType() (f field.LiquidityIndTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWtAverageLiquidity gets WtAverageLiquidity, Tag 410
func (m BidRequest) GetWtAverageLiquidity() (f field.WtAverageLiquidityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExchangeForPhysical gets ExchangeForPhysical, Tag 411
func (m BidRequest) GetExchangeForPhysical() (f field.ExchangeForPhysicalField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOutMainCntryUIndex gets OutMainCntryUIndex, Tag 412
func (m BidRequest) GetOutMainCntryUIndex() (f field.OutMainCntryUIndexField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCrossPercent gets CrossPercent, Tag 413
func (m BidRequest) GetCrossPercent() (f field.CrossPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProgRptReqs gets ProgRptReqs, Tag 414
func (m BidRequest) GetProgRptReqs() (f field.ProgRptReqsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProgPeriodInterval gets ProgPeriodInterval, Tag 415
func (m BidRequest) GetProgPeriodInterval() (f field.ProgPeriodIntervalField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIncTaxInd gets IncTaxInd, Tag 416
func (m BidRequest) GetIncTaxInd() (f field.IncTaxIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNumBidders gets NumBidders, Tag 417
func (m BidRequest) GetNumBidders() (f field.NumBiddersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidTradeType gets BidTradeType, Tag 418
func (m BidRequest) GetBidTradeType() (f field.BidTradeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBasisPxType gets BasisPxType, Tag 419
func (m BidRequest) GetBasisPxType() (f field.BasisPxTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoBidComponents gets NoBidComponents, Tag 420
func (m BidRequest) GetNoBidComponents() (NoBidComponentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoBidComponentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetStrikeTime gets StrikeTime, Tag 443
func (m BidRequest) GetStrikeTime() (f field.StrikeTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//HasTotNoRelatedSym returns true if TotNoRelatedSym is present, Tag 393
func (m BidRequest) HasTotNoRelatedSym() bool {
	return m.Has(tag.TotNoRelatedSym)
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

//HasBidTradeType returns true if BidTradeType is present, Tag 418
func (m BidRequest) HasBidTradeType() bool {
	return m.Has(tag.BidTradeType)
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
	quickfix.Group
}

//SetBidDescriptorType sets BidDescriptorType, Tag 399
func (m NoBidDescriptors) SetBidDescriptorType(v int) {
	m.Set(field.NewBidDescriptorType(v))
}

//SetBidDescriptor sets BidDescriptor, Tag 400
func (m NoBidDescriptors) SetBidDescriptor(v string) {
	m.Set(field.NewBidDescriptor(v))
}

//SetSideValueInd sets SideValueInd, Tag 401
func (m NoBidDescriptors) SetSideValueInd(v int) {
	m.Set(field.NewSideValueInd(v))
}

//SetLiquidityValue sets LiquidityValue, Tag 404
func (m NoBidDescriptors) SetLiquidityValue(v float64) {
	m.Set(field.NewLiquidityValue(v))
}

//SetLiquidityNumSecurities sets LiquidityNumSecurities, Tag 441
func (m NoBidDescriptors) SetLiquidityNumSecurities(v int) {
	m.Set(field.NewLiquidityNumSecurities(v))
}

//SetLiquidityPctLow sets LiquidityPctLow, Tag 402
func (m NoBidDescriptors) SetLiquidityPctLow(v float64) {
	m.Set(field.NewLiquidityPctLow(v))
}

//SetLiquidityPctHigh sets LiquidityPctHigh, Tag 403
func (m NoBidDescriptors) SetLiquidityPctHigh(v float64) {
	m.Set(field.NewLiquidityPctHigh(v))
}

//SetEFPTrackingError sets EFPTrackingError, Tag 405
func (m NoBidDescriptors) SetEFPTrackingError(v float64) {
	m.Set(field.NewEFPTrackingError(v))
}

//SetFairValue sets FairValue, Tag 406
func (m NoBidDescriptors) SetFairValue(v float64) {
	m.Set(field.NewFairValue(v))
}

//SetOutsideIndexPct sets OutsideIndexPct, Tag 407
func (m NoBidDescriptors) SetOutsideIndexPct(v float64) {
	m.Set(field.NewOutsideIndexPct(v))
}

//SetValueOfFutures sets ValueOfFutures, Tag 408
func (m NoBidDescriptors) SetValueOfFutures(v float64) {
	m.Set(field.NewValueOfFutures(v))
}

//GetBidDescriptorType gets BidDescriptorType, Tag 399
func (m NoBidDescriptors) GetBidDescriptorType() (f field.BidDescriptorTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBidDescriptor gets BidDescriptor, Tag 400
func (m NoBidDescriptors) GetBidDescriptor() (f field.BidDescriptorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSideValueInd gets SideValueInd, Tag 401
func (m NoBidDescriptors) GetSideValueInd() (f field.SideValueIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLiquidityValue gets LiquidityValue, Tag 404
func (m NoBidDescriptors) GetLiquidityValue() (f field.LiquidityValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLiquidityNumSecurities gets LiquidityNumSecurities, Tag 441
func (m NoBidDescriptors) GetLiquidityNumSecurities() (f field.LiquidityNumSecuritiesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLiquidityPctLow gets LiquidityPctLow, Tag 402
func (m NoBidDescriptors) GetLiquidityPctLow() (f field.LiquidityPctLowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLiquidityPctHigh gets LiquidityPctHigh, Tag 403
func (m NoBidDescriptors) GetLiquidityPctHigh() (f field.LiquidityPctHighField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEFPTrackingError gets EFPTrackingError, Tag 405
func (m NoBidDescriptors) GetEFPTrackingError() (f field.EFPTrackingErrorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFairValue gets FairValue, Tag 406
func (m NoBidDescriptors) GetFairValue() (f field.FairValueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOutsideIndexPct gets OutsideIndexPct, Tag 407
func (m NoBidDescriptors) GetOutsideIndexPct() (f field.OutsideIndexPctField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValueOfFutures gets ValueOfFutures, Tag 408
func (m NoBidDescriptors) GetValueOfFutures() (f field.ValueOfFuturesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
}

//SetListID sets ListID, Tag 66
func (m NoBidComponents) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetSide sets Side, Tag 54
func (m NoBidComponents) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoBidComponents) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoBidComponents) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetNetGrossInd sets NetGrossInd, Tag 430
func (m NoBidComponents) SetNetGrossInd(v int) {
	m.Set(field.NewNetGrossInd(v))
}

//SetSettlType sets SettlType, Tag 63
func (m NoBidComponents) SetSettlType(v string) {
	m.Set(field.NewSettlType(v))
}

//SetSettlDate sets SettlDate, Tag 64
func (m NoBidComponents) SetSettlDate(v string) {
	m.Set(field.NewSettlDate(v))
}

//SetAccount sets Account, Tag 1
func (m NoBidComponents) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAcctIDSource sets AcctIDSource, Tag 660
func (m NoBidComponents) SetAcctIDSource(v int) {
	m.Set(field.NewAcctIDSource(v))
}

//GetListID gets ListID, Tag 66
func (m NoBidComponents) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m NoBidComponents) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoBidComponents) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoBidComponents) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetGrossInd gets NetGrossInd, Tag 430
func (m NoBidComponents) GetNetGrossInd() (f field.NetGrossIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlType gets SettlType, Tag 63
func (m NoBidComponents) GetSettlType() (f field.SettlTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlDate gets SettlDate, Tag 64
func (m NoBidComponents) GetSettlDate() (f field.SettlDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccount gets Account, Tag 1
func (m NoBidComponents) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAcctIDSource gets AcctIDSource, Tag 660
func (m NoBidComponents) GetAcctIDSource() (f field.AcctIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoBidComponents) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasNetGrossInd returns true if NetGrossInd is present, Tag 430
func (m NoBidComponents) HasNetGrossInd() bool {
	return m.Has(tag.NetGrossInd)
}

//HasSettlType returns true if SettlType is present, Tag 63
func (m NoBidComponents) HasSettlType() bool {
	return m.Has(tag.SettlType)
}

//HasSettlDate returns true if SettlDate is present, Tag 64
func (m NoBidComponents) HasSettlDate() bool {
	return m.Has(tag.SettlDate)
}

//HasAccount returns true if Account is present, Tag 1
func (m NoBidComponents) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAcctIDSource returns true if AcctIDSource is present, Tag 660
func (m NoBidComponents) HasAcctIDSource() bool {
	return m.Has(tag.AcctIDSource)
}

//NoBidComponentsRepeatingGroup is a repeating group, Tag 420
type NoBidComponentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoBidComponentsRepeatingGroup returns an initialized, NoBidComponentsRepeatingGroup
func NewNoBidComponentsRepeatingGroup() NoBidComponentsRepeatingGroup {
	return NoBidComponentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoBidComponents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.NetGrossInd), quickfix.GroupElement(tag.SettlType), quickfix.GroupElement(tag.SettlDate), quickfix.GroupElement(tag.Account), quickfix.GroupElement(tag.AcctIDSource)})}
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
