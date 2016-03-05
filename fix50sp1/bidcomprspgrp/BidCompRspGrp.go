package bidcomprspgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
)

//NoBidComponents is a repeating group in BidCompRspGrp
type NoBidComponents struct {
	//CommissionData Component
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

//BidCompRspGrp is a fix50sp1 Component
type BidCompRspGrp struct {
	//NoBidComponents is a required field for BidCompRspGrp.
	NoBidComponents []NoBidComponents `fix:"420"`
}

func (m *BidCompRspGrp) SetNoBidComponents(v []NoBidComponents) { m.NoBidComponents = v }
