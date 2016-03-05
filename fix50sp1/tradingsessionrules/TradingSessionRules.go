package tradingsessionrules

//NoOrdTypeRules is a repeating group in TradingSessionRules
type NoOrdTypeRules struct {
	//OrdType is a non-required field for NoOrdTypeRules.
	OrdType *string `fix:"40"`
}

//NoTimeInForceRules is a repeating group in TradingSessionRules
type NoTimeInForceRules struct {
	//TimeInForce is a non-required field for NoTimeInForceRules.
	TimeInForce *string `fix:"59"`
}

//NoExecInstRules is a repeating group in TradingSessionRules
type NoExecInstRules struct {
	//ExecInstValue is a non-required field for NoExecInstRules.
	ExecInstValue *string `fix:"1308"`
}

//NoMatchRules is a repeating group in TradingSessionRules
type NoMatchRules struct {
	//MatchAlgorithm is a non-required field for NoMatchRules.
	MatchAlgorithm *string `fix:"1142"`
	//MatchType is a non-required field for NoMatchRules.
	MatchType *string `fix:"574"`
}

//NoMDFeedTypes is a repeating group in TradingSessionRules
type NoMDFeedTypes struct {
	//MDFeedType is a non-required field for NoMDFeedTypes.
	MDFeedType *string `fix:"1022"`
	//MarketDepth is a non-required field for NoMDFeedTypes.
	MarketDepth *int `fix:"264"`
	//MDBookType is a non-required field for NoMDFeedTypes.
	MDBookType *int `fix:"1021"`
}

//TradingSessionRules is a fix50sp1 Component
type TradingSessionRules struct {
	//NoOrdTypeRules is a non-required field for TradingSessionRules.
	NoOrdTypeRules []NoOrdTypeRules `fix:"1237,omitempty"`
	//NoTimeInForceRules is a non-required field for TradingSessionRules.
	NoTimeInForceRules []NoTimeInForceRules `fix:"1239,omitempty"`
	//NoExecInstRules is a non-required field for TradingSessionRules.
	NoExecInstRules []NoExecInstRules `fix:"1232,omitempty"`
	//NoMatchRules is a non-required field for TradingSessionRules.
	NoMatchRules []NoMatchRules `fix:"1235,omitempty"`
	//NoMDFeedTypes is a non-required field for TradingSessionRules.
	NoMDFeedTypes []NoMDFeedTypes `fix:"1141,omitempty"`
}

func (m *TradingSessionRules) SetNoOrdTypeRules(v []NoOrdTypeRules)         { m.NoOrdTypeRules = v }
func (m *TradingSessionRules) SetNoTimeInForceRules(v []NoTimeInForceRules) { m.NoTimeInForceRules = v }
func (m *TradingSessionRules) SetNoExecInstRules(v []NoExecInstRules)       { m.NoExecInstRules = v }
func (m *TradingSessionRules) SetNoMatchRules(v []NoMatchRules)             { m.NoMatchRules = v }
func (m *TradingSessionRules) SetNoMDFeedTypes(v []NoMDFeedTypes)           { m.NoMDFeedTypes = v }
