package allocgrp

import (
	"github.com/quickfixgo/quickfix/fix50/clrinstgrp"
	"github.com/quickfixgo/quickfix/fix50/commissiondata"
	"github.com/quickfixgo/quickfix/fix50/miscfeesgrp"
	"github.com/quickfixgo/quickfix/fix50/nestedparties"
	"github.com/quickfixgo/quickfix/fix50/settlinstructionsdata"
)

//New returns an initialized AllocGrp instance
func New() *AllocGrp {
	var m AllocGrp
	return &m
}

//NoAllocs is a repeating group in AllocGrp
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocAcctIDSource is a non-required field for NoAllocs.
	AllocAcctIDSource *int `fix:"661"`
	//MatchStatus is a non-required field for NoAllocs.
	MatchStatus *string `fix:"573"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//NestedParties is a non-required component for NoAllocs.
	NestedParties *nestedparties.NestedParties
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *bool `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//CommissionData is a non-required component for NoAllocs.
	CommissionData *commissiondata.CommissionData
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//AllocSettlCurrAmt is a non-required field for NoAllocs.
	AllocSettlCurrAmt *float64 `fix:"737"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//AllocSettlCurrency is a non-required field for NoAllocs.
	AllocSettlCurrency *string `fix:"736"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AllocAccruedInterestAmt is a non-required field for NoAllocs.
	AllocAccruedInterestAmt *float64 `fix:"742"`
	//AllocInterestAtMaturity is a non-required field for NoAllocs.
	AllocInterestAtMaturity *float64 `fix:"741"`
	//MiscFeesGrp is a non-required component for NoAllocs.
	MiscFeesGrp *miscfeesgrp.MiscFeesGrp
	//ClrInstGrp is a non-required component for NoAllocs.
	ClrInstGrp *clrinstgrp.ClrInstGrp
	//AllocSettlInstType is a non-required field for NoAllocs.
	AllocSettlInstType *int `fix:"780"`
	//SettlInstructionsData is a non-required component for NoAllocs.
	SettlInstructionsData *settlinstructionsdata.SettlInstructionsData
	//SecondaryIndividualAllocID is a non-required field for NoAllocs.
	SecondaryIndividualAllocID *string `fix:"989"`
	//AllocMethod is a non-required field for NoAllocs.
	AllocMethod *int `fix:"1002"`
	//AllocCustomerCapacity is a non-required field for NoAllocs.
	AllocCustomerCapacity *string `fix:"993"`
	//IndividualAllocType is a non-required field for NoAllocs.
	IndividualAllocType *int `fix:"992"`
	//AllocPositionEffect is a non-required field for NoAllocs.
	AllocPositionEffect *string `fix:"1047"`
	//ClearingFeeIndicator is a non-required field for NoAllocs.
	ClearingFeeIndicator *string `fix:"635"`
}

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string)                          { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocAcctIDSource(v int)                        { m.AllocAcctIDSource = &v }
func (m *NoAllocs) SetMatchStatus(v string)                           { m.MatchStatus = &v }
func (m *NoAllocs) SetAllocPrice(v float64)                           { m.AllocPrice = &v }
func (m *NoAllocs) SetAllocQty(v float64)                             { m.AllocQty = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)                     { m.IndividualAllocID = &v }
func (m *NoAllocs) SetProcessCode(v string)                           { m.ProcessCode = &v }
func (m *NoAllocs) SetNestedParties(v nestedparties.NestedParties)    { m.NestedParties = &v }
func (m *NoAllocs) SetNotifyBrokerOfCredit(v bool)                    { m.NotifyBrokerOfCredit = &v }
func (m *NoAllocs) SetAllocHandlInst(v int)                           { m.AllocHandlInst = &v }
func (m *NoAllocs) SetAllocText(v string)                             { m.AllocText = &v }
func (m *NoAllocs) SetEncodedAllocTextLen(v int)                      { m.EncodedAllocTextLen = &v }
func (m *NoAllocs) SetEncodedAllocText(v string)                      { m.EncodedAllocText = &v }
func (m *NoAllocs) SetCommissionData(v commissiondata.CommissionData) { m.CommissionData = &v }
func (m *NoAllocs) SetAllocAvgPx(v float64)                           { m.AllocAvgPx = &v }
func (m *NoAllocs) SetAllocNetMoney(v float64)                        { m.AllocNetMoney = &v }
func (m *NoAllocs) SetSettlCurrAmt(v float64)                         { m.SettlCurrAmt = &v }
func (m *NoAllocs) SetAllocSettlCurrAmt(v float64)                    { m.AllocSettlCurrAmt = &v }
func (m *NoAllocs) SetSettlCurrency(v string)                         { m.SettlCurrency = &v }
func (m *NoAllocs) SetAllocSettlCurrency(v string)                    { m.AllocSettlCurrency = &v }
func (m *NoAllocs) SetSettlCurrFxRate(v float64)                      { m.SettlCurrFxRate = &v }
func (m *NoAllocs) SetSettlCurrFxRateCalc(v string)                   { m.SettlCurrFxRateCalc = &v }
func (m *NoAllocs) SetAllocAccruedInterestAmt(v float64)              { m.AllocAccruedInterestAmt = &v }
func (m *NoAllocs) SetAllocInterestAtMaturity(v float64)              { m.AllocInterestAtMaturity = &v }
func (m *NoAllocs) SetMiscFeesGrp(v miscfeesgrp.MiscFeesGrp)          { m.MiscFeesGrp = &v }
func (m *NoAllocs) SetClrInstGrp(v clrinstgrp.ClrInstGrp)             { m.ClrInstGrp = &v }
func (m *NoAllocs) SetAllocSettlInstType(v int)                       { m.AllocSettlInstType = &v }
func (m *NoAllocs) SetSettlInstructionsData(v settlinstructionsdata.SettlInstructionsData) {
	m.SettlInstructionsData = &v
}
func (m *NoAllocs) SetSecondaryIndividualAllocID(v string) { m.SecondaryIndividualAllocID = &v }
func (m *NoAllocs) SetAllocMethod(v int)                   { m.AllocMethod = &v }
func (m *NoAllocs) SetAllocCustomerCapacity(v string)      { m.AllocCustomerCapacity = &v }
func (m *NoAllocs) SetIndividualAllocType(v int)           { m.IndividualAllocType = &v }
func (m *NoAllocs) SetAllocPositionEffect(v string)        { m.AllocPositionEffect = &v }
func (m *NoAllocs) SetClearingFeeIndicator(v string)       { m.ClearingFeeIndicator = &v }

//AllocGrp is a fix50 Component
type AllocGrp struct {
	//NoAllocs is a non-required field for AllocGrp.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
}

func (m *AllocGrp) SetNoAllocs(v []NoAllocs) { m.NoAllocs = v }
