package executionreport

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//ExecutionReport is the fix42 ExecutionReport type, MsgType = 8
type ExecutionReport struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a ExecutionReport from a quickfix.Message instance
func FromMessage(m quickfix.Message) ExecutionReport {
	return ExecutionReport{
		Header:      fix42.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix42.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m ExecutionReport) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a ExecutionReport initialized with the required fields for ExecutionReport
func New(orderid field.OrderIDField, execid field.ExecIDField, exectranstype field.ExecTransTypeField, exectype field.ExecTypeField, ordstatus field.OrdStatusField, symbol field.SymbolField, side field.SideField, leavesqty field.LeavesQtyField, cumqty field.CumQtyField, avgpx field.AvgPxField) (m ExecutionReport) {
	m.Header = fix42.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("8"))
	m.Set(orderid)
	m.Set(execid)
	m.Set(exectranstype)
	m.Set(exectype)
	m.Set(ordstatus)
	m.Set(symbol)
	m.Set(side)
	m.Set(leavesqty)
	m.Set(cumqty)
	m.Set(avgpx)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg ExecutionReport, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "8", r
}

//SetAccount sets Account, Tag 1
func (m ExecutionReport) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetAvgPx sets AvgPx, Tag 6
func (m ExecutionReport) SetAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewAvgPx(value, scale))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m ExecutionReport) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetCommission sets Commission, Tag 12
func (m ExecutionReport) SetCommission(value decimal.Decimal, scale int32) {
	m.Set(field.NewCommission(value, scale))
}

//SetCommType sets CommType, Tag 13
func (m ExecutionReport) SetCommType(v string) {
	m.Set(field.NewCommType(v))
}

//SetCumQty sets CumQty, Tag 14
func (m ExecutionReport) SetCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCumQty(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m ExecutionReport) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetExecID sets ExecID, Tag 17
func (m ExecutionReport) SetExecID(v string) {
	m.Set(field.NewExecID(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m ExecutionReport) SetExecInst(v string) {
	m.Set(field.NewExecInst(v))
}

//SetExecRefID sets ExecRefID, Tag 19
func (m ExecutionReport) SetExecRefID(v string) {
	m.Set(field.NewExecRefID(v))
}

//SetExecTransType sets ExecTransType, Tag 20
func (m ExecutionReport) SetExecTransType(v string) {
	m.Set(field.NewExecTransType(v))
}

//SetHandlInst sets HandlInst, Tag 21
func (m ExecutionReport) SetHandlInst(v string) {
	m.Set(field.NewHandlInst(v))
}

//SetIDSource sets IDSource, Tag 22
func (m ExecutionReport) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetLastCapacity sets LastCapacity, Tag 29
func (m ExecutionReport) SetLastCapacity(v string) {
	m.Set(field.NewLastCapacity(v))
}

//SetLastMkt sets LastMkt, Tag 30
func (m ExecutionReport) SetLastMkt(v string) {
	m.Set(field.NewLastMkt(v))
}

//SetLastPx sets LastPx, Tag 31
func (m ExecutionReport) SetLastPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastPx(value, scale))
}

//SetLastShares sets LastShares, Tag 32
func (m ExecutionReport) SetLastShares(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastShares(value, scale))
}

//SetOrderID sets OrderID, Tag 37
func (m ExecutionReport) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m ExecutionReport) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrdStatus sets OrdStatus, Tag 39
func (m ExecutionReport) SetOrdStatus(v string) {
	m.Set(field.NewOrdStatus(v))
}

//SetOrdType sets OrdType, Tag 40
func (m ExecutionReport) SetOrdType(v string) {
	m.Set(field.NewOrdType(v))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m ExecutionReport) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetPrice sets Price, Tag 44
func (m ExecutionReport) SetPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewPrice(value, scale))
}

