package tradingsessionlist

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionList is the fix50 TradingSessionList type, MsgType = BJ
type TradingSessionList struct {
	fixt11.Header
	quickfix.Body
	fixt11.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a TradingSessionList from a quickfix.Message instance
func FromMessage(m quickfix.Message) TradingSessionList {
	return TradingSessionList{
		Header:      fixt11.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fixt11.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionList) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a TradingSessionList initialized with the required fields for TradingSessionList
func New() (m TradingSessionList) {
	m.Header = fixt11.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("BJ"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionList, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "7", "BJ", r
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionList) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m TradingSessionList) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionList) GetTradSesReqID() (f field.TradSesReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m TradingSessionList) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionList) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m TradingSessionList) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v string) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoTradingSessions) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m NoTradingSessions) SetTradSesMethod(v int) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m NoTradingSessions) SetTradSesMode(v int) {
	m.Set(field.NewTradSesMode(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m NoTradingSessions) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetTradSesStatus sets TradSesStatus, Tag 340
func (m NoTradingSessions) SetTradSesStatus(v int) {
	m.Set(field.NewTradSesStatus(v))
}

//SetTradSesStatusRejReason sets TradSesStatusRejReason, Tag 567
func (m NoTradingSessions) SetTradSesStatusRejReason(v int) {
	m.Set(field.NewTradSesStatusRejReason(v))
}

//SetTradSesStartTime sets TradSesStartTime, Tag 341
func (m NoTradingSessions) SetTradSesStartTime(v time.Time) {
	m.Set(field.NewTradSesStartTime(v))
}

//SetTradSesOpenTime sets TradSesOpenTime, Tag 342
func (m NoTradingSessions) SetTradSesOpenTime(v time.Time) {
	m.Set(field.NewTradSesOpenTime(v))
}

//SetTradSesPreCloseTime sets TradSesPreCloseTime, Tag 343
func (m NoTradingSessions) SetTradSesPreCloseTime(v time.Time) {
	m.Set(field.NewTradSesPreCloseTime(v))
}

//SetTradSesCloseTime sets TradSesCloseTime, Tag 344
func (m NoTradingSessions) SetTradSesCloseTime(v time.Time) {
	m.Set(field.NewTradSesCloseTime(v))
}

//SetTradSesEndTime sets TradSesEndTime, Tag 345
func (m NoTradingSessions) SetTradSesEndTime(v time.Time) {
	m.Set(field.NewTradSesEndTime(v))
}

//SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387
func (m NoTradingSessions) SetTotalVolumeTraded(v float64) {
	m.Set(field.NewTotalVolumeTraded(v))
}

//SetText sets Text, Tag 58
func (m NoTradingSessions) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoTradingSessions) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoTradingSessions) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (f field.TradingSessionSubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoTradingSessions) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m NoTradingSessions) GetTradSesMethod() (f field.TradSesMethodField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m NoTradingSessions) GetTradSesMode() (f field.TradSesModeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m NoTradingSessions) GetUnsolicitedIndicator() (f field.UnsolicitedIndicatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStatus gets TradSesStatus, Tag 340
func (m NoTradingSessions) GetTradSesStatus() (f field.TradSesStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStatusRejReason gets TradSesStatusRejReason, Tag 567
func (m NoTradingSessions) GetTradSesStatusRejReason() (f field.TradSesStatusRejReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesStartTime gets TradSesStartTime, Tag 341
func (m NoTradingSessions) GetTradSesStartTime() (f field.TradSesStartTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesOpenTime gets TradSesOpenTime, Tag 342
func (m NoTradingSessions) GetTradSesOpenTime() (f field.TradSesOpenTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343
func (m NoTradingSessions) GetTradSesPreCloseTime() (f field.TradSesPreCloseTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesCloseTime gets TradSesCloseTime, Tag 344
func (m NoTradingSessions) GetTradSesCloseTime() (f field.TradSesCloseTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradSesEndTime gets TradSesEndTime, Tag 345
func (m NoTradingSessions) GetTradSesEndTime() (f field.TradSesEndTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m NoTradingSessions) GetTotalVolumeTraded() (f field.TotalVolumeTradedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoTradingSessions) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoTradingSessions) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoTradingSessions) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoTradingSessions) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoTradingSessions) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoTradingSessions) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasTradSesMethod returns true if TradSesMethod is present, Tag 338
func (m NoTradingSessions) HasTradSesMethod() bool {
	return m.Has(tag.TradSesMethod)
}

//HasTradSesMode returns true if TradSesMode is present, Tag 339
func (m NoTradingSessions) HasTradSesMode() bool {
	return m.Has(tag.TradSesMode)
}

//HasUnsolicitedIndicator returns true if UnsolicitedIndicator is present, Tag 325
func (m NoTradingSessions) HasUnsolicitedIndicator() bool {
	return m.Has(tag.UnsolicitedIndicator)
}

//HasTradSesStatus returns true if TradSesStatus is present, Tag 340
func (m NoTradingSessions) HasTradSesStatus() bool {
	return m.Has(tag.TradSesStatus)
}

//HasTradSesStatusRejReason returns true if TradSesStatusRejReason is present, Tag 567
func (m NoTradingSessions) HasTradSesStatusRejReason() bool {
	return m.Has(tag.TradSesStatusRejReason)
}

//HasTradSesStartTime returns true if TradSesStartTime is present, Tag 341
func (m NoTradingSessions) HasTradSesStartTime() bool {
	return m.Has(tag.TradSesStartTime)
}

//HasTradSesOpenTime returns true if TradSesOpenTime is present, Tag 342
func (m NoTradingSessions) HasTradSesOpenTime() bool {
	return m.Has(tag.TradSesOpenTime)
}

//HasTradSesPreCloseTime returns true if TradSesPreCloseTime is present, Tag 343
func (m NoTradingSessions) HasTradSesPreCloseTime() bool {
	return m.Has(tag.TradSesPreCloseTime)
}

//HasTradSesCloseTime returns true if TradSesCloseTime is present, Tag 344
func (m NoTradingSessions) HasTradSesCloseTime() bool {
	return m.Has(tag.TradSesCloseTime)
}

//HasTradSesEndTime returns true if TradSesEndTime is present, Tag 345
func (m NoTradingSessions) HasTradSesEndTime() bool {
	return m.Has(tag.TradSesEndTime)
}

//HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387
func (m NoTradingSessions) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
}

//HasText returns true if Text is present, Tag 58
func (m NoTradingSessions) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoTradingSessions) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoTradingSessions) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.TradSesMethod), quickfix.GroupElement(tag.TradSesMode), quickfix.GroupElement(tag.UnsolicitedIndicator), quickfix.GroupElement(tag.TradSesStatus), quickfix.GroupElement(tag.TradSesStatusRejReason), quickfix.GroupElement(tag.TradSesStartTime), quickfix.GroupElement(tag.TradSesOpenTime), quickfix.GroupElement(tag.TradSesPreCloseTime), quickfix.GroupElement(tag.TradSesCloseTime), quickfix.GroupElement(tag.TradSesEndTime), quickfix.GroupElement(tag.TotalVolumeTraded), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
}

//Add create and append a new NoTradingSessions to this group
func (m NoTradingSessionsRepeatingGroup) Add() NoTradingSessions {
	g := m.RepeatingGroup.Add()
	return NoTradingSessions{g}
}

//Get returns the ith NoTradingSessions in the NoTradingSessionsRepeatinGroup
func (m NoTradingSessionsRepeatingGroup) Get(i int) NoTradingSessions {
	return NoTradingSessions{m.RepeatingGroup.Get(i)}
}
