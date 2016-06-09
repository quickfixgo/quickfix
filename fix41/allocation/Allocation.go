package allocation

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix41"
	"github.com/quickfixgo/quickfix/tag"
)

//Allocation is the fix41 Allocation type, MsgType = J
type Allocation struct {
	fix41.Header
	quickfix.Body
	fix41.Trailer
}

//FromMessage creates a Allocation from a quickfix.Message instance
func FromMessage(m quickfix.Message) Allocation {
	return Allocation{
		Header:  fix41.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix41.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m Allocation) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a Allocation initialized with the required fields for Allocation
func New(allocid field.AllocIDField, alloctranstype field.AllocTransTypeField, side field.SideField, symbol field.SymbolField, shares field.SharesField, avgpx field.AvgPxField, tradedate field.TradeDateField) (m Allocation) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("J"))
	m.Set(allocid)
	m.Set(alloctranstype)
	m.Set(side)
	m.Set(symbol)
	m.Set(shares)
	m.Set(avgpx)
	m.Set(tradedate)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg Allocation, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.1", "J", r
}

//SetAvgPx sets AvgPx, Tag 6
func (m Allocation) SetAvgPx(v float64) {
	m.Set(field.NewAvgPx(v))
}

//SetCurrency sets Currency, Tag 15
func (m Allocation) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m Allocation) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m Allocation) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Allocation) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetShares sets Shares, Tag 53
func (m Allocation) SetShares(v float64) {
	m.Set(field.NewShares(v))
}

//SetSide sets Side, Tag 54
func (m Allocation) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m Allocation) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m Allocation) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m Allocation) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m Allocation) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Allocation) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Allocation) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetAllocID sets AllocID, Tag 70
func (m Allocation) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetAllocTransType sets AllocTransType, Tag 71
func (m Allocation) SetAllocTransType(v string) {
	m.Set(field.NewAllocTransType(v))
}

//SetRefAllocID sets RefAllocID, Tag 72
func (m Allocation) SetRefAllocID(v string) {
	m.Set(field.NewRefAllocID(v))
}

//SetNoOrders sets NoOrders, Tag 73
func (m Allocation) SetNoOrders(f NoOrdersRepeatingGroup) {
	m.SetGroup(f)
}

//SetAvgPrxPrecision sets AvgPrxPrecision, Tag 74
func (m Allocation) SetAvgPrxPrecision(v int) {
	m.Set(field.NewAvgPrxPrecision(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m Allocation) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetOpenClose sets OpenClose, Tag 77
func (m Allocation) SetOpenClose(v string) {
	m.Set(field.NewOpenClose(v))
}

//SetNoAllocs sets NoAllocs, Tag 78
func (m Allocation) SetNoAllocs(f NoAllocsRepeatingGroup) {
	m.SetGroup(f)
}

//SetIssuer sets Issuer, Tag 106
func (m Allocation) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m Allocation) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetNetMoney sets NetMoney, Tag 118
func (m Allocation) SetNetMoney(v float64) {
	m.Set(field.NewNetMoney(v))
}

//SetNoExecs sets NoExecs, Tag 124
func (m Allocation) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNumDaysInterest sets NumDaysInterest, Tag 157
func (m Allocation) SetNumDaysInterest(v int) {
	m.Set(field.NewNumDaysInterest(v))
}

//SetAccruedInterestRate sets AccruedInterestRate, Tag 158
func (m Allocation) SetAccruedInterestRate(v float64) {
	m.Set(field.NewAccruedInterestRate(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m Allocation) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetAllocLinkID sets AllocLinkID, Tag 196
func (m Allocation) SetAllocLinkID(v string) {
	m.Set(field.NewAllocLinkID(v))
}

//SetAllocLinkType sets AllocLinkType, Tag 197
func (m Allocation) SetAllocLinkType(v int) {
	m.Set(field.NewAllocLinkType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m Allocation) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m Allocation) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m Allocation) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m Allocation) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m Allocation) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m Allocation) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//GetAvgPx gets AvgPx, Tag 6
