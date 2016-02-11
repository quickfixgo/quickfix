package rgstdtlsgrp

import (
	"github.com/quickfixgo/quickfix/fix50sp1/nestedparties"
)

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
	//NestedParties Component
	NestedParties nestedparties.Component
	//OwnerType is a non-required field for NoRegistDtls.
	OwnerType *int `fix:"522"`
	//DateOfBirth is a non-required field for NoRegistDtls.
	DateOfBirth *string `fix:"486"`
	//InvestorCountryOfResidence is a non-required field for NoRegistDtls.
	InvestorCountryOfResidence *string `fix:"475"`
}

//Component is a fix50sp1 RgstDtlsGrp Component
type Component struct {
	//NoRegistDtls is a non-required field for RgstDtlsGrp.
	NoRegistDtls []NoRegistDtls `fix:"473,omitempty"`
}

func New() *Component { return new(Component) }
