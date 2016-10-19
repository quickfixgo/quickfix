package tradingsessionlist

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//TradingSessionList is the fix50sp1 TradingSessionList type, MsgType = BJ
type TradingSessionList struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a TradingSessionList from a quickfix.Message instance
func FromMessage(m *quickfix.Message) TradingSessionList {
	return TradingSessionList{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m TradingSessionList) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a TradingSessionList initialized with the required fields for TradingSessionList
func New() (m TradingSessionList) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BJ"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg TradingSessionList, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "8", "BJ", r
}

//SetTradSesReqID sets TradSesReqID, Tag 335
func (m TradingSessionList) SetTradSesReqID(v string) {
	m.Set(field.NewTradSesReqID(v))
}

//SetNoTradingSessions sets NoTradingSessions, Tag 386
func (m TradingSessionList) SetNoTradingSessions(f NoTradingSessionsRepeatingGroup) {
	m.SetGroup(f)
}

//SetApplID sets ApplID, Tag 1180
func (m TradingSessionList) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m TradingSessionList) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m TradingSessionList) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m TradingSessionList) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//GetTradSesReqID gets TradSesReqID, Tag 335
func (m TradingSessionList) GetTradSesReqID() (v string, err quickfix.MessageRejectError) {
	var f field.TradSesReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTradingSessions gets NoTradingSessions, Tag 386
func (m TradingSessionList) GetNoTradingSessions() (NoTradingSessionsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTradingSessionsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetApplID gets ApplID, Tag 1180
func (m TradingSessionList) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m TradingSessionList) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m TradingSessionList) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m TradingSessionList) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasTradSesReqID returns true if TradSesReqID is present, Tag 335
func (m TradingSessionList) HasTradSesReqID() bool {
	return m.Has(tag.TradSesReqID)
}

//HasNoTradingSessions returns true if NoTradingSessions is present, Tag 386
func (m TradingSessionList) HasNoTradingSessions() bool {
	return m.Has(tag.NoTradingSessions)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m TradingSessionList) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m TradingSessionList) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m TradingSessionList) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m TradingSessionList) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//NoTradingSessions is a repeating group element, Tag 386
type NoTradingSessions struct {
	*quickfix.Group
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoTradingSessions) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoTradingSessions) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoTradingSessions) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetTradSesMethod sets TradSesMethod, Tag 338
func (m NoTradingSessions) SetTradSesMethod(v enum.TradSesMethod) {
	m.Set(field.NewTradSesMethod(v))
}

//SetTradSesMode sets TradSesMode, Tag 339
func (m NoTradingSessions) SetTradSesMode(v enum.TradSesMode) {
	m.Set(field.NewTradSesMode(v))
}

//SetUnsolicitedIndicator sets UnsolicitedIndicator, Tag 325
func (m NoTradingSessions) SetUnsolicitedIndicator(v bool) {
	m.Set(field.NewUnsolicitedIndicator(v))
}

//SetTradSesStatus sets TradSesStatus, Tag 340
func (m NoTradingSessions) SetTradSesStatus(v enum.TradSesStatus) {
	m.Set(field.NewTradSesStatus(v))
}

//SetTradSesStatusRejReason sets TradSesStatusRejReason, Tag 567
func (m NoTradingSessions) SetTradSesStatusRejReason(v enum.TradSesStatusRejReason) {
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
func (m NoTradingSessions) SetTotalVolumeTraded(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalVolumeTraded(value, scale))
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

//SetMarketID sets MarketID, Tag 1301
func (m NoTradingSessions) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m NoTradingSessions) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetTradingSessionDesc sets TradingSessionDesc, Tag 1326
func (m NoTradingSessions) SetTradingSessionDesc(v string) {
	m.Set(field.NewTradingSessionDesc(v))
}