func (m Allocation) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m Allocation) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m Allocation) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m Allocation) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Allocation) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetShares gets Shares, Tag 53
func (m Allocation) GetShares() (f field.SharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m Allocation) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Allocation) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m Allocation) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Allocation) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m Allocation) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Allocation) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Allocation) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocID gets AllocID, Tag 70
func (m Allocation) GetAllocID() (f field.AllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocTransType gets AllocTransType, Tag 71
func (m Allocation) GetAllocTransType() (f field.AllocTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRefAllocID gets RefAllocID, Tag 72
func (m Allocation) GetRefAllocID() (f field.RefAllocIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m Allocation) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAvgPrxPrecision gets AvgPrxPrecision, Tag 74
func (m Allocation) GetAvgPrxPrecision() (f field.AvgPrxPrecisionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m Allocation) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m Allocation) GetOpenClose() (f field.OpenCloseField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m Allocation) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetIssuer gets Issuer, Tag 106
func (m Allocation) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Allocation) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m Allocation) GetNetMoney() (f field.NetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoExecs gets NoExecs, Tag 124
func (m Allocation) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNumDaysInterest gets NumDaysInterest, Tag 157
func (m Allocation) GetNumDaysInterest() (f field.NumDaysInterestField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestRate gets AccruedInterestRate, Tag 158
func (m Allocation) GetAccruedInterestRate() (f field.AccruedInterestRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m Allocation) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocLinkID gets AllocLinkID, Tag 196
func (m Allocation) GetAllocLinkID() (f field.AllocLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocLinkType gets AllocLinkType, Tag 197
func (m Allocation) GetAllocLinkType() (f field.AllocLinkTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m Allocation) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m Allocation) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m Allocation) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m Allocation) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m Allocation) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m Allocation) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m Allocation) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m Allocation) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m Allocation) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m Allocation) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m Allocation) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasShares returns true if Shares is present, Tag 53
func (m Allocation) HasShares() bool {
	return m.Has(tag.Shares)
}

//HasSide returns true if Side is present, Tag 54
func (m Allocation) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m Allocation) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m Allocation) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m Allocation) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m Allocation) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m Allocation) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m Allocation) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasAllocID returns true if AllocID is present, Tag 70
func (m Allocation) HasAllocID() bool {
	return m.Has(tag.AllocID)
}

//HasAllocTransType returns true if AllocTransType is present, Tag 71
func (m Allocation) HasAllocTransType() bool {
	return m.Has(tag.AllocTransType)
}

//HasRefAllocID returns true if RefAllocID is present, Tag 72
func (m Allocation) HasRefAllocID() bool {
	return m.Has(tag.RefAllocID)
}

//HasNoOrders returns true if NoOrders is present, Tag 73
func (m Allocation) HasNoOrders() bool {
	return m.Has(tag.NoOrders)
}

//HasAvgPrxPrecision returns true if AvgPrxPrecision is present, Tag 74
func (m Allocation) HasAvgPrxPrecision() bool {
	return m.Has(tag.AvgPrxPrecision)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m Allocation) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasOpenClose returns true if OpenClose is present, Tag 77
func (m Allocation) HasOpenClose() bool {
	return m.Has(tag.OpenClose)
}

//HasNoAllocs returns true if NoAllocs is present, Tag 78
func (m Allocation) HasNoAllocs() bool {
	return m.Has(tag.NoAllocs)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m Allocation) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m Allocation) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasNetMoney returns true if NetMoney is present, Tag 118
func (m Allocation) HasNetMoney() bool {
	return m.Has(tag.NetMoney)
}

//HasNoExecs returns true if NoExecs is present, Tag 124
func (m Allocation) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

//HasNumDaysInterest returns true if NumDaysInterest is present, Tag 157
func (m Allocation) HasNumDaysInterest() bool {
	return m.Has(tag.NumDaysInterest)
}

//HasAccruedInterestRate returns true if AccruedInterestRate is present, Tag 158
func (m Allocation) HasAccruedInterestRate() bool {
	return m.Has(tag.AccruedInterestRate)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m Allocation) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasAllocLinkID returns true if AllocLinkID is present, Tag 196
func (m Allocation) HasAllocLinkID() bool {
	return m.Has(tag.AllocLinkID)
}

//HasAllocLinkType returns true if AllocLinkType is present, Tag 197
func (m Allocation) HasAllocLinkType() bool {
	return m.Has(tag.AllocLinkType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m Allocation) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m Allocation) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m Allocation) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m Allocation) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m Allocation) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m Allocation) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoOrders) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m NoOrders) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetListID sets ListID, Tag 66
