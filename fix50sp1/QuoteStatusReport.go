package fix50sp1

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteStatusReport struct {
	quickfixgo.Message
}

func (m *QuoteStatusReport) QuoteStatusReqID() (*field.QuoteStatusReqID, error) {
	f := new(field.QuoteStatusReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteReqID() (*field.QuoteReqID, error) {
	f := new(field.QuoteReqID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteRespID() (*field.QuoteRespID, error) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ExpireTime() (*field.ExpireTime, error) {
	f := new(field.ExpireTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidPx() (*field.BidPx, error) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferPx() (*field.OfferPx, error) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MktBidPx() (*field.MktBidPx, error) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MktOfferPx() (*field.MktOfferPx, error) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MinBidSize() (*field.MinBidSize, error) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidSize() (*field.BidSize, error) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MinOfferSize() (*field.MinOfferSize, error) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferSize() (*field.OfferSize, error) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidSpotRate() (*field.BidSpotRate, error) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferSpotRate() (*field.OfferSpotRate, error) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidForwardPoints() (*field.BidForwardPoints, error) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferForwardPoints() (*field.OfferForwardPoints, error) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MidPx() (*field.MidPx, error) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidYield() (*field.BidYield, error) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MidYield() (*field.MidYield, error) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferYield() (*field.OfferYield, error) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) BidForwardPoints2() (*field.BidForwardPoints2, error) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) OfferForwardPoints2() (*field.OfferForwardPoints2, error) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, error) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, error) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteStatus() (*field.QuoteStatus, error) {
	f := new(field.QuoteStatus)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteCancelType() (*field.QuoteCancelType, error) {
	f := new(field.QuoteCancelType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteMsgID() (*field.QuoteMsgID, error) {
	f := new(field.QuoteMsgID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) QuoteRejectReason() (*field.QuoteRejectReason, error) {
	f := new(field.QuoteRejectReason)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteStatusReport) MinQty() (*field.MinQty, error) {
	f := new(field.MinQty)
	err := m.Body.Get(f)
	return f, err
}
