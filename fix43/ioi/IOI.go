package ioi

import (
	"time"

	"github.com/quickfixgo/quickfix"
	"github.com/quickfixgo/quickfix/field"
	"github.com/quickfixgo/quickfix/fix43"
	"github.com/quickfixgo/quickfix/tag"
)

//IOI is the fix43 IOI type, MsgType = 6
type IOI struct {
	fix43.Header
	quickfix.Body
	fix43.Trailer
	//ReceiveTime is the time that this message was read from the socket connection
	ReceiveTime time.Time
}

//FromMessage creates a IOI from a quickfix.Message instance
func FromMessage(m quickfix.Message) IOI {
	return IOI{
		Header:      fix43.Header{Header: m.Header},
		Body:        m.Body,
		Trailer:     fix43.Trailer{Trailer: m.Trailer},
		ReceiveTime: m.ReceiveTime,
	}
}

//ToMessage returns a quickfix.Message instance
func (m IOI) ToMessage() quickfix.Message {
	return quickfix.Message{
		Header:      m.Header.Header,
		Body:        m.Body,
		Trailer:     m.Trailer.Trailer,
		ReceiveTime: m.ReceiveTime,
	}
}

//New returns a IOI initialized with the required fields for IOI
func New(ioiid field.IOIidField, ioitranstype field.IOITransTypeField, side field.SideField, ioiqty field.IOIQtyField) (m IOI) {
	m.Header = fix43.NewHeader()
	m.Init()
	m.Trailer.Init()

	m.Header.Set(field.NewMsgType("6"))
	m.Header.Set(field.NewBeginString("FIX.4.3"))
	m.Set(ioiid)
	m.Set(ioitranstype)
	m.Set(side)
	m.Set(ioiqty)

	return
}

//A RouteOut is the callback type that should be implemented for routing Message
type RouteOut func(msg IOI, sessionID quickfix.SessionID) quickfix.MessageRejectError

//Route returns the beginstring, message type, and MessageRoute for this Message type
func Route(router RouteOut) (string, string, quickfix.MessageRoute) {
	r := func(msg quickfix.Message, sessionID quickfix.SessionID) quickfix.MessageRejectError {
		return router(FromMessage(msg), sessionID)
	}
	return "FIX.4.3", "6", r
}

//SetCurrency sets Currency, Tag 15
func (m IOI) SetCurrency(v string) {
	m.Set(field.NewCurrency(v))
}

//SetSecurityIDSource sets SecurityIDSource, Tag 22
func (m IOI) SetSecurityIDSource(v string) {
	m.Set(field.NewSecurityIDSource(v))
}

//SetIOIid sets IOIid, Tag 23
func (m IOI) SetIOIid(v string) {
	m.Set(field.NewIOIid(v))
}

//SetIOIQltyInd sets IOIQltyInd, Tag 25
func (m IOI) SetIOIQltyInd(v string) {
	m.Set(field.NewIOIQltyInd(v))
}

//SetIOIRefID sets IOIRefID, Tag 26
func (m IOI) SetIOIRefID(v string) {
	m.Set(field.NewIOIRefID(v))
}

//SetIOIQty sets IOIQty, Tag 27
func (m IOI) SetIOIQty(v string) {
	m.Set(field.NewIOIQty(v))
}

//SetIOITransType sets IOITransType, Tag 28
func (m IOI) SetIOITransType(v string) {
	m.Set(field.NewIOITransType(v))
}

//SetPrice sets Price, Tag 44
func (m IOI) SetPrice(v float64) {
	m.Set(field.NewPrice(v))
}

//SetSecurityID sets SecurityID, Tag 48
func (m IOI) SetSecurityID(v string) {
	m.Set(field.NewSecurityID(v))
}

//SetSide sets Side, Tag 54
func (m IOI) SetSide(v string) {
	m.Set(field.NewSide(v))
}

//SetSymbol sets Symbol, Tag 55
func (m IOI) SetSymbol(v string) {
	m.Set(field.NewSymbol(v))
}

//SetText sets Text, Tag 58
func (m IOI) SetText(v string) {
	m.Set(field.NewText(v))
}

//SetTransactTime sets TransactTime, Tag 60
func (m IOI) SetTransactTime(v time.Time) {
	m.Set(field.NewTransactTime(v))
}

