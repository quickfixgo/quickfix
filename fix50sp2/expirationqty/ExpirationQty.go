package expirationqty

func New() *ExpirationQty {
	var m ExpirationQty
	return &m
}

//NoExpiration is a repeating group in ExpirationQty
type NoExpiration struct {
	//ExpirationQtyType is a non-required field for NoExpiration.
	ExpirationQtyType *int `fix:"982"`
	//ExpQty is a non-required field for NoExpiration.
	ExpQty *float64 `fix:"983"`
}

//ExpirationQty is a fix50sp2 Component
type ExpirationQty struct {
	//NoExpiration is a non-required field for ExpirationQty.
	NoExpiration []NoExpiration `fix:"981,omitempty"`
}

func (m *ExpirationQty) SetNoExpiration(v []NoExpiration) { m.NoExpiration = v }
