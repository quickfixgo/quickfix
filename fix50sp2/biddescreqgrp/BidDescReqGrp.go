package biddescreqgrp

//New returns an initialized BidDescReqGrp instance
func New() *BidDescReqGrp {
	var m BidDescReqGrp
	return &m
}

//NoBidDescriptors is a repeating group in BidDescReqGrp
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

//BidDescReqGrp is a fix50sp2 Component
type BidDescReqGrp struct {
	//NoBidDescriptors is a non-required field for BidDescReqGrp.
	NoBidDescriptors []NoBidDescriptors `fix:"398,omitempty"`
}

func (m *BidDescReqGrp) SetNoBidDescriptors(v []NoBidDescriptors) { m.NoBidDescriptors = v }
