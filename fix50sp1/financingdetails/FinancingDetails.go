package financingdetails

//New returns an initialized FinancingDetails instance
func New() *FinancingDetails {
	var m FinancingDetails
	return &m
}

//FinancingDetails is a fix50sp1 Component
type FinancingDetails struct {
	//AgreementDesc is a non-required field for FinancingDetails.
	AgreementDesc *string `fix:"913"`
	//AgreementID is a non-required field for FinancingDetails.
	AgreementID *string `fix:"914"`
	//AgreementDate is a non-required field for FinancingDetails.
	AgreementDate *string `fix:"915"`
	//AgreementCurrency is a non-required field for FinancingDetails.
	AgreementCurrency *string `fix:"918"`
	//TerminationType is a non-required field for FinancingDetails.
	TerminationType *int `fix:"788"`
	//StartDate is a non-required field for FinancingDetails.
	StartDate *string `fix:"916"`
	//EndDate is a non-required field for FinancingDetails.
	EndDate *string `fix:"917"`
	//DeliveryType is a non-required field for FinancingDetails.
	DeliveryType *int `fix:"919"`
	//MarginRatio is a non-required field for FinancingDetails.
	MarginRatio *float64 `fix:"898"`
}

func (m *FinancingDetails) SetAgreementDesc(v string)     { m.AgreementDesc = &v }
func (m *FinancingDetails) SetAgreementID(v string)       { m.AgreementID = &v }
func (m *FinancingDetails) SetAgreementDate(v string)     { m.AgreementDate = &v }
func (m *FinancingDetails) SetAgreementCurrency(v string) { m.AgreementCurrency = &v }
func (m *FinancingDetails) SetTerminationType(v int)      { m.TerminationType = &v }
func (m *FinancingDetails) SetStartDate(v string)         { m.StartDate = &v }
func (m *FinancingDetails) SetEndDate(v string)           { m.EndDate = &v }
func (m *FinancingDetails) SetDeliveryType(v int)         { m.DeliveryType = &v }
func (m *FinancingDetails) SetMarginRatio(v float64)      { m.MarginRatio = &v }