//SetValidUntilTime sets ValidUntilTime, Tag 62
func (m IOI) SetValidUntilTime(v time.Time) {
	m.Set(field.NewValidUntilTime(v))
}

//SetSymbolSfx sets SymbolSfx, Tag 65
func (m IOI) SetSymbolSfx(v string) {
	m.Set(field.NewSymbolSfx(v))
}

//SetIssuer sets Issuer, Tag 106
func (m IOI) SetIssuer(v string) {
	m.Set(field.NewIssuer(v))
}

//SetSecurityDesc sets SecurityDesc, Tag 107
func (m IOI) SetSecurityDesc(v string) {
	m.Set(field.NewSecurityDesc(v))
}

//SetIOINaturalFlag sets IOINaturalFlag, Tag 130
func (m IOI) SetIOINaturalFlag(v bool) {
	m.Set(field.NewIOINaturalFlag(v))
}

//SetURLLink sets URLLink, Tag 149
func (m IOI) SetURLLink(v string) {
	m.Set(field.NewURLLink(v))
}

//SetSecurityType sets SecurityType, Tag 167
func (m IOI) SetSecurityType(v string) {
	m.Set(field.NewSecurityType(v))
}

//SetNoIOIQualifiers sets NoIOIQualifiers, Tag 199
func (m IOI) SetNoIOIQualifiers(f NoIOIQualifiersRepeatingGroup) {
	m.SetGroup(f)
}

//SetMaturityMonthYear sets MaturityMonthYear, Tag 200
func (m IOI) SetMaturityMonthYear(v string) {
	m.Set(field.NewMaturityMonthYear(v))
}

//SetStrikePrice sets StrikePrice, Tag 202
func (m IOI) SetStrikePrice(v float64) {
	m.Set(field.NewStrikePrice(v))
}

//SetOptAttribute sets OptAttribute, Tag 206
func (m IOI) SetOptAttribute(v string) {
	m.Set(field.NewOptAttribute(v))
}

//SetSecurityExchange sets SecurityExchange, Tag 207
func (m IOI) SetSecurityExchange(v string) {
	m.Set(field.NewSecurityExchange(v))
}

//SetNoRoutingIDs sets NoRoutingIDs, Tag 215
func (m IOI) SetNoRoutingIDs(f NoRoutingIDsRepeatingGroup) {
	m.SetGroup(f)
}

//SetSpread sets Spread, Tag 218
func (m IOI) SetSpread(v float64) {
	m.Set(field.NewSpread(v))
}

//SetBenchmark sets Benchmark, Tag 219
func (m IOI) SetBenchmark(v string) {
	m.Set(field.NewBenchmark(v))
}

//SetBenchmarkCurveCurrency sets BenchmarkCurveCurrency, Tag 220
func (m IOI) SetBenchmarkCurveCurrency(v string) {
	m.Set(field.NewBenchmarkCurveCurrency(v))
}

//SetBenchmarkCurveName sets BenchmarkCurveName, Tag 221
func (m IOI) SetBenchmarkCurveName(v string) {
	m.Set(field.NewBenchmarkCurveName(v))
}

//SetBenchmarkCurvePoint sets BenchmarkCurvePoint, Tag 222
func (m IOI) SetBenchmarkCurvePoint(v string) {
	m.Set(field.NewBenchmarkCurvePoint(v))
}

//SetCouponRate sets CouponRate, Tag 223
func (m IOI) SetCouponRate(v float64) {
	m.Set(field.NewCouponRate(v))
}

//SetCouponPaymentDate sets CouponPaymentDate, Tag 224
func (m IOI) SetCouponPaymentDate(v string) {
	m.Set(field.NewCouponPaymentDate(v))
}

//SetIssueDate sets IssueDate, Tag 225
func (m IOI) SetIssueDate(v string) {
	m.Set(field.NewIssueDate(v))
}

//SetRepurchaseTerm sets RepurchaseTerm, Tag 226
func (m IOI) SetRepurchaseTerm(v int) {
	m.Set(field.NewRepurchaseTerm(v))
}

//SetRepurchaseRate sets RepurchaseRate, Tag 227
func (m IOI) SetRepurchaseRate(v float64) {
	m.Set(field.NewRepurchaseRate(v))
}

