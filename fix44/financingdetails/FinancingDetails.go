package financingdetails

//Component is a fix44 FinancingDetails Component
type Component struct {
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

func New() *Component { return new(Component) }