//SetRule80A sets Rule80A, Tag 47
func (m ExecutionReport) SetRule80A(v string) {
	m.Set(field.NewRule80A(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m ExecutionReport) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m ExecutionReport) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m ExecutionReport) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m ExecutionReport) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m ExecutionReport) SetTimeInForce(v string) {
	m.Set(field.NewTimeInForce(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m ExecutionReport) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSettlmntTyp sets SettlmntTyp, Tag 63
func (m ExecutionReport) SetSettlmntTyp(v string) {
	m.Set(field.NewSettlmntTyp(v))
}

//SetFutSettDate sets FutSettDate, Tag 64
func (m ExecutionReport) SetFutSettDate(v string) {
	m.Set(field.NewFutSettDate(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m ExecutionReport) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m ExecutionReport) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetTradeDate sets TradeDate, Tag 75
func (m ExecutionReport) SetTradeDate(v string) {
	m.Set(field.NewTradeDate(v))
}

//SetExecBroker sets ExecBroker, Tag 76
func (m ExecutionReport) SetExecBroker(v string) {
	m.Set(field.NewExecBroker(v))
}

//SetOpenClose sets OpenClose, Tag 77
func (m ExecutionReport) SetOpenClose(v string) {
	m.Set(field.NewOpenClose(v))
}

//SetStopPx sets StopPx, Tag 99
func (m ExecutionReport) SetStopPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewStopPx(value, scale))
}

//SetOrdRejReason sets OrdRejReason, Tag 103
func (m ExecutionReport) SetOrdRejReason(v int) {
	m.Set(field.NewOrdRejReason(v))
}

//SetIssuer sets Issuer, Tag 106
func (m ExecutionReport) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m ExecutionReport) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetClientID sets ClientID, Tag 109
func (m ExecutionReport) SetClientID(v string) {
	m.Set(field.NewClientID(v))
}

//SetMinQty sets MinQty, Tag 110
func (m ExecutionReport) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetMaxFloor sets MaxFloor, Tag 111
func (m ExecutionReport) SetMaxFloor(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxFloor(value, scale))
}

//SetReportToExch sets ReportToExch, Tag 113
func (m ExecutionReport) SetReportToExch(v bool) {
	m.Set(field.NewReportToExch(v))
}

//SetSettlCurrAmt sets SettlCurrAmt, Tag 119
func (m ExecutionReport) SetSettlCurrAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrAmt(value, scale))
}

//SetSettlCurrency sets SettlCurrency, Tag 120
func (m ExecutionReport) SetSettlCurrency(v string) {
	m.Set(field.NewSettlCurrency(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m ExecutionReport) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetExecType sets ExecType, Tag 150
func (m ExecutionReport) SetExecType(v string) {
	m.Set(field.NewExecType(v))
}

//SetLeavesQty sets LeavesQty, Tag 151
func (m ExecutionReport) SetLeavesQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewLeavesQty(value, scale))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m ExecutionReport) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSettlCurrFxRate sets SettlCurrFxRate, Tag 155
func (m ExecutionReport) SetSettlCurrFxRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewSettlCurrFxRate(value, scale))
}

//SetSettlCurrFxRateCalc sets SettlCurrFxRateCalc, Tag 156
func (m ExecutionReport) SetSettlCurrFxRateCalc(v string) {
	m.Set(field.NewSettlCurrFxRateCalc(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m ExecutionReport) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetEffectiveTime sets EffectiveTime, Tag 168
func (m ExecutionReport) SetEffectiveTime(v time.Time) {
	m.Set(field.NewEffectiveTime(v))
}

//SetOrderQty2 sets OrderQty2, Tag 192
func (m ExecutionReport) SetOrderQty2(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty2(value, scale))
}

//SetFutSettDate2 sets FutSettDate2, Tag 193
func (m ExecutionReport) SetFutSettDate2(v string) {
	m.Set(field.NewFutSettDate2(v))
}

//SetLastSpotRate sets LastSpotRate, Tag 194
func (m ExecutionReport) SetLastSpotRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastSpotRate(value, scale))
}

//SetLastForwardPoints sets LastForwardPoints, Tag 195
func (m ExecutionReport) SetLastForwardPoints(value decimal.Decimal, scale int32) {
	m.Set(field.NewLastForwardPoints(value, scale))
}

