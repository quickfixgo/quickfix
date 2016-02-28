package trdgsesgrp

//NoTradingSessions is a repeating group in TrdgSesGrp
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
}

//Component is a fix50 TrdgSesGrp Component
type Component struct {
	//NoTradingSessions is a non-required field for TrdgSesGrp.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
