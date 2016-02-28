package orderqtydata

//Component is a fix44 OrderQtyData Component
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

func (m *Component) SetOrderQty(v float64)         { m.OrderQty = &v }
func (m *Component) SetCashOrderQty(v float64)     { m.CashOrderQty = &v }
func (m *Component) SetOrderPercent(v float64)     { m.OrderPercent = &v }
func (m *Component) SetRoundingDirection(v string) { m.RoundingDirection = &v }
func (m *Component) SetRoundingModulus(v float64)  { m.RoundingModulus = &v }