//SetSecondaryOrderID sets SecondaryOrderID, Tag 198
func (m ExecutionReport) SetSecondaryOrderID(v string) {
	m.Set(field.NewSecondaryOrderID(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m ExecutionReport) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m ExecutionReport) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m ExecutionReport) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m ExecutionReport) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m ExecutionReport) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m ExecutionReport) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetMaxShow sets MaxShow, Tag 210
func (m ExecutionReport) SetMaxShow(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxShow(value, scale))
}

//SetPegDifference sets PegDifference, Tag 211
func (m ExecutionReport) SetPegDifference(value decimal.Decimal, scale int32) {
	m.Set(field.NewPegDifference(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m ExecutionReport) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m ExecutionReport) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m ExecutionReport) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m ExecutionReport) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m ExecutionReport) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m ExecutionReport) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m ExecutionReport) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m ExecutionReport) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m ExecutionReport) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m ExecutionReport) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetSolicitedFlag sets SolicitedFlag, Tag 377
func (m ExecutionReport) SetSolicitedFlag(v bool) {
	m.Set(field.NewSolicitedFlag(v))
}

//SetExecRestatementReason sets ExecRestatementReason, Tag 378
func (m ExecutionReport) SetExecRestatementReason(v int) {
	m.Set(field.NewExecRestatementReason(v))
}

//SetGrossTradeAmt sets GrossTradeAmt, Tag 381
func (m ExecutionReport) SetGrossTradeAmt(value decimal.Decimal, scale int32) {
	m.Set(field.NewGrossTradeAmt(value, scale))
}

//SetNoContraBrokers sets NoContraBrokers, Tag 382
func (m ExecutionReport) SetNoContraBrokers(f NoContraBrokersRepeatingGroup) {
	m.SetGroup(f)
}

//SetDiscretionInst sets DiscretionInst, Tag 388
func (m ExecutionReport) SetDiscretionInst(v string) {
	m.Set(field.NewDiscretionInst(v))
}

//SetDiscretionOffset sets DiscretionOffset, Tag 389
func (m ExecutionReport) SetDiscretionOffset(value decimal.Decimal, scale int32) {
	m.Set(field.NewDiscretionOffset(value, scale))
}

//SetDayOrderQty sets DayOrderQty, Tag 424
func (m ExecutionReport) SetDayOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayOrderQty(value, scale))
}

//SetDayCumQty sets DayCumQty, Tag 425
func (m ExecutionReport) SetDayCumQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayCumQty(value, scale))
}

//SetDayAvgPx sets DayAvgPx, Tag 426
func (m ExecutionReport) SetDayAvgPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewDayAvgPx(value, scale))
}

