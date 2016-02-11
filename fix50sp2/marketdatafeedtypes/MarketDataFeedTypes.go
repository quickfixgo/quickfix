package marketdatafeedtypes

//NoMDFeedTypes is a repeating group in MarketDataFeedTypes
type NoMDFeedTypes struct {
	//MDFeedType is a non-required field for NoMDFeedTypes.
	MDFeedType *string `fix:"1022"`
	//MarketDepth is a non-required field for NoMDFeedTypes.
	MarketDepth *int `fix:"264"`
	//MDBookType is a non-required field for NoMDFeedTypes.
	MDBookType *int `fix:"1021"`
}

//Component is a fix50sp2 MarketDataFeedTypes Component
type Component struct {
	//NoMDFeedTypes is a non-required field for MarketDataFeedTypes.
	NoMDFeedTypes []NoMDFeedTypes `fix:"1141,omitempty"`
}

func New() *Component { return new(Component) }
