package marketdataincrementalrefresh

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix42"
	"github.com/quickfixgo/quickfix/tag"
)

//MarketDataIncrementalRefresh is the fix42 MarketDataIncrementalRefresh type, MsgType = X
type MarketDataIncrementalRefresh struct {
	fix42.Header
	quickfix.Body
	fix42.Trailer
}

//FromMessage creates a MarketDataIncrementalRefresh from a quickfix.Message instance
func FromMessage(m quickfix.Message) MarketDataIncrementalRefresh {
	return MarketDataIncrementalRefresh{
		Header:  fix42.Header{Header: m.Header},
		Body:    m.Body,
		Trailer: fix42.Trailer{Trailer: m.Trailer},
	}
}

//ToMessage returns a quickfix.Message instance
func (m MarketDataIncrementalRefresh) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:  m.Header.Header,
		Body:    m.Body,
		Trailer: m.Trailer.Trailer,
	}
}

//New returns a MarketDataIncrementalRefresh initialized with the required fields for MarketDataIncrementalRefresh
func New() (m MarketDataIncrementalRefresh) {
	m.Header.Init()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("X"))

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg MarketDataIncrementalRefresh, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.2", "X", r
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
func (m MarketDataIncrementalRefresh) GetMDReqID() (f field.MDReqIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
	quickfix.Group
}

//SetMDUpdateAction sets MDUpdateAction, Tag 279
func (m NoMDEntries) SetMDUpdateAction(v string) {
	m.Set(field.NewMDUpdateAction(v))
}

//SetDeleteReason sets DeleteReason, Tag 285
func (m NoMDEntries) SetDeleteReason(v string) {
	m.Set(field.NewDeleteReason(v))
}

//SetMDEntryType sets MDEntryType, Tag 269
func (m NoMDEntries) SetMDEntryType(v string) {
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
func (m NoMDEntries) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m NoMDEntries) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetIDSource sets IDSource, Tag 22
func (m NoMDEntries) SetIDSource(v string) {
	m.Set(field.NewIDSource(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m NoMDEntries) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m NoMDEntries) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetMaturityDay sets MaturityDay, Tag 205
func (m NoMDEntries) SetMaturityDay(v int) {
	m.Set(field.NewMaturityDay(v))
}

//SetPutOrCall sets PutOrCall, Tag 201
func (m NoMDEntries) SetPutOrCall(v int) {
	m.Set(field.NewPutOrCall(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m NoMDEntries) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m NoMDEntries) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m NoMDEntries) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m NoMDEntries) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
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
func (m NoMDEntries) SetFinancialStatus(v string) {
	m.Set(field.NewFinancialStatus(v))
}

//SetCorporateAction sets CorporateAction, Tag 292
func (m NoMDEntries) SetCorporateAction(v string) {
	m.Set(field.NewCorporateAction(v))
}

//SetMDEntryPx sets MDEntryPx, Tag 270
func (m NoMDEntries) SetMDEntryPx(v float64) {
	m.Set(field.NewMDEntryPx(v))
}

