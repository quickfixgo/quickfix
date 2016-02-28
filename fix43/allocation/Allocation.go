//Package allocation msg type = J.
package allocation

import (
	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/fix43/commissiondata"
	"github.com/quickfixgo/quickfix/fix43/instrument"
	"github.com/quickfixgo/quickfix/fix43/nestedparties"
	"github.com/quickfixgo/quickfix/fix43/parties"
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
	//SecondaryClOrdID is a non-required field for NoOrders.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for NoOrders.
	ListID *string `fix:"66"`
}

func (m *NoOrders) SetClOrdID(v string)          { m.ClOrdID = &v }
func (m *NoOrders) SetOrderID(v string)          { m.OrderID = &v }
func (m *NoOrders) SetSecondaryOrderID(v string) { m.SecondaryOrderID = &v }
func (m *NoOrders) SetSecondaryClOrdID(v string) { m.SecondaryClOrdID = &v }
func (m *NoOrders) SetListID(v string)           { m.ListID = &v }

//NoExecs is a repeating group in Allocation
type NoExecs struct {
	//LastQty is a non-required field for NoExecs.
	LastQty *float64 `fix:"32"`
	//ExecID is a non-required field for NoExecs.
	ExecID *string `fix:"17"`
	//SecondaryExecID is a non-required field for NoExecs.
	SecondaryExecID *string `fix:"527"`
	//LastPx is a non-required field for NoExecs.
	LastPx *float64 `fix:"31"`
	//LastCapacity is a non-required field for NoExecs.
	LastCapacity *string `fix:"29"`
}

func (m *NoExecs) SetLastQty(v float64)        { m.LastQty = &v }
func (m *NoExecs) SetExecID(v string)          { m.ExecID = &v }
func (m *NoExecs) SetSecondaryExecID(v string) { m.SecondaryExecID = &v }
func (m *NoExecs) SetLastPx(v float64)         { m.LastPx = &v }
func (m *NoExecs) SetLastCapacity(v string)    { m.LastCapacity = &v }

//NoAllocs is a repeating group in Allocation
type NoAllocs struct {
	//AllocAccount is a non-required field for NoAllocs.
	AllocAccount *string `fix:"79"`
	//AllocPrice is a non-required field for NoAllocs.
	AllocPrice *float64 `fix:"366"`
	//AllocQty is a non-required field for NoAllocs.
	AllocQty *float64 `fix:"80"`
	//IndividualAllocID is a non-required field for NoAllocs.
	IndividualAllocID *string `fix:"467"`
	//ProcessCode is a non-required field for NoAllocs.
	ProcessCode *string `fix:"81"`
	//NestedParties Component
	NestedParties nestedparties.Component
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
	//CommissionData Component
	CommissionData commissiondata.Component
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
func (m *NoAllocs) SetAllocQty(v float64)           { m.AllocQty = &v }
func (m *NoAllocs) SetIndividualAllocID(v string)   { m.IndividualAllocID = &v }
func (m *NoAllocs) SetProcessCode(v string)         { m.ProcessCode = &v }
func (m *NoAllocs) SetNotifyBrokerOfCredit(v bool)  { m.NotifyBrokerOfCredit = &v }
func (m *NoAllocs) SetAllocHandlInst(v int)         { m.AllocHandlInst = &v }
func (m *NoAllocs) SetAllocText(v string)           { m.AllocText = &v }
func (m *NoAllocs) SetEncodedAllocTextLen(v int)    { m.EncodedAllocTextLen = &v }
func (m *NoAllocs) SetEncodedAllocText(v string)    { m.EncodedAllocText = &v }
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
	Header     fix43.Header
	//AllocID is a required field for Allocation.
	AllocID string `fix:"70"`
	//AllocTransType is a required field for Allocation.
	AllocTransType string `fix:"71"`
	//AllocType is a required field for Allocation.
	AllocType int `fix:"626"`
	//RefAllocID is a non-required field for Allocation.
	RefAllocID *string `fix:"72"`
	//AllocLinkID is a non-required field for Allocation.
	AllocLinkID *string `fix:"196"`
	//AllocLinkType is a non-required field for Allocation.
	AllocLinkType *int `fix:"197"`
	//BookingRefID is a non-required field for Allocation.
	BookingRefID *string `fix:"466"`
	//NoOrders is a non-required field for Allocation.
	NoOrders []NoOrders `fix:"73,omitempty"`
	//NoExecs is a non-required field for Allocation.
	NoExecs []NoExecs `fix:"124,omitempty"`
	//Side is a required field for Allocation.
	Side string `fix:"54"`
	//Instrument Component
	Instrument instrument.Component
	//Quantity is a required field for Allocation.
	Quantity float64 `fix:"53"`
	//LastMkt is a non-required field for Allocation.
	LastMkt *string `fix:"30"`
	//TradeOriginationDate is a non-required field for Allocation.
	TradeOriginationDate *string `fix:"229"`
	//TradingSessionID is a non-required field for Allocation.
	TradingSessionID *string `fix:"336"`
	//TradingSessionSubID is a non-required field for Allocation.
	TradingSessionSubID *string `fix:"625"`
	//PriceType is a non-required field for Allocation.
	PriceType *int `fix:"423"`
	//AvgPx is a required field for Allocation.
	AvgPx float64 `fix:"6"`
	//Currency is a non-required field for Allocation.
	Currency *string `fix:"15"`
	//AvgPrxPrecision is a non-required field for Allocation.
	AvgPrxPrecision *int `fix:"74"`
	//Parties Component
	Parties parties.Component
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
	//Concession is a non-required field for Allocation.
	Concession *float64 `fix:"238"`
	//TotalTakedown is a non-required field for Allocation.
	TotalTakedown *float64 `fix:"237"`
	//NetMoney is a non-required field for Allocation.
	NetMoney *float64 `fix:"118"`
	//PositionEffect is a non-required field for Allocation.
	PositionEffect *string `fix:"77"`
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
	//TotalAccruedInterestAmt is a non-required field for Allocation.
	TotalAccruedInterestAmt *float64 `fix:"540"`
	//LegalConfirm is a non-required field for Allocation.
	LegalConfirm *bool `fix:"650"`
	//NoAllocs is a non-required field for Allocation.
	NoAllocs []NoAllocs `fix:"78,omitempty"`
	Trailer  fix43.Trailer
}

