package fix50

import (
	"github.com/cbusbey/quickfixgo"
	"github.com/cbusbey/quickfixgo/field"
)

type QuoteResponse struct {
	quickfixgo.Message
}

func (m *QuoteResponse) QuoteRespID() (*field.QuoteRespID, error) {
	f := new(field.QuoteRespID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) QuoteID() (*field.QuoteID, error) {
	f := new(field.QuoteID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) QuoteRespType() (*field.QuoteRespType, error) {
	f := new(field.QuoteRespType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) ClOrdID() (*field.ClOrdID, error) {
	f := new(field.ClOrdID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OrderCapacity() (*field.OrderCapacity, error) {
	f := new(field.OrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) IOIID() (*field.IOIID, error) {
	f := new(field.IOIID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) QuoteType() (*field.QuoteType, error) {
	f := new(field.QuoteType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) TradingSessionID() (*field.TradingSessionID, error) {
	f := new(field.TradingSessionID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) TradingSessionSubID() (*field.TradingSessionSubID, error) {
	f := new(field.TradingSessionSubID)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Side() (*field.Side, error) {
	f := new(field.Side)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlType() (*field.SettlType, error) {
	f := new(field.SettlType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlDate() (*field.SettlDate, error) {
	f := new(field.SettlDate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlDate2() (*field.SettlDate2, error) {
	f := new(field.SettlDate2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OrderQty2() (*field.OrderQty2, error) {
	f := new(field.OrderQty2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Currency() (*field.Currency, error) {
	f := new(field.Currency)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Account() (*field.Account, error) {
	f := new(field.Account)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) AcctIDSource() (*field.AcctIDSource, error) {
	f := new(field.AcctIDSource)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) AccountType() (*field.AccountType, error) {
	f := new(field.AccountType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidPx() (*field.BidPx, error) {
	f := new(field.BidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferPx() (*field.OfferPx, error) {
	f := new(field.OfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MktBidPx() (*field.MktBidPx, error) {
	f := new(field.MktBidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MktOfferPx() (*field.MktOfferPx, error) {
	f := new(field.MktOfferPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MinBidSize() (*field.MinBidSize, error) {
	f := new(field.MinBidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidSize() (*field.BidSize, error) {
	f := new(field.BidSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MinOfferSize() (*field.MinOfferSize, error) {
	f := new(field.MinOfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferSize() (*field.OfferSize, error) {
	f := new(field.OfferSize)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) ValidUntilTime() (*field.ValidUntilTime, error) {
	f := new(field.ValidUntilTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidSpotRate() (*field.BidSpotRate, error) {
	f := new(field.BidSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferSpotRate() (*field.OfferSpotRate, error) {
	f := new(field.OfferSpotRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidForwardPoints() (*field.BidForwardPoints, error) {
	f := new(field.BidForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferForwardPoints() (*field.OfferForwardPoints, error) {
	f := new(field.OfferForwardPoints)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MidPx() (*field.MidPx, error) {
	f := new(field.MidPx)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidYield() (*field.BidYield, error) {
	f := new(field.BidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) MidYield() (*field.MidYield, error) {
	f := new(field.MidYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferYield() (*field.OfferYield, error) {
	f := new(field.OfferYield)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) TransactTime() (*field.TransactTime, error) {
	f := new(field.TransactTime)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OrdType() (*field.OrdType, error) {
	f := new(field.OrdType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) BidForwardPoints2() (*field.BidForwardPoints2, error) {
	f := new(field.BidForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) OfferForwardPoints2() (*field.OfferForwardPoints2, error) {
	f := new(field.OfferForwardPoints2)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlCurrBidFxRate() (*field.SettlCurrBidFxRate, error) {
	f := new(field.SettlCurrBidFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlCurrOfferFxRate() (*field.SettlCurrOfferFxRate, error) {
	f := new(field.SettlCurrOfferFxRate)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) SettlCurrFxRateCalc() (*field.SettlCurrFxRateCalc, error) {
	f := new(field.SettlCurrFxRateCalc)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Commission() (*field.Commission, error) {
	f := new(field.Commission)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) CommType() (*field.CommType, error) {
	f := new(field.CommType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) CustOrderCapacity() (*field.CustOrderCapacity, error) {
	f := new(field.CustOrderCapacity)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) ExDestination() (*field.ExDestination, error) {
	f := new(field.ExDestination)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Text() (*field.Text, error) {
	f := new(field.Text)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) EncodedTextLen() (*field.EncodedTextLen, error) {
	f := new(field.EncodedTextLen)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) EncodedText() (*field.EncodedText, error) {
	f := new(field.EncodedText)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) Price() (*field.Price, error) {
	f := new(field.Price)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) PriceType() (*field.PriceType, error) {
	f := new(field.PriceType)
	err := m.Body.Get(f)
	return f, err
}
func (m *QuoteResponse) ExDestinationIDSource() (*field.ExDestinationIDSource, error) {
	f := new(field.ExDestinationIDSource)
	err := m.Body.Get(f)
	return f, err
}