//SetCurrency sets Currency, Tag 15
func (m NoMDEntries) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetMDEntrySize sets MDEntrySize, Tag 271
func (m NoMDEntries) SetMDEntrySize(v float64) {
	m.Set(field.NewMDEntrySize(v))
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
func (m NoMDEntries) SetTickDirection(v string) {
	m.Set(field.NewTickDirection(v))
}

//SetMDMkt sets MDMkt, Tag 275
func (m NoMDEntries) SetMDMkt(v string) {
	m.Set(field.NewMDMkt(v))
}

//SetTradingSessionID sets TradingSessionID, Tag 336
func (m NoMDEntries) SetTradingSessionID(v string) {
	m.Set(field.NewTradingSessionID(v))
}

//SetQuoteCondition sets QuoteCondition, Tag 276
func (m NoMDEntries) SetQuoteCondition(v string) {
	m.Set(field.NewQuoteCondition(v))
}

//SetTradeCondition sets TradeCondition, Tag 277
func (m NoMDEntries) SetTradeCondition(v string) {
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
func (m NoMDEntries) SetOpenCloseSettleFlag(v string) {
	m.Set(field.NewOpenCloseSettleFlag(v))
}

//SetTimeInForce sets TimeInForce, Tag 59
func (m NoMDEntries) SetTimeInForce(v string) {
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
func (m NoMDEntries) SetMinQty(v float64) {
	m.Set(field.NewMinQty(v))
}

//SetExecInst sets ExecInst, Tag 18
func (m NoMDEntries) SetExecInst(v string) {
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

//SetTotalVolumeTraded sets TotalVolumeTraded, Tag 387
func (m NoMDEntries) SetTotalVolumeTraded(v float64) {
	m.Set(field.NewTotalVolumeTraded(v))
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
func (m NoMDEntries) GetMDUpdateAction() (f field.MDUpdateActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeleteReason gets DeleteReason, Tag 285
func (m NoMDEntries) GetDeleteReason() (f field.DeleteReasonField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryType gets MDEntryType, Tag 269
func (m NoMDEntries) GetMDEntryType() (f field.MDEntryTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryID gets MDEntryID, Tag 278
func (m NoMDEntries) GetMDEntryID() (f field.MDEntryIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryRefID gets MDEntryRefID, Tag 280
func (m NoMDEntries) GetMDEntryRefID() (f field.MDEntryRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m NoMDEntries) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m NoMDEntries) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m NoMDEntries) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIDSource gets IDSource, Tag 22
func (m NoMDEntries) GetIDSource() (f field.IDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m NoMDEntries) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m NoMDEntries) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDay gets MaturityDay, Tag 205
func (m NoMDEntries) GetMaturityDay() (f field.MaturityDayField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPutOrCall gets PutOrCall, Tag 201
func (m NoMDEntries) GetPutOrCall() (f field.PutOrCallField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m NoMDEntries) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m NoMDEntries) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m NoMDEntries) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m NoMDEntries) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m NoMDEntries) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m NoMDEntries) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m NoMDEntries) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m NoMDEntries) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m NoMDEntries) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m NoMDEntries) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m NoMDEntries) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFinancialStatus gets FinancialStatus, Tag 291
func (m NoMDEntries) GetFinancialStatus() (f field.FinancialStatusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCorporateAction gets CorporateAction, Tag 292
func (m NoMDEntries) GetCorporateAction() (f field.CorporateActionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryPx gets MDEntryPx, Tag 270
func (m NoMDEntries) GetMDEntryPx() (f field.MDEntryPxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCurrency gets Currency, Tag 15
func (m NoMDEntries) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntrySize gets MDEntrySize, Tag 271
func (m NoMDEntries) GetMDEntrySize() (f field.MDEntrySizeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryDate gets MDEntryDate, Tag 272
func (m NoMDEntries) GetMDEntryDate() (f field.MDEntryDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryTime gets MDEntryTime, Tag 273
func (m NoMDEntries) GetMDEntryTime() (f field.MDEntryTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTickDirection gets TickDirection, Tag 274
func (m NoMDEntries) GetTickDirection() (f field.TickDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDMkt gets MDMkt, Tag 275
func (m NoMDEntries) GetMDMkt() (f field.MDMktField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradingSessionID gets TradingSessionID, Tag 336
func (m NoMDEntries) GetTradingSessionID() (f field.TradingSessionIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteCondition gets QuoteCondition, Tag 276
func (m NoMDEntries) GetQuoteCondition() (f field.QuoteConditionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTradeCondition gets TradeCondition, Tag 277
func (m NoMDEntries) GetTradeCondition() (f field.TradeConditionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryOriginator gets MDEntryOriginator, Tag 282
func (m NoMDEntries) GetMDEntryOriginator() (f field.MDEntryOriginatorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocationID gets LocationID, Tag 283
func (m NoMDEntries) GetLocationID() (f field.LocationIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetDeskID gets DeskID, Tag 284
func (m NoMDEntries) GetDeskID() (f field.DeskIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOpenCloseSettleFlag gets OpenCloseSettleFlag, Tag 286
func (m NoMDEntries) GetOpenCloseSettleFlag() (f field.OpenCloseSettleFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTimeInForce gets TimeInForce, Tag 59
func (m NoMDEntries) GetTimeInForce() (f field.TimeInForceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireDate gets ExpireDate, Tag 432
func (m NoMDEntries) GetExpireDate() (f field.ExpireDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExpireTime gets ExpireTime, Tag 126
func (m NoMDEntries) GetExpireTime() (f field.ExpireTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMinQty gets MinQty, Tag 110
func (m NoMDEntries) GetMinQty() (f field.MinQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetExecInst gets ExecInst, Tag 18
func (m NoMDEntries) GetExecInst() (f field.ExecInstField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSellerDays gets SellerDays, Tag 287
func (m NoMDEntries) GetSellerDays() (f field.SellerDaysField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m NoMDEntries) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuoteEntryID gets QuoteEntryID, Tag 299
func (m NoMDEntries) GetQuoteEntryID() (f field.QuoteEntryIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryBuyer gets MDEntryBuyer, Tag 288
func (m NoMDEntries) GetMDEntryBuyer() (f field.MDEntryBuyerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntrySeller gets MDEntrySeller, Tag 289
func (m NoMDEntries) GetMDEntrySeller() (f field.MDEntrySellerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNumberOfOrders gets NumberOfOrders, Tag 346
func (m NoMDEntries) GetNumberOfOrders() (f field.NumberOfOrdersField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMDEntryPositionNo gets MDEntryPositionNo, Tag 290
func (m NoMDEntries) GetMDEntryPositionNo() (f field.MDEntryPositionNoField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTotalVolumeTraded gets TotalVolumeTraded, Tag 387
func (m NoMDEntries) GetTotalVolumeTraded() (f field.TotalVolumeTradedField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m NoMDEntries) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m NoMDEntries) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m NoMDEntries) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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

//HasIDSource returns true if IDSource is present, Tag 22
func (m NoMDEntries) HasIDSource() bool {
	return m.Has(tag.IDSource)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m NoMDEntries) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m NoMDEntries) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasMaturityDay returns true if MaturityDay is present, Tag 205
func (m NoMDEntries) HasMaturityDay() bool {
	return m.Has(tag.MaturityDay)
}

//HasPutOrCall returns true if PutOrCall is present, Tag 201
func (m NoMDEntries) HasPutOrCall() bool {
	return m.Has(tag.PutOrCall)
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

//HasTotalVolumeTraded returns true if TotalVolumeTraded is present, Tag 387
func (m NoMDEntries) HasTotalVolumeTraded() bool {
	return m.Has(tag.TotalVolumeTraded)
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

//NoMDEntriesRepeatingGroup is a repeating group, Tag 268
type NoMDEntriesRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoMDEntriesRepeatingGroup returns an initialized, NoMDEntriesRepeatingGroup
func NewNoMDEntriesRepeatingGroup() NoMDEntriesRepeatingGroup {
	return NoMDEntriesRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoMDEntries,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.MDUpdateAction), quickfix.GroupElement(tag.DeleteReason), quickfix.GroupElement(tag.MDEntryType), quickfix.GroupElement(tag.MDEntryID), quickfix.GroupElement(tag.MDEntryRefID), quickfix.GroupElement(tag.Symbol), quickfix.GroupElement(tag.SymbolSfx), quickfix.GroupElement(tag.SecurityID), quickfix.GroupElement(tag.IDSource), quickfix.GroupElement(tag.SecurityType), quickfix.GroupElement(tag.MaturityMonthYear), quickfix.GroupElement(tag.MaturityDay), quickfix.GroupElement(tag.PutOrCall), quickfix.GroupElement(tag.StrikePrice), quickfix.GroupElement(tag.OptAttribute), quickfix.GroupElement(tag.ContractMultiplier), quickfix.GroupElement(tag.CouponRate), quickfix.GroupElement(tag.SecurityExchange), quickfix.GroupElement(tag.Issuer), quickfix.GroupElement(tag.EncodedIssuerLen), quickfix.GroupElement(tag.EncodedIssuer), quickfix.GroupElement(tag.SecurityDesc), quickfix.GroupElement(tag.EncodedSecurityDescLen), quickfix.GroupElement(tag.EncodedSecurityDesc), quickfix.GroupElement(tag.FinancialStatus), quickfix.GroupElement(tag.CorporateAction), quickfix.GroupElement(tag.MDEntryPx), quickfix.GroupElement(tag.Currency), quickfix.GroupElement(tag.MDEntrySize), quickfix.GroupElement(tag.MDEntryDate), quickfix.GroupElement(tag.MDEntryTime), quickfix.GroupElement(tag.TickDirection), quickfix.GroupElement(tag.MDMkt), quickfix.GroupElement(tag.TradingSessionID), quickfix.GroupElement(tag.QuoteCondition), quickfix.GroupElement(tag.TradeCondition), quickfix.GroupElement(tag.MDEntryOriginator), quickfix.GroupElement(tag.LocationID), quickfix.GroupElement(tag.DeskID), quickfix.GroupElement(tag.OpenCloseSettleFlag), quickfix.GroupElement(tag.TimeInForce), quickfix.GroupElement(tag.ExpireDate), quickfix.GroupElement(tag.ExpireTime), quickfix.GroupElement(tag.MinQty), quickfix.GroupElement(tag.ExecInst), quickfix.GroupElement(tag.SellerDays), quickfix.GroupElement(tag.OrderID), quickfix.GroupElement(tag.QuoteEntryID), quickfix.GroupElement(tag.MDEntryBuyer), quickfix.GroupElement(tag.MDEntrySeller), quickfix.GroupElement(tag.NumberOfOrders), quickfix.GroupElement(tag.MDEntryPositionNo), quickfix.GroupElement(tag.TotalVolumeTraded), quickfix.GroupElement(tag.Text), quickfix.GroupElement(tag.EncodedTextLen), quickfix.GroupElement(tag.EncodedText)})}
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