func (m NoOrders) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetWaveNo sets WaveNo, Tag 105
func (m NoOrders) SetWaveNo(v string) {
	m.Set(field.NewWaveNo(v))
}

//GetClOrdID gets ClOrdID, Tag 11
func (m NoOrders) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoOrders) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m NoOrders) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m NoOrders) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m NoOrders) GetWaveNo() (f field.WaveNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m NoOrders) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoOrders) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m NoOrders) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasListID returns true if ListID is present, Tag 66
func (m NoOrders) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasWaveNo returns true if WaveNo is present, Tag 105
func (m NoOrders) HasWaveNo() bool {
	return m.Has(tag.WaveNo)
}

//NoOrdersRepeatingGroup is a repeating group, Tag 73
type NoOrdersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdersRepeatingGroup returns an initialized, NoOrdersRepeatingGroup
func NewNoOrdersRepeatingGroup() NoOrdersRepeatingGroup {
	return NoOrdersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrders,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.SecondaryOrderID), quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.WaveNo)})}
}

//Add create and append a new NoOrders to this group
func (m NoOrdersRepeatingGroup) Add() NoOrders {
	g := m.RepeatingGroup.Add()
	return NoOrders{g}
}

//Get returns the ith NoOrders in the NoOrdersRepeatinGroup
func (m NoOrdersRepeatingGroup) Get(i int) NoOrders {
	return NoOrders{m.RepeatingGroup.Get(i)}
}

