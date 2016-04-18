package execallocgrp

//New returns an initialized ExecAllocGrp instance
func New() *ExecAllocGrp {
	var m ExecAllocGrp
	return &m
}

//NoExecs is a repeating group in ExecAllocGrp
type NoExecs struct {
	//LastQty is a non-required field for NoExecs.
	LastQty *float64 `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for NoExecs.
	SecondaryExecID *string `fix:"527"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastParPx is a non-required field for NoExecs.
	LastParPx *float64 `fix:"669"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
	//TradeID is a non-required field for NoExecs.
	TradeID *string `fix:"1003"`
	//FirmTradeID is a non-required field for NoExecs.
	FirmTradeID *string `fix:"1041"`
}

//NewNoExecs returns an initialized NoExecs instance
func NewNoExecs() *NoExecs {
	var m NoExecs
	return &m
}

func (m *NoExecs) SetLastQty(v float64)        { m.LastQty = &v }
func (m *NoExecs) SetExecID(v string)          { m.ExecID = &v }
func (m *NoExecs) SetSecondaryExecID(v string) { m.SecondaryExecID = &v }
func (m *NoExecs) SetLastPx(v float64)         { m.LastPx = &v }
func (m *NoExecs) SetLastParPx(v float64)      { m.LastParPx = &v }
func (m *NoExecs) SetLastCapacity(v string)    { m.LastCapacity = &v }
func (m *NoExecs) SetTradeID(v string)         { m.TradeID = &v }
func (m *NoExecs) SetFirmTradeID(v string)     { m.FirmTradeID = &v }

//ExecAllocGrp is a fix50sp2 Component
type ExecAllocGrp struct {
	//NoExecs is a non-required field for ExecAllocGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func (m *ExecAllocGrp) SetNoExecs(v []NoExecs) { m.NoExecs = v }