//SetGTBookingInst sets GTBookingInst, Tag 427
func (m ExecutionReport) SetGTBookingInst(v int) {
	m.Set(field.NewGTBookingInst(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m ExecutionReport) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetClearingFirm sets ClearingFirm, Tag 439
func (m ExecutionReport) SetClearingFirm(v string) {
	m.Set(field.NewClearingFirm(v))
}

//SetClearingAccount sets ClearingAccount, Tag 440
func (m ExecutionReport) SetClearingAccount(v string) {
	m.Set(field.NewClearingAccount(v))
}

//SetMultiLegReportingType sets MultiLegReportingType, Tag 442
func (m ExecutionReport) SetMultiLegReportingType(v string) {
	m.Set(field.NewMultiLegReportingType(v))
}

//GetAccount gets Account, Tag 1
func (m ExecutionReport) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAvgPx gets AvgPx, Tag 6
func (m ExecutionReport) GetAvgPx() (f field.AvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m ExecutionReport) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommission gets Commission, Tag 12
func (m ExecutionReport) GetCommission() (f field.CommissionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCommType gets CommType, Tag 13
func (m ExecutionReport) GetCommType() (f field.CommTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCumQty gets CumQty, Tag 14
func (m ExecutionReport) GetCumQty() (f field.CumQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m ExecutionReport) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecID gets ExecID, Tag 17
func (m ExecutionReport) GetExecID() (f field.ExecIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m ExecutionReport) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRefID gets ExecRefID, Tag 19
func (m ExecutionReport) GetExecRefID() (f field.ExecRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecTransType gets ExecTransType, Tag 20
func (m ExecutionReport) GetExecTransType() (f field.ExecTransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetHandlInst gets HandlInst, Tag 21
func (m ExecutionReport) GetHandlInst() (f field.HandlInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m ExecutionReport) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastCapacity gets LastCapacity, Tag 29
func (m ExecutionReport) GetLastCapacity() (f field.LastCapacityField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastMkt gets LastMkt, Tag 30
func (m ExecutionReport) GetLastMkt() (f field.LastMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastPx gets LastPx, Tag 31
func (m ExecutionReport) GetLastPx() (f field.LastPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastShares gets LastShares, Tag 32
func (m ExecutionReport) GetLastShares() (f field.LastSharesField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m ExecutionReport) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m ExecutionReport) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdStatus gets OrdStatus, Tag 39
func (m ExecutionReport) GetOrdStatus() (f field.OrdStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdType gets OrdType, Tag 40
func (m ExecutionReport) GetOrdType() (f field.OrdTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m ExecutionReport) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m ExecutionReport) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRule80A gets Rule80A, Tag 47
func (m ExecutionReport) GetRule80A() (f field.Rule80AField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m ExecutionReport) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m ExecutionReport) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m ExecutionReport) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m ExecutionReport) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m ExecutionReport) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m ExecutionReport) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlmntTyp gets SettlmntTyp, Tag 63
func (m ExecutionReport) GetSettlmntTyp() (f field.SettlmntTypField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate gets FutSettDate, Tag 64
func (m ExecutionReport) GetFutSettDate() (f field.FutSettDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m ExecutionReport) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m ExecutionReport) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeDate gets TradeDate, Tag 75
func (m ExecutionReport) GetTradeDate() (f field.TradeDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecBroker gets ExecBroker, Tag 76
func (m ExecutionReport) GetExecBroker() (f field.ExecBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOpenClose gets OpenClose, Tag 77
func (m ExecutionReport) GetOpenClose() (f field.OpenCloseField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStopPx gets StopPx, Tag 99
func (m ExecutionReport) GetStopPx() (f field.StopPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrdRejReason gets OrdRejReason, Tag 103
func (m ExecutionReport) GetOrdRejReason() (f field.OrdRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m ExecutionReport) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m ExecutionReport) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClientID gets ClientID, Tag 109
func (m ExecutionReport) GetClientID() (f field.ClientIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m ExecutionReport) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxFloor gets MaxFloor, Tag 111
func (m ExecutionReport) GetMaxFloor() (f field.MaxFloorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetReportToExch gets ReportToExch, Tag 113
func (m ExecutionReport) GetReportToExch() (f field.ReportToExchField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrAmt gets SettlCurrAmt, Tag 119
func (m ExecutionReport) GetSettlCurrAmt() (f field.SettlCurrAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrency gets SettlCurrency, Tag 120
func (m ExecutionReport) GetSettlCurrency() (f field.SettlCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m ExecutionReport) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecType gets ExecType, Tag 150
func (m ExecutionReport) GetExecType() (f field.ExecTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLeavesQty gets LeavesQty, Tag 151
func (m ExecutionReport) GetLeavesQty() (f field.LeavesQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m ExecutionReport) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRate gets SettlCurrFxRate, Tag 155
func (m ExecutionReport) GetSettlCurrFxRate() (f field.SettlCurrFxRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSettlCurrFxRateCalc gets SettlCurrFxRateCalc, Tag 156
func (m ExecutionReport) GetSettlCurrFxRateCalc() (f field.SettlCurrFxRateCalcField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m ExecutionReport) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEffectiveTime gets EffectiveTime, Tag 168
func (m ExecutionReport) GetEffectiveTime() (f field.EffectiveTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty2 gets OrderQty2, Tag 192
func (m ExecutionReport) GetOrderQty2() (f field.OrderQty2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFutSettDate2 gets FutSettDate2, Tag 193
func (m ExecutionReport) GetFutSettDate2() (f field.FutSettDate2Field, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastSpotRate gets LastSpotRate, Tag 194
func (m ExecutionReport) GetLastSpotRate() (f field.LastSpotRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLastForwardPoints gets LastForwardPoints, Tag 195
func (m ExecutionReport) GetLastForwardPoints() (f field.LastForwardPointsField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryOrderID gets SecondaryOrderID, Tag 198
func (m ExecutionReport) GetSecondaryOrderID() (f field.SecondaryOrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m ExecutionReport) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m ExecutionReport) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m ExecutionReport) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m ExecutionReport) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m ExecutionReport) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m ExecutionReport) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaxShow gets MaxShow, Tag 210
func (m ExecutionReport) GetMaxShow() (f field.MaxShowField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPegDifference gets PegDifference, Tag 211
func (m ExecutionReport) GetPegDifference() (f field.PegDifferenceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m ExecutionReport) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m ExecutionReport) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m ExecutionReport) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m ExecutionReport) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m ExecutionReport) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m ExecutionReport) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m ExecutionReport) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m ExecutionReport) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m ExecutionReport) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m ExecutionReport) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSolicitedFlag gets SolicitedFlag, Tag 377
func (m ExecutionReport) GetSolicitedFlag() (f field.SolicitedFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecRestatementReason gets ExecRestatementReason, Tag 378
func (m ExecutionReport) GetExecRestatementReason() (f field.ExecRestatementReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGrossTradeAmt gets GrossTradeAmt, Tag 381
func (m ExecutionReport) GetGrossTradeAmt() (f field.GrossTradeAmtField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoContraBrokers gets NoContraBrokers, Tag 382
func (m ExecutionReport) GetNoContraBrokers() (NoContraBrokersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoContraBrokersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetDiscretionInst gets DiscretionInst, Tag 388
func (m ExecutionReport) GetDiscretionInst() (f field.DiscretionInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDiscretionOffset gets DiscretionOffset, Tag 389
func (m ExecutionReport) GetDiscretionOffset() (f field.DiscretionOffsetField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayOrderQty gets DayOrderQty, Tag 424
func (m ExecutionReport) GetDayOrderQty() (f field.DayOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayCumQty gets DayCumQty, Tag 425
func (m ExecutionReport) GetDayCumQty() (f field.DayCumQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDayAvgPx gets DayAvgPx, Tag 426
func (m ExecutionReport) GetDayAvgPx() (f field.DayAvgPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetGTBookingInst gets GTBookingInst, Tag 427
func (m ExecutionReport) GetGTBookingInst() (f field.GTBookingInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m ExecutionReport) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingFirm gets ClearingFirm, Tag 439
func (m ExecutionReport) GetClearingFirm() (f field.ClearingFirmField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClearingAccount gets ClearingAccount, Tag 440
func (m ExecutionReport) GetClearingAccount() (f field.ClearingAccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMultiLegReportingType gets MultiLegReportingType, Tag 442
func (m ExecutionReport) GetMultiLegReportingType() (f field.MultiLegReportingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m ExecutionReport) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasAvgPx returns true if AvgPx is present, Tag 6
func (m ExecutionReport) HasAvgPx() bool {
	return m.Has(tag.AvgPx)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m ExecutionReport) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasCommission returns true if Commission is present, Tag 12
func (m ExecutionReport) HasCommission() bool {
	return m.Has(tag.Commission)
}

//HasCommType returns true if CommType is present, Tag 13
func (m ExecutionReport) HasCommType() bool {
	return m.Has(tag.CommType)
}

//HasCumQty returns true if CumQty is present, Tag 14
func (m ExecutionReport) HasCumQty() bool {
	return m.Has(tag.CumQty)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m ExecutionReport) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasExecID returns true if ExecID is present, Tag 17
func (m ExecutionReport) HasExecID() bool {
	return m.Has(tag.ExecID)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m ExecutionReport) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasExecRefID returns true if ExecRefID is present, Tag 19
func (m ExecutionReport) HasExecRefID() bool {
	return m.Has(tag.ExecRefID)
}

//HasExecTransType returns true if ExecTransType is present, Tag 20
func (m ExecutionReport) HasExecTransType() bool {
	return m.Has(tag.ExecTransType)
}

//HasHandlInst returns true if HandlInst is present, Tag 21
func (m ExecutionReport) HasHandlInst() bool {
	return m.Has(tag.HandlInst)
}

//HasIDSource returns true if IDSource is present, Tag 22
func (m ExecutionReport) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasLastCapacity returns true if LastCapacity is present, Tag 29
func (m ExecutionReport) HasLastCapacity() bool {
	return m.Has(tag.LastCapacity)
}

//HasLastMkt returns true if LastMkt is present, Tag 30
func (m ExecutionReport) HasLastMkt() bool {
	return m.Has(tag.LastMkt)
}

//HasLastPx returns true if LastPx is present, Tag 31
func (m ExecutionReport) HasLastPx() bool {
	return m.Has(tag.LastPx)
}

//HasLastShares returns true if LastShares is present, Tag 32
func (m ExecutionReport) HasLastShares() bool {
	return m.Has(tag.LastShares)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m ExecutionReport) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m ExecutionReport) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrdStatus returns true if OrdStatus is present, Tag 39
func (m ExecutionReport) HasOrdStatus() bool {
	return m.Has(tag.OrdStatus)
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m ExecutionReport) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m ExecutionReport) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasPrice returns true if Price is present, Tag 44
func (m ExecutionReport) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasRule80A returns true if Rule80A is present, Tag 47
func (m ExecutionReport) HasRule80A() bool {
	return m.Has(tag.Rule80A)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m ExecutionReport) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m ExecutionReport) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m ExecutionReport) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m ExecutionReport) HasText() bool {
	return m.Has(tag.Text)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m ExecutionReport) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m ExecutionReport) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSettlmntTyp returns true if SettlmntTyp is present, Tag 63
func (m ExecutionReport) HasSettlmntTyp() bool {
	return m.Has(tag.SettlmntTyp)
}

//HasFutSettDate returns true if FutSettDate is present, Tag 64
func (m ExecutionReport) HasFutSettDate() bool {
	return m.Has(tag.FutSettDate)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m ExecutionReport) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m ExecutionReport) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasTradeDate returns true if TradeDate is present, Tag 75
func (m ExecutionReport) HasTradeDate() bool {
	return m.Has(tag.TradeDate)
}

//HasExecBroker returns true if ExecBroker is present, Tag 76
func (m ExecutionReport) HasExecBroker() bool {
	return m.Has(tag.ExecBroker)
}

//HasOpenClose returns true if OpenClose is present, Tag 77
func (m ExecutionReport) HasOpenClose() bool {
	return m.Has(tag.OpenClose)
}

//HasStopPx returns true if StopPx is present, Tag 99
func (m ExecutionReport) HasStopPx() bool {
	return m.Has(tag.StopPx)
}

//HasOrdRejReason returns true if OrdRejReason is present, Tag 103
func (m ExecutionReport) HasOrdRejReason() bool {
	return m.Has(tag.OrdRejReason)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m ExecutionReport) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m ExecutionReport) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasClientID returns true if ClientID is present, Tag 109
func (m ExecutionReport) HasClientID() bool {
	return m.Has(tag.ClientID)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m ExecutionReport) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasMaxFloor returns true if MaxFloor is present, Tag 111
func (m ExecutionReport) HasMaxFloor() bool {
	return m.Has(tag.MaxFloor)
}

//HasReportToExch returns true if ReportToExch is present, Tag 113
func (m ExecutionReport) HasReportToExch() bool {
	return m.Has(tag.ReportToExch)
}

//HasSettlCurrAmt returns true if SettlCurrAmt is present, Tag 119
func (m ExecutionReport) HasSettlCurrAmt() bool {
	return m.Has(tag.SettlCurrAmt)
}

//HasSettlCurrency returns true if SettlCurrency is present, Tag 120
func (m ExecutionReport) HasSettlCurrency() bool {
	return m.Has(tag.SettlCurrency)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m ExecutionReport) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasExecType returns true if ExecType is present, Tag 150
func (m ExecutionReport) HasExecType() bool {
	return m.Has(tag.ExecType)
}

//HasLeavesQty returns true if LeavesQty is present, Tag 151
func (m ExecutionReport) HasLeavesQty() bool {
	return m.Has(tag.LeavesQty)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m ExecutionReport) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSettlCurrFxRate returns true if SettlCurrFxRate is present, Tag 155
func (m ExecutionReport) HasSettlCurrFxRate() bool {
	return m.Has(tag.SettlCurrFxRate)
}

//HasSettlCurrFxRateCalc returns true if SettlCurrFxRateCalc is present, Tag 156
func (m ExecutionReport) HasSettlCurrFxRateCalc() bool {
	return m.Has(tag.SettlCurrFxRateCalc)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m ExecutionReport) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasEffectiveTime returns true if EffectiveTime is present, Tag 168
func (m ExecutionReport) HasEffectiveTime() bool {
	return m.Has(tag.EffectiveTime)
}

//HasOrderQty2 returns true if OrderQty2 is present, Tag 192
func (m ExecutionReport) HasOrderQty2() bool {
	return m.Has(tag.OrderQty2)
}

//HasFutSettDate2 returns true if FutSettDate2 is present, Tag 193
func (m ExecutionReport) HasFutSettDate2() bool {
	return m.Has(tag.FutSettDate2)
}

//HasLastSpotRate returns true if LastSpotRate is present, Tag 194
func (m ExecutionReport) HasLastSpotRate() bool {
	return m.Has(tag.LastSpotRate)
}

//HasLastForwardPoints returns true if LastForwardPoints is present, Tag 195
func (m ExecutionReport) HasLastForwardPoints() bool {
	return m.Has(tag.LastForwardPoints)
}

//HasSecondaryOrderID returns true if SecondaryOrderID is present, Tag 198
func (m ExecutionReport) HasSecondaryOrderID() bool {
	return m.Has(tag.SecondaryOrderID)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m ExecutionReport) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m ExecutionReport) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m ExecutionReport) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m ExecutionReport) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m ExecutionReport) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m ExecutionReport) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasMaxShow returns true if MaxShow is present, Tag 210
func (m ExecutionReport) HasMaxShow() bool {
	return m.Has(tag.MaxShow)
}

//HasPegDifference returns true if PegDifference is present, Tag 211
func (m ExecutionReport) HasPegDifference() bool {
	return m.Has(tag.PegDifference)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m ExecutionReport) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m ExecutionReport) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m ExecutionReport) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m ExecutionReport) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m ExecutionReport) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m ExecutionReport) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m ExecutionReport) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m ExecutionReport) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m ExecutionReport) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m ExecutionReport) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasSolicitedFlag returns true if SolicitedFlag is present, Tag 377
