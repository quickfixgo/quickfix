package tradingsessionrules

import (
	"github.com/quickfixgo/quickfix/fix50sp1/execinstrules"
	"github.com/quickfixgo/quickfix/fix50sp1/marketdatafeedtypes"
	"github.com/quickfixgo/quickfix/fix50sp1/matchrules"
	"github.com/quickfixgo/quickfix/fix50sp1/ordtyperules"
	"github.com/quickfixgo/quickfix/fix50sp1/timeinforcerules"
)

//TradingSessionRules is a fix50sp1 Component
type TradingSessionRules struct {
	//OrdTypeRules Component
	ordtyperules.OrdTypeRules
	//TimeInForceRules Component
	timeinforcerules.TimeInForceRules
	//ExecInstRules Component
	execinstrules.ExecInstRules
	//MatchRules Component
	matchrules.MatchRules
	//MarketDataFeedTypes Component
	marketdatafeedtypes.MarketDataFeedTypes
}
