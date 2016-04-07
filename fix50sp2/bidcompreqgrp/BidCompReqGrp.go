package bidcompreqgrp

func New() *BidCompReqGrp {
	var m BidCompReqGrp
	return &m
}

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

func (m *NoBidComponents) SetListID(v string)              { m.ListID = &v }
func (m *NoBidComponents) SetSide(v string)                { m.Side = &v }
func (m *NoBidComponents) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoBidComponents) SetTradingSessionSubID(v string) { m.TradingSessionSubID = &v }
func (m *NoBidComponents) SetNetGrossInd(v int)            { m.NetGrossInd = &v }
func (m *NoBidComponents) SetSettlType(v string)           { m.SettlType = &v }
func (m *NoBidComponents) SetSettlDate(v string)           { m.SettlDate = &v }
func (m *NoBidComponents) SetAccount(v string)             { m.Account = &v }
func (m *NoBidComponents) SetAcctIDSource(v int)           { m.AcctIDSource = &v }

//BidCompReqGrp is a fix50sp2 Component
type BidCompReqGrp struct {
	//NoBidComponents is a non-required field for BidCompReqGrp.
	NoBidComponents []NoBidComponents `fix:"420,omitempty"`
}

func (m *BidCompReqGrp) SetNoBidComponents(v []NoBidComponents) { m.NoBidComponents = v }