//Marshal converts Message to a quickfix.Message instance
func (m Message) Marshal() quickfix.Message { return quickfix.Marshal(m) }

func (m *Message) SetAllocID(v string)                  { m.AllocID = v }
func (m *Message) SetAllocTransType(v string)           { m.AllocTransType = v }
func (m *Message) SetAllocType(v int)                   { m.AllocType = v }
func (m *Message) SetRefAllocID(v string)               { m.RefAllocID = &v }
func (m *Message) SetAllocLinkID(v string)              { m.AllocLinkID = &v }
func (m *Message) SetAllocLinkType(v int)               { m.AllocLinkType = &v }
func (m *Message) SetBookingRefID(v string)             { m.BookingRefID = &v }
func (m *Message) SetNoOrders(v []NoOrders)             { m.NoOrders = v }
func (m *Message) SetNoExecs(v []NoExecs)               { m.NoExecs = v }
func (m *Message) SetSide(v string)                     { m.Side = v }
func (m *Message) SetQuantity(v float64)                { m.Quantity = v }
func (m *Message) SetLastMkt(v string)                  { m.LastMkt = &v }
func (m *Message) SetTradeOriginationDate(v string)     { m.TradeOriginationDate = &v }
func (m *Message) SetTradingSessionID(v string)         { m.TradingSessionID = &v }
func (m *Message) SetTradingSessionSubID(v string)      { m.TradingSessionSubID = &v }
func (m *Message) SetPriceType(v int)                   { m.PriceType = &v }
func (m *Message) SetAvgPx(v float64)                   { m.AvgPx = v }
func (m *Message) SetCurrency(v string)                 { m.Currency = &v }
func (m *Message) SetAvgPrxPrecision(v int)             { m.AvgPrxPrecision = &v }
func (m *Message) SetTradeDate(v string)                { m.TradeDate = v }
func (m *Message) SetTransactTime(v time.Time)          { m.TransactTime = &v }
func (m *Message) SetSettlmntTyp(v string)              { m.SettlmntTyp = &v }
func (m *Message) SetFutSettDate(v string)              { m.FutSettDate = &v }
func (m *Message) SetGrossTradeAmt(v float64)           { m.GrossTradeAmt = &v }
func (m *Message) SetConcession(v float64)              { m.Concession = &v }
func (m *Message) SetTotalTakedown(v float64)           { m.TotalTakedown = &v }
func (m *Message) SetNetMoney(v float64)                { m.NetMoney = &v }
func (m *Message) SetPositionEffect(v string)           { m.PositionEffect = &v }
func (m *Message) SetText(v string)                     { m.Text = &v }
func (m *Message) SetEncodedTextLen(v int)              { m.EncodedTextLen = &v }
func (m *Message) SetEncodedText(v string)              { m.EncodedText = &v }
func (m *Message) SetNumDaysInterest(v int)             { m.NumDaysInterest = &v }
func (m *Message) SetAccruedInterestRate(v float64)     { m.AccruedInterestRate = &v }
func (m *Message) SetTotalAccruedInterestAmt(v float64) { m.TotalAccruedInterestAmt = &v }
func (m *Message) SetLegalConfirm(v bool)               { m.LegalConfirm = &v }
func (m *Message) SetNoAllocs(v []NoAllocs)             { m.NoAllocs = v }

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
	return enum.BeginStringFIX43, "J", r
}
