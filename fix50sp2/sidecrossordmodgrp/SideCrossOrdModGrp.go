package sidecrossordmodgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp2/commissiondata"
	"github.com/quickfixgo/quickfix/fix50sp2/orderqtydata"
	"github.com/quickfixgo/quickfix/fix50sp2/parties"
	"github.com/quickfixgo/quickfix/fix50sp2/preallocgrp"
	"time"
)

//New returns an initialized SideCrossOrdModGrp instance
func New(nosides []NoSides) *SideCrossOrdModGrp {
	var m SideCrossOrdModGrp
	m.SetNoSides(nosides)
	return &m
}

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

//NewNoSides returns an initialized NoSides instance
func NewNoSides(side string, clordid string, orderqtydata orderqtydata.OrderQtyData) *NoSides {
	var m NoSides
	m.SetSide(side)
	m.SetClOrdID(clordid)
	m.SetOrderQtyData(orderqtydata)
	return &m
}

func (m *NoSides) SetSide(v string)                                  { m.Side = v }
func (m *NoSides) SetClOrdID(v string)                               { m.ClOrdID = v }
func (m *NoSides) SetSecondaryClOrdID(v string)                      { m.SecondaryClOrdID = &v }
func (m *NoSides) SetClOrdLinkID(v string)                           { m.ClOrdLinkID = &v }
func (m *NoSides) SetParties(v parties.Parties)                      { m.Parties = &v }
func (m *NoSides) SetTradeOriginationDate(v string)                  { m.TradeOriginationDate = &v }
func (m *NoSides) SetTradeDate(v string)                             { m.TradeDate = &v }
func (m *NoSides) SetAccount(v string)                               { m.Account = &v }
func (m *NoSides) SetAcctIDSource(v int)                             { m.AcctIDSource = &v }
func (m *NoSides) SetAccountType(v int)                              { m.AccountType = &v }
func (m *NoSides) SetDayBookingInst(v string)                        { m.DayBookingInst = &v }
func (m *NoSides) SetBookingUnit(v string)                           { m.BookingUnit = &v }
func (m *NoSides) SetPreallocMethod(v string)                        { m.PreallocMethod = &v }
func (m *NoSides) SetAllocID(v string)                               { m.AllocID = &v }
func (m *NoSides) SetPreAllocGrp(v preallocgrp.PreAllocGrp)          { m.PreAllocGrp = &v }
func (m *NoSides) SetQtyType(v int)                                  { m.QtyType = &v }
func (m *NoSides) SetOrderQtyData(v orderqtydata.OrderQtyData)       { m.OrderQtyData = v }
func (m *NoSides) SetCommissionData(v commissiondata.CommissionData) { m.CommissionData = &v }
func (m *NoSides) SetOrderCapacity(v string)                         { m.OrderCapacity = &v }
func (m *NoSides) SetOrderRestrictions(v string)                     { m.OrderRestrictions = &v }
func (m *NoSides) SetCustOrderCapacity(v int)                        { m.CustOrderCapacity = &v }
func (m *NoSides) SetForexReq(v bool)                                { m.ForexReq = &v }
func (m *NoSides) SetSettlCurrency(v string)                         { m.SettlCurrency = &v }
func (m *NoSides) SetBookingType(v int)                              { m.BookingType = &v }
func (m *NoSides) SetText(v string)                                  { m.Text = &v }
func (m *NoSides) SetEncodedTextLen(v int)                           { m.EncodedTextLen = &v }
func (m *NoSides) SetEncodedText(v string)                           { m.EncodedText = &v }
func (m *NoSides) SetPositionEffect(v string)                        { m.PositionEffect = &v }
func (m *NoSides) SetCoveredOrUncovered(v int)                       { m.CoveredOrUncovered = &v }
func (m *NoSides) SetCashMargin(v string)                            { m.CashMargin = &v }
func (m *NoSides) SetClearingFeeIndicator(v string)                  { m.ClearingFeeIndicator = &v }
func (m *NoSides) SetSolicitedFlag(v bool)                           { m.SolicitedFlag = &v }
func (m *NoSides) SetSideComplianceID(v string)                      { m.SideComplianceID = &v }
func (m *NoSides) SetSideTimeInForce(v time.Time)                    { m.SideTimeInForce = &v }
func (m *NoSides) SetPreTradeAnonymity(v bool)                       { m.PreTradeAnonymity = &v }
func (m *NoSides) SetOrigClOrdID(v string)                           { m.OrigClOrdID = &v }

//SideCrossOrdModGrp is a fix50sp2 Component
type SideCrossOrdModGrp struct {
	//NoSides is a required field for SideCrossOrdModGrp.
	NoSides []NoSides `fix:"552"`
}

func (m *SideCrossOrdModGrp) SetNoSides(v []NoSides) { m.NoSides = v }
