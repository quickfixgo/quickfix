package marketdataincrementalrefresh

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/enum"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataIncrementalRefresh is the fix43 MarketDataIncrementalRefresh type, MsgType = X
type MarketDataIncrementalRefresh struct {
	fix43.Header
	*quickfix.Body
	fix43.Trailer
	Message *quickfix.Message
}

//FromMessage creates a MarketDataIncrementalRefresh from a quickfix.Message instance
func FromMessage(m *quickfix.Message) MarketDataIncrementalRefresh {
	return MarketDataIncrementalRefresh{
		Header:  fix43.Header{&m.Header},
		Body:    &m.Body,
		Trailer: fix43.Trailer{&m.Trailer},
		Message: m,
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataIncrementalRefresh) ToMessage() *quickfix.Message {
	return m.Message
}

//New returns a MarketDataIncrementalRefresh initialized with the required fields for MarketDataIncrementalRefresh
func New() (m MarketDataIncrementalRefresh) {
	m.Message = quickfix.NewMessage()
	m.Header = fix43.NewHeader(&m.Message.Header)
	m.Body = &m.Message.Body
	m.Trailer.Trailer = &m.Message.Trailer

	m.Header.Set(field.NewMsgType("X"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg *quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "X", r
}

//SetMDReqID sets MDReqID, Tag 262
func (m MarketDataIncrementalRefresh) SetMDReqID(v string) {
	m.Set(field.NewMDReqID(v))
}

//SetNoMDEntries sets NoMDEntries, Tag 268
func (m MarketDataIncrementalRefresh) SetNoMDEntries(f NoMDEntriesRepeatingGroup) {
	m.SetGroup(f)
}

//GetMDReqID gets MDReqID, Tag 262
func (m MarketDataIncrementalRefresh) GetMDReqID() (v string, err quickfix.MessageRejectError) {
	var f field.MDReqIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoMDEntries gets NoMDEntries, Tag 268
func (m MarketDataIncrementalRefresh) GetNoMDEntries() (NoMDEntriesRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoMDEntriesRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//HasMDReqID returns true if MDReqID is present, Tag 262
func (m MarketDataIncrementalRefresh) HasMDReqID() bool {
	return m.Has(tag.MDReqID)
}

//HasNoMDEntries returns true if NoMDEntries is present, Tag 268
func (m MarketDataIncrementalRefresh) HasNoMDEntries() bool {
	return m.Has(tag.NoMDEntries)
}

//NoMDEntries is a repeating group element, Tag 268
type NoMDEntries struct {
	*quickfix.Group
}

//SetMDUpdateAction sets MDUpdateAction, Tag 279
func (m NoMDEntries) SetMDUpdateAction(v enum.MDUpdateAction) {
	m.Set(field.NewMDUpdateAction(v))
}

//SetDeleteReason sets DeleteReason, Tag 285
func (m NoMDEntries) SetDeleteReason(v enum.DeleteReason) {
	m.Set(field.NewDeleteReason(v))
}

//SetMDEntryType sets MDEntryType, Tag 269
func (m NoMDEntries) SetMDEntryType(v enum.MDEntryType) {
	m.Set(field.NewMDEntryType(v))
}

//SetMDEntryID sets MDEntryID, Tag 278
func (m NoMDEntries) SetMDEntryID(v string) {
	m.Set(field.NewMDEntryID(v))
}

//SetMDEntryRefID sets MDEntryRefID, Tag 280
func (m NoMDEntries) SetMDEntryRefID(v string) {
	m.Set(field.NewMDEntryRefID(v))
}

//SetSymbol sets Symbol, Tag 55
func (m NoMDEntries) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m NoMDEntries) SetSymbolSfx(v enum.SymbolSfx) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoMDEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m NoMDEntries) SetSecurityIDSource(v enum.SecurityIDSource) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m NoMDEntries) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m NoMDEntries) SetProduct(v enum.Product) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m NoMDEntries) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoMDEntries) SetSecurityType(v enum.SecurityType) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoMDEntries) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m NoMDEntries) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m NoMDEntries) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m NoMDEntries) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m NoMDEntries) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m NoMDEntries) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m NoMDEntries) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m NoMDEntries) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetCreditRating sets CreditRating, Tag 255
func (m NoMDEntries) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m NoMDEntries) SetInstrRegistry(v enum.InstrRegistry) {
	m.Set(field.NewInstrRegistry(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m NoMDEntries) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m NoMDEntries) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m NoMDEntries) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m NoMDEntries) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoMDEntries) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoMDEntries) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoMDEntries) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoMDEntries) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m NoMDEntries) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetIssuer sets Issuer, Tag 106
