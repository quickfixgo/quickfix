package marketdatafeedtypes

//New returns an initialized MarketDataFeedTypes instance
func New() *MarketDataFeedTypes {
	var m MarketDataFeedTypes
	return &m
}

//NoMDFeedTypes is a repeating group in MarketDataFeedTypes
type NoMDFeedTypes struct {
	//MDFeedType is a non-required field for NoMDFeedTypes.
	MDFeedType *string `fix:"1022"`
	//MarketDepth is a non-required field for NoMDFeedTypes.
	MarketDepth *int `fix:"264"`
	//MDBookType is a non-required field for NoMDFeedTypes.
	MDBookType *int `fix:"1021"`
}

//NewNoMDFeedTypes returns an initialized NoMDFeedTypes instance
func NewNoMDFeedTypes() *NoMDFeedTypes {
	var m NoMDFeedTypes
	return &m
}

func (m *NoMDFeedTypes) SetMDFeedType(v string) { m.MDFeedType = &v }
func (m *NoMDFeedTypes) SetMarketDepth(v int)   { m.MarketDepth = &v }
func (m *NoMDFeedTypes) SetMDBookType(v int)    { m.MDBookType = &v }

//MarketDataFeedTypes is a fix50sp2 Component
type MarketDataFeedTypes struct {
	//NoMDFeedTypes is a non-required field for MarketDataFeedTypes.
	NoMDFeedTypes []NoMDFeedTypes `fix:"1141,omitempty"`
}

func (m *MarketDataFeedTypes) SetNoMDFeedTypes(v []NoMDFeedTypes) { m.NoMDFeedTypes = v }
