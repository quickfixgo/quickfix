package expirationqty

//NoExpiration is a repeating group in ExpirationQty
type NoExpiration struct {
	//ExpirationQtyType is a non-required field for NoExpiration.
	ExpirationQtyType *int `fix:"982"`
	//ExpQty is a non-required field for NoExpiration.
	ExpQty *float64 `fix:"983"`
}

//Component is a fix50sp1 ExpirationQty Component
type Component struct {
	//NoExpiration is a non-required field for ExpirationQty.
	NoExpiration []NoExpiration `fix:"981,omitempty"`
}

func New() *Component { return new(Component) }
