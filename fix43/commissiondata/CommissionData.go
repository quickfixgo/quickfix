package commissiondata

func New() *CommissionData {
	var m CommissionData
	return &m
}

//CommissionData is a fix43 Component
type CommissionData struct {
	//Commission is a non-required field for CommissionData.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for CommissionData.
	CommType *string `fix:"13"`
	//CommCurrency is a non-required field for CommissionData.
	CommCurrency *string `fix:"479"`
	//FundRenewWaiv is a non-required field for CommissionData.
	FundRenewWaiv *string `fix:"497"`
}

func (m *CommissionData) SetCommission(v float64)   { m.Commission = &v }
func (m *CommissionData) SetCommType(v string)      { m.CommType = &v }
func (m *CommissionData) SetCommCurrency(v string)  { m.CommCurrency = &v }
func (m *CommissionData) SetFundRenewWaiv(v string) { m.FundRenewWaiv = &v }
