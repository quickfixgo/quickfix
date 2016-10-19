package allocation

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix40"
	"github.com/quickfixgo/quickfix/tag"
)

//Allocation is the fix40 Allocation type, MsgType = J
type Allocation struct {
	fix40.Header
	*quickfix.Body
	fix40.Trailer
	Message *quickfix.Message
}

//FromMessage creates a Allocation from a quickfix.Message instance
func FromMessage(m *quickfix.Message) Allocation {
	return Allocation{
		Header:  fix40.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix40.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m Allocation) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a Allocation initialized with the required fields for Allocation
func New(allocid field.AllocIDField, alloctranstype field.AllocTransTypeField, side field.SideField, symbol field.SymbolField, shares field.SharesField, avgpx field.AvgPxField, tradedate field.TradeDateField) (m Allocation) {
	m.Message = quickfix.NewMessage()
	m.Header = fix40.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

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
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.0", "J", r
}

//SetAvgPx sets AvgPx, Tag 6
func (m Allocation) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m Allocation) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetIDSource sets IDSource, Tag 22
func (m Allocation) SetIDSource(v enum.IDSource) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m Allocation) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetShares sets Shares, Tag 53
func (m Allocation) SetShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewShares(value, scale))
}

//SetSide sets Side, Tag 54
func (m Allocation) SetSide(v enum.Side) {
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
func (m Allocation) SetSettlmntTyp(v enum.SettlmntTyp) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m Allocation) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m Allocation) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetAllocID sets AllocID, Tag 70
func (m Allocation) SetAllocID(v string) {
	m.Set(field.NewAllocID(v))
}

//SetAllocTransType sets AllocTransType, Tag 71
func (m Allocation) SetAllocTransType(v enum.AllocTransType) {
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
func (m Allocation) SetOpenClose(v enum.OpenClose) {
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
func (m Allocation) SetNetMoney(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetMoney(value, scale))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m Allocation) SetSettlCurrAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrAmt(value, scale))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m Allocation) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetNoExecs sets NoExecs, Tag 124
func (m Allocation) SetNoExecs(f NoExecsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMiscFees sets NoMiscFees, Tag 136
func (m Allocation) SetNoMiscFees(f NoMiscFeesRepeatingGroup) {
	m.SetGroup(f)
}

//GetAvgPx gets AvgPx, Tag 6
func (m Allocation) GetAvgPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AvgPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m Allocation) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIDSource gets IDSource, Tag 22
func (m Allocation) GetIDSource() (v enum.IDSource, err quickfix.MessageRejectError) {
	var f field.IDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m Allocation) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetShares gets Shares, Tag 53
func (m Allocation) GetShares() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSide gets Side, Tag 54
func (m Allocation) GetSide() (v enum.Side, err quickfix.MessageRejectError) {
	var f field.SideField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m Allocation) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m Allocation) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m Allocation) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m Allocation) GetSettlmntTyp() (v enum.SettlmntTyp, err quickfix.MessageRejectError) {
	var f field.SettlmntTypField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m Allocation) GetFutSettDate() (v string, err quickfix.MessageRejectError) {
	var f field.FutSettDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m Allocation) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocID gets AllocID, Tag 70
