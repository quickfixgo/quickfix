package rgstdistinstgrp

//New returns an initialized RgstDistInstGrp instance
func New() *RgstDistInstGrp {
	var m RgstDistInstGrp
	return &m
}

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

//NewNoDistribInsts returns an initialized NoDistribInsts instance
func NewNoDistribInsts() *NoDistribInsts {
	var m NoDistribInsts
	return &m
}

func (m *NoDistribInsts) SetDistribPaymentMethod(v int)          { m.DistribPaymentMethod = &v }
func (m *NoDistribInsts) SetDistribPercentage(v float64)         { m.DistribPercentage = &v }
func (m *NoDistribInsts) SetCashDistribCurr(v string)            { m.CashDistribCurr = &v }
func (m *NoDistribInsts) SetCashDistribAgentName(v string)       { m.CashDistribAgentName = &v }
func (m *NoDistribInsts) SetCashDistribAgentCode(v string)       { m.CashDistribAgentCode = &v }
func (m *NoDistribInsts) SetCashDistribAgentAcctNumber(v string) { m.CashDistribAgentAcctNumber = &v }
func (m *NoDistribInsts) SetCashDistribPayRef(v string)          { m.CashDistribPayRef = &v }
func (m *NoDistribInsts) SetCashDistribAgentAcctName(v string)   { m.CashDistribAgentAcctName = &v }

//RgstDistInstGrp is a fix50 Component
type RgstDistInstGrp struct {
	//NoDistribInsts is a non-required field for RgstDistInstGrp.
	NoDistribInsts []NoDistribInsts `fix:"510,omitempty"`
}

func (m *RgstDistInstGrp) SetNoDistribInsts(v []NoDistribInsts) { m.NoDistribInsts = v }