func (m ExecutionReport) HasSolicitedFlag() bool {
	return m.Has(tag.SolicitedFlag)
}

//HasExecRestatementReason returns true if ExecRestatementReason is present, Tag 378
func (m ExecutionReport) HasExecRestatementReason() bool {
	return m.Has(tag.ExecRestatementReason)
}

//HasGrossTradeAmt returns true if GrossTradeAmt is present, Tag 381
func (m ExecutionReport) HasGrossTradeAmt() bool {
	return m.Has(tag.GrossTradeAmt)
}

//HasNoContraBrokers returns true if NoContraBrokers is present, Tag 382
func (m ExecutionReport) HasNoContraBrokers() bool {
	return m.Has(tag.NoContraBrokers)
}

//HasDiscretionInst returns true if DiscretionInst is present, Tag 388
func (m ExecutionReport) HasDiscretionInst() bool {
	return m.Has(tag.DiscretionInst)
}

//HasDiscretionOffset returns true if DiscretionOffset is present, Tag 389
func (m ExecutionReport) HasDiscretionOffset() bool {
	return m.Has(tag.DiscretionOffset)
}

//HasDayOrderQty returns true if DayOrderQty is present, Tag 424
func (m ExecutionReport) HasDayOrderQty() bool {
	return m.Has(tag.DayOrderQty)
}