func (m NoMDEntries) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m NoMDEntries) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m NoMDEntries) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m NoMDEntries) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m NoMDEntries) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m NoMDEntries) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetFinancialStatus sets FinancialStatus, Tag 291
func (m NoMDEntries) SetFinancialStatus(v enum.FinancialStatus) {
	m.Set(field.NewFinancialStatus(v))
}

//SetCorporateAction sets CorporateAction, Tag 292
func (m NoMDEntries) SetCorporateAction(v enum.CorporateAction) {
	m.Set(field.NewCorporateAction(v))
}

//SetMDEntryPx sets MDEntryPx, Tag 270
func (m NoMDEntries) SetMDEntryPx(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntryPx(value, scale))
}

//SetCurrency sets Currency, Tag 15
func (m NoMDEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetMDEntrySize sets MDEntrySize, Tag 271
func (m NoMDEntries) SetMDEntrySize(value decimal.Decimal, scale int32) {
	m.Set(field.NewMDEntrySize(value, scale))
}

//SetMDEntryDate sets MDEntryDate, Tag 272
func (m NoMDEntries) SetMDEntryDate(v string) {
	m.Set(field.NewMDEntryDate(v))
}

//SetMDEntryTime sets MDEntryTime, Tag 273
func (m NoMDEntries) SetMDEntryTime(v string) {
	m.Set(field.NewMDEntryTime(v))
}

//SetTickDirection sets TickDirection, Tag 274
func (m NoMDEntries) SetTickDirection(v enum.TickDirection) {
	m.Set(field.NewTickDirection(v))
}

//SetMDMkt sets MDMkt, Tag 275
func (m NoMDEntries) SetMDMkt(v string) {
	m.Set(field.NewMDMkt(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoMDEntries) SetTradingSessionID(v enum.TradingSessionID) {
	m.Set(field.NewTradingSessionID(v))
}

//SetTradingSessionSubID sets TradingSessionSubID, Tag 625
func (m NoMDEntries) SetTradingSessionSubID(v enum.TradingSessionSubID) {
	m.Set(field.NewTradingSessionSubID(v))
}

//SetQuoteCondition sets QuoteCondition, Tag 276
func (m NoMDEntries) SetQuoteCondition(v enum.QuoteCondition) {
	m.Set(field.NewQuoteCondition(v))
}

//SetTradeCondition sets TradeCondition, Tag 277
func (m NoMDEntries) SetTradeCondition(v enum.TradeCondition) {
	m.Set(field.NewTradeCondition(v))
}

//SetMDEntryOriginator sets MDEntryOriginator, Tag 282
func (m NoMDEntries) SetMDEntryOriginator(v string) {
	m.Set(field.NewMDEntryOriginator(v))
}

//SetLocationID sets LocationID, Tag 283
func (m NoMDEntries) SetLocationID(v string) {
	m.Set(field.NewLocationID(v))
}

//SetDeskID sets DeskID, Tag 284
func (m NoMDEntries) SetDeskID(v string) {
	m.Set(field.NewDeskID(v))
}

//SetOpenCloseSettleFlag sets OpenCloseSettleFlag, Tag 286
func (m NoMDEntries) SetOpenCloseSettleFlag(v enum.OpenCloseSettleFlag) {
	m.Set(field.NewOpenCloseSettleFlag(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NoMDEntries) SetTimeInForce(v enum.TimeInForce) {
	m.Set(field.NewTimeInForce(v))
}

//SetExpireDate sets ExpireDate, Tag 432
func (m NoMDEntries) SetExpireDate(v string) {
	m.Set(field.NewExpireDate(v))
}

//SetExpireTime sets ExpireTime, Tag 126
func (m NoMDEntries) SetExpireTime(v time.Time) {
	m.Set(field.NewExpireTime(v))
}

//SetMinQty sets MinQty, Tag 110
func (m NoMDEntries) SetMinQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewMinQty(value, scale))
}

//SetExecInst sets ExecInst, Tag 18
func (m NoMDEntries) SetExecInst(v enum.ExecInst) {
	m.Set(field.NewExecInst(v))
}

//SetSellerDays sets SellerDays, Tag 287
func (m NoMDEntries) SetSellerDays(v int) {
	m.Set(field.NewSellerDays(v))
}

//SetOrderID sets OrderID, Tag 37
func (m NoMDEntries) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetQuoteEntryID sets QuoteEntryID, Tag 299
func (m NoMDEntries) SetQuoteEntryID(v string) {
	m.Set(field.NewQuoteEntryID(v))
}

//SetMDEntryBuyer sets MDEntryBuyer, Tag 288
func (m NoMDEntries) SetMDEntryBuyer(v string) {
	m.Set(field.NewMDEntryBuyer(v))
}

//SetMDEntrySeller sets MDEntrySeller, Tag 289
func (m NoMDEntries) SetMDEntrySeller(v string) {
	m.Set(field.NewMDEntrySeller(v))
}

//SetNumberOfOrders sets NumberOfOrders, Tag 346
func (m NoMDEntries) SetNumberOfOrders(v int) {
	m.Set(field.NewNumberOfOrders(v))
}

//SetMDEntryPositionNo sets MDEntryPositionNo, Tag 290
func (m NoMDEntries) SetMDEntryPositionNo(v int) {
	m.Set(field.NewMDEntryPositionNo(v))
}

//SetScope sets Scope, Tag 546
func (m NoMDEntries) SetScope(v enum.Scope) {
	m.Set(field.NewScope(v))
}

//SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387
func (m NoMDEntries) SetTotalVolumeTraded(value decimal.Decimal, scale int32) {
	m.Set(field.NewTotalVolumeTraded(value, scale))
}

//SetTotalVolumeTradedDate sets TotalVolumeTradedDate, Tag 449
func (m NoMDEntries) SetTotalVolumeTradedDate(v string) {
	m.Set(field.NewTotalVolumeTradedDate(v))
}

//SetTotalVolumeTradedTime sets TotalVolumeTradedTime, Tag 450
func (m NoMDEntries) SetTotalVolumeTradedTime(v string) {
	m.Set(field.NewTotalVolumeTradedTime(v))
}

//SetNetChgPrevDay sets NetChgPrevDay, Tag 451
func (m NoMDEntries) SetNetChgPrevDay(value decimal.Decimal, scale int32) {
	m.Set(field.NewNetChgPrevDay(value, scale))
}

//SetText sets Text, Tag 58
func (m NoMDEntries) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m NoMDEntries) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m NoMDEntries) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//GetMDUpdateAction gets MDUpdateAction, Tag 279
func (m NoMDEntries) GetMDUpdateAction() (v enum.MDUpdateAction, err quickfix.MessageRejectError) {
	var f field.MDUpdateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeleteReason gets DeleteReason, Tag 285
func (m NoMDEntries) GetDeleteReason() (v enum.DeleteReason, err quickfix.MessageRejectError) {
	var f field.DeleteReasonField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryType gets MDEntryType, Tag 269
func (m NoMDEntries) GetMDEntryType() (v enum.MDEntryType, err quickfix.MessageRejectError) {
	var f field.MDEntryTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryID gets MDEntryID, Tag 278
func (m NoMDEntries) GetMDEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryRefID gets MDEntryRefID, Tag 280
func (m NoMDEntries) GetMDEntryRefID() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryRefIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NoMDEntries) GetSymbol() (v string, err quickfix.MessageRejectError) {
	var f field.SymbolField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoMDEntries) GetSymbolSfx() (v enum.SymbolSfx, err quickfix.MessageRejectError) {
	var f field.SymbolSfxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoMDEntries) GetSecurityID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m NoMDEntries) GetSecurityIDSource() (v enum.SecurityIDSource, err quickfix.MessageRejectError) {
	var f field.SecurityIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m NoMDEntries) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m NoMDEntries) GetProduct() (v enum.Product, err quickfix.MessageRejectError) {
	var f field.ProductField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCFICode gets CFICode, Tag 461
func (m NoMDEntries) GetCFICode() (v string, err quickfix.MessageRejectError) {
	var f field.CFICodeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoMDEntries) GetSecurityType() (v enum.SecurityType, err quickfix.MessageRejectError) {
	var f field.SecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoMDEntries) GetMaturityMonthYear() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityMonthYearField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m NoMDEntries) GetMaturityDate() (v string, err quickfix.MessageRejectError) {
	var f field.MaturityDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m NoMDEntries) GetCouponPaymentDate() (v string, err quickfix.MessageRejectError) {
	var f field.CouponPaymentDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m NoMDEntries) GetIssueDate() (v string, err quickfix.MessageRejectError) {
	var f field.IssueDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m NoMDEntries) GetRepoCollateralSecurityType() (v int, err quickfix.MessageRejectError) {
	var f field.RepoCollateralSecurityTypeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m NoMDEntries) GetRepurchaseTerm() (v int, err quickfix.MessageRejectError) {
	var f field.RepurchaseTermField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m NoMDEntries) GetRepurchaseRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.RepurchaseRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFactor gets Factor, Tag 228
func (m NoMDEntries) GetFactor() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.FactorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m NoMDEntries) GetCreditRating() (v string, err quickfix.MessageRejectError) {
	var f field.CreditRatingField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m NoMDEntries) GetInstrRegistry() (v enum.InstrRegistry, err quickfix.MessageRejectError) {
	var f field.InstrRegistryField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m NoMDEntries) GetCountryOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.CountryOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m NoMDEntries) GetStateOrProvinceOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.StateOrProvinceOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m NoMDEntries) GetLocaleOfIssue() (v string, err quickfix.MessageRejectError) {
	var f field.LocaleOfIssueField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m NoMDEntries) GetRedemptionDate() (v string, err quickfix.MessageRejectError) {
	var f field.RedemptionDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoMDEntries) GetStrikePrice() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.StrikePriceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoMDEntries) GetOptAttribute() (v string, err quickfix.MessageRejectError) {
	var f field.OptAttributeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoMDEntries) GetContractMultiplier() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.ContractMultiplierField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoMDEntries) GetCouponRate() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.CouponRateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoMDEntries) GetSecurityExchange() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityExchangeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoMDEntries) GetIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.IssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoMDEntries) GetEncodedIssuerLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoMDEntries) GetEncodedIssuer() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedIssuerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoMDEntries) GetSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoMDEntries) GetEncodedSecurityDescLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoMDEntries) GetEncodedSecurityDesc() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedSecurityDescField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetFinancialStatus gets FinancialStatus, Tag 291
