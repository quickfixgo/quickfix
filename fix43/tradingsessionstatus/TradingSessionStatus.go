package tradingsessionstatus

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionStatus is the fix43 TradingSessionStatus type, MsgType = h
type TradingSessionStatus struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
}

//FromMessage creates a TradingSessionStatus from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradingSessionStatus {
	return TradingSessionStatus{
		Header:  fix43.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix43.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionStatus) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a TradingSessionStatus initialized with the required fields for TradingSessionStatus
func New(tradingsessionid field.TradingSessionIDField, tradsesstatus field.TradSesStatusField) (m TradingSessionStatus) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("h"))
	m.Set(tradingsessionid)
	m.Set(tradsesstatus)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionStatus, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "h", r
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
func (m TradingSessionStatus) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m TradingSessionStatus) SetTradSesMethod(v int) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m TradingSessionStatus) SetTradSesMode(v int) {
	m.Set(field.NewTradSesMode(v))
}

//SetTradSesStatus sets TradSesStatus, Tag 340
func (m TradingSessionStatus) SetTradSesStatus(v int) {
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
func (m TradingSessionStatus) SetTotalVolumeTraded(v float64) {
	m.Set(field.NewTotalVolumeTraded(v))
}

//SetTradSesStatusRejReason sets TradSesStatusRejReason, Tag 567
func (m TradingSessionStatus) SetTradSesStatusRejReason(v int) {
	m.Set(field.NewTradSesStatusRejReason(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m TradingSessionStatus) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//GetText gets Text, Tag 58
func (m TradingSessionStatus) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m TradingSessionStatus) GetUnsolicitedIndicator() (f field.UnsolicitedIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionStatus) GetTradSesReqID() (f field.TradSesReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m TradingSessionStatus) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m TradingSessionStatus) GetTradSesMethod() (f field.TradSesMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m TradingSessionStatus) GetTradSesMode() (f field.TradSesModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStatus gets TradSesStatus, Tag 340
func (m TradingSessionStatus) GetTradSesStatus() (f field.TradSesStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStartTime gets TradSesStartTime, Tag 341
func (m TradingSessionStatus) GetTradSesStartTime() (f field.TradSesStartTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesOpenTime gets TradSesOpenTime, Tag 342
func (m TradingSessionStatus) GetTradSesOpenTime() (f field.TradSesOpenTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343
func (m TradingSessionStatus) GetTradSesPreCloseTime() (f field.TradSesPreCloseTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesCloseTime gets TradSesCloseTime, Tag 344
func (m TradingSessionStatus) GetTradSesCloseTime() (f field.TradSesCloseTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesEndTime gets TradSesEndTime, Tag 345
func (m TradingSessionStatus) GetTradSesEndTime() (f field.TradSesEndTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m TradingSessionStatus) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m TradingSessionStatus) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m TradingSessionStatus) GetTotalVolumeTraded() (f field.TotalVolumeTradedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStatusRejReason gets TradSesStatusRejReason, Tag 567
func (m TradingSessionStatus) GetTradSesStatusRejReason() (f field.TradSesStatusRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m TradingSessionStatus) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//HasTradSesStatusRejReason returns true if TradSesStatusRejReason is present, Tag 567
func (m TradingSessionStatus) HasTradSesStatusRejReason() bool {
	return m.Has(tag.TradSesStatusRejReason)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m TradingSessionStatus) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}
