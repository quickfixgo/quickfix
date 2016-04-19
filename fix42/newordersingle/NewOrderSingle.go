//Package newordersingle msg type = D.
package newordersingle

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix42"
	"time"
)

//NoAllocs is a repeating group in NewOrderSingle
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocShares is a non-required field for NoAllocs.
	AllocShares *float64 `fix:"80"`
}

//NewNoAllocs returns an initialized NoAllocs instance
func NewNoAllocs() *NoAllocs {
	var m NoAllocs
	return &m
}

func (m *NoAllocs) SetAllocAccount(v string) { m.AllocAccount = &v }
func (m *NoAllocs) SetAllocShares(v float64) { m.AllocShares = &v }

//NoTradingSessions is a repeating group in NewOrderSingle
type NoTradingSessions struct {
	//TradingSessionID is a non-required field for NoTradingSessions.
	TradingSessionID *string `fix:"336"`
}

//NewNoTradingSessions returns an initialized NoTradingSessions instance
func NewNoTradingSessions() *NoTradingSessions {
	var m NoTradingSessions
	return &m
}

func (m *NoTradingSessions) SetTradingSessionID(v string) { m.TradingSessionID = &v }

//Message is a NewOrderSingle FIX Message
type Message struct {
	FIXMsgType string `fix:"D"`
	fix42.Header
	//ClOrdID is a required field for NewOrderSingle.
	ClOrdID string `fix:"11"`
	//ClientID is a non-required field for NewOrderSingle.
	ClientID *string `fix:"109"`
	//ExecBroker is a non-required field for NewOrderSingle.
	ExecBroker *string `fix:"76"`
	//Account is a non-required field for NewOrderSingle.
	Account *string `fix:"1"`
	//NoAllocs is a non-required field for NewOrderSingle.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	//SettlmntTyp is a non-required field for NewOrderSingle.
	SettlmntTyp *string `fix:"63"`
	//FutSettDate is a non-required field for NewOrderSingle.
	FutSettDate *string `fix:"64"`
	//HandlInst is a required field for NewOrderSingle.
	HandlInst string `fix:"21"`
	//ExecInst is a non-required field for NewOrderSingle.
	ExecInst *string `fix:"18"`
	//MinQty is a non-required field for NewOrderSingle.
	MinQty *float64 `fix:"110"`
	//MaxFloor is a non-required field for NewOrderSingle.
	MaxFloor *float64 `fix:"111"`
	//ExDestination is a non-required field for NewOrderSingle.
	ExDestination *string `fix:"100"`
	//NoTradingSessions is a non-required field for NewOrderSingle.
	NoTradingSessions []NoTradingSessions `fix:"386,omitempty"`
	//ProcessCode is a non-required field for NewOrderSingle.
	ProcessCode *string `fix:"81"`
	//Symbol is a required field for NewOrderSingle.
	Symbol string `fix:"55"`
	//SymbolSfx is a non-required field for NewOrderSingle.
	SymbolSfx *string `fix:"65"`
	//SecurityID is a non-required field for NewOrderSingle.
	SecurityID *string `fix:"48"`
	//IDSource is a non-required field for NewOrderSingle.
	IDSource *string `fix:"22"`
	//SecurityType is a non-required field for NewOrderSingle.
	SecurityType *string `fix:"167"`
	//MaturityMonthYear is a non-required field for NewOrderSingle.
	MaturityMonthYear *string `fix:"200"`
	//MaturityDay is a non-required field for NewOrderSingle.
	MaturityDay *int `fix:"205"`
	//PutOrCall is a non-required field for NewOrderSingle.
	PutOrCall *int `fix:"201"`
	//StrikePrice is a non-required field for NewOrderSingle.
	StrikePrice *float64 `fix:"202"`
	//OptAttribute is a non-required field for NewOrderSingle.
	OptAttribute *string `fix:"206"`
	//ContractMultiplier is a non-required field for NewOrderSingle.
	ContractMultiplier *float64 `fix:"231"`
	//CouponRate is a non-required field for NewOrderSingle.
	CouponRate *float64 `fix:"223"`
	//SecurityExchange is a non-required field for NewOrderSingle.
	SecurityExchange *string `fix:"207"`
	//Issuer is a non-required field for NewOrderSingle.
	Issuer *string `fix:"106"`
	//EncodedIssuerLen is a non-required field for NewOrderSingle.
	EncodedIssuerLen *int `fix:"348"`
	//EncodedIssuer is a non-required field for NewOrderSingle.
	EncodedIssuer *string `fix:"349"`
	//SecurityDesc is a non-required field for NewOrderSingle.
	SecurityDesc *string `fix:"107"`
	//EncodedSecurityDescLen is a non-required field for NewOrderSingle.
	EncodedSecurityDescLen *int `fix:"350"`
	//EncodedSecurityDesc is a non-required field for NewOrderSingle.
	EncodedSecurityDesc *string `fix:"351"`
	//PrevClosePx is a non-required field for NewOrderSingle.
	PrevClosePx *float64 `fix:"140"`
	//Side is a required field for NewOrderSingle.
	Side string `fix:"54"`
	//LocateReqd is a non-required field for NewOrderSingle.
	LocateReqd *bool `fix:"114"`
	//TransactTime is a required field for NewOrderSingle.
	TransactTime time.Time `fix:"60"`
	//OrderQty is a non-required field for NewOrderSingle.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for NewOrderSingle.
	CashOrderQty *float64 `fix:"152"`
	//OrdType is a required field for NewOrderSingle.
	OrdType string `fix:"40"`
	//Price is a non-required field for NewOrderSingle.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for NewOrderSingle.
	StopPx *float64 `fix:"99"`
	//Currency is a non-required field for NewOrderSingle.
	Currency *string `fix:"15"`
	//ComplianceID is a non-required field for NewOrderSingle.
	ComplianceID *string `fix:"376"`
	//SolicitedFlag is a non-required field for NewOrderSingle.
	SolicitedFlag *bool `fix:"377"`
	//IOIid is a non-required field for NewOrderSingle.
	IOIid *string `fix:"23"`
	//QuoteID is a non-required field for NewOrderSingle.
	QuoteID *string `fix:"117"`
	//TimeInForce is a non-required field for NewOrderSingle.
	TimeInForce *string `fix:"59"`
	//EffectiveTime is a non-required field for NewOrderSingle.
	EffectiveTime *time.Time `fix:"168"`
	//ExpireDate is a non-required field for NewOrderSingle.
	ExpireDate *string `fix:"432"`
	//ExpireTime is a non-required field for NewOrderSingle.
	ExpireTime *time.Time `fix:"126"`
	//GTBookingInst is a non-required field for NewOrderSingle.
	GTBookingInst *int `fix:"427"`
	//Commission is a non-required field for NewOrderSingle.
	Commission *float64 `fix:"12"`
	//CommType is a non-required field for NewOrderSingle.
	CommType *string `fix:"13"`
	//Rule80A is a non-required field for NewOrderSingle.
	Rule80A *string `fix:"47"`
	//ForexReq is a non-required field for NewOrderSingle.
	ForexReq *bool `fix:"121"`
	//SettlCurrency is a non-required field for NewOrderSingle.
	SettlCurrency *string `fix:"120"`
	//Text is a non-required field for NewOrderSingle.
	Text *string `fix:"58"`
	//EncodedTextLen is a non-required field for NewOrderSingle.
	EncodedTextLen *int `fix:"354"`
	//EncodedText is a non-required field for NewOrderSingle.
	EncodedText *string `fix:"355"`
	//FutSettDate2 is a non-required field for NewOrderSingle.
	FutSettDate2 *string `fix:"193"`
	//OrderQty2 is a non-required field for NewOrderSingle.
	OrderQty2 *float64 `fix:"192"`
	//OpenClose is a non-required field for NewOrderSingle.
	OpenClose *string `fix:"77"`
	//CoveredOrUncovered is a non-required field for NewOrderSingle.
	CoveredOrUncovered *int `fix:"203"`
	//CustomerOrFirm is a non-required field for NewOrderSingle.
	CustomerOrFirm *int `fix:"204"`
	//MaxShow is a non-required field for NewOrderSingle.
	MaxShow *float64 `fix:"210"`
	//PegDifference is a non-required field for NewOrderSingle.
	PegDifference *float64 `fix:"211"`
	//DiscretionInst is a non-required field for NewOrderSingle.
	DiscretionInst *string `fix:"388"`
	//DiscretionOffset is a non-required field for NewOrderSingle.
	DiscretionOffset *float64 `fix:"389"`
	//ClearingFirm is a non-required field for NewOrderSingle.
	ClearingFirm *string `fix:"439"`
	//ClearingAccount is a non-required field for NewOrderSingle.
	ClearingAccount *string `fix:"440"`
	fix42.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

//New returns an initialized NewOrderSingle instance
func New(clordid string, handlinst string, symbol string, side string, transacttime time.Time, ordtype string) *Message {
	var m Message
	m.SetClOrdID(clordid)
	m.SetHandlInst(handlinst)
	m.SetSymbol(symbol)
	m.SetSide(side)
	m.SetTransactTime(transacttime)
	m.SetOrdType(ordtype)
	return &m
}

func (m *Message) SetClOrdID(v string)                        { m.ClOrdID = v }
func (m *Message) SetClientID(v string)                       { m.ClientID = &v }
func (m *Message) SetExecBroker(v string)                     { m.ExecBroker = &v }
func (m *Message) SetAccount(v string)                        { m.Account = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)                   { m.NoAllocs = v }
func (m *Message) SetSettlmntTyp(v string)                    { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)                    { m.FutSettDate = &v }
func (m *Message) SetHandlInst(v string)                      { m.HandlInst = v }
func (m *Message) SetExecInst(v string)                       { m.ExecInst = &v }
func (m *Message) SetMinQty(v float64)                        { m.MinQty = &v }
func (m *Message) SetMaxFloor(v float64)                      { m.MaxFloor = &v }
func (m *Message) SetExDestination(v string)                  { m.ExDestination = &v }
func (m *Message) SetNoTradingSessions(v []NoTradingSessions) { m.NoTradingSessions = v }
func (m *Message) SetProcessCode(v string)                    { m.ProcessCode = &v }
func (m *Message) SetSymbol(v string)                         { m.Symbol = v }
func (m *Message) SetSymbolSfx(v string)                      { m.SymbolSfx = &v }
func (m *Message) SetSecurityID(v string)                     { m.SecurityID = &v }
func (m *Message) SetIDSource(v string)                       { m.IDSource = &v }
func (m *Message) SetSecurityType(v string)                   { m.SecurityType = &v }
func (m *Message) SetMaturityMonthYear(v string)              { m.MaturityMonthYear = &v }
func (m *Message) SetMaturityDay(v int)                       { m.MaturityDay = &v }
func (m *Message) SetPutOrCall(v int)                         { m.PutOrCall = &v }
func (m *Message) SetStrikePrice(v float64)                   { m.StrikePrice = &v }
func (m *Message) SetOptAttribute(v string)                   { m.OptAttribute = &v }
func (m *Message) SetContractMultiplier(v float64)            { m.ContractMultiplier = &v }
func (m *Message) SetCouponRate(v float64)                    { m.CouponRate = &v }
func (m *Message) SetSecurityExchange(v string)               { m.SecurityExchange = &v }
func (m *Message) SetIssuer(v string)                         { m.Issuer = &v }
func (m *Message) SetEncodedIssuerLen(v int)                  { m.EncodedIssuerLen = &v }
func (m *Message) SetEncodedIssuer(v string)                  { m.EncodedIssuer = &v }
func (m *Message) SetSecurityDesc(v string)                   { m.SecurityDesc = &v }
func (m *Message) SetEncodedSecurityDescLen(v int)            { m.EncodedSecurityDescLen = &v }
func (m *Message) SetEncodedSecurityDesc(v string)            { m.EncodedSecurityDesc = &v }
func (m *Message) SetPrevClosePx(v float64)                   { m.PrevClosePx = &v }
func (m *Message) SetSide(v string)                           { m.Side = v }
func (m *Message) SetLocateReqd(v bool)                       { m.LocateReqd = &v }
func (m *Message) SetTransactTime(v time.Time)                { m.TransactTime = v }
func (m *Message) SetOrderQty(v float64)                      { m.OrderQty = &v }
func (m *Message) SetCashOrderQty(v float64)                  { m.CashOrderQty = &v }
func (m *Message) SetOrdType(v string)                        { m.OrdType = v }
func (m *Message) SetPrice(v float64)                         { m.Price = &v }
func (m *Message) SetStopPx(v float64)                        { m.StopPx = &v }
func (m *Message) SetCurrency(v string)                       { m.Currency = &v }
func (m *Message) SetComplianceID(v string)                   { m.ComplianceID = &v }
func (m *Message) SetSolicitedFlag(v bool)                    { m.SolicitedFlag = &v }
func (m *Message) SetIOIid(v string)                          { m.IOIid = &v }
func (m *Message) SetQuoteID(v string)                        { m.QuoteID = &v }
func (m *Message) SetTimeInForce(v string)                    { m.TimeInForce = &v }
func (m *Message) SetEffectiveTime(v time.Time)               { m.EffectiveTime = &v }
func (m *Message) SetExpireDate(v string)                     { m.ExpireDate = &v }
func (m *Message) SetExpireTime(v time.Time)                  { m.ExpireTime = &v }
func (m *Message) SetGTBookingInst(v int)                     { m.GTBookingInst = &v }
func (m *Message) SetCommission(v float64)                    { m.Commission = &v }
func (m *Message) SetCommType(v string)                       { m.CommType = &v }
func (m *Message) SetRule80A(v string)                        { m.Rule80A = &v }
func (m *Message) SetForexReq(v bool)                         { m.ForexReq = &v }
func (m *Message) SetSettlCurrency(v string)                  { m.SettlCurrency = &v }
func (m *Message) SetText(v string)                           { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)                    { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)                    { m.EncodedText = &v }
func (m *Message) SetFutSettDate2(v string)                   { m.FutSettDate2 = &v }
func (m *Message) SetOrderQty2(v float64)                     { m.OrderQty2 = &v }
func (m *Message) SetOpenClose(v string)                      { m.OpenClose = &v }
func (m *Message) SetCoveredOrUncovered(v int)                { m.CoveredOrUncovered = &v }
func (m *Message) SetCustomerOrFirm(v int)                    { m.CustomerOrFirm = &v }
func (m *Message) SetMaxShow(v float64)                       { m.MaxShow = &v }
func (m *Message) SetPegDifference(v float64)                 { m.PegDifference = &v }
func (m *Message) SetDiscretionInst(v string)                 { m.DiscretionInst = &v }
func (m *Message) SetDiscretionOffset(v float64)              { m.DiscretionOffset = &v }
func (m *Message) SetClearingFirm(v string)                   { m.ClearingFirm = &v }
func (m *Message) SetClearingAccount(v string)                { m.ClearingAccount = &v }

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
	return enum.BeginStringFIX42, "D", r
}
