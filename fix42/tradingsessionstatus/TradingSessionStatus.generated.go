package tradingsessionstatus

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionStatus is the fix42 TradingSessionStatus type, MsgType = h
type TradingSessionStatus struct {
	fix42.Header
	*quickfix.Body
	fix42.Trailer
	Message *quickfix.Message
}

//FromMessage creates a TradingSessionStatus from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TradingSessionStatus {
	return TradingSessionStatus{
		Header:  fix42.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix42.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionStatus) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a TradingSessionStatus initialized with the required fields for TradingSessionStatus
func New(tradingsessionid field.TradingSessionIDField, tradsesstatus field.TradSesStatusField) (m TradingSessionStatus) {
	m.Message = quickfix.NewMessage()
	m.Header = fix42.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("h"))
	m.Set(tradingsessionid)
	m.Set(tradsesstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "h", r
}

//SetText sets Text, Tag 58
func (m TradingSessionStatus) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m TradingSessionStatus) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionStatus) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m TradingSessionStatus) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionStatus) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionStatus) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

//SetTradSesStatus sets TradSesStatus, Tag 340
func (m TradingSessionStatus) SetTradSesStatus(v enum.TradSesStatus) {
	m.Set(field.NewTradSesStatus(v))
}

//SetTradSesStartTime sets TradSesStartTime, Tag 341
func (m TradingSessionStatus) SetTradSesStartTime(v time.Time) {
	m.Set(field.NewTradSesStartTime(v))
}

//SetTradSesOpenTime sets TradSesOpenTime, Tag 342
func (m TradingSessionStatus) SetTradSesOpenTime(v time.Time) {
	m.Set(field.NewTradSesOpenTime(v))
}

//SetTradSesPreCloseTime sets TradSesPreCloseTime, Tag 343
func (m TradingSessionStatus) SetTradSesPreCloseTime(v time.Time) {
	m.Set(field.NewTradSesPreCloseTime(v))
}

//SetTradSesCloseTime sets TradSesCloseTime, Tag 344
func (m TradingSessionStatus) SetTradSesCloseTime(v time.Time) {
	m.Set(field.NewTradSesCloseTime(v))
}

//SetTradSesEndTime sets TradSesEndTime, Tag 345
func (m TradingSessionStatus) SetTradSesEndTime(v time.Time) {
	m.Set(field.NewTradSesEndTime(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m TradingSessionStatus) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m TradingSessionStatus) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387
func (m TradingSessionStatus) SetTotalVolumeTraded(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalVolumeTraded(value, scale))
}

//GetText gets Text, Tag 58
func (m TradingSessionStatus) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m TradingSessionStatus) GetUnsolicitedIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.UnsolicitedIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionStatus) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionStatus) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionStatus) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionStatus) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesStatus gets TradSesStatus, Tag 340
func (m TradingSessionStatus) GetTradSesStatus() (v enum.TradSesStatus, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesStartTime gets TradSesStartTime, Tag 341
func (m TradingSessionStatus) GetTradSesStartTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesOpenTime gets TradSesOpenTime, Tag 342
func (m TradingSessionStatus) GetTradSesOpenTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesOpenTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343
func (m TradingSessionStatus) GetTradSesPreCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesPreCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesCloseTime gets TradSesCloseTime, Tag 344
func (m TradingSessionStatus) GetTradSesCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesEndTime gets TradSesEndTime, Tag 345
func (m TradingSessionStatus) GetTradSesEndTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m TradingSessionStatus) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m TradingSessionStatus) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m TradingSessionStatus) GetTotalVolumeTraded() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasText returns true if Text is present, Tag 58
func (m TradingSessionStatus) HasText() bool {
	return m.Has(tag.Text)
}

//HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m TradingSessionStatus) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

//HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionStatus) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m TradingSessionStatus) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradSesMethod returns true if TradSesMethod is present, Tag 338
func (m TradingSessionStatus) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

//HasTradSesMode returns true if TradSesMode is present, Tag 339
func (m TradingSessionStatus) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

//HasTradSesStatus returns true if TradSesStatus is present, Tag 340
func (m TradingSessionStatus) HasTradSesStatus() bool {
	return m.Has(tag.TradSesStatus)
}

//HasTradSesStartTime returns true if TradSesStartTime is present, Tag 341
func (m TradingSessionStatus) HasTradSesStartTime() bool {
	return m.Has(tag.TradSesStartTime)
}

//HasTradSesOpenTime returns true if TradSesOpenTime is present, Tag 342
func (m TradingSessionStatus) HasTradSesOpenTime() bool {
	return m.Has(tag.TradSesOpenTime)
}

//HasTradSesPreCloseTime returns true if TradSesPreCloseTime is present, Tag 343
func (m TradingSessionStatus) HasTradSesPreCloseTime() bool {
	return m.Has(tag.TradSesPreCloseTime)
}

//HasTradSesCloseTime returns true if TradSesCloseTime is present, Tag 344
func (m TradingSessionStatus) HasTradSesCloseTime() bool {
	return m.Has(tag.TradSesCloseTime)
}

//HasTradSesEndTime returns true if TradSesEndTime is present, Tag 345
func (m TradingSessionStatus) HasTradSesEndTime() bool {
	return m.Has(tag.TradSesEndTime)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m TradingSessionStatus) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m TradingSessionStatus) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387
func (m TradingSessionStatus) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
}