func (m NoMDEntries) GetFinancialStatus() (v enum.FinancialStatus, err quickfix.MessageRejectError) {
	var f field.FinancialStatusField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCorporateAction gets CorporateAction, Tag 292
func (m NoMDEntries) GetCorporateAction() (v enum.CorporateAction, err quickfix.MessageRejectError) {
	var f field.CorporateActionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryPx gets MDEntryPx, Tag 270
func (m NoMDEntries) GetMDEntryPx() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntryPxField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoMDEntries) GetCurrency() (v string, err quickfix.MessageRejectError) {
	var f field.CurrencyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntrySize gets MDEntrySize, Tag 271
func (m NoMDEntries) GetMDEntrySize() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MDEntrySizeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryDate gets MDEntryDate, Tag 272
func (m NoMDEntries) GetMDEntryDate() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryTime gets MDEntryTime, Tag 273
func (m NoMDEntries) GetMDEntryTime() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTickDirection gets TickDirection, Tag 274
func (m NoMDEntries) GetTickDirection() (v enum.TickDirection, err quickfix.MessageRejectError) {
	var f field.TickDirectionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDMkt gets MDMkt, Tag 275
func (m NoMDEntries) GetMDMkt() (v string, err quickfix.MessageRejectError) {
	var f field.MDMktField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoMDEntries) GetTradingSessionID() (v enum.TradingSessionID, err quickfix.MessageRejectError) {
	var f field.TradingSessionIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradingSessionSubID gets TradingSessionSubID, Tag 625
func (m NoMDEntries) GetTradingSessionSubID() (v enum.TradingSessionSubID, err quickfix.MessageRejectError) {
	var f field.TradingSessionSubIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteCondition gets QuoteCondition, Tag 276
func (m NoMDEntries) GetQuoteCondition() (v enum.QuoteCondition, err quickfix.MessageRejectError) {
	var f field.QuoteConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTradeCondition gets TradeCondition, Tag 277
func (m NoMDEntries) GetTradeCondition() (v enum.TradeCondition, err quickfix.MessageRejectError) {
	var f field.TradeConditionField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryOriginator gets MDEntryOriginator, Tag 282
func (m NoMDEntries) GetMDEntryOriginator() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryOriginatorField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetLocationID gets LocationID, Tag 283
func (m NoMDEntries) GetLocationID() (v string, err quickfix.MessageRejectError) {
	var f field.LocationIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetDeskID gets DeskID, Tag 284
func (m NoMDEntries) GetDeskID() (v string, err quickfix.MessageRejectError) {
	var f field.DeskIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOpenCloseSettleFlag gets OpenCloseSettleFlag, Tag 286
func (m NoMDEntries) GetOpenCloseSettleFlag() (v enum.OpenCloseSettleFlag, err quickfix.MessageRejectError) {
	var f field.OpenCloseSettleFlagField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NoMDEntries) GetTimeInForce() (v enum.TimeInForce, err quickfix.MessageRejectError) {
	var f field.TimeInForceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m NoMDEntries) GetExpireDate() (v string, err quickfix.MessageRejectError) {
	var f field.ExpireDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoMDEntries) GetExpireTime() (v time.Time, err quickfix.MessageRejectError) {
	var f field.ExpireTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NoMDEntries) GetMinQty() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.MinQtyField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NoMDEntries) GetExecInst() (v enum.ExecInst, err quickfix.MessageRejectError) {
	var f field.ExecInstField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSellerDays gets SellerDays, Tag 287
func (m NoMDEntries) GetSellerDays() (v int, err quickfix.MessageRejectError) {
	var f field.SellerDaysField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoMDEntries) GetOrderID() (v string, err quickfix.MessageRejectError) {
	var f field.OrderIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetQuoteEntryID gets QuoteEntryID, Tag 299
func (m NoMDEntries) GetQuoteEntryID() (v string, err quickfix.MessageRejectError) {
	var f field.QuoteEntryIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryBuyer gets MDEntryBuyer, Tag 288
func (m NoMDEntries) GetMDEntryBuyer() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntryBuyerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntrySeller gets MDEntrySeller, Tag 289
func (m NoMDEntries) GetMDEntrySeller() (v string, err quickfix.MessageRejectError) {
	var f field.MDEntrySellerField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNumberOfOrders gets NumberOfOrders, Tag 346
func (m NoMDEntries) GetNumberOfOrders() (v int, err quickfix.MessageRejectError) {
	var f field.NumberOfOrdersField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetMDEntryPositionNo gets MDEntryPositionNo, Tag 290
func (m NoMDEntries) GetMDEntryPositionNo() (v int, err quickfix.MessageRejectError) {
	var f field.MDEntryPositionNoField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetScope gets Scope, Tag 546
func (m NoMDEntries) GetScope() (v enum.Scope, err quickfix.MessageRejectError) {
	var f field.ScopeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m NoMDEntries) GetTotalVolumeTraded() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalVolumeTradedDate gets TotalVolumeTradedDate, Tag 449
func (m NoMDEntries) GetTotalVolumeTradedDate() (v string, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedDateField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetTotalVolumeTradedTime gets TotalVolumeTradedTime, Tag 450
func (m NoMDEntries) GetTotalVolumeTradedTime() (v string, err quickfix.MessageRejectError) {
	var f field.TotalVolumeTradedTimeField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetNetChgPrevDay gets NetChgPrevDay, Tag 451
func (m NoMDEntries) GetNetChgPrevDay() (v decimal.Decimal, err quickfix.MessageRejectError) {
	var f field.NetChgPrevDayField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetText gets Text, Tag 58
func (m NoMDEntries) GetText() (v string, err quickfix.MessageRejectError) {
	var f field.TextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoMDEntries) GetEncodedTextLen() (v int, err quickfix.MessageRejectError) {
	var f field.EncodedTextLenField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoMDEntries) GetEncodedText() (v string, err quickfix.MessageRejectError) {
	var f field.EncodedTextField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasMDUpdateAction returns true if MDUpdateAction is present, Tag 279
func (m NoMDEntries) HasMDUpdateAction() bool {
	return m.Has(tag.MDUpdateAction)
}

//HasDeleteReason returns true if DeleteReason is present, Tag 285
func (m NoMDEntries) HasDeleteReason() bool {
	return m.Has(tag.DeleteReason)
}

//HasMDEntryType returns true if MDEntryType is present, Tag 269
func (m NoMDEntries) HasMDEntryType() bool {
	return m.Has(tag.MDEntryType)
}

//HasMDEntryID returns true if MDEntryID is present, Tag 278
func (m NoMDEntries) HasMDEntryID() bool {
	return m.Has(tag.MDEntryID)
}

//HasMDEntryRefID returns true if MDEntryRefID is present, Tag 280
func (m NoMDEntries) HasMDEntryRefID() bool {
	return m.Has(tag.MDEntryRefID)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m NoMDEntries) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m NoMDEntries) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m NoMDEntries) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m NoMDEntries) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m NoMDEntries) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m NoMDEntries) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m NoMDEntries) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoMDEntries) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoMDEntries) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m NoMDEntries) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m NoMDEntries) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m NoMDEntries) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m NoMDEntries) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m NoMDEntries) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m NoMDEntries) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m NoMDEntries) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m NoMDEntries) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m NoMDEntries) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m NoMDEntries) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m NoMDEntries) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m NoMDEntries) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m NoMDEntries) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m NoMDEntries) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m NoMDEntries) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m NoMDEntries) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m NoMDEntries) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m NoMDEntries) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m NoMDEntries) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m NoMDEntries) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m NoMDEntries) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m NoMDEntries) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m NoMDEntries) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m NoMDEntries) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasFinancialStatus returns true if FinancialStatus is present, Tag 291
