package trdsesslstgrp

import (
	"time"
)

//New returns an initialized TrdSessLstGrp instance
func New(notradingsessions []NoTradingSessions) *TrdSessLstGrp {
	var m TrdSessLstGrp
	m.SetNoTradingSessions(notradingsessions)
	return &m
}

//NoTradingSessions is a repeating group in TrdSessLstGrp
type NoTradingSessions struct {
	//TradingSessionID is a required field for NoTradingSessions.
	TradingSessionID string `fix:"336"`
	//TradingSessionSubID is a non-required field for NoTradingSessions.
	TradingSessionSubID *string `fix:"625"`
	//SecurityExchange is a non-required field for NoTradingSessions.
	SecurityExchange *string `fix:"207"`
	//TradSesMethod is a non-required field for NoTradingSessions.
	TradSesMethod *int `fix:"338"`
	//TradSesMode is a non-required field for NoTradingSessions.
	TradSesMode *int `fix:"339"`
	//UnsolicitedIndicator is a non-required field for NoTradingSessions.
	UnsolicitedIndicator *bool `fix:"325"`
	//TradSesStatus is a required field for NoTradingSessions.
	TradSesStatus int `fix:"340"`
	//TradSesStatusRejReason is a non-required field for NoTradingSessions.
	TradSesStatusRejReason *int `fix:"567"`
	//TradSesStartTime is a non-required field for NoTradingSessions.
	TradSesStartTime *time.Time `fix:"341"`
	//TradSesOpenTime is a non-required field for NoTradingSessions.
	TradSesOpenTime *time.Time `fix:"342"`
	//TradSesPreCloseTime is a non-required field for NoTradingSessions.
	TradSesPreCloseTime *time.Time `fix:"343"`
	//TradSesCloseTime is a non-required field for NoTradingSessions.
	TradSesCloseTime *time.Time `fix:"344"`
	//TradSesEndTime is a non-required field for NoTradingSessions.
	TradSesEndTime *time.Time `fix:"345"`
	//TotalVolumeTraded is a non-required field for NoTradingSessions.
	TotalVolumeTraded *float64 `fix:"387"`
	//Text is a non-required field for NoTradingSessions.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoTradingSessions.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoTradingSessions.
	EncodedText *string `fix:"355"`
}

//NewNoTradingSessions returns an initialized NoTradingSessions instance
func NewNoTradingSessions(tradingsessionid string, tradsesstatus int) *NoTradingSessions {
	var m NoTradingSessions
	m.SetTradingSessionID(tradingsessionid)
	m.SetTradSesStatus(tradsesstatus)
	return &m
}

func (m *NoTradingSessions) SetTradingSessionID(v string)       { m.TradingSessionID = v }
func (m *NoTradingSessions) SetTradingSessionSubID(v string)    { m.TradingSessionSubID = &v }
func (m *NoTradingSessions) SetSecurityExchange(v string)       { m.SecurityExchange = &v }
func (m *NoTradingSessions) SetTradSesMethod(v int)             { m.TradSesMethod = &v }
func (m *NoTradingSessions) SetTradSesMode(v int)               { m.TradSesMode = &v }
func (m *NoTradingSessions) SetUnsolicitedIndicator(v bool)     { m.UnsolicitedIndicator = &v }
func (m *NoTradingSessions) SetTradSesStatus(v int)             { m.TradSesStatus = v }
func (m *NoTradingSessions) SetTradSesStatusRejReason(v int)    { m.TradSesStatusRejReason = &v }
func (m *NoTradingSessions) SetTradSesStartTime(v time.Time)    { m.TradSesStartTime = &v }
func (m *NoTradingSessions) SetTradSesOpenTime(v time.Time)     { m.TradSesOpenTime = &v }
func (m *NoTradingSessions) SetTradSesPreCloseTime(v time.Time) { m.TradSesPreCloseTime = &v }
func (m *NoTradingSessions) SetTradSesCloseTime(v time.Time)    { m.TradSesCloseTime = &v }
func (m *NoTradingSessions) SetTradSesEndTime(v time.Time)      { m.TradSesEndTime = &v }
func (m *NoTradingSessions) SetTotalVolumeTraded(v float64)     { m.TotalVolumeTraded = &v }
func (m *NoTradingSessions) SetText(v string)                   { m.Text = &v }
func (m *NoTradingSessions) SetEncodedTextLen(v int)            { m.EncodedTextLen = &v }
func (m *NoTradingSessions) SetEncodedText(v string)            { m.EncodedText = &v }

//TrdSessLstGrp is a fix50 Component
type TrdSessLstGrp struct {
	//NoTradingSessions is a required field for TrdSessLstGrp.
	NoTradingSessions []NoTradingSessions `fix:"386"`
}

func (m *TrdSessLstGrp) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
