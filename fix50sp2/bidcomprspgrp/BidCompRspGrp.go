package bidcomprspgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
)

//New returns an initialized BidCompRspGrp instance
func New(nobidcomponents []NoBidComponents) *BidCompRspGrp {
	var m BidCompRspGrp
	m.SetNoBidComponents(nobidcomponents)
	return &m
}

//NoBidComponents is a repeating group in BidCompRspGrp
type NoBidComponents struct {
	//CommissionData is a required component for NoBidComponents.
	commissiondata.CommissionData
	//ListID is a non-required field for NoBidComponents.
	ListID *string `fix:"66"`
	//Country is a non-required field for NoBidComponents.
	Country *string `fix:"421"`
	//Side is a non-required field for NoBidComponents.
	Side *string `fix:"54"`
	//Price is a non-required field for NoBidComponents.
	Price *float64 `fix:"44"`
	//PriceType is a non-required field for NoBidComponents.
	PriceType *int `fix:"423"`
	//FairValue is a non-required field for NoBidComponents.
	FairValue *float64 `fix:"406"`
	//NetGrossInd is a non-required field for NoBidComponents.
	NetGrossInd *int `fix:"430"`
	//SettlType is a non-required field for NoBidComponents.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoBidComponents.
	SettlDate *string `fix:"64"`
	//TradingSessionID is a non-required field for NoBidComponents.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoBidComponents.
	TradingSessionSubID *string `fix:"625"`
	//Text is a non-required field for NoBidComponents.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoBidComponents.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoBidComponents.
	EncodedText *string `fix:"355"`
}

//NewNoBidComponents returns an initialized NoBidComponents instance
func NewNoBidComponents(commissiondata commissiondata.CommissionData) *NoBidComponents {
	var m NoBidComponents
	m.SetCommissionData(commissiondata)
	return &m
}

func (m *NoBidComponents) SetCommissionData(v commissiondata.CommissionData) { m.CommissionData = v }
func (m *NoBidComponents) SetListID(v string)                                { m.ListID = &v }
func (m *NoBidComponents) SetCountry(v string)                               { m.Country = &v }
func (m *NoBidComponents) SetSide(v string)                                  { m.Side = &v }
func (m *NoBidComponents) SetPrice(v float64)                                { m.Price = &v }
func (m *NoBidComponents) SetPriceType(v int)                                { m.PriceType = &v }
func (m *NoBidComponents) SetFairValue(v float64)                            { m.FairValue = &v }
func (m *NoBidComponents) SetNetGrossInd(v int)                              { m.NetGrossInd = &v }
func (m *NoBidComponents) SetSettlType(v string)                             { m.SettlType = &v }
func (m *NoBidComponents) SetSettlDate(v string)                             { m.SettlDate = &v }
func (m *NoBidComponents) SetTradingSessionID(v string)                      { m.TradingSessionID = &v }
func (m *NoBidComponents) SetTradingSessionSubID(v string)                   { m.TradingSessionSubID = &v }
func (m *NoBidComponents) SetText(v string)                                  { m.Text = &v }
func (m *NoBidComponents) SetEncodedTextLen(v int)                           { m.EncodedTextLen = &v }
func (m *NoBidComponents) SetEncodedText(v string)                           { m.EncodedText = &v }

//BidCompRspGrp is a fix50sp2 Component
type BidCompRspGrp struct {
	//NoBidComponents is a required field for BidCompRspGrp.
	NoBidComponents []NoBidComponents `fix:"420"`
}

func (m *BidCompRspGrp) SetNoBidComponents(v []NoBidComponents) { m.NoBidComponents = v }