func (m NoMDEntries) HasFinancialStatus() bool {
	return m.Has(tag.FinancialStatus)
}

//HasCorporateAction returns true if CorporateAction is present, Tag 292
func (m NoMDEntries) HasCorporateAction() bool {
	return m.Has(tag.CorporateAction)
}

//HasMDEntryPx returns true if MDEntryPx is present, Tag 270
func (m NoMDEntries) HasMDEntryPx() bool {
	return m.Has(tag.MDEntryPx)
}

//HasCurrency returns true if Currency is present, Tag 15
func (m NoMDEntries) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasMDEntrySize returns true if MDEntrySize is present, Tag 271
func (m NoMDEntries) HasMDEntrySize() bool {
	return m.Has(tag.MDEntrySize)
}

//HasMDEntryDate returns true if MDEntryDate is present, Tag 272
func (m NoMDEntries) HasMDEntryDate() bool {
	return m.Has(tag.MDEntryDate)
}

//HasMDEntryTime returns true if MDEntryTime is present, Tag 273
func (m NoMDEntries) HasMDEntryTime() bool {
	return m.Has(tag.MDEntryTime)
}

//HasTickDirection returns true if TickDirection is present, Tag 274
func (m NoMDEntries) HasTickDirection() bool {
	return m.Has(tag.TickDirection)
}

