package bidresponse

import (
	"github.com/shopspring/decimal"

	"github.com/alpacahq/quickfix"
	"github.com/alpacahq/quickfix/enum"
	"github.com/alpacahq/quickfix/field"
	"github.com/alpacahq/quickfix/fix42"
	"github.com/alpacahq/quickfix/tag"
)

//BidResponse is the fix42 BidResponse type, MsgType = l
type BidResponse struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a BidResponse from a quickfix.Message instance
func FromMessage(m *quickfix.Message) BidResponse {
	return BidResponse{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m BidResponse) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a BidResponse initialized with the required fields for BidResponse
func New() (m BidResponse) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("l"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg BidResponse, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "l", r
}

//SetBidID sets BidID, Tag 390
func (m BidResponse) SetBidID(v string) {
	m.Set(field.NewBidID(v))
}

//SetClientBidID sets ClientBidID, Tag 391
func (m BidResponse) SetClientBidID(v string) {
	m.Set(field.NewClientBidID(v))
}

//SetNoBidComponents sets NoBidComponents, Tag 420
func (m BidResponse) SetNoBidComponents(f NoBidComponentsRepeatingGroup) {
	m.SetGroup(f)
}

//GetBidID gets BidID, Tag 390
func (m BidResponse) GetBidID() (v string, err quickfix.MessageRejectError) {
	var f field.BidIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClientBidID gets ClientBidID, Tag 391
func (m BidResponse) GetClientBidID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientBidIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoBidComponents gets NoBidComponents, Tag 420
func (m BidResponse) GetNoBidComponents() (NoBidComponentsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoBidComponentsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasBidID returns true if BidID is present, Tag 390
func (m BidResponse) HasBidID() bool {
	return m.Has(tag.BidID)
}

//HasClientBidID returns true if ClientBidID is present, Tag 391
func (m BidResponse) HasClientBidID() bool {
	return m.Has(tag.ClientBidID)
}

//HasNoBidComponents returns true if NoBidComponents is present, Tag 420
func (m BidResponse) HasNoBidComponents() bool {
	return m.Has(tag.NoBidComponents)
}

//NoBidComponents is a repeating group element, Tag 420
type NoBidComponents struct {
	*quickfix.Group
}

//SetCommission sets Commission, Tag 12
func (m NoBidComponents) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NoBidComponents) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetListID sets ListID, Tag 66
func (m NoBidComponents) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetCountry sets Country, Tag 421
func (m NoBidComponents) SetCountry(v string) {
	m.Set(field.NewCountry(v))
}

//SetSide sets Side, Tag 54
func (m NoBidComponents) SetSide(v enum.Side) {
	m.Set(field.NewSide(v))
}

//SetPrice sets Price, Tag 44
func (m NoBidComponents) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetPriceType sets PriceType, Tag 423
func (m NoBidComponents) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetFairValue sets FairValue, Tag 406
func (m NoBidComponents) SetFairValue(value decimal.Decimal, scale int32) {
	m.Set(field.NewFairValue(value, scale))
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

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoBidComponents) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetText sets Text, Tag 58
func (m NoBidComponents) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoBidComponents) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoBidComponents) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetCommission gets Commission, Tag 12
func (m NoBidComponents) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m NoBidComponents) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m NoBidComponents) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountry gets Country, Tag 421
func (m NoBidComponents) GetCountry() (v string, err quickfix.MessageRejectError) {
	var f field.CountryField
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

//GetPrice gets Price, Tag 44
func (m NoBidComponents) GetPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.PriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m NoBidComponents) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFairValue gets FairValue, Tag 406
func (m NoBidComponents) GetFairValue() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FairValueField
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

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoBidComponents) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m NoBidComponents) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoBidComponents) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoBidComponents) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoBidComponents) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoBidComponents) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasListID returns true if ListID is present, Tag 66
func (m NoBidComponents) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasCountry returns true if Country is present, Tag 421
func (m NoBidComponents) HasCountry() bool {
	return m.Has(tag.Country)
}

//HasSide returns true if Side is present, Tag 54
func (m NoBidComponents) HasSide() bool {
	return m.Has(tag.Side)
}

//HasPrice returns true if Price is present, Tag 44
func (m NoBidComponents) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m NoBidComponents) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasFairValue returns true if FairValue is present, Tag 406
func (m NoBidComponents) HasFairValue() bool {
	return m.Has(tag.FairValue)
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

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoBidComponents) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasText returns true if Text is present, Tag 58
func (m NoBidComponents) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoBidComponents) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoBidComponents) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//NoBidComponentsRepeatingGroup is a repeating group, Tag 420
type NoBidComponentsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoBidComponentsRepeatingGroup returns an initialized, NoBidComponentsRepeatingGroup
func NewNoBidComponentsRepeatingGroup() NoBidComponentsRepeatingGroup {
	return NoBidComponentsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoBidComponents,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.Country), quickfix.GroupElement(tag.Side), quickfix.GroupElement(tag.Price), quickfix.GroupElement(tag.PriceType), quickfix.GroupElement(tag.FairValue), quickfix.GroupElement(tag.NetGrossInd), quickfix.GroupElement(tag.SettlmntTyp), quickfix.GroupElement(tag.FutSettDate), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
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