func (m Allocation) GetAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.AllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocTransType gets AllocTransType, Tag 71
func (m Allocation) GetAllocTransType() (v enum.AllocTransType, err quickfix.MessageRejectError) {
	var f field.AllocTransTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRefAllocID gets RefAllocID, Tag 72
func (m Allocation) GetRefAllocID() (v string, err quickfix.MessageRejectError) {
	var f field.RefAllocIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoOrders gets NoOrders, Tag 73
func (m Allocation) GetNoOrders() (NoOrdersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetAvgPrxPrecision gets AvgPrxPrecision, Tag 74
func (m Allocation) GetAvgPrxPrecision() (v int, err quickfix.MessageRejectError) {
	var f field.AvgPrxPrecisionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m Allocation) GetTradeDate() (v string, err quickfix.MessageRejectError) {
	var f field.TradeDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m Allocation) GetOpenClose() (v enum.OpenClose, err quickfix.MessageRejectError) {
	var f field.OpenCloseField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoAllocs gets NoAllocs, Tag 78
func (m Allocation) GetNoAllocs() (NoAllocsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoAllocsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetIssuer gets Issuer, Tag 106
func (m Allocation) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m Allocation) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetMoney gets NetMoney, Tag 118
func (m Allocation) GetNetMoney() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NetMoneyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m Allocation) GetSettlCurrAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.SettlCurrAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m Allocation) GetSettlCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.SettlCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoExecs gets NoExecs, Tag 124
func (m Allocation) GetNoExecs() (NoExecsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMiscFees gets NoMiscFees, Tag 136
func (m Allocation) GetNoMiscFees() (NoMiscFeesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMiscFeesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m Allocation) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m Allocation) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasNoExecs returns true if NoExecs is present, Tag 124
func (m Allocation) HasNoExecs() bool {
	return m.Has(tag.NoExecs)
}

//HasNoMiscFees returns true if NoMiscFees is present, Tag 136
func (m Allocation) HasNoMiscFees() bool {
	return m.Has(tag.NoMiscFees)
}

//NoOrders is a repeating group element, Tag 73
type NoOrders struct {
	*quickfix.Group
}

//SetClOrdID sets ClOrdID, Tag 11
func (m NoOrders) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoOrders) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
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
func (m NoOrders) GetClOrdID() (v string, err quickfix.MessageRejectError) {
	var f field.ClOrdIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoOrders) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetListID gets ListID, Tag 66
func (m NoOrders) GetListID() (v string, err quickfix.MessageRejectError) {
	var f field.ListIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetWaveNo gets WaveNo, Tag 105
func (m NoOrders) GetWaveNo() (v string, err quickfix.MessageRejectError) {
	var f field.WaveNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ClOrdID), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.ListID), quickfix.GroupElement(tag.WaveNo)})}
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
	*quickfix.Group
}

//SetAllocAccount sets AllocAccount, Tag 79
func (m NoAllocs) SetAllocAccount(v string) {
	m.Set(field.NewAllocAccount(v))
}

//SetAllocShares sets AllocShares, Tag 80
func (m NoAllocs) SetAllocShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewAllocShares(value, scale))
}

//SetProcessCode sets ProcessCode, Tag 81
func (m NoAllocs) SetProcessCode(v enum.ProcessCode) {
	m.Set(field.NewProcessCode(v))
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
func (m NoAllocs) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m NoAllocs) SetCommType(v enum.CommType) {
	m.Set(field.NewCommType(v))
}

//SetNoDlvyInst sets NoDlvyInst, Tag 85
func (m NoAllocs) SetNoDlvyInst(v int) {
	m.Set(field.NewNoDlvyInst(v))
}

//SetBrokerOfCredit sets BrokerOfCredit, Tag 92
func (m NoAllocs) SetBrokerOfCredit(v string) {
	m.Set(field.NewBrokerOfCredit(v))
}

//SetDlvyInst sets DlvyInst, Tag 86
func (m NoAllocs) SetDlvyInst(v string) {
	m.Set(field.NewDlvyInst(v))
}

