//Package allocation msg type = J.
package allocation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoOrders is a repeating group in Allocation
type NoOrders struct {
	//ClOrdID is a non-required field for NoOrders.
	ClOrdID *string `fix:"11"`
	//OrderID is a non-required field for NoOrders.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for NoOrders.
	SecondaryOrderID *string `fix:"198"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
	//WaveNo is a non-required field for NoOrders.
	WaveNo *string `fix:"105"`
}

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }
func (m *NoOrders) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *NoOrders) SetListID(v string)           { m.ListID = &v }
func (m *NoOrders) SetWaveNo(v string)           { m.WaveNo = &v }

//NoExecs is a repeating group in Allocation
type NoExecs struct {
	//LastShares is a non-required field for NoExecs.
	LastShares *float64 `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
}

func (m *NoExecs) SetLastShares(v float64)  { m.LastShares = &v }
func (m *NoExecs) SetExecID(v string)       { m.ExecID = &v }
func (m *NoExecs) SetLastPx(v float64)      { m.LastPx = &v }
func (m *NoExecs) SetLastCapacity(v string) { m.LastCapacity = &v }

//NoAllocs is a repeating group in Allocation
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocShares is a required field for NoAllocs.
	AllocShares float64 `fix:"80"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//BrokerOfCredit is a non-required field for NoAllocs.
	BrokerOfCredit *string `fix:"92"`
	//NotifyBrokerOfCredit is a non-required field for NoAllocs.
	NotifyBrokerOfCredit *bool `fix:"208"`
	//AllocHandlInst is a non-required field for NoAllocs.
	AllocHandlInst *int `fix:"209"`
	//AllocText is a non-required field for NoAllocs.
	AllocText *string `fix:"161"`
	//EncodedAllocTextLen is a non-required field for NoAllocs.
	EncodedAllocTextLen *int `fix:"360"`
	//EncodedAllocText is a non-required field for NoAllocs.
	EncodedAllocText *string `fix:"361"`
	//ExecBroker is a non-required field for NoAllocs.
	ExecBroker *string `fix:"76"`
	//ClientID is a non-required field for NoAllocs.
	ClientID *string `fix:"109"`
	//Commission is a non-required field for NoAllocs.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NoAllocs.
	CommType *string `fix:"13"`
	//AllocAvgPx is a non-required field for NoAllocs.
	AllocAvgPx *float64 `fix:"153"`
	//AllocNetMoney is a non-required field for NoAllocs.
	AllocNetMoney *float64 `fix:"154"`
	//SettlCurrAmt is a non-required field for NoAllocs.
	SettlCurrAmt *float64 `fix:"119"`
	//SettlCurrency is a non-required field for NoAllocs.
	SettlCurrency *string `fix:"120"`
	//SettlCurrFxRate is a non-required field for NoAllocs.
	SettlCurrFxRate *float64 `fix:"155"`
	//SettlCurrFxRateCalc is a non-required field for NoAllocs.
	SettlCurrFxRateCalc *string `fix:"156"`
	//AccruedInterestAmt is a non-required field for NoAllocs.
	AccruedInterestAmt *float64 `fix:"159"`
	//SettlInstMode is a non-required field for NoAllocs.
	SettlInstMode *string `fix:"160"`
	//NoMiscFees is a non-required field for NoAllocs.
	NoMiscFees []NoMiscFees `fix:"136,omitempty"`
}

func (m *NoAllocs) SetAllocAccount(v string)        { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocPrice(v float64)         { m.AllocPrice = &v }
func (m *NoAllocs) SetAllocShares(v float64)        { m.AllocShares = v }
func (m *NoAllocs) SetProcessCode(v string)         { m.ProcessCode = &v }
func (m *NoAllocs) SetBrokerOfCredit(v string)      { m.BrokerOfCredit = &v }
func (m *NoAllocs) SetNotifyBrokerOfCredit(v bool)  { m.NotifyBrokerOfCredit = &v }
func (m *NoAllocs) SetAllocHandlInst(v int)         { m.AllocHandlInst = &v }
func (m *NoAllocs) SetAllocText(v string)           { m.AllocText = &v }
func (m *NoAllocs) SetEncodedAllocTextLen(v int)    { m.EncodedAllocTextLen = &v }
func (m *NoAllocs) SetEncodedAllocText(v string)    { m.EncodedAllocText = &v }
func (m *NoAllocs) SetExecBroker(v string)          { m.ExecBroker = &v }
func (m *NoAllocs) SetClientID(v string)            { m.ClientID = &v }
func (m *NoAllocs) SetCommission(v float64)         { m.Commission = &v }
func (m *NoAllocs) SetCommType(v string)            { m.CommType = &v }
func (m *NoAllocs) SetAllocAvgPx(v float64)         { m.AllocAvgPx = &v }
func (m *NoAllocs) SetAllocNetMoney(v float64)      { m.AllocNetMoney = &v }
func (m *NoAllocs) SetSettlCurrAmt(v float64)       { m.SettlCurrAmt = &v }
func (m *NoAllocs) SetSettlCurrency(v string)       { m.SettlCurrency = &v }
func (m *NoAllocs) SetSettlCurrFxRate(v float64)    { m.SettlCurrFxRate = &v }
func (m *NoAllocs) SetSettlCurrFxRateCalc(v string) { m.SettlCurrFxRateCalc = &v }
func (m *NoAllocs) SetAccruedInterestAmt(v float64) { m.AccruedInterestAmt = &v }
func (m *NoAllocs) SetSettlInstMode(v string)       { m.SettlInstMode = &v }
func (m *NoAllocs) SetNoMiscFees(v []NoMiscFees)    { m.NoMiscFees = v }

//NoMiscFees is a repeating group in NoAllocs
type NoMiscFees struct {
	//MiscFeeAmt is a non-required field for NoMiscFees.
	MiscFeeAmt *float64 `fix:"137"`
	//MiscFeeCurr is a non-required field for NoMiscFees.
	MiscFeeCurr *string `fix:"138"`
	//MiscFeeType is a non-required field for NoMiscFees.
	MiscFeeType *string `fix:"139"`
}

func (m *NoMiscFees) SetMiscFeeAmt(v float64) { m.MiscFeeAmt = &v }
func (m *NoMiscFees) SetMiscFeeCurr(v string) { m.MiscFeeCurr = &v }
func (m *NoMiscFees) SetMiscFeeType(v string) { m.MiscFeeType = &v }

//Message is a Allocation FIX Message
type Message struct {
	FIXMsgType string `fix:"J"`
	fix42.Header
	//AllocID is a required field for Allocation.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for Allocation.
	AllocTransType string `fix:"71"`
	//RefAllocID is a non-required field for Allocation.
	RefAllocID *string `fix:"72"`
	//AllocLinkID is a non-required field for Allocation.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for Allocation.
	AllocLinkType *int `fix:"197"`
	//NoOrders is a non-required field for Allocation.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//NoExecs is a non-required field for Allocation.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//Side is a required field for Allocation.
	Side string `fix:"54"`
	//Symbol is a required field for Allocation.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for Allocation.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for Allocation.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for Allocation.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for Allocation.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for Allocation.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for Allocation.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for Allocation.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for Allocation.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for Allocation.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for Allocation.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for Allocation.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for Allocation.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for Allocation.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for Allocation.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for Allocation.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for Allocation.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for Allocation.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for Allocation.
	EncodedSecurityDesc *string `fix:"351"`
	//Shares is a required field for Allocation.
	Shares float64 `fix:"53"`
	//LastMkt is a non-required field for Allocation.
	LastMkt *string `fix:"30"`
	//TradingSessionID is a non-required field for Allocation.
	TradingSessionID *string `fix:"336"`
	//AvgPx is a required field for Allocation.
	AvgPx float64 `fix:"6"`
	//Currency is a non-required field for Allocation.
	Currency *string `fix:"15"`
	//AvgPrxPrecision is a non-required field for Allocation.
	AvgPrxPrecision *int `fix:"74"`
	//TradeDate is a required field for Allocation.
	TradeDate string `fix:"75"`
	//TransactTime is a non-required field for Allocation.
	TransactTime *time.Time `fix:"60"`
	//SettlmntTyp is a non-required field for Allocation.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for Allocation.
	FutSettDate *string `fix:"64"`
	//GrossTradeAmt is a non-required field for Allocation.
	GrossTradeAmt *float64 `fix:"381"`
	//NetMoney is a non-required field for Allocation.
	NetMoney *float64 `fix:"118"`
	//OpenClose is a non-required field for Allocation.
	OpenClose *string `fix:"77"`
	//Text is a non-required field for Allocation.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for Allocation.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for Allocation.
	EncodedText *string `fix:"355"`
	//NumDaysInterest is a non-required field for Allocation.
	NumDaysInterest *int `fix:"157"`
	//AccruedInterestRate is a non-required field for Allocation.
	AccruedInterestRate *float64 `fix:"158"`
	//NoAllocs is a non-required field for Allocation.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAllocID(v string)              { m.AllocID = v }
func (m *Message) SetAllocTransType(v string)       { m.AllocTransType = v }
func (m *Message) SetRefAllocID(v string)           { m.RefAllocID = &v }
func (m *Message) SetAllocLinkID(v string)          { m.AllocLinkID = &v }
func (m *Message) SetAllocLinkType(v int)           { m.AllocLinkType = &v }
func (m *Message) SetNoOrders(v []NoOrders)         { m.NoOrders = v }
func (m *Message) SetNoExecs(v []NoExecs)           { m.NoExecs = v }
func (m *Message) SetSide(v string)                 { m.Side = v }
func (m *Message) SetSymbol(v string)               { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)            { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)           { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)             { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)         { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string)    { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)             { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)               { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)         { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)         { m.OptAttribute = &v }
func (m *Message) SetContractMultiplier(v float64)  { m.ContractMultiplier = &v }
func (m *Message) SetCouponRate(v float64)          { m.CouponRate = &v }
func (m *Message) SetSecurityExchange(v string)     { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)               { m.Issuer = &v }
func (m *Message) SetEncodedIssuerLen(v int)        { m.EncodedIssuerLen = &v }
func (m *Message) SetEncodedIssuer(v string)        { m.EncodedIssuer = &v }
func (m *Message) SetSecurityDesc(v string)         { m.SecurityDesc = &v }
func (m *Message) SetEncodedSecurityDescLen(v int)  { m.EncodedSecurityDescLen = &v }
func (m *Message) SetEncodedSecurityDesc(v string)  { m.EncodedSecurityDesc = &v }
func (m *Message) SetShares(v float64)              { m.Shares = v }
func (m *Message) SetLastMkt(v string)              { m.LastMkt = &v }
func (m *Message) SetTradingSessionID(v string)     { m.TradingSessionID = &v }
func (m *Message) SetAvgPx(v float64)               { m.AvgPx = v }
func (m *Message) SetCurrency(v string)             { m.Currency = &v }
func (m *Message) SetAvgPrxPrecision(v int)         { m.AvgPrxPrecision = &v }
func (m *Message) SetTradeDate(v string)            { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)      { m.TransactTime = &v }
func (m *Message) SetSettlmntTyp(v string)          { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)          { m.FutSettDate = &v }
func (m *Message) SetGrossTradeAmt(v float64)       { m.GrossTradeAmt = &v }
func (m *Message) SetNetMoney(v float64)            { m.NetMoney = &v }
func (m *Message) SetOpenClose(v string)            { m.OpenClose = &v }
func (m *Message) SetText(v string)                 { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)          { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)          { m.EncodedText = &v }
func (m *Message) SetNumDaysInterest(v int)         { m.NumDaysInterest = &v }
func (m *Message) SetAccruedInterestRate(v float64) { m.AccruedInterestRate = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)         { m.NoAllocs = v }

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
	return enum.BeginStringFIX42, "J", r
}
