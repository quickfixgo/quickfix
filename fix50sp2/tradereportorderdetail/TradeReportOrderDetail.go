package tradereportorderdetail

import (
	"time"
)

//TradeReportOrderDetail is a fix50sp2 Component
type TradeReportOrderDetail struct {
	//OrderID is a non-required field for TradeReportOrderDetail.
	OrderID *string `fix:"37"`
	//SecondaryOrderID is a non-required field for TradeReportOrderDetail.
	SecondaryOrderID *string `fix:"198"`
	//ClOrdID is a non-required field for TradeReportOrderDetail.
	ClOrdID *string `fix:"11"`
	//SecondaryClOrdID is a non-required field for TradeReportOrderDetail.
	SecondaryClOrdID *string `fix:"526"`
	//ListID is a non-required field for TradeReportOrderDetail.
	ListID *string `fix:"66"`
	//RefOrderID is a non-required field for TradeReportOrderDetail.
	RefOrderID *string `fix:"1080"`
	//RefOrderIDSource is a non-required field for TradeReportOrderDetail.
	RefOrderIDSource *string `fix:"1081"`
	//RefOrdIDReason is a non-required field for TradeReportOrderDetail.
	RefOrdIDReason *int `fix:"1431"`
	//OrdType is a non-required field for TradeReportOrderDetail.
	OrdType *string `fix:"40"`
	//Price is a non-required field for TradeReportOrderDetail.
	Price *float64 `fix:"44"`
	//StopPx is a non-required field for TradeReportOrderDetail.
	StopPx *float64 `fix:"99"`
	//ExecInst is a non-required field for TradeReportOrderDetail.
	ExecInst *string `fix:"18"`
	//OrdStatus is a non-required field for TradeReportOrderDetail.
	OrdStatus *string `fix:"39"`
	//OrderQty is a non-required field for TradeReportOrderDetail.
	OrderQty *float64 `fix:"38"`
	//CashOrderQty is a non-required field for TradeReportOrderDetail.
	CashOrderQty *float64 `fix:"152"`
	//OrderPercent is a non-required field for TradeReportOrderDetail.
	OrderPercent *float64 `fix:"516"`
	//RoundingDirection is a non-required field for TradeReportOrderDetail.
	RoundingDirection *string `fix:"468"`
	//RoundingModulus is a non-required field for TradeReportOrderDetail.
	RoundingModulus *float64 `fix:"469"`
	//LeavesQty is a non-required field for TradeReportOrderDetail.
	LeavesQty *float64 `fix:"151"`
	//CumQty is a non-required field for TradeReportOrderDetail.
	CumQty *float64 `fix:"14"`
	//TimeInForce is a non-required field for TradeReportOrderDetail.
	TimeInForce *string `fix:"59"`
	//ExpireTime is a non-required field for TradeReportOrderDetail.
	ExpireTime *time.Time `fix:"126"`
	//SecondaryDisplayQty is a non-required field for TradeReportOrderDetail.
	SecondaryDisplayQty *float64 `fix:"1082"`
	//DisplayWhen is a non-required field for TradeReportOrderDetail.
	DisplayWhen *string `fix:"1083"`
	//DisplayMethod is a non-required field for TradeReportOrderDetail.
	DisplayMethod *string `fix:"1084"`
	//DisplayLowQty is a non-required field for TradeReportOrderDetail.
	DisplayLowQty *float64 `fix:"1085"`
	//DisplayHighQty is a non-required field for TradeReportOrderDetail.
	DisplayHighQty *float64 `fix:"1086"`
	//DisplayMinIncr is a non-required field for TradeReportOrderDetail.
	DisplayMinIncr *float64 `fix:"1087"`
	//RefreshQty is a non-required field for TradeReportOrderDetail.
	RefreshQty *float64 `fix:"1088"`
	//DisplayQty is a non-required field for TradeReportOrderDetail.
	DisplayQty *float64 `fix:"1138"`
	//OrderCapacity is a non-required field for TradeReportOrderDetail.
	OrderCapacity *string `fix:"528"`
	//OrderRestrictions is a non-required field for TradeReportOrderDetail.
	OrderRestrictions *string `fix:"529"`
	//OrigCustOrderCapacity is a non-required field for TradeReportOrderDetail.
	OrigCustOrderCapacity *int `fix:"1432"`
	//OrderInputDevice is a non-required field for TradeReportOrderDetail.
	OrderInputDevice *string `fix:"821"`
	//LotType is a non-required field for TradeReportOrderDetail.
	LotType *string `fix:"1093"`
	//TransBkdTime is a non-required field for TradeReportOrderDetail.
	TransBkdTime *time.Time `fix:"483"`
	//OrigOrdModTime is a non-required field for TradeReportOrderDetail.
	OrigOrdModTime *time.Time `fix:"586"`
	//BookingType is a non-required field for TradeReportOrderDetail.
	BookingType *int `fix:"775"`
}

