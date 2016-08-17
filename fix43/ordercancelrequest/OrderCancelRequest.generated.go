package ordercancelrequest

import (
	"github.com/shopspring/decimal"
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//OrderCancelRequest is the fix43 OrderCancelRequest type, MsgType = F
type OrderCancelRequest struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a OrderCancelRequest from a quickfix.Message instance
func FromMessage(m quickfix.Message) OrderCancelRequest {
	return OrderCancelRequest{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m OrderCancelRequest) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a OrderCancelRequest initialized with the required fields for OrderCancelRequest
func New(origclordid field.OrigClOrdIDField, clordid field.ClOrdIDField, side field.SideField, transacttime field.TransactTimeField) (m OrderCancelRequest) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("F"))
	m.Set(origclordid)
	m.Set(clordid)
	m.Set(side)
	m.Set(transacttime)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg OrderCancelRequest, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "F", r
}

//SetAccount sets Account, Tag 1
func (m OrderCancelRequest) SetAccount(v string) {
	m.Set(field.NewAccount(v))
}

//SetClOrdID sets ClOrdID, Tag 11
func (m OrderCancelRequest) SetClOrdID(v string) {
	m.Set(field.NewClOrdID(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m OrderCancelRequest) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetOrderID sets OrderID, Tag 37
func (m OrderCancelRequest) SetOrderID(v string) {
	m.Set(field.NewOrderID(v))
}

//SetOrderQty sets OrderQty, Tag 38
func (m OrderCancelRequest) SetOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderQty(value, scale))
}

//SetOrigClOrdID sets OrigClOrdID, Tag 41
func (m OrderCancelRequest) SetOrigClOrdID(v string) {
	m.Set(field.NewOrigClOrdID(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m OrderCancelRequest) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m OrderCancelRequest) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m OrderCancelRequest) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m OrderCancelRequest) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m OrderCancelRequest) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m OrderCancelRequest) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetListID sets ListID, Tag 66
func (m OrderCancelRequest) SetListID(v string) {
	m.Set(field.NewListID(v))
}

//SetIssuer sets Issuer, Tag 106
func (m OrderCancelRequest) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m OrderCancelRequest) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetCashOrderQty sets CashOrderQty, Tag 152
func (m OrderCancelRequest) SetCashOrderQty(value decimal.Decimal, scale int32) {
	m.Set(field.NewCashOrderQty(value, scale))
}

//SetSecurityType sets SecurityType, Tag 167
func (m OrderCancelRequest) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m OrderCancelRequest) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m OrderCancelRequest) SetStrikePrice(value decimal.Decimal, scale int32) {
	m.Set(field.NewStrikePrice(value, scale))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m OrderCancelRequest) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m OrderCancelRequest) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m OrderCancelRequest) SetCouponRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewCouponRate(value, scale))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m OrderCancelRequest) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m OrderCancelRequest) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m OrderCancelRequest) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m OrderCancelRequest) SetRepurchaseRate(value decimal.Decimal, scale int32) {
	m.Set(field.NewRepurchaseRate(value, scale))
}

//SetFactor sets Factor, Tag 228
func (m OrderCancelRequest) SetFactor(value decimal.Decimal, scale int32) {
	m.Set(field.NewFactor(value, scale))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m OrderCancelRequest) SetContractMultiplier(value decimal.Decimal, scale int32) {
	m.Set(field.NewContractMultiplier(value, scale))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m OrderCancelRequest) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m OrderCancelRequest) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m OrderCancelRequest) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m OrderCancelRequest) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m OrderCancelRequest) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m OrderCancelRequest) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m OrderCancelRequest) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m OrderCancelRequest) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m OrderCancelRequest) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetComplianceID sets ComplianceID, Tag 376
func (m OrderCancelRequest) SetComplianceID(v string) {
	m.Set(field.NewComplianceID(v))
}

//SetNoPartyIDs sets NoPartyIDs, Tag 453
func (m OrderCancelRequest) SetNoPartyIDs(f NoPartyIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m OrderCancelRequest) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m OrderCancelRequest) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m OrderCancelRequest) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetRoundingDirection sets RoundingDirection, Tag 468
func (m OrderCancelRequest) SetRoundingDirection(v string) {
	m.Set(field.NewRoundingDirection(v))
}