//SetNoOrdTypeRules sets NoOrdTypeRules, Tag 1237
func (m NoTradingSessions) SetNoOrdTypeRules(f NoOrdTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoTimeInForceRules sets NoTimeInForceRules, Tag 1239
func (m NoTradingSessions) SetNoTimeInForceRules(f NoTimeInForceRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoExecInstRules sets NoExecInstRules, Tag 1232
func (m NoTradingSessions) SetNoExecInstRules(f NoExecInstRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMatchRules sets NoMatchRules, Tag 1235
func (m NoTradingSessions) SetNoMatchRules(f NoMatchRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoMDFeedTypes sets NoMDFeedTypes, Tag 1141
func (m NoTradingSessions) SetNoMDFeedTypes(f NoMDFeedTypesRepeatingGroup) {
	m.SetGroup(f)
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoTradingSessions) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoTradingSessions) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoTradingSessions) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMethod gets TradSesMethod, Tag 338
func (m NoTradingSessions) GetTradSesMethod() (v enum.TradSesMethod, err quickfix.MessageRejectError) {
	var f field.TradSesMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesMode gets TradSesMode, Tag 339
func (m NoTradingSessions) GetTradSesMode() (v enum.TradSesMode, err quickfix.MessageRejectError) {
	var f field.TradSesModeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetUnsolicitedIndicator gets UnsolicitedIndicator, Tag 325
func (m NoTradingSessions) GetUnsolicitedIndicator() (v bool, err quickfix.MessageRejectError) {
	var f field.UnsolicitedIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesStatus gets TradSesStatus, Tag 340
func (m NoTradingSessions) GetTradSesStatus() (v enum.TradSesStatus, err quickfix.MessageRejectError) {
	var f field.TradSesStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesStatusRejReason gets TradSesStatusRejReason, Tag 567
func (m NoTradingSessions) GetTradSesStatusRejReason() (v enum.TradSesStatusRejReason, err quickfix.MessageRejectError) {
	var f field.TradSesStatusRejReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesStartTime gets TradSesStartTime, Tag 341
func (m NoTradingSessions) GetTradSesStartTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesStartTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesOpenTime gets TradSesOpenTime, Tag 342
func (m NoTradingSessions) GetTradSesOpenTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesOpenTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesPreCloseTime gets TradSesPreCloseTime, Tag 343
func (m NoTradingSessions) GetTradSesPreCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesPreCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesCloseTime gets TradSesCloseTime, Tag 344
func (m NoTradingSessions) GetTradSesCloseTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesCloseTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradSesEndTime gets TradSesEndTime, Tag 345
func (m NoTradingSessions) GetTradSesEndTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TradSesEndTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m NoTradingSessions) GetTotalVolumeTraded() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m NoTradingSessions) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoTradingSessions) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoTradingSessions) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m NoTradingSessions) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m NoTradingSessions) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionDesc gets TradingSessionDesc, Tag 1326
func (m NoTradingSessions) GetTradingSessionDesc() (v string, err quickfix.MessageRejectError) {
	var f field.TradingSessionDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoOrdTypeRules gets NoOrdTypeRules, Tag 1237
func (m NoTradingSessions) GetNoOrdTypeRules() (NoOrdTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoTimeInForceRules gets NoTimeInForceRules, Tag 1239
func (m NoTradingSessions) GetNoTimeInForceRules() (NoTimeInForceRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTimeInForceRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoExecInstRules gets NoExecInstRules, Tag 1232
func (m NoTradingSessions) GetNoExecInstRules() (NoExecInstRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecInstRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMatchRules gets NoMatchRules, Tag 1235
func (m NoTradingSessions) GetNoMatchRules() (NoMatchRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMatchRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoMDFeedTypes gets NoMDFeedTypes, Tag 1141
func (m NoTradingSessions) GetNoMDFeedTypes() (NoMDFeedTypesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDFeedTypesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
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

//HasMarketID returns true if MarketID is present, Tag 1301
func (m NoTradingSessions) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m NoTradingSessions) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasTradingSessionDesc returns true if TradingSessionDesc is present, Tag 1326
func (m NoTradingSessions) HasTradingSessionDesc() bool {
	return m.Has(tag.TradingSessionDesc)
}

//HasNoOrdTypeRules returns true if NoOrdTypeRules is present, Tag 1237
func (m NoTradingSessions) HasNoOrdTypeRules() bool {
	return m.Has(tag.NoOrdTypeRules)
}

//HasNoTimeInForceRules returns true if NoTimeInForceRules is present, Tag 1239
func (m NoTradingSessions) HasNoTimeInForceRules() bool {
	return m.Has(tag.NoTimeInForceRules)
}

//HasNoExecInstRules returns true if NoExecInstRules is present, Tag 1232
func (m NoTradingSessions) HasNoExecInstRules() bool {
	return m.Has(tag.NoExecInstRules)
}

//HasNoMatchRules returns true if NoMatchRules is present, Tag 1235
func (m NoTradingSessions) HasNoMatchRules() bool {
	return m.Has(tag.NoMatchRules)
}

//HasNoMDFeedTypes returns true if NoMDFeedTypes is present, Tag 1141
func (m NoTradingSessions) HasNoMDFeedTypes() bool {
	return m.Has(tag.NoMDFeedTypes)
}

//NoOrdTypeRules is a repeating group element, Tag 1237
type NoOrdTypeRules struct {
	*quickfix.Group
}

//SetOrdType sets OrdType, Tag 40
func (m NoOrdTypeRules) SetOrdType(v enum.OrdType) {
	m.Set(field.NewOrdType(v))
}

//GetOrdType gets OrdType, Tag 40
func (m NoOrdTypeRules) GetOrdType() (v enum.OrdType, err quickfix.MessageRejectError) {
	var f field.OrdTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasOrdType returns true if OrdType is present, Tag 40
func (m NoOrdTypeRules) HasOrdType() bool {
	return m.Has(tag.OrdType)
}

//NoOrdTypeRulesRepeatingGroup is a repeating group, Tag 1237
type NoOrdTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoOrdTypeRulesRepeatingGroup returns an initialized, NoOrdTypeRulesRepeatingGroup
func NewNoOrdTypeRulesRepeatingGroup() NoOrdTypeRulesRepeatingGroup {
	return NoOrdTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoOrdTypeRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.OrdType)})}
}

//Add create and append a new NoOrdTypeRules to this group
func (m NoOrdTypeRulesRepeatingGroup) Add() NoOrdTypeRules {
	g := m.RepeatingGroup.Add()
	return NoOrdTypeRules{g}
}

//Get returns the ith NoOrdTypeRules in the NoOrdTypeRulesRepeatinGroup
func (m NoOrdTypeRulesRepeatingGroup) Get(i int) NoOrdTypeRules {
	return NoOrdTypeRules{m.RepeatingGroup.Get(i)}
}

//NoTimeInForceRules is a repeating group element, Tag 1239
type NoTimeInForceRules struct {
	*quickfix.Group
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NoTimeInForceRules) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NoTimeInForceRules) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NoTimeInForceRules) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//NoTimeInForceRulesRepeatingGroup is a repeating group, Tag 1239
type NoTimeInForceRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTimeInForceRulesRepeatingGroup returns an initialized, NoTimeInForceRulesRepeatingGroup
func NewNoTimeInForceRulesRepeatingGroup() NoTimeInForceRulesRepeatingGroup {
	return NoTimeInForceRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTimeInForceRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TimeInForce)})}
}