func (m *TradeReportOrderDetail) SetOrderID(v string)              { m.OrderID = &v }
func (m *TradeReportOrderDetail) SetSecondaryOrderID(v string)     { m.SecondaryOrderID = &v }
func (m *TradeReportOrderDetail) SetClOrdID(v string)              { m.ClOrdID = &v }
func (m *TradeReportOrderDetail) SetSecondaryClOrdID(v string)     { m.SecondaryClOrdID = &v }
func (m *TradeReportOrderDetail) SetListID(v string)               { m.ListID = &v }
func (m *TradeReportOrderDetail) SetRefOrderID(v string)           { m.RefOrderID = &v }
func (m *TradeReportOrderDetail) SetRefOrderIDSource(v string)     { m.RefOrderIDSource = &v }
func (m *TradeReportOrderDetail) SetRefOrdIDReason(v int)          { m.RefOrdIDReason = &v }
func (m *TradeReportOrderDetail) SetOrdType(v string)              { m.OrdType = &v }
func (m *TradeReportOrderDetail) SetPrice(v float64)               { m.Price = &v }
func (m *TradeReportOrderDetail) SetStopPx(v float64)              { m.StopPx = &v }
func (m *TradeReportOrderDetail) SetExecInst(v string)             { m.ExecInst = &v }
func (m *TradeReportOrderDetail) SetOrdStatus(v string)            { m.OrdStatus = &v }
func (m *TradeReportOrderDetail) SetOrderQty(v float64)            { m.OrderQty = &v }
func (m *TradeReportOrderDetail) SetCashOrderQty(v float64)        { m.CashOrderQty = &v }
func (m *TradeReportOrderDetail) SetOrderPercent(v float64)        { m.OrderPercent = &v }
func (m *TradeReportOrderDetail) SetRoundingDirection(v string)    { m.RoundingDirection = &v }
func (m *TradeReportOrderDetail) SetRoundingModulus(v float64)     { m.RoundingModulus = &v }
func (m *TradeReportOrderDetail) SetLeavesQty(v float64)           { m.LeavesQty = &v }
func (m *TradeReportOrderDetail) SetCumQty(v float64)              { m.CumQty = &v }
func (m *TradeReportOrderDetail) SetTimeInForce(v string)          { m.TimeInForce = &v }
func (m *TradeReportOrderDetail) SetExpireTime(v time.Time)        { m.ExpireTime = &v }
func (m *TradeReportOrderDetail) SetSecondaryDisplayQty(v float64) { m.SecondaryDisplayQty = &v }
func (m *TradeReportOrderDetail) SetDisplayWhen(v string)          { m.DisplayWhen = &v }
func (m *TradeReportOrderDetail) SetDisplayMethod(v string)        { m.DisplayMethod = &v }
func (m *TradeReportOrderDetail) SetDisplayLowQty(v float64)       { m.DisplayLowQty = &v }
func (m *TradeReportOrderDetail) SetDisplayHighQty(v float64)      { m.DisplayHighQty = &v }
func (m *TradeReportOrderDetail) SetDisplayMinIncr(v float64)      { m.DisplayMinIncr = &v }
func (m *TradeReportOrderDetail) SetRefreshQty(v float64)          { m.RefreshQty = &v }
func (m *TradeReportOrderDetail) SetDisplayQty(v float64)          { m.DisplayQty = &v }
func (m *TradeReportOrderDetail) SetOrderCapacity(v string)        { m.OrderCapacity = &v }
func (m *TradeReportOrderDetail) SetOrderRestrictions(v string)    { m.OrderRestrictions = &v }
func (m *TradeReportOrderDetail) SetOrigCustOrderCapacity(v int)   { m.OrigCustOrderCapacity = &v }
func (m *TradeReportOrderDetail) SetOrderInputDevice(v string)     { m.OrderInputDevice = &v }
func (m *TradeReportOrderDetail) SetLotType(v string)              { m.LotType = &v }
func (m *TradeReportOrderDetail) SetTransBkdTime(v time.Time)      { m.TransBkdTime = &v }
func (m *TradeReportOrderDetail) SetOrigOrdModTime(v time.Time)    { m.OrigOrdModTime = &v }
func (m *TradeReportOrderDetail) SetBookingType(v int)             { m.BookingType = &v }
