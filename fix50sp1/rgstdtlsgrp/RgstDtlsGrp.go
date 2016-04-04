package rgstdtlsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

func New() *RgstDtlsGrp {
	var m RgstDtlsGrp
	return &m
}

//NoRegistDtls is a repeating group in RgstDtlsGrp
type NoRegistDtls struct {
	//RegistDtls is a non-required field for NoRegistDtls.
	RegistDtls *string `fix:"509"`
	//RegistEmail is a non-required field for NoRegistDtls.
	RegistEmail *string `fix:"511"`
	//MailingDtls is a non-required field for NoRegistDtls.
	MailingDtls *string `fix:"474"`
	//MailingInst is a non-required field for NoRegistDtls.
	MailingInst *string `fix:"482"`
	//NestedParties is a non-required component for NoRegistDtls.
	NestedParties *nestedparties.NestedParties
	//OwnerType is a non-required field for NoRegistDtls.
	OwnerType *int `fix:"522"`
	//DateOfBirth is a non-required field for NoRegistDtls.
	DateOfBirth *string `fix:"486"`
	//InvestorCountryOfResidence is a non-required field for NoRegistDtls.
	InvestorCountryOfResidence *string `fix:"475"`
}

func (m *NoRegistDtls) SetRegistDtls(v string)                         { m.RegistDtls = &v }
func (m *NoRegistDtls) SetRegistEmail(v string)                        { m.RegistEmail = &v }
func (m *NoRegistDtls) SetMailingDtls(v string)                        { m.MailingDtls = &v }
func (m *NoRegistDtls) SetMailingInst(v string)                        { m.MailingInst = &v }
func (m *NoRegistDtls) SetNestedParties(v nestedparties.NestedParties) { m.NestedParties = &v }
func (m *NoRegistDtls) SetOwnerType(v int)                             { m.OwnerType = &v }
func (m *NoRegistDtls) SetDateOfBirth(v string)                        { m.DateOfBirth = &v }
func (m *NoRegistDtls) SetInvestorCountryOfResidence(v string)         { m.InvestorCountryOfResidence = &v }

//RgstDtlsGrp is a fix50sp1 Component
type RgstDtlsGrp struct {
	//NoRegistDtls is a non-required field for RgstDtlsGrp.
	NoRegistDtls []NoRegistDtls `fix:"473,omitempty"`
}

func (m *RgstDtlsGrp) SetNoRegistDtls(v []NoRegistDtls) { m.NoRegistDtls = v }
