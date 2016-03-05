package sidecrossordcxlgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"time"
)

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
	//Parties Component
	parties.Parties
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoSides.
	TradeDate *string `fix:"75"`
	//OrderQtyData Component
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

//SideCrossOrdCxlGrp is a fix50sp2 Component
type SideCrossOrdCxlGrp struct {
	//NoSides is a required field for SideCrossOrdCxlGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *SideCrossOrdCxlGrp) SetNoSides(v []NoSides) { m.NoSides = v }