//SetRoundingModulus sets RoundingModulus, Tag 469
func (m OrderCancelRequest) SetRoundingModulus(value decimal.Decimal, scale int32) {
	m.Set(field.NewRoundingModulus(value, scale))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m OrderCancelRequest) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelRequest) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m OrderCancelRequest) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetOrderPercent sets OrderPercent, Tag 516
func (m OrderCancelRequest) SetOrderPercent(value decimal.Decimal, scale int32) {
	m.Set(field.NewOrderPercent(value, scale))
}

//SetSecondaryClOrdID sets SecondaryClOrdID, Tag 526
func (m OrderCancelRequest) SetSecondaryClOrdID(v string) {
	m.Set(field.NewSecondaryClOrdID(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m OrderCancelRequest) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m OrderCancelRequest) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//SetAccountType sets AccountType, Tag 581
func (m OrderCancelRequest) SetAccountType(v int) {
	m.Set(field.NewAccountType(v))
}

//SetClOrdLinkID sets ClOrdLinkID, Tag 583
func (m OrderCancelRequest) SetClOrdLinkID(v string) {
	m.Set(field.NewClOrdLinkID(v))
}

//SetOrigOrdModTime sets OrigOrdModTime, Tag 586
func (m OrderCancelRequest) SetOrigOrdModTime(v time.Time) {
	m.Set(field.NewOrigOrdModTime(v))
}

//GetAccount gets Account, Tag 1
func (m OrderCancelRequest) GetAccount() (f field.AccountField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdID gets ClOrdID, Tag 11
func (m OrderCancelRequest) GetClOrdID() (f field.ClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m OrderCancelRequest) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderID gets OrderID, Tag 37
func (m OrderCancelRequest) GetOrderID() (f field.OrderIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderQty gets OrderQty, Tag 38
func (m OrderCancelRequest) GetOrderQty() (f field.OrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigClOrdID gets OrigClOrdID, Tag 41
func (m OrderCancelRequest) GetOrigClOrdID() (f field.OrigClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m OrderCancelRequest) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m OrderCancelRequest) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m OrderCancelRequest) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m OrderCancelRequest) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m OrderCancelRequest) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m OrderCancelRequest) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetListID gets ListID, Tag 66
func (m OrderCancelRequest) GetListID() (f field.ListIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m OrderCancelRequest) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m OrderCancelRequest) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCashOrderQty gets CashOrderQty, Tag 152
func (m OrderCancelRequest) GetCashOrderQty() (f field.CashOrderQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m OrderCancelRequest) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m OrderCancelRequest) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m OrderCancelRequest) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m OrderCancelRequest) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m OrderCancelRequest) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m OrderCancelRequest) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m OrderCancelRequest) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m OrderCancelRequest) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m OrderCancelRequest) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m OrderCancelRequest) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m OrderCancelRequest) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m OrderCancelRequest) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m OrderCancelRequest) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m OrderCancelRequest) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m OrderCancelRequest) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m OrderCancelRequest) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m OrderCancelRequest) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m OrderCancelRequest) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m OrderCancelRequest) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m OrderCancelRequest) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m OrderCancelRequest) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetComplianceID gets ComplianceID, Tag 376
func (m OrderCancelRequest) GetComplianceID() (f field.ComplianceIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoPartyIDs gets NoPartyIDs, Tag 453
func (m OrderCancelRequest) GetNoPartyIDs() (NoPartyIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoPartyIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m OrderCancelRequest) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m OrderCancelRequest) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m OrderCancelRequest) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingDirection gets RoundingDirection, Tag 468
func (m OrderCancelRequest) GetRoundingDirection() (f field.RoundingDirectionField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoundingModulus gets RoundingModulus, Tag 469
func (m OrderCancelRequest) GetRoundingModulus() (f field.RoundingModulusField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m OrderCancelRequest) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m OrderCancelRequest) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m OrderCancelRequest) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrderPercent gets OrderPercent, Tag 516
func (m OrderCancelRequest) GetOrderPercent() (f field.OrderPercentField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecondaryClOrdID gets SecondaryClOrdID, Tag 526
func (m OrderCancelRequest) GetSecondaryClOrdID() (f field.SecondaryClOrdIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m OrderCancelRequest) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m OrderCancelRequest) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetAccountType gets AccountType, Tag 581
func (m OrderCancelRequest) GetAccountType() (f field.AccountTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetClOrdLinkID gets ClOrdLinkID, Tag 583
func (m OrderCancelRequest) GetClOrdLinkID() (f field.ClOrdLinkIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOrigOrdModTime gets OrigOrdModTime, Tag 586
func (m OrderCancelRequest) GetOrigOrdModTime() (f field.OrigOrdModTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasAccount returns true if Account is present, Tag 1
func (m OrderCancelRequest) HasAccount() bool {
	return m.Has(tag.Account)
}

//HasClOrdID returns true if ClOrdID is present, Tag 11
func (m OrderCancelRequest) HasClOrdID() bool {
	return m.Has(tag.ClOrdID)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m OrderCancelRequest) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasOrderID returns true if OrderID is present, Tag 37
func (m OrderCancelRequest) HasOrderID() bool {
	return m.Has(tag.OrderID)
}

//HasOrderQty returns true if OrderQty is present, Tag 38
func (m OrderCancelRequest) HasOrderQty() bool {
	return m.Has(tag.OrderQty)
}

//HasOrigClOrdID returns true if OrigClOrdID is present, Tag 41
func (m OrderCancelRequest) HasOrigClOrdID() bool {
	return m.Has(tag.OrigClOrdID)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m OrderCancelRequest) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m OrderCancelRequest) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m OrderCancelRequest) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m OrderCancelRequest) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m OrderCancelRequest) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m OrderCancelRequest) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasListID returns true if ListID is present, Tag 66