//HasDayCumQty returns true if DayCumQty is present, Tag 425
func (m ExecutionReport) HasDayCumQty() bool {
	return m.Has(tag.DayCumQty)
}

//HasDayAvgPx returns true if DayAvgPx is present, Tag 426
func (m ExecutionReport) HasDayAvgPx() bool {
	return m.Has(tag.DayAvgPx)
}

//HasGTBookingInst returns true if GTBookingInst is present, Tag 427
func (m ExecutionReport) HasGTBookingInst() bool {
	return m.Has(tag.GTBookingInst)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m ExecutionReport) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasClearingFirm returns true if ClearingFirm is present, Tag 439
func (m ExecutionReport) HasClearingFirm() bool {
	return m.Has(tag.ClearingFirm)
}

//HasClearingAccount returns true if ClearingAccount is present, Tag 440
func (m ExecutionReport) HasClearingAccount() bool {
	return m.Has(tag.ClearingAccount)
}

//HasMultiLegReportingType returns true if MultiLegReportingType is present, Tag 442
func (m ExecutionReport) HasMultiLegReportingType() bool {
	return m.Has(tag.MultiLegReportingType)
}

//NoContraBrokers is a repeating group element, Tag 382
type NoContraBrokers struct {
	quickfix.Group
}

//SetContraBroker sets ContraBroker, Tag 375
func (m NoContraBrokers) SetContraBroker(v string) {
	m.Set(field.NewContraBroker(v))
}

