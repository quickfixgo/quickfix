package sidecrossordcxlgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"time"
)

//New returns an initialized SideCrossOrdCxlGrp instance
func New(nosides []NoSides) *SideCrossOrdCxlGrp {
	var m SideCrossOrdCxlGrp
	m.SetNoSides(nosides)
	return &m
}

//NoSides is a repeating group in SideCrossOrdCxlGrp
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//OrigClOrdID is a non-required field for NoSides.
	OrigClOrdID *string `fix:"41"`
	//ClOrdID is a required field for NoSides.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NoSides.
	ClOrdLinkID *string `fix:"583"`
	//OrigOrdModTime is a non-required field for NoSides.
	OrigOrdModTime *time.Time `fix:"586"`
	//Parties is a non-required component for NoSides.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoSides.
	TradeDate *string `fix:"75"`
	//OrderQtyData is a required component for NoSides.
	orderqtydata.OrderQtyData
	//ComplianceID is a non-required field for NoSides.
	ComplianceID *string `fix:"376"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
}

//NewNoSides returns an initialized NoSides instance
func NewNoSides(side string, clordid string, orderqtydata orderqtydata.OrderQtyData) *NoSides {
	var m NoSides
	m.SetSide(side)
	m.SetClOrdID(clordid)
	m.SetOrderQtyData(orderqtydata)
	return &m
}

func (m *NoSides) SetSide(v string)                            { m.Side = v }
func (m *NoSides) SetOrigClOrdID(v string)                     { m.OrigClOrdID = &v }
func (m *NoSides) SetClOrdID(v string)                         { m.ClOrdID = v }
func (m *NoSides) SetSecondaryClOrdID(v string)                { m.SecondaryClOrdID = &v }
func (m *NoSides) SetClOrdLinkID(v string)                     { m.ClOrdLinkID = &v }
func (m *NoSides) SetOrigOrdModTime(v time.Time)               { m.OrigOrdModTime = &v }
func (m *NoSides) SetParties(v parties.Parties)                { m.Parties = &v }
func (m *NoSides) SetTradeOriginationDate(v string)            { m.TradeOriginationDate = &v }
func (m *NoSides) SetTradeDate(v string)                       { m.TradeDate = &v }
func (m *NoSides) SetOrderQtyData(v orderqtydata.OrderQtyData) { m.OrderQtyData = v }
func (m *NoSides) SetComplianceID(v string)                    { m.ComplianceID = &v }
func (m *NoSides) SetText(v string)                            { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)                     { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)                     { m.EncodedText = &v }

//SideCrossOrdCxlGrp is a fix50sp1 Component
type SideCrossOrdCxlGrp struct {
	//NoSides is a required field for SideCrossOrdCxlGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *SideCrossOrdCxlGrp) SetNoSides(v []NoSides) { m.NoSides = v }