//SetFactor sets Factor, Tag 228
func (m IOI) SetFactor(v float64) {
	m.Set(field.NewFactor(v))
}

//SetContractMultiplier sets ContractMultiplier, Tag 231
func (m IOI) SetContractMultiplier(v float64) {
	m.Set(field.NewContractMultiplier(v))
}

//SetRepoCollateralSecurityType sets RepoCollateralSecurityType, Tag 239
func (m IOI) SetRepoCollateralSecurityType(v int) {
	m.Set(field.NewRepoCollateralSecurityType(v))
}

//SetRedemptionDate sets RedemptionDate, Tag 240
func (m IOI) SetRedemptionDate(v string) {
	m.Set(field.NewRedemptionDate(v))
}

//SetCreditRating sets CreditRating, Tag 255
func (m IOI) SetCreditRating(v string) {
	m.Set(field.NewCreditRating(v))
}

//SetEncodedIssuerLen sets EncodedIssuerLen, Tag 348
func (m IOI) SetEncodedIssuerLen(v int) {
	m.Set(field.NewEncodedIssuerLen(v))
}

//SetEncodedIssuer sets EncodedIssuer, Tag 349
func (m IOI) SetEncodedIssuer(v string) {
	m.Set(field.NewEncodedIssuer(v))
}

//SetEncodedSecurityDescLen sets EncodedSecurityDescLen, Tag 350
func (m IOI) SetEncodedSecurityDescLen(v int) {
	m.Set(field.NewEncodedSecurityDescLen(v))
}

//SetEncodedSecurityDesc sets EncodedSecurityDesc, Tag 351
func (m IOI) SetEncodedSecurityDesc(v string) {
	m.Set(field.NewEncodedSecurityDesc(v))
}

//SetEncodedTextLen sets EncodedTextLen, Tag 354
func (m IOI) SetEncodedTextLen(v int) {
	m.Set(field.NewEncodedTextLen(v))
}

//SetEncodedText sets EncodedText, Tag 355
func (m IOI) SetEncodedText(v string) {
	m.Set(field.NewEncodedText(v))
}

//SetPriceType sets PriceType, Tag 423
func (m IOI) SetPriceType(v int) {
	m.Set(field.NewPriceType(v))
}

//SetNoSecurityAltID sets NoSecurityAltID, Tag 454
func (m IOI) SetNoSecurityAltID(f NoSecurityAltIDRepeatingGroup) {
	m.SetGroup(f)
}

//SetProduct sets Product, Tag 460
func (m IOI) SetProduct(v int) {
	m.Set(field.NewProduct(v))
}

//SetCFICode sets CFICode, Tag 461
func (m IOI) SetCFICode(v string) {
	m.Set(field.NewCFICode(v))
}

//SetQuantityType sets QuantityType, Tag 465
func (m IOI) SetQuantityType(v int) {
	m.Set(field.NewQuantityType(v))
}

//SetCountryOfIssue sets CountryOfIssue, Tag 470
func (m IOI) SetCountryOfIssue(v string) {
	m.Set(field.NewCountryOfIssue(v))
}

//SetStateOrProvinceOfIssue sets StateOrProvinceOfIssue, Tag 471
func (m IOI) SetStateOrProvinceOfIssue(v string) {
	m.Set(field.NewStateOrProvinceOfIssue(v))
}

//SetLocaleOfIssue sets LocaleOfIssue, Tag 472
func (m IOI) SetLocaleOfIssue(v string) {
	m.Set(field.NewLocaleOfIssue(v))
}

//SetMaturityDate sets MaturityDate, Tag 541
func (m IOI) SetMaturityDate(v string) {
	m.Set(field.NewMaturityDate(v))
}

//SetInstrRegistry sets InstrRegistry, Tag 543
func (m IOI) SetInstrRegistry(v string) {
	m.Set(field.NewInstrRegistry(v))
}

