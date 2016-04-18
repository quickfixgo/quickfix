package underlyingamount

//New returns an initialized UnderlyingAmount instance
func New() *UnderlyingAmount {
	var m UnderlyingAmount
	return &m
}

//NoUnderlyingAmounts is a repeating group in UnderlyingAmount
type NoUnderlyingAmounts struct {
	//UnderlyingPayAmount is a non-required field for NoUnderlyingAmounts.
	UnderlyingPayAmount *float64 `fix:"985"`
	//UnderlyingCollectAmount is a non-required field for NoUnderlyingAmounts.
	UnderlyingCollectAmount *float64 `fix:"986"`
	//UnderlyingSettlementDate is a non-required field for NoUnderlyingAmounts.
	UnderlyingSettlementDate *string `fix:"987"`
	//UnderlyingSettlementStatus is a non-required field for NoUnderlyingAmounts.
	UnderlyingSettlementStatus *string `fix:"988"`
}

//NewNoUnderlyingAmounts returns an initialized NoUnderlyingAmounts instance
func NewNoUnderlyingAmounts() *NoUnderlyingAmounts {
	var m NoUnderlyingAmounts
	return &m
}

func (m *NoUnderlyingAmounts) SetUnderlyingPayAmount(v float64)     { m.UnderlyingPayAmount = &v }
func (m *NoUnderlyingAmounts) SetUnderlyingCollectAmount(v float64) { m.UnderlyingCollectAmount = &v }
func (m *NoUnderlyingAmounts) SetUnderlyingSettlementDate(v string) { m.UnderlyingSettlementDate = &v }
func (m *NoUnderlyingAmounts) SetUnderlyingSettlementStatus(v string) {
	m.UnderlyingSettlementStatus = &v
}

//UnderlyingAmount is a fix50sp2 Component
type UnderlyingAmount struct {
	//NoUnderlyingAmounts is a non-required field for UnderlyingAmount.
	NoUnderlyingAmounts []NoUnderlyingAmounts `fix:"984,omitempty"`
}

func (m *UnderlyingAmount) SetNoUnderlyingAmounts(v []NoUnderlyingAmounts) { m.NoUnderlyingAmounts = v }
