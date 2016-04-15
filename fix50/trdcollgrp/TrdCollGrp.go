package trdcollgrp

func New() *TrdCollGrp {
	var m TrdCollGrp
	return &m
}

//NoTrades is a repeating group in TrdCollGrp
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

func (m *NoTrades) SetTradeReportID(v string)          { m.TradeReportID = &v }
func (m *NoTrades) SetSecondaryTradeReportID(v string) { m.SecondaryTradeReportID = &v }

//TrdCollGrp is a fix50 Component
type TrdCollGrp struct {
	//NoTrades is a non-required field for TrdCollGrp.
	NoTrades []NoTrades `fix:"897,omitempty"`
}

func (m *TrdCollGrp) SetNoTrades(v []NoTrades) { m.NoTrades = v }
