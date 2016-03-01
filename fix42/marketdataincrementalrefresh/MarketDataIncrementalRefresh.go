//Package marketdataincrementalrefresh msg type = X.
package marketdataincrementalrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoMDEntries is a repeating group in MarketDataIncrementalRefresh
type NoMDEntries struct {
	//MDUpdateAction is a required field for NoMDEntries.
	MDUpdateAction string `fix:"279"`
	//DeleteReason is a non-required field for NoMDEntries.
	DeleteReason *string `fix:"285"`
	//MDEntryType is a non-required field for NoMDEntries.
	MDEntryType *string `fix:"269"`
	//MDEntryID is a non-required field for NoMDEntries.
	MDEntryID *string `fix:"278"`
	//MDEntryRefID is a non-required field for NoMDEntries.
	MDEntryRefID *string `fix:"280"`
	//Symbol is a non-required field for NoMDEntries.
	Symbol *string `fix:"55"`
	//SymbolSfx is a non-required field for NoMDEntries.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NoMDEntries.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NoMDEntries.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NoMDEntries.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NoMDEntries.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NoMDEntries.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NoMDEntries.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NoMDEntries.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NoMDEntries.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NoMDEntries.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NoMDEntries.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NoMDEntries.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NoMDEntries.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NoMDEntries.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NoMDEntries.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NoMDEntries.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NoMDEntries.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NoMDEntries.
	EncodedSecurityDesc *string `fix:"351"`
	//FinancialStatus is a non-required field for NoMDEntries.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for NoMDEntries.
	CorporateAction *string `fix:"292"`
	//MDEntryPx is a non-required field for NoMDEntries.
	MDEntryPx *float64 `fix:"270"`
	//Currency is a non-required field for NoMDEntries.
	Currency *string `fix:"15"`
	//MDEntrySize is a non-required field for NoMDEntries.
	MDEntrySize *float64 `fix:"271"`
	//MDEntryDate is a non-required field for NoMDEntries.
	MDEntryDate *string `fix:"272"`
	//MDEntryTime is a non-required field for NoMDEntries.
	MDEntryTime *string `fix:"273"`
	//TickDirection is a non-required field for NoMDEntries.
	TickDirection *string `fix:"274"`
	//MDMkt is a non-required field for NoMDEntries.
	MDMkt *string `fix:"275"`
	//TradingSessionID is a non-required field for NoMDEntries.
	TradingSessionID *string `fix:"336"`
	//QuoteCondition is a non-required field for NoMDEntries.
	QuoteCondition *string `fix:"276"`
	//TradeCondition is a non-required field for NoMDEntries.
	TradeCondition *string `fix:"277"`
	//MDEntryOriginator is a non-required field for NoMDEntries.
	MDEntryOriginator *string `fix:"282"`
	//LocationID is a non-required field for NoMDEntries.
	LocationID *string `fix:"283"`
	//DeskID is a non-required field for NoMDEntries.
	DeskID *string `fix:"284"`
	//OpenCloseSettleFlag is a non-required field for NoMDEntries.
	OpenCloseSettleFlag *string `fix:"286"`
	//TimeInForce is a non-required field for NoMDEntries.
	TimeInForce *string `fix:"59"`
	//ExpireDate is a non-required field for NoMDEntries.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NoMDEntries.
	ExpireTime *time.Time `fix:"126"`
	//MinQty is a non-required field for NoMDEntries.
	MinQty *float64 `fix:"110"`
	//ExecInst is a non-required field for NoMDEntries.
	ExecInst *string `fix:"18"`
	//SellerDays is a non-required field for NoMDEntries.
	SellerDays *int `fix:"287"`
	//OrderID is a non-required field for NoMDEntries.
	OrderID *string `fix:"37"`
	//QuoteEntryID is a non-required field for NoMDEntries.
	QuoteEntryID *string `fix:"299"`
	//MDEntryBuyer is a non-required field for NoMDEntries.
	MDEntryBuyer *string `fix:"288"`
	//MDEntrySeller is a non-required field for NoMDEntries.
	MDEntrySeller *string `fix:"289"`
	//NumberOfOrders is a non-required field for NoMDEntries.
	NumberOfOrders *int `fix:"346"`
	//MDEntryPositionNo is a non-required field for NoMDEntries.
	MDEntryPositionNo *int `fix:"290"`
	//TotalVolumeTraded is a non-required field for NoMDEntries.
	TotalVolumeTraded *float64 `fix:"387"`
	//Text is a non-required field for NoMDEntries.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoMDEntries.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoMDEntries.
	EncodedText *string `fix:"355"`
}

//Message is a MarketDataIncrementalRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"X"`
	Header     fix42.Header
	//MDReqID is a non-required field for MarketDataIncrementalRefresh.
	MDReqID *string `fix:"262"`
	//NoMDEntries is a required field for MarketDataIncrementalRefresh.
	NoMDEntries []NoMDEntries `fix:"268"`
	Trailer     fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "X", r
}