//GetAllocAccount gets AllocAccount, Tag 79
func (m NoAllocs) GetAllocAccount() (v string, err quickfix.MessageRejectError) {
	var f field.AllocAccountField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetAllocShares gets AllocShares, Tag 80
func (m NoAllocs) GetAllocShares() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.AllocSharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetProcessCode gets ProcessCode, Tag 81
func (m NoAllocs) GetProcessCode() (v enum.ProcessCode, err quickfix.MessageRejectError) {
	var f field.ProcessCodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m NoAllocs) GetExecBroker() (v string, err quickfix.MessageRejectError) {
	var f field.ExecBrokerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetClientID gets ClientID, Tag 109
func (m NoAllocs) GetClientID() (v string, err quickfix.MessageRejectError) {
	var f field.ClientIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommission gets Commission, Tag 12
func (m NoAllocs) GetCommission() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CommissionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCommType gets CommType, Tag 13
func (m NoAllocs) GetCommType() (v enum.CommType, err quickfix.MessageRejectError) {
	var f field.CommTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoDlvyInst gets NoDlvyInst, Tag 85
func (m NoAllocs) GetNoDlvyInst() (v int, err quickfix.MessageRejectError) {
	var f field.NoDlvyInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetBrokerOfCredit gets BrokerOfCredit, Tag 92
func (m NoAllocs) GetBrokerOfCredit() (v string, err quickfix.MessageRejectError) {
	var f field.BrokerOfCreditField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDlvyInst gets DlvyInst, Tag 86
func (m NoAllocs) GetDlvyInst() (v string, err quickfix.MessageRejectError) {
	var f field.DlvyInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
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

//HasNoDlvyInst returns true if NoDlvyInst is present, Tag 85
func (m NoAllocs) HasNoDlvyInst() bool {
	return m.Has(tag.NoDlvyInst)
}

//HasBrokerOfCredit returns true if BrokerOfCredit is present, Tag 92
func (m NoAllocs) HasBrokerOfCredit() bool {
	return m.Has(tag.BrokerOfCredit)
}

//HasDlvyInst returns true if DlvyInst is present, Tag 86
func (m NoAllocs) HasDlvyInst() bool {
	return m.Has(tag.DlvyInst)
}

//NoAllocsRepeatingGroup is a repeating group, Tag 78
type NoAllocsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoAllocsRepeatingGroup returns an initialized, NoAllocsRepeatingGroup
func NewNoAllocsRepeatingGroup() NoAllocsRepeatingGroup {
	return NoAllocsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoAllocs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.AllocAccount), quickfix.GroupElement(tag.AllocShares), quickfix.GroupElement(tag.ProcessCode), quickfix.GroupElement(tag.ExecBroker), quickfix.GroupElement(tag.ClientID), quickfix.GroupElement(tag.Commission), quickfix.GroupElement(tag.CommType), quickfix.GroupElement(tag.NoDlvyInst), quickfix.GroupElement(tag.BrokerOfCredit), quickfix.GroupElement(tag.DlvyInst)})}
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
	*quickfix.Group
}

//SetExecID sets ExecID, Tag 17
func (m NoExecs) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetLastShares sets LastShares, Tag 32
func (m NoExecs) SetLastShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastShares(value, scale))
}

//SetLastPx sets LastPx, Tag 31
func (m NoExecs) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastMkt sets LastMkt, Tag 30
func (m NoExecs) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//GetExecID gets ExecID, Tag 17
func (m NoExecs) GetExecID() (v string, err quickfix.MessageRejectError) {
	var f field.ExecIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastShares gets LastShares, Tag 32
func (m NoExecs) GetLastShares() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastSharesField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastPx gets LastPx, Tag 31
func (m NoExecs) GetLastPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LastPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m NoExecs) GetLastMkt() (v string, err quickfix.MessageRejectError) {
	var f field.LastMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasExecID returns true if ExecID is present, Tag 17
func (m NoExecs) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasLastShares returns true if LastShares is present, Tag 32
func (m NoExecs) HasLastShares() bool {
	return m.Has(tag.LastShares)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m NoExecs) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m NoExecs) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//NoExecsRepeatingGroup is a repeating group, Tag 124
type NoExecsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecsRepeatingGroup returns an initialized, NoExecsRepeatingGroup
func NewNoExecsRepeatingGroup() NoExecsRepeatingGroup {
	return NoExecsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecID), quickfix.GroupElement(tag.LastShares), quickfix.GroupElement(tag.LastPx), quickfix.GroupElement(tag.LastMkt)})}
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

//NoMiscFees is a repeating group element, Tag 136
type NoMiscFees struct {
	*quickfix.Group
}

//SetMiscFeeAmt sets MiscFeeAmt, Tag 137
func (m NoMiscFees) SetMiscFeeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewMiscFeeAmt(value, scale))
}

//SetMiscFeeCurr sets MiscFeeCurr, Tag 138
func (m NoMiscFees) SetMiscFeeCurr(v string) {
	m.Set(field.NewMiscFeeCurr(v))
}

//SetMiscFeeType sets MiscFeeType, Tag 139
func (m NoMiscFees) SetMiscFeeType(v enum.MiscFeeType) {
	m.Set(field.NewMiscFeeType(v))
}

//GetMiscFeeAmt gets MiscFeeAmt, Tag 137
func (m NoMiscFees) GetMiscFeeAmt() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MiscFeeAmtField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeCurr gets MiscFeeCurr, Tag 138
func (m NoMiscFees) GetMiscFeeCurr() (v string, err quickfix.MessageRejectError) {
	var f field.MiscFeeCurrField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMiscFeeType gets MiscFeeType, Tag 139
func (m NoMiscFees) GetMiscFeeType() (v enum.MiscFeeType, err quickfix.MessageRejectError) {
	var f field.MiscFeeTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
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