//HasMDMkt returns true if MDMkt is present, Tag 275
func (m NoMDEntries) HasMDMkt() bool {
	return m.Has(tag.MDMkt)
}

//HasTradingSessionID returns true if TradingSessionID is present, Tag 336
func (m NoMDEntries) HasTradingSessionID() bool {
	return m.Has(tag.TradingSessionID)
}

//HasTradingSessionSubID returns true if TradingSessionSubID is present, Tag 625
func (m NoMDEntries) HasTradingSessionSubID() bool {
	return m.Has(tag.TradingSessionSubID)
}

//HasQuoteCondition returns true if QuoteCondition is present, Tag 276
func (m NoMDEntries) HasQuoteCondition() bool {
	return m.Has(tag.QuoteCondition)
}

//HasTradeCondition returns true if TradeCondition is present, Tag 277
func (m NoMDEntries) HasTradeCondition() bool {
	return m.Has(tag.TradeCondition)
}

//HasMDEntryOriginator returns true if MDEntryOriginator is present, Tag 282
func (m NoMDEntries) HasMDEntryOriginator() bool {
	return m.Has(tag.MDEntryOriginator)
}

//HasLocationID returns true if LocationID is present, Tag 283
func (m NoMDEntries) HasLocationID() bool {
	return m.Has(tag.LocationID)
}

