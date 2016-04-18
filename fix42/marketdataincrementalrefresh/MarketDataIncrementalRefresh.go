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

//NewNoMDEntries returns an initialized NoMDEntries instance
func NewNoMDEntries(mdupdateaction string) *NoMDEntries {
	var m NoMDEntries
	m.SetMDUpdateAction(mdupdateaction)
	return &m
}

func (m *NoMDEntries) SetMDUpdateAction(v string)      { m.MDUpdateAction = v }
func (m *NoMDEntries) SetDeleteReason(v string)        { m.DeleteReason = &v }
func (m *NoMDEntries) SetMDEntryType(v string)         { m.MDEntryType = &v }
func (m *NoMDEntries) SetMDEntryID(v string)           { m.MDEntryID = &v }
func (m *NoMDEntries) SetMDEntryRefID(v string)        { m.MDEntryRefID = &v }
func (m *NoMDEntries) SetSymbol(v string)              { m.Symbol = &v }
func (m *NoMDEntries) SetSymbolSfx(v string)           { m.SymbolSfx = &v }
func (m *NoMDEntries) SetSecurityID(v string)          { m.SecurityID = &v }
func (m *NoMDEntries) SetIDSource(v string)            { m.IDSource = &v }
func (m *NoMDEntries) SetSecurityType(v string)        { m.SecurityType = &v }
func (m *NoMDEntries) SetMaturityMonthYear(v string)   { m.MaturityMonthYear = &v }
func (m *NoMDEntries) SetMaturityDay(v int)            { m.MaturityDay = &v }
func (m *NoMDEntries) SetPutOrCall(v int)              { m.PutOrCall = &v }
func (m *NoMDEntries) SetStrikePrice(v float64)        { m.StrikePrice = &v }
func (m *NoMDEntries) SetOptAttribute(v string)        { m.OptAttribute = &v }
func (m *NoMDEntries) SetContractMultiplier(v float64) { m.ContractMultiplier = &v }
func (m *NoMDEntries) SetCouponRate(v float64)         { m.CouponRate = &v }
func (m *NoMDEntries) SetSecurityExchange(v string)    { m.SecurityExchange = &v }
func (m *NoMDEntries) SetIssuer(v string)              { m.Issuer = &v }
func (m *NoMDEntries) SetEncodedIssuerLen(v int)       { m.EncodedIssuerLen = &v }
func (m *NoMDEntries) SetEncodedIssuer(v string)       { m.EncodedIssuer = &v }
func (m *NoMDEntries) SetSecurityDesc(v string)        { m.SecurityDesc = &v }
func (m *NoMDEntries) SetEncodedSecurityDescLen(v int) { m.EncodedSecurityDescLen = &v }
func (m *NoMDEntries) SetEncodedSecurityDesc(v string) { m.EncodedSecurityDesc = &v }
func (m *NoMDEntries) SetFinancialStatus(v string)     { m.FinancialStatus = &v }
func (m *NoMDEntries) SetCorporateAction(v string)     { m.CorporateAction = &v }
func (m *NoMDEntries) SetMDEntryPx(v float64)          { m.MDEntryPx = &v }
func (m *NoMDEntries) SetCurrency(v string)            { m.Currency = &v }
func (m *NoMDEntries) SetMDEntrySize(v float64)        { m.MDEntrySize = &v }
func (m *NoMDEntries) SetMDEntryDate(v string)         { m.MDEntryDate = &v }
func (m *NoMDEntries) SetMDEntryTime(v string)         { m.MDEntryTime = &v }
func (m *NoMDEntries) SetTickDirection(v string)       { m.TickDirection = &v }
func (m *NoMDEntries) SetMDMkt(v string)               { m.MDMkt = &v }
func (m *NoMDEntries) SetTradingSessionID(v string)    { m.TradingSessionID = &v }
func (m *NoMDEntries) SetQuoteCondition(v string)      { m.QuoteCondition = &v }
func (m *NoMDEntries) SetTradeCondition(v string)      { m.TradeCondition = &v }
func (m *NoMDEntries) SetMDEntryOriginator(v string)   { m.MDEntryOriginator = &v }
func (m *NoMDEntries) SetLocationID(v string)          { m.LocationID = &v }
func (m *NoMDEntries) SetDeskID(v string)              { m.DeskID = &v }
func (m *NoMDEntries) SetOpenCloseSettleFlag(v string) { m.OpenCloseSettleFlag = &v }
func (m *NoMDEntries) SetTimeInForce(v string)         { m.TimeInForce = &v }
func (m *NoMDEntries) SetExpireDate(v string)          { m.ExpireDate = &v }
func (m *NoMDEntries) SetExpireTime(v time.Time)       { m.ExpireTime = &v }
func (m *NoMDEntries) SetMinQty(v float64)             { m.MinQty = &v }
func (m *NoMDEntries) SetExecInst(v string)            { m.ExecInst = &v }
func (m *NoMDEntries) SetSellerDays(v int)             { m.SellerDays = &v }
func (m *NoMDEntries) SetOrderID(v string)             { m.OrderID = &v }
func (m *NoMDEntries) SetQuoteEntryID(v string)        { m.QuoteEntryID = &v }
func (m *NoMDEntries) SetMDEntryBuyer(v string)        { m.MDEntryBuyer = &v }
func (m *NoMDEntries) SetMDEntrySeller(v string)       { m.MDEntrySeller = &v }
func (m *NoMDEntries) SetNumberOfOrders(v int)         { m.NumberOfOrders = &v }
func (m *NoMDEntries) SetMDEntryPositionNo(v int)      { m.MDEntryPositionNo = &v }
func (m *NoMDEntries) SetTotalVolumeTraded(v float64)  { m.TotalVolumeTraded = &v }
func (m *NoMDEntries) SetText(v string)                { m.Text = &v }
func (m *NoMDEntries) SetEncodedTextLen(v int)         { m.EncodedTextLen = &v }
func (m *NoMDEntries) SetEncodedText(v string)         { m.EncodedText = &v }

//Message is a MarketDataIncrementalRefresh FIX Message
type Message struct {
	FIXMsgType string `fix:"X"`
	fix42.Header
	//MDReqID is a non-required field for MarketDataIncrementalRefresh.
	MDReqID *string `fix:"262"`
	//NoMDEntries is a required field for MarketDataIncrementalRefresh.
	NoMDEntries []NoMDEntries `fix:"268"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized MarketDataIncrementalRefresh instance
func New(nomdentries []NoMDEntries) *Message {
	var m Message
	m.SetNoMDEntries(nomdentries)
	return &m
}

func (m *Message) SetMDReqID(v string)            { m.MDReqID = &v }
func (m *Message) SetNoMDEntries(v []NoMDEntries) { m.NoMDEntries = v }

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
