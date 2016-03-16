package sidecrossordmodgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp1/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp1/parties"
	"github.com/quickfixgo/quickfix/fix50sp1/preallocgrp"
	"time"
)

//NoSides is a repeating group in SideCrossOrdModGrp
type NoSides struct {
	//Side is a required field for NoSides.
	Side string `fix:"54"`
	//ClOrdID is a required field for NoSides.
	ClOrdID string `fix:"11"`
	//SecondaryClOrdID is a non-required field for NoSides.
	SecondaryClOrdID *string `fix:"526"`
	//ClOrdLinkID is a non-required field for NoSides.
	ClOrdLinkID *string `fix:"583"`
	//Parties is a non-required component for NoSides.
	Parties *parties.Parties
	//TradeOriginationDate is a non-required field for NoSides.
	TradeOriginationDate *string `fix:"229"`
	//TradeDate is a non-required field for NoSides.
	TradeDate *string `fix:"75"`
	//Account is a non-required field for NoSides.
	Account *string `fix:"1"`
	//AcctIDSource is a non-required field for NoSides.
	AcctIDSource *int `fix:"660"`
	//AccountType is a non-required field for NoSides.
	AccountType *int `fix:"581"`
	//DayBookingInst is a non-required field for NoSides.
	DayBookingInst *string `fix:"589"`
	//BookingUnit is a non-required field for NoSides.
	BookingUnit *string `fix:"590"`
	//PreallocMethod is a non-required field for NoSides.
	PreallocMethod *string `fix:"591"`
	//AllocID is a non-required field for NoSides.
	AllocID *string `fix:"70"`
	//PreAllocGrp is a non-required component for NoSides.
	PreAllocGrp *preallocgrp.PreAllocGrp
	//QtyType is a non-required field for NoSides.
	QtyType *int `fix:"854"`
	//OrderQtyData is a required component for NoSides.
	orderqtydata.OrderQtyData
	//CommissionData is a non-required component for NoSides.
	CommissionData *commissiondata.CommissionData
	//OrderCapacity is a non-required field for NoSides.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for NoSides.
	OrderRestrictions *string `fix:"529"`
	//CustOrderCapacity is a non-required field for NoSides.
	CustOrderCapacity *int `fix:"582"`
	//ForexReq is a non-required field for NoSides.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NoSides.
	SettlCurrency *string `fix:"120"`
	//BookingType is a non-required field for NoSides.
	BookingType *int `fix:"775"`
	//Text is a non-required field for NoSides.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoSides.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoSides.
	EncodedText *string `fix:"355"`
	//PositionEffect is a non-required field for NoSides.
	PositionEffect *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NoSides.
	CoveredOrUncovered *int `fix:"203"`
	//CashMargin is a non-required field for NoSides.
	CashMargin *string `fix:"544"`
	//ClearingFeeIndicator is a non-required field for NoSides.
	ClearingFeeIndicator *string `fix:"635"`
	//SolicitedFlag is a non-required field for NoSides.
	SolicitedFlag *bool `fix:"377"`
	//SideComplianceID is a non-required field for NoSides.
	SideComplianceID *string `fix:"659"`
	//SideTimeInForce is a non-required field for NoSides.
	SideTimeInForce *time.Time `fix:"962"`
	//PreTradeAnonymity is a non-required field for NoSides.
	PreTradeAnonymity *bool `fix:"1091"`
	//OrigClOrdID is a non-required field for NoSides.
	OrigClOrdID *string `fix:"41"`
}

//SideCrossOrdModGrp is a fix50sp1 Component
type SideCrossOrdModGrp struct {
	//NoSides is a required field for SideCrossOrdModGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *SideCrossOrdModGrp) SetNoSides(v []NoSides) { m.NoSides = v }