func (m OrderCancelRequest) HasListID() bool {
	return m.Has(tag.ListID)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m OrderCancelRequest) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m OrderCancelRequest) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasCashOrderQty returns true if CashOrderQty is present, Tag 152
func (m OrderCancelRequest) HasCashOrderQty() bool {
	return m.Has(tag.CashOrderQty)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m OrderCancelRequest) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m OrderCancelRequest) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m OrderCancelRequest) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m OrderCancelRequest) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m OrderCancelRequest) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m OrderCancelRequest) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m OrderCancelRequest) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m OrderCancelRequest) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m OrderCancelRequest) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m OrderCancelRequest) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m OrderCancelRequest) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m OrderCancelRequest) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m OrderCancelRequest) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m OrderCancelRequest) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m OrderCancelRequest) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m OrderCancelRequest) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m OrderCancelRequest) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m OrderCancelRequest) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m OrderCancelRequest) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m OrderCancelRequest) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m OrderCancelRequest) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasComplianceID returns true if ComplianceID is present, Tag 376
func (m OrderCancelRequest) HasComplianceID() bool {
	return m.Has(tag.ComplianceID)
}

//HasNoPartyIDs returns true if NoPartyIDs is present, Tag 453
func (m OrderCancelRequest) HasNoPartyIDs() bool {
	return m.Has(tag.NoPartyIDs)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m OrderCancelRequest) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m OrderCancelRequest) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m OrderCancelRequest) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasRoundingDirection returns true if RoundingDirection is present, Tag 468
func (m OrderCancelRequest) HasRoundingDirection() bool {
	return m.Has(tag.RoundingDirection)
}

//HasRoundingModulus returns true if RoundingModulus is present, Tag 469
func (m OrderCancelRequest) HasRoundingModulus() bool {
	return m.Has(tag.RoundingModulus)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m OrderCancelRequest) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m OrderCancelRequest) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m OrderCancelRequest) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasOrderPercent returns true if OrderPercent is present, Tag 516
func (m OrderCancelRequest) HasOrderPercent() bool {
	return m.Has(tag.OrderPercent)
}