//NoAllocs is a repeating group element, Tag 78
type NoAllocs struct {
	quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocShares sets AllocShares, Tag 80
func (m NoAllocs) SetAllocShares(v float64) {
	m.Set(field.NewAllocShares(v))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoAllocs) SetProcessCode(v string) {
	m.Set(field.NewProcessCode(v))
}

//SetBrokerOfCredit sets BrokerOfCredit, Tag 92
func (m NoAllocs) SetBrokerOfCredit(v string) {
	m.Set(field.NewBrokerOfCredit(v))
}

//SetNotifyBrokerOfCredit sets NotifyBrokerOfCredit, Tag 208
func (m NoAllocs) SetNotifyBrokerOfCredit(v bool) {
	m.Set(field.NewNotifyBrokerOfCredit(v))
}

//SetAllocHandlInst sets AllocHandlInst, Tag 209
func (m NoAllocs) SetAllocHandlInst(v int) {
	m.Set(field.NewAllocHandlInst(v))
}

//SetAllocText sets AllocText, Tag 161
func (m NoAllocs) SetAllocText(v string) {
	m.Set(field.NewAllocText(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m NoAllocs) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetClientID sets ClientID, Tag 109
func (m NoAllocs) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetCommission sets Commission, Tag 12
func (m NoAllocs) SetCommission(v float64) {
	m.Set(field.NewCommission(v))
}

//SetCommType sets CommType, Tag 13
func (m NoAllocs) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetAllocAvgPx sets AllocAvgPx, Tag 153
func (m NoAllocs) SetAllocAvgPx(v float64) {
	m.Set(field.NewAllocAvgPx(v))
}

//SetAllocNetMoney sets AllocNetMoney, Tag 154
func (m NoAllocs) SetAllocNetMoney(v float64) {
	m.Set(field.NewAllocNetMoney(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m NoAllocs) SetSettlCurrAmt(v float64) {
	m.Set(field.NewSettlCurrAmt(v))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m NoAllocs) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m NoAllocs) SetSettlCurrFxRate(v float64) {
	m.Set(field.NewSettlCurrFxRate(v))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m NoAllocs) SetSettlCurrFxRateCalc(v string) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetAccruedInterestAmt sets AccruedInterestAmt, Tag 159
func (m NoAllocs) SetAccruedInterestAmt(v float64) {
	m.Set(field.NewAccruedInterestAmt(v))
}

//SetSettlInstMode sets SettlInstMode, Tag 160
func (m NoAllocs) SetSettlInstMode(v string) {
	m.Set(field.NewSettlInstMode(v))
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m NoAllocs) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (f field.AllocAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocShares gets AllocShares, Tag 80
func (m NoAllocs) GetAllocShares() (f field.AllocSharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoAllocs) GetProcessCode() (f field.ProcessCodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBrokerOfCredit gets BrokerOfCredit, Tag 92
func (m NoAllocs) GetBrokerOfCredit() (f field.BrokerOfCreditField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNotifyBrokerOfCredit gets NotifyBrokerOfCredit, Tag 208
func (m NoAllocs) GetNotifyBrokerOfCredit() (f field.NotifyBrokerOfCreditField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocHandlInst gets AllocHandlInst, Tag 209
func (m NoAllocs) GetAllocHandlInst() (f field.AllocHandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocText gets AllocText, Tag 161
func (m NoAllocs) GetAllocText() (f field.AllocTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m NoAllocs) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m NoAllocs) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m NoAllocs) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m NoAllocs) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocAvgPx gets AllocAvgPx, Tag 153
func (m NoAllocs) GetAllocAvgPx() (f field.AllocAvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAllocNetMoney gets AllocNetMoney, Tag 154
func (m NoAllocs) GetAllocNetMoney() (f field.AllocNetMoneyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m NoAllocs) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m NoAllocs) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m NoAllocs) GetSettlCurrFxRate() (f field.SettlCurrFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m NoAllocs) GetSettlCurrFxRateCalc() (f field.SettlCurrFxRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccruedInterestAmt gets AccruedInterestAmt, Tag 159
func (m NoAllocs) GetAccruedInterestAmt() (f field.AccruedInterestAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlInstMode gets SettlInstMode, Tag 160
func (m NoAllocs) GetSettlInstMode() (f field.SettlInstModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m NoAllocs) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasAllocAccount returns true if AllocAccount is present, Tag 79
func (m NoAllocs) HasAllocAccount() bool {
	return m.Has(tag.AllocAccount)
}

//HasAllocShares returns true if AllocShares is present, Tag 80
func (m NoAllocs) HasAllocShares() bool {
	return m.Has(tag.AllocShares)
}

//HasProcessCode returns true if ProcessCode is present, Tag 81
func (m NoAllocs) HasProcessCode() bool {
	return m.Has(tag.ProcessCode)
}

//HasBrokerOfCredit returns true if BrokerOfCredit is present, Tag 92
func (m NoAllocs) HasBrokerOfCredit() bool {
	return m.Has(tag.BrokerOfCredit)
}

//HasNotifyBrokerOfCredit returns true if NotifyBrokerOfCredit is present, Tag 208
func (m NoAllocs) HasNotifyBrokerOfCredit() bool {
	return m.Has(tag.NotifyBrokerOfCredit)
}

//HasAllocHandlInst returns true if AllocHandlInst is present, Tag 209
func (m NoAllocs) HasAllocHandlInst() bool {
	return m.Has(tag.AllocHandlInst)
}

//HasAllocText returns true if AllocText is present, Tag 161
func (m NoAllocs) HasAllocText() bool {
	return m.Has(tag.AllocText)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m NoAllocs) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m NoAllocs) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m NoAllocs) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m NoAllocs) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasAllocAvgPx returns true if AllocAvgPx is present, Tag 153
func (m NoAllocs) HasAllocAvgPx() bool {
	return m.Has(tag.AllocAvgPx)
}

//HasAllocNetMoney returns true if AllocNetMoney is present, Tag 154
func (m NoAllocs) HasAllocNetMoney() bool {
	return m.Has(tag.AllocNetMoney)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m NoAllocs) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m NoAllocs) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m NoAllocs) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m NoAllocs) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasAccruedInterestAmt returns true if AccruedInterestAmt is present, Tag 159
func (m NoAllocs) HasAccruedInterestAmt() bool {
	return m.Has(tag.AccruedInterestAmt)
}

//HasSettlInstMode returns true if SettlInstMode is present, Tag 160
func (m NoAllocs) HasSettlInstMode() bool {
	return m.Has(tag.SettlInstMode)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m NoAllocs) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(v float64) {
	m.Set(field.NewMiscFeeAmt(v))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v string) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (f field.MiscFeeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (f field.MiscFeeCurrField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (f field.MiscFeeTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasMiscFeeAmt returns true if MiscFeeAmt is present, Tag 137
func (m NoMiscFees) HasMiscFeeAmt() bool {
	return m.Has(tag.MiscFeeAmt)
}

//HasMiscFeeCurr returns true if MiscFeeCurr is present, Tag 138
func (m NoMiscFees) HasMiscFeeCurr() bool {
	return m.Has(tag.MiscFeeCurr)
}

//HasMiscFeeType returns true if MiscFeeType is present, Tag 139
func (m NoMiscFees) HasMiscFeeType() bool {
	return m.Has(tag.MiscFeeType)
}

//NoMiscFeesRepeatingGroup is a repeating group, Tag 136
type NoMiscFeesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMiscFeesRepeatingGroup returns an initialized, NoMiscFeesRepeatingGroup
func NewNoMiscFeesRepeatingGroup() NoMiscFeesRepeatingGroup {
	return NoMiscFeesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMiscFees,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MiscFeeAmt), quickfix.GroupElement(tag.MiscFeeCurr), quickfix.GroupElement(tag.MiscFeeType)})}
}

