package settlinstgrp

import (
	"github.com/quickfixgo/quickfix/fix50/parties"
	"github.com/quickfixgo/quickfix/fix50/settlinstructionsdata"
	"time"
)

func New() *SettlInstGrp {
	var m SettlInstGrp
	return &m
}

//NoSettlInst is a repeating group in SettlInstGrp
type NoSettlInst struct {
	//SettlInstID is a non-required field for NoSettlInst.
	SettlInstID *string `fix:"162"`
	//SettlInstTransType is a non-required field for NoSettlInst.
	SettlInstTransType *string `fix:"163"`
	//SettlInstRefID is a non-required field for NoSettlInst.
	SettlInstRefID *string `fix:"214"`
	//Parties is a non-required component for NoSettlInst.
	Parties *parties.Parties
	//Side is a non-required field for NoSettlInst.
	Side *string `fix:"54"`
	//Product is a non-required field for NoSettlInst.
	Product *int `fix:"460"`
	//SecurityType is a non-required field for NoSettlInst.
	SecurityType *string `fix:"167"`
	//CFICode is a non-required field for NoSettlInst.
	CFICode *string `fix:"461"`
	//EffectiveTime is a non-required field for NoSettlInst.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireTime is a non-required field for NoSettlInst.
	ExpireTime *time.Time `fix:"126"`
	//LastUpdateTime is a non-required field for NoSettlInst.
	LastUpdateTime *time.Time `fix:"779"`
	//SettlInstructionsData is a non-required component for NoSettlInst.
	SettlInstructionsData *settlinstructionsdata.SettlInstructionsData
	//PaymentMethod is a non-required field for NoSettlInst.
	PaymentMethod *int `fix:"492"`
	//PaymentRef is a non-required field for NoSettlInst.
	PaymentRef *string `fix:"476"`
	//CardHolderName is a non-required field for NoSettlInst.
	CardHolderName *string `fix:"488"`
	//CardNumber is a non-required field for NoSettlInst.
	CardNumber *string `fix:"489"`
	//CardStartDate is a non-required field for NoSettlInst.
	CardStartDate *string `fix:"503"`
	//CardExpDate is a non-required field for NoSettlInst.
	CardExpDate *string `fix:"490"`
	//CardIssNum is a non-required field for NoSettlInst.
	CardIssNum *string `fix:"491"`
	//PaymentDate is a non-required field for NoSettlInst.
	PaymentDate *string `fix:"504"`
	//PaymentRemitterID is a non-required field for NoSettlInst.
	PaymentRemitterID *string `fix:"505"`
	//SettlCurrency is a non-required field for NoSettlInst.
	SettlCurrency *string `fix:"120"`
}

func (m *NoSettlInst) SetSettlInstID(v string)        { m.SettlInstID = &v }
func (m *NoSettlInst) SetSettlInstTransType(v string) { m.SettlInstTransType = &v }
func (m *NoSettlInst) SetSettlInstRefID(v string)     { m.SettlInstRefID = &v }
func (m *NoSettlInst) SetParties(v parties.Parties)   { m.Parties = &v }
func (m *NoSettlInst) SetSide(v string)               { m.Side = &v }
func (m *NoSettlInst) SetProduct(v int)               { m.Product = &v }
func (m *NoSettlInst) SetSecurityType(v string)       { m.SecurityType = &v }
func (m *NoSettlInst) SetCFICode(v string)            { m.CFICode = &v }
func (m *NoSettlInst) SetEffectiveTime(v time.Time)   { m.EffectiveTime = &v }
func (m *NoSettlInst) SetExpireTime(v time.Time)      { m.ExpireTime = &v }
func (m *NoSettlInst) SetLastUpdateTime(v time.Time)  { m.LastUpdateTime = &v }
func (m *NoSettlInst) SetSettlInstructionsData(v settlinstructionsdata.SettlInstructionsData) {
	m.SettlInstructionsData = &v
}
func (m *NoSettlInst) SetPaymentMethod(v int)        { m.PaymentMethod = &v }
func (m *NoSettlInst) SetPaymentRef(v string)        { m.PaymentRef = &v }
func (m *NoSettlInst) SetCardHolderName(v string)    { m.CardHolderName = &v }
func (m *NoSettlInst) SetCardNumber(v string)        { m.CardNumber = &v }
func (m *NoSettlInst) SetCardStartDate(v string)     { m.CardStartDate = &v }
func (m *NoSettlInst) SetCardExpDate(v string)       { m.CardExpDate = &v }
func (m *NoSettlInst) SetCardIssNum(v string)        { m.CardIssNum = &v }
func (m *NoSettlInst) SetPaymentDate(v string)       { m.PaymentDate = &v }
func (m *NoSettlInst) SetPaymentRemitterID(v string) { m.PaymentRemitterID = &v }
func (m *NoSettlInst) SetSettlCurrency(v string)     { m.SettlCurrency = &v }

//SettlInstGrp is a fix50 Component
type SettlInstGrp struct {
	//NoSettlInst is a non-required field for SettlInstGrp.
	NoSettlInst []NoSettlInst `fix:"778,omitempty"`
}

func (m *SettlInstGrp) SetNoSettlInst(v []NoSettlInst) { m.NoSettlInst = v }