//Add create and append a new NoTimeInForceRules to this group
func (m NoTimeInForceRulesRepeatingGroup) Add() NoTimeInForceRules {
	g := m.RepeatingGroup.Add()
	return NoTimeInForceRules{g}
}

//Get returns the ith NoTimeInForceRules in the NoTimeInForceRulesRepeatinGroup
func (m NoTimeInForceRulesRepeatingGroup) Get(i int) NoTimeInForceRules {
	return NoTimeInForceRules{m.RepeatingGroup.Get(i)}
}

//NoExecInstRules is a repeating group element, Tag 1232
type NoExecInstRules struct {
	*quickfix.Group
}

//SetExecInstValue sets ExecInstValue, Tag 1308
func (m NoExecInstRules) SetExecInstValue(v string) {
	m.Set(field.NewExecInstValue(v))
}

//GetExecInstValue gets ExecInstValue, Tag 1308
func (m NoExecInstRules) GetExecInstValue() (v string, err quickfix.MessageRejectError) {
	var f field.ExecInstValueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasExecInstValue returns true if ExecInstValue is present, Tag 1308
func (m NoExecInstRules) HasExecInstValue() bool {
	return m.Has(tag.ExecInstValue)
}

//NoExecInstRulesRepeatingGroup is a repeating group, Tag 1232
type NoExecInstRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoExecInstRulesRepeatingGroup returns an initialized, NoExecInstRulesRepeatingGroup
func NewNoExecInstRulesRepeatingGroup() NoExecInstRulesRepeatingGroup {
	return NoExecInstRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoExecInstRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.ExecInstValue)})}
}

