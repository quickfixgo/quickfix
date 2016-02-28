package rgstdistinstgrp

//NoDistribInsts is a repeating group in RgstDistInstGrp
type NoDistribInsts struct {
	//DistribPaymentMethod is a non-required field for NoDistribInsts.
	DistribPaymentMethod *int `fix:"477"`
	//DistribPercentage is a non-required field for NoDistribInsts.
	DistribPercentage *float64 `fix:"512"`
	//CashDistribCurr is a non-required field for NoDistribInsts.
	CashDistribCurr *string `fix:"478"`
	//CashDistribAgentName is a non-required field for NoDistribInsts.
	CashDistribAgentName *string `fix:"498"`
	//CashDistribAgentCode is a non-required field for NoDistribInsts.
	CashDistribAgentCode *string `fix:"499"`
	//CashDistribAgentAcctNumber is a non-required field for NoDistribInsts.
	CashDistribAgentAcctNumber *string `fix:"500"`
	//CashDistribPayRef is a non-required field for NoDistribInsts.
	CashDistribPayRef *string `fix:"501"`
	//CashDistribAgentAcctName is a non-required field for NoDistribInsts.
	CashDistribAgentAcctName *string `fix:"502"`
}

//Component is a fix50sp1 RgstDistInstGrp Component
type Component struct {
	//NoDistribInsts is a non-required field for RgstDistInstGrp.
	NoDistribInsts []NoDistribInsts `fix:"510,omitempty"`
}

func New() *Component { return new(Component) }

func (m *Component) SetNoDistribInsts(v []NoDistribInsts) { m.NoDistribInsts = v }