//GetCurrency gets Currency, Tag 15
func (m IOI) GetCurrency() (f field.CurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityIDSource gets SecurityIDSource, Tag 22
func (m IOI) GetSecurityIDSource() (f field.SecurityIDSourceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIid gets IOIid, Tag 23
func (m IOI) GetIOIid() (f field.IOIidField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIQltyInd gets IOIQltyInd, Tag 25
func (m IOI) GetIOIQltyInd() (f field.IOIQltyIndField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIRefID gets IOIRefID, Tag 26
func (m IOI) GetIOIRefID() (f field.IOIRefIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOIQty gets IOIQty, Tag 27
func (m IOI) GetIOIQty() (f field.IOIQtyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOITransType gets IOITransType, Tag 28
func (m IOI) GetIOITransType() (f field.IOITransTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPrice gets Price, Tag 44
func (m IOI) GetPrice() (f field.PriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityID gets SecurityID, Tag 48
func (m IOI) GetSecurityID() (f field.SecurityIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSide gets Side, Tag 54
func (m IOI) GetSide() (f field.SideField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbol gets Symbol, Tag 55
func (m IOI) GetSymbol() (f field.SymbolField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetText gets Text, Tag 58
func (m IOI) GetText() (f field.TextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetTransactTime gets TransactTime, Tag 60
func (m IOI) GetTransactTime() (f field.TransactTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetValidUntilTime gets ValidUntilTime, Tag 62
func (m IOI) GetValidUntilTime() (f field.ValidUntilTimeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSymbolSfx gets SymbolSfx, Tag 65
func (m IOI) GetSymbolSfx() (f field.SymbolSfxField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssuer gets Issuer, Tag 106
func (m IOI) GetIssuer() (f field.IssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityDesc gets SecurityDesc, Tag 107
func (m IOI) GetSecurityDesc() (f field.SecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIOINaturalFlag gets IOINaturalFlag, Tag 130
func (m IOI) GetIOINaturalFlag() (f field.IOINaturalFlagField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetURLLink gets URLLink, Tag 149
func (m IOI) GetURLLink() (f field.URLLinkField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityType gets SecurityType, Tag 167
func (m IOI) GetSecurityType() (f field.SecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoIOIQualifiers gets NoIOIQualifiers, Tag 199
func (m IOI) GetNoIOIQualifiers() (NoIOIQualifiersRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoIOIQualifiersRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetMaturityMonthYear gets MaturityMonthYear, Tag 200
func (m IOI) GetMaturityMonthYear() (f field.MaturityMonthYearField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStrikePrice gets StrikePrice, Tag 202
func (m IOI) GetStrikePrice() (f field.StrikePriceField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetOptAttribute gets OptAttribute, Tag 206
func (m IOI) GetOptAttribute() (f field.OptAttributeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetSecurityExchange gets SecurityExchange, Tag 207
func (m IOI) GetSecurityExchange() (f field.SecurityExchangeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoRoutingIDs gets NoRoutingIDs, Tag 215
func (m IOI) GetNoRoutingIDs() (NoRoutingIDsRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoRoutingIDsRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetSpread gets Spread, Tag 218
func (m IOI) GetSpread() (f field.SpreadField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmark gets Benchmark, Tag 219
func (m IOI) GetBenchmark() (f field.BenchmarkField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveCurrency gets BenchmarkCurveCurrency, Tag 220
func (m IOI) GetBenchmarkCurveCurrency() (f field.BenchmarkCurveCurrencyField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurveName gets BenchmarkCurveName, Tag 221
func (m IOI) GetBenchmarkCurveName() (f field.BenchmarkCurveNameField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetBenchmarkCurvePoint gets BenchmarkCurvePoint, Tag 222
func (m IOI) GetBenchmarkCurvePoint() (f field.BenchmarkCurvePointField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponRate gets CouponRate, Tag 223
func (m IOI) GetCouponRate() (f field.CouponRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCouponPaymentDate gets CouponPaymentDate, Tag 224
func (m IOI) GetCouponPaymentDate() (f field.CouponPaymentDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetIssueDate gets IssueDate, Tag 225
func (m IOI) GetIssueDate() (f field.IssueDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseTerm gets RepurchaseTerm, Tag 226
func (m IOI) GetRepurchaseTerm() (f field.RepurchaseTermField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepurchaseRate gets RepurchaseRate, Tag 227
func (m IOI) GetRepurchaseRate() (f field.RepurchaseRateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetFactor gets Factor, Tag 228
func (m IOI) GetFactor() (f field.FactorField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetContractMultiplier gets ContractMultiplier, Tag 231
func (m IOI) GetContractMultiplier() (f field.ContractMultiplierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRepoCollateralSecurityType gets RepoCollateralSecurityType, Tag 239
func (m IOI) GetRepoCollateralSecurityType() (f field.RepoCollateralSecurityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRedemptionDate gets RedemptionDate, Tag 240
func (m IOI) GetRedemptionDate() (f field.RedemptionDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCreditRating gets CreditRating, Tag 255
func (m IOI) GetCreditRating() (f field.CreditRatingField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuerLen gets EncodedIssuerLen, Tag 348
func (m IOI) GetEncodedIssuerLen() (f field.EncodedIssuerLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedIssuer gets EncodedIssuer, Tag 349
func (m IOI) GetEncodedIssuer() (f field.EncodedIssuerField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDescLen gets EncodedSecurityDescLen, Tag 350
func (m IOI) GetEncodedSecurityDescLen() (f field.EncodedSecurityDescLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedSecurityDesc gets EncodedSecurityDesc, Tag 351
func (m IOI) GetEncodedSecurityDesc() (f field.EncodedSecurityDescField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedTextLen gets EncodedTextLen, Tag 354
func (m IOI) GetEncodedTextLen() (f field.EncodedTextLenField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetEncodedText gets EncodedText, Tag 355
func (m IOI) GetEncodedText() (f field.EncodedTextField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetPriceType gets PriceType, Tag 423
func (m IOI) GetPriceType() (f field.PriceTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetNoSecurityAltID gets NoSecurityAltID, Tag 454
func (m IOI) GetNoSecurityAltID() (NoSecurityAltIDRepeatingGroup, quickfix.MessageRejectError) {
	f := NewNoSecurityAltIDRepeatingGroup()
	err := m.GetGroup(f)
	return f, err
}

//GetProduct gets Product, Tag 460
func (m IOI) GetProduct() (f field.ProductField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCFICode gets CFICode, Tag 461
func (m IOI) GetCFICode() (f field.CFICodeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetQuantityType gets QuantityType, Tag 465
func (m IOI) GetQuantityType() (f field.QuantityTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetCountryOfIssue gets CountryOfIssue, Tag 470
func (m IOI) GetCountryOfIssue() (f field.CountryOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetStateOrProvinceOfIssue gets StateOrProvinceOfIssue, Tag 471
func (m IOI) GetStateOrProvinceOfIssue() (f field.StateOrProvinceOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetLocaleOfIssue gets LocaleOfIssue, Tag 472
func (m IOI) GetLocaleOfIssue() (f field.LocaleOfIssueField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetMaturityDate gets MaturityDate, Tag 541
func (m IOI) GetMaturityDate() (f field.MaturityDateField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetInstrRegistry gets InstrRegistry, Tag 543
func (m IOI) GetInstrRegistry() (f field.InstrRegistryField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasCurrency returns true if Currency is present, Tag 15
func (m IOI) HasCurrency() bool {
	return m.Has(tag.Currency)
}

//HasSecurityIDSource returns true if SecurityIDSource is present, Tag 22
func (m IOI) HasSecurityIDSource() bool {
	return m.Has(tag.SecurityIDSource)
}

//HasIOIid returns true if IOIid is present, Tag 23
func (m IOI) HasIOIid() bool {
	return m.Has(tag.IOIid)
}

//HasIOIQltyInd returns true if IOIQltyInd is present, Tag 25
func (m IOI) HasIOIQltyInd() bool {
	return m.Has(tag.IOIQltyInd)
}

//HasIOIRefID returns true if IOIRefID is present, Tag 26
func (m IOI) HasIOIRefID() bool {
	return m.Has(tag.IOIRefID)
}

//HasIOIQty returns true if IOIQty is present, Tag 27
func (m IOI) HasIOIQty() bool {
	return m.Has(tag.IOIQty)
}

//HasIOITransType returns true if IOITransType is present, Tag 28
func (m IOI) HasIOITransType() bool {
	return m.Has(tag.IOITransType)
}

//HasPrice returns true if Price is present, Tag 44
func (m IOI) HasPrice() bool {
	return m.Has(tag.Price)
}

//HasSecurityID returns true if SecurityID is present, Tag 48
func (m IOI) HasSecurityID() bool {
	return m.Has(tag.SecurityID)
}

//HasSide returns true if Side is present, Tag 54
func (m IOI) HasSide() bool {
	return m.Has(tag.Side)
}

//HasSymbol returns true if Symbol is present, Tag 55
func (m IOI) HasSymbol() bool {
	return m.Has(tag.Symbol)
}

//HasText returns true if Text is present, Tag 58
func (m IOI) HasText() bool {
	return m.Has(tag.Text)
}

//HasTransactTime returns true if TransactTime is present, Tag 60
func (m IOI) HasTransactTime() bool {
	return m.Has(tag.TransactTime)
}

//HasValidUntilTime returns true if ValidUntilTime is present, Tag 62
func (m IOI) HasValidUntilTime() bool {
	return m.Has(tag.ValidUntilTime)
}

//HasSymbolSfx returns true if SymbolSfx is present, Tag 65
func (m IOI) HasSymbolSfx() bool {
	return m.Has(tag.SymbolSfx)
}

//HasIssuer returns true if Issuer is present, Tag 106
func (m IOI) HasIssuer() bool {
	return m.Has(tag.Issuer)
}

//HasSecurityDesc returns true if SecurityDesc is present, Tag 107
func (m IOI) HasSecurityDesc() bool {
	return m.Has(tag.SecurityDesc)
}

//HasIOINaturalFlag returns true if IOINaturalFlag is present, Tag 130
func (m IOI) HasIOINaturalFlag() bool {
	return m.Has(tag.IOINaturalFlag)
}

//HasURLLink returns true if URLLink is present, Tag 149
func (m IOI) HasURLLink() bool {
	return m.Has(tag.URLLink)
}

//HasSecurityType returns true if SecurityType is present, Tag 167
func (m IOI) HasSecurityType() bool {
	return m.Has(tag.SecurityType)
}

//HasNoIOIQualifiers returns true if NoIOIQualifiers is present, Tag 199
func (m IOI) HasNoIOIQualifiers() bool {
	return m.Has(tag.NoIOIQualifiers)
}

//HasMaturityMonthYear returns true if MaturityMonthYear is present, Tag 200
func (m IOI) HasMaturityMonthYear() bool {
	return m.Has(tag.MaturityMonthYear)
}

//HasStrikePrice returns true if StrikePrice is present, Tag 202
func (m IOI) HasStrikePrice() bool {
	return m.Has(tag.StrikePrice)
}

//HasOptAttribute returns true if OptAttribute is present, Tag 206
func (m IOI) HasOptAttribute() bool {
	return m.Has(tag.OptAttribute)
}

//HasSecurityExchange returns true if SecurityExchange is present, Tag 207
func (m IOI) HasSecurityExchange() bool {
	return m.Has(tag.SecurityExchange)
}

//HasNoRoutingIDs returns true if NoRoutingIDs is present, Tag 215
func (m IOI) HasNoRoutingIDs() bool {
	return m.Has(tag.NoRoutingIDs)
}

//HasSpread returns true if Spread is present, Tag 218
func (m IOI) HasSpread() bool {
	return m.Has(tag.Spread)
}

//HasBenchmark returns true if Benchmark is present, Tag 219
func (m IOI) HasBenchmark() bool {
	return m.Has(tag.Benchmark)
}

//HasBenchmarkCurveCurrency returns true if BenchmarkCurveCurrency is present, Tag 220
func (m IOI) HasBenchmarkCurveCurrency() bool {
	return m.Has(tag.BenchmarkCurveCurrency)
}

//HasBenchmarkCurveName returns true if BenchmarkCurveName is present, Tag 221
func (m IOI) HasBenchmarkCurveName() bool {
	return m.Has(tag.BenchmarkCurveName)
}

//HasBenchmarkCurvePoint returns true if BenchmarkCurvePoint is present, Tag 222
func (m IOI) HasBenchmarkCurvePoint() bool {
	return m.Has(tag.BenchmarkCurvePoint)
}

//HasCouponRate returns true if CouponRate is present, Tag 223
func (m IOI) HasCouponRate() bool {
	return m.Has(tag.CouponRate)
}

//HasCouponPaymentDate returns true if CouponPaymentDate is present, Tag 224
func (m IOI) HasCouponPaymentDate() bool {
	return m.Has(tag.CouponPaymentDate)
}

//HasIssueDate returns true if IssueDate is present, Tag 225
func (m IOI) HasIssueDate() bool {
	return m.Has(tag.IssueDate)
}

//HasRepurchaseTerm returns true if RepurchaseTerm is present, Tag 226
func (m IOI) HasRepurchaseTerm() bool {
	return m.Has(tag.RepurchaseTerm)
}

//HasRepurchaseRate returns true if RepurchaseRate is present, Tag 227
func (m IOI) HasRepurchaseRate() bool {
	return m.Has(tag.RepurchaseRate)
}

//HasFactor returns true if Factor is present, Tag 228
func (m IOI) HasFactor() bool {
	return m.Has(tag.Factor)
}

//HasContractMultiplier returns true if ContractMultiplier is present, Tag 231
func (m IOI) HasContractMultiplier() bool {
	return m.Has(tag.ContractMultiplier)
}

//HasRepoCollateralSecurityType returns true if RepoCollateralSecurityType is present, Tag 239
func (m IOI) HasRepoCollateralSecurityType() bool {
	return m.Has(tag.RepoCollateralSecurityType)
}

//HasRedemptionDate returns true if RedemptionDate is present, Tag 240
func (m IOI) HasRedemptionDate() bool {
	return m.Has(tag.RedemptionDate)
}

//HasCreditRating returns true if CreditRating is present, Tag 255
func (m IOI) HasCreditRating() bool {
	return m.Has(tag.CreditRating)
}

//HasEncodedIssuerLen returns true if EncodedIssuerLen is present, Tag 348
func (m IOI) HasEncodedIssuerLen() bool {
	return m.Has(tag.EncodedIssuerLen)
}

//HasEncodedIssuer returns true if EncodedIssuer is present, Tag 349
func (m IOI) HasEncodedIssuer() bool {
	return m.Has(tag.EncodedIssuer)
}

//HasEncodedSecurityDescLen returns true if EncodedSecurityDescLen is present, Tag 350
func (m IOI) HasEncodedSecurityDescLen() bool {
	return m.Has(tag.EncodedSecurityDescLen)
}

//HasEncodedSecurityDesc returns true if EncodedSecurityDesc is present, Tag 351
func (m IOI) HasEncodedSecurityDesc() bool {
	return m.Has(tag.EncodedSecurityDesc)
}

//HasEncodedTextLen returns true if EncodedTextLen is present, Tag 354
func (m IOI) HasEncodedTextLen() bool {
	return m.Has(tag.EncodedTextLen)
}

//HasEncodedText returns true if EncodedText is present, Tag 355
func (m IOI) HasEncodedText() bool {
	return m.Has(tag.EncodedText)
}

//HasPriceType returns true if PriceType is present, Tag 423
func (m IOI) HasPriceType() bool {
	return m.Has(tag.PriceType)
}

//HasNoSecurityAltID returns true if NoSecurityAltID is present, Tag 454
func (m IOI) HasNoSecurityAltID() bool {
	return m.Has(tag.NoSecurityAltID)
}

//HasProduct returns true if Product is present, Tag 460
func (m IOI) HasProduct() bool {
	return m.Has(tag.Product)
}

//HasCFICode returns true if CFICode is present, Tag 461
func (m IOI) HasCFICode() bool {
	return m.Has(tag.CFICode)
}

//HasQuantityType returns true if QuantityType is present, Tag 465
func (m IOI) HasQuantityType() bool {
	return m.Has(tag.QuantityType)
}

//HasCountryOfIssue returns true if CountryOfIssue is present, Tag 470
func (m IOI) HasCountryOfIssue() bool {
	return m.Has(tag.CountryOfIssue)
}

//HasStateOrProvinceOfIssue returns true if StateOrProvinceOfIssue is present, Tag 471
func (m IOI) HasStateOrProvinceOfIssue() bool {
	return m.Has(tag.StateOrProvinceOfIssue)
}

//HasLocaleOfIssue returns true if LocaleOfIssue is present, Tag 472
func (m IOI) HasLocaleOfIssue() bool {
	return m.Has(tag.LocaleOfIssue)
}

//HasMaturityDate returns true if MaturityDate is present, Tag 541
func (m IOI) HasMaturityDate() bool {
	return m.Has(tag.MaturityDate)
}

//HasInstrRegistry returns true if InstrRegistry is present, Tag 543
func (m IOI) HasInstrRegistry() bool {
	return m.Has(tag.InstrRegistry)
}

//NoIOIQualifiers is a repeating group element, Tag 199
type NoIOIQualifiers struct {
	quickfix.Group
}

//SetIOIQualifier sets IOIQualifier, Tag 104
func (m NoIOIQualifiers) SetIOIQualifier(v string) {
	m.Set(field.NewIOIQualifier(v))
}

//GetIOIQualifier gets IOIQualifier, Tag 104
func (m NoIOIQualifiers) GetIOIQualifier() (f field.IOIQualifierField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasIOIQualifier returns true if IOIQualifier is present, Tag 104
func (m NoIOIQualifiers) HasIOIQualifier() bool {
	return m.Has(tag.IOIQualifier)
}

//NoIOIQualifiersRepeatingGroup is a repeating group, Tag 199
type NoIOIQualifiersRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoIOIQualifiersRepeatingGroup returns an initialized, NoIOIQualifiersRepeatingGroup
func NewNoIOIQualifiersRepeatingGroup() NoIOIQualifiersRepeatingGroup {
	return NoIOIQualifiersRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoIOIQualifiers,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.IOIQualifier)})}
}

//Add create and append a new NoIOIQualifiers to this group
func (m NoIOIQualifiersRepeatingGroup) Add() NoIOIQualifiers {
	g := m.RepeatingGroup.Add()
	return NoIOIQualifiers{g}
}

//Get returns the ith NoIOIQualifiers in the NoIOIQualifiersRepeatinGroup
func (m NoIOIQualifiersRepeatingGroup) Get(i int) NoIOIQualifiers {
	return NoIOIQualifiers{m.RepeatingGroup.Get(i)}
}

//NoRoutingIDs is a repeating group element, Tag 215
type NoRoutingIDs struct {
	quickfix.Group
}

//SetRoutingType sets RoutingType, Tag 216
func (m NoRoutingIDs) SetRoutingType(v int) {
	m.Set(field.NewRoutingType(v))
}

//SetRoutingID sets RoutingID, Tag 217
func (m NoRoutingIDs) SetRoutingID(v string) {
	m.Set(field.NewRoutingID(v))
}

//GetRoutingType gets RoutingType, Tag 216
func (m NoRoutingIDs) GetRoutingType() (f field.RoutingTypeField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//GetRoutingID gets RoutingID, Tag 217
func (m NoRoutingIDs) GetRoutingID() (f field.RoutingIDField, err quickfix.MessageRejectError) {
	err = m.Get(&f)
	return
}

//HasRoutingType returns true if RoutingType is present, Tag 216
func (m NoRoutingIDs) HasRoutingType() bool {
	return m.Has(tag.RoutingType)
}

//HasRoutingID returns true if RoutingID is present, Tag 217
func (m NoRoutingIDs) HasRoutingID() bool {
	return m.Has(tag.RoutingID)
}

//NoRoutingIDsRepeatingGroup is a repeating group, Tag 215
type NoRoutingIDsRepeatingGroup struct {
	*quickfix.RepeatingGroup
}

//NewNoRoutingIDsRepeatingGroup returns an initialized, NoRoutingIDsRepeatingGroup
func NewNoRoutingIDsRepeatingGroup() NoRoutingIDsRepeatingGroup {
	return NoRoutingIDsRepeatingGroup{
		quickfix.NewRepeatingGroup(tag.NoRoutingIDs,
			quickfix.GroupTemplate{quickfix.GroupElement(tag.RoutingType), quickfix.GroupElement(tag.RoutingID)})}
}

//Add create and append a new NoRoutingIDs to this group
func (m NoRoutingIDsRepeatingGroup) Add() NoRoutingIDs {
	g := m.RepeatingGroup.Add()
	return NoRoutingIDs{g}
}

//Get returns the ith NoRoutingIDs in the NoRoutingIDsRepeatinGroup
func (m NoRoutingIDsRepeatingGroup) Get(i int) NoRoutingIDs {
	return NoRoutingIDs{m.RepeatingGroup.Get(i)}
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
