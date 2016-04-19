package sectypesgrp

import (
	"time"
)

//New returns an initialized SecTypesGrp instance
func New() *SecTypesGrp {
	var m SecTypesGrp
	return &m
}

//NoSecurityTypes is a repeating group in SecTypesGrp
type NoSecurityTypes struct {
	//SecurityType is a non-required field for NoSecurityTypes.
	SecurityType *string `fix:"167"`
	//SecuritySubType is a non-required field for NoSecurityTypes.
	SecuritySubType *string `fix:"762"`
	//Product is a non-required field for NoSecurityTypes.
	Product *int `fix:"460"`
	//CFICode is a non-required field for NoSecurityTypes.
	CFICode *string `fix:"461"`
	//TransactTime is a non-required field for NoSecurityTypes.
	TransactTime *time.Time `fix:"60"`
}

//NewNoSecurityTypes returns an initialized NoSecurityTypes instance
func NewNoSecurityTypes() *NoSecurityTypes {
	var m NoSecurityTypes
	return &m
}

func (m *NoSecurityTypes) SetSecurityType(v string)    { m.SecurityType = &v }
func (m *NoSecurityTypes) SetSecuritySubType(v string) { m.SecuritySubType = &v }
func (m *NoSecurityTypes) SetProduct(v int)            { m.Product = &v }
func (m *NoSecurityTypes) SetCFICode(v string)         { m.CFICode = &v }
func (m *NoSecurityTypes) SetTransactTime(v time.Time) { m.TransactTime = &v }

//SecTypesGrp is a fix50sp2 Component
type SecTypesGrp struct {
	//NoSecurityTypes is a non-required field for SecTypesGrp.
	NoSecurityTypes []NoSecurityTypes `fix:"558,omitempty"`
}

func (m *SecTypesGrp) SetNoSecurityTypes(v []NoSecurityTypes) { m.NoSecurityTypes = v }
