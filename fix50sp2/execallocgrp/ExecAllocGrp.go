package execallocgrp

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

//Component is a fix50sp2 ExecAllocGrp Component
type Component struct {
	//NoExecs is a non-required field for ExecAllocGrp.
	NoExecs []NoExecs `fix:"124,omitempty"`
}

func New() *Component { return new(Component) }
