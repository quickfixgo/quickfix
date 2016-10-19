package marketdefinition

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fixt11"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDefinition is the fix50sp2 MarketDefinition type, MsgType = BU
type MarketDefinition struct {
	fixt11.Header
	*quickfix.Body
	fixt11.Trailer
	Message *quickfix.Message
}

//FromMessage creates a MarketDefinition from a quickfix.Message instance
func FromMessage(m *quickfix.Message) MarketDefinition {
	return MarketDefinition{
		Header:  fixt11.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fixt11.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDefinition) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MarketDefinition initialized with the required fields for MarketDefinition
func New(marketreportid field.MarketReportIDField, marketid field.MarketIDField) (m MarketDefinition) {
	m.Message = quickfix.NewMessage()
	m.Header = fixt11.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("BU"))
	m.Set(marketreportid)
	m.Set(marketid)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDefinition, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "9", "BU", r
}

//SetCurrency sets Currency, Tag 15
func (m MarketDefinition) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetText sets Text, Tag 58
func (m MarketDefinition) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m MarketDefinition) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m MarketDefinition) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m MarketDefinition) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPriceType sets PriceType, Tag 423
func (m MarketDefinition) SetPriceType(v enum.PriceType) {
	m.Set(field.NewPriceType(v))
}

//SetRoundLot sets RoundLot, Tag 561
func (m MarketDefinition) SetRoundLot(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundLot(value, scale))
}

//SetMinTradeVol sets MinTradeVol, Tag 562
func (m MarketDefinition) SetMinTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinTradeVol(value, scale))
}

//SetExpirationCycle sets ExpirationCycle, Tag 827
func (m MarketDefinition) SetExpirationCycle(v enum.ExpirationCycle) {
	m.Set(field.NewExpirationCycle(v))
}

//SetMaxTradeVol sets MaxTradeVol, Tag 1140
func (m MarketDefinition) SetMaxTradeVol(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxTradeVol(value, scale))
}

//SetMaxPriceVariation sets MaxPriceVariation, Tag 1143
func (m MarketDefinition) SetMaxPriceVariation(value decimal.Decimal, scale int32) {
	m.Set(field.NewMaxPriceVariation(value, scale))
}

//SetImpliedMarketIndicator sets ImpliedMarketIndicator, Tag 1144
func (m MarketDefinition) SetImpliedMarketIndicator(v enum.ImpliedMarketIndicator) {
	m.Set(field.NewImpliedMarketIndicator(v))
}

//SetLowLimitPrice sets LowLimitPrice, Tag 1148
func (m MarketDefinition) SetLowLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewLowLimitPrice(value, scale))
}

//SetHighLimitPrice sets HighLimitPrice, Tag 1149
func (m MarketDefinition) SetHighLimitPrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewHighLimitPrice(value, scale))
}

//SetTradingReferencePrice sets TradingReferencePrice, Tag 1150
func (m MarketDefinition) SetTradingReferencePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewTradingReferencePrice(value, scale))
}

//SetApplID sets ApplID, Tag 1180
func (m MarketDefinition) SetApplID(v string) {
	m.Set(field.NewApplID(v))
}

//SetApplSeqNum sets ApplSeqNum, Tag 1181
func (m MarketDefinition) SetApplSeqNum(v int) {
	m.Set(field.NewApplSeqNum(v))
}

