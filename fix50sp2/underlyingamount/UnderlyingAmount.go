package underlyingamount

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

//UnderlyingAmount is a fix50sp2 Component
type UnderlyingAmount struct {
	//NoUnderlyingAmounts is a non-required field for UnderlyingAmount.
	NoUnderlyingAmounts []NoUnderlyingAmounts `fix:"984,omitempty"`
}

func (m *UnderlyingAmount) SetNoUnderlyingAmounts(v []NoUnderlyingAmounts) { m.NoUnderlyingAmounts = v }