//HasDeskID returns true if DeskID is present, Tag 284
func (m NoMDEntries) HasDeskID() bool {
	return m.Has(tag.DeskID)
}

//HasOpenCloseSettleFlag returns true if OpenCloseSettleFlag is present, Tag 286
func (m NoMDEntries) HasOpenCloseSettleFlag() bool {
	return m.Has(tag.OpenCloseSettleFlag)
}

//HasTimeInForce returns true if TimeInForce is present, Tag 59
func (m NoMDEntries) HasTimeInForce() bool {
	return m.Has(tag.TimeInForce)
}

//HasExpireDate returns true if ExpireDate is present, Tag 432
func (m NoMDEntries) HasExpireDate() bool {
	return m.Has(tag.ExpireDate)
}

//HasExpireTime returns true if ExpireTime is present, Tag 126
func (m NoMDEntries) HasExpireTime() bool {
	return m.Has(tag.ExpireTime)
}

//HasMinQty returns true if MinQty is present, Tag 110
func (m NoMDEntries) HasMinQty() bool {
	return m.Has(tag.MinQty)
}

//HasExecInst returns true if ExecInst is present, Tag 18
func (m NoMDEntries) HasExecInst() bool {
	return m.Has(tag.ExecInst)
}

//HasSellerDays returns true if SellerDays is present, Tag 287
func (m NoMDEntries) HasSellerDays() bool {
	return m.Has(tag.SellerDays)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m NoMDEntries) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasQuoteEntryID returns true if QuoteEntryID is present, Tag 299