//Add create and append a new NoMiscFees to this group
func (m NoMiscFeesRepeatingGroup) Add() NoMiscFees {
	g := m.RepeatingGroup.Add()
	return NoMiscFees{g}
}

//Get returns the ith NoMiscFees in the NoMiscFeesRepeatinGroup
func (m NoMiscFeesRepeatingGroup) Get(i int) NoMiscFees {
	return NoMiscFees{m.RepeatingGroup.Get(i)}
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocShares), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.BrokerOfCredit), quickfix.GroupElement(tag.NotifyBrokerOfCredit), quickfix.GroupElement(tag.AllocHandlInst), quickfix.GroupElement(tag.AllocText), quickfix.GroupElement(tag.ExecBroker), quickfix.GroupElement(tag.ClientID), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.AllocAvgPx), quickfix.GroupElement(tag.AllocNetMoney), quickfix.GroupElement(tag.SettlCurrAmt), quickfix.GroupElement(tag.SettlCurrency), quickfix.GroupElement(tag.SettlCurrFxRate), quickfix.GroupElement(tag.SettlCurrFxRateCalc), quickfix.GroupElement(tag.AccruedInterestAmt), quickfix.GroupElement(tag.SettlInstMode), quickfix.GroupElement(tag.NoMiscFees)})}
}

//Add create and append a new NoAllocs to this group
func (m NoAllocsRepeatingGroup) Add() NoAllocs {
	g := m.RepeatingGroup.Add()
	return NoAllocs{g}
}

//Get returns the ith NoAllocs in the NoAllocsRepeatinGroup
func (m NoAllocsRepeatingGroup) Get(i int) NoAllocs {
	return NoAllocs{m.RepeatingGroup.Get(i)}
}

//NoExecs is a repeating group element, Tag 124
type NoExecs struct {
	quickfix.Group
}

//SetLastShares sets LastShares, Tag 32
func (m NoExecs) SetLastShares(v float64) {
	m.Set(field.NewLastShares(v))
}

//SetExecID sets ExecID, Tag 17
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetLastPx sets LastPx, Tag 31
func (m NoExecs) SetLastPx(v float64) {
	m.Set(field.NewLastPx(v))
}

//SetLastCapacity sets LastCapacity, Tag 29
func (m NoExecs) SetLastCapacity(v string) {
	m.Set(field.NewLastCapacity(v))
}

//GetLastShares gets LastShares, Tag 32
func (m NoExecs) GetLastShares() (f field.LastSharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m NoExecs) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m NoExecs) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastCapacity gets LastCapacity, Tag 29
func (m NoExecs) GetLastCapacity() (f field.LastCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasLastShares returns true if LastShares is present, Tag 32
func (m NoExecs) HasLastShares() bool {
	return m.Has(tag.LastShares)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m NoExecs) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastCapacity returns true if LastCapacity is present, Tag 29
func (m NoExecs) HasLastCapacity() bool {
	return m.Has(tag.LastCapacity)
}

//NoExecsRepeatingGroup is a repeating group, Tag 124
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LastShares), quickfix.GroupElement(tag.ExecID), quickfix.GroupElement(tag.LastPx), quickfix.GroupElement(tag.LastCapacity)})}
}

//Add create and append a new NoExecs to this group
func (m NoExecsRepeatingGroup) Add() NoExecs {
	g := m.RepeatingGroup.Add()
	return NoExecs{g}
}

//Get returns the ith NoExecs in the NoExecsRepeatinGroup
func (m NoExecsRepeatingGroup) Get(i int) NoExecs {
	return NoExecs{m.RepeatingGroup.Get(i)}
}
