package bidcompreqgrp

//NoBidComponents is a repeating group in BidCompReqGrp
type NoBidComponents struct {
	//ListID is a non-required field for NoBidComponents.
	ListID *string `fix:"66"`
	//Side is a non-required field for NoBidComponents.
	Side *string `fix:"54"`
	//TradingSessionID is a non-required field for NoBidComponents.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoBidComponents.
	TradingSessionSubID *string `fix:"625"`
	//NetGrossInd is a non-required field for NoBidComponents.
	NetGrossInd *int `fix:"430"`
	//SettlType is a non-required field for NoBidComponents.
	SettlType *string `fix:"63"`
	//SettlDate is a non-required field for NoBidComponents.
	SettlDate *string `fix:"64"`
	//Account is a non-required field for NoBidComponents.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoBidComponents.
	AcctIDSource *int `fix:"660"`
}

//BidCompReqGrp is a fix50sp2 Component
type BidCompReqGrp struct {
	//NoBidComponents is a non-required field for BidCompReqGrp.
	NoBidComponents []NoBidComponents `fix:"420,omitempty"`
}

func (m *BidCompReqGrp) SetNoBidComponents(v []NoBidComponents) { m.NoBidComponents = v }