func (m NoMDEntries) HasQuoteEntryID() bool {
	return m.Has(tag.QuoteEntryID)
}

//HasMDEntryBuyer returns true if MDEntryBuyer is present, Tag 288
func (m NoMDEntries) HasMDEntryBuyer() bool {
	return m.Has(tag.MDEntryBuyer)
}

//HasMDEntrySeller returns true if MDEntrySeller is present, Tag 289
func (m NoMDEntries) HasMDEntrySeller() bool {
	return m.Has(tag.MDEntrySeller)
}

//HasNumberOfOrders returns true if NumberOfOrders is present, Tag 346
func (m NoMDEntries) HasNumberOfOrders() bool {
	return m.Has(tag.NumberOfOrders)
}

//HasMDEntryPositionNo returns true if MDEntryPositionNo is present, Tag 290
func (m NoMDEntries) HasMDEntryPositionNo() bool {
	return m.Has(tag.MDEntryPositionNo)
}

//HasScope returns true if Scope is present, Tag 546
func (m NoMDEntries) HasScope() bool {
	return m.Has(tag.Scope)
}

//HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387
func (m NoMDEntries) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
}

//HasTotalVolumeTradedDate returns true if TotalVolumeTradedDate is present, Tag 449
func (m NoMDEntries) HasTotalVolumeTradedDate() bool {
	return m.Has(tag.TotalVolumeTradedDate)
}

//HasTotalVolumeTradedTime returns true if TotalVolumeTradedTime is present, Tag 450
func (m NoMDEntries) HasTotalVolumeTradedTime() bool {
	return m.Has(tag.TotalVolumeTradedTime)
}

//HasNetChgPrevDay returns true if NetChgPrevDay is present, Tag 451
func (m NoMDEntries) HasNetChgPrevDay() bool {
	return m.Has(tag.NetChgPrevDay)
}

//HasText returns true if Text is present, Tag 58
func (m NoMDEntries) HasText() bool {
	return m.Has(tag.Text)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m NoMDEntries) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m NoMDEntries) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	*quickfix.Group
}

//SetSecurityAltID sets SecurityAltID, Tag 455
func (m NoSecurityAltID) SetSecurityAltID(v string) {
	m.Set(field.NewSecurityAltID(v))
}

//SetSecurityAltIDSource sets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) SetSecurityAltIDSource(v string) {
	m.Set(field.NewSecurityAltIDSource(v))
}

//GetSecurityAltID gets SecurityAltID, Tag 455
func (m NoSecurityAltID) GetSecurityAltID() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (v string, err quickfix.MessageRejectError) {
	var f field.SecurityAltIDSourceField
	if err = m.Get(&f); err == nil {
		v = f.Value()
	}
	return
}

//HasSecurityAltID returns true if SecurityAltID is present, Tag 455
func (m NoSecurityAltID) HasSecurityAltID() bool {
	return m.Has(tag.SecurityAltID)
}

//HasSecurityAltIDSource returns true if SecurityAltIDSource is present, Tag 456
func (m NoSecurityAltID) HasSecurityAltIDSource() bool {
	return m.Has(tag.SecurityAltIDSource)
}