//SetContraTrader sets ContraTrader, Tag 337
func (m NoContraBrokers) SetContraTrader(v string) {
	m.Set(field.NewContraTrader(v))
}

//SetContraTradeQty sets ContraTradeQty, Tag 437
func (m NoContraBrokers) SetContraTradeQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewContraTradeQty(value, scale))
}

//SetContraTradeTime sets ContraTradeTime, Tag 438
func (m NoContraBrokers) SetContraTradeTime(v time.Time) {
	m.Set(field.NewContraTradeTime(v))
}

//GetContraBroker gets ContraBroker, Tag 375
func (m NoContraBrokers) GetContraBroker() (f field.ContraBrokerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContraTrader gets ContraTrader, Tag 337
func (m NoContraBrokers) GetContraTrader() (f field.ContraTraderField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContraTradeQty gets ContraTradeQty, Tag 437
func (m NoContraBrokers) GetContraTradeQty() (f field.ContraTradeQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContraTradeTime gets ContraTradeTime, Tag 438
func (m NoContraBrokers) GetContraTradeTime() (f field.ContraTradeTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasContraBroker returns true if ContraBroker is present, Tag 375
func (m NoContraBrokers) HasContraBroker() bool {
	return m.Has(tag.ContraBroker)
}

//HasContraTrader returns true if ContraTrader is present, Tag 337
func (m NoContraBrokers) HasContraTrader() bool {
	return m.Has(tag.ContraTrader)
}

//HasContraTradeQty returns true if ContraTradeQty is present, Tag 437
func (m NoContraBrokers) HasContraTradeQty() bool {
	return m.Has(tag.ContraTradeQty)
}

//HasContraTradeTime returns true if ContraTradeTime is present, Tag 438
func (m NoContraBrokers) HasContraTradeTime() bool {
	return m.Has(tag.ContraTradeTime)
}

//NoContraBrokersRepeatingGroup is a repeating group, Tag 382
type NoContraBrokersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoContraBrokersRepeatingGroup returns an initialized, NoContraBrokersRepeatingGroup
func NewNoContraBrokersRepeatingGroup() NoContraBrokersRepeatingGroup {
	return NoContraBrokersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoContraBrokers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ContraBroker), quickfix.GroupElement(tag.ContraTrader), quickfix.GroupElement(tag.ContraTradeQty), quickfix.GroupElement(tag.ContraTradeTime)})}
}

//Add create and append a new NoContraBrokers to this group
func (m NoContraBrokersRepeatingGroup) Add() NoContraBrokers {
	g := m.RepeatingGroup.Add()
	return NoContraBrokers{g}
}

//Get returns the ith NoContraBrokers in the NoContraBrokersRepeatinGroup
func (m NoContraBrokersRepeatingGroup) Get(i int) NoContraBrokers {
	return NoContraBrokers{m.RepeatingGroup.Get(i)}
}
