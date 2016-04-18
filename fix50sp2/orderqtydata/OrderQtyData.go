package orderqtydata

//New returns an initialized OrderQtyData instance
func New() *OrderQtyData {
	var m OrderQtyData
	return &m
}

//OrderQtyData is a fix50sp2 Component
type OrderQtyData struct {
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

func (m *OrderQtyData) SetOrderQty(v float64)         { m.OrderQty = &v }
func (m *OrderQtyData) SetCashOrderQty(v float64)     { m.CashOrderQty = &v }
func (m *OrderQtyData) SetOrderPercent(v float64)     { m.OrderPercent = &v }
func (m *OrderQtyData) SetRoundingDirection(v string) { m.RoundingDirection = &v }
func (m *OrderQtyData) SetRoundingModulus(v float64)  { m.RoundingModulus = &v }