//HasSecondaryClOrdID returns true if SecondaryClOrdID is present, Tag 526
func (m OrderCancelRequest) HasSecondaryClOrdID() bool {
	return m.Has(tag.SecondaryClOrdID)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m OrderCancelRequest) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m OrderCancelRequest) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//HasAccountType returns true if AccountType is present, Tag 581
func (m OrderCancelRequest) HasAccountType() bool {
	return m.Has(tag.AccountType)
}

//HasClOrdLinkID returns true if ClOrdLinkID is present, Tag 583
func (m OrderCancelRequest) HasClOrdLinkID() bool {
	return m.Has(tag.ClOrdLinkID)
}

//HasOrigOrdModTime returns true if OrigOrdModTime is present, Tag 586
func (m OrderCancelRequest) HasOrigOrdModTime() bool {
	return m.Has(tag.OrigOrdModTime)
}

//NoPartyIDs is a repeating group element, Tag 453
type NoPartyIDs struct {
	quickfix.Group
}

//SetPartyID sets PartyID, Tag 448
func (m NoPartyIDs) SetPartyID(v string) {
	m.Set(field.NewPartyID(v))
}

//SetPartyIDSource sets PartyIDSource, Tag 447
func (m NoPartyIDs) SetPartyIDSource(v string) {
	m.Set(field.NewPartyIDSource(v))
}

//SetPartyRole sets PartyRole, Tag 452
func (m NoPartyIDs) SetPartyRole(v int) {
	m.Set(field.NewPartyRole(v))
}

//SetPartySubID sets PartySubID, Tag 523
func (m NoPartyIDs) SetPartySubID(v string) {
	m.Set(field.NewPartySubID(v))
}

//GetPartyID gets PartyID, Tag 448
func (m NoPartyIDs) GetPartyID() (f field.PartyIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyIDSource gets PartyIDSource, Tag 447
func (m NoPartyIDs) GetPartyIDSource() (f field.PartyIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartyRole gets PartyRole, Tag 452
func (m NoPartyIDs) GetPartyRole() (f field.PartyRoleField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPartySubID gets PartySubID, Tag 523
func (m NoPartyIDs) GetPartySubID() (f field.PartySubIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasPartyID returns true if PartyID is present, Tag 448
func (m NoPartyIDs) HasPartyID() bool {
	return m.Has(tag.PartyID)
}

//HasPartyIDSource returns true if PartyIDSource is present, Tag 447
func (m NoPartyIDs) HasPartyIDSource() bool {
	return m.Has(tag.PartyIDSource)
}

//HasPartyRole returns true if PartyRole is present, Tag 452
func (m NoPartyIDs) HasPartyRole() bool {
	return m.Has(tag.PartyRole)
}

//HasPartySubID returns true if PartySubID is present, Tag 523
func (m NoPartyIDs) HasPartySubID() bool {
	return m.Has(tag.PartySubID)
}

//NoPartyIDsRepeatingGroup is a repeating group, Tag 453
type NoPartyIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoPartyIDsRepeatingGroup returns an initialized, NoPartyIDsRepeatingGroup
func NewNoPartyIDsRepeatingGroup() NoPartyIDsRepeatingGroup {
	return NoPartyIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoPartyIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.PartyID), quickfix.GroupElement(tag.PartyIDSource), quickfix.GroupElement(tag.PartyRole), quickfix.GroupElement(tag.PartySubID)})}
}

//Add create and append a new NoPartyIDs to this group
func (m NoPartyIDsRepeatingGroup) Add() NoPartyIDs {
	g := m.RepeatingGroup.Add()
	return NoPartyIDs{g}
}

//Get returns the ith NoPartyIDs in the NoPartyIDsRepeatinGroup
func (m NoPartyIDsRepeatingGroup) Get(i int) NoPartyIDs {
	return NoPartyIDs{m.RepeatingGroup.Get(i)}
}

//NoSecurityAltID is a repeating group element, Tag 454
type NoSecurityAltID struct {
	quickfix.Group
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
func (m NoSecurityAltID) GetSecurityAltID() (f field.SecurityAltIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityAltIDSource gets SecurityAltIDSource, Tag 456
func (m NoSecurityAltID) GetSecurityAltIDSource() (f field.SecurityAltIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
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
