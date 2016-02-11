//Package marketdatasnapshotfullrefresh msg type = W.
package marketdatasnapshotfullrefresh

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoMDEntries is a repeating group in MarketDataSnapshotFullRefresh
type NoMDEntries struct {
	//MDEntryType is a required field for NoMDEntries.
	MDEntryType string `fix:"269"`
	//MDEntryPx is a required field for NoMDEntries.
	MDEntryPx float64 `fix:"270"`
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
	//Text is a non-required field for NoMDEntries.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NoMDEntries.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NoMDEntries.
	EncodedText *string `fix:"355"`
}

//Message is a MarketDataSnapshotFullRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"W"`
	Header     fix42.Header
	//MDReqID is a non-required field for MarketDataSnapshotFullRefresh.
	MDReqID *string `fix:"262"`
	//Symbol is a required field for MarketDataSnapshotFullRefresh.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for MarketDataSnapshotFullRefresh.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for MarketDataSnapshotFullRefresh.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for MarketDataSnapshotFullRefresh.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for MarketDataSnapshotFullRefresh.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for MarketDataSnapshotFullRefresh.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for MarketDataSnapshotFullRefresh.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for MarketDataSnapshotFullRefresh.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for MarketDataSnapshotFullRefresh.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for MarketDataSnapshotFullRefresh.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for MarketDataSnapshotFullRefresh.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for MarketDataSnapshotFullRefresh.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for MarketDataSnapshotFullRefresh.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for MarketDataSnapshotFullRefresh.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for MarketDataSnapshotFullRefresh.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for MarketDataSnapshotFullRefresh.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for MarketDataSnapshotFullRefresh.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for MarketDataSnapshotFullRefresh.
	EncodedSecurityDesc *string `fix:"351"`
	//FinancialStatus is a non-required field for MarketDataSnapshotFullRefresh.
	FinancialStatus *string `fix:"291"`
	//CorporateAction is a non-required field for MarketDataSnapshotFullRefresh.
	CorporateAction *string `fix:"292"`
	//TotalVolumeTraded is a non-required field for MarketDataSnapshotFullRefresh.
	TotalVolumeTraded *float64 `fix:"387"`
	//NoMDEntries is a required field for MarketDataSnapshotFullRefresh.
	NoMDEntries []NoMDEntries `fix:"268"`
	Trailer     fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Message, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Mesage type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		m := new(Message)
		if err := quickfix.Unmarshal(msg, m); err != nil {
			return err
		}
		return router(*m, sessionID)
	}
	return enum.BeginStringFIX42, "W", r
}