//NoSecurityAltIDRepeatingGroup is a repeating group, Tag 454
type NoSecurityAltIDRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoSecurityAltIDRepeatingGroup returns an initialized, NoSecurityAltIDRepeatingGroup
func NewNoSecurityAltIDRepeatingGroup() NoSecurityAltIDRepeatingGroup {
	return NoSecurityAltIDRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoSecurityAltID,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.SecurityAltID), quickfix.GroupElement(tag.SecurityAltIDSource)})}
}

//Add create and append a new NoSecurityAltID to this group
func (m NoSecurityAltIDRepeatingGroup) Add() NoSecurityAltID {
	g := m.RepeatingGroup.Add()
	return NoSecurityAltID{g}
}

//Get returns the ith NoSecurityAltID in the NoSecurityAltIDRepeatinGroup
func (m NoSecurityAltIDRepeatingGroup) Get(i int) NoSecurityAltID {
	return NoSecurityAltID{m.RepeatingGroup.Get(i)}
}

//NoMDEntriesRepeatingGroup is a repeating group, Tag 268
type NoMDEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntriesRepeatingGroup returns an initialized, NoMDEntriesRepeatingGroup
func NewNoMDEntriesRepeatingGroup() NoMDEntriesRepeatingGroup {
	return NoMDEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDEntries,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDUpdateAction), quickfix.GroupElement(tag.DeleteReason), quickfix.GroupElement(tag.MDEntryType), quickfix.GroupElement(tag.MDEntryID), quickfix.GroupElement(tag.MDEntryRefID), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.SecurityIDSource), NewNoSecurityAltIDRepeatingGroup(), quickfix.GroupElement(tag.Product), quickfix.GroupElement(tag.CFICode), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDate), quickfix.GroupElement(tag.CouponPaymentDate), quickfix.GroupElement(tag.IssueDate), quickfix.GroupElement(tag.RepoCollateralSecurityType), quickfix.GroupElement(tag.RepurchaseTerm), quickfix.GroupElement(tag.RepurchaseRate), quickfix.GroupElement(tag.Factor), quickfix.GroupElement(tag.CreditRating), quickfix.GroupElement(tag.InstrRegistry), quickfix.GroupElement(tag.CountryOfIssue), quickfix.GroupElement(tag.StateOrProvinceOfIssue), quickfix.GroupElement(tag.LocaleOfIssue), quickfix.GroupElement(tag.RedemptionDate), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.FinancialStatus), quickfix.GroupElement(tag.CorporateAction), quickfix.GroupElement(tag.MDEntryPx), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.MDEntrySize), quickfix.GroupElement(tag.MDEntryDate), quickfix.GroupElement(tag.MDEntryTime), quickfix.GroupElement(tag.TickDirection), quickfix.GroupElement(tag.MDMkt), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.TradingSessionSubID), quickfix.GroupElement(tag.QuoteCondition), quickfix.GroupElement(tag.TradeCondition), quickfix.GroupElement(tag.MDEntryOriginator), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID), quickfix.GroupElement(tag.OpenCloseSettleFlag), quickfix.GroupElement(tag.TimeInForce), quickfix.GroupElement(tag.ExpireDate), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.MinQty), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.SellerDays), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.QuoteEntryID), quickfix.GroupElement(tag.MDEntryBuyer), quickfix.GroupElement(tag.MDEntrySeller), quickfix.GroupElement(tag.NumberOfOrders), quickfix.GroupElement(tag.MDEntryPositionNo), quickfix.GroupElement(tag.Scope), quickfix.GroupElement(tag.TotalVolumeTraded), quickfix.GroupElement(tag.TotalVolumeTradedDate), quickfix.GroupElement(tag.TotalVolumeTradedTime), quickfix.GroupElement(tag.NetChgPrevDay), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
}

//Add create and append a new NoMDEntries to this group
func (m NoMDEntriesRepeatingGroup) Add() NoMDEntries {
	g := m.RepeatingGroup.Add()
	return NoMDEntries{g}
}

//Get returns the ith NoMDEntries in the NoMDEntriesRepeatinGroup
func (m NoMDEntriesRepeatingGroup) Get(i int) NoMDEntries {
	return NoMDEntries{m.RepeatingGroup.Get(i)}
}