//Add create and append a new NoExecInstRules to this group
func (m NoExecInstRulesRepeatingGroup) Add() NoExecInstRules {
	g := m.RepeatingGroup.Add()
	return NoExecInstRules{g}
}

//Get returns the ith NoExecInstRules in the NoExecInstRulesRepeatinGroup
func (m NoExecInstRulesRepeatingGroup) Get(i int) NoExecInstRules {
	return NoExecInstRules{m.RepeatingGroup.Get(i)}
}

//NoMatchRules is a repeating group element, Tag 1235
type NoMatchRules struct {
	*quickfix.Group
}

//SetMatchAlgorithm sets MatchAlgorithm, Tag 1142
func (m NoMatchRules) SetMatchAlgorithm(v string) {
	m.Set(field.NewMatchAlgorithm(v))
}

//SetMatchType sets MatchType, Tag 574
func (m NoMatchRules) SetMatchType(v enum.MatchType) {
	m.Set(field.NewMatchType(v))
}

//GetMatchAlgorithm gets MatchAlgorithm, Tag 1142
func (m NoMatchRules) GetMatchAlgorithm() (v string, err quickfix.MessageRejectError) {
	var f field.MatchAlgorithmField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMatchType gets MatchType, Tag 574
func (m NoMatchRules) GetMatchType() (v enum.MatchType, err quickfix.MessageRejectError) {
	var f field.MatchTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMatchAlgorithm returns true if MatchAlgorithm is present, Tag 1142
func (m NoMatchRules) HasMatchAlgorithm() bool {
	return m.Has(tag.MatchAlgorithm)
}

//HasMatchType returns true if MatchType is present, Tag 574
func (m NoMatchRules) HasMatchType() bool {
	return m.Has(tag.MatchType)
}

//NoMatchRulesRepeatingGroup is a repeating group, Tag 1235
type NoMatchRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMatchRulesRepeatingGroup returns an initialized, NoMatchRulesRepeatingGroup
func NewNoMatchRulesRepeatingGroup() NoMatchRulesRepeatingGroup {
	return NoMatchRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMatchRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MatchAlgorithm), quickfix.GroupElement(tag.MatchType)})}
}

//Add create and append a new NoMatchRules to this group
func (m NoMatchRulesRepeatingGroup) Add() NoMatchRules {
	g := m.RepeatingGroup.Add()
	return NoMatchRules{g}
}

//Get returns the ith NoMatchRules in the NoMatchRulesRepeatinGroup
func (m NoMatchRulesRepeatingGroup) Get(i int) NoMatchRules {
	return NoMatchRules{m.RepeatingGroup.Get(i)}
}

//NoMDFeedTypes is a repeating group element, Tag 1141
type NoMDFeedTypes struct {
	*quickfix.Group
}

//SetMDFeedType sets MDFeedType, Tag 1022
func (m NoMDFeedTypes) SetMDFeedType(v string) {
	m.Set(field.NewMDFeedType(v))
}

