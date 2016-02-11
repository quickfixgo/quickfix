package trdcollgrp

//NoTrades is a repeating group in TrdCollGrp
type NoTrades struct {
	//TradeReportID is a non-required field for NoTrades.
	TradeReportID *string `fix:"571"`
	//SecondaryTradeReportID is a non-required field for NoTrades.
	SecondaryTradeReportID *string `fix:"818"`
}

//Component is a fix50sp2 TrdCollGrp Component
type Component struct {
	//NoTrades is a non-required field for TrdCollGrp.
	NoTrades []NoTrades `fix:"897,omitempty"`
}

func New() *Component { return new(Component) }
