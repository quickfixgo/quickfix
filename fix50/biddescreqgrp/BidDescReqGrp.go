package biddescreqgrp

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

//BidDescReqGrp is a fix50 Component
type BidDescReqGrp struct {
	//NoBidDescriptors is a non-required field for BidDescReqGrp.
	NoBidDescriptors []NoBidDescriptors `fix:"398,omitempty"`
}

func (m *BidDescReqGrp) SetNoBidDescriptors(v []NoBidDescriptors) { m.NoBidDescriptors = v }