//SetNoTickRules sets NoTickRules, Tag 1205
func (m MarketDefinition) SetNoTickRules(f NoTickRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoExecInstRules sets NoExecInstRules, Tag 1232
func (m MarketDefinition) SetNoExecInstRules(f NoExecInstRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoLotTypeRules sets NoLotTypeRules, Tag 1234
func (m MarketDefinition) SetNoLotTypeRules(f NoLotTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoOrdTypeRules sets NoOrdTypeRules, Tag 1237
func (m MarketDefinition) SetNoOrdTypeRules(f NoOrdTypeRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoTimeInForceRules sets NoTimeInForceRules, Tag 1239
func (m MarketDefinition) SetNoTimeInForceRules(f NoTimeInForceRulesRepeatingGroup) {
	m.SetGroup(f)
}

//SetTradingCurrency sets TradingCurrency, Tag 1245
func (m MarketDefinition) SetTradingCurrency(v string) {
	m.Set(field.NewTradingCurrency(v))
}

//SetMarketSegmentID sets MarketSegmentID, Tag 1300
func (m MarketDefinition) SetMarketSegmentID(v string) {
	m.Set(field.NewMarketSegmentID(v))
}

//SetMarketID sets MarketID, Tag 1301
func (m MarketDefinition) SetMarketID(v string) {
	m.Set(field.NewMarketID(v))
}

//SetPriceLimitType sets PriceLimitType, Tag 1306
func (m MarketDefinition) SetPriceLimitType(v enum.PriceLimitType) {
	m.Set(field.NewPriceLimitType(v))
}

//SetParentMktSegmID sets ParentMktSegmID, Tag 1325
func (m MarketDefinition) SetParentMktSegmID(v string) {
	m.Set(field.NewParentMktSegmID(v))
}

//SetApplLastSeqNum sets ApplLastSeqNum, Tag 1350
func (m MarketDefinition) SetApplLastSeqNum(v int) {
	m.Set(field.NewApplLastSeqNum(v))
}

//SetApplResendFlag sets ApplResendFlag, Tag 1352
func (m MarketDefinition) SetApplResendFlag(v bool) {
	m.Set(field.NewApplResendFlag(v))
}

//SetMultilegModel sets MultilegModel, Tag 1377
func (m MarketDefinition) SetMultilegModel(v enum.MultilegModel) {
	m.Set(field.NewMultilegModel(v))
}

//SetMultilegPriceMethod sets MultilegPriceMethod, Tag 1378
func (m MarketDefinition) SetMultilegPriceMethod(v enum.MultilegPriceMethod) {
	m.Set(field.NewMultilegPriceMethod(v))
}

//SetMarketReqID sets MarketReqID, Tag 1393
func (m MarketDefinition) SetMarketReqID(v string) {
	m.Set(field.NewMarketReqID(v))
}

//SetMarketReportID sets MarketReportID, Tag 1394
func (m MarketDefinition) SetMarketReportID(v string) {
	m.Set(field.NewMarketReportID(v))
}

//SetMarketSegmentDesc sets MarketSegmentDesc, Tag 1396
func (m MarketDefinition) SetMarketSegmentDesc(v string) {
	m.Set(field.NewMarketSegmentDesc(v))
}

//SetEncodedMktSegmDescLen sets EncodedMktSegmDescLen, Tag 1397
func (m MarketDefinition) SetEncodedMktSegmDescLen(v int) {
	m.Set(field.NewEncodedMktSegmDescLen(v))
}

//SetEncodedMktSegmDesc sets EncodedMktSegmDesc, Tag 1398
func (m MarketDefinition) SetEncodedMktSegmDesc(v string) {
	m.Set(field.NewEncodedMktSegmDesc(v))
}

//GetCurrency gets Currency, Tag 15
func (m MarketDefinition) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m MarketDefinition) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m MarketDefinition) GetTransactTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.TransactTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m MarketDefinition) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m MarketDefinition) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceType gets PriceType, Tag 423
func (m MarketDefinition) GetPriceType() (v enum.PriceType, err quickfix.MessageRejectError) {
	var f field.PriceTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRoundLot gets RoundLot, Tag 561
func (m MarketDefinition) GetRoundLot() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RoundLotField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinTradeVol gets MinTradeVol, Tag 562
func (m MarketDefinition) GetMinTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpirationCycle gets ExpirationCycle, Tag 827
func (m MarketDefinition) GetExpirationCycle() (v enum.ExpirationCycle, err quickfix.MessageRejectError) {
	var f field.ExpirationCycleField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxTradeVol gets MaxTradeVol, Tag 1140
func (m MarketDefinition) GetMaxTradeVol() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxTradeVolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaxPriceVariation gets MaxPriceVariation, Tag 1143
func (m MarketDefinition) GetMaxPriceVariation() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MaxPriceVariationField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetImpliedMarketIndicator gets ImpliedMarketIndicator, Tag 1144
func (m MarketDefinition) GetImpliedMarketIndicator() (v enum.ImpliedMarketIndicator, err quickfix.MessageRejectError) {
	var f field.ImpliedMarketIndicatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLowLimitPrice gets LowLimitPrice, Tag 1148
func (m MarketDefinition) GetLowLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.LowLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetHighLimitPrice gets HighLimitPrice, Tag 1149
func (m MarketDefinition) GetHighLimitPrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.HighLimitPriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingReferencePrice gets TradingReferencePrice, Tag 1150
func (m MarketDefinition) GetTradingReferencePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TradingReferencePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplID gets ApplID, Tag 1180
func (m MarketDefinition) GetApplID() (v string, err quickfix.MessageRejectError) {
	var f field.ApplIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplSeqNum gets ApplSeqNum, Tag 1181
func (m MarketDefinition) GetApplSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoTickRules gets NoTickRules, Tag 1205
func (m MarketDefinition) GetNoTickRules() (NoTickRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTickRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoExecInstRules gets NoExecInstRules, Tag 1232
func (m MarketDefinition) GetNoExecInstRules() (NoExecInstRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoExecInstRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoLotTypeRules gets NoLotTypeRules, Tag 1234
func (m MarketDefinition) GetNoLotTypeRules() (NoLotTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoLotTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoOrdTypeRules gets NoOrdTypeRules, Tag 1237
func (m MarketDefinition) GetNoOrdTypeRules() (NoOrdTypeRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoOrdTypeRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoTimeInForceRules gets NoTimeInForceRules, Tag 1239
func (m MarketDefinition) GetNoTimeInForceRules() (NoTimeInForceRulesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoTimeInForceRulesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetTradingCurrency gets TradingCurrency, Tag 1245
func (m MarketDefinition) GetTradingCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.TradingCurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketSegmentID gets MarketSegmentID, Tag 1300
func (m MarketDefinition) GetMarketSegmentID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketID gets MarketID, Tag 1301
func (m MarketDefinition) GetMarketID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetPriceLimitType gets PriceLimitType, Tag 1306
func (m MarketDefinition) GetPriceLimitType() (v enum.PriceLimitType, err quickfix.MessageRejectError) {
	var f field.PriceLimitTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetParentMktSegmID gets ParentMktSegmID, Tag 1325
func (m MarketDefinition) GetParentMktSegmID() (v string, err quickfix.MessageRejectError) {
	var f field.ParentMktSegmIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplLastSeqNum gets ApplLastSeqNum, Tag 1350
func (m MarketDefinition) GetApplLastSeqNum() (v int, err quickfix.MessageRejectError) {
	var f field.ApplLastSeqNumField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetApplResendFlag gets ApplResendFlag, Tag 1352
func (m MarketDefinition) GetApplResendFlag() (v bool, err quickfix.MessageRejectError) {
	var f field.ApplResendFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultilegModel gets MultilegModel, Tag 1377
func (m MarketDefinition) GetMultilegModel() (v enum.MultilegModel, err quickfix.MessageRejectError) {
	var f field.MultilegModelField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMultilegPriceMethod gets MultilegPriceMethod, Tag 1378
func (m MarketDefinition) GetMultilegPriceMethod() (v enum.MultilegPriceMethod, err quickfix.MessageRejectError) {
	var f field.MultilegPriceMethodField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketReqID gets MarketReqID, Tag 1393
func (m MarketDefinition) GetMarketReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketReportID gets MarketReportID, Tag 1394
func (m MarketDefinition) GetMarketReportID() (v string, err quickfix.MessageRejectError) {
	var f field.MarketReportIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMarketSegmentDesc gets MarketSegmentDesc, Tag 1396
func (m MarketDefinition) GetMarketSegmentDesc() (v string, err quickfix.MessageRejectError) {
	var f field.MarketSegmentDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedMktSegmDescLen gets EncodedMktSegmDescLen, Tag 1397
func (m MarketDefinition) GetEncodedMktSegmDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedMktSegmDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedMktSegmDesc gets EncodedMktSegmDesc, Tag 1398
func (m MarketDefinition) GetEncodedMktSegmDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedMktSegmDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m MarketDefinition) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasText returns true if Text is present, Tag 58
func (m MarketDefinition) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m MarketDefinition) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m MarketDefinition) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m MarketDefinition) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m MarketDefinition) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasRoundLot returns true if RoundLot is present, Tag 561
func (m MarketDefinition) HasRoundLot() bool {
	return m.Has(tag.RoundLot)
}

//HasMinTradeVol returns true if MinTradeVol is present, Tag 562
func (m MarketDefinition) HasMinTradeVol() bool {
	return m.Has(tag.MinTradeVol)
}

//HasExpirationCycle returns true if ExpirationCycle is present, Tag 827
func (m MarketDefinition) HasExpirationCycle() bool {
	return m.Has(tag.ExpirationCycle)
}

//HasMaxTradeVol returns true if MaxTradeVol is present, Tag 1140
func (m MarketDefinition) HasMaxTradeVol() bool {
	return m.Has(tag.MaxTradeVol)
}

//HasMaxPriceVariation returns true if MaxPriceVariation is present, Tag 1143
func (m MarketDefinition) HasMaxPriceVariation() bool {
	return m.Has(tag.MaxPriceVariation)
}

//HasImpliedMarketIndicator returns true if ImpliedMarketIndicator is present, Tag 1144
func (m MarketDefinition) HasImpliedMarketIndicator() bool {
	return m.Has(tag.ImpliedMarketIndicator)
}

//HasLowLimitPrice returns true if LowLimitPrice is present, Tag 1148
func (m MarketDefinition) HasLowLimitPrice() bool {
	return m.Has(tag.LowLimitPrice)
}

//HasHighLimitPrice returns true if HighLimitPrice is present, Tag 1149
func (m MarketDefinition) HasHighLimitPrice() bool {
	return m.Has(tag.HighLimitPrice)
}

//HasTradingReferencePrice returns true if TradingReferencePrice is present, Tag 1150
func (m MarketDefinition) HasTradingReferencePrice() bool {
	return m.Has(tag.TradingReferencePrice)
}

//HasApplID returns true if ApplID is present, Tag 1180
func (m MarketDefinition) HasApplID() bool {
	return m.Has(tag.ApplID)
}

//HasApplSeqNum returns true if ApplSeqNum is present, Tag 1181
func (m MarketDefinition) HasApplSeqNum() bool {
	return m.Has(tag.ApplSeqNum)
}

//HasNoTickRules returns true if NoTickRules is present, Tag 1205
func (m MarketDefinition) HasNoTickRules() bool {
	return m.Has(tag.NoTickRules)
}

//HasNoExecInstRules returns true if NoExecInstRules is present, Tag 1232
func (m MarketDefinition) HasNoExecInstRules() bool {
	return m.Has(tag.NoExecInstRules)
}

//HasNoLotTypeRules returns true if NoLotTypeRules is present, Tag 1234
func (m MarketDefinition) HasNoLotTypeRules() bool {
	return m.Has(tag.NoLotTypeRules)
}

//HasNoOrdTypeRules returns true if NoOrdTypeRules is present, Tag 1237
func (m MarketDefinition) HasNoOrdTypeRules() bool {
	return m.Has(tag.NoOrdTypeRules)
}

//HasNoTimeInForceRules returns true if NoTimeInForceRules is present, Tag 1239
func (m MarketDefinition) HasNoTimeInForceRules() bool {
	return m.Has(tag.NoTimeInForceRules)
}

//HasTradingCurrency returns true if TradingCurrency is present, Tag 1245
func (m MarketDefinition) HasTradingCurrency() bool {
	return m.Has(tag.TradingCurrency)
}

//HasMarketSegmentID returns true if MarketSegmentID is present, Tag 1300
func (m MarketDefinition) HasMarketSegmentID() bool {
	return m.Has(tag.MarketSegmentID)
}

//HasMarketID returns true if MarketID is present, Tag 1301
func (m MarketDefinition) HasMarketID() bool {
	return m.Has(tag.MarketID)
}

//HasPriceLimitType returns true if PriceLimitType is present, Tag 1306
func (m MarketDefinition) HasPriceLimitType() bool {
	return m.Has(tag.PriceLimitType)
}

//HasParentMktSegmID returns true if ParentMktSegmID is present, Tag 1325
func (m MarketDefinition) HasParentMktSegmID() bool {
	return m.Has(tag.ParentMktSegmID)
}

//HasApplLastSeqNum returns true if ApplLastSeqNum is present, Tag 1350
func (m MarketDefinition) HasApplLastSeqNum() bool {
	return m.Has(tag.ApplLastSeqNum)
}

//HasApplResendFlag returns true if ApplResendFlag is present, Tag 1352
func (m MarketDefinition) HasApplResendFlag() bool {
	return m.Has(tag.ApplResendFlag)
}

//HasMultilegModel returns true if MultilegModel is present, Tag 1377
func (m MarketDefinition) HasMultilegModel() bool {
	return m.Has(tag.MultilegModel)
}

//HasMultilegPriceMethod returns true if MultilegPriceMethod is present, Tag 1378
func (m MarketDefinition) HasMultilegPriceMethod() bool {
	return m.Has(tag.MultilegPriceMethod)
}

//HasMarketReqID returns true if MarketReqID is present, Tag 1393
func (m MarketDefinition) HasMarketReqID() bool {
	return m.Has(tag.MarketReqID)
}

//HasMarketReportID returns true if MarketReportID is present, Tag 1394
func (m MarketDefinition) HasMarketReportID() bool {
	return m.Has(tag.MarketReportID)
}

//HasMarketSegmentDesc returns true if MarketSegmentDesc is present, Tag 1396
func (m MarketDefinition) HasMarketSegmentDesc() bool {
	return m.Has(tag.MarketSegmentDesc)
}

//HasEncodedMktSegmDescLen returns true if EncodedMktSegmDescLen is present, Tag 1397
func (m MarketDefinition) HasEncodedMktSegmDescLen() bool {
	return m.Has(tag.EncodedMktSegmDescLen)
}

//HasEncodedMktSegmDesc returns true if EncodedMktSegmDesc is present, Tag 1398
func (m MarketDefinition) HasEncodedMktSegmDesc() bool {
	return m.Has(tag.EncodedMktSegmDesc)
}

//NoTickRules is a repeating group element, Tag 1205
type NoTickRules struct {
	*quickfix.Group
}

//SetStartTickPriceRange sets StartTickPriceRange, Tag 1206
func (m NoTickRules) SetStartTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewStartTickPriceRange(value, scale))
}

//SetEndTickPriceRange sets EndTickPriceRange, Tag 1207
func (m NoTickRules) SetEndTickPriceRange(value decimal.Decimal, scale int32) {
	m.Set(field.NewEndTickPriceRange(value, scale))
}

//SetTickIncrement sets TickIncrement, Tag 1208
func (m NoTickRules) SetTickIncrement(value decimal.Decimal, scale int32) {
	m.Set(field.NewTickIncrement(value, scale))
}

//SetTickRuleType sets TickRuleType, Tag 1209
func (m NoTickRules) SetTickRuleType(v enum.TickRuleType) {
	m.Set(field.NewTickRuleType(v))
}

//GetStartTickPriceRange gets StartTickPriceRange, Tag 1206
func (m NoTickRules) GetStartTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StartTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEndTickPriceRange gets EndTickPriceRange, Tag 1207
func (m NoTickRules) GetEndTickPriceRange() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.EndTickPriceRangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTickIncrement gets TickIncrement, Tag 1208
func (m NoTickRules) GetTickIncrement() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TickIncrementField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTickRuleType gets TickRuleType, Tag 1209
func (m NoTickRules) GetTickRuleType() (v enum.TickRuleType, err quickfix.MessageRejectError) {
	var f field.TickRuleTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasStartTickPriceRange returns true if StartTickPriceRange is present, Tag 1206
func (m NoTickRules) HasStartTickPriceRange() bool {
	return m.Has(tag.StartTickPriceRange)
}

//HasEndTickPriceRange returns true if EndTickPriceRange is present, Tag 1207
func (m NoTickRules) HasEndTickPriceRange() bool {
	return m.Has(tag.EndTickPriceRange)
}

//HasTickIncrement returns true if TickIncrement is present, Tag 1208
func (m NoTickRules) HasTickIncrement() bool {
	return m.Has(tag.TickIncrement)
}

//HasTickRuleType returns true if TickRuleType is present, Tag 1209
func (m NoTickRules) HasTickRuleType() bool {
	return m.Has(tag.TickRuleType)
}

//NoTickRulesRepeatingGroup is a repeating group, Tag 1205
type NoTickRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoTickRulesRepeatingGroup returns an initialized, NoTickRulesRepeatingGroup
func NewNoTickRulesRepeatingGroup() NoTickRulesRepeatingGroup {
	return NoTickRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoTickRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.StartTickPriceRange), quickfix.GroupElement(tag.EndTickPriceRange), quickfix.GroupElement(tag.TickIncrement), quickfix.GroupElement(tag.TickRuleType)})}
}

//Add create and append a new NoTickRules to this group
func (m NoTickRulesRepeatingGroup) Add() NoTickRules {
	g := m.RepeatingGroup.Add()
	return NoTickRules{g}
}

//Get returns the ith NoTickRules in the NoTickRulesRepeatinGroup
func (m NoTickRulesRepeatingGroup) Get(i int) NoTickRules {
	return NoTickRules{m.RepeatingGroup.Get(i)}
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

//NoLotTypeRules is a repeating group element, Tag 1234
type NoLotTypeRules struct {
	*quickfix.Group
}

//SetLotType sets LotType, Tag 1093
func (m NoLotTypeRules) SetLotType(v enum.LotType) {
	m.Set(field.NewLotType(v))
}

//SetMinLotSize sets MinLotSize, Tag 1231
func (m NoLotTypeRules) SetMinLotSize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinLotSize(value, scale))
}

//GetLotType gets LotType, Tag 1093
func (m NoLotTypeRules) GetLotType() (v enum.LotType, err quickfix.MessageRejectError) {
	var f field.LotTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinLotSize gets MinLotSize, Tag 1231
func (m NoLotTypeRules) GetMinLotSize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinLotSizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasLotType returns true if LotType is present, Tag 1093
func (m NoLotTypeRules) HasLotType() bool {
	return m.Has(tag.LotType)
}

//HasMinLotSize returns true if MinLotSize is present, Tag 1231
func (m NoLotTypeRules) HasMinLotSize() bool {
	return m.Has(tag.MinLotSize)
}

//NoLotTypeRulesRepeatingGroup is a repeating group, Tag 1234
type NoLotTypeRulesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoLotTypeRulesRepeatingGroup returns an initialized, NoLotTypeRulesRepeatingGroup
func NewNoLotTypeRulesRepeatingGroup() NoLotTypeRulesRepeatingGroup {
	return NoLotTypeRulesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoLotTypeRules,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.LotType), quickfix.GroupElement(tag.MinLotSize)})}
}

//Add create and append a new NoLotTypeRules to this group
func (m NoLotTypeRulesRepeatingGroup) Add() NoLotTypeRules {
	g := m.RepeatingGroup.Add()
	return NoLotTypeRules{g}
}

//Get returns the ith NoLotTypeRules in the NoLotTypeRulesRepeatinGroup
func (m NoLotTypeRulesRepeatingGroup) Get(i int) NoLotTypeRules {
	return NoLotTypeRules{m.RepeatingGroup.Get(i)}
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
