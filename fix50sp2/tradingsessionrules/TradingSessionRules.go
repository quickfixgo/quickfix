package tradingsessionrules

import (
	"github.com/quickfixgo/quickfix/fix50sp2/execinstrules"
	"github.com/quickfixgo/quickfix/fix50sp2/marketdatafeedtypes"
	"github.com/quickfixgo/quickfix/fix50sp2/matchrules"
	"github.com/quickfixgo/quickfix/fix50sp2/ordtyperules"
	"github.com/quickfixgo/quickfix/fix50sp2/timeinforcerules"
)

//TradingSessionRules is a fix50sp2 Component
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