//SetMarketDepth sets MarketDepth, Tag 264
func (m NoMDFeedTypes) SetMarketDepth(v int) {
	m.Set(field.NewMarketDepth(v))
}

//SetMDBookType sets MDBookType, Tag 1021
func (m NoMDFeedTypes) SetMDBookType(v enum.MDBookType) {
	m.Set(field.NewMDBookType(v))
}

//GetMDFeedType gets MDFeedType, Tag 1022
func (m NoMDFeedTypes) GetMDFeedType() (v string, err quickfix.MessageRejectError) {
	var f field.MDFeedTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketDepth gets MarketDepth, Tag 264
func (m NoMDFeedTypes) GetMarketDepth() (v int, err quickfix.MessageRejectError) {
	var f field.MarketDepthField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDBookType gets MDBookType, Tag 1021
func (m NoMDFeedTypes) GetMDBookType() (v enum.MDBookType, err quickfix.MessageRejectError) {
	var f field.MDBookTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMDFeedType returns true if MDFeedType is present, Tag 1022
func (m NoMDFeedTypes) HasMDFeedType() bool {
	return m.Has(tag.MDFeedType)
}

//HasMarketDepth returns true if MarketDepth is present, Tag 264
func (m NoMDFeedTypes) HasMarketDepth() bool {
	return m.Has(tag.MarketDepth)
}

//HasMDBookType returns true if MDBookType is present, Tag 1021
func (m NoMDFeedTypes) HasMDBookType() bool {
	return m.Has(tag.MDBookType)
}

//NoMDFeedTypesRepeatingGroup is a repeating group, Tag 1141
type NoMDFeedTypesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDFeedTypesRepeatingGroup returns an initialized, NoMDFeedTypesRepeatingGroup
func NewNoMDFeedTypesRepeatingGroup() NoMDFeedTypesRepeatingGroup {
	return NoMDFeedTypesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDFeedTypes,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDFeedType), quickfix.GroupElement(tag.MarketDepth), quickfix.GroupElement(tag.MDBookType)})}
}

//Add create and append a new NoMDFeedTypes to this group
func (m NoMDFeedTypesRepeatingGroup) Add() NoMDFeedTypes {
	g := m.RepeatingGroup.Add()
	return NoMDFeedTypes{g}
}

//Get returns the ith NoMDFeedTypes in the NoMDFeedTypesRepeatinGroup
func (m NoMDFeedTypesRepeatingGroup) Get(i int) NoMDFeedTypes {
	return NoMDFeedTypes{m.RepeatingGroup.Get(i)}
}

//NoTradingSessionsRepeatingGroup is a repeating group, Tag 386
type NoTradingSessionsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTradingSessionsRepeatingGroup returns an initialized, NoTradingSessionsRepeatingGroup
func NewNoTradingSessionsRepeatingGroup() NoTradingSessionsRepeatingGroup {
	return NoTradingSessionsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTradingSessions,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.TradSesMethod), quickfix.GroupElement(tag.TradSesMode), quickfix.GroupElement(tag.UnsolicitedIndicator), quickfix.GroupElement(tag.TradSesStatus), quickfix.GroupElement(tag.TradSesStatusRejReason), quickfix.GroupElement(tag.TradSesStartTime), quickfix.GroupElement(tag.TradSesOpenTime), quickfix.GroupElement(tag.TradSesPreCloseTime), quickfix.GroupElement(tag.TradSesCloseTime), quickfix.GroupElement(tag.TradSesEndTime), quickfix.GroupElement(tag.TotalVolumeTraded), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText), quickfix.GroupElement(tag.MarketID), quickfix.GroupElement(tag.MarketSegmentID), quickfix.GroupElement(tag.TradingSessionDesc), NewNoOrdTypeRulesRepeatingGroup(), NewNoTimeInForceRulesRepeatingGroup(), NewNoExecInstRulesRepeatingGroup(), NewNoMatchRulesRepeatingGroup(), NewNoMDFeedTypesRepeatingGroup()})}
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
