package orderqtydata

//Component is a fix50 OrderQtyData Component
type Component struct {
	//OrderQty is a non-required field for OrderQtyData.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for OrderQtyData.
	CashOrderQty *float64 `fix:"152"`
	//OrderPercent is a non-required field for OrderQtyData.
	OrderPercent *float64 `fix:"516"`
	//RoundingDirection is a non-required field for OrderQtyData.
	RoundingDirection *string `fix:"468"`
	//RoundingModulus is a non-required field for OrderQtyData.
	RoundingModulus *float64 `fix:"469"`
}

func New() *Component { return new(Component) }
